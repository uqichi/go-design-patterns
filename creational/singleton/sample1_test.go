package singleton

import "testing"

func Test(t *testing.T) {
	GetDB()
	GetDB()
	GetDB()
}
