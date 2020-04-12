package mydatabase

import (
	"fmt"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	postgresDB *gorm.DB = nil
	mutex               = &sync.Mutex{}
)

// Initialize DB connection
func Initialize() *gorm.DB {
	mutex.Lock()
	if postgresDB == nil {
		fmt.Println("Creating new postgres connection")
		db, err := gorm.Open("postgres", "host=localhost port=5434 user=superadmin dbname=template1 password=tympass sslmode=disable")
		if err != nil {
			fmt.Println("Error while connecting to db")
			fmt.Println(err)
			return nil
		}
		postgresDB = db
	}
	mutex.Unlock()
	return postgresDB
}
