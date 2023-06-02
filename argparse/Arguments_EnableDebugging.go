package argparse

import "log"

func (args *Arguments) EnableDebugging() {
	log.Println("Debugging enabled")
	args.debug = true
}
