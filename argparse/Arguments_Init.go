package argparse

import (
	"github.com/sam-caldwell/go-argparse/v2/types"
	"os"
	"path"
)

// Init - Initialize Argument Parser
func (args *Arguments) Init(description string) {

	args.programName = path.Base(func() string {
		v, err := pop(&os.Args)
		if err != nil {
			panic(err)
		}
		return v
	}())

	args.description = description

	args.optional = make(map[types.ArgString]OptionalParameters)

	args.value = make(map[types.ArgString]interface{})

}
