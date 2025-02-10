package usecase

import (
	"context"

	"github.com/Maksim646/ozon_link_shorter/internal/model"
)

type Usecase struct {
	originalLinkRepository model.IOriginalLinkRepository
}

func New(originalLinkRepository model.IOriginalLinkRepository) model.IOriginalLinkUsecase {
	return &Usecase{
		originalLinkRepository: originalLinkRepository,
	}
}

func (u *Usecase) CreateOriginalLink(ctx context.Context, originalLink string, shorterLinkID int64) error {
	return u.originalLinkRepository.CreateOriginalLink(ctx, originalLink, shorterLinkID)
}

func (u *Usecase) GetOriginalLinkByShorterLinkID(ctx context.Context, shorterLinkID int64) (string, error) {
	return u.originalLinkRepository.GetOriginalLinkByShorterLinkID(ctx, shorterLinkID)
}
