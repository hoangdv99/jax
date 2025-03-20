package main

import (
	"net/http"
	"strconv"
	"strings"
)

func (app *application) getProductsHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	storeIdsStr := queryParams.Get("storeIds")
	var storeIds []int64
	if storeIdsStr != "" {
		storeIdsStrSlice := strings.Split(storeIdsStr, ",")
		for _, idStr := range storeIdsStrSlice {
			id, err := strconv.ParseInt(idStr, 10, 64)
			if err != nil {
				http.Error(w, "Invalid store ID", http.StatusBadRequest)
				return
			}
			storeIds = append(storeIds, id)
		}
	}
	page := queryParams.Get("page")
	limit := queryParams.Get("limit")

	user := app.contextGetUser(r)

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 1
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 10
	}

	products, err := app.models.Products.GetProducts(user.Id, storeIds, pageInt, limitInt)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	app.writeJSON(w, http.StatusOK, envelop{"products": products}, nil)
}
