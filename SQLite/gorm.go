package main

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

type POST struct{
	gorm.Model
	Title string
	Slug string `gorm:"uniqueIndex:idx_slug"`
	Likes uint
}

func (p POST) String() string{
	return fmt.Sprintf("Post Title : %s \nSlug : %s", p.Title, p.Slug)
}

func createPost(title string, slug string) *POST{
	var returnPost POST = POST{Title:title, Slug:slug}
	if res:= db.Create(&returnPost); res.Error != nil{
		panic(res.Error)
	}

	return &returnPost
}

func getPost(slug string) *POST{
	var targetPost = POST{Slug : slug}
	if res:= db.First(&targetPost); res.Error != nil{
		panic(res.Error)
	}

	return &targetPost
}