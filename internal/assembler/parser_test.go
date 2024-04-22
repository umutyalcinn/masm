package assembler_test

import (
	"testing"

	"github.com/umutyalcinn/masm/internal/assembler"
)

func TestParser(t *testing.T) {
    instructions := []assembler.Instruction {
        { Operation: "AND", Operand: 0 },
        { Operation: "ADD", Operand: 0 },
        { Operation: "LDA", Operand: 0x333 },
        { Operation: "STA", Operand: 0 },
        { Operation: "BUN", Operand: 0 },
        { Operation: "BSA", Operand: 0 },
        { Operation: "ISZ", Operand: 0 },
        { Operation: "ANDI", Operand: 0 },
        { Operation: "ADDI", Operand: 0 },
        { Operation: "LDAI", Operand: 0x333 },
        { Operation: "STAI", Operand: 0 },
        { Operation: "BUNI", Operand: 0 },
        { Operation: "BSAI", Operand: 0 },
        { Operation: "ISZI", Operand: 0 },
        { Operation: "CLA", Operand: 0 },
        { Operation: "CLE", Operand: 0 },
        { Operation: "CMA", Operand: 0 },
        { Operation: "CME", Operand: 0 },
        { Operation: "CIR", Operand: 0 },
        { Operation: "CIL", Operand: 0 },
        { Operation: "INC", Operand: 0 },
        { Operation: "SPA", Operand: 0 },
        { Operation: "SNA", Operand: 0 },
        { Operation: "SZA", Operand: 0 },
        { Operation: "SZE", Operand: 0 },
        { Operation: "HLT", Operand: 0 },
        { Operation: "INP", Operand: 0 },
        { Operation: "OUT", Operand: 0 },
        { Operation: "SKI", Operand: 0 },
        { Operation: "SKO", Operand: 0 },
        { Operation: "ION", Operand: 0 },
        { Operation: "IOF", Operand: 0 },
    }

    expected := []assembler.Opcode {
        0x0000,
        0x1000,
        0x2333,
        0x3000,
        0x4000,
        0x5000,
        0x6000,
        0x8000,
        0x9000,
        0xA333,
        0xB000,
        0xC000,
        0xD000,
        0xE000,
        0x7800,
        0x7400,
        0x7200,
        0x7100,
        0x7080,
        0x7040,
        0x7020,
        0x7010,
        0x7008,
        0x7004,
        0x7002,
        0x7001,
        0xF800,
        0xF400,
        0xF200,
        0xF100,
        0xF080,
        0xF040,
    }

    for i, _ := range instructions {
        res, err := assembler.ParseInstruction(instructions[i])

        if err != nil {
            t.Error(err)
        }

        if res != expected[i] {
            t.Fail()
        }
    }
}
