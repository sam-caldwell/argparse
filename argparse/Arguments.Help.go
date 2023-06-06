package argparse

import (
	"fmt"
	"strings"
)

// Help - Print Argument usage (help text)
func (arg *Arguments) Help() (text string) {
	var lines []string
	post := func(indentSz int, line string) {
		indent := func(sz int) string {
			const indentation = "  "
			prefix := ""
			for i := 0; i < sz; i++ {
				prefix = indentation
			}
			return prefix
		}
		lines = append(lines, fmt.Sprintf("%s", indent(indentSz)+line))
	}
	printList := func(i int, list []any) {
		for _, line := range arg.preambleText.List() {
			post(i, line.(string))
		}
	}
	pad := func(c rune, a int, b int) (p string) {
		sz := b - a
		for i := 0; i < sz; i++ {
			p += string(c)
		}
		return p
	}
	calcMaxWidth := func(s string, cw int) int {
		if w := len(s); w > cw {
			return w
		}
		return cw
	}
	post(0, "")
	post(0, arg.programName)
	printList(1, arg.preambleText.List())

	post(0, "")
	post(0, "Usage:")
	post(1, fmt.Sprintf("%s [positional args] [optional args]", arg.programName))
	post(0, "")
	func() {
		//get positional arguments
		post(1, "Positional Arguments")
		pos := 0
		for name, argument := range arg.descriptors.List() {
			if (argument.GetShort() != "") || (argument.GetLong() != "") {
				continue
			}
			post(2, fmt.Sprintf(" %s - %s", name, argument.GetHelp()))
			pos++
		}
		if pos == 0 {
			post(2, "<none>")
		}
	}()
	post(0, "")

	func() {
		const space = ' '
		var shorts []string
		var longs []string
		var types []string
		var fields []string
		var defaults []string
		var help []string
		var shortWidth = 0
		var longWidth = 0
		var typeWidth = 0
		var fieldWidth = 0
		var defaultWidth = 0

		post(1, "Optional Arguments")
		shorts = append(shorts, "-h")
		longs = append(longs, "--help")
		types = append(types, "")
		fields = append(fields, "help")
		defaults = append(defaults, "")
		help = append(help, "show this help message")

		for name, argument := range arg.descriptors.List() {
			if (argument.GetShort() == "") && (argument.GetLong() == "") {
				continue
			}

			var localShort string = argument.GetShort()
			shortWidth = calcMaxWidth(localShort, shortWidth)
			shorts = append(shorts, localShort)

			var localLong string = argument.GetLong()
			longWidth = calcMaxWidth(localLong, longWidth)
			longs = append(longs, localLong)

			var localType string = argument.GetType()
			if localType == "Flag" {
				localType = ""
			}
			typeWidth = calcMaxWidth(localType, typeWidth) + 2
			types = append(types, localType)

			fields = append(fields, name)
			if thisWidth := len(name); thisWidth > fieldWidth {
				fieldWidth = thisWidth
			}

			var localDefault string = "default:" + argument.GetDefault()
			if localType == "" {
				localDefault = ""
			}
			defaultWidth = calcMaxWidth(localDefault, defaultWidth)
			defaults = append(defaults, localDefault)

			argHelp := argument.GetHelp()
			help = append(help, argHelp)
		}
		makeLine := func(f string, w int) string {
			return f + pad(space, len(f), w)
		}

		for i, _ := range shorts {
			thisShort := makeLine(shorts[i], shortWidth)
			thisLong := makeLine(longs[i], longWidth)
			thisType := func() string {
				if types[i] == "" {
					return makeLine("", typeWidth)
				}
				return makeLine("<"+types[i]+">", typeWidth)
			}()
			thisField := makeLine(fields[i], fieldWidth)
			thisDefault := makeLine(defaults[i], defaultWidth)

			post(3, fmt.Sprintf("%s %s %s %s (%s) [%s] - %s",
				thisShort, thisType, thisLong, thisType, thisField, thisDefault, help[i]))
		}
	}()

	post(0, "")

	for _, line := range arg.postscriptText.List() {
		post(1, line.(string))
	}
	post(0, "")
	return strings.Join(lines, "\n")
}
