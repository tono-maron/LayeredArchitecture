package application

import (
	"infrastructure"
	"domain"
)



func SelectUserByPrimaryKey(userID string) (*domain.User, error) {
	//アプリケーションレイヤから
	//直接インフラストラクチャレイヤの実装を利⽤できる。
	dto, err := infrastructure.SelectUserByPrimaryKey(infrastructure.DB, userID)
	if err != nil {
		return nil, err
	}
	//インフラストラクチャレイヤの DTO からモデルを⽣成する
	//直接domain.Userにマッピングしても構わない
	return domain.NewUserFromDTO(dto), nil
}