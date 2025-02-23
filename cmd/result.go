package cmd

import (
	"fmt"
	"time"

	ics "github.com/arran4/golang-ical"
	"github.com/pterm/pterm"
)

type Result struct {
	Calendar *ics.Calendar
}

func (r Result) Print() {
	fmt.Println(r.Calendar)
}

func (r Result) Data() any {
	return r.Calendar
}

func (r Result) String() string {
	return r.Table()
}

func (r Result) Table() string {
	table := pterm.TableData{}
	table = append(table, []string{
		"Start",
		"End",
		"Summary",
	})

	for _, event := range r.Calendar.Events() {
		start, err := event.GetStartAt()
		if err != nil {
			start = time.Time{}
		}

		end, err := event.GetEndAt()
		if err != nil {
			end = time.Time{}
		}

		if start.Before(time.Now()) {
			continue
		}

		table = append(table, []string{
			start.Format("2006-01-02 15:04"),
			end.Format("2006-01-02 15:04"),
			event.GetProperty(ics.ComponentPropertySummary).Value,
		})
	}

	render, _ := pterm.DefaultTable.WithHasHeader().WithData(table).Srender()

	return render
}
