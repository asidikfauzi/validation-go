package domain

type Users struct {
	ID       int    `gorm:"type:int;primary_key;auto_increments" json:"id"`
	Username string `gorm:"type:varchar(50);not null" json:"username"`
	Email    string `gorm:"type:varchar(50);not null" json:"email"`
	Password string `gorm:"type:text;not null" json:"password"`
}
