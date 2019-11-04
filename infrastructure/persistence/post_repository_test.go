package persistence

import (
	"LayeredArchitecture/infrastructure"
	"database/sql"
	"reflect"
	"testing"

	"LayeredArchitecture/domain"

	fixture "github.com/takashabe/go-fixture"
	_ "github.com/takashabe/go-fixture/mysql"
)

//loadFixture
func loadFixture(t *testing.T, DB *sql.DB, file string) {
	fixture, err := fixture.NewFixture(DB, "mysql")
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	err = fixture.Load(file)
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
}

//SELECTのテスト
/*１件だけ取得
全てのテストデータを別にする：selectされたデータを明確にするため
条件を満たさない場合取得できないことを確認する
複数件取得(検索条件に主キーを指定しない場合も含む)
データは３件以上(検索条件に合致２件以上、合致しない１件以上)：複数件取得、データマッピングが正しいか確認するため
「<」「>=」を使用している場合、同値ケースのテストをする
AND、ORをつなげてる場合、その組み合わせも網羅的にテストする
*/
func TestSelectByPrimaryKey(t *testing.T) {
	err := infrastructure.NewDBConnection()
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	loadFixture(t, infrastructure.DB, "testdata/posts.yml")

	cases := []struct {
		input      int
		expectPost *domain.Post
		expectErr  error
	}{
		{
			1,
			nil,
			nil,
		},
		{
			7,
			&domain.Post{
				PostID:       7,
				Content:      "この上腕二頭筋もこれもう最高やろ？",
				CreateUserID: "319a3327-ae1a-4fa7-ab86-edca1d45c8c2",
			},
			nil,
		},
	}
	for _, c := range cases {
		repo := NewPostRepository(infrastructure.DB)
		post, err := repo.SelectByPrimaryKey(c.input)
		if err != c.expectErr {
			t.Fatalf("#%d: want error %#v, got %#v", c.input, c.expectErr, err)
		}
		if err != nil {
			continue
		}
		if !reflect.DeepEqual(post, c.expectPost) {
			t.Errorf("#%d: want %#v, got %#v", c.input, c.expectPost, post)
		}
	}
}

func TestGetAll(t *testing.T) {
	err := infrastructure.NewDBConnection()
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	loadFixture(t, infrastructure.DB, "testdata/posts.yml")

	cases := []struct {
		expectIDs []int
	}{
		{[]int{2, 3, 4, 5, 6, 7, 8, 10}},
	}
	for i, c := range cases {
		repo := NewPostRepository(infrastructure.DB)
		posts, err := repo.GetAll()
		if err != nil {
			t.Fatalf("#%d: want non error, got %#v", i, err)
		}
		ids := make([]int, 0)
		for _, post := range posts {
			ids = append(ids, post.PostID)
		}
		if !reflect.DeepEqual(ids, c.expectIDs) {
			t.Errorf("#%d: want %#v, got %#v", i, c.expectIDs, ids)
		}
	}
}

/*INSERTのテスト
null以外のデータを登録
nullを許可するカラムにnullを登録
カラム全てに最大桁のデータを登録
一意制約に違反するデータを登録
*/
func TestInsert(t *testing.T) {
	err := infrastructure.NewDBConnection()
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	loadFixture(t, infrastructure.DB, "testdata/posts.yml")

	cases := []struct {
		inputContent string
		inputUserID  string
	}{
		{"I like soccer!", "b52e63b5-42a4-471b-ae36-a0508206cd31"},
		{"I like soccer!", "b52e63b5-42a4-471b-ae36-a0508206cd31"}, // duplicate
		{"", "b52e63b5-42a4-471b-ae36-a0508206cd31"},
	}
	for i, c := range cases {
		repo := NewPostRepository(infrastructure.DB)
		err := repo.Insert(c.inputContent, c.inputUserID)
		if err != nil {
			t.Fatalf("#%d: want non error, got %#v", i, err)
		}

		posts, err := repo.GetAll()
		if err != nil {
			t.Fatalf("#%d: want non error, got %#v", i, err)
		}
		find := false
		for _, post := range posts {
			if post.Content == c.inputContent {
				find = true
				break
			}
		}
		if !find {
			t.Errorf("#%d: want contain content %s, but not found it", i, c.inputContent)
		}
	}
}

/*UPDATAのテスト
null以外のデータを登録
nullを許可するカラムにnullを登録
カラム全てに最大桁のデータを登録
更新対象のレコードがないときの更新
複数件のレコード更新
*/
func TestUpdateByPrimaryKey(t *testing.T) {
	err := infrastructure.NewDBConnection()
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	loadFixture(t, infrastructure.DB, "testdata/posts.yml")

	cases := []struct {
		inputPostID  int
		inputContent string
	}{
		{6, "血管うねうねマスクメロン"},
		{7, "キミプロテイン持ってない？"},
	}
	for i, c := range cases {
		repo := NewPostRepository(infrastructure.DB)
		err := repo.UpdateByPrimaryKey(c.inputPostID, c.inputContent)
		if err != nil {
			t.Fatalf("#%d: want non error, got %#v", i, err)
		}

		//postデータ全件取得
		posts, err := repo.GetAll()
		if err != nil {
			t.Fatalf("#%d: want non error, got %#v", i, err)
		}
		//更新したデータがあるかどうかのチェック
		find := false
		for _, post := range posts {
			if post.Content == c.inputContent {
				find = true
				break
			}
		}
		if !find {
			t.Errorf("#%d: want contain content %s, but not found it", i, c.inputContent)
		}
	}
}

/*DELETEのテスト
存在しないレコードの削除
*/
func TestDeleteByPrimaryKey(t *testing.T) {
	err := infrastructure.NewDBConnection()
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	loadFixture(t, infrastructure.DB, "testdata/posts.yml")

	cases := []struct {
		input []int
	}{
		{[]int{2, 3, 5}},
	}
	//テスト実行
	for i, c := range cases {
		repo := NewPostRepository(infrastructure.DB)
		err := repo.DeleteByPrimaryKey(c.input[i])
		if err != nil {
			t.Fatalf("#%d: want non error, got %#v", i, err)
		}

		//テーブル上のデータを全件取得
		posts, err := repo.GetAll()
		if err != nil {
			t.Fatalf("#%d: want non error, got %#v", i, err)
		}

		//削除した値がないかチェック
		find := false
		for _, post := range posts {
			if post.PostID == c.input[i] {
				find = true
				break
			}
		}
		if find {
			t.Errorf("#%d: want contain content %d, but not found it", i, c.input[i])
		}
	}
}
