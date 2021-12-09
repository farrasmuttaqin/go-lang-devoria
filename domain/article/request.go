package article

// CreateArticleRequest is model for creating article.
type CreateArticleRequest struct {
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Content  string `json:"content"`
	Author   int64  `json:"author"`
}

// EditArticleRequest is model for modified article.
type EditArticleRequest struct {
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Content  string `json:"content"`
	Author   int64  `json:"author"`
}

// PublishArticleRequest is model for modified article.
type PublishArticleRequest struct {
	Id     int64 `json:"id"`
	Author int64 `json:"author"`
}

// DraftArticleRequest is model for modified article.
type ArchiveArticleRequest struct {
	Id     int64 `json:"id"`
	Author int64 `json:"author"`
}
