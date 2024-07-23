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

type TerminalHandler struct {
	service *services.TerminalService
}

func NewTerminalHandler(service *services.TerminalService) *TerminalHandler {
	return &TerminalHandler{service: service}
}

// @Summary Get all terminals
// @Tags terminal
// @Security BearerAuth
// @Produce json
// @Success 200 {array} models.Terminal
// @Failure 500 {object} map[string]string
// @Router /terminal [get]
func (h *TerminalHandler) GetAllTerminals(w http.ResponseWriter, r *http.Request) {
	terminals, err := h.service.GetAllTerminals(r.Context())
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Could not retrieve terminals")
		logger.ErrorLogger.Printf("Error in GetAllTerminals handler: %v", err)
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, terminals)
}

// @Summary Get terminal by ID
// @Tags terminal
// @Security BearerAuth
// @Produce json
// @Param id path int true "Terminal ID"
// @Success 200 {object} models.Terminal
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /terminal/{id} [get]
func (h *TerminalHandler) GetTerminalByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid terminal ID")
		logger.ErrorLogger.Printf("Invalid terminal ID: %v", err)
		return
	}

	terminal, err := h.service.GetTerminalByID(r.Context(), id)
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Terminal not found")
		logger.ErrorLogger.Printf("Error in GetTerminalByID handler: %v", err)
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, terminal)
}

// @Summary Create a new terminal
// @Tags terminal
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param terminal body models.Terminal true "New terminal data"
// @Success 201 {object} models.Terminal
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /terminal [post]
func (h *TerminalHandler) CreateTerminal(w http.ResponseWriter, r *http.Request) {
	var terminal models.Terminal
	if err := json.NewDecoder(r.Body).Decode(&terminal); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		logger.ErrorLogger.Printf("Error decoding terminal creation request: %v", err)
		return
	}

	if err := h.service.CreateTerminal(r.Context(), &terminal); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to create terminal")
		logger.ErrorLogger.Printf("Error in CreateTerminal handler: %v", err)
		return
	}
	utils.RespondWithJSON(w, http.StatusCreated, terminal)
}

// @Summary Update terminal
// @Tags terminal
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "Terminal ID"
// @Param terminal body models.Terminal true "Updated terminal data"
// @Success 200 {object} models.Terminal
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /terminal/{id} [put]
func (h *TerminalHandler) UpdateTerminal(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid terminal ID")
		logger.ErrorLogger.Printf("Invalid terminal ID: %v", err)
		return
	}

	var terminal models.Terminal
	if err := json.NewDecoder(r.Body).Decode(&terminal); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		logger.ErrorLogger.Printf("Error decoding terminal update request: %v", err)
		return
	}

	terminal.ID = id

	if err := h.service.UpdateTerminal(r.Context(), &terminal); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to update terminal")
		logger.ErrorLogger.Printf("Error in UpdateTerminal handler: %v", err)
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, terminal)
}

// @Summary Delete terminal
// @Tags terminal
// @Security BearerAuth
// @Produce json
// @Param id path int true "Terminal ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /terminal/{id} [delete]
func (h *TerminalHandler) DeleteTerminal(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid terminal ID")
		logger.ErrorLogger.Printf("Invalid terminal ID: %v", err)
		return
	}

	if err := h.service.DeleteTerminal(r.Context(), id); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to delete terminal")
		logger.ErrorLogger.Printf("Error in DeleteTerminal handler: %v", err)
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, "Terminal deleted")
}

func (h *TerminalHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", h.GetAllTerminals)
	r.Post("/", h.CreateTerminal)
	r.Get("/{id}", h.GetTerminalByID)
	r.Put("/{id}", h.UpdateTerminal)
	r.Delete("/{id}", h.DeleteTerminal)

	return r
}
