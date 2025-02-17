package structs

// User struct for registration
type UserSignUpDetails struct {
	Email     string   `json:"email" binding:"required"`
	Password  string   `json:"password" binding:"required"`
	Gender    string   `json:"gender"`
	Languages []string `json:"languages"` // field to store multiple languages
}

// Login struct for authentication
type UserLoginDetails struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// User struct stores user details
type UserDetails struct {
	Email     string   `bson:"email"`
	Gender    string   `bson:"gender"`
	Languages []string `bson:"languages"`
}
