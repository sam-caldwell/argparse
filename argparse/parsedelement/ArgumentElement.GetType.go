package parsedelement

import (
	"github.com/sam-caldwell/argparse/v2/types"
)

func (arg *ArgumentElement) GetType() types.ArgTypes {
	return arg.typ
}
