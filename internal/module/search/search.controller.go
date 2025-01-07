package search

import (
	"github.com/LeMinh0706/lala-song/internal/handler"
	"github.com/LeMinh0706/lala-song/res"
	"github.com/gofiber/fiber/v2"
)

type SearchController struct {
	service ISearchService
}

func NewSearchController(service ISearchService) *SearchController {
	return &SearchController{
		service: service,
	}
}

// Song godoc
// @Summary      Tìm kiếm bài hát thông qua lời
// @Description  Tìm kiếm bài hát thông qua lời
// @Tags         Search
// @Accept       json
// @Produce      json
// @Param        lyric query string true "Song's lyric"
// @Param        page query int true "Page"
// @Param        page_size query int true "Page Size"s
// @Success      200  {object}  []db.GetSongRow
// @Failure      500  {object}  res.ErrSwaggerJson
// @Router       /search/fts [get]
func (s *SearchController) SearchByLyrics(f *fiber.Ctx) error {
	lyric := f.Query("lyric")
	pageStr := f.Query("page")
	pageSizeStr := f.Query("page_size")

	page, pageSize, numErr := handler.CheckQuery(f, pageStr, pageSizeStr)
	if numErr != 0 {
		return res.ErrorResponse(f, numErr)
	}
	list, err := s.service.SearchLyric(f.Context(), lyric, page, pageSize)
	if err != nil {
		return res.ErrorNonKnow(f, err.Error())
	}
	return res.SuccessResponse(f, 200, list)
}

// Song godoc
// @Summary      Tìm theo tên bài hát
// @Description  Tìm theo tên bài hát
// @Tags         Search
// @Accept       json
// @Produce      json
// @Param        name query string true "Song's name"
// @Param        page query int true "Page"
// @Param        page_size query int true "Page Size"s
// @Success      200  {object}  []db.GetSongRow
// @Failure      500  {object}  res.ErrSwaggerJson
// @Router       /search/song [get]
func (s *SearchController) SearchBySong(f *fiber.Ctx) error {
	name := f.Query("name")
	pageStr := f.Query("page")
	pageSizeStr := f.Query("page_size")

	page, pageSize, numErr := handler.CheckQuery(f, pageStr, pageSizeStr)
	if numErr != 0 {
		return res.ErrorResponse(f, numErr)
	}
	list, err := s.service.SearchSong(f.Context(), name, page, pageSize)
	if err != nil {
		return res.ErrorNonKnow(f, err.Error())
	}
	return res.SuccessResponse(f, 200, list)
}
