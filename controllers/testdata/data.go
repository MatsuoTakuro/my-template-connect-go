package testdata

import "github.com/MatsuoTakuro/my-template-connect-go/models"

var articleTestData = []models.Article{
	models.Article{
		ID:          1,
		Title:       "firstPost",
		Contents:    "This is my first blog",
		UserName:    "Takuro Matsuo",
		NiceNum:     2,
		CommentList: commentTestData,
	},
	models.Article{
		ID:       2,
		Title:    "2nd",
		Contents: "Second blog post",
		UserName: "Takuro Matsuo",
		NiceNum:  4,
	},
}
}
