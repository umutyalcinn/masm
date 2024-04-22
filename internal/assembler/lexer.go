package assembler

import (
	"strconv"
	"strings"
)

func GetInstructions(input string) []Instruction {

    instructions := []Instruction {}

    lines := strings.Split(input, "\n")

    for _, line := range lines {
        words := strings.Fields(line)

        var operand uint16

        if len(words) == 0 {
            continue
        }

        if len(words) > 1 {
            parsed, err := strconv.ParseUint(words[1], 0, 16)

            operand = uint16(parsed)

            if err != nil {
                panic("Can't parse operand")
            }
        } else {
            operand = 0
        }

        instructions = append(instructions, Instruction {
            Operation: words[0],
            Operand: operand,
        })
    }

    return instructions
}
