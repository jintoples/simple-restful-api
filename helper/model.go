package helper

import (
	"github.com/jintoples/simple-restful-api/model/domain"
	"github.com/jintoples/simple-restful-api/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		category.Id,
		category.Name,
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var CategoryResponses []web.CategoryResponse
	for _, category := range categories {
		CategoryResponses = append(CategoryResponses, ToCategoryResponse(category))
	}
	return CategoryResponses
}
