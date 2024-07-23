package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/idkOybek/internal/logger"
	"github.com/idkOybek/internal/models"
	"github.com/idkOybek/internal/services"
	"github.com/idkOybek/internal/utils"
)

type AuthHandler struct {
	service *services.AuthService
}

func NewAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

// @Summary Register a new user
// @Tags auth
// @Accept json
// @Produce json
// @Param user body models.UserRegistrationRequest true "User registration request"
// @Success 201 {object} models.User
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /auth/register [post]
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var userReq models.UserRegistrationRequest
	if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		logger.ErrorLogger.Printf("Error decoding user registration request: %v", err)
		return
	}

	user := models.User{
		INN:      userReq.INN,
		Username: userReq.Username,
		Password: userReq.Password,
		IsActive: userReq.IsActive,
		IsAdmin:  userReq.IsAdmin,
	}

	if err := h.service.RegisterUser(r.Context(), user); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to register user")
		logger.ErrorLogger.Printf("Error registering user: %v", err)
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, user)
}

// @Summary Authenticate a user
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body models.UserLoginRequest true "User login request"
// @Success 200 {object} models.UserLoginResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /auth/login [post]
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var creds models.UserLoginRequest
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		logger.ErrorLogger.Printf("Error decoding login request: %v", err)
		return
	}

	user, token, err := h.service.AuthenticateUser(r.Context(), creds.Username, creds.Password)
	if err != nil {
		utils.RespondWithError(w, http.StatusUnauthorized, "Invalid credentials")
		logger.ErrorLogger.Printf("Error authenticating user: %v", err)
		return
	}

	response := models.UserLoginResponse{
		User:  *user,
		Token: token,
	}

	utils.RespondWithJSON(w, http.StatusOK, response)
}

func (h *AuthHandler) AuthRoutes() chi.Router {
	r := chi.NewRouter()
	r.Post("/register", h.Register)
	r.Post("/login", h.Login)
	return r
}
