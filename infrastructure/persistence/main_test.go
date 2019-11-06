package persistence

import (
	"LayeredArchitecture/infrastructure"
	"log"
	"os"
	"testing"

	fixture "github.com/takashabe/go-fixture"
)

func MainTest(m *testing.M) {
	setup()
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
