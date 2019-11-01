package persistence

import (
	"LayeredArchitecture/domain"
	"LayeredArchitecture/domain/repository"
	"database/sql"
)

type UserPersistence struct{}

func NewUserPersistence() repository.UserRepository {
	return &UserPersistence{}
}

func (userPersistence UserPersistence) SelectByPrimaryKey(DB *sql.DB, userID string) (*domain.User, error) {
	row := DB.QueryRow("SELECT * FROM user WHERE user_id = ?", userID)
	return convertToUser(row)
}

func (userPersistence UserPersistence) Insert(DB *sql.DB, userID, name, email, password string, admin bool) error {
	stmt, err := DB.Prepare("INSERT INTO user(user_id, name, email, password, admin) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(userID, name, email, password, admin)
	return err
}

func (userPersistence UserPersistence) SelectByEmail(DB *sql.DB, email string) (*domain.User, error) {
	row := DB.QueryRow("SELECT * FROM user WHERE email = ?", email)
	return convertToUser(row)
}

func convertToUser(row *sql.Row) (*domain.User, error) {
	user := domain.User{}
	err := row.Scan(&user.UserID, &user.Name, &user.Email, &user.Password, &user.Admin)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
