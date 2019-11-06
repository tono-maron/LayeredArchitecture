package persistence

import (
	"LayeredArchitecture/domain"
	"LayeredArchitecture/infrastructure"
	"reflect"
	"testing"
)

func TestSelectUserByPrimaryKey(t *testing.T) {
	err := infrastructure.NewDBConnection()
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	loadFixture(t, infrastructure.DB, "testdata/users.yml")

	cases := []struct {
		input      string
		expectPost *domain.User
		expectErr  error
	}{
		{
			"",
			nil,
			nil,
		},
		{
			"319a3327-ae1a-4fa7-ab86-edca1d45c8c2",
			&domain.User{
				UserID:   "319a3327-ae1a-4fa7-ab86-edca1d45c8c2",
				Name:     "takumi",
				Email:    "takumi@takumi",
				Password: "",
				Admin:    false,
			},
			nil,
		},
	}
	for _, c := range cases {
		repo := NewUserPersistence(infrastructure.DB)
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

func TestUserInsert(t *testing.T) {

}
