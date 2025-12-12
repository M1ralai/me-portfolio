package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/M1ralai/me-portfolio/internal/infrasturacture/logger"
	"github.com/M1ralai/me-portfolio/internal/modules/contact/domain"
	"github.com/M1ralai/me-portfolio/internal/modules/contact/service"
	"github.com/go-playground/validator"
)

type ContactHandler struct {
	service  service.ContactService
	logger   *logger.ZapLogger
	validate *validator.Validate
}

func NewContactHandler(service service.ContactService, logger *logger.ZapLogger, validate *validator.Validate) ContactHandler {
	return ContactHandler{
		service:  service,
		logger:   logger,
		validate: validate,
	}
}

func (h ContactHandler) List(w http.ResponseWriter, r *http.Request) {
	resp, err := h.service.List()
	if err != nil {
		h.logger.Error("Failed to get contacts", err, nil)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err = json.NewEncoder(w).Encode(&resp); err != nil {
		h.logger.Error("Failed to write json to a responsewriter", err, nil)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.logger.Info("all contacts getted successfully", nil)
}

func (h ContactHandler) Create(w http.ResponseWriter, r *http.Request) {
	var c domain.Contact
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		h.logger.Error("Failed to read request", err, nil)
		http.Error(w, "Failed to read request", http.StatusInternalServerError)
		return
	}
	if err := h.validate.Struct(c); err != nil {
		h.logger.Error("Validator error while creating contact", err, nil)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.Create(c); err != nil {
		h.logger.Error("failed to create contact", err, nil)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.logger.Info("Contact created successfully", map[string]any{"created": c, "action": "CREATE"})
}

func (h ContactHandler) Delete(w http.ResponseWriter, r *http.Request) {
	sid := r.URL.Query().Get("id")
	id, err := strconv.Atoi(sid)
	if err != nil {
		h.logger.Error("There is no id field in url", err, nil)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.Delete(id); err != nil {
		h.logger.Error("Failed to delete contact", err, nil)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.logger.Info("Contact delete successfully", map[string]any{"action": "DELETE"})
}
