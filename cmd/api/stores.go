package main

import (
	"errors"
	"net/http"

	"jax.hoangdv99/internal/data"
	"jax.hoangdv99/internal/validator"
)

func (app *application) addStoreHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Url    string  `json:"url"`
		TagIds []int64 `json:"tagIds"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()
	v.Check(input.Url != "", "url", "must be provided")
	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	platform, err := app.models.Stores.GetStorePlatform(input.Url)
	if err != nil {
		app.badRequestResponse(w, r, errors.New("Invalid url or unsupported platform"))
		return
	}

	store, err := app.models.Stores.GetByUrl(input.Url)
	user := app.contextGetUser(r)
	if store == nil {
		store := &data.Store{
			Url:      input.Url,
			Platform: platform,
			IsActive: true,
		}
		err := app.models.Stores.Insert(user, store, input.TagIds)
		if err != nil {
			switch {
			case errors.Is(err, data.ErrDuplicatedStore):
				app.badRequestResponse(w, r, err)
			default:
				app.serverErrorResponse(w, r, err)
			}
			return
		}
	} else {
		err := app.models.Stores.UpdateUserStore(user, store, input.TagIds)
		if err != nil {
			switch {
			case errors.Is(err, data.ErrDuplicatedStore):
				app.badRequestResponse(w, r, err)
			default:
				app.serverErrorResponse(w, r, err)
			}
			return
		}
	}

	err = app.writeJSON(w, http.StatusOK, envelop{"message": "OK"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) listStoreHandler(w http.ResponseWriter, r *http.Request) {
	user := app.contextGetUser(r)
	stores, err := app.models.Stores.List(user.Id)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelop{"data": stores}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) updateStoreHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		TagIds []int64 `json:"tagIds"`
	}

	storeId, err := app.readIdParam(r)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	err = app.models.Stores.UpdateStoreTags(storeId, input.TagIds)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelop{"message": "OK"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteStoreHandler(w http.ResponseWriter, r *http.Request) {
	storeId, err := app.readIdParam(r)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	user := app.contextGetUser(r)

	err = app.models.Stores.DeleteStore(user.Id, storeId)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelop{"message": "OK"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getCollectionsHandler(w http.ResponseWriter, r *http.Request) {
	storeId, err := app.readIdParam(r)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	store, err := app.models.Stores.GetById(storeId)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	collections, err := app.models.Stores.GetCollections(*store)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelop{"data": collections}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
