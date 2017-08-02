#!/bin/sh -x
go test -c -o test_cover main_test.go -coverprofile p.out -covermode atomic --coverpkg github.com/AZavorotnii/go-coverpkg-initissue/foo,github.com/AZavorotnii/go-coverpkg-initissue/bar
./test_cover
