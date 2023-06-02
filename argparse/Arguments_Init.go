package argparse

import (
	"github.com/sam-caldwell/go-argparse/v2/types"
	"os"
	"path"
)

// Init - Initialize Argument Parser (memory plus parameter list)
func (args *Arguments) Init(params ...string) {

	args.programName = path.Base(func() string {
		v, err := pop(&os.Args)
		if err != nil {
			panic(err)
		}
		return v
	}())

	for _, parameter := range params {
		args.descriptor = append(args.descriptor, parameter)
	}

	args.optional = make(map[types.ArgString]OptionalParameters)

	args.value = make(map[types.ArgString]interface{})

}
