// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package example_zoo

// Injectors from wire.gen.go:

func InitializeZoo() (Zoo, func(), error) {
	cat := Cat{}
	dog := ProvideDog()
	lion := NewLion()
	bird := &Bird{}
	zoo := Zoo{
		Cat:       cat,
		Dog:       dog,
		Lion:      lion,
		FlyAnimal: bird,
	}
	return zoo, func() {
	}, nil
}