package main

import (
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
		app.badRequestResponse(w, r, err)
		return
	}

	user := app.contextGetUser(r)

	tag := &data.Tag{
		UserId: user.Id,
		Name:   input.Name,
	}

	v := validator.New()
	v.Check(tag.Name != "", "name", "must be provided")
	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	id, err := app.models.Tags.Insert(tag)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusCreated, envelop{"id": id}, nil)
}

func (app *application) getListTagHandler(w http.ResponseWriter, r *http.Request) {
	user := app.contextGetUser(r)
	tags, err := app.models.Tags.ListByUserId(user.Id)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.writeJSON(w, http.StatusOK, envelop{"tags": tags}, nil)
}
