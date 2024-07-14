package user

import (
	"net/http"

	"github.com/chavocito/ecom/types"
	"github.com/chavocito/ecom/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	store *types.UserStore
}

func (h *Handler) RegisterRoutes(subrouter *mux.Router) {
	subrouter.HandleFunc("/login", h.handleLogin)
	subrouter.HandleFunc("/register", h.handleRegistration)
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin)
	router.HandleFunc("/register", h.handleRegistration)
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	err := r.Body.Close()
	if err != nil {
		return
	}
	//collect user credentials from request body
	//hash and salt password
	//compare to hashed and salted password
	//auth if it's a match with auth token and refresh token
}

func (h *Handler) handleRegistration(w http.ResponseWriter, r *http.Request) {
	//get json payload
	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(r, payload); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, err)
	}

	//check if user exists
	user, err := h.store.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, err)
	}
}
