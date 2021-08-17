package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type Translation struct {
  Order int
  Original string
  Translation string
}

type TranslationCollection []Translation

type model struct {
  translations TranslationCollection
  currentIdx int
}

var initialModel = model{
    // Our to-do list is just a grocery list
    translations: make(TranslationCollection, 1000),

    // A map which indicates which choices are selected. We're using
    // the  map like a mathematical set. The keys refer to the indexes
    // of the `choices` slice, above.
    currentIdx: 0,
}

func (m model) Init() tea.Cmd {
    // Just return `nil`, which means "no I/O right now, please."
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {

    // Is it a key press?
    case tea.KeyMsg:

        // Cool, what was the actual key pressed?
        switch msg.String() {

        // These keys should exit the program.
        case "ctrl+c", "q":
          return m, tea.Quit
        // The "enter" key and the spacebar (a literal space) toggle
        // the selected state for the item that the cursor is pointing at.
        case "enter", " ":
          m.currentIdx++         
        }
    }

    // Return the updated model to the Bubble Tea runtime for processing.
    // Note that we're not returning a command.
    return m, nil
}

func (m model) View() string {
  // The header
  s := "Click next?\n\n"

  cursor := ">" // cursor!

  // Render the row
  s += fmt.Sprintf("%s %s\n", cursor, m.translations, m.currentIdx)

  // The footer
  s += "\nPress q to quit.\n"

  // Send the UI for rendering
  return s
}

var chosenLanguage string

func Execute() {
  var mostCommonWords TranslationCollection

  chosenLanguage = "italian"
  
  if mostCommonWords = ReadDataFile(chosenLanguage); mostCommonWords == nil {
    fetchedWords, err := RetreiveWordsForLang(chosenLanguage)
    if err != nil {
      ErrorLogger.Println(err.Error())
    }

    fetchedWords.WriteToDataFile(chosenLanguage)
    mostCommonWords = fetchedWords
  }
  initialModel.translations = mostCommonWords

  p := tea.NewProgram(initialModel)
  if err := p.Start(); err != nil {
    fmt.Printf("Alas, there's been an error: %v", err)
    os.Exit(1)
  }
}

func check(e error) {
  if e != nil {
    panic(e)
  }
}

