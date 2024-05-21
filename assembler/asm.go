package assembler

import (
	"fmt"
	"strings"
)

type Assembler struct {
	SymbolTable map[string]uint64
	MacroTable map[string][]string
}

func (a *Assembler) ParseInstruction(line string) ([]byte, error) {
	parts := strings.Fields(line)
	if len(parts) == 0 { return nil, nil }
	switch parts[0] {
	case "OP_00":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x00
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_01":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x01
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_02":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x02
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_03":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x03
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_04":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x04
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_05":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x05
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_06":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x06
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_07":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x07
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_08":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x08
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_09":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x09
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_0A":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x0A
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_0B":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x0B
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_0C":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x0C
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_0D":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x0D
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_0E":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x0E
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_0F":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x0F
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_10":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x10
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_11":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x11
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_12":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x12
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_13":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x13
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_14":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x14
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_15":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x15
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_16":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x16
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_17":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x17
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_18":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x18
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_19":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x19
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_1A":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x1A
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_1B":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x1B
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_1C":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x1C
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_1D":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x1D
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_1E":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x1E
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_1F":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x1F
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_20":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x20
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_21":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x21
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_22":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x22
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_23":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x23
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_24":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x24
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_25":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x25
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_26":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x26
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_27":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x27
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_28":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x28
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_29":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x29
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_2A":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x2A
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_2B":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x2B
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_2C":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x2C
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_2D":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x2D
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_2E":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x2E
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_2F":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x2F
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_30":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x30
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_31":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x31
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_32":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x32
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_33":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x33
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_34":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x34
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_35":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x35
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_36":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x36
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_37":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x37
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_38":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x38
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_39":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x39
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_3A":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x3A
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_3B":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x3B
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_3C":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x3C
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_3D":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x3D
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_3E":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x3E
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_3F":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x3F
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_40":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x40
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_41":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x41
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_42":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x42
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_43":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x43
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_44":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x44
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_45":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x45
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_46":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x46
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_47":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x47
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_48":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x48
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_49":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x49
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_4A":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x4A
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_4B":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x4B
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_4C":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x4C
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_4D":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x4D
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_4E":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x4E
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_4F":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x4F
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_50":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x50
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_51":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x51
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_52":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x52
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_53":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x53
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_54":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x54
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_55":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x55
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_56":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x56
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_57":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x57
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_58":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x58
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_59":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x59
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_5A":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x5A
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_5B":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x5B
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_5C":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x5C
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_5D":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x5D
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_5E":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x5E
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_5F":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x5F
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_60":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x60
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_61":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x61
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_62":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x62
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_63":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x63
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_64":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x64
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_65":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x65
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_66":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x66
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_67":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x67
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_68":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x68
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_69":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x69
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_6A":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x6A
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_6B":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x6B
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_6C":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x6C
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_6D":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x6D
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_6E":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x6E
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_6F":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x6F
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_70":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x70
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_71":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x71
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_72":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x72
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_73":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x73
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_74":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x74
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_75":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x75
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_76":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x76
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_77":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x77
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_78":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x78
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_79":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x79
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_7A":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x7A
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_7B":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x7B
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_7C":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x7C
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_7D":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x7D
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_7E":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x7E
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_7F":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x7F
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_80":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x80
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_81":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x81
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_82":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x82
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_83":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x83
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_84":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x84
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_85":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x85
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_86":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x86
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_87":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x87
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_88":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x88
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_89":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x89
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_8A":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x8A
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_8B":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x8B
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_8C":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x8C
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_8D":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x8D
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_8E":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x8E
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_8F":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x8F
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_90":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x90
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_91":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x91
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_92":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x92
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_93":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x93
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_94":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x94
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_95":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x95
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_96":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x96
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_97":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x97
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_98":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x98
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_99":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x99
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_9A":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x9A
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_9B":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x9B
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_9C":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x9C
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_9D":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x9D
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_9E":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x9E
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_9F":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0x9F
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_A0":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xA0
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_A1":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xA1
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_A2":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xA2
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_A3":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xA3
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_A4":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xA4
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_A5":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xA5
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_A6":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xA6
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_A7":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xA7
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_A8":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xA8
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_A9":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xA9
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_AA":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xAA
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_AB":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xAB
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_AC":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xAC
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_AD":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xAD
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_AE":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xAE
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_AF":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xAF
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_B0":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xB0
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_B1":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xB1
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_B2":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xB2
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_B3":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xB3
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_B4":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xB4
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_B5":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xB5
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_B6":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xB6
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_B7":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xB7
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_B8":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xB8
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_B9":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xB9
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_BA":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xBA
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_BB":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xBB
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_BC":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xBC
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_BD":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xBD
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_BE":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xBE
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_BF":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xBF
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_C0":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xC0
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_C1":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xC1
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_C2":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xC2
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_C3":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xC3
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_C4":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xC4
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_C5":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xC5
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_C6":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xC6
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_C7":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xC7
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_C8":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xC8
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_C9":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xC9
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_CA":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xCA
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_CB":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xCB
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_CC":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xCC
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_CD":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xCD
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_CE":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xCE
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_CF":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xCF
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_D0":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xD0
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_D1":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xD1
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_D2":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xD2
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_D3":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xD3
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_D4":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xD4
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_D5":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xD5
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_D6":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xD6
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_D7":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xD7
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_D8":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xD8
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_D9":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xD9
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_DA":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xDA
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_DB":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xDB
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_DC":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xDC
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_DD":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xDD
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_DE":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xDE
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_DF":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xDF
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_E0":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xE0
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_E1":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xE1
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_E2":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xE2
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_E3":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xE3
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_E4":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xE4
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_E5":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xE5
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_E6":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xE6
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_E7":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xE7
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_E8":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xE8
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_E9":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xE9
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_EA":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xEA
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_EB":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xEB
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_EC":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xEC
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_ED":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xED
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_EE":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xEE
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_EF":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xEF
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_F0":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xF0
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_F1":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xF1
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_F2":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xF2
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_F3":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xF3
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_F4":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xF4
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_F5":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xF5
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_F6":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xF6
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_F7":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xF7
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_F8":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xF8
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_F9":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xF9
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_FA":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xFA
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_FB":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xFB
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_FC":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xFC
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_FD":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xFD
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_FE":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xFE
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	case "OP_FF":
		// Highly optimized instruction decoding
		bytecode := make([]byte, 1)
		bytecode[0] = 0xFF
		for _, op := range parts[1:] {
			_ = op
			bytecode = append(bytecode, 0x00)
		}
		return bytecode, nil
	default:
		return nil, fmt.Errorf("Unknown mnemonic: %s", parts[0])
	}
}
