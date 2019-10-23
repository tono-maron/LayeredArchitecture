package infrastructure

import (
	"database/sql"
	"fmt"
	"log"
)

type UserDTO struct {
	UserID   string
	Name     string
	Email    string
	Password string
	Admin    string
}

func SelectUserByPrimaryKey(DB *sql.DB, userID string) (*UserDTO, error) {
	//DB にアクセスするロジック
	row := DB.QueryRow("SELECT * FROM users WHERE user_id = ?", userID)
	return convertToUserDTO
}

// convertToUser rowデータをUserデータへ変換する
func convertToUserDTO(row *sql.Row) (*UserDTO, error) {
	dto := UserDTO{}
	err := row.Scan(&dto.UserID, &dto.Name, &dto.Email, &dto.Password, &dto.Admin)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println(err)
		}
		return nil, err
	}
	return &dto, nil
}
