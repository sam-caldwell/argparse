package argparse

import (
	"log"
)

// Debug - Log a debug message
func (args *Arguments) Debug(msg ...any) {
	if args.debug {
		log.Printf("DEBUG: %v", msg...)
	}
}

// Debugf - Log a formatted debug message
func (args *Arguments) Debugf(format string, msg ...any) {
	if args.debug {
		log.Printf("DEBUG:"+format, msg...)
	}
}
