package TUI

import (
	"fmt"
	"log"

	. "github.com/PHillemans/LearnLangs/internal/FileActions"
	. "github.com/PHillemans/LearnLangs/internal/Scraper"
	. "github.com/PHillemans/LearnLangs/internal/Types"
	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
  TranslationHeader Translation
  Translations TranslationCollection
  CurrentIdx int
  Language string
  Spinner spinner.Model
  Paginator paginator.Model
}

func (m Model) Init() tea.Cmd {
  var cmds []tea.Cmd
  cmds = append(cmds, fetchLangData(m.Language), spinner.Tick)
  return tea.Batch(cmds...)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  switch msg := msg.(type) {

  case tea.KeyMsg:
    switch msg.String(){

    case "h":
      m.Paginator.PrevPage()
    return m, nil

    case "l":
      m.Paginator.NextPage()
      return m, nil

    case "q", "ctrl+c":
      return m, tea.Quit

    default:
      return m, nil
    }

  case TranslationMsg:
      m.TranslationHeader = msg[0]
      msg = msg[1:]
      m.Paginator.SetTotalPages(len(msg))
      m.Translations = TranslationCollection(msg)
      return m, nil

  }
  var firstCmd tea.Cmd
  var secondCmd tea.Cmd
  var cmds []tea.Cmd
  m.Spinner, firstCmd = m.Spinner.Update(msg)
  m.Paginator, secondCmd = m.Paginator.Update(msg)
  cmds = append(cmds, firstCmd, secondCmd)
  return m, tea.Batch(cmds...)
}

func (m Model) View() string {
  var s string
  if len(m.Translations) == 0 {
    s += fmt.Sprintf("%v Loading words for: %s\n", m.Spinner.View(), m.Language)
  } else {
    start, end := m.Paginator.GetSliceBounds(len(m.Translations))

    s += fmt.Sprintf("    |  %s  |  %s\n", 
      m.TranslationHeader.Original, m.TranslationHeader.Translation)

    for _, item := range m.Translations[start:end] {
      s += fmt.Sprintf(" %v  |  %s  |  %s\n",
        item.Order, item.Original, item.Translation)
    }
    s += fmt.Sprintf(m.Paginator.View())
  }
  return s
}

func fetchLangData(lang string) tea.Cmd {
  return func() tea.Msg {
    var mostCommonWords TranslationCollection

    if mostCommonWords = ReadDataFile(lang); mostCommonWords == nil {
      fetchedWords, err := RetreiveWordsForLang(lang)
      if err != nil {
        log.Panic(
          `this language doesn't have a words record, 
          please try another language`)
      }

      WriteToDataFile(fetchedWords, lang)

      mostCommonWords = fetchedWords
    }
    return TranslationMsg(mostCommonWords)
  }
}
