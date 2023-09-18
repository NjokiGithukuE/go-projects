package main

type User struct {
	ID int `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role string `json:"role"`
}

type Role struct {
	Name string `json:"name"`
	Access []string `json:"access"`
}