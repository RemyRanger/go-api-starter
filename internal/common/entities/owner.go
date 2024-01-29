package entities

type Owner struct {
	ID   string `gorm:"size:255;column:uuid"`
	Code string `gorm:"size:255;column:code"`
}
