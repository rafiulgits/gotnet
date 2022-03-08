package gotnet

type Server struct {
	router *Router
}

func newServer() *Server {
	return &Server{router: newRouter()}
}

func (server *Server) Router() *Router {
	return server.router
}
