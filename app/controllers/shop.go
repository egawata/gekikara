package controllers

import (
	"github.com/egawata/gekikara/app/models"
	"github.com/revel/revel"
)

type Shop struct {
	App
}

func (c *Shop) CreateForm() revel.Result {
	return c.Render()
}

func (c *Shop) CreateComplete() revel.Result {
	var name, address, businessHour string
	c.Params.Bind(&name, "name")
	c.Params.Bind(&address, "address")
	c.Params.Bind(&businessHour, "business_hour")

	Db.Create(&models.Shop{Name: name, Address: address, BusinessHour: businessHour})

	return c.Render()
}