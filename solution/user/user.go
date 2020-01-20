package user

import (
	"github.com/fobilow/import-cycle/solution/service"
	"net/http"
)

type User struct {
	name string
	email string
}

type Service service.Service

func (u *Service) Create() {
	u.Client.Execute("api/create", http.MethodPost, nil)

}