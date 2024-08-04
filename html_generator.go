// Copyright (c) 2024 Mattia Cabrini
// SPDX-License-Identifier: MIT

package utility

import (
	"fmt"
	"os"
)

type HtmlGenerator struct {
	fp *os.File

	partial      int
	writtenBytes int
	openedTags   []string

	defined bool

	style map[string]string

	err error
}

func NewHtmlGenerator(fp *os.File) (h *HtmlGenerator) {
	if fp == nil {
		fp = os.Stdout
	}

	return &HtmlGenerator{
		fp:    fp,
		style: make(map[string]string),
	}
}

func (h *HtmlGenerator) push(tag string) {
	h.openedTags = append(h.openedTags, tag)
}

func (h *HtmlGenerator) pop() (tag string) {
	if l := len(h.openedTags); l > 0 {
		tag = h.openedTags[l-1]
		h.openedTags = h.openedTags[:l-1]
		h.defined = true
	}

	return
}

func (h *HtmlGenerator) printf(format string, args ...any) (w int, _ error) {
	w, h.err = fmt.Fprintf(h.fp, format, args...)
	h.writtenBytes += w
	h.partial += w
	return w, h.err
}

func (h *HtmlGenerator) Err() error {
	return h.err
}

func (h *HtmlGenerator) Open(tag string) (w int, _ error) {
	h.closeDefinition()

	w, h.err = h.printf("<%s ", tag)

	if h.err == nil {
		h.push(tag)
		h.defined = false
	}

	return w, h.err
}

func (h *HtmlGenerator) AddAttribute(name string, value string) {
	_, h.err = h.printf(" %s=\"%s\" ", name, value)
}

func (h *HtmlGenerator) AddStyle(name string, value string) {
	h.style[name] = value
}

func (h *HtmlGenerator) AddText(text string) {
	h.closeDefinition()

	if h.err != nil {
		h.printf("%s", text)
	}
}

func (h *HtmlGenerator) closeDefinition() {
	if len(h.openedTags) == 0 {
		return
	}

	if h.defined {
		return
	}

	h.printStyle()

	if h.err == nil {
		h.printf(">")
	}

	h.style = make(map[string]string)
	h.defined = true
}

func (h *HtmlGenerator) printStyle() {
	if !h.defined && len(h.style) > 0 {
		h.printf("style=\"")

		for k, v := range h.style {
			h.printf("%s: %s; ", k, v)
		}

		h.printf("\"")
	}
}

func (h *HtmlGenerator) OpenInput(class string, typ string, name string, placeholder string, value string) {
	h.Open("input")

	h.AddAttribute("class", Quote(class))
	h.AddAttribute("type", Quote(typ))
	h.AddAttribute("name", Quote(name))
	h.AddAttribute("placeholder", Quote(placeholder))
	h.AddAttribute("value", Quote(value))
}

func (h *HtmlGenerator) Close() {
	if len(h.openedTags) == 0 {
		return
	}

	h.closeDefinition()

	h.printf("</%s>", h.pop())
}

func (h *HtmlGenerator) CloseAll() {
	for len(h.openedTags) != 0 {
		h.Close()
	}
}

func (h *HtmlGenerator) GetPartialBytesCount() int {
	return h.partial
}

func (h *HtmlGenerator) GetWrittenBytesCount() int {
	return h.writtenBytes
}

func (h *HtmlGenerator) ResetPartialBytesCount() {
	h.partial = 0
}
