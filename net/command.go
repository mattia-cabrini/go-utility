// Copyright (c) 2023 Mattia Cabrini
// SPDX-License-Identifier: MIT

package net

import (
	"fmt"
	"io"
	"strconv"
)

type Command struct {
	Params map[string]string
	Data   []byte
}

func (c *Command) readData(r io.Reader) (err error) {
	var nTot, n int
	c.Data = make([]byte, c.GetParamInt("Content-Length"))

	if cap(c.Data) == 0 {
		return
	}

	for nTot < cap(c.Data) && err == nil {
		n, err = r.Read(c.Data[nTot:])
		nTot += n
	}

	return
}

func (c *Command) GetParamInt(param string) (n int) {
	n, _ = strconv.Atoi(c.GetParam(param))
	return
}

func (c *Command) GetParam(param string) (str string) {
	str, b := c.Params[param]

	if !b {
		str = ""
	}

	return
}

func (c *Command) Print(w io.Writer) (err error) {
	for k, v := range c.Params {
		str := fmt.Sprintf("%s: %s\r\n", k, v)

		_, err = w.Write([]byte(str))
		if err != nil {
			break
		}
	}

	_, err = w.Write([]byte("\r\n"))
	if err == nil {
		_, err = w.Write(c.Data)
	}

	return
}

func splitCommandValue(str string) (comm string, value string) {
	var i int

	for _, commx := range str {
		if commx == ':' {
			break
		}

		comm += string(commx)
		i++
	}

	if len(str) >= i+2 {
		value = str[i+2:]
	}

	return
}

func (c *Command) InheritParameters(oth *Command, paramNames ...string) {
	for _, px := range paramNames {
		c.Params[px] = oth.Params[px]
	}
}
