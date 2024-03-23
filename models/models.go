package models

import "gopkg.in/mgo.v2/bson"

type (
	User struct {
		UserID   bson.ObjectId `bson:"_id,omitempty" json:"userID"`
		Username string        `json:"username"`
		Password string        `json:"password"`
		Session  string        `json:"session"`
	}

	Session struct {
		Cookie string `json:"cookie"`
		UserID string `json:"userID"`
	}
	OnlyUserID struct {
		UserID string `json:"userID"`
	}
	OnlyUserIDBson struct {
		UserID bson.ObjectId `bson:"_id" json:"userID"`
	}
	OnlySession struct {
		Session string `json:"session"`
	}
	OnlyUsername struct {
		Username string `json:"username"`
	}
)

type (
	Recipe struct {
		RecipeID        bson.ObjectId `bson:"_id,omitempty" json:"recipeID"`
		RecipeName      string        `json:"recipeName"`
		RecipeDetail    string        `json:"recipeDetail"`
		ImageName       string        `json:"imageName"`
		CurrentUsername string        `json:"currentUsername"`
		Comments        []Comment     `json:"comments"`
	}
	Comment struct {
		User     string `json:"user"`
		Comment  string `json:"comment"`
		IsEdited bool   `json:"isEdited"`
	}
)
