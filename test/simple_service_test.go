package test

import (
	"fmt"
	"testing"
	"xriot/learn-golang-dependency-injection/simple"
)

func TestSimpleService(t *testing.T) {
	simpleService, err := simple.InitializeService()
	fmt.Println(err)
	fmt.Println(simpleService)
}
