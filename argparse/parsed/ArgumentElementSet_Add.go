package parsed

import "github.com/sam-caldwell/argparse/v2/types"

func (arg *ArgumentElementSet) Add(n string, t types.ArgTypes, v any) error {
	(*arg)[n] = ArgumentElement{
		_type:  t,
		_value: v,
	}
	return nil
}
