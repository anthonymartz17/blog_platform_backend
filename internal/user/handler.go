package user

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	httpResponse "github.com/anthonymartz17/blog_platform_backend.git/internal/transport/http"

	"github.com/gorilla/mux"
)

const (
	msgInvalidBody       = "invalid request body"
	msgInvalidCredential = "invalid email or password"
	msgInternal          = "internal server error"
)

var ErrInvalidCredentials = errors.New("invalid email or password")

// AuthService defines the user authentication operations used by the handler.
type AuthService interface {
	Signup(ctx context.Context, email, password string) (*AuthResponse, error)
	Login(ctx context.Context, email, password string) (*AuthResponse, error)
}


// Handler handles HTTP requests for user authentication.
type Handler struct {
	svc AuthService
}

// NewHandler creates a new user handler.
func NewHandler(svc AuthService) *Handler {
	return &Handler{svc: svc}
}

// RegisterRoutes registers user authentication routes.
func (h *Handler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/auth/signup", h.Signup).Methods(http.MethodPost)
	r.HandleFunc("/auth/login", h.Login).Methods(http.MethodPost)
}

// Signup handles account creation requests.
func (h *Handler) Signup(w http.ResponseWriter, r *http.Request) {
	payload, err := decodeAuthRequest(r)
	if err != nil {
		httpResponse.Error(w, http.StatusBadRequest, msgInvalidBody)
		return
	}

	if err := validateAuthRequest(payload); err != nil {
		httpResponse.Error(w, http.StatusBadRequest, msgInvalidCredential)
		return
	}

	response, err := h.svc.Signup(r.Context(), payload.Email, payload.Password)
	if err != nil {
		httpResponse.Error(w, http.StatusInternalServerError, msgInternal)
		return
	}

	httpResponse.JSON(w, http.StatusCreated, response)
}

// Login handles authentication requests.
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	payload, err := decodeAuthRequest(r)
	if err != nil {
		httpResponse.Error(w, http.StatusBadRequest, msgInvalidBody)
		return
	}

	if err := validateAuthRequest(payload); err != nil {
		httpResponse.Error(w, http.StatusBadRequest, msgInvalidCredential)
		return
	}

	response, err := h.svc.Login(r.Context(), payload.Email, payload.Password)
	if err != nil {
		if errors.Is(err, ErrInvalidCredentials) {
			httpResponse.Error(w, http.StatusUnauthorized, msgInvalidCredential)
			return
		}

		httpResponse.Error(w, http.StatusInternalServerError, msgInternal)
		return
	}

	httpResponse.JSON(w, http.StatusOK, response)
}

func decodeAuthRequest(req *http.Request) (*AuthRequest, error) {
	defer req.Body.Close()

	var payload AuthRequest
	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&payload); err != nil {
		return nil, err
	}

	if err := decoder.Decode(&struct{}{}); err != io.EOF {
		return nil, errors.New("request body must contain a single JSON object")
	}

	return &payload, nil
}

func validateAuthRequest(payload *AuthRequest) error {
	payload.Email = strings.TrimSpace(payload.Email)
	payload.Password = strings.TrimSpace(payload.Password)

	if payload.Email == "" || payload.Password == "" {
		return ErrInvalidCredentials
	}

	return nil
}


