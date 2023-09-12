package tests

import (
	"fmt"
	"github.com/projects/person-crud/initializers"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// os.Exit skips defer calls
	// so we need to call another function
	code, err := run(m)
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(code)
}

func run(m *testing.M) (code int, err error) {
	err = os.Setenv("APP_ENV", "testing")
	if err != nil {
		return 0, err
	}

	err = os.Setenv("DB_CONNECTION", "sqlite")
	if err != nil {
		return 0, err
	}
	initializers.InitializeDatabaseConnection()

	return m.Run(), nil
}
