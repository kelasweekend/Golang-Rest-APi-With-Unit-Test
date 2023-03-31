package Todo

type Todo struct {
	Id         uint   `gorm:"primaryKey" json:"id"`
	Name       string `json:"name"`
	Status     string `json:"status"`
	Priority   string `json:"priority"`
	ActivityID uint
	Activity   Activity `gorm:"foreignKey:ActivityID:references:Id:constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Updated    int64    `gorm:"autoUpdateTime:milli"`
	Created    int64    `gorm:"autoCreateTime"`
}

type Activity struct {
	Id    uint
	Title string
}

func (b *Todo) TableName() string {
	return "todo"
}
