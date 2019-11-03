package persistence

import (
	"LayeredArchitecture/domain"
	"LayeredArchitecture/domain/repository"
	"LayeredArchitecture/infrastructure"
	"database/sql"
)

type userPersistence struct{}

func NewUserPersistence() repository.UserRepository {
	return &userPersistence{}
}

func (up userPersistence) SelectByPrimaryKey(userID string) (*domain.User, error) {
	row := infrastructure.DB.QueryRow("SELECT * FROM user WHERE user_id = ?", userID)
	return convertToUser(row)
}

func (up userPersistence) Insert(userID, name, email, password string, admin bool) error {
	stmt, err := infrastructure.DB.Prepare("INSERT INTO user(user_id, name, email, password, admin) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(userID, name, email, password, admin)
	return err
}

func (up userPersistence) SelectByEmail(email string) (*domain.User, error) {
	row := infrastructure.DB.QueryRow("SELECT * FROM user WHERE email = ?", email)
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
