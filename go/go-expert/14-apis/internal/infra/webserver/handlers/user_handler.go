package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Lucioschenkel/goexpert/14-apis/internal/dto"
	"github.com/Lucioschenkel/goexpert/14-apis/internal/entity"
	"github.com/Lucioschenkel/goexpert/14-apis/internal/infra/database"
	"github.com/go-chi/jwtauth"
)

type Error struct {
	Message string `json:"message"`
}

type UserHandler struct {
	UserDB database.UserInterface
}

func NewUserHandler(userDB database.UserInterface) *UserHandler {
	return &UserHandler{
		UserDB: userDB,
	}
}

// GetJWT godoc
// @Summary 		Get a user JWT
// @Description Get a user JWT
// @Tags 				users
// @Accept 			json
// @Produce 		json
// @Param 			request 			body 			dto.GetJWTInput 	true 	"user credentials"
// @Success 		200 					{object} 	dto.GetJWTOutput
// @Failure 		400
// @Failure 		401
// @Router 			/users/login 	[post]
func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	jwtExpiresIn := r.Context().Value("JwtExpiresIn").(int)
	var user dto.GetJWTInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	u, err := h.UserDB.FindByEmail(user.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if !u.ValidatePassword(user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	_, tokenString, _ := jwt.Encode(map[string]interface{}{
		"sub": u.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(jwtExpiresIn)).Unix(),
	})
	accessToken := dto.GetJWTOutput{
		AccessToken: tokenString,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}

// Create user godoc
// @Summary 		Create a user
// @Description Create a user
// @Tags 				users
// @Accept 			json
// @Produce 		json
// @Param 			request 			body 			dto.CreateUserInput true "user request"
// @Success 		201
// @Failure 		500 					{object} 	Error
// @Router 			/users 				[post]
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}
	err = h.UserDB.Create(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
