package service

import (
	"github.com/katerji/UserAuthKit/db"
	"github.com/katerji/UserAuthKit/db/query"
	"github.com/katerji/UserAuthKit/input"
)

type UserService struct{}

func (service UserService) registerUser(input input.CreateUserInput) (int, error) {
	return db.GetDbInstance().Insert(query.InsertUserQuery, input.Email, input.Password)
}
