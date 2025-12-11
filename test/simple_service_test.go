package test

import (
	"fmt"
	"testing"
	"xriot/learn-golang-dependency-injection/simple"
)

func TestSimpleService(t *testing.T) {
	simpleService := simple.InitializeService()
	fmt.Println(simpleService.SimpleRepository)
}
