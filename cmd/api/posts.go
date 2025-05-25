package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/omkar-nag/socialapp/internal/store"
)

type CreatePostPayload struct {
	Title   string `json:"title" validate:"required,max=100`
	Content string `json:"content" validate:"required,max=1000"`
}

func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {

	userId := 1
	var payload CreatePostPayload

	if err := readJSON(w, r, &payload); err != nil {
		app.badRequest(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		fmt.Println(Validate.Struct(payload))
		app.badRequest(w, r, err)
		return
	}

	post := &store.Post{
		Title:   payload.Title,
		Content: payload.Content,
		UserID:  int64(userId),
		Tags:    []string{},
	}

	ctx := r.Context()
	if err := app.store.Posts.Create(ctx, post); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := writeJSON(w, http.StatusCreated, post); err != nil {
		app.internalServerError(w, r, err)
		return
	}

}

func (app *application) getPostHandler(w http.ResponseWriter, r *http.Request) {
	postIDStr := chi.URLParam(r, "postID")

	ctx := r.Context()

	postID, err := strconv.ParseInt(postIDStr, 10, 64)

	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	post, err := app.store.Posts.GetById(ctx, postID)

	if err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			writeJSONError(w, http.StatusNotFound, err.Error())
		default:
			app.internalServerError(w, r, err)
		}
		return
	}

	comments, err := app.store.Comments.GetByPostId(ctx, postID)
	if err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			post.Comments = []store.Comment{}
		default:
			app.internalServerError(w, r, err)
			return
		}
	}
	post.Comments = comments

	if err := writeJSON(w, http.StatusOK, post); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
