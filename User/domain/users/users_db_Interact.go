package users

import (
	"fmt"
	"strings"
	"time"

	"github.com/ankitanwar/GoAPIUtils/errors"
	userdb "github.com/ankitanwar/user-api/databasource/postgres"
	cryptos "github.com/ankitanwar/user-api/utils/cryptoUtils"
)

const (
	insertUser                = "INSERT INTO users(first_name,last_name,email,date_created,status,password)VALUES(?,?,?,?,?,?) "
	getUser                   = "SELECT id,first_name,last_name,email,date_created FROM users WHERE id=?;"
	errNoRows                 = "no rows in result set"
	updateUser                = "UPDATE users SET first_name=?,last_name=?,email=? WHERE id=?"
	deleteUser                = "DELETE FROM users WHERE id=?"
	getUserByStatus           = "SELECT id,first_name,last_name,email,date_created FROM users WHERE status=?;"
	getUserByEmailAndPassword = "SELECT id,first_name,last_name,email,date_created FROM users WHERE email=? AND password=?;"
)

//Save : To save the user into the database
func (user *User) Save() *errors.RestError {
	stmt, err := userdb.Client.Prepare(insertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	now := time.Now()
	user.DateCreated = now.Format("02-01-2006 15:04")
	user.Password = cryptos.GetMd5(user.Password)
	user.Status = "Active"
	insert, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)

	// insert, err := stmt.Exec(insertUser, user.FirstName, user.LastName, user.Email, user.DateCreated) we can also do it like this

	if err != nil {
		if strings.Contains(err.Error(), "users.email_UNIQUE") {
			return errors.NewBadRequest(fmt.Sprintf("User with %s already exist in the database", user.Email))
		}
		return errors.NewBadRequest(err.Error())
	}
	userid, err := insert.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	user.ID = int(userid)
	return nil

}

//Get : To get the user from the database by given id
func (user *User) Get() *errors.RestError {
	stmt, err := userdb.Client.Prepare(getUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.ID) //to query the single row
	if err := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		if strings.Contains(err.Error(), errNoRows) {
			return errors.NewNotFound(fmt.Sprintf("No user with exist with id %v ", user.ID))
		}
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}

//Update : To  update the value of the existing users
func (user *User) Update() *errors.RestError {
	stmt, err := userdb.Client.Prepare(updateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.ID)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil

}

//Delete : To delete the user from the database
func (user *User) Delete() *errors.RestError {
	stmt, err := userdb.Client.Prepare(deleteUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.ID)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}

//FindByStatus : To find all the users according to their status
func (user *User) FindByStatus(status string) ([]User, *errors.RestError) {
	stmt, err := userdb.Client.Prepare(getUserByStatus)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	rows, err := stmt.Query(status)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer rows.Close()
	result := []User{}
	for rows.Next() {
		var user User
		rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated)
		result = append(result, user)
	}

	if len(result) == 0 {
		return nil, errors.NewNotFound("No User Found With Status")
	}

	return result, nil
}

// GetUserByEmailAndPassword : To reterive the user by email id and password
func (user *User) GetUserByEmailAndPassword() *errors.RestError {
	user.Password = cryptos.GetMd5(user.Password)
	stmt, err := userdb.Client.Prepare(getUserByEmailAndPassword)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.Email, user.Password)
	if err := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		if strings.Contains(err.Error(), errNoRows) {
			return errors.NewNotFound(fmt.Sprintf("No user with exist with id %v ", user.ID))
		}
		return errors.NewInternalServerError(err.Error())
	}
	return nil

}
