package user

import (
	"github.com/fobilow/import-cycle/problem/client"
	"net/http"
)

type User struct {
	name string
	email string
}

type UserService client.Service

func (u *UserService) Create() {
	u.Client.Execute("api/create", http.MethodPost, nil)
}


