package users

import "encoding/json"

//PublicUser : Public User Marshall Struct
type PublicUser struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

//PrivateUser : When user is interacting with itself
type PrivateUser struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

//MarshallUser : to Marshal the slice of users
func (user Users) MarshallUser(isPublic bool) []interface{} {
	result := make([]interface{}, len(user))
	for index, value := range user {
		result[index] = value.MarshallUser(isPublic)
	}
	return result
}

//MarshallUser : To marshall the User
func (user *User) MarshallUser(isPublic bool) interface{} {

	if isPublic {
		user := &PublicUser{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email, //if json filed of Public user is diff than json fielf of user then this method is Useful
		}
		return user
	}
	userJSON, _ := json.Marshal(user)
	privateuser := &PrivateUser{}
	if err := json.Unmarshal(userJSON, privateuser); err != nil {
		return nil
	}
	return privateuser

}
