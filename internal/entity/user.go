package entity

type User struct {
	ID           int     `json:"ID" db:"id"`
	Email        *string `json:"email,omitempty" db:"email"`
	FirstName    string  `json:"firstName" db:"first_name"`
	LastName     string  `json:"lastName" db:"last_name"`
	PasswordHash string  `json:"password" db:"password"`
	Phone        string  `json:"phone" db:"phone"`
	Username     string  `json:"username" db:"username"`
}

type UserToSignUp struct {
	Email     *string `json:"email,omitempty"`
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	Password  string  `json:"password"`
	Phone     string  `json:"phone"`
	Username  string  `json:"username"`
}

type UserToSignIn struct {
	Password string `json:"password"`
	Username string `json:"username"`
}
