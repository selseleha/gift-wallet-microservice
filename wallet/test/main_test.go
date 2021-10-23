package test

import (
	"os"
	"task/pkg/test"
	"testing"
)

var fakeRepo = test.NewFakeRepo()

func TestMain(m *testing.M) {

	os.Exit(m.Run())

}
