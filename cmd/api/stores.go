package main

import (
	"net/http"

	"jax.hoangdv99/internal/validator"
)

func (app *application) getStorePlatformHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		url string
	}

	v := validator.New()

	qs := r.URL.Query()
	input.url = app.readString(qs, "url", "")

	v.Check(input.url != "", "url", "url required")
	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	platform, err := app.models.Stores.GetStorePlatform(input.url)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelop{"platform": platform}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
