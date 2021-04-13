package commands

import (
	"github.com/DavidGamba/go-getoptions/option"
	"github.com/hashicorp/vagrant-plugin-sdk/component"
	plugincore "github.com/hashicorp/vagrant-plugin-sdk/core"
	"github.com/hashicorp/vagrant-plugin-sdk/docs"
	"github.com/hashicorp/vagrant-plugin-sdk/terminal"
)

// DoThing is a Command implementation for myplugin
// It is a subcommand of myplugin
type DoThing struct {
	*Command
}

func (c *DoThing) ConfigSet(v interface{}) error {
	return nil
}

func (c *DoThing) CommandFunc() interface{} {
	return nil
}

func (c *DoThing) Config() (interface{}, error) {
	return &c.config, nil
}

func (c *DoThing) Documentation() (*docs.Documentation, error) {
	doc, err := docs.New(docs.FromConfig(&CommandConfig{}))
	if err != nil {
		return nil, err
	}
	return doc, nil
}

// SynopsisFunc implements component.Command
func (c *DoThing) SynopsisFunc() interface{} {
	return c.Synopsis
}

// HelpFunc implements component.Command
func (c *DoThing) HelpFunc() interface{} {
	return c.Help
}

// FlagsFunc implements component.Command
func (c *DoThing) FlagsFunc() interface{} {
	return c.Flags
}

// ExecuteFunc implements component.Command
func (c *DoThing) ExecuteFunc() interface{} {
	return c.Execute
}

// SubcommandFunc implements component.Command
func (c *DoThing) SubcommandsFunc() interface{} {
	return c.Subcommands
}

// CommandInfoFunc implements component.Command
func (c *DoThing) CommandInfoFunc() interface{} {
	return c.CommandInfo
}

func (c *DoThing) CommandInfo() *plugincore.CommandInfo {
	return &plugincore.CommandInfo{
		Name:     []string{"myplugin", "donothing"},
		Help:     c.Help(),
		Synopsis: c.Synopsis(),
		Flags:    c.Flags(),
	}
}

func (c *DoThing) Synopsis() string {
	return "Really important *stuff*"
}

func (c *DoThing) Help() string {
	return "I do really important work"
}

func (c *DoThing) Flags() []*option.Option {
	booltest := option.New("booltest", option.BoolType)
	booltest.Description = "a test flag for bools"
	booltest.DefaultStr = "true"
	booltest.Aliases = append(booltest.Aliases, "bt")

	stringflag := option.New("stringflag", option.StringType)
	stringflag.Description = "a test flag for strings"
	stringflag.DefaultStr = "message"
	stringflag.Aliases = append(stringflag.Aliases, "sf")

	return []*option.Option{booltest, stringflag}
}

func (c *DoThing) Subcommands() []string {
	return []string{}
}

func (c *DoThing) Execute(trm terminal.UI) int64 {
	trm.Output("Tricked ya! I actually do nothing :P")
	return 0
}

var (
	_ component.Command = (*DoThing)(nil)
)
