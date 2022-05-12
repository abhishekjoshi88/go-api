package repository

import (
	"cleanarch/domain"
	"context"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

// StickerConstructor will create an object that represent the StickerRepository interface
func StickerConstructor(db *gorm.DB) domain.StickerRepository {
	return &repository{db: db}
}

//GetTrendingStickers will fetch trending stickers from database
func (repository *repository) GetTrendingStickers(currentTime string) ([]*domain.Sticker, error) {

	trendingStickers := make([]*domain.Sticker, 0)
	err := repository.db.WithContext(context.Background()).Table("stickers").
		Where("stickers.startTime <= ?", currentTime).
		Where("stickers.endTime >= ?", currentTime).
		Order("priority desc").
		Find(&trendingStickers).Error
	if err != nil {

		err = domain.UnexpectedError
	}
	return trendingStickers, err
}

//GetStickerPackStickers will fetch pack stickers from database
func (repository *repository) GetStickerPackStickers(id int) ([]*domain.Sticker, error) {

	packStickers := make([]*domain.Sticker, 0)
	err := repository.db.WithContext(context.Background()).Table("stickers").
		Where("stickers.packid = ?", id).
		Find(&packStickers).Error

	return packStickers, err
}
