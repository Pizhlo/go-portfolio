// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"time"
)

type Admin struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Education struct {
	ID          int32     `json:"id"`
	Name        string    `json:"name"`
	Link        string    `json:"link"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}

type Invitation struct {
	ID      int32     `json:"id"`
	HrName  string    `json:"hr_name"`
	Phone   string    `json:"phone"`
	Email   string    `json:"email"`
	Company string    `json:"company"`
	Message string    `json:"message"`
	Date    time.Time `json:"date"`
}

type Project struct {
	ID          int32     `json:"id"`
	Name        string    `json:"name"`
	Link        string    `json:"link"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}
