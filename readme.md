# Go-AutoWire
> helps you to generate wire files with easy annotate

[中文文档](./readme_zh.md)


this project is base on [wire](https://github.com/google/wire)

but it did `simplify` the wire usage and make wire `much more stronger `

## Installation

Install Wire by running:
```sh
go get github.com/google/wire/cmd/wire
```
then 

Install Gutowire by running:
```sh
go get github.com/windrivder/go-autowire/cmd/gutowire
```
and ensuring that $GOPATH/bin is added to your $PATH.

## Usage example

If you want to build a `zoo`,you may need some dependencies like animals
```go
package example

type Zoo struct{ 
    Cat         Cat
    Dog         Dog
    FlyAnimal FlyAnimal
}

type Cat struct{
}

type FlyAnimal interface{
    Fly()
}

type Bird struct{
}

func (b Bird)Fly(){
}

type Dog struct{
}
```

in traditional `wire`,you need to write some files to explain the wire relation to google/wire

```go
package example_zoo

import (
	"github.com/google/wire"
)

var zooSet = wire.NewSet(
	wire.Struct(new(Zoo), "*"),
)

var animalsSet = wire.NewSet(
	wire.Struct(new(Cat), "*"),
	wire.Struct(new(Dog), "*"),

	wire.Struct(new(Bird), "*"),
	wire.Bind(new(FlyAnimal), new(Bird)),
)

var sets = wire.NewSet(zooSet, animalsSet)

func InitZoo() Zoo {
	panic(wire.Build(sets))
}
```

you need to rewrite your `wire.go` and comes much more harder to manager all the dependencies

as your zoo goes bigger and bigger 

### but now

you can use `gutowire`

write annotate as `below`
```go
package example

// @autowire.init(set=zoo)
// it will be collect into zooSet (this comment is not necessary)
type Zoo struct{ 
    Cat         Cat
    Dog         Dog
    FlyAnimal FlyAnimal
}

// it will be collect into animalsSet (this comment is not necessary)
// @autowire(set=animals)
type Cat struct{
}


type FlyAnimal interface{
    Fly()
}

// it will be collect into animalsSet and wire as interface FlyAnimal (this comment is not necessary)
// @autowire(set=animals,FlyAnimal)
type Bird struct{
}

func (b Bird)Fly(){
}

// it will be collect into animalsSet (this comment is not necessary)
// @autowire(set=animals)
type Dog struct{
}
```


`.init` in `@autowire.init(set=zoo)` will auto write InitializeZoo func in wire.gen.go like below:
```go
// Code generated by go-autowire. DO NOT EDIT.

// +build wireinject
//
package example_zoo

import "github.com/google/wire"

func InitializeZoo() (*Zoo, func(), error) {
	panic(wire.Build(Sets))
}
```

and run
```sh
gutowire -s ./example_zoo ./example_zoo
```

`-s` means scope to look up build dependencies

all the wire files you need will generate

look at file generated in [example_zoo](./example_zoo)
