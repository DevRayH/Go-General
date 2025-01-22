package main

import (
	"fmt"
)

func main() {
	// Entry point

	var post POST = POST{}
	//db.AutoMigrate(&post)

	if !db.Migrator().HasTable("posts") {
    	// Create database
    	err = db.AutoMigrate(&post)
    	if err != nil {
        	panic("failed to create database")
    	}
	}

	createPost("This Post", "You musha fucka")

	var oldPost *POST = getPost("You musha fucka")
	fmt.Println(*oldPost)
}

