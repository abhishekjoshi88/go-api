package domain

import "time"

//StickersTable Table Names
const (
	StickersTable = "stickers"
)

//Sticker tables Structs
type Sticker struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Priority  int       `json:"priority"`
	Packid    int       `json:"packId"`
	StartTime string    `json:"-"`
	EndTime   string    `json:"-"`
	UpdatedAt time.Time `json:"-"`
	CreatedAt time.Time `json:"-"`
}

//TableName table name
func (sticker *Sticker) TableName() string {
	return StickersTable
}

// StickerUseCase interface Structs
type StickerUseCase interface {
	GetTrendingStickers() ([]*Sticker, error)
	GetStickerPackStickers(id int) ([]*Sticker, error)
}

// StickerRepository interface Structs
type StickerRepository interface {
	GetTrendingStickers(currentTime string) ([]*Sticker, error)
	GetStickerPackStickers(id int) ([]*Sticker, error)
}
