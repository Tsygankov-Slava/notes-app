package notes

type User struct {
	Id       int    `json:"int" db:"id"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
