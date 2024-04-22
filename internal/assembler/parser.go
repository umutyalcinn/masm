package assembler

import "errors"

type Instruction struct {
    Operation string
    Operand uint16
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

func ParseInstruction(instruction Instruction) (Opcode, error) {
    val, nonMemory := nonMemInstructionMap[instruction.Operation]

    if nonMemory {
        return val, nil
    }

    val, memory := memInstructionMap [instruction.Operation]

    if memory {
        return val | Opcode(instruction.Operand), nil
    }

    return 0, errors.New("Undefined operation")
}
