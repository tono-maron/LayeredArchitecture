package persistence


type PostPersistence struct{}

func (postPersistence PostPersistence) SelectByPrimaryKey(DB *sql.DB, postID int) (*domain.Post, error) {
	row := DB.QueryRow("SELECT * FROM post WHERE post_id = ?", userID)
	return convertToUser(row)
}
func (postPersistence PostPersistence) GetAll(DB *sql.DB) ([]domain.Post, error) {
	rows, err := DB.Query("SELECT * FROM post")
	posts := make([]domain.Post)
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

func (postPersistence PostPersistence) Insert(DB *sql.DB, content, userID string) error {
	stmt, err := DB.Prepare("INSERT INTO post(content, create_user_id) VALUES(?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(content, userID)
	return err
}

func (postPersistence PostPersistence) UpdateByPrimaryKey(DB *sql.DB, postID int, content string) error {
	stmt, err := DB.Prepare("UPDATE post SET content=? WHERE post_id=?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(content, postID)
	return err
}

func (postPersistence PostPersistence) Delete(DB *sql.DB, postID int) error {
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
		log.Println(err)
		return nil, err
	}
	return &post, nil
}