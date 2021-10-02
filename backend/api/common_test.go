package api

type fixture struct {
	api *Server
}

func setTestFixture() *fixture {
	srv := New("", "", "")
	srv.Routes()
	return &fixture{
		api: srv,
	}
}
