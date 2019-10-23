package persistence

import (
	"LayeredArchitecture/domain"
	"database/sql"
	"log"
)

type UserPersistence struct{}

func (userPersistence UserPersistence) SelectByPrimaryKey(DB *sql.DB, userID string) (*domain.User, error) {
	row := DB.QueryRow("SELECT * FROM user WHERE user_id = ?", userID)
	return convertToUser(row)
}

func convertToUser(row *sql.Row) (*domain.User, error) {
	user := domain.User{}
	err := row.Scan(&user.UserID, &user.Name, &user.Email, &user.Password, &user.Admin)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Println(err)
		return nil, err
	}
	return &user, nil
}
