// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package fraktur

import (
	"image/color"

	"charm.land/fang/v2"
	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/exp/charmtone"
)

func CharmFangColorSchemeFunc(c lipgloss.LightDarkFunc) fang.ColorScheme {
	return fang.ColorScheme{
		Base:           c(charmtone.Charcoal, charmtone.Ash),
		Title:          charmtone.Charple,
		Codeblock:      c(charmtone.Salt, lipgloss.Color("#2F2E36")),
		Program:        c(charmtone.Malibu, charmtone.Guppy),
		Command:        c(charmtone.Pony, charmtone.Cheeky),
		DimmedArgument: c(charmtone.Squid, charmtone.Oyster),
		Comment:        c(charmtone.Squid, lipgloss.Color("#747282")),
		Flag:           c(lipgloss.Color("#0CB37F"), charmtone.Guac),
		Argument:       c(charmtone.Charcoal, charmtone.Ash),
		Description:    c(charmtone.Charcoal, charmtone.Ash), // flag and command descriptions
		FlagDefault:    c(charmtone.Smoke, charmtone.Squid),  // flag default values in descriptions
		QuotedString:   c(charmtone.Coral, charmtone.Salmon),
		ErrorHeader: [2]color.Color{
			charmtone.Butter,
			charmtone.Cherry,
		},
	}
}
