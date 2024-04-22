package assembler_test

import (
	"reflect"
	"testing"

	"github.com/umutyalcinn/masm/internal/assembler"
)

func TestLexer(t *testing.T) {
    input := "ADD 0x10\nCLA\nCIL\nAND 0x1"

    expected := []assembler.Instruction {
        { Operation: "ADD", Operand: 0x10},
        { Operation: "CLA", Operand: 0x0},
        { Operation: "CIL", Operand: 0x0},
        { Operation: "AND", Operand: 0x1},
    }

    result := assembler.GetInstructions(input)

    if !reflect.DeepEqual(result, expected) {
        t.Fail()
    }
}
