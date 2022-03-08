package gotnet

type AppBuilder struct {
	Service *Service
}

func NewAppBuilder() *AppBuilder {
	return &AppBuilder{
		Service: newService(),
	}
}

func (builder *AppBuilder) Build() *App {
	app := &App{
		Service: builder.Service,
		Router:  NewRouter(),
	}
	return app
}
