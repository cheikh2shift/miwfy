package quick_test

import (
	"math/rand"
	"reflect"
	"testing"
	"testing/quick"

	quack "github.com/cheikh2shift/miwfy/quick"
)

func TestCalculateAverageTime(t *testing.T) {

	c := &quack.ChatRoom{
		Usernames:         []string{"Foo", "Bar"},
		TotalActionOnline: 10,
	}

	c.CalculateAverageAction()

	if c.AverageActionPerUser != 5 {
		t.Errorf(
			"Incorrect value received, expected %v got %v",
			5,
			c.AverageActionPerUser,
		)
	}

	// Add fuzzing
}

func TestGetAverageAction(t *testing.T) {

	for i := int64(0); i < 100; i++ {

		//construct RAND
		// with index
		rand := rand.New(
			rand.NewSource(i),
		)

		cType, ok := quick.Value(
			reflect.TypeOf(quack.ChatRoom{}),
			rand,
		)

		if !ok {
			t.Error("Error generating a value")
		}

		c := cType.Interface().(quack.ChatRoom)

		t.Log(
			quack.GetAverageAction(&c),
		)
	}

}
