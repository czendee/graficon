package main

import (
	"banwire/dash/dashboard_front/db"
)

// BatchTest is a function only for test
func BatchTest() {

	defer func() {
		db.Connection.Close(nil)
	}()

	/*
		bla bla bla ... more code
	*/
}
