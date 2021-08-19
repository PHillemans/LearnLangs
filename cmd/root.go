package cmd

import (
	"fmt"
	"os"

	. "github.com/PHillemans/LearnLangs/internal/TUI"
	. "github.com/PHillemans/LearnLangs/internal/Types"
	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var root = &cobra.Command{
  Use: "langs <language>",
  Short: "langs fetches the 1000 most used words in a language to learn words",
  Args: cobra.MaximumNArgs(1),
  RunE: func(cmd *cobra.Command, args []string) error {
    var err error
    var language string

    if len(args) > 0 {
      language = args[0]
    } else {
      fmt.Println("Please specify a language!")
      os.Exit(0)
    }

    var collection TranslationCollection

    s := spinner.NewModel()
    s.Spinner = spinner.Line

    p := paginator.NewModel()
    p.Type = paginator.Dots
    p.PerPage = 15
    p.SetTotalPages(1000)

    translation := Model{
      Translations: collection,
      CurrentIdx: 0,
      Language: language,
      Spinner: s,
      Paginator: p,
    }

    program := tea.NewProgram(translation, tea.WithAltScreen())
    err = program.Start()
    return err
  },
}

func Execute() {
  err := root.Execute()
  if err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }
}

