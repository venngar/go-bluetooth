package types

import (
	"fmt"
	"strings"
)

type ApiGroup struct {
	FileName    string
	Name        string
	Description string
	Api         []*Api
	debug       bool
}

type Api struct {
	Title       string
	Description string
	Service     string
	Interface   string
	ObjectPath  string
	Methods     []*Method
	// those are currently avail only in health-api
	Signals    []*Method
	Properties []*Property
}

type Flag int

const (
	FlagReadOnly Flag = iota + 1
	FlagWriteOnly
	FlagReadWrite
	FlagExperimental
	FlagOptional
	FlagServerOnly
)

type Arg struct {
	Type string
	Name string
}

func (a *Arg) String() string {
	return fmt.Sprintf("%s %s", a.Type, a.Name)
}

type Method struct {
	Name       string
	ReturnType string
	Args       []Arg
	Errors     []string
	Docs       string
}

func (m *Method) String() string {
	args := []string{}
	for _, arg := range m.Args {
		args = append(args, arg.String())
	}
	return fmt.Sprintf("%s %s(%s)", m.ReturnType, m.Name, strings.Join(args, ", "))
}

type Property struct {
	Name  string
	Type  string
	Docs  string
	Flags []Flag
}

func (p *Property) String() string {
	flags := []string{}
	for _, flag := range p.Flags {
		flagLabel := ""
		switch flag {
		case FlagReadOnly:
			flagLabel = "readonly"
			break
		case FlagWriteOnly:
			flagLabel = "writeonly"
			break
		case FlagReadWrite:
			flagLabel = "readwrite"
			break
		case FlagExperimental:
			flagLabel = "experimental"
			break
		case FlagOptional:
			flagLabel = "optional"
			break
		case FlagServerOnly:
			flagLabel = "server-only"
			break
		}
		if flagLabel != "" {
			flags = append(flags, flagLabel)
		}
	}

	flagsStr := ""
	if len(flags) > 0 {
		flagsStr = fmt.Sprintf("[%s]", strings.Join(flags, ", "))
	}

	return fmt.Sprintf("%s %s %s", p.Type, p.Name, flagsStr)
}
