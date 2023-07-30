package noodle_test

import (
    "testing"
    "github.com/ironsmith58/gotest/noodle"
)

func TestBoil(t *testing.T) {

    want := "Boil fish"
    expected := noodle.Boil("fish")

    if expected != want {
        t.Fatalf("wanted '%s', received '%s'", want, expected)
    }

}
