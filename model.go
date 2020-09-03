package main

import (
	"database/sql"
)

// Env for DB
type Env struct {
	db *sql.DB
}

// DropdownRes : Dropdown response
type DropdownRes struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

// AuthRes : Auth with Web Portal
type AuthRes struct {
	Status *Status `json:"status"`
	Data   *Data   `json:"data"`
}

// Status from Auth response
type Status struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Data from Auth response
type Data struct {
	IsLoggedIn bool     `json:"isLoggedIn"`
	UserID     string   `json:"userId,omitempty"`
	Roles      []string `json:"roles,omitempty"`
	Token      string   `json:"token,omitempty"`
}
