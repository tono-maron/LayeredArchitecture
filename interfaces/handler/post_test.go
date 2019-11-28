package handler

import (
	"LayeredArchitecture/domain"
	"LayeredArchitecture/domain/mock_repository"
	"LayeredArchitecture/usecase"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/takashabe/go-router"
)

func (ph postHandler) Routes() *router.Router {
	r := router.NewRouter()
	r.Get("/post/:id", ph.HandlePostGet)
	r.Post("/post/:id", ph.HandlePostCreate)
	r.Post("/posts", ph.HandlePostsGet)
	r.Delete("/post/:id", ph.HandlePostDelete)
	r.Put("/post/:id", ph.HandlePostUpdate)
	return r
}

func prepareServer(t *testing.T, ph postHandler) *httptest.Server {
	return httptest.NewServer(ph.Routes())
}

func sendRequest(t *testing.T, method, url string, body io.Reader) *http.Response {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	return res
}

func TestHandlePostGet(t *testing.T) {
	cases := []struct {
		requestBody string
		wantBody    string
		wantCode    int
		mock        func(r *mock_repository.MockPostRepository)
	}{
		{
			requestBody: "",
			wantBody:    "InternalServerError",
			wantCode:    200,
			mock: func(r *mock_repository.MockPostRepository) {
				r.EXPECT().SelectByPrimaryKey(10).Return("", errors.New("Internal Server Error!"))
			},
		},
		{
			requestBody: "",
			wantBody:    "",
			wantCode:    500,
			mock: func(r *mock_repository.MockPostRepository) {
				r.EXPECT().SelectByPrimaryKey(1).Return("", nil)
			},
		},
		{
			requestBody: "",
			wantBody:    `PostID: 2, Content: "I like humbarger.", CreateUserID: "aafhisfh"`,
			wantCode:    200,
			mock: func(r *mock_repository.MockPostRepository) {
				r.EXPECT().SelectByPrimaryKey(1).Return(&domain.Post{PostID: 2, Content: "I like humbarger.", CreateUserID: "aafhisfh"}, nil)
			},
		},
		{
			requestBody: "",
			wantBody:    `PostID: 1, Content: "I like soccer.", CreateUserID: "abcd"`,
			wantCode:    200,
			mock: func(r *mock_repository.MockPostRepository) {
				r.EXPECT().SelectByPrimaryKey(1).Return(&domain.Post{PostID: 1, Content: "I like soccer.", CreateUserID: "abcd"}, nil)
			},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case#%d", i), func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			postRepo := mock_repository.NewMockPostRepository(ctrl)
			c.mock(postRepo)
			pu := usecase.NewPostUsecase(postRepo)
			ph := NewPostHandler(pu)

			ts := prepareServer(t, ph)
			defer ts.Close()

			res := sendRequest(t, "GET", fmt.Sprintf("%s/post/%d", ts.URL), nil)
			defer res.Body.Close()

			if c.wantCode != res.StatusCode {
				t.Errorf("want %d, got %d", c.wantCode, res.StatusCode)
			}
			if res.StatusCode != http.StatusOK {
				return
			}

			payload, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("want non error, got %#v", err)
			}
			if diff := cmp.Diff(c.wantBody, payload); diff != "" {
				t.Errorf("body mismatch %s", string(diff))
			}
		})
	}
}

func TestHandlePostsGet(t *testing.T) {

}

func TestHandlePostCreate(t *testing.T) {

}

func TestHandlePostUpdate(t *testing.T) {

}

func TestHandlePostDelete(t *testing.T) {

}
