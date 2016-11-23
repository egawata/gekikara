package controllers

import (
	"github.com/egawata/gekikara/app/models"
	"github.com/revel/revel"
	"log"
	"strconv"
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
		ConvertErrorI18n(c, c.Validation)
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.SignUp)
	}

	user.SignUp()

	return c.Redirect(App.Login)
}

func (c App) Login() revel.Result {
	return c.Render()
}

func (c App) LoginCheck() revel.Result {
	var username, password string
	c.Params.Bind(&username, "username")
	c.Params.Bind(&password, "password")
	user := models.User{
		Name:     username,
		Password: password,
	}

	user = user.ValidateLogin(c.Validation)
	if c.Validation.HasErrors() {
		ConvertErrorI18n(c, c.Validation)
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Login)
	}

	c.Session["userId"] = strconv.Itoa(int(user.ID))

	return c.Redirect("/")
}

func (c App) Logout() revel.Result {
	delete(c.Session, "userId")

	log.Printf("Logout")
	return c.Redirect(App.Index)
}
