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

func TestGetLineType(t *testing.T) {

	assert := assert.New(t)
	lines, err := readLines("./test.chordpro")

	assert.Nil(err)

	// Testing title type
	titleType, _, _ := getLineType(lines[0])
	assert.Equal("title", titleType, "they should have the same type")

	// Testing subtitle type
	stType, _, _ := getLineType(lines[1])
	assert.Equal("subtitle", stType, "they should have the same type")

	// Testing comment type
	commentType, _, _ := getLineType(lines[3])
	assert.Equal("comment", commentType, "they should have the same type")

	// Testing verse type
	verseType, _, _ := getLineType(lines[5])
	assert.Equal("verse", verseType, "they should have the same type")
}
