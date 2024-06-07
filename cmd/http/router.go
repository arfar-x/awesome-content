package http

func (s *Server) InitRoutes() {
	v1Group := s.engine.Group("api/v1")
	{
		s.DefineContentRoutes(v1Group)
	}
}
