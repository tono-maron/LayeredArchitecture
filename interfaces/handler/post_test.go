package handler

import (
	"LayeredArchitecture/domain"
	"LayeredArchitecture/domain/mock_repository"
	"database/sql"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/julienschmidt/httprouter"
)

func Routes() *httprouter.Router {

	router := httprouter.New()

	// Post Route
	router.GET("/post/:id", HandlePostGet)
	// router.GET("/posts/index", middleware.Authenticate(postHandler.HandlePostsGet))
	// router.POST("/post/create", middleware.Authenticate(postHandler.HandlePostCreate))
	// router.PUT("/post/:id", middleware.Authenticate(postHandler.HandlePostUpdate))
	// router.DELETE("/post/:id", middleware.Authenticate(postHandler.HandlePostDelete))

	return router
}

func prepareServer(t *testing.T) *httptest.Server {
	return httptest.NewServer(Routes())
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
		input    int
		wantBody []byte
		wantCode int
		mock     func(*mock_repository.MockPostRepository)
	}{
		{
			input:    1,
			wantBody: []byte(`{"PostID":1, "Content":"I like soccer."}, "CreateUserID":"abcd"`),
			wantCode: http.StatusOK,
			mock: func(r *mock_repository.MockPostRepository) {
				r.EXPECT().SelectByPrimaryKey(1).Return(&domain.Post{PostID: 1, Content: "I like soccer.", CreateUserID: "abcd"}, nil)
			},
		},
		{
			input:    0,
			wantBody: nil,
			wantCode: http.StatusNotFound,
			mock: func(r *mock_repository.MockPostRepository) {
				r.EXPECT().SelectByPrimaryKey(0).Return(nil, sql.ErrNoRows)
			},
		},
	}
	for i, tt := range cases {
		t.Run(fmt.Sprintf("case#%d", i), func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			postRepo := mock_repository.NewMockPostRepository(ctrl)
			tt.mock(postRepo)

			// ph := &postHandler{
			// 	postUsecase: postRepo,
			// }

			ts := prepareServer(t)
			defer ts.Close()

			res := sendRequest(t, "GET", fmt.Sprintf("%s/post/%d", ts.URL, tt.input), nil)
			defer res.Body.Close()

			if tt.wantCode != res.StatusCode {
				t.Errorf("want %d, got %d", tt.wantCode, res.StatusCode)
			}
			if res.StatusCode != http.StatusOK {
				return
			}

			payload, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("want non error, got %#v", err)
			}
			if diff := cmp.Diff(tt.wantBody, payload); diff != "" {
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
