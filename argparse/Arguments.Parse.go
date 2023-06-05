package argparse

// Parse - Parse os.Args
func (arg *Arguments) Parse() *Arguments {
	//var foundOptionalArgs bool = false
	//Remove arg[0] - program name
	//if arg.hasHelpFlag() {
	//	arg.value[argHelpName] = true
	//}
	//
	//for pos, descriptor := range os.Args[1:] {
	//	if foundOptionalArgs {
	//		//We cannot do positional arguments
	//		if thisArg, ok := arg.searchOptionalArguments(&descriptor); ok {
	//			if err := arg.value.Set(thisArg); err != nil {
	//
	//			}
	//		}
	//
	//	} else {
	//		//We can do either positional or optional arguments
	//	}
	//}
	//for pos := 0; pos < len(os.Args[1:]); pos++ {
	//	thisArg := os.Args[pos]
	//
	//	if err := valid.IsShortArg(&os.Args[pos]); err == nil {
	//		foundOptionalArgs = true
	//	}
	//	if err := valid.IsLongArg(&os.Args[pos]); err == nil {
	//		foundOptionalArgs = true
	//	}
	//	if foundOptionalArgs {
	//
	//	} else {
	//		arg.value = os.Args[pos]
	//	}
	//}
	return arg
}
