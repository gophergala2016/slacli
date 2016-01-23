package main

import "strings"

// Parser is the main user input handler
type Parser struct {
	currentInput string
	IsAction     bool
	IsExit       bool
}

// ParseCurrentInput parses current input and decides what to do with it
func (p *Parser) ParseCurrentInput(s string) {
	p.currentInput = ""

	if idx := strings.Index(s, "/"); idx == 0 {
		p.IsAction = true
		p.proceedAction(removeEOL(s[1:]))
	} else {
		p.IsAction = false
		p.currentInput = cleanString(s)
	}
}

// GetCurrentInput retrun current input in case it's not an action
func (p *Parser) GetCurrentInput() string {
	return p.currentInput
}

func removeEOL(s string) string {
	return strings.TrimRight(s, "\n")
}

func cleanString(s string) string {
	return removeEOL(s)
}

func (p *Parser) proceedAction(s string) {
	action := strings.ToLower(s)

	p.IsExit = false
	switch action {
	case "exit", "quit":
		p.IsExit = true
		break
	default:
		break
	}
}
