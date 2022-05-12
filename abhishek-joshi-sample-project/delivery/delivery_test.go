package delivery_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	StickerHttp "cleanarch/delivery"
	"cleanarch/domain"
	mocks "cleanarch/mocks"
	_repo "cleanarch/repository"
)

// Test code for Sticker Handler
func TestGetTrendingStickers(t *testing.T) {

	mockSticker := &domain.Sticker{
		ID: 1, Name: "name1", StartTime: "11:00:00",
		EndTime: "20:00:00", CreatedAt: time.Now(), UpdatedAt: time.Now(),
	}

	mockUCase := new(mocks.StickerUseCase)
	mockListSticker := make([]*domain.Sticker, 0)
	mockListSticker = append(mockListSticker, mockSticker)

	mockUCase.On("GetTrendingStickers").Return(mockListSticker, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/v1/trendingsticker", strings.NewReader(""))
	//err = errors.New("hello")
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := StickerHttp.Delivery{
		UseCase: mockUCase,
	}
	err = handler.GetTrendingStickers(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}

//TestGetStickerPackStickersDetails testing
func TestGetStickerPackStickersDetails(t *testing.T) {
	DB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	dialector := mysql.New(mysql.Config{
		DSN:                       "sqlmock_db_0",
		DriverName:                "mysql",
		Conn:                      DB,
		SkipInitializeWithVersion: true,
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})
	defer DB.Close()
	mockStickerPack := domain.Sticker{
		ID: 1, Name: "A1", Priority: 1, Packid: 1,
	}
	rows := sqlmock.NewRows([]string{"id", "name", "priority", "packid"}).
		AddRow(mockStickerPack.ID, mockStickerPack.Name, mockStickerPack.Priority, mockStickerPack.Packid)
	query := "[select id, name, priority, packid from stickers where id = 1]"
	mock.ExpectQuery(query).WillReturnRows(rows)
	a := _repo.StickerConstructor(db)
	pack, err := a.GetStickerPackStickers(0)
	assert.NoError(t, err)
	assert.Len(t, pack, 1)

}
