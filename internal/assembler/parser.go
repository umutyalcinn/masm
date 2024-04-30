package assembler

import "errors"

type Instruction struct {
    Address uint16
    Operation string
    Operand uint16
    Tag string
    OpTag string
}

type Opcode uint16

var nonMemInstructionMap = map[string]Opcode {
    "CLA": 0x7800,
    "CLE": 0x7400,
    "CMA": 0x7200,
    "CME": 0x7100,
    "CIR": 0x7080,
    "CIL": 0x7040,
    "INC": 0x7020,
    "SPA": 0x7010,
    "SNA": 0x7008,
    "SZA": 0x7004,
    "SZE": 0x7002,
    "HLT": 0x7001,
    "INP": 0xF800,
    "OUT": 0xF400,
    "SKI": 0xF200,
    "SKO": 0xF100,
    "ION": 0xF080,
    "IOF": 0xF040,
}

var memInstructionMap = map[string]Opcode {
    "AND": 0x0000,
    "ADD": 0x1000,
    "LDA": 0x2000,
    "STA": 0x3000,
    "BUN": 0x4000,
    "BSA": 0x5000,
    "ISZ": 0x6000,
    "ANDI": 0x8000,
    "ADDI": 0x9000,
    "LDAI": 0xA000,
    "STAI": 0xB000,
    "BUNI": 0xC000,
    "BSAI": 0xD000,
    "ISZI": 0xE000,
}

const (
    PSEDUO_UNKNOWN = 0
    PSEDUO_ORIGIN = 1
    PSEDUO_HEX = 2
)

func ParsePseudoInstruction(instruction Instruction) (int, uint16) {
    switch (instruction.Operation) {
        case "ORG": {
            return PSEDUO_ORIGIN, instruction.Operand
        }
        case "HEX": {
            return PSEDUO_HEX, instruction.Operand
        }
        default: {
            return PSEDUO_UNKNOWN, 0
        }
    }
}

func ParseInstruction(instruction Instruction) (Opcode, int, uint16, error) {
    val, nonMemory := nonMemInstructionMap[instruction.Operation]

    if nonMemory {
        return val, 0, 0, nil
    }

    val, memory := memInstructionMap [instruction.Operation]

    if memory {
        return val | Opcode(instruction.Operand), 0, 0, nil
    }

    pseudo_instruction, pseduo_operand := ParsePseudoInstruction(instruction)

    if pseudo_instruction != PSEDUO_UNKNOWN {
        return 0, pseudo_instruction, pseduo_operand, nil
    }

    return 0, 0, 0, errors.New("Undefined operation")
}
