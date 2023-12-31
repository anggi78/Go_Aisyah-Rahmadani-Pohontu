package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id int `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users = []User{
	{Id: 1, Email: "anggikiyowo@gmail.com", Name: "Anggi Pohontu", Password: "anggiii"},
	{Id: 2, Email: "aisyahrahmadani@gmail.com", Name: "Aisyah Rahmadani Pohontu", Password: "aisyahahaha"},
	{Id: 3, Email: "anggita@gmail.com", Name: "anggita", Password: "dkdrl"},
}

// get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users": users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var foundUser User
	for _, user := range users {
		if user.Id == id {
			foundUser = user
			break
		}
	}

	if foundUser.Id == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "User not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user by ID",
		"user":    foundUser,
	})
}


// delete user by id
func DeleteUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	for i, user := range users {
		if user.Id == id {
			users = append(users[:i], users[i+1:]...)
			return c.JSON(http.StatusOK, map[string]interface{}{
				"message": "success delete user",
			})
		}
	}

	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"message": "user not found",
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	for i, user := range users {
		if user.Id == id {
			updatedUser := new(User)
			if err := c.Bind(updatedUser); err != nil {
				return err
			}

			users[i].Name = updatedUser.Name
			users[i].Email = updatedUser.Email
			users[i].Password = updatedUser.Password

			return c.JSON(http.StatusOK, map[string]interface{}{
				"message": "success update user",
				"user": users[i],
			})
		}
	}

	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"message": "user not found",
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	// binding data
	user := User{}
	c.Bind(&user)
	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user": user,
	})
}

func main() {
	e := echo.New()

	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUserController)
	e.POST("/users", CreateUserController)
	e.DELETE("/users/:id", DeleteUserController)
	e.PUT("/users/:id", UpdateUserController)

	e.Logger.Fatal(e.Start(":8000"))
}