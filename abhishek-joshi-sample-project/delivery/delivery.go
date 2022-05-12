package delivery

import (
	"cleanarch/domain"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

//Delivery represents the http handlers for stickers
type Delivery struct {
	UseCase domain.StickerUseCase
}

// StickerConstructor will initialize the endpoints and their respective handlers
func StickerConstructor(e *echo.Echo, useCase domain.StickerUseCase) {
	handler := &Delivery{
		UseCase: useCase,
	}

	e.GET("/v1/trendingStickers", handler.GetTrendingStickers)
	e.GET("/stickerPacks/:id", handler.GetPackStickers)
}

//GetTrendingStickers will fetch trending stickers
func (delivery *Delivery) GetTrendingStickers(ctx echo.Context) error {
	sticker, err := delivery.UseCase.GetTrendingStickers()
	if err != nil {
		if err.Error() == "databaseError" {
			return ctx.JSON(http.StatusInternalServerError, domain.ErrDB)
		}
		return ctx.JSON(http.StatusInternalServerError, domain.UnexpectedError)

	}
	return ctx.JSON(http.StatusOK, sticker)
}

//GetPackStickers will fetch pack stickers
func (delivery *Delivery) GetPackStickers(ctx echo.Context) error {

	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil || (id == 0) {
		return ctx.JSON(http.StatusBadRequest, domain.ErrInvalidID)
	}
	stickerPackStickers, err := delivery.UseCase.GetStickerPackStickers(id)
	if err != nil {
		if err.Error() == "databaseError" {
			return ctx.JSON(http.StatusNotFound, domain.ErrDB)
		}
		return ctx.JSON(http.StatusInternalServerError, domain.UnexpectedError)
	}
	return ctx.JSON(http.StatusOK, stickerPackStickers)
}
