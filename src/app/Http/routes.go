package Routes

import (
	. "app/Http/Controllers"

	_routing "github.com/qiangxue/fasthttp-routing"
)

func Handle(router *_routing.Router) {
	api := router.Group("/api")

	userController := UserController{}
	api.Get("/users", userController.Index)
	api.Get("/users/<userId>", userController.Show)
	api.Post("/users", userController.Store)
	api.Put("/users/<userId>", userController.Update)

	countryController := CountryController{}
	api.Get("/countries", countryController.Index)
	api.Get("/countries/<countryId>", countryController.Show)
	api.Post("/countries", countryController.Store)
	api.Put("/countries/<countryId>", countryController.Update)
}
