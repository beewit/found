package main

import (
	"testing"
	"github.com/beewit/beekit/utils"
)

func TestRedisList(t *testing.T) {
	err := utils.CloseChrome()
	if err != nil {
		t.Error(err.Error())
	}
}
