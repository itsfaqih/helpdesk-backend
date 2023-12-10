package permission

type Permission struct {
	Code        string `json:"code" gorm:"primaryKey;type:varchar"`
	Name        string `json:"name" gorm:"not null;type:varchar"`
	Description string `json:"description" gorm:"not null;type:text;default:''"`
}
