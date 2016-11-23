package controllers

import (
	"github.com/egawata/gekikara/app/models"
	"github.com/revel/revel"
)

func init() {
	revel.OnAppStart(models.InitDB)
}
