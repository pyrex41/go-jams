package jams

import (
	"errors"
	"log"
)

type ParserState struct {
	bytes []byte
	i     int
}

func (ps ParserState) hasmore() bool {
	return ps.i < len(ps.bytes)
}

func (ps *ParserState) incr() {
	ps.i += 1
}

func (ps ParserState) current() byte {
	if ps.hasmore() {
		return ps.bytes[ps.i]
	} else {
		return 0
	}
}

func (ps *ParserState) advance() byte {
	current := ps.current()
	ps.incr()
	return current
}

func isjamspace(b byte) bool {
	return b == byte(' ') || b == '\n' || b == '\t'
}

func (ps *ParserState) chompspace() {
	for {
		if !(ps.hasmore() && isjamspace(ps.current())) {
			break
		}
		ps.incr()
	}
}

func isws(b byte) bool {
	return b == byte(' ') || b == byte('\t') || b == byte('\n') || b == byte('\r')
}

func issyn(b byte) bool {
	switch b {
	case byte('{'), byte('}'), byte('['), byte(']'):
		return true
	default:
		return false
	}
}

func isany(b byte) bool {
	return issafe(b) || isws(b) || issyn(b) || b == 92
}

func issafe(b byte) bool {
	return b == 33 || (35 <= b && b <= 90) || (94 <= b && b <= 122) || b == 124 || b == 126
}

func (ps *ParserState) parse_bare(close byte) string {
	ps.chompspace()
	out := make([]byte, 0)
	m := len(ps.bytes) + 1
	for i := ps.i; i < m; i++ {
		c := ps.current()
		if issafe(c) {
			ps.incr()
			out = append(out, c)
		} else if isws(c) {
			ps.incr()
			return string(out)
		} else if c == close {
			return string(out)
		} else {
			err := errors.New("Error should not surface")
			log.Fatal(err)
			return "null"
		}
	}
	err := errors.New("Something went wrong")
	log.Fatal(err)
	return "null"
}

func (ps *ParserState) parse_quote(close byte) string {
	ps.chompspace()
	out := make([]byte, 0)
	m := len(ps.bytes) + 1
	c := ps.advance()
	if !(c == byte('"')) {
		err := errors.New("Quote must start with quote sign")
		log.Fatal(err)
		return "null"
	}
	ps.chompspace()
	for i := 1; i < m; i++ {
		c = ps.current()
		if c == byte('\\') {
			ps.incr()
			out = append(out, c)
			c = ps.advance()
			out = append(out, c)
		} else if c == byte('"') {
			ps.incr()
			return string(out)
		} else if c == close {
			return string(out)

		} else {
			ps.incr()
			out = append(out, c)
		}
	}
	err := errors.New("This timed out; should never happen")
	log.Fatal(err)
	return string(out)
}

func (ps *ParserState) parse_str(close byte) string {
	ps.chompspace()
	if ps.current() == byte('"') {
		return ps.parse_quote(close)
	} else {
		return ps.parse_bare(close)
	}
}

func (ps *ParserState) parse_arr() []interface{} {
	a := make([]interface{}, 0)
	c := ps.advance()
	ps.chompspace()
	m := len(ps.bytes) + 1
	if !(c == byte('[')) {
		err := errors.New("cannot parse arr that doesn't begin with [")
		log.Fatal(err)
		return a
	}
	ps.chompspace()
	if ps.current() == byte(']') {
		ps.incr()
		return a
	}
	for i := 0; i < m; i++ {
		ps.chompspace()
		var jam interface{}
		jam = ps.parse_jam(']')
		a = append(a, jam)
		ps.chompspace()
		if ps.current() == byte(']') {
			break
		}

	}
	ps.incr()
	return a
}

func (ps *ParserState) parse_obj() map[string]interface{} {
	ps.chompspace()
	c := ps.advance()
	object := make(map[string]interface{})
	m := len(ps.bytes) + 1
	if !(c == byte('{')) {
		err := errors.New("cannot parse object that doesn't begin with {")
		log.Fatal(err)
		return object
	}
	ps.chompspace()
	if ps.current() == byte('}') {
		ps.incr()
		return object
	}
	for i := 0; i < m; i++ {
		key := ps.parse_str(' ')
		ps.chompspace()
		var val interface{}
		val = ps.parse_jam('}')
		ps.chompspace()
		object[key] = val
		if ps.current() == '}' {
			break
		}
	}
	ps.incr()
	return object
}

func (ps *ParserState) parse_jam(close byte) interface{} {
	var jam interface{}
	switch ps.current() {
	case byte('{'):
		jam = ps.parse_obj()
	case byte('['):
		jam = ps.parse_arr()
	default:
		jam = ps.parse_str(close)
	}
	return jam
}

func Parse(content []byte) interface{} {
	ps := ParserState{content, 0}
	ps.chompspace()
	var out interface{}
	out = ps.parse_jam(' ')
	return out
}
