package services

type serviceContainer struct {
	services map[string]interface{}
}

type initializerFunc func() (interface{}, error)

func (sc *serviceContainer) get(name string, initializer initializerFunc) (any, error) {
	var err error
	if sc.services[name] == nil {
		sc.services[name], err = initializer()
	}
	return sc.services[name], err
}

func (sc *serviceContainer) getOrPanic(name string, initializer initializerFunc) any {
	service, err := sc.get(name, initializer)
	if err != nil {
		panic(err)
	}
	return service
}

func newServiceContainer() *serviceContainer {
	return &serviceContainer{
		services: make(map[string]interface{}),
	}
}

var Container = newServiceContainer()
