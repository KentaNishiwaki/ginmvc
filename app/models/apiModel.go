package models

import (
	"time"
)

type ApiParam struct {
	ID         int
	Username   string
	Token      string `json:"token"`
	Department string
	SelDate    time.Time `json:"selDate"`
}
