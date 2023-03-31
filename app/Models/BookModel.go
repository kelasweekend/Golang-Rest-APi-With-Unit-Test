package Models

type Book struct {
	Id          uint    `gorm:"primaryKey" json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Updated     int64   `gorm:"autoUpdateTime:milli"`
	Created     int64   `gorm:"autoCreateTime"`
}

func (b *Book) TableName() string {
	return "book"
}
