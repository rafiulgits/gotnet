package gotnet

type Service struct {
	container *Container
}

func newService() *Service {
	return &Service{
		container: newContainer(),
	}
}

func (service *Service) AddSingleton(constructor interface{}) {
	service.container.Provide(constructor)
}

func (service *Service) Container() *Container {
	return service.container
}
