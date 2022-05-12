package usecase_test

import (
	"cleanarch/domain"
	"cleanarch/mocks"
	"testing"
	"time"

	usecase "cleanarch/usecase"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

//TestGetTrendingStickers testing
func TestGetTrendingStickers(t *testing.T) {
	mockStickerRepo := new(mocks.StickerRepository)
	mockSticker := &domain.Sticker{
		ID: 1, Name: "name1", StartTime: "11:00:00",
		EndTime: "20:00:00", CreatedAt: time.Now(), UpdatedAt: time.Now(),
	}

	mockListSticker := make([]*domain.Sticker, 0)
	mockListSticker = append(mockListSticker, mockSticker)

	t.Run("success", func(t *testing.T) {
		mockStickerRepo.On("GetTrendingStickers", mock.AnythingOfType("string")).Return(mockListSticker, nil).Once()

		u := usecase.StickerConstructor(mockStickerRepo)

		list, err := u.GetTrendingStickers()
		assert.NoError(t, err)
		assert.Len(t, list, len(mockListSticker))

		mockStickerRepo.AssertExpectations(t)

	})

}

//TestGetStickerPackStickersDetails testing
func TestGetStickerPackStickersDetails(t *testing.T) {
	mockStickerRepo := new(mocks.StickerRepository)
	mockStickerPack := &domain.Sticker{
		ID: 1, Name: "A1", Priority: 1, Packid: 1,
	}

	mockListSticker := make([]*domain.Sticker, 0)
	mockListSticker = append(mockListSticker, mockStickerPack)

	t.Run("success", func(t *testing.T) {
		mockStickerRepo.On("GetStickerPackStickers", mock.AnythingOfType("int")).Return(mockListSticker, nil).Once()

		u := usecase.StickerConstructor(mockStickerRepo)

		list, err := u.GetStickerPackStickers(mockStickerPack.Packid)

		assert.NoError(t, err)
		assert.Len(t, list, len(mockListSticker))

		mockStickerRepo.AssertExpectations(t)

	})

}
