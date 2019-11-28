package persistence

import (
	"LayeredArchitecture/infrastructure"
	"database/sql"
	"log"
	"os"
	"testing"

	fixture "github.com/takashabe/go-fixture"
)

type TestDB struct {
	DB *sql.DB
}

func MainTest(m *testing.M) {
	setup()
	//m.Run()の前後にテストの前処理、後処理を書く。
	os.Exit(m.Run())
}

func setup() {
	err := infrastructure.NewDBConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer infrastructure.DB.Close()

	fixture, err := fixture.NewFixture(infrastructure.DB, "mysql")
	if err != nil {
		log.Println(err)
	}
	err = fixture.Load("testdata/schema.sql")
	if err != nil {
		log.Println(err)
	}
}
