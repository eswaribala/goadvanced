package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name        string `json:"full_name"`
	Age         int    `json:"age,omitempty"`
	Active      bool   `json:"-"`
	lastLoginAt string
}

func main() {
	u, err := json.Marshal(User{Name: "Bob", Age: 0, Active: true, lastLoginAt: "today"})
	if err != nil {
		panic(err)
	}
	fmt.Println(string(u)) // {"full_name":"Bob"}
}
