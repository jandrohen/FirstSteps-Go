package users

import (
	"WebstormProjects/UDEMY/GO/FirstSteps-GO/models"
	"fmt"
	"time"
)

func CreateUser() {
	u := new(models.User)
	u.AddUser(1, "Eduardo", time.Now(), true)
	fmt.Println(u)
}
