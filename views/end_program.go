package views

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/EdoRguez/business-manager-admin/utils"
	"github.com/charmbracelet/lipgloss"
)

const (
	progressBarWidth  = 71
	progressFullChar  = "█"
	progressEmptyChar = "░"
)

var (
	keywordStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("211"))
	ticksStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("79"))
	progressEmpty = subtleStyle.Render(progressEmptyChar)

	// Gradient colors we'll use for the progress bar
	ramp = utils.MakeRampStyles("#B14FFF", "#00FFA3", progressBarWidth)
)

type EndProgramModel struct {
	optionSelected int
	isLoaded       bool
	progress       float64
	ticks          int
}

func NewEndProgramModal(optionSelected int, isLoaded bool, progress float64, ticks int) EndProgramModel {
	return EndProgramModel{
		optionSelected: optionSelected,
		isLoaded:       isLoaded,
		progress:       progress,
	}
}

func (m EndProgramModel) EndProgramView() string {
	var msg string

	switch m.optionSelected {
	case 0:
		msg = fmt.Sprintf("Carrot planting?\n\nCool, we'll need %s and %s...", keywordStyle.Render("libgarden"), keywordStyle.Render("vegeutils"))
	case 1:
		msg = fmt.Sprintf("A trip to the market?\n\nOkay, then we should install %s and %s...", keywordStyle.Render("marketkit"), keywordStyle.Render("libshopping"))
	case 2:
		msg = fmt.Sprintf("Reading time?\n\nOkay, cool, then we’ll need a library. Yes, an %s.", keywordStyle.Render("actual library"))
	default:
		msg = fmt.Sprintf("It’s always good to see friends.\n\nFetching %s and %s...", keywordStyle.Render("social-skills"), keywordStyle.Render("conversationutils"))
	}

	label := "Downloading..."
	if m.isLoaded {
		label = fmt.Sprintf("Downloaded. Exiting in %s seconds...", ticksStyle.Render(strconv.Itoa(m.ticks)))
	}

	return msg + "\n\n" + label + "\n" + progressbar(m.progress) + "%"
}

func progressbar(percent float64) string {
	w := float64(progressBarWidth)

	fullSize := int(math.Round(w * percent))
	var fullCells string
	for i := 0; i < fullSize; i++ {
		fullCells += ramp[i].Render(progressFullChar)
	}

	emptySize := int(w) - fullSize
	emptyCells := strings.Repeat(progressEmpty, emptySize)

	return fmt.Sprintf("%s%s %3.0f", fullCells, emptyCells, math.Round(percent*100))
}
