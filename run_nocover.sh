#!/bin/sh -x
go test -c -o test_nocover main_test.go
./test_nocover
