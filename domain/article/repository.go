package article

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/sangianpatrick/devoria-article-service/exception"
)

type ArticleRepository interface {
	Save(ctx context.Context, article Article) (ID int64, err error)
	// Update(ctx context.Context, ID int64, updatedArticle Article) (err error)
	// Publish(ctx context.Context, ID int64) (err error)
	// Archive(ctx context.Context, ID int64, updatedArticle Article) (err error)
	FindByTitle(ctx context.Context, Title string) (article Article, err error)
	FindByID(ctx context.Context, ID int64) (article Article, err error)
	// FindMany(ctx context.Context) (bunchOfArticles []Article, err error)
	// FindManySpecificProfile(ctx context.Context, articleID int64) (bunchOfArticles []Article, err error)
}

type articleRepositoryImpl struct {
	db        *sql.DB
	tableName string
}

func NewArticleRepository(db *sql.DB, tableName string) ArticleRepository {
	return &articleRepositoryImpl{
		db:        db,
		tableName: tableName,
	}
}

func (r *articleRepositoryImpl) Save(ctx context.Context, article Article) (ID int64, err error) {
	command := fmt.Sprintf("INSERT INTO %s (title, author, subtitle, content, status, createdAt) VALUES (?, ?, ?, ?, ?, ?)", r.tableName)
	stmt, err := r.db.PrepareContext(ctx, command)
	if err != nil {
		log.Println(err)
		return
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(
		ctx,
		article.Title,
		article.Author,
		article.Subtitle,
		article.Content,
		article.Status,
		article.CreatedAt,
	)

	if err != nil {
		log.Println(err)
		return
	}

	ID, _ = result.LastInsertId()

	return
}

func (r *articleRepositoryImpl) FindByID(ctx context.Context, ID int64) (article Article, err error) {
	query := fmt.Sprintf(`SELECT id,title,subtitle,content,createdAt,lastModifiedAt,publishedAt FROM %s WHERE id = ?`, r.tableName)
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		log.Println(err)
		err = exception.ErrInternalServer
		return
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, ID)

	err = row.Scan(
		&article.ID,
		&article.Title,
		&article.Subtitle,
		&article.Content,
		&article.CreatedAt,
		&article.LastModifiedAt,
		&article.PublishedAt,
	)

	if err != nil {
		log.Println(err)
		err = exception.ErrNotFound
		return
	}

	return
}

// func (r *articleRepositoryImpl) Update(ctx context.Context, ID int64, updatedArticle Article) (ID int64, err error) {
// 	command := fmt.Sprintf(`UPDATE %s SET title = ?, subtitle = ?, content = ?, lastModifiedAt = ? WHERE id = ?`, r.tableName)
// 	stmt, err := r.db.PrepareContext(ctx, command)
// 	if err != nil {
// 		log.Println(err)
// 		err = exception.ErrInternalServer
// 		return
// 	}
// 	defer stmt.Close()

// 	result, err := stmt.ExecContext(
// 		ctx,
// 		updatedArticle.Title,
// 		updatedArticle.Subtitle,
// 		updatedArticle.Content,
// 		updatedArticle.LastModifiedAt,
// 	)

// 	if err != nil {
// 		log.Println(err)
// 		err = exception.ErrInternalServer
// 		return
// 	}

// 	rowsAffected, _ := result.RowsAffected()
// 	if rowsAffected < 1 {
// 		err = exception.ErrNotFound
// 		return
// 	}

// 	return
// }

// func (r *articleRepositoryImpl) Publish(ctx context.Context, ID int64, updatedArticle Article) (err error) {
// 	command := fmt.Sprintf(`UPDATE %s SET status = %s, lastModifiedAt = ? WHERE id = ?`, r.tableName, ArticleStatusPublished)
// 	stmt, err := r.db.PrepareContext(ctx, command)
// 	if err != nil {
// 		log.Println(err)
// 		err = exception.ErrInternalServer
// 		return
// 	}
// 	defer stmt.Close()

// 	result, err := stmt.ExecContext(
// 		ctx,
// 		updatedArticle.LastModifiedAt,
// 		updatedArticle.ID,
// 	)

// 	if err != nil {
// 		log.Println(err)
// 		err = exception.ErrInternalServer
// 		return
// 	}

// 	rowsAffected, _ := result.RowsAffected()
// 	if rowsAffected < 1 {
// 		err = exception.ErrNotFound
// 		return
// 	}

// 	return
// }

// func (r *articleRepositoryImpl) Archive(ctx context.Context, ID int64, updatedArticle Article) (err error) {
// 	command := fmt.Sprintf(`UPDATE %s SET status = %s, lastModifiedAt = ? WHERE id = ?`, r.tableName, ArticleStatusArchive)
// 	stmt, err := r.db.PrepareContext(ctx, command)
// 	if err != nil {
// 		log.Println(err)
// 		err = exception.ErrInternalServer
// 		return
// 	}
// 	defer stmt.Close()

// 	result, err := stmt.ExecContext(
// 		ctx,
// 		updatedArticle.Status,
// 		updatedArticle.LastModifiedAt,
// 	)

// 	if err != nil {
// 		log.Println(err)
// 		err = exception.ErrInternalServer
// 		return
// 	}

// 	rowsAffected, _ := result.RowsAffected()
// 	if rowsAffected < 1 {
// 		err = exception.ErrNotFound
// 		return
// 	}

// 	return
// }

// func (r *articleRepositoryImpl) FindByID(ctx context.Context, id int64) (article Article, err error) {
// 	query := fmt.Sprintf(`SELECT id, title, subtitle, content, createdAt, publishedAt, lastModified FROM %s WHERE id = ?`, r.tableName)
// 	stmt, err := r.db.PrepareContext(ctx, query)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	defer stmt.Close()

// 	row := stmt.QueryRowContext(ctx, id)

// 	var lastModifiedAt sql.NullTime

// 	err = row.Scan(
// 		&article.ID,
// 		&article.Title,
// 		&article.Subtitle,
// 		&article.Content,
// 		&article.CreatedAt,
// 		&article.PublishedAt,
// 		&lastModifiedAt,
// 	)

// 	if err != nil {
// 		log.Println(err)
// 		err = exception.ErrNotFound
// 		return
// 	}

// 	if lastModifiedAt.Valid {
// 		article.LastModifiedAt = &lastModifiedAt.Time
// 	}

// 	return
// }

func (r *articleRepositoryImpl) FindByTitle(ctx context.Context, title string) (article Article, err error) {
	query := fmt.Sprintf(`SELECT id,title,subtitle,content,createdAt,lastModifiedAt,publishedAt FROM %s WHERE title = ?`, r.tableName)
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		log.Println(err)
		err = exception.ErrInternalServer
		return
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, title)

	err = row.Scan(
		&article.ID,
		&article.Title,
		&article.Subtitle,
		&article.Content,
		&article.CreatedAt,
		&article.LastModifiedAt,
		&article.PublishedAt,
	)

	if err != nil {
		log.Println(err)
		err = exception.ErrNotFound
		return
	}

	return
}
