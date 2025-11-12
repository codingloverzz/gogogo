package main

type GetName interface {
	Name() string
}

type Inner struct{}

func (i Inner) Name() string {
	return "我是Inner"
}

func (i Inner) Hello() string {

	return "hello" + i.Name()
}

type Outer struct {
	Inner
}

func (i Outer) Name() string {
	return "我是outer"
}

func Compose() {
	o := Outer{
		Inner: Inner{},
	}
	println(o.Hello())

}
