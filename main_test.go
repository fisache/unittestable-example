package main

import (
	"testing"

	"github.com/fisache/unittestable-example/hello"
	"github.com/fisache/unittestable-example/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHelloFunc(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockHello := mock.NewMockInterface(mockCtrl)

	hello.New(mockHello)

	mockHello.EXPECT().SayHello("inki").Return("gomock").Times(1)

	assert.Equal(t, "gomock", helloFunc())
}
