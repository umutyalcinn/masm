# Project Description

MASM is an assembly language designed in Digital Design lecture in Turkish-German  University. This project contains a assembler binary. For future plans, there will be a C compiler, a VM, an operation system for that specific design of hardware.

# Instruction List

## Memory Reference Instructions

Format: <Symbol> <Address>

Symbol | Adrress
AND | 0x0xxx
ADD | 0x1xxx
LDA | 0x2xxx
STA | 0x3xxx
BUN | 0x4xxx
BAS | 0x5xxx
ISZ | 0x6xxx

## Register Reference Instructions

Format: <Symbol>

Symbol | Adrress
CLA | 0x7800
CLE | 0x7400
CMA | 0x7200
CME | 0x7100
CIR | 0x7080
CIL | 0x7040
INC | 0x7020
SPA | 0x7010
SNA | 0x7008
SZA | 0x7004
SZE | 0x7002
HLT | 0x7001

## Output Instrcutions

Symbol | Adrress
INP | 0xF800
OUT | 0xF400
SKI | 0xF200
SKO | 0xF100
ION | 0xF080
IOF | 0xF040

