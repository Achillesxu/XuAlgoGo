// Package itertor
// Time    : 2022/8/17 09:44
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package itertor

type Collection interface {
	createIterator() Iterator
}

type Iterator interface {
	HasNext() bool
	GetNext() *User
}

type User struct {
	Name string
	Age  int
}

type UserCollection struct {
	Users []*User
}

type UserIterator struct {
	Index int
	Users []*User
}

func (u *UserIterator) HasNext() bool {
	if u.Index < len(u.Users) {
		return true
	}
	return false
}

func (u *UserIterator) GetNext() *User {
	if u.HasNext() {
		user := u.Users[u.Index]
		u.Index++
		return user
	}
	return nil
}

func (uc *UserCollection) CreateIterator() Iterator {
	return &UserIterator{
		Index: 0,
		Users: uc.Users,
	}
}
