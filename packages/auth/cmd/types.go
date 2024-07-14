package cmd

type User struct {
	ID       uint   `json:"id"       gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email"    gorm:"unique"`
	Password string `json:"password"`
}
