package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"image/color"
)

type ThemeDefault struct{}

var _ fyne.Theme = (*ThemeDefault)(nil)

func (m ThemeDefault) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	return theme.DefaultTheme().Color(name, variant)
}

func (m ThemeDefault) Font(style fyne.TextStyle) fyne.Resource {
	if style.Bold {
		return resourceMsyhbdTtc
	}
	if style.Italic {
		return theme.DefaultTheme().Font(style)
	}
	if style.Monospace {
		return theme.DefaultTheme().Font(style)
	}
	return resourceMsyhTtc
}

func (m ThemeDefault) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (m ThemeDefault) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}