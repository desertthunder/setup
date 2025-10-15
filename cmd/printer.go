package main

import (
	"github.com/charmbracelet/lipgloss"
)

// Style definitions for terminal output using lipgloss.
// These provide colored and formatted text output for the CLI.
var (
	bolded      = bold()
	dimmed      = faint()
	boldCyan    = bold().Foreground(lipgloss.Color("14"))
	boldYellow  = bold().Foreground(lipgloss.Color("11"))
	boldMagenta = bold().Foreground(lipgloss.Color("13"))
	boldGreen   = bold().Foreground(lipgloss.Color("10"))
	boldRed     = bold().Foreground(lipgloss.Color("9"))
)

func bold() lipgloss.Style {
	return lipgloss.NewStyle().Bold(true)
}

func faint() lipgloss.Style {
	return lipgloss.NewStyle().Faint(true)
}

// Bold applies bold formatting to text.
func Bold(text string) string {
	return bolded.Render(text)
}

// Dim applies dim formatting to text.
func Dim(text string) string {
	return dimmed.Render(text)
}

// BoldCyan applies bold cyan color to text.
func BoldCyan(text string) string {
	return boldCyan.Render(text)
}

// BoldYellow applies bold yellow color to text.
func BoldYellow(text string) string {
	return boldYellow.Render(text)
}

// BoldMagenta applies bold magenta color to text.
func BoldMagenta(text string) string {
	return boldMagenta.Render(text)
}

// BoldGreen applies bold green color to text.
func BoldGreen(text string) string {
	return boldGreen.Render(text)
}

// BoldRed applies bold red color to text.
func BoldRed(text string) string {
	return boldRed.Render(text)
}
