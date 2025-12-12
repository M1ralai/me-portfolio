package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/M1ralai/me-portfolio/internal/infrasturacture/logger"
	"github.com/M1ralai/me-portfolio/internal/modules/post/domain"
	"github.com/M1ralai/me-portfolio/internal/modules/post/service"
	"github.com/go-playground/validator"
)

type PostHandler struct {
	service  service.PostService
	logger   *logger.ZapLogger
	validate *validator.Validate
}

func NewPostHandler(service service.PostService, logger *logger.ZapLogger, validate *validator.Validate) *PostHandler {
	return &PostHandler{
		service:  service,
		logger:   logger,
		validate: validate,
	}
}

func (h PostHandler) List(w http.ResponseWriter, r *http.Request) {
	resp, err := h.service.List()
	if err != nil {
		h.logger.Error("Failed to get list of posts", err, nil)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err = json.NewEncoder(w).Encode(&resp); err != nil {
		h.logger.Error("Failed to write json to a responsewriter", err, nil)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.logger.Info("All posts getted successfully", nil)
}

func (h PostHandler) GetById(w http.ResponseWriter, r *http.Request) {
	sid := r.URL.Query().Get("id")
	id, err := strconv.Atoi(sid)
	if err != nil {
		h.logger.Error("id field in url query is not integer", err, nil)
		http.Error(w, "id field in url query is not integer", http.StatusBadRequest)
		return
	}
	resp, err := h.service.GetById(id)
	if err != nil {
		h.logger.Error("Failed to get post", err, nil)
		http.Error(w, "Failed to get post", http.StatusInternalServerError)
		return
	}
	if err = json.NewEncoder(w).Encode(resp); err != nil {
		h.logger.Error("Failed to write json to a responsewriter", err, nil)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.logger.Info("Post successfully getted by id", nil)
}

func (h PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var p domain.Post
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		h.logger.Error("Failed to read json from a responsewriter", err, nil)
		http.Error(w, "Failed to read json from a responsewriter", http.StatusBadRequest)
		return
	}
	if err := h.validate.Struct(p); err != nil {
		h.logger.Error("Validator error", err, nil)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var ctx context.Context
	if err = h.service.Create(ctx, p); err != nil {
		h.logger.Error("Failed to create post", err, nil)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.logger.Info("Post successfully created", map[string]any{"created": p, "action": "CREATE"})
}

func (h PostHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	//TODO contexti middlewarede requeste gom
	var ctx context.Context
	sid := r.URL.Query().Get("id")
	id, err := strconv.Atoi(sid)
	if err != nil {
		h.logger.Error("id field in url query is not integer", err, nil)
		http.Error(w, "id field in url query is not integer", http.StatusBadRequest)
		return
	}
	err = h.service.Delete(ctx, id)
	if err != nil {
		h.logger.Error("Failed to delete post", err, nil)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.logger.Info("Post successfully deleted", map[string]any{"action": "DELETE"})
}

func (h PostHandler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	//TODO contexti middlewarede requeste gom
	var ctx context.Context
	var p domain.Post
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		h.logger.Error("Failed to read json from a responsewriter", err, nil)
		http.Error(w, "Failed to read json from a responsewriter", http.StatusBadRequest)
		return
	}
	if err := h.validate.Struct(p); err != nil {
		h.logger.Error("Validator error", err, nil)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.Update(ctx, p); err != nil {
		h.logger.Error("Failed to update post", err, nil)
		http.Error(w, "Failed to update post", http.StatusInternalServerError)
		return
	}
	h.logger.Info("Post successfully updated", map[string]any{"action": "UPDATE"})
}
