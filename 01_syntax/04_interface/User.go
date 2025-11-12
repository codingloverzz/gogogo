package main

type Runnable interface {
	Run() string
	Eat() string
}

type User struct {
	name string
	age  int
}

func (u *User) Run() string {
	//TODO implement me
	panic("implement me")
}

func (u *User) Eat() string {
	//TODO implement me
	panic("implement me")
}

func (u *User) ChangeAge(newAge int) {
	u.age = newAge

}

func (u User) ChangeName(newName string) {
	u.name = newName
}
