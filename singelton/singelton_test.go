package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSingle(t *testing.T) {

	var singelton *single
	singelton = getInstance()
	var sec_singelton *single
	sec_singelton = getInstance()
	assert.Equal(t, singelton, sec_singelton, "they should not be equal")
}
