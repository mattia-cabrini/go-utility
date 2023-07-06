/*
MIT License

Copyright (c) 2023 Mattia Cabrini

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
SOFTWARE.
*/

package net

import (
	"io"
)

type Scanner struct {
	r io.Reader

	buffer  []byte
	pos     int
	len     int
	command *Command
}

func NewScanner(r io.Reader, netBufSize int) *Scanner {
	return &Scanner{
		r:      r,
		pos:    0,
		len:    0,
		buffer: make([]byte, netBufSize),
	}
}

func (s *Scanner) Scan() (err error) {
	var line string
	var lineComplete bool

	s.command = &Command{
		Params: make(map[string]string),
	}

	for err == nil && line != "\r\n" {
		if err = s.read(); err != nil {
			break
		}

		line, lineComplete = s.lineFromBuf(s.buffer, line)

		for lineComplete {
			if line == "\r\n" {
				break
			}

			line = line[:len(line)-2] // removing trailing \r\n

			command, value := splitCommandValue(line)
			s.command.Params[command] = value
			line = ""

			line, lineComplete = s.lineFromBuf(s.buffer, line)
		}
	}

	if err == nil {
		err = s.command.readData(s.r)
		s.pos = s.len
	}

	return
}

func (s *Scanner) lineFromBuf(buf []byte, part string) (line string, complete bool) {
	var readCR bool
	var rx byte
	var i int

	line = part

	if s.pos == s.len {
		return
	}

BUFER_LOOP:
	for i, rx = range s.buffer[s.pos:s.len] {
		line += string(rx)

		switch rx {
		case '\r':
			readCR = true
		case '\n':
			if readCR {
				complete = true
				break BUFER_LOOP
			}
		default:
			readCR = false
		}
	}

	s.pos += i + 1

	return
}

func (s *Scanner) read() (err error) {
	if s.pos == s.len {
		s.len, err = s.r.Read(s.buffer)
		s.pos = 0
	}

	return
}

func (s *Scanner) Command() *Command {
	return s.command
}
