package song

import (
	"fmt"
	"path/filepath"
	"strconv"
	"time"

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
	const maxSize = 10 << 20
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
