package main

import "fmt"

type user struct {
	id       int
	username string
	password string
}

type userservice struct {
	t []user
}

func (u userservice) getallusers() []user {
	return u.t
}

func (u userservice) getuserbyid(id int) user {
	for _, r := range u.t {
		if id == r.id {
			return r
		}
	}

	return user{}
}

func main() {
	users := userservice{
		t: []user{
			{1, "user1", "pass1"},
			{2, "user2", "pass2"},
			{3, "user3", "pass3"},
		},
	}

	allUsers := users.getallusers()
	fmt.Println("Semua Pengguna:")
	for _, u := range allUsers {
		fmt.Printf("ID: %d, Username: %s, Password: %s\n", u.id, u.username, u.password)
	}

	idToFind := 2
	foundUser := users.getuserbyid(idToFind)
	if foundUser.id != 0 {
		fmt.Printf("\nPengguna dengan ID %d ditemukan:\n", idToFind)
		fmt.Printf("ID: %d, Username: %s, Password: %s\n", foundUser.id, foundUser.username, foundUser.password)
	} else {
		fmt.Printf("\nPengguna dengan ID %d tidak ditemukan.\n", idToFind)
	}
}
