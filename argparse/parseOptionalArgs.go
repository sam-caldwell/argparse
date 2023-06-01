package argparse

import (
	"fmt"
	"github.com/sam-caldwell/go-argparse/v2/types"
	"os"
	"strconv"
	"strings"
)

// parseOptionalArgs - Parse the optional arguments
func (args *Arguments) parseOptionalArgs() (err error) {
	var thisToken string
	var thisArg types.ArgString
	var thisType types.ArgTypes
	var expectArg bool = true
	for {
		//Get the next value from os.Args
		if thisToken, err = pop(&os.Args); err != nil {
			return err
		}
		if expectArg {
			//Get argument term (--arg)
			thisArg = types.ArgString(thisToken)
			if param, found := args.optional[thisArg]; found {
				thisType = param.argType
			} else {
				return fmt.Errorf(errMissingArgument, thisArg)
			}
			if thisType == types.Flag {
				args.value[thisArg] = true
				expectArg = true //Get an argument next
			} else {
				expectArg = false //Get a value next
			}
		} else {
			// Get argument value
			expectArg = true //Get the next term (argument)
			switch thisType {
			case types.Boolean:
				value, err := strconv.ParseBool(thisToken)
				if err != nil {
					return err
				}
				args.value[thisArg] = value
			case types.Float:
				value, err := strconv.ParseFloat(thisToken, 64)
				if err != nil {
					return err
				}
				args.value[thisArg] = value
			case types.Integer:
				value, err := strconv.ParseInt(thisToken, 10, 64)
				if err != nil {
					return err
				}
				args.value[thisArg] = value
			case types.String:
				args.value[thisArg] = strings.TrimSpace(thisToken)
			default:
				return fmt.Errorf(errTypeMismatch, thisType.String(), thisArg)
			}
		}

	}

}
