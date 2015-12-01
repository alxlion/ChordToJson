ChordToJson
=====
[![Build Status](https://travis-ci.org/lnalex/ChordToJson.svg)](https://travis-ci.org/lnalex/ChordToJson)

ChordPro to JSON converter

## Usage

    ./ChordToJson file [options]

| Option name | Description        |
|-------------|--------------------|
| --pretty    | Output pretty JSON |

## Build

    $ export GOPATH=$PWD
    $ go get github.com/lnalex/ChordToJson


## Demo
_test.chorpro_

	{t: Hello world}
	{st: Foo Bar}

	{c: © 2015 FooBar Ltd}

	[G]Lorem ipsum dolor sit amet, [D/F#]consectetur adipiscing elit. Don[Em]ec a diam lectus.
	Sed s[G]it amet ipsum mauris.
	Maecenas con[G]gue ligula ac[D/F#] quam viverra nec consectetur [G]ante hendrerit.
	Donec e[Bm7]t mollis dolor.[D/F#]

_Output (Pretty option)_

	{
	 "lines": [
	  {
	   "type": {
	    "t": "title"
	   },
	   "content": " Hello world"
	  },
	  {
	   "type": {
	    "st": "subtitle"
	   },
	   "content": " Foo Bar"
	  },
	  {
	   "type": {
	    "b": "blank"
	   }
	  },
	  {
	   "type": {
	    "c": "comment"
	   },
	   "content": " © 2015 FooBar Ltd"
	  },
	  {
	   "type": {
	    "b": "blank"
	   }
	  },
	  {
	   "type": {
	    "v": "verse"
	   },
	   "content": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec a diam lectus.",
	   "chords": {
	    "0": "[G]",
	    "31": "[D/F#]",
	    "69": "[Em]"
	   }
	  },
	  {
	   "type": {
	    "v": "verse"
	   },
	   "content": "Sed sit amet ipsum mauris.",
	   "chords": {
	    "5": "[G]"
	   }
	  },
	  {
	   "type": {
	    "v": "verse"
	   },
	   "content": "Maecenas congue ligula ac quam viverra nec consectetur ante hendrerit.",
	   "chords": {
	    "12": "[G]",
	    "28": "[D/F#]",
	    "64": "[G]"
	   }
	  },
	  {
	   "type": {
	    "v": "verse"
	   },
	   "content": "Donec et mollis dolor.",
	   "chords": {
	    "27": "[D/F#]",
	    "7": "[Bm7]"
	   }
	  }
	 ]
	}

## License
[MIT License](https://github.com/lnalex/ChordToJson/blob/master/LICENSE)
