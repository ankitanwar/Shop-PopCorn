package domain

//AccessToken : Fields for accessToken
type AccessToken struct {
	UserID string `json:"user_id" bson:"_id"`
	Email  string `json:"email" bson:"email"`
	Token  string `json:"access_token" bson:"access_token"`
}
