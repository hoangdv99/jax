package main

import (
	"errors"
	"net/http"

	"jax.hoangdv99/internal/data"
	"jax.hoangdv99/internal/validator"
)

func (app *application) createTagHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name string `json:"name"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	v := validator.New()
	v.Check(input.Name != "", "name", "name is required")
	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	user := app.contextGetUser(r)

	tag := &data.Tag{
		Name:   input.Name,
		UserId: user.Id,
	}

	id, err := app.models.Tags.Insert(tag)
	tag.Id = id

	err = app.writeJSON(w, http.StatusOK, envelop{"data": tag}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getTagsHandler(w http.ResponseWriter, r *http.Request) {
	user := app.contextGetUser(r)
	tags, err := app.models.Tags.List(user.Id)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelop{"tags": tags}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) editTagHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIdParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	tag, err := app.models.Tags.GetById(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	var input struct {
		Name string `json:"name"`
	}
	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	tag.Name = input.Name
	v := validator.New()
	v.Check(tag.Name != "", "name", "name is required")

	user := app.contextGetUser(r)
	tag.UserId = user.Id

	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

}
