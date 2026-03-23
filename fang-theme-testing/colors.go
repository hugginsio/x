// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package main

import (
	"fmt"
	"image/color"
	"math"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/exp/charmtone"
	"github.com/spf13/cobra"
)

const (
	paletteColumns = 8
	swatchWidth    = 10 // chars per swatch cell (name + trailing spaces)
)

func newColorsCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "colors",
		Short: "Preview all CharmTone palette colors with Lip Gloss",
		Long:  "Taskwarrior-inspired color chart for the CharmTone palette, rendered with Lip Gloss.",
		Run:   runColors,
	}
}

func runColors(cmd *cobra.Command, _ []string) {
	w := cmd.OutOrStdout()

	heading := func(title, sub string) {
		h := lipgloss.NewStyle().Bold(true).Render(title)
		if sub != "" {
			h += "  " + lipgloss.NewStyle().Faint(true).Render(sub)
		}
		fmt.Fprintf(w, "\n%s\n", h)
	}

	// printRows renders two rows for a batch of colors: first as foreground,
	// then as background, mirroring the Taskwarrior `task colors` layout.
	printRows := func(keys []charmtone.Key) {
		fmt.Fprint(w, "  ")
		for _, k := range keys {
			fmt.Fprint(w, lipgloss.NewStyle().Foreground(k).Render(fmt.Sprintf("%-*s", swatchWidth, k.String())))
		}
		fmt.Fprintln(w)

		fmt.Fprint(w, "  ")
		for _, k := range keys {
			fg := contrastFor(k)
			fmt.Fprint(w, lipgloss.NewStyle().Background(k).Foreground(fg).Render(fmt.Sprintf("%-*s", swatchWidth, k.String())))
		}
		fmt.Fprintln(w)
	}

	keys := charmtone.Keys()
	heading("CharmTone Palette", fmt.Sprintf("(%d colors)", len(keys)))
	for i := 0; i < len(keys); i += paletteColumns {
		end := min(i+paletteColumns, len(keys))
		printRows(keys[i:end])
		fmt.Fprintln(w)
	}

	heading("Diff Colors", "(additions / deletions)")
	printRows([]charmtone.Key{charmtone.Julep, charmtone.Pickle, charmtone.Gator, charmtone.Spinach})
	fmt.Fprintln(w)
	printRows([]charmtone.Key{charmtone.Cherry, charmtone.Pom, charmtone.Steak, charmtone.Toast})
	fmt.Fprintln(w)

	heading("Effects", "")
	type namedEffect struct {
		label string
		style lipgloss.Style
	}

	effects := []namedEffect{
		{"Charple", lipgloss.NewStyle().Foreground(charmtone.Charple)},
		{"bold Charple", lipgloss.NewStyle().Foreground(charmtone.Charple).Bold(true)},
		{"italic Coral", lipgloss.NewStyle().Foreground(charmtone.Coral).Italic(true)},
		{"underline Malibu", lipgloss.NewStyle().Foreground(charmtone.Malibu).Underline(true)},
		{"strike Pom", lipgloss.NewStyle().Foreground(charmtone.Pom).Strikethrough(true)},
		{"on Guac", lipgloss.NewStyle().Background(charmtone.Guac).Foreground(charmtone.Pepper)},
		{"on Butter", lipgloss.NewStyle().Background(charmtone.Butter).Foreground(charmtone.Pepper)},
		{"inverse Charple", lipgloss.NewStyle().Foreground(charmtone.Charple).Reverse(true)},
		{"faint Squid", lipgloss.NewStyle().Foreground(charmtone.Squid).Faint(true)},
	}

	rendered := make([]string, len(effects))
	for i, e := range effects {
		rendered[i] = e.style.Render(e.label)
	}

	fmt.Fprintf(w, "  %s\n", strings.Join(rendered, "  "))
}

func themeSwatch(c color.Color, name string, asBg bool) string {
	if asBg {
		return lipgloss.NewStyle().Background(c).Foreground(charmtone.Butter).Render(" " + name + " ")
	}
	return lipgloss.NewStyle().Foreground(c).Render(name)
}

// contrastFor returns Pepper (dark) or Butter (light) as a foreground color
// that provides readable contrast against the given background.
func contrastFor(bg color.Color) color.Color {
	r, g, b, _ := bg.RGBA()
	lum := 0.2126*linearize(float64(r)/65535) +
		0.7152*linearize(float64(g)/65535) +
		0.0722*linearize(float64(b)/65535)
	if lum > 0.35 {
		return charmtone.Pepper
	}
	return charmtone.Butter
}

func linearize(v float64) float64 {
	if v <= 0.04045 {
		return v / 12.92
	}
	return math.Pow((v+0.055)/1.055, 2.4)
}
