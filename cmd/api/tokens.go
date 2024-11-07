package main

import (
	"errors"
	"net/http"
	"time"

	"jax.hoangdv99/internal/constant"
	"jax.hoangdv99/internal/data"
	"jax.hoangdv99/internal/validator"
)

func (app *application) createActivationToken(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email string `json:"email"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
	}

	v := validator.New()

	if data.ValidateEmail(v, input.Email); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	user, err := app.models.Users.GetByEmail(input.Email)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			v.AddError("email", "no matching email address found")
			app.failedValidationResponse(w, r, v.Errors)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	if user.Status != constant.USER_STATUS_WAITING_ACTIVATION {
		v.AddError("email", "cannot active this user")
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	token, err := app.models.Tokens.New(user.Id, 3*24*time.Hour, data.ScopeActivation)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.background(func() {
		data := map[string]interface{}{
			"activationToken": token.Plaintext,
		}
		err = app.mailer.Send(user.Email, "token_activation.html", data)
		if err != nil {
			app.logger.PrintError(err, nil)
		}
	})

	err = app.writeJSON(
		w,
		http.StatusAccepted,
		envelop{"message": "an email will be sent to you containing activation instructions"},
		nil,
	)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
