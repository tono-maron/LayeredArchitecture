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
	err := infrastructure.NewDBConnection()
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	loadFixture(t, infrastructure.DB, "testdata/posts.yml")

	cases := []struct {
		inputUserID   string
		inputName     string
		inputEmail    string
		inputPassword string
		inputAdmin    bool
	}{
		{"ilikesoccer", "takeshi", "takeshi@takeshi", "taketake", false},
		{"Ilikesoccer", "hiroshi", "hiroshi@hiroshi", "hirohiro", false}, // duplicate
		{"", "tono", "tono@tono", "tonotono", false},
	}
	for i, c := range cases {
		repo := NewUserPersistence(infrastructure.DB)
		err := repo.Insert(c.inputUserID, c.inputName, c.inputEmail, c.inputPassword, c.inputAdmin)
		if err != nil {
			t.Fatalf("#%d: want non error, got %#v", i, err)
		}

		user, err := repo.SelectByPrimaryKey(c.inputUserID)
		if user.UserID == c.inputUserID {
			t.Fatalf("#%d: want error %#v, got %#v", i, c.inputUserID, user.UserID)
		}
		if user.Name == c.inputName {
			t.Fatalf("#%d: want error %#v, got %#v", i, c.inputName, user.Name)
		}
		if user.Email == c.inputEmail {
			t.Fatalf("#%d: want error %#v, got %#v", i, c.inputEmail, user.Email)
		}
		if user.Password == c.inputPassword {
			t.Fatalf("#%d: want error %#v, got %#v", i, c.inputPassword, user.Password)
		}
		if user.Admin == c.inputAdmin {
			t.Fatalf("#%d: want error %#v, got %#v", i, c.inputAdmin, user.Admin)
		}
	}
}
