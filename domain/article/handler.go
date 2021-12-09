package article

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/sangianpatrick/devoria-article-service/middleware"
	"github.com/sangianpatrick/devoria-article-service/response"
)

type ArticleHTTPHandler struct {
	Validate *validator.Validate
	Usecase  ArticleUseCase
}

func NewArticleHTTPHandler(
	router *mux.Router,
	basicAuthMiddleware middleware.RouteMiddleware,
	validate *validator.Validate,
	usecase ArticleUseCase,
) {
	handler := &ArticleHTTPHandler{
		Validate: validate,
		Usecase:  usecase,
	}

	router.HandleFunc("/v1/article/create", basicAuthMiddleware.Verify(handler.Create)).Methods(http.MethodPost)
	router.HandleFunc("/v1/article/get", basicAuthMiddleware.Verify(handler.GetByID)).Methods(http.MethodGet)
}

func (handler *ArticleHTTPHandler) Create(w http.ResponseWriter, r *http.Request) {
	var resp response.Response
	var params CreateArticleRequest
	var ctx = r.Context()

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		resp = response.Error(response.StatusUnprocessabelEntity, nil, err)
		resp.JSON(w)
		return
	}

	err = handler.Validate.StructCtx(ctx, params)
	if err != nil {
		resp = response.Error(response.StatusInvalidPayload, nil, err)
		resp.JSON(w)
		return
	}

	resp = handler.Usecase.Create(ctx, params)
	resp.JSON(w)
}

func (handler *ArticleHTTPHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	var resp response.Response
	var ctx = r.Context()

	resp = handler.Usecase.GetByID(ctx, r)
	resp.JSON(w)
}

// func (handler *AccountHTTPHandler) Create(w http.ResponseWriter, r *http.Request) {
// 	var resp response.Response
// 	var params CreateArticleRequest
// 	var ctx = r.Context()

// 	err := json.NewDecoder(r.Body).Decode(&params)
// 	if err != nil {
// 		resp = response.Error(response.StatusUnprocessabelEntity, nil, err)
// 		resp.JSON(w)
// 		return
// 	}

// 	err = handler.Validate.StructCtx(ctx, params)
// 	if err != nil {
// 		resp = response.Error(response.StatusInvalidPayload, nil, err)
// 		resp.JSON(w)
// 		return
// 	}

// 	resp = handler.Usecase.Create(ctx, params)
// 	resp.JSON(w)
// }
