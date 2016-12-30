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
	models.Db.Raw(`
		SELECT s.id AS id,  
		       s.name AS name, 
			   s.address AS address, 
			   s.business_hour AS business_hour, 
			   s.post_user_id AS post_user_id,
			   u.name AS post_user_name
		  FROM shops s INNER JOIN users u ON (s.post_user_id = u.id)
		 ORDER BY s.created_at DESC
	`).Scan(&shops)
	c.RenderArgs["shops"] = &shops

	return c.Render()
}
