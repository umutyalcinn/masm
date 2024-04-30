package assembler

import (
	"strconv"
	"strings"
)

func GetInstructions(input string) ([]Instruction, map[string]uint16) {

    instructions := []Instruction {}
    var current_address uint16 = 0
    tagAddressTable := make(map[string]uint16)

    lines := strings.Split(input, "\n")

    for _, line := range lines {

        parts := strings.Split(line, ",")

        var operation string 
        var tag string 
        var op_tag string 

        if len(parts) == 1 {
            tag = ""
            operation = parts[0]
        } else {
            tag = parts[0]
            operation = parts[1]
            tagAddressTable[tag] = current_address
        }

        words := strings.Fields(operation)

        var operand uint16

        if len(words) == 0 {
            continue
        }

        if len(words) > 1 {
            parsed, err := strconv.ParseUint(words[1], 0, 16)

            operand = uint16(parsed)

            if err != nil {
                op_tag = words[1]
            } else {
                op_tag = ""
            }
        } else {
            operand = 0
        }

        if words[0] == "ORG" {
            current_address = operand
        } else {
            instructions = append(instructions, Instruction {
                Operation: words[0],
                Operand: operand,
                Tag: tag,
                OpTag: op_tag,
                Address: current_address,
            })

            current_address += 1
        }

    }

    return instructions, tagAddressTable
}
