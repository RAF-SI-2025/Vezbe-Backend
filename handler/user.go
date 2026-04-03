package handler

import (
	"encoding/json"
	"net/http"

	"github.com/RAF-SI-2025/Vezbe-Backend/service"
	"github.com/google/uuid"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := service.GetAllUsers()
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, "failed to encode users", http.StatusInternalServerError)
		return
	}
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	user, err := service.GetUserByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "failed to encode user", http.StatusInternalServerError)
		return
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var body struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		Password  string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	user, err := service.CreateUser(body.FirstName, body.LastName, body.Email, body.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "failed to encode user", http.StatusInternalServerError)
		return
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var body struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	user, err := service.UpdateUser(id, body.FirstName, body.LastName, body.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "failed to encode user", http.StatusInternalServerError)
		return
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	if err := service.DeleteUser(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func ChangePassword(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var body struct {
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := service.ChangePassword(id, body.OldPassword, body.NewPassword); err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "invalid credentials" {
			status = http.StatusUnauthorized
		} else if err.Error() == "user not found" {
			status = http.StatusNotFound
		}
		http.Error(w, err.Error(), status)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
