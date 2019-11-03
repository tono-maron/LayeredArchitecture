package persistence

import (
	"LayeredArchitecture/domain"
	"LayeredArchitecture/domain/repository"
	"LayeredArchitecture/infrastructure"
	"database/sql"
)

type postPersistence struct{}

func NewPostPersistence() repository.PostRepository {
	return &postPersistence{}
}

func (pp postPersistence) SelectByPrimaryKey(postID int) (*domain.Post, error) {
	DB := infrastructure.NewDBConnection()
	row := DB.QueryRow("SELECT * FROM post WHERE post_id = ?", postID)
	return convertToPost(row)

}

func (pp postPersistence) GetAll() ([]domain.Post, error) {
	DB := infrastructure.NewDBConnection()
	rows, err := DB.Query("SELECT * FROM post")
	if err != nil {
		return nil, err
	}
	posts := make([]domain.Post, 0)
	for rows.Next() {
		var post domain.Post
		err := rows.Scan(&post.PostID, &post.Content, &post.CreateUserID)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (pp postPersistence) Insert(content, userID string) error {
	DB := infrastructure.NewDBConnection()
	stmt, err := DB.Prepare("INSERT INTO post(content, user_id) VALUES(?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(content, userID)
	return err
}

func (pp postPersistence) UpdateByPrimaryKey(postID int, content string) error {
	DB := infrastructure.NewDBConnection()
	stmt, err := DB.Prepare("UPDATE post SET content=? WHERE post_id=?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(content, postID)
	return err
}

func (pp postPersistence) DeleteByPrimaryKey(postID int) error {
	DB := infrastructure.NewDBConnection()
	stmt, err := DB.Prepare("DELETE FROM post WHERE post_id=?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(postID)
	return err
}

func convertToPost(row *sql.Row) (*domain.Post, error) {
	post := domain.Post{}
	err := row.Scan(&post.PostID, &post.Content, &post.CreateUserID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &post, nil
}
