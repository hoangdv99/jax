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

	existedTag, err := app.models.Tags.GetByName(tag.Name)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if existedTag != nil {
		v.AddError("name", "duplicated tag")
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	id, err := app.models.Tags.Insert(tag)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusCreated, envelop{"id": id}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getListTagHandler(w http.ResponseWriter, r *http.Request) {
	user := app.contextGetUser(r)
	tags, err := app.models.Tags.ListByUserId(user.Id)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelop{"tags": tags}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) updateTagHandler(w http.ResponseWriter, r *http.Request) {
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
		Name *string `json:"name"`
	}
	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()
	existedTag, err := app.models.Tags.GetByName(*input.Name)
	if existedTag != nil && existedTag.Id != id {
		v.AddError("name", "duplicated tag")
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	if input.Name != nil {
		tag.Name = *input.Name
	}

	v.Check(tag.Name != "", "name", "name is required")
	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Tags.Update(tag)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelop{"message": "OK"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteTagHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIdParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	_, err = app.models.Tags.GetById(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.models.Tags.Delete(id)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelop{"message": "OK"}, nil)
}
