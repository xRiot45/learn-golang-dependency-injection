package test

import (
	"fmt"
	"testing"
	"xriot/learn-golang-dependency-injection/simple"

	"github.com/stretchr/testify/assert"
)

func TestSimpleServiceError(t *testing.T) {
	simpleService, err := simple.InitializeService(true)
	assert.Nil(t, simpleService)
	assert.NotNil(t, err)
	fmt.Println("Error:", err.Error())
}

func TestSimpleServiceSuccess(t *testing.T) {
	simpleService, err := simple.InitializeService(false)
	assert.NotNil(t, simpleService)
	assert.Nil(t, err)
	fmt.Println("Service:", simpleService)
}
