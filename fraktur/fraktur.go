// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package fraktur

import (
	"image/color"

	"charm.land/fang/v2"
	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/exp/charmtone"
)

// Theme holds the Fraktur color palette as raw Lip Gloss colors, usable in
// any Lip Gloss or Bubble Tea project.
type Theme struct {
	Base           color.Color
	Title          color.Color
	Codeblock      color.Color
	Program        color.Color
	Command        color.Color
	DimmedArgument color.Color
	Comment        color.Color
	Flag           color.Color
	Argument       color.Color
	Description    color.Color
	FlagDefault    color.Color
	QuotedString   color.Color
	ErrorHeader    [2]color.Color // 0=fg, 1=bg
}

// NewTheme returns the Fraktur color palette resolved for the given light/dark
// context. Use [charm.land/lipgloss/v2.LightDark] to create c.
func NewTheme(c lipgloss.LightDarkFunc) Theme {
	// TODO: this needs work
	return Theme{
		Base:           c(charmtone.Charcoal, charmtone.Ash),
		Title:          charmtone.Charple,
		Codeblock:      c(charmtone.Salt, charmtone.Salt),
		Program:        c(charmtone.Malibu, charmtone.Guppy),
		Command:        c(charmtone.Pony, charmtone.Cheeky),
		DimmedArgument: c(charmtone.Squid, charmtone.Oyster),
		Comment:        charmtone.Squid,
		Flag:           charmtone.Guac,
		Argument:       c(charmtone.Charcoal, charmtone.Ash),
		Description:    c(charmtone.Charcoal, charmtone.Ash),
		FlagDefault:    c(charmtone.Smoke, charmtone.Squid),
		QuotedString:   c(charmtone.Coral, charmtone.Salmon),
		ErrorHeader: [2]color.Color{
			charmtone.Salt,
			charmtone.Cherry,
		},
	}
}

// CharmFangColorSchemeFunc returns the Fraktur theme as a [charm.land/fang/v2.ColorScheme].
func CharmFangColorSchemeFunc(c lipgloss.LightDarkFunc) fang.ColorScheme {
	t := NewTheme(c)
	return fang.ColorScheme{
		Base:           t.Base,
		Title:          t.Title,
		Codeblock:      t.Codeblock,
		Program:        t.Program,
		Command:        t.Command,
		DimmedArgument: t.DimmedArgument,
		Comment:        t.Comment,
		Flag:           t.Flag,
		Argument:       t.Argument,
		Description:    t.Description,
		FlagDefault:    t.FlagDefault,
		QuotedString:   t.QuotedString,
		ErrorHeader:    t.ErrorHeader,
	}
}
