package persistence

import (
	"LayeredArchitecture/domain"
	"LayeredArchitecture/domain/repository"
	"database/sql"
)

type postRepository struct {
	DB *sql.DB
}

func NewPostRepository(DB *sql.DB) repository.PostRepository {
	return &postRepository{DB: DB}
}

func (pr postRepository) SelectByPrimaryKey(postID int) (*domain.Post, error) {
	row := pr.DB.QueryRow("SELECT * FROM post WHERE post_id = ?", postID)
	return convertToPost(row)

}

func (pr postRepository) GetAll() ([]domain.Post, error) {
	rows, err := pr.DB.Query("SELECT * FROM post")
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

func (pr postRepository) Insert(content, userID string) error {

	stmt, err := pr.DB.Prepare("INSERT INTO post(content, user_id) VALUES(?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(content, userID)
	return err
}

func (pr postRepository) UpdateByPrimaryKey(postID int, content string) error {
	stmt, err := pr.DB.Prepare("UPDATE post SET content=? WHERE post_id=?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(content, postID)
	return err
}

func (pr postRepository) DeleteByPrimaryKey(postID int) error {
	stmt, err := pr.DB.Prepare("DELETE FROM post WHERE post_id=?")
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
