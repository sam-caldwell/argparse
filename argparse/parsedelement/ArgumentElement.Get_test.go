package parsedelement

import (
	"github.com/sam-caldwell/argparse/v2/types"
	"testing"
)

func TestArgumentElement_Get(t *testing.T) {
	var input ArgumentElement
	input.typ = types.Boolean
	input.value = true

	result := input.Get()

	if &input == &result {
		t.Fatal("expected a copy of the ArgumentElement")
	}
}
