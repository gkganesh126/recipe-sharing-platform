package controllers

import (
	"github.com/gkganesh126/recipe-sharing-platform/models"
)

type (
	// For Get
	UserResources struct {
		Data []models.User `json:"data"`
	}
	// For Post/Put
	UserResource struct {
		Data models.User `json:"data"`
	}
	OnlyUserResource struct {
		Data models.OnlyUserID `json:"data"`
	}
)
