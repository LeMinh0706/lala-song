package search

import "github.com/gofiber/fiber/v2"

func NewSearchRouter(f fiber.Router, service ISearchService) {
	sr := NewSearchController(service)
	searchGroup := f.Group("/search")
	{
		searchGroup.Get("/fts", sr.SearchByLyrics)
		searchGroup.Get("/song", sr.SearchBySong)
	}
}
