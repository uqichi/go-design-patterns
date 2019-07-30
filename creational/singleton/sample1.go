package singleton

import (
	"fmt"
	"sync"
)

type db struct{}

var once sync.Once

var instance *db

func initDB() {
	fmt.Println("initialize db")
	instance = new(db)
}

func GetDB() *db {
	once.Do(func() {
		initDB()
	})
	return instance
}
