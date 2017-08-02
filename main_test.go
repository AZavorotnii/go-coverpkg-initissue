package main

import (
        "fmt"

        "github.com/AZavorotnii/go-coverpkg-initissue/foo"
        "testing"
        "os"
)

func init() {
        fmt.Println("main.")
}

func TestMain(m *testing.M) {
        fmt.Println("main main.")
	foo.Foo()

        os.Exit(m.Run())
}
