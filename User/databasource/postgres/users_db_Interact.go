package userdb

import (
	"fmt"
	"strings"
	"time"

	"github.com/ankitanwar/GoAPIUtils/errors"
	"github.com/ankitanwar/e-Commerce/User/domain/users"
	User "github.com/ankitanwar/e-Commerce/User/domain/users"
	cryptos "github.com/ankitanwar/e-Commerce/User/utils/cryptoUtils"
)

const (
	insertUser                = "INSERT INTO users(first_name,last_name,email,date_created,status,password,phone)VALUES(?,?,?,?,?,?,?) "
	getUser                   = "SELECT id,first_name,last_name,email,date_created,phone FROM users WHERE id=?;"
	errNoRows                 = "no rows in result set"
	updateUser                = "UPDATE users SET first_name=?,last_name=?,email=?,phone=? WHERE id=?"
	deleteUser                = "DELETE FROM users WHERE id=?"
	getUserByStatus           = "SELECT id,first_name,last_name,email,date_created FROM users WHERE status=?;"
	getUserByEmailAndPassword = "SELECT id,first_name,last_name,email,date_created FROM users WHERE email=? AND password=?;"
	mongoNotFound             = "mongo: no documents in result"
)

//Save : To save the user into the database
func Save(user *User.User) *errors.RestError {
	stmt, err := Client.Prepare(insertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	now := time.Now()
	user.DateCreated = now.Format("02-01-2006 15:04")
	user.Password = cryptos.GetMd5(user.Password)
	user.Status = "Active"
	insert, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password, user.PhoneNo)

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

	user.ID = fmt.Sprint(userid)
	return nil

}

//Get : To get the user from the database by given id
func Get(userID string) (*users.User, *errors.RestError) {
	user := &User.User{}
	stmt, err := Client.Prepare(getUser)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	result := stmt.QueryRow(userID) //to query the single row
	if err := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.PhoneNo); err != nil {
		if strings.Contains(err.Error(), errNoRows) {
			return nil, errors.NewNotFound(fmt.Sprintf("No user with exist with id %v ", user.ID))
		}
		return nil, errors.NewInternalServerError(err.Error())
	}
	return user, nil
}

//Update : To  update the value of the existing users
func Update(user *users.User) *errors.RestError {
	stmt, err := Client.Prepare(updateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.ID, user.PhoneNo)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil

}

//Delete : To delete the user from the database
func Delete(userID string) *errors.RestError {
	stmt, err := Client.Prepare(deleteUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(userID)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}

//FindByStatus : To find all the users according to their status
func FindByStatus(status string) ([]User.User, *errors.RestError) {
	stmt, err := Client.Prepare(getUserByStatus)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	rows, err := stmt.Query(status)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer rows.Close()
	result := []User.User{}
	for rows.Next() {
		var user User.User
		rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated)
		result = append(result, user)
	}

	if len(result) == 0 {
		return nil, errors.NewNotFound("No User Found With Status")
	}

	return result, nil
}

// GetUserByEmailAndPassword : To reterive the user by email id and password
func GetUserByEmailAndPassword(user *users.User) *errors.RestError {
	user.Password = cryptos.GetMd5(user.Password)
	stmt, err := Client.Prepare(getUserByEmailAndPassword)
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
