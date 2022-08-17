// Package itertor
// Time    : 2022/8/17 10:25
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package itertor

import (
	"github.com/stretchr/testify/suite"
	"log"
	"testing"
)

type UserIteratorTestSuit struct {
	uc *UserCollection
	suite.Suite
}

func (s *UserIteratorTestSuit) SetupTest() {
	log.Println("SetupTest start to load data")
	user1 := &User{
		Name: "xushiyin",
		Age:  18,
	}
	user2 := &User{
		Name: "xushiyin2",
		Age:  18,
	}
	user3 := &User{
		Name: "xushiyin3",
		Age:  18,
	}
	s.uc = &UserCollection{
		Users: []*User{
			user1, user2, user3,
		},
	}
}

func (s *UserIteratorTestSuit) TearDownTest() {
	log.Println("TearDownTest over")
}

// TestUserIterator is the testing entrance of UserIteratorTestSuit
func TestUserIterator(t *testing.T) {
	suite.Run(t, new(UserIteratorTestSuit))
}

func (s *UserIteratorTestSuit) TestUserIteratorFunc() {
	ui := s.uc.CreateIterator()
	for ui.HasNext() {
		user := ui.GetNext()
		log.Printf("user: %#v\n", user)
	}
}
