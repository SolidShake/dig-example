package main

import (
	"fmt"

	"go.uber.org/dig"
)

type testService struct {
	config *config
}

type config struct {
	field string
}

func initConfig() *config {
	return &config{field: "value"}
}

func initTestService(config *config) *testService {
	return &testService{config: config}
}

func (s *testService) SomeAction() {
	fmt.Println("some action...")
}

func BuildContainer() (*dig.Container, error) {
	container := dig.New()

	err := container.Provide(func() *config {
		return initConfig()
	})
	if err != nil {
		return nil, err
	}

	err = container.Provide(func(cfg *config) *testService {
		return initTestService(cfg)
	})
	if err != nil {
		return nil, err
	}

	fmt.Println("container build finished...")

	return container, nil
}

func main() {
	fmt.Println("start...")

	container, err := BuildContainer()
	if err != nil {
		panic(err)
	}

	err = container.Invoke(func(service *testService) {
		service.SomeAction()
	})

	if err != nil {
		panic(err)
	}

	fmt.Println("end...")
}
