package link

import (
	"context"

	linkv2 "github.com/Maksim646/protos/gen/go/link"
	// "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc"
)

type serverAPI struct {
	linkv2.UnimplementedLinkServer
}

func Register(gRPC *grpc.Server) {
	linkv2.RegisterLinkServer(gRPC, &serverAPI{})
}

func (s *serverAPI) CreateShorterLink(ctx context.Context, req *linkv2.CreateShorterLinkRequest) (*linkv2.CreateShorterLinkResponse, error) {
	// if req == nil {
	// 	return nil, status.Error(codes.InvalidArgument, "request is nil")
	// }
	// originalURL := req.GetOriginalLink()
	// if originalURL == "" {
	// 	return nil, status.Error(codes.InvalidArgument, "original_url is required")
	// }

	// // 2. Генерация сокращённой ссылки
	// shortenedLink, err := s.generateShortLink(originalURL) // Предполагаем, что s - это структура с методом generateShortLink
	// if err != nil {
	// 	// Логирование ошибки
	// 	s.log.Error("failed to generate short link", slog.Error(err)) // Используем логгер s.log (предполагается, что он у вас есть)
	// 	return nil, status.Error(codes.Internal, "failed to generate short link")
	// }

	// // 3. Сохранение оригинальной и сокращённой ссылок в базе данных
	// err = s.linkService.SaveLink(ctx, originalURL, shortenedLink) // Предполагаем, что s.linkService - это объект, реализующий интерфейс для работы с БД
	// if err != nil {
	// 	// Логирование ошибки
	// 	slog.Error("failed to save link", slog.Error(err))
	// 	return nil, status.Error(codes.Internal, "failed to save link")
	// }

	// // 4. Формирование ответа
	// resp := &linkv2.CreateShorterLinkResponse{
	// 	ShorterLink: shortenedLink,
	// }

	// // 5. Логирование успешного создания
	// s.log.Info("link created", sl.String("original_url", originalURL), sl.String("shortened_url", shortenedLink))

	return resp, nil
}

func (s *serverAPI) GetOriginalLink(ctx context.Context, req *linkv2.GetOriginalLinkRequest) (*linkv2.GetOriginalLinkResponse, error) {

	panic("implement me")
}
