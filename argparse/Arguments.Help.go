package argparse

import (
	"fmt"
	"github.com/sam-caldwell/argparse/v2/argparse/types"
	"github.com/sam-caldwell/utilities/v2"
	"sort"
	"strings"
)

// Help - Print Argument usage (help text)
func (arg *Arguments) Help() (text string) {

	const space = ' '
	const empty = ""

	//this is our output buffer
	var lines []string

	//post text to our lines buffer
	post := func(indentSz int, line string) {
		lines = append(lines, utilities.LeftPad(space, line, len(line)+indentSz))
	}
	newLine := func() {
		post(0, empty)
	}

	type ListFunc func() []any
	printList := func(i int, f ListFunc) {
		for _, line := range f() {
			post(i, line.(string))
		}
	}

	/*
	 * Generate output
	 */
	newLine()
	post(0, arg.programName)
	printList(2, arg.preambleText.List)
	newLine()
	post(0, "Usage:")
	post(2, fmt.Sprintf("%s [positional args] [optional args]", arg.programName))
	newLine()
	//start: show positional arguments
	post(2, "Positional Arguments")
	if positionalArgs := arg.descriptors.ListPositionalArgs("%s [%s] - %s"); len(positionalArgs) == 0 {
		post(4, "<none>")
	} else {
		printList(4, func() []any { return positionalArgs })
	}
	newLine()
	//finished: show positional arguments

	//start: show optional arguments
	post(1, "Optional Arguments")

	type optionalArgs struct {
		short  string
		long   string
		typ    string
		name   string
		dValue string
		help   string
	}
	var rows []optionalArgs

	var columnSizes struct {
		short  int
		long   int
		typ    int
		dValue int
		help   int
	}
	rows = append(rows, optionalArgs{
		short:  "-h",
		long:   "--help",
		typ:    "",
		dValue: "",
		help:   "show this help message",
	})
	//Get our columnar data
	for _, argument := range arg.descriptors.List() {
		if (argument.GetShort() == "") && (argument.GetLong() == "") {
			continue
		}
		rows = append(rows, optionalArgs{
			short: argument.GetShort(),
			long:  argument.GetLong(),
			typ: func() string {
				a := argument.GetType()
				if a == types.Flag {
					return ""
				}
				return "[" + a.String() + "]"
			}(),
			dValue: func() string {
				if argument.GetType() == types.Flag {
					return ""
				}
				return "[default:" + argument.GetDefault() + "]"
			}(),
			help: argument.GetHelp(),
		})
	}
	//Sort the data
	sort.Slice(rows, func(i, j int) bool {
		// Specify your sorting criteria here
		// For example, sorting by the "short" field in ascending order
		return rows[i].short < rows[j].short
	})
	// Print the sorted rows
	for _, row := range rows {
		fmt.Println(row)
	}
	//Calculate column widths
	for _, row := range rows {
		if sz := len(row.short); sz > columnSizes.short {
			columnSizes.short = sz
		}
		if sz := len(row.long); sz > columnSizes.long {
			columnSizes.long = sz
		}
		if sz := len(row.typ); sz > columnSizes.typ {
			columnSizes.typ = sz
		}
		if sz := len(row.dValue); sz > columnSizes.dValue {
			columnSizes.dValue = sz
		}
		if sz := len(row.help); sz > columnSizes.help {
			columnSizes.help = sz
		}
	}
	for _, row := range rows {
		showIfNotEmpty := func(showThis string, thisNotEmpty string) string {
			if thisNotEmpty != "" {
				return showThis
			}
			return utilities.RepeatChars(string(space), len(showThis))
		}
		post(4, fmt.Sprintf("%s %s %s %s %s - %s",
			utilities.RightPad(space, row.short, columnSizes.short),
			utilities.RightPad(space, showIfNotEmpty(row.typ, row.short), columnSizes.typ),
			utilities.RightPad(space, row.long, columnSizes.long),
			utilities.RightPad(space, showIfNotEmpty(row.typ, row.long), columnSizes.typ),
			utilities.RightPad(space, row.dValue, columnSizes.dValue),
			utilities.RightPad(space, row.help, columnSizes.help)))
	}
	newLine()

	printList(1, arg.postscriptText.List)

	newLine()
	return strings.Join(lines, "\n")
}
