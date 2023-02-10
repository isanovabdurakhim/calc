
package main

import (
	"fmt"

	"app/controller"
	"app/models"
)

func main() {

	controller.CreateUser(
		models.User{
			Id: 1,
			Name: "Shohruh",
			Surname: "Safarov",
		},
	)

	controller.CreateUser(
		models.User{
			Id: 2,
			Name: "Abduqodir",
			Surname: "Musayev",
		},
	)

	for _, user := range controller.GetListUser() {
		fmt.Println(user)
	}

	res:=controller.Getbyid(1)
	fmt.Println(res)
	fmt.Println(controller.Update(models.User{1,"yangi","yangi"}))

}