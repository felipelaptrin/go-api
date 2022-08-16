package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/felipelaptrin/go-api/pkg/models"
	"github.com/felipelaptrin/go-api/pkg/repository"
	"github.com/felipelaptrin/go-api/pkg/schemas"
	"github.com/felipelaptrin/go-api/pkg/utils"
	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating user...")
	body, _ := io.ReadAll(r.Body)

	var user models.User
	json.Unmarshal(body, &user)
	repository.CreateUser(&user)

	response := schemas.CreatedUser{
		Id:   user.ID,
		Name: user.Name,
	}
	utils.SetResponse(w, 201, "User created with success", response)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Retrieving all users...")

	users, _ := repository.GetAllUsers()

	var response []schemas.CreatedUser
	for i := range *users {
		response = append(response, schemas.CreatedUser{
			Id:   (*users)[i].ID,
			Name: (*users)[i].Name,
		})
	}

	if len(*users) == 0 {
		utils.SetResponse(w, 200, "There are no users registered!", nil)
	} else {
		utils.SetResponse(w, 200, "Users info retrieved with success", response)
	}
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	log.Println("Retrieving user...")
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	user, err := repository.GetUserById(id)

	if err != nil {
		log.Println(err)
		utils.SetResponse(w, 404, "User not found", nil)
		return
	}

	response := schemas.CreatedUser{
		Id:   user.ID,
		Name: user.Name,
	}
	utils.SetResponse(w, 200, "User info retrieved with success", response)
}

func UpdateUserById(w http.ResponseWriter, r *http.Request) {
	log.Println("Updating user...")
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	body, _ := io.ReadAll(r.Body)

	hasUser := repository.CheckUserExists(id)
	if !hasUser {
		utils.SetResponse(w, 404, "The given user does not exist", nil)
		return
	}

	var user models.User
	json.Unmarshal(body, &user)
	repository.UpdateUserById(id, &user)

	utils.SetResponse(w, 200, "User updated with success", nil)
}

func DeleteUserById(w http.ResponseWriter, r *http.Request) {
	log.Println("Deleting user...")
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	hasUser := repository.CheckUserExists(id)
	if !hasUser {
		utils.SetResponse(w, 404, "The given user does not exist", nil)
		return
	}

	repository.DeleteUserById(id)
	utils.SetResponse(w, 200, "User deleted with success", nil)
}
