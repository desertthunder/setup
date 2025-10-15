package cmd

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

// Style definitions for terminal output using lipgloss
var (
	bolded      = bold()
	dimmed      = faint()
	boldCyan    = bold().Foreground(lipgloss.Color("14"))
	boldYellow  = bold().Foreground(lipgloss.Color("11"))
	boldMagenta = bold().Foreground(lipgloss.Color("13"))
	boldGreen   = bold().Foreground(lipgloss.Color("10"))
	boldRed     = bold().Foreground(lipgloss.Color("9"))
)

// Style represents different output styles for printer methods.
type Style int

const (
	StyleInfo Style = iota
	StyleInfoC
	StyleSuccess
	StyleWarn
	StyleErr
	StyleDim
)

var Print printer = printer{}

// printer provides styled terminal output methods for consistent formatting across the CLI application.
type printer struct{}

// Info prints informational messages without special formatting.
//
// When no arguments are provided, prints an empty line.
func (p printer) Info(a ...any) {
	if len(a) == 0 {
		fmt.Println()
		return
	}
	fmt.Println(a...)
}

// InfoC prints informational messages in bold cyan.
func (p printer) InfoC(a ...any) {
	if len(a) == 0 {
		return
	}
	msg := fmt.Sprint(a...)
	fmt.Println(BoldCyan(msg))
}

// Warn prints warning messages in bold yellow.
func (p printer) Warn(a ...any) {
	if len(a) == 0 {
		return
	}
	msg := fmt.Sprint(a...)
	fmt.Println(BoldYellow(msg))
}

// Err prints error messages in bold red.
func (p printer) Err(a ...any) {
	if len(a) == 0 {
		return
	}
	msg := fmt.Sprint(a...)
	fmt.Println(BoldRed(msg))
}

// Success prints success messages in bold green.
func (p printer) Success(a ...any) {
	if len(a) == 0 {
		return
	}
	msg := fmt.Sprint(a...)
	fmt.Println(BoldGreen(msg))
}

// Dimmed prints dimmed/faint text.
func (p printer) Dimmed(a ...any) {
	if len(a) == 0 {
		return
	}
	msg := fmt.Sprint(a...)
	fmt.Println(Dim(msg))
}

// NewLns prints styled message followed by an extra newline for adding visual spacing in output.
//
// The style parameter determines which printer method to dispatch to.
func (p printer) NewLns(style Style, a ...any) {
	switch style {
	case StyleInfo:
		p.Info(a...)
	case StyleInfoC:
		p.InfoC(a...)
	case StyleSuccess:
		p.Success(a...)
	case StyleWarn:
		p.Warn(a...)
	case StyleErr:
		p.Err(a...)
	case StyleDim:
		p.Dimmed(a...)
	default:
		p.Info(a...)
	}
	fmt.Println()
}

func (p printer) Beforeln(style Style, a ...any) {
	fmt.Println()
	switch style {
	case StyleInfo:
		p.Info(a...)
	case StyleInfoC:
		p.InfoC(a...)
	case StyleSuccess:
		p.Success(a...)
	case StyleWarn:
		p.Warn(a...)
	case StyleErr:
		p.Err(a...)
	case StyleDim:
		p.Dimmed(a...)
	default:
		p.Info(a...)
	}
}

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
