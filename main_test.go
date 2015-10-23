package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadLines(t *testing.T) {

	assert := assert.New(t)
	lines, err := readLines("./test.chordpro")

	assert.Nil(err)
	assert.Equal("{t: Hello world}", lines[0], "they should be equal")
	assert.Equal("", lines[4], "they should be equal")
	assert.Equal("[G]Lorem ipsum dolor sit amet, [D/F#]consectetur adipiscing elit. Don[Em]ec a diam lectus.", lines[5], "they should be equal")
}
