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

func SelectUserByPrimaryKey(DB *sql.DB, userID string) (*User, error) {
	//インフラストラクチャレイヤの実装を利⽤する。
	dto, err := infrastructure.GetUserByID(DB, id)
	if err != nil {
		return nil, err
	}
	user := &User{
		UserID:   dto.UserID,
		Name:     dto.Name,
		Email:    dto.Email,
		Password: dto.Password,
		Admin:    dto.Admin,
	}
	return user, nil
}