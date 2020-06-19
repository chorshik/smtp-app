package server

// setHandlers ...
func (s *Server ) setHandlers()  {
	//h := handlers.NewHandlers()

	s.router.HandleFunc("/", s.handler.Index.HandleIndex()).Methods("GET")
	s.router.HandleFunc("/send", s.handler.Sender.Send()).Methods("POST")
}
