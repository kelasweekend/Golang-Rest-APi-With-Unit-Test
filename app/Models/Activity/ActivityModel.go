package Activity

import "time"

type Activity struct {
	Id        uint   `gorm:"primaryKey" json:"id"`
	Title     string `json:"title"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (b *Activity) TableName() string {
	return "activity"
}
