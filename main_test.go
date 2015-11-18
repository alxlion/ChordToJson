/* The MIT License (MIT)

Copyright (c) 2015 Alexandre Lion <contact@alexandrelion.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE. */

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

func TestParseLine(t *testing.T) {

	assert := assert.New(t)
	lines, err := readLines("./test.chordpro")

	assert.Nil(err)

	// Testing title type
	titleType, titleContent, err := parseLine(lines[0])
	assert.Nil(err)
	assert.Equal(map[string]string{"t": "title"}, titleType, "they should have the same type")
	assert.Equal(" Hello world", titleContent, "they should have the same content")

	// Testing subtitle type
	stType, stContent, err := parseLine(lines[1])
	assert.Nil(err)
	assert.Equal(map[string]string{"st": "subtitle"}, stType, "they should have the same type")
	assert.Equal(" Foo Bar", stContent, "they should have the same content")

	// Testing comment type
	commentType, commentContent, err := parseLine(lines[3])
	assert.Nil(err)
	assert.Equal(map[string]string{"c": "comment"}, commentType, "they should have the same type")
	assert.Equal(" © 2015 FooBar Ltd", commentContent, "they should have the same content")

	// Testing verse type
	verseType, verseContent, err := parseLine(lines[5])
	assert.Nil(err)
	assert.Equal(map[string]string{"v": "verse"}, verseType, "they should have the same type")
	assert.Equal("[G]Lorem ipsum dolor sit amet, [D/F#]consectetur adipiscing elit. Don[Em]ec a diam lectus.", verseContent, "they should have the same type")
}
