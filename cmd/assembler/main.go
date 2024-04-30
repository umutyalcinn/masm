package main

import (
	"encoding/binary"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/umutyalcinn/masm/internal/assembler"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Printf("usage: [<options>] <inputs>\n")
    }

    var args = DefaultArgs()

    if err := args.ParseArgs(os.Args); err != nil {
        fmt.Printf("Error parsing arguments: %v\n", err)
    }

    if len(args.Inputs) == 0 {
        fmt.Printf("No input provided\n")
        os.Exit(1)
    } 

    for _, v := range args.Inputs {
        content, err := os.ReadFile(v)

        if err != nil {
            fmt.Printf("Could not open file: %s\n", v)
        }

        instructions, mapAddressTable := assembler.GetInstructions(string(content))

        for i, instruction := range instructions {
            if instruction.Operand == 0 && instruction.OpTag != "" {
                address, ok := mapAddressTable[instruction.OpTag]
                if ok {
                    instruction.Operand = address
                    instructions[i] = instruction
                }
            }
        }

        output := make([]byte, 4096)

        for _, instruction := range instructions {
            opcode, pseudo_instruction, pseudo_operand, err := assembler.ParseInstruction(instruction)
            address := instruction.Address * 2

            if err != nil {
                fmt.Printf("Error parsing intruction %s\n", instruction.Operation)
            }

            switch pseudo_instruction {
                case assembler.PSEDUO_UNKNOWN: {
                    output[address] = byte(opcode)
                    output[address + 1] = byte(opcode >> 8)
                    break
                }
                case assembler.PSEDUO_HEX: {
                    output[address] = byte(pseudo_operand)
                    output[address + 1] = byte(pseudo_operand >> 8)
                }
            }
        }

        out_file, err := os.Create(args.Output)

        if err != nil {
            fmt.Printf("Could not open output file: %s", args.Output)
        }

        binary.Write(out_file, binary.LittleEndian, output)
    }
}

type Args struct {
    Command string
    Inputs []string
    Output string
}

func DefaultArgs() Args {
    return Args {
        Command: "",
        Inputs: make([]string, 0),
        Output: "./out",
    }
}

func (s *Args) ParseArgs(args []string) error {
    i := 0

    s.Command = args[i]
    i += 1

    for i < len(args) {
        currArg := args[i]
        switch currArg {
            case "-o": {
                i += 1
                s.Output = args[i]
                i += 1
                break
            }
            case "--version": {
                fmt.Println("version: 1.0.0")
                i += 1
                os.Exit(0)
                break
            }

            default: {
                if strings.HasPrefix(currArg, "-") {
                    return errors.New("Invalid argument")
                }

                s.Inputs = append(s.Inputs, currArg)
                i += 1
            }
        }
    }

    return nil
}
