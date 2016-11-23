package models

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/revel/revel"
	"log"
	"strconv"
)

type User struct {
	gorm.Model
	Name        string `gorm:"type:varchar(255);unique_index;not null"`
	Password    string `gorm:"type:varchar(255);not null"`
	DisplayName string `gorm:"type:varchar(255)"`

	PasswordConfirm string `gorm:"-"`
}

func (user User) Validate(v *revel.Validation) {
	v.Required(user.Name).Message("Please enter a username")
	v.Required(user.Password).Message("Please enter a password")
	v.MinSize(user.Name, 4).Message("Username must at lease 4 chars long")
	v.MinSize(user.Password, 4).Message("Password must at lease 4 chars long")
	v.MaxSize(user.Name, 15).Message("Username must at most 15 chars long")
	v.MaxSize(user.Password, 15).Message("Password must at most 15 chars long")
}

func (user User) ValidateSignUp(v *revel.Validation) {
	user.Validate(v)

	if user.Password != user.PasswordConfirm {
		v.Error("Passwords do not match")
	}

	if !v.HasErrors() {
		Db.Where("name = ?", user.Name).First(&user)
		log.Printf("User : %#v", user)
		if user.ID != 0 {
			v.Error("User %s is already exists", user.Name)
		}
	}
}

func (user User) SignUp() User {
	var n uint64
	binary.Read(rand.Reader, binary.LittleEndian, &n)
	salt := strconv.FormatUint(n, 36)

	orig := user.Password + salt
	digest := sha256.Sum256([]byte(orig))
	user.Password = fmt.Sprintf("%s$%x", salt, digest)

	log.Println("Salted: %s", user.Password)

	Db.Create(&user)

	return user
}
