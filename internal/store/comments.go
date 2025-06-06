package store

import (
	"context"
	"database/sql"
)

type Comment struct {
	ID        int64  `json:"id"`
	PostID    int64  `json:"post_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	User      *User  `json:"user"`
}

type CommentStore struct {
	db *sql.DB
}

func (s *CommentStore) Create(ctx context.Context, comment *Comment) error {
	query := `
	INSERT INTO comments (post_id, content, user_id)
	VALUES ($1, $2, $3) RETURNING id, created_at
	`
	err := s.db.QueryRowContext(ctx,
		query,
		comment.PostID,
		comment.Content,
		comment.User.ID, // Assuming User.ID is set before calling Create
	).Scan(
		&comment.ID,
		&comment.CreatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

func (s *CommentStore) GetByPostId(ctx context.Context, postID int64) ([]Comment, error) {
	query := `
	SELECT c.id, c.post_id, c.content, c.created_at, users.username, users.id, users.email
	FROM comments c 
	JOIN users ON c.user_id = users.id
	WHERE c.post_id = $1
	ORDER BY c.created_at DESC
	`

	rows, err := s.db.QueryContext(ctx, query, postID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	comments := []Comment{}
	for rows.Next() {
		var c Comment
		c.User = &User{}
		err := rows.Scan(
			&c.ID,
			&c.PostID,
			&c.Content,
			&c.CreatedAt,
			&c.User.Username,
			&c.User.ID,
			&c.User.Email,
		)
		if err != nil {
			return nil, err
		}

		comments = append(comments, c)
	}

	return comments, nil
}
