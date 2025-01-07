package song

import (
	"fmt"
	"path/filepath"
	"strconv"
	"time"

	"github.com/LeMinh0706/lala-song/internal/db"
	"github.com/LeMinh0706/lala-song/internal/handler"
	"github.com/LeMinh0706/lala-song/res"
	"github.com/LeMinh0706/lala-song/util"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type SongController struct {
	service ISongService
}

func NewSongController(service ISongService) *SongController {
	return &SongController{
		service: service,
	}
}

// Songs godoc
// @Summary      Create song
// @Description  Create song
// @Tags         Songs
// @Accept       multipart/form-data
// @Produce      json
// @Param        name formData string true "name"
// @Param        album_id formData int true "album id"
// @Param        mp3 formData file true "Mp3 file"
// @Param        lyric formData file true "Lyric file"
// @Security BearerAuth
// @Success      201  {object} 	db.CreateSongRow
// @Failure      500  {object}  res.ErrSwaggerJson
// @Router       /songs [post]
func (s *SongController) CreateSong(f *fiber.Ctx) error {
	songfile, err := f.FormFile("mp3")
	if err != nil {
		return res.ErrorNonKnow(f, err.Error())
	}
	const maxSize = 16 << 20
	if songfile.Size > maxSize {
		return res.ErrorResponse(f, 41300)
	}
	if songfile != nil {
		if !util.Mp3Check(songfile.Filename) {
			return res.ErrorResponse(f, res.ErrBadRequestMime)
		}
	}
	filename := fmt.Sprintf("upload/%s/%d%s", "song", time.Now().Unix(), filepath.Ext(songfile.Filename))

	err = f.SaveFile(songfile, filename)
	if err != nil {
		return res.ErrorNonKnow(f, err.Error())
	}

	lyric, err := f.FormFile("lyric")
	if err != nil {
		return res.ErrorNonKnow(f, err.Error())
	}
	if lyric.Size > maxSize {
		return res.ErrorResponse(f, 41300)
	}
	if lyric != nil {
		if !util.LyricCheck(lyric.Filename) {
			return res.ErrorResponse(f, res.ErrBadRequestMime)
		}
	}
	lyricname := fmt.Sprintf("upload/%s/%d%s", "lyrics", time.Now().Unix(), filepath.Ext(lyric.Filename))

	err = f.SaveFile(lyric, lyricname)
	if err != nil {
		return res.ErrorNonKnow(f, err.Error())
	}
	name := f.FormValue("name")

	idStr := f.FormValue("album_id")
	album_id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return res.ErrorResponse(f, res.ErrBadRequestId)
	}

	uuid, _ := uuid.NewRandom()

	song, err := s.service.CreateSong(f.Context(), uuid, name, filename, lyricname, album_id)
	if err != nil {
		return res.ErrorNonKnow(f, err.Error())
	}
	return res.SuccessResponse(f, 201, song)
}

// song godoc
// @Summary      Get song with id
// @Description  Get song with id
// @Tags         Songs
// @Accept       json
// @Produce      json
// @Param        id path string true "ID"
// @Success      200  {object}  SongResponse
// @Failure      500  {object}  res.ErrSwaggerJson
// @Router       /songs/{id} [get]
func (s *SongController) GetSongById(f *fiber.Ctx) error {
	idStr := f.Params("id")

	uuid, err := uuid.Parse(idStr)

	if err != nil {
		return res.ErrorResponse(f, res.ErrBadRequestId)
	}

	genre, err := s.service.GetSong(f.Context(), uuid)
	if err != nil {
		return res.ErrorResponse(f, res.ErrGenreNotFound)
	}
	return res.SuccessResponse(f, 200, genre)
}

// song godoc
// @Summary      Thêm nghệ sĩ khác vào bài hát
// @Description  Thêm nghệ sĩ khác vào bài hát
// @Tags         Songs
// @Accept       json
// @Produce      json
// @Param        request body db.AddSongSingerParams true "request"
// @Security BearerAuth
// @Success      200  {object}  SongResponse
// @Failure      500  {object}  res.ErrSwaggerJson
// @Router       /songs/feature [post]
func (s *SongController) AddFeatureSong(f *fiber.Ctx) error {

	var req db.AddSongSingerParams

	if err := f.BodyParser(&req); err != nil {
		return res.ErrorResponse(f, 40000)
	}

	feature, err := s.service.AddFeatureSong(f.Context(), req.SongID, req.SingerID)
	if err != nil {
		return res.ErrorResponse(f, res.ErrAddSingerSong)
	}

	return res.SuccessResponse(f, 200, feature)
}

// song godoc
// @Summary      Thêm thể loại nhạc cho bài hát khác vào bài hát
// @Description  Thêm thể loại nhạc cho bài hát khác vào bài hát
// @Tags         Songs
// @Accept       json
// @Produce      json
// @Param        request body db.AddSongGenreParams true "request"
// @Security BearerAuth
// @Success      200  {object}  SongResponse
// @Failure      500  {object}  res.ErrSwaggerJson
// @Router       /songs/genre [post]
func (s *SongController) AddGenre(f *fiber.Ctx) error {

	var req db.AddSongGenreParams

	if err := f.BodyParser(&req); err != nil {
		return res.ErrorResponse(f, 40000)
	}

	feature, err := s.service.AddGenreSong(f.Context(), req.SongID, req.GenresID)
	if err != nil {
		return res.ErrorResponse(f, res.ErrAddGenreSong)
	}

	return res.SuccessResponse(f, 200, feature)
}

// Song godoc
// @Summary      Get list songs
// @Description  Get list songs with page and page size (Limit-Offset)
// @Tags         Songs
// @Accept       json
// @Produce      json
// @Param        filter query string false "Your filter"
// @Param        singer query int false "Singer ID"
// @Param        album query int false "Album ID"
// @Param        genres query int false "Genre ID"
// @Param        page query int true "Page"
// @Param        page_size query int true "Page Size"
// @Success      200  {object}  []db.GetSongRow
// @Failure      500  {object}  res.ErrSwaggerJson
// @Router       /songs [get]
func (s *SongController) GetListSong(f *fiber.Ctx) error {
	pageStr := f.Query("page")
	pageSizeStr := f.Query("page_size")

	page, pageSize, numErr := handler.CheckQuery(f, pageStr, pageSizeStr)
	if numErr != 0 {
		return res.ErrorResponse(f, numErr)
	}

	filter := f.Query("filter")
	singer := f.Query("singer")
	album := f.Query("album")
	genres := f.Query("genres")

	song, code, err := s.service.GetListSong(f.Context(), singer, album, genres, filter, page, pageSize)
	if err != nil {
		return res.ErrorResponse(f, code)
	}
	return res.SuccessResponse(f, code, song)
}

// song godoc
// @Summary      Delete song with id
// @Description  Delete song with id
// @Tags         Songs
// @Accept       json
// @Produce      json
// @Param        id path string true "ID"
// @Security BearerAuth
// @Success      200  "No content"
// @Failure      500  {object}  res.ErrSwaggerJson
// @Router       /songs/soft/{id} [post]
func (s *SongController) DeleteSong(f *fiber.Ctx) error {
	idStr := f.Params("id")

	uuid, err := uuid.Parse(idStr)

	if err != nil {
		return res.ErrorResponse(f, res.ErrBadRequestId)
	}

	err = s.service.DeleteSong(f.Context(), uuid)
	if err != nil {
		return res.ErrorNonKnow(f, err.Error())
	}

	return res.SuccessResponse(f, 204, nil)
}
