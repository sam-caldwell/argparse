package parsed

import (
	"github.com/sam-caldwell/argparse/v2/argparse/parsedelement"
)

// Argument - a map of ArgumentElement objects
type Argument struct {
	data map[string]parsedelement.ArgumentElement
}
