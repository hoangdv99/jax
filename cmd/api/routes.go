package main

import (
	"expvar"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	publicRouter := httprouter.New()
	publicRouter.NotFound = http.HandlerFunc(app.notFoundResponse)
	publicRouter.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	publicRouter.HandlerFunc(http.MethodGet, "/healthcheck", app.healthcheckHandler)
	publicRouter.Handler(http.MethodGet, "/debug/vars", expvar.Handler())

	publicRouter.HandlerFunc(http.MethodPost, "/users", app.registerUserHandler)
	publicRouter.HandlerFunc(http.MethodPut, "/users/activation", app.activateUserHandler)

	publicRouter.HandlerFunc(http.MethodPost, "/tokens/activation", app.createActivationTokenHandler)

	publicRouter.HandlerFunc(http.MethodPost, "/login", app.loginHandler)
	publicRouter.HandlerFunc(http.MethodPost, "/logout", app.logoutHandler)

	// Protected routes (start with /v1) which need authentication
	protectedRouter := httprouter.New()
	protectedRouter.NotFound = http.HandlerFunc(app.notFoundResponse)
	protectedRouter.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	protectedRouter.HandlerFunc(http.MethodGet, "/v1/users", app.getListUserHandler)
	protectedRouter.HandlerFunc(http.MethodGet, "/v1/current-user", app.getCurrentUserHandler)

	protectedRouter.HandlerFunc(http.MethodPost, "/v1/store", app.addStoreHandler)
	protectedRouter.HandlerFunc(http.MethodGet, "/v1/stores", app.listStoreHandler)
	protectedRouter.HandlerFunc(http.MethodPut, "/v1/store/:id", app.updateStoreHandler)
	protectedRouter.HandlerFunc(http.MethodDelete, "/v1/store/:id", app.deleteStoreHandler)

	protectedRouter.HandlerFunc(http.MethodPost, "/v1/tag", app.createTagHandler)
	protectedRouter.HandlerFunc(http.MethodGet, "/v1/tags", app.getListTagHandler)
	protectedRouter.HandlerFunc(http.MethodPatch, "/v1/tag/:id", app.updateTagHandler)
	protectedRouter.HandlerFunc(http.MethodDelete, "/v1/tag/:id", app.deleteTagHandler)

	protectedRouter.HandlerFunc(http.MethodGet, "/v1/products", app.getProductsHandler)

	router := http.NewServeMux()
	router.Handle("/", publicRouter)
	router.Handle("/v1/", app.authenticate(protectedRouter))

	return app.recoverPanic(app.enableCORS(router))
}
