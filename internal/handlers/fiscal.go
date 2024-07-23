package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/idkOybek/internal/models"
	"github.com/idkOybek/internal/services"
	"github.com/idkOybek/internal/utils"
)

type FiscalHandler struct {
	service *services.FiscalService
}

func NewFiscalHandler(service *services.FiscalService) *FiscalHandler {
	return &FiscalHandler{service: service}
}

func (h *FiscalHandler) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/", h.GetAllFiscalModules)
	r.Get("/{id}", h.GetFiscalModuleByID)
	r.Post("/", h.CreateFiscalModule)
	r.Put("/{id}", h.UpdateFiscalModule)
	r.Delete("/{id}", h.DeleteFiscalModule)
	return r
}

// GetAllFiscalModules обрабатывает запрос на получение всех фискальных модулей
// @Summary Get all fiscal modules
// @Description Get all fiscal modules
// @Tags fiscal
// @Produce json
// @Success 200 {array} models.FiscalModuleResponse
// @Failure 500 {object} map[string]string
// @Router /fiscal [get]
// @Security BearerAuth
func (h *FiscalHandler) GetAllFiscalModules(w http.ResponseWriter, r *http.Request) {
	fiscalModules, err := h.service.GetAll(r.Context())
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Could not retrieve fiscal modules")
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, fiscalModules)
}

// GetFiscalModuleByID обрабатывает запрос на получение фискального модуля по ID
// @Summary Get fiscal module by ID
// @Description Get fiscal module by ID
// @Tags fiscal
// @Produce json
// @Param id path int true "Fiscal module ID"
// @Success 200 {object} models.FiscalModuleResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /fiscal/{id} [get]
// @Security BearerAuth
func (h *FiscalHandler) GetFiscalModuleByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid fiscal module ID")
		return
	}

	module, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Fiscal module not found")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, module)
}

// CreateFiscalModule обрабатывает запрос на создание нового фискального модуля
// @Summary Create a new fiscal module
// @Description Create a new fiscal module
// @Tags fiscal
// @Accept json
// @Produce json
// @Param module body models.FiscalModuleCreateRequest true "New fiscal module"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /fiscal [post]
// @Security BearerAuth
func (h *FiscalHandler) CreateFiscalModule(w http.ResponseWriter, r *http.Request) {
	var moduleReq models.FiscalModuleCreateRequest
	err := json.NewDecoder(r.Body).Decode(&moduleReq)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	module := models.FiscalModule{
		FactoryNumber: moduleReq.FactoryNumber,
		FiscalNumber:  moduleReq.FiscalNumber,
		UserID:        moduleReq.UserID,
	}

	err = h.service.Create(r.Context(), module)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Could not create fiscal module")
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"message": "Fiscal module created"})
}

// UpdateFiscalModule обрабатывает запрос на обновление фискального модуля
// @Summary Update fiscal module
// @Description Update fiscal module
// @Tags fiscal
// @Accept json
// @Produce json
// @Param id path int true "Fiscal module ID"
// @Param module body models.FiscalModuleUpdateRequest true "Fiscal module data"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /fiscal/{id} [put]
// @Security BearerAuth
func (h *FiscalHandler) UpdateFiscalModule(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid fiscal module ID")
		return
	}

	var moduleReq models.FiscalModuleUpdateRequest
	err = json.NewDecoder(r.Body).Decode(&moduleReq)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	module := models.FiscalModule{
		ID:            id,
		FactoryNumber: moduleReq.FactoryNumber,
		FiscalNumber:  moduleReq.FiscalNumber,
		UserID:        moduleReq.UserID,
	}

	err = h.service.Update(r.Context(), module)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Could not update fiscal module")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Fiscal module updated"})
}

// DeleteFiscalModule обрабатывает запрос на удаление фискального модуля
// @Summary Delete fiscal module
// @Description Delete fiscal module
// @Tags fiscal
// @Produce json
// @Param id path int true "Fiscal module ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /fiscal/{id} [delete]
// @Security BearerAuth
func (h *FiscalHandler) DeleteFiscalModule(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid fiscal module ID")
		return
	}

	err = h.service.Delete(r.Context(), id)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Could not delete fiscal module")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Fiscal module deleted"})
}
