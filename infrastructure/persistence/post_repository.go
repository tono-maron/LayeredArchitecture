package persistence


type PostPersistence struct{}

func (postPersistence PostPersistence) SelectByPrimaryKey(DB *sql.DB, postID int) (*domain.Post, error) {

}
func (postPersistence PostPersistence) GetAll(DB *sql.DB) ([]domain.Post, error) {
		
}

func (postPersistence PostPersistence) Insert(DB *sql.DB, postID int, content, userID string) error {

}

func (postPersistence PostPersistence) Update(DB *sql.DB, postID int, content string) error {

}

func (postPersistence PostPersistence) Delete(DB *sql.DB, postID int) error {

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