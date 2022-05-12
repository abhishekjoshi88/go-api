package usecase

import (
	"cleanarch/config"
	"cleanarch/domain"

	"time"
)

// Usecase represents the uscase for stickers
type useCase struct {
	repository domain.StickerRepository
}

// StickerConstructor will create an object that represent the domain.StickerUsecase interface
func StickerConstructor(repository domain.StickerRepository) domain.StickerUseCase {
	return &useCase{repository: repository}
}

//GetTrendingStickers func
func (useCase useCase) GetTrendingStickers() (result []*domain.Sticker, err error) {
	now := time.Now()
	currentTime := now.Format(config.TimeFormat)

	result, err = useCase.repository.GetTrendingStickers(currentTime)

	return
}

//GetStickerPackStickers func
func (useCase useCase) GetStickerPackStickers(id int) (resultPack []*domain.Sticker, err error) {

	resultPack, err = useCase.repository.GetStickerPackStickers(id)

	return
}
