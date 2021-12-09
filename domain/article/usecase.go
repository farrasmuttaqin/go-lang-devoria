package article

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/sangianpatrick/devoria-article-service/exception"
	"github.com/sangianpatrick/devoria-article-service/response"
)

type ArticleUseCase interface {
	Create(ctx context.Context, params CreateArticleRequest) (resp response.Response)
	GetByID(ctx context.Context, r *http.Request) (resp response.Response)
	// Publish(ctx context.Context, params PublishArticleRequest) (resp response.Response)
	// Update(ctx context.Context, params EditArticleRequest) (resp response.Response)
	// Archive(ctx context.Context, params ArchiveArticleRequest) (resp response.Response)
}

type articleUsecaseImpl struct {
	location   *time.Location
	repository ArticleRepository
}

func NewArticleUsecase(
	location *time.Location,
	repository ArticleRepository,
) ArticleUseCase {
	return &articleUsecaseImpl{
		location:   location,
		repository: repository,
	}
}

func (u *articleUsecaseImpl) Create(ctx context.Context, params CreateArticleRequest) (resp response.Response) {
	fmt.Println(params.Title)
	_, err := u.repository.FindByTitle(ctx, params.Title)
	if err == nil {
		return response.Error(response.StatusConflicted, nil, exception.ErrConflicted)
	}

	newArticle := Article{}
	newArticle.Author = params.Author
	newArticle.Title = params.Title
	newArticle.Subtitle = params.Subtitle
	newArticle.Content = params.Content
	newArticle.Status = ArticleStatus(ArticleStatusArchive)
	newArticle.CreatedAt = time.Now().In(u.location)

	ID, err := u.repository.Save(ctx, newArticle)
	if err != nil {
		return response.Error(response.StatusUnexpectedError, nil, exception.ErrInternalServer)
	}
	newArticle.ID = ID

	createArticleResponse := ArticleResponse{}
	createArticleResponse.Data = newArticle

	return response.Success(response.StatusCreated, createArticleResponse)
}

func (u *articleUsecaseImpl) GetByID(ctx context.Context, r *http.Request) (resp response.Response) {

	keys, ok := r.URL.Query()["id"]

	if !ok || len(keys[0]) < 1 {
		return response.Error("Url Param 'id' is missing", nil, exception.ErrNotFound)
	}

	id := keys[0]

	//convert id from string to integer
	articleID, err := strconv.ParseInt(id, 0, 8)
	if err != nil {
		return response.Error("id must be integer", nil, exception.ErrDataType)
	}
	fmt.Println(articleID)
	article, err := u.repository.FindByID(ctx, articleID)

	if err != nil {
		if err == exception.ErrNotFound {
			return response.Error(response.StatusNotFound, nil, exception.ErrBadRequest)
		}
		return response.Error(response.StatusUnexpectedError, nil, exception.ErrInternalServer)
	}

	articleResponse := ArticleResponse{}
	articleResponse.Data = article

	return response.Success(response.StatusOK, articleResponse)
}

// func (u *articleUsecaseImpl) Publish(ctx context.Context, params PublishArticleRequest) (resp response.Response) {
// 	_, err := u.repository.FindByID(ctx, params.Id)
// 	if err == nil {
// 		return response.Error(response.StatusConflicted, nil, exception.ErrConflicted)
// 	}

// 	if err != exception.ErrNotFound {
// 		return response.Error(response.StatusUnexpectedError, nil, exception.ErrInternalServer)
// 	}

// 	publishArticle := Article{}
// 	publishArticle.LastModifiedAt = time.Now()
// 	publishArticle.ID = params.Id

// 	article, err := u.repository.Save(ctx, publishArticle)
// 	if err != nil {
// 		return response.Error(response.StatusUnexpectedError, nil, exception.ErrInternalServer)
// 	}

// 	publishArticleResponse := ArticleResponse{}
// 	publishArticleResponse.Data = article

// 	return response.Success(response.StatusCreated, publishArticleResponse)
// }

// func (u *articleUsecaseImpl) Archive(ctx context.Context, params ArchiveArticleRequest) (resp response.Response) {
// 	fmt.Println(u.globalIV)
// 	_, err := u.repository.FindByID(ctx, params.Id)
// 	if err == nil {
// 		return response.Error(response.StatusConflicted, nil, exception.ErrConflicted)
// 	}

// 	if err != exception.ErrNotFound {
// 		return response.Error(response.StatusUnexpectedError, nil, exception.ErrInternalServer)
// 	}

// 	article, err := u.repository.Archive(ctx, params.Id)
// 	if err != nil {
// 		return response.Error(response.StatusUnexpectedError, nil, exception.ErrInternalServer)
// 	}

// 	archiveArticleResponse := ArticleResponse{}
// 	archiveArticleResponse.Data = article

// 	return response.Success(response.StatusCreated, archiveArticleResponse)
// }
