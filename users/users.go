package users

// User represents a user with a username, password, and a TOTP secret.
type User struct {
	Username string
	Password string
	Secret   string
}

// In-memory "database" of users
var Users = map[string]*User{
	"john": {Username: "john", Password: "password"},

}
