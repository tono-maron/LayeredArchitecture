package repository

import (
	"LayeredArchitecture/config"
	"LayeredArchitecture/mock"
	"log"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestSelectByPrimaryKey(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testMock := mock.NewMockUserRepository(ctrl)
	call := testMock.EXPECT().SelectByPrimaryKey(config.DB, "mock")
	log.Println(call)
}
