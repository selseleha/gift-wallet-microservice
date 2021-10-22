package test

import (
	"fmt"
	"math/rand"
	"os"
	"task/pkg/test"
	"testing"
)

var fakeRepo = test.NewFakeRepo()

func TestMain(m *testing.M) {

	os.Exit(m.Run())

}

func randomPhoneNumber() string {
	return fmt.Sprintf("09%d", random(100000000, 900000000))
}
func random(min int, max int) int {
	return rand.Intn(max-min) + min
}
