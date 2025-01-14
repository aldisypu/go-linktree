package entity

type Link struct {
	ID        string `gorm:"column:id;primaryKey"`
	Title     string `gorm:"column:title"`
	Url       string `gorm:"column:url"`
	Username  string `gorm:"column:username"`
	CreatedAt int64  `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt int64  `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
	User      User   `gorm:"foreignKey:username;references:username"`
}

func (r *Link) TableName() string {
	return "links"
}
