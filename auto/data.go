package auto

import (
	"../api/models"
)

var users = []models.User{
	models.User{Nickname: "kick1930", Email: "kick-1930@hotmail.com", Password: "123456"},
}

var posts = []models.Post{
	models.Post{
		Title:   "Title",
		Content: "Hello World",
	},
}
