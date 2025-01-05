package server

func (s *Server) Static() {
	s.Router.Static("upload/avatar", "./upload/avatar")
	s.Router.Static("upload/song", "./upload/song")
	s.Router.Static("upload/album", "./upload/album")
	s.Router.Static("upload/singers", "./upload/singers")
	s.Router.Static("upload/lyrics", "./upload/lyrics")
	s.Router.Static("upload/genres", "./upload/genres")
}
