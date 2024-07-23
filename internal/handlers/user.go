package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/idkOybek/internal/logger"
	"github.com/idkOybek/internal/models"
	"github.com/idkOybek/internal/services"
	"github.com/idkOybek/internal/utils"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// @Summary Get all users
// @Tags user
// @Security BearerAuth
// @Produce json
// @Success 200 {array} models.User
// @Failure 500 {object} map[string]string
// @Router /users [get]
func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetAllUsers(r.Context())
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to retrieve users")
		logger.ErrorLogger.Printf("Error retrieving users: %v", err)
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, users)
}

// @Summary Get user by ID
// @Tags user
// @Security BearerAuth
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/{id} [get]
func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid user ID")
		logger.ErrorLogger.Printf("Invalid user ID: %v", err)
		return
	}

	user, err := h.service.GetUserByID(r.Context(), id)
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "User not found")
		logger.ErrorLogger.Printf("User not found: %v", err)
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, user)
}

// @Summary Create a new user
// @Tags user
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param user body models.UserRegistrationRequest true "New user data"
// @Success 201 {object} models.User
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users [post]
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userRequest models.UserRegistrationRequest
	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		logger.ErrorLogger.Printf("Error decoding user creation request: %v", err)
		return
	}

	user := &models.User{
		INN:      userRequest.INN,
		Username: userRequest.Username,
		Password: userRequest.Password,
		IsActive: userRequest.IsActive,
		IsAdmin:  userRequest.IsAdmin,
	}

	if err := h.service.CreateUser(r.Context(), user); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to create user")
		logger.ErrorLogger.Printf("Error creating user: %v", err)
		return
	}
	utils.RespondWithJSON(w, http.StatusCreated, user)
}

// @Summary Update user
// @Tags user
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body models.UserUpdateRequest true "Updated user data"
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/{id} [put]
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid user ID")
		logger.ErrorLogger.Printf("Invalid user ID: %v", err)
		return
	}

	var userRequest models.UserUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		logger.ErrorLogger.Printf("Error decoding user update request: %v", err)
		return
	}

	user := &models.User{
		ID:       id,
		INN:      userRequest.INN,
		Username: userRequest.Username,
		Password: userRequest.Password,
		IsActive: userRequest.IsActive,
		IsAdmin:  userRequest.IsAdmin,
	}

	if err := h.service.UpdateUser(r.Context(), user); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to update user")
		logger.ErrorLogger.Printf("Error updating user: %v", err)
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, user)
}

// @Summary Delete user
// @Tags user
// @Security BearerAuth
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/{id} [delete]
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid user ID")
		logger.ErrorLogger.Printf("Invalid user ID: %v", err)
		return
	}

	if err := h.service.DeleteUser(r.Context(), id); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to delete user")
		logger.ErrorLogger.Printf("Error deleting user: %v", err)
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, "User deleted")
}

func (h *UserHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", h.GetAllUsers)
	r.Post("/", h.CreateUser)
	r.Get("/{id}", h.GetUserByID)
	r.Put("/{id}", h.UpdateUser)
	r.Delete("/{id}", h.DeleteUser)

	return r
}
