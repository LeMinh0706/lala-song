package server

func (s *Server) Static() {
	s.Router.Static("upload/avatar", "./upload/avatar")
	s.Router.Static("upload/background", "./upload/background")
	s.Router.Static("upload/song", "./upload/song")
}
