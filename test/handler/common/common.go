package common

import (
	"os"
)

func SetupDB() (oldValue string) {
	oldValue = os.Getenv("SQLITE_PATH")
	os.Setenv("SQLITE_PATH", "testdatabase.db")
	return
}

func TearDownDb(oldValue string) {
	os.Setenv("SQLITE_PATH", oldValue)
	os.Remove("testdatabase.db")
}
