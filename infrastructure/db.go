package infrastructure

import(
	"database/sql"
	"fmt"
	"log"
	"os"

	// blank import for MySQL driver
	_ "github.com/go-sql-driver/mysql"
)

type UserDTO struct {
	UserID   string
	Name     string
	Email    string
	Password string
	Admin    string
}

// Driver名
const driverName = "mysql"

// DB 各repositoryで利用するDB接続情報
var DB *sql.DB

func init() {
	/* ===== データベースへ接続する. ===== */
	// ユーザ
	user := os.Getenv("MYSQL_USER")
	// パスワード
	password := os.Getenv("MYSQL_PASSWORD")
	// 接続先ホスト
	host := os.Getenv("MYSQL_HOST")
	// 接続先ポート
	port := os.Getenv("MYSQL_PORT")
	// 接続先データベース
	database := os.Getenv("MYSQL_DATABASE")

	// 接続情報は以下のように指定する.
	// user:password@tcp(host:port)/database
	var err error
	DB, err = sql.Open(driverName,
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database))
	if err != nil {
		log.Fatal(err)
	}
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