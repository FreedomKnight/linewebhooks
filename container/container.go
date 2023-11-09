package container

import "go.uber.org/dig"

var container *dig.Container = dig.New()

func Get() *dig.Container {
    return container
}

func Invoke(function interface{}, opts ...dig.InvokeOption) error {
	return Get().Invoke(function, opts...)
}

func Provide(constructor interface{}, opts ...dig.ProvideOption) error {
	return Get().Provide(constructor, opts...)
}
