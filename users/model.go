package users

type User struct {
	ID       int    `bson:"_id"`
	FullName string `json:"fullname"`
	Age      int    `json:"age"`
	Phone    int32  `json:"phone"`
	Address  string `json:"address"`
}

type Users []User

const (
	HOST     = "localhost"
	PORT     = 5432
	USER     = "postgres"
	PASSWORD = "YOUR_PASSWORD"
	DBNAME   = "YOUR_DATABASE"
)