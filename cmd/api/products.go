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

func (app *application) getCollectionProductsHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	collectionIdStr := queryParams.Get("collectionId")
	handle := queryParams.Get("handle")
	page := queryParams.Get("page")
	if page == "" {
		page = "1"
	}
	limit := queryParams.Get("limit")
	if limit == "" {
		limit = "30"
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 1
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 30
	}

	var collectionIdInt int64
	if collectionIdStr != "" {
		collectionIdInt, err = strconv.ParseInt(collectionIdStr, 10, 64)
		if err != nil {
			app.badRequestResponse(w, r, err)
			return
		}
	}

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

	if collectionIdStr == "" && handle == "" {
		app.badRequestResponse(w, r, err)
		return
	}

	products, err := app.models.Products.GetCollectionProducts(*store, collectionIdInt, handle, pageInt, limitInt)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelop{"products": products}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
