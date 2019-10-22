package domain

type User struct {
	UserID   int
	Name     string
	Email    string
	Password string
	Admin    bool
}

func (u *User) IsAdmin() bool {
   return u.Admin
}