package app

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/Maksim646/ozon_link_shorter/internal/config"
	"github.com/Maksim646/ozon_link_shorter/internal/database"
	"github.com/Maksim646/ozon_link_shorter/internal/database/postgresql"
	grpcapp "github.com/Maksim646/ozon_link_shorter/internal/grpc_app"
	"github.com/Maksim646/ozon_link_shorter/internal/model"
	"github.com/heetch/sqalx"
	"github.com/jmoiron/sqlx"

	_originalLinkRepo "github.com/Maksim646/ozon_link_shorter/internal/domain/original_link/repository/postgresql"
	_originalLinkUsecase "github.com/Maksim646/ozon_link_shorter/internal/domain/original_link/usecase"
)

type App struct {
	GRPCSrv *grpcapp.App
}

type Usecases struct {
	OriginalLinkUsecase model.IOriginalLinkUsecase
	// Добавьте другие юзкейсы здесь
	// UserUsecase         _userUsecase.Usecase
	// AdminUsecase        _adminUsecase.Usecase
	// ...
}

func New(log *slog.Logger, cfg *config.Config) (*App, func(), error) { // take config as param
	// Initialize repository
	usecases, dbConn, cleanup, err := initRepository(log, cfg) // Изменено
	if err != nil {
		return nil, nil, fmt.Errorf("failed to initialize repository: %w", err)
	}

	// Initialize gRPC server
	grpcServer := grpcapp.New(log, cfg.GRPC.Prot, usecases)

	return &App{
			GRPCSrv: grpcServer,
		}, func() {
			cleanup() // Call cleanup to close resources
			if dbConn != nil {
				if err := dbConn.Close(); err != nil {
					log.Error("failed to close db connection", slog.Any("error", err))
				}
			}
		}, nil
}

func initRepository(log *slog.Logger, cfg *config.Config) (Usecases, *sqlx.DB, func(), error) { // Изменено
	switch cfg.Data {
	case "postgresql":
		log.Info("initializing postgresql database")
		dbConn, err := sqlx.Connect("postgres", cfg.PostgresURI)
		if err != nil {
			return Usecases{}, nil, nil, fmt.Errorf("failed to connect to postgres: %w", err) // Изменено
		}

		dbConn.SetMaxOpenConns(100)
		dbConn.SetMaxIdleConns(100)
		dbConn.SetConnMaxLifetime(5 * time.Minute)

		sqalxConn, err := sqalx.New(dbConn)
		if err != nil {
			return Usecases{}, nil, nil, fmt.Errorf("failed to create sqalx connection: %w", err) // Изменено
		}

		migrator := postgresql.NewMigrator(cfg.PostgresURI, cfg.MigrationDir)
		if err := migrator.Apply(); err != nil {
			return Usecases{}, nil, nil, fmt.Errorf("failed to apply migrations: %w", err) // Изменено
		}

		originalLinkRepo := _originalLinkRepo.New(sqalxConn)
		originalLinkUsecase := _originalLinkUsecase.New(originalLinkRepo)

		// Инициализируйте другие репозитории и юзкейсы здесь
		// userRepo := _userRepo.New(sqalxConn)
		// userUsecase := _userUsecase.New(userRepo)
		// ...

		return Usecases{ // Изменено
				OriginalLinkUsecase: originalLinkUsecase,
				// Заполните другие юзкейсы здесь
				// UserUsecase: userUsecase,
				// ...
			}, dbConn, func() {
				sqalxConn.Close()
				if err := dbConn.Close(); err != nil {
					log.Error("failed to close db connection", slog.Any("error", err))
				}
			}, nil
	case "inmemory":
		log.Info("initializing in-memory storage")
		repo := database.NewInMemoryRepository()
		inMemoryUsecase := inmemory.NewInMemoryUsecase(repo)

		// Инициализируйте другие in-memory юзкейсы и репозитории здесь

		return Usecases{ // Изменено
			OriginalLinkUsecase: inMemoryUsecase,
			// Заполните другие in-memory юзкейсы здесь
			// UserUsecase: inMemoryUserUsecase,
			// ...
		}, nil, func() {}, nil
	default:
		return Usecases{}, nil, nil, fmt.Errorf("invalid data source: %s", cfg.Data) // Изменено
	}
}
