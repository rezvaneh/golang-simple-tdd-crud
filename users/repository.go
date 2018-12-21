package users

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Repository struct{}

var db *sql.DB

func (r *Repository) InitDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", HOST, PORT, USER, PASSWORD, DBNAME)
	dbp, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	db = dbp

	return db, nil
}

func (r *Repository) connectDB(databse *sql.DB) {
	db = databse

	return
}

func (r *Repository) GetUsers() (results Users) {
	rows, err := db.Query("SELECT ID, FullName, Age, Phone, Address FROM users")
	checkErr(err)
	var row User

	for rows.Next() {
		err = rows.Scan(&row.ID, &row.FullName, &row.Age, &row.Phone, &row.Address)
		checkErr(err)
		results = append(results, row)
	}

	return results
}

func (r *Repository) GetUserById(id int) User {
	var row User
	err := db.QueryRow("SELECT ID, FullName, Age, Phone, Address FROM users WHERE Id=$1;", id).Scan(&row.ID, &row.FullName, &row.Age, &row.Phone, &row.Address)
	checkErr(err)

	return row
}

func (r *Repository) GetUsersByString(query string) Users {
	rows, err := db.Query("SELECT id, fullName, Age, Phone, Address FROM users WHERE LOWER(FullName) like '%' || LOWER($1) || '%';", query)
	checkErr(err)
	results := Users{}
	var row User

	for rows.Next() {
		err = rows.Scan(&row.ID, &row.FullName, &row.Age, &row.Phone, &row.Address)
		checkErr(err)
		results = append(results, row)
	}

	return results
}

func (r *Repository) InsertUser(u User) error {
	_, err := db.Exec("INSERT INTO users(fullname, age, phone, address) VALUES ($1 ,$2, $3,$4);", u.FullName, u.Age, u.Phone, u.Address)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateUser(u *User) error {
	stmt, err := db.Prepare("UPDATE users SET FullName=$1, Age=$2, Phone=$3, Address=$4 WHERE id=$5;")
	checkErr(err)
	_, err = stmt.Exec(u.FullName, u.Age, u.Phone, u.Address, u.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) DeleteUser(id int) error {
	stmt, err := db.Prepare("DELETE FROM users WHERE id=$1")
	checkErr(err)

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
