package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/kireeti-28/moms-music/pkg/auth"
	"github.com/kireeti-28/moms-music/pkg/database"
	"github.com/kireeti-28/moms-music/pkg/utils"
)

func Login(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	user := database.User{}
	err := decoder.Decode(&user)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = validateUserCredentials(user.Email, user.Password)
	if err != nil {
		utils.RespondWithError(w, http.StatusUnauthorized, "invalid credentials")
		return
	}

	type resp struct {
		Message string `json:"message"`
	}

	response := resp{
		Message: "Logged In Sucessfully",
	}

	utils.RespondWithJSON(w, http.StatusOK, response) // in future we might also send token.
}

// ValidateUserCredentials validates user credentials
func validateUserCredentials(email, password string) error {
	email = strings.ToLower(email) // make to lower case
	dbUser, err := database.GetUser(email)
	if err != nil {
		return err
	}

	err = auth.ComparePasswordHash([]byte(dbUser.HashPassword), []byte(password))
	if err != nil {
		return err
	}

	return nil
}

// Register stores user into database with password being hashed
func Register(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	data := database.User{}
	err := decoder.Decode(&data)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	dat, _ := database.GetUser(data.Email) 

	// if user is not nil return error
	if dat.Email == data.Email {
		utils.RespondWithError(w, http.StatusBadRequest, "email already existed")
		return
	}

	hash, err := auth.HashPassword(data.Password)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	user := database.User{}
	user.Email = strings.ToLower(data.Email)
	user.Password = hash

	err = database.InsertUser(user)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	type resp struct {
		Message string `json:"message"`
	}

	response := resp{
		Message: "Registered Sucessfully",
	}

	utils.RespondWithJSON(w, http.StatusOK, response)
}
