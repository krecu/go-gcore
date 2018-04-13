package gcore

import (
	"context"
	"net/http"
)

var accountDetailsURL = "/clients/me"

type AccountService service

type Account struct {
	// Your account ID
	ID int `json:"id"`

	// ID of user who requested information
	CurrentUser int `json:"currentUser"`

	// An array which contains information about all users of the requested account
	Users []User `json:"users"`
}

type User struct {
	ID       int     `json:"id"`
	Deleted  bool    `json:"deleted"`
	Email    string  `json:"email"`
	Name     string  `json:"name"`
	Client   int     `json:"client"`
	Company  string  `json:"company"`
	Lang     string  `json:"lang"`
	Phone    string  `json:"phone"`
	Reseller int     `json:"reseller,omitempty"`
	Groups   []Group `json:"groups"`
}

type Group struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (s *AccountService) Details(ctx context.Context) (*Account, *http.Response, error) {

	req, err := s.client.NewRequest(ctx, "GET", accountDetailsURL, nil)
	if err != nil {
		return nil, nil, err
	}

	account := &Account{}

	resp, err := s.client.Do(req, account)
	if err != nil {
		return nil, resp, err
	}

	return account, resp, nil
}