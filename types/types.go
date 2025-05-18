package types

type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

type User struct {
	Id       string `json:"-" db:"id"`
	Email    string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignInData struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type NewUserData struct {
	Email string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
