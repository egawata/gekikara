package controllers

import (
	"github.com/egawata/gekikara/app/models"
	"github.com/revel/revel"
	"strconv"
)

type Shop struct {
	App
}

func (c *Shop) CreateForm() revel.Result {
	if c.Session["userId"] == "" {
		return c.Redirect("/login")
	}

	return c.Render()
}

func (c *Shop) CreateComplete() revel.Result {
	userIdStr := c.Session["userId"]
	if userIdStr == "" {
		return c.Redirect("/login")
	}
	userId, err := strconv.Atoi(c.Session["userId"])
	if err != nil {
		return c.Redirect("/login")
	}

	var name, address, businessHour string
	c.Params.Bind(&name, "name")
	c.Params.Bind(&address, "address")
	c.Params.Bind(&businessHour, "business_hour")

	models.Db.Create(&models.Shop{
		Name:         name,
		Address:      address,
		BusinessHour: businessHour,
		PostUserId:   uint64(userId),
	})

	return c.Render()
}

func (c *Shop) List() revel.Result {
	var shops []models.Shop
	models.Db.Find(&shops)
	c.RenderArgs["shops"] = &shops

	return c.Render()
}
