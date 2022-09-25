package main

import (
	"encoding/json"
	"net/http"
)

const (
	ErrMarshallUsers     = "Error marshalling the users array"
	ErrMarshallUser      = "Error marshalling the user"
	ErrEmptyArrayUsers   = "Error there are not users"
	ErrUnmarshallingUser = "Error unmarshalling data"
	ErrWithUserFormat    = "Error with the user Format"
)

type User struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var Users []User

func init() {
	Users = []User{
		{
			ID:       1,
			Name:     "Martin",
			Email:    "elTincho@hotmail.com",
			Password: "123456",
		},
		{
			ID:       2,
			Name:     "Pablo",
			Email:    "elMasCapo@hotmail.com",
			Password: "123456",
		},
		{
			ID:       3,
			Name:     "Ana",
			Email:    "anita_la_hobbit@hotmail.com",
			Password: "123456",
		},
	}
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func getUsers(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-type", "application/json")

	if Users == nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{Message: ErrEmptyArrayUsers})
		return
	}

	res, err := json.Marshal(Users)
	if err != nil || string(res) == "null" {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{Message: ErrMarshallUsers})
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func addUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	user := User{}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{Message: ErrUnmarshallingUser})
		return
	}

	if user.Name == "" ||
		user.Email == "" ||
		user.Password == "" {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{Message: ErrWithUserFormat})
		return
	}
	user.ID = uint64(len(Users) + 1)

	res, err := json.Marshal(user)
	if err != nil || string(res) == "null" {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{Message: ErrMarshallUser})
		return
	}
	Users = append(Users, user)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
