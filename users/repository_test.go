package users

import (
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
)

func TestGetUsers(t *testing.T) {
	r := Repository{}
	db, mock, err := sqlmock.New()
	assert.Nil(t, err)
	defer db.Close()
	r.connectDB(db)

	rows := sqlmock.NewRows([]string{"Id", "FullName", "Age", "Phone", "Address"})
	rows.AddRow(1, "George", 23, 11223344, "Tehran, Iran")
	rows.AddRow(3, "Harry", 12, 121212, "US, CA")
	var people = Users{
		User{
			ID:       1,
			FullName: "George",
			Age:      23,
			Phone:    11223344,
			Address:  "Tehran, Iran",
		},
		User{
			ID:       3,
			FullName: "Harry",
			Age:      12,
			Phone:    121212,
			Address:  "US, CA",
		},
	}

	mock.ExpectQuery("^SELECT (.+) FROM users$").WillReturnRows(rows)
	users := r.GetUsers()
	assert.Equal(t, people, users)

	err = mock.ExpectationsWereMet()
	assert.Nil(t, err, "There were unfulfilled expectations.")
}

func TestGetUserById(t *testing.T) {
	r := Repository{}
	db, mock, err := sqlmock.New()
	assert.Nil(t, err)
	defer db.Close()
	r.connectDB(db)

	rows := sqlmock.NewRows([]string{"ID", "FullName", "Age", "Phone", "Address"})
	rows.AddRow(4, "David", 29, 123456, "indonesia")
	var user = User{
		ID:       4,
		FullName: "David",
		Age:      29,
		Phone:    123456,
		Address:  "indonesia",
		}
	mock.ExpectQuery("^SELECT (.+) FROM users").WithArgs(4).WillReturnRows(rows)


	userById := r.GetUserById(4)
	assert.Equal(t, user, userById)

	err = mock.ExpectationsWereMet()
	assert.Nil(t, err, "There were unfulfilled expectations.")
}

func TestGetUsersByString(t *testing.T) {
	r := Repository{}
	db, mock, err := sqlmock.New()
	assert.Nil(t, err)
	defer db.Close()
	r.connectDB(db)
	rows := sqlmock.NewRows([]string{"ID", "FullName", "Age", "Phone", "Address"})
	rows.AddRow(4, "David", 29, 123456, "indonesia")

	var people = Users{
		User{
			ID:       4,
			FullName: "David",
			Age:      29,
			Phone:    123456,
			Address:  "indonesia",
		},

	}

	mock.ExpectQuery("^SELECT (.+) FROM").WithArgs("David").WillReturnRows(rows)
	userById := r.GetUsersByString("David")
	assert.Equal(t, people, userById)

	err = mock.ExpectationsWereMet()
	assert.Nil(t, err, "There were unfulfilled expectations.")
}

func TestInsertUser(t *testing.T) {
	r := Repository{}
	db, mock, err := sqlmock.New()
	assert.Nil(t, err)
	defer db.Close()
	r.connectDB(db)

	var insertedId = 4

	rows := sqlmock.NewRows([]string{"id"})
	rows.AddRow(insertedId)

	mock.ExpectExec("^INSERT INTO").WithArgs("David", 29, 123456, "indonesia").WillReturnResult(sqlmock.NewResult(0, 1))

	var user = User{
		ID:       4,
		FullName: "David",
		Age:      29,
		Phone:    123456,
		Address:  "indonesia",
	}

	err = r.InsertUser(user)
	assert.Nil(t, err)
	err = mock.ExpectationsWereMet()
	assert.Nil(t, err, "There were unfulfilled expectations.")
}


func TestUpdateUser(t *testing.T) {
	r := Repository{}
	db, mock, err := sqlmock.New()
	assert.Nil(t, err)
	defer db.Close()
	r.connectDB(db)

	var user = &User{
		ID:       4,
		FullName: "David",
		Age:      29,
		Phone:    123456,
		Address:  "indonesia",
	}
	mock.ExpectPrepare("UPDATE").ExpectExec().WithArgs("David", 29, 123456, "indonesia", 4).WillReturnResult(sqlmock.NewResult(0, 1))

	err = r.UpdateUser(user)
	assert.Nil(t, err)
	err = mock.ExpectationsWereMet()
	assert.Nil(t, err, "There were unfulfilled expectations.")
}

func TestDeleteUser(t *testing.T) {
	r := Repository{}
	db, mock, err := sqlmock.New()
	assert.Nil(t, err)
	defer db.Close()
	r.connectDB(db)

	var userID = 4

	mock.ExpectPrepare("DELETE FROM users").ExpectExec().WithArgs(userID).WillReturnResult(sqlmock.NewResult(0, 1))
	err = r.DeleteUser(userID)
	assert.Nil(t, err)

	err = mock.ExpectationsWereMet()
	assert.Nil(t, err, "There were unfulfilled expectations.")
}
