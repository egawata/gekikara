package controllers

import (
	"github.com/egawata/gekikara/app/models"
	"github.com/revel/revel"
	"log"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) SignUp() revel.Result {
	return c.Render()
}

func (c App) SignUpCheck() revel.Result {
	var username, password, password_confirm string
	c.Params.Bind(&username, "username")
	c.Params.Bind(&password, "password")
	c.Params.Bind(&password_confirm, "password_confirm")
	user := &models.User{
		Name:            username,
		Password:        password,
		PasswordConfirm: password_confirm,
	}

	user.ValidateSignUp(c.Validation)
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.SignUp)
	}

	return c.Redirect(App.Login)
}

func (c App) Login() revel.Result {
	return c.Render()
}

func (c App) LoginCheck() revel.Result {
	var username, password string
	c.Params.Bind(&username, "username")
	c.Params.Bind(&password, "password")
	log.Printf("Username: %s, Password: %s", username, password)

	return c.Redirect("/")
}
