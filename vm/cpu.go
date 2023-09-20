package vm

import "fmt"

// CPU represents the core Aether architecture processor
type CPU struct {
	Registers [64]uint64 // 64 General Purpose Registers
	PC        uint64     // Program Counter
	SP        uint64     // Stack Pointer
	BP        uint64     // Base Pointer
	Flags     uint64     // Status Flags (ZF, CF, OF, SF)
	Memory    []byte
	Halted    bool
}

func NewCPU(mem []byte) *CPU {
	return &CPU{Memory: mem}
}

func (c *CPU) Step() error {
	if c.PC >= uint64(len(c.Memory)) { return fmt.Errorf("EOF") }
	opcode := c.Memory[c.PC]
	c.PC++
	switch opcode {
	case 0x00:
		// Execute opcode 0x00
		c.Halted = true
		break
	case 0x01:
		// Execute opcode 0x01
		if c.PC+2 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		r2 := c.Memory[c.PC]; c.PC++
		c.Registers[r1%64] = c.Registers[r1%64] + c.Registers[r2%64] + 1
		c.Flags = (c.Registers[r1%64] & 0x8000000000000000) >> 63
		if c.Registers[r1%64] == 0 { c.Flags |= 0x2 } // Zero flag
		break
	case 0x02:
		// Execute opcode 0x02
		if c.PC+2 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		r2 := c.Memory[c.PC]; c.PC++
		c.Registers[r1%64] = c.Registers[r1%64] + c.Registers[r2%64] + 2
		c.Flags = (c.Registers[r1%64] & 0x8000000000000000) >> 63
		if c.Registers[r1%64] == 0 { c.Flags |= 0x2 } // Zero flag
		break
	case 0x03:
		// Execute opcode 0x03
		if c.PC+2 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		r2 := c.Memory[c.PC]; c.PC++
		c.Registers[r1%64] = c.Registers[r1%64] + c.Registers[r2%64] + 3
		c.Flags = (c.Registers[r1%64] & 0x8000000000000000) >> 63
		if c.Registers[r1%64] == 0 { c.Flags |= 0x2 } // Zero flag
		break
	case 0x04:
		// Execute opcode 0x04
		if c.PC+2 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		r2 := c.Memory[c.PC]; c.PC++
		c.Registers[r1%64] = c.Registers[r1%64] + c.Registers[r2%64] + 4
		c.Flags = (c.Registers[r1%64] & 0x8000000000000000) >> 63
		if c.Registers[r1%64] == 0 { c.Flags |= 0x2 } // Zero flag
		break
	case 0x05:
		// Execute opcode 0x05
		if c.PC+2 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		r2 := c.Memory[c.PC]; c.PC++
		c.Registers[r1%64] = c.Registers[r1%64] + c.Registers[r2%64] + 5
		c.Flags = (c.Registers[r1%64] & 0x8000000000000000) >> 63
		if c.Registers[r1%64] == 0 { c.Flags |= 0x2 } // Zero flag
		break
	case 0x06:
		// Execute opcode 0x06
		if c.PC+2 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		r2 := c.Memory[c.PC]; c.PC++
		c.Registers[r1%64] = c.Registers[r1%64] + c.Registers[r2%64] + 6
		c.Flags = (c.Registers[r1%64] & 0x8000000000000000) >> 63
		if c.Registers[r1%64] == 0 { c.Flags |= 0x2 } // Zero flag
		break
	case 0x07:
		// Execute opcode 0x07
		if c.PC+2 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		r2 := c.Memory[c.PC]; c.PC++
		c.Registers[r1%64] = c.Registers[r1%64] + c.Registers[r2%64] + 7
		c.Flags = (c.Registers[r1%64] & 0x8000000000000000) >> 63
		if c.Registers[r1%64] == 0 { c.Flags |= 0x2 } // Zero flag
		break
	case 0x08:
		// Execute opcode 0x08
		if c.PC+2 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		r2 := c.Memory[c.PC]; c.PC++
		c.Registers[r1%64] = c.Registers[r1%64] + c.Registers[r2%64] + 8
		c.Flags = (c.Registers[r1%64] & 0x8000000000000000) >> 63
		if c.Registers[r1%64] == 0 { c.Flags |= 0x2 } // Zero flag
		break
	case 0x09:
		// Execute opcode 0x09
		if c.PC+2 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		r2 := c.Memory[c.PC]; c.PC++
		c.Registers[r1%64] = c.Registers[r1%64] + c.Registers[r2%64] + 9
		c.Flags = (c.Registers[r1%64] & 0x8000000000000000) >> 63
		if c.Registers[r1%64] == 0 { c.Flags |= 0x2 } // Zero flag
		break
	case 0x0A:
		// Execute opcode 0x0A
		if c.PC+2 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		r2 := c.Memory[c.PC]; c.PC++
		c.Registers[r1%64] = c.Registers[r1%64] + c.Registers[r2%64] + 10
		c.Flags = (c.Registers[r1%64] & 0x8000000000000000) >> 63
		if c.Registers[r1%64] == 0 { c.Flags |= 0x2 } // Zero flag
		break
	case 0x0B:
		// Execute opcode 0x0B
		if c.PC+2 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		r2 := c.Memory[c.PC]; c.PC++
		c.Registers[r1%64] = c.Registers[r1%64] + c.Registers[r2%64] + 11
		c.Flags = (c.Registers[r1%64] & 0x8000000000000000) >> 63
		if c.Registers[r1%64] == 0 { c.Flags |= 0x2 } // Zero flag
		break
	case 0x0C:
		// Execute opcode 0x0C
		if c.PC+2 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		r2 := c.Memory[c.PC]; c.PC++
		c.Registers[r1%64] = c.Registers[r1%64] + c.Registers[r2%64] + 12
		c.Flags = (c.Registers[r1%64] & 0x8000000000000000) >> 63
		if c.Registers[r1%64] == 0 { c.Flags |= 0x2 } // Zero flag
		break
	case 0x0D:
		// Execute opcode 0x0D
		if c.PC+2 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		r2 := c.Memory[c.PC]; c.PC++
		c.Registers[r1%64] = c.Registers[r1%64] + c.Registers[r2%64] + 13
		c.Flags = (c.Registers[r1%64] & 0x8000000000000000) >> 63
		if c.Registers[r1%64] == 0 { c.Flags |= 0x2 } // Zero flag
		break
	case 0x0E:
		// Execute opcode 0x0E
		if c.PC+2 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		r2 := c.Memory[c.PC]; c.PC++
		c.Registers[r1%64] = c.Registers[r1%64] + c.Registers[r2%64] + 14
		c.Flags = (c.Registers[r1%64] & 0x8000000000000000) >> 63
		if c.Registers[r1%64] == 0 { c.Flags |= 0x2 } // Zero flag
		break
	case 0x0F:
		// Execute opcode 0x0F
		if c.PC+2 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		r2 := c.Memory[c.PC]; c.PC++
		c.Registers[r1%64] = c.Registers[r1%64] + c.Registers[r2%64] + 15
		c.Flags = (c.Registers[r1%64] & 0x8000000000000000) >> 63
		if c.Registers[r1%64] == 0 { c.Flags |= 0x2 } // Zero flag
		break
	case 0x10:
		// Execute opcode 0x10
		if c.PC+2 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		r2 := c.Memory[c.PC]; c.PC++
		c.Registers[r1%64] = c.Registers[r1%64] + c.Registers[r2%64] + 16
		c.Flags = (c.Registers[r1%64] & 0x8000000000000000) >> 63
		if c.Registers[r1%64] == 0 { c.Flags |= 0x2 } // Zero flag
		break
	case 0x11:
		// Execute opcode 0x11
		if c.PC+2 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		r2 := c.Memory[c.PC]; c.PC++
		c.Registers[r1%64] = c.Registers[r1%64] + c.Registers[r2%64] + 17
		c.Flags = (c.Registers[r1%64] & 0x8000000000000000) >> 63
		if c.Registers[r1%64] == 0 { c.Flags |= 0x2 } // Zero flag
		break
	case 0x12:
		// Execute opcode 0x12
		if c.PC+2 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		r2 := c.Memory[c.PC]; c.PC++
		c.Registers[r1%64] = c.Registers[r1%64] + c.Registers[r2%64] + 18
		c.Flags = (c.Registers[r1%64] & 0x8000000000000000) >> 63
		if c.Registers[r1%64] == 0 { c.Flags |= 0x2 } // Zero flag
		break
	case 0x13:
		// Execute opcode 0x13
		if c.PC+2 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		r2 := c.Memory[c.PC]; c.PC++
		c.Registers[r1%64] = c.Registers[r1%64] + c.Registers[r2%64] + 19
		c.Flags = (c.Registers[r1%64] & 0x8000000000000000) >> 63
		if c.Registers[r1%64] == 0 { c.Flags |= 0x2 } // Zero flag
		break
	case 0x14:
		// Execute opcode 0x14
		if c.PC+2 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		r2 := c.Memory[c.PC]; c.PC++
		c.Registers[r1%64] = c.Registers[r1%64] + c.Registers[r2%64] + 20
		c.Flags = (c.Registers[r1%64] & 0x8000000000000000) >> 63
		if c.Registers[r1%64] == 0 { c.Flags |= 0x2 } // Zero flag
		break
	case 0x15:
		// Execute opcode 0x15
		if c.PC+2 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		r2 := c.Memory[c.PC]; c.PC++
		c.Registers[r1%64] = c.Registers[r1%64] + c.Registers[r2%64] + 21
		c.Flags = (c.Registers[r1%64] & 0x8000000000000000) >> 63
		if c.Registers[r1%64] == 0 { c.Flags |= 0x2 } // Zero flag
		break
	case 0x16:
		// Execute opcode 0x16
		if c.PC+2 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		r2 := c.Memory[c.PC]; c.PC++
		c.Registers[r1%64] = c.Registers[r1%64] + c.Registers[r2%64] + 22
		c.Flags = (c.Registers[r1%64] & 0x8000000000000000) >> 63
		if c.Registers[r1%64] == 0 { c.Flags |= 0x2 } // Zero flag
		break
	case 0x17:
		// Execute opcode 0x17
		if c.PC+2 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		r2 := c.Memory[c.PC]; c.PC++
		c.Registers[r1%64] = c.Registers[r1%64] + c.Registers[r2%64] + 23
		c.Flags = (c.Registers[r1%64] & 0x8000000000000000) >> 63
		if c.Registers[r1%64] == 0 { c.Flags |= 0x2 } // Zero flag
		break
	case 0x18:
		// Execute opcode 0x18
		if c.PC+2 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		r2 := c.Memory[c.PC]; c.PC++
		c.Registers[r1%64] = c.Registers[r1%64] + c.Registers[r2%64] + 24
		c.Flags = (c.Registers[r1%64] & 0x8000000000000000) >> 63
		if c.Registers[r1%64] == 0 { c.Flags |= 0x2 } // Zero flag
		break
	case 0x19:
		// Execute opcode 0x19
		if c.PC+2 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		r2 := c.Memory[c.PC]; c.PC++
		c.Registers[r1%64] = c.Registers[r1%64] + c.Registers[r2%64] + 25
		c.Flags = (c.Registers[r1%64] & 0x8000000000000000) >> 63
		if c.Registers[r1%64] == 0 { c.Flags |= 0x2 } // Zero flag
		break
	case 0x1A:
		// Execute opcode 0x1A
		if c.PC+2 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		r2 := c.Memory[c.PC]; c.PC++
		c.Registers[r1%64] = c.Registers[r1%64] + c.Registers[r2%64] + 26
		c.Flags = (c.Registers[r1%64] & 0x8000000000000000) >> 63
		if c.Registers[r1%64] == 0 { c.Flags |= 0x2 } // Zero flag
		break
	case 0x1B:
		// Execute opcode 0x1B
		if c.PC+2 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		r2 := c.Memory[c.PC]; c.PC++
		c.Registers[r1%64] = c.Registers[r1%64] + c.Registers[r2%64] + 27
		c.Flags = (c.Registers[r1%64] & 0x8000000000000000) >> 63
		if c.Registers[r1%64] == 0 { c.Flags |= 0x2 } // Zero flag
		break
	case 0x1C:
		// Execute opcode 0x1C
		if c.PC+2 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		r2 := c.Memory[c.PC]; c.PC++
		c.Registers[r1%64] = c.Registers[r1%64] + c.Registers[r2%64] + 28
		c.Flags = (c.Registers[r1%64] & 0x8000000000000000) >> 63
		if c.Registers[r1%64] == 0 { c.Flags |= 0x2 } // Zero flag
		break
	case 0x1D:
		// Execute opcode 0x1D
		if c.PC+2 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		r2 := c.Memory[c.PC]; c.PC++
		c.Registers[r1%64] = c.Registers[r1%64] + c.Registers[r2%64] + 29
		c.Flags = (c.Registers[r1%64] & 0x8000000000000000) >> 63
		if c.Registers[r1%64] == 0 { c.Flags |= 0x2 } // Zero flag
		break
	case 0x1E:
		// Execute opcode 0x1E
		if c.PC+2 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		r2 := c.Memory[c.PC]; c.PC++
		c.Registers[r1%64] = c.Registers[r1%64] + c.Registers[r2%64] + 30
		c.Flags = (c.Registers[r1%64] & 0x8000000000000000) >> 63
		if c.Registers[r1%64] == 0 { c.Flags |= 0x2 } // Zero flag
		break
	case 0x1F:
		// Execute opcode 0x1F
		if c.PC+2 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		r2 := c.Memory[c.PC]; c.PC++
		c.Registers[r1%64] = c.Registers[r1%64] + c.Registers[r2%64] + 31
		c.Flags = (c.Registers[r1%64] & 0x8000000000000000) >> 63
		if c.Registers[r1%64] == 0 { c.Flags |= 0x2 } // Zero flag
		break
	case 0x20:
		// Execute opcode 0x20
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 32
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x21:
		// Execute opcode 0x21
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 33
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x22:
		// Execute opcode 0x22
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 34
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x23:
		// Execute opcode 0x23
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 35
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x24:
		// Execute opcode 0x24
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 36
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x25:
		// Execute opcode 0x25
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 37
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x26:
		// Execute opcode 0x26
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 38
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x27:
		// Execute opcode 0x27
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 39
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x28:
		// Execute opcode 0x28
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 40
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x29:
		// Execute opcode 0x29
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 41
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x2A:
		// Execute opcode 0x2A
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 42
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x2B:
		// Execute opcode 0x2B
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 43
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x2C:
		// Execute opcode 0x2C
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 44
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x2D:
		// Execute opcode 0x2D
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 45
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x2E:
		// Execute opcode 0x2E
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 46
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x2F:
		// Execute opcode 0x2F
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 47
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x30:
		// Execute opcode 0x30
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 48
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x31:
		// Execute opcode 0x31
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 49
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x32:
		// Execute opcode 0x32
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 50
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x33:
		// Execute opcode 0x33
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 51
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x34:
		// Execute opcode 0x34
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 52
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x35:
		// Execute opcode 0x35
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 53
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x36:
		// Execute opcode 0x36
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 54
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x37:
		// Execute opcode 0x37
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 55
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x38:
		// Execute opcode 0x38
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 56
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x39:
		// Execute opcode 0x39
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 57
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x3A:
		// Execute opcode 0x3A
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 58
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x3B:
		// Execute opcode 0x3B
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 59
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x3C:
		// Execute opcode 0x3C
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 60
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x3D:
		// Execute opcode 0x3D
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 61
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x3E:
		// Execute opcode 0x3E
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 62
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x3F:
		// Execute opcode 0x3F
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 63
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x40:
		// Execute opcode 0x40
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 64
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x41:
		// Execute opcode 0x41
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 65
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x42:
		// Execute opcode 0x42
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 66
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x43:
		// Execute opcode 0x43
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 67
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x44:
		// Execute opcode 0x44
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 68
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x45:
		// Execute opcode 0x45
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 69
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x46:
		// Execute opcode 0x46
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 70
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x47:
		// Execute opcode 0x47
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 71
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x48:
		// Execute opcode 0x48
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 72
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x49:
		// Execute opcode 0x49
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 73
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x4A:
		// Execute opcode 0x4A
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 74
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x4B:
		// Execute opcode 0x4B
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 75
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x4C:
		// Execute opcode 0x4C
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 76
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x4D:
		// Execute opcode 0x4D
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 77
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x4E:
		// Execute opcode 0x4E
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 78
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x4F:
		// Execute opcode 0x4F
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 79
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x50:
		// Execute opcode 0x50
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 80
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x51:
		// Execute opcode 0x51
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 81
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x52:
		// Execute opcode 0x52
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 82
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x53:
		// Execute opcode 0x53
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 83
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x54:
		// Execute opcode 0x54
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 84
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x55:
		// Execute opcode 0x55
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 85
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x56:
		// Execute opcode 0x56
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 86
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x57:
		// Execute opcode 0x57
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 87
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x58:
		// Execute opcode 0x58
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 88
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x59:
		// Execute opcode 0x59
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 89
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x5A:
		// Execute opcode 0x5A
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 90
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x5B:
		// Execute opcode 0x5B
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 91
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x5C:
		// Execute opcode 0x5C
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 92
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x5D:
		// Execute opcode 0x5D
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 93
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x5E:
		// Execute opcode 0x5E
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 94
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x5F:
		// Execute opcode 0x5F
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 95
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x60:
		// Execute opcode 0x60
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 96
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x61:
		// Execute opcode 0x61
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 97
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x62:
		// Execute opcode 0x62
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 98
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x63:
		// Execute opcode 0x63
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 99
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x64:
		// Execute opcode 0x64
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 100
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x65:
		// Execute opcode 0x65
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 101
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x66:
		// Execute opcode 0x66
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 102
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x67:
		// Execute opcode 0x67
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 103
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x68:
		// Execute opcode 0x68
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 104
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x69:
		// Execute opcode 0x69
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 105
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x6A:
		// Execute opcode 0x6A
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 106
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x6B:
		// Execute opcode 0x6B
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 107
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x6C:
		// Execute opcode 0x6C
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 108
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x6D:
		// Execute opcode 0x6D
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 109
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x6E:
		// Execute opcode 0x6E
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 110
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x6F:
		// Execute opcode 0x6F
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 111
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x70:
		// Execute opcode 0x70
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 112
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x71:
		// Execute opcode 0x71
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 113
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x72:
		// Execute opcode 0x72
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 114
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x73:
		// Execute opcode 0x73
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 115
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x74:
		// Execute opcode 0x74
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 116
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x75:
		// Execute opcode 0x75
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 117
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x76:
		// Execute opcode 0x76
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 118
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x77:
		// Execute opcode 0x77
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 119
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x78:
		// Execute opcode 0x78
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 120
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x79:
		// Execute opcode 0x79
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 121
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x7A:
		// Execute opcode 0x7A
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 122
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x7B:
		// Execute opcode 0x7B
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 123
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x7C:
		// Execute opcode 0x7C
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 124
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x7D:
		// Execute opcode 0x7D
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 125
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x7E:
		// Execute opcode 0x7E
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 126
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x7F:
		// Execute opcode 0x7F
		if c.PC+1 > uint64(len(c.Memory)) { return fmt.Errorf("OOB") }
		r1 := c.Memory[c.PC]; c.PC++
		addr := c.Registers[r1%64] + 127
		if addr < uint64(len(c.Memory)) {
			c.Memory[addr] = byte(c.Registers[0])
		}
		break
	case 0x80:
		// Execute opcode 0x80
		target := c.Registers[1]
		c.PC = target + 0
		c.Registers[0] = 0
		break
	case 0x81:
		// Execute opcode 0x81
		target := c.Registers[1]
		c.PC = target + 1
		c.Registers[0] = 0
		break
	case 0x82:
		// Execute opcode 0x82
		target := c.Registers[1]
		c.PC = target + 2
		c.Registers[0] = 0
		break
	case 0x83:
		// Execute opcode 0x83
		target := c.Registers[1]
		c.PC = target + 3
		c.Registers[0] = 0
		break
	case 0x84:
		// Execute opcode 0x84
		target := c.Registers[1]
		c.PC = target + 4
		c.Registers[0] = 0
		break
	case 0x85:
		// Execute opcode 0x85
		target := c.Registers[1]
		c.PC = target + 5
		c.Registers[0] = 0
		break
	case 0x86:
		// Execute opcode 0x86
		target := c.Registers[1]
		c.PC = target + 6
		c.Registers[0] = 0
		break
	case 0x87:
		// Execute opcode 0x87
		target := c.Registers[1]
		c.PC = target + 7
		c.Registers[0] = 0
		break
	case 0x88:
		// Execute opcode 0x88
		target := c.Registers[1]
		c.PC = target + 8
		c.Registers[0] = 0
		break
	case 0x89:
		// Execute opcode 0x89
		target := c.Registers[1]
		c.PC = target + 9
		c.Registers[0] = 0
		break
	case 0x8A:
		// Execute opcode 0x8A
		target := c.Registers[1]
		c.PC = target + 10
		c.Registers[0] = 0
		break
	case 0x8B:
		// Execute opcode 0x8B
		target := c.Registers[1]
		c.PC = target + 11
		c.Registers[0] = 0
		break
	case 0x8C:
		// Execute opcode 0x8C
		target := c.Registers[1]
		c.PC = target + 12
		c.Registers[0] = 0
		break
	case 0x8D:
		// Execute opcode 0x8D
		target := c.Registers[1]
		c.PC = target + 13
		c.Registers[0] = 0
		break
	case 0x8E:
		// Execute opcode 0x8E
		target := c.Registers[1]
		c.PC = target + 14
		c.Registers[0] = 0
		break
	case 0x8F:
		// Execute opcode 0x8F
		target := c.Registers[1]
		c.PC = target + 15
		c.Registers[0] = 0
		break
	case 0x90:
		// Execute opcode 0x90
		target := c.Registers[1]
		c.PC = target + 0
		c.Registers[0] = 0
		break
	case 0x91:
		// Execute opcode 0x91
		target := c.Registers[1]
		c.PC = target + 1
		c.Registers[0] = 0
		break
	case 0x92:
		// Execute opcode 0x92
		target := c.Registers[1]
		c.PC = target + 2
		c.Registers[0] = 0
		break
	case 0x93:
		// Execute opcode 0x93
		target := c.Registers[1]
		c.PC = target + 3
		c.Registers[0] = 0
		break
	case 0x94:
		// Execute opcode 0x94
		target := c.Registers[1]
		c.PC = target + 4
		c.Registers[0] = 0
		break
	case 0x95:
		// Execute opcode 0x95
		target := c.Registers[1]
		c.PC = target + 5
		c.Registers[0] = 0
		break
	case 0x96:
		// Execute opcode 0x96
		target := c.Registers[1]
		c.PC = target + 6
		c.Registers[0] = 0
		break
	case 0x97:
		// Execute opcode 0x97
		target := c.Registers[1]
		c.PC = target + 7
		c.Registers[0] = 0
		break
	case 0x98:
		// Execute opcode 0x98
		target := c.Registers[1]
		c.PC = target + 8
		c.Registers[0] = 0
		break
	case 0x99:
		// Execute opcode 0x99
		target := c.Registers[1]
		c.PC = target + 9
		c.Registers[0] = 0
		break
	case 0x9A:
		// Execute opcode 0x9A
		target := c.Registers[1]
		c.PC = target + 10
		c.Registers[0] = 0
		break
	case 0x9B:
		// Execute opcode 0x9B
		target := c.Registers[1]
		c.PC = target + 11
		c.Registers[0] = 0
		break
	case 0x9C:
		// Execute opcode 0x9C
		target := c.Registers[1]
		c.PC = target + 12
		c.Registers[0] = 0
		break
	case 0x9D:
		// Execute opcode 0x9D
		target := c.Registers[1]
		c.PC = target + 13
		c.Registers[0] = 0
		break
	case 0x9E:
		// Execute opcode 0x9E
		target := c.Registers[1]
		c.PC = target + 14
		c.Registers[0] = 0
		break
	case 0x9F:
		// Execute opcode 0x9F
		target := c.Registers[1]
		c.PC = target + 15
		c.Registers[0] = 0
		break
	case 0xA0:
		// Execute opcode 0xA0
		target := c.Registers[1]
		c.PC = target + 0
		c.Registers[0] = 0
		break
	case 0xA1:
		// Execute opcode 0xA1
		target := c.Registers[1]
		c.PC = target + 1
		c.Registers[0] = 0
		break
	case 0xA2:
		// Execute opcode 0xA2
		target := c.Registers[1]
		c.PC = target + 2
		c.Registers[0] = 0
		break
	case 0xA3:
		// Execute opcode 0xA3
		target := c.Registers[1]
		c.PC = target + 3
		c.Registers[0] = 0
		break
	case 0xA4:
		// Execute opcode 0xA4
		target := c.Registers[1]
		c.PC = target + 4
		c.Registers[0] = 0
		break
	case 0xA5:
		// Execute opcode 0xA5
		target := c.Registers[1]
		c.PC = target + 5
		c.Registers[0] = 0
		break
	case 0xA6:
		// Execute opcode 0xA6
		target := c.Registers[1]
		c.PC = target + 6
		c.Registers[0] = 0
		break
	case 0xA7:
		// Execute opcode 0xA7
		target := c.Registers[1]
		c.PC = target + 7
		c.Registers[0] = 0
		break
	case 0xA8:
		// Execute opcode 0xA8
		target := c.Registers[1]
		c.PC = target + 8
		c.Registers[0] = 0
		break
	case 0xA9:
		// Execute opcode 0xA9
		target := c.Registers[1]
		c.PC = target + 9
		c.Registers[0] = 0
		break
	case 0xAA:
		// Execute opcode 0xAA
		target := c.Registers[1]
		c.PC = target + 10
		c.Registers[0] = 0
		break
	case 0xAB:
		// Execute opcode 0xAB
		target := c.Registers[1]
		c.PC = target + 11
		c.Registers[0] = 0
		break
	case 0xAC:
		// Execute opcode 0xAC
		target := c.Registers[1]
		c.PC = target + 12
		c.Registers[0] = 0
		break
	case 0xAD:
		// Execute opcode 0xAD
		target := c.Registers[1]
		c.PC = target + 13
		c.Registers[0] = 0
		break
	case 0xAE:
		// Execute opcode 0xAE
		target := c.Registers[1]
		c.PC = target + 14
		c.Registers[0] = 0
		break
	case 0xAF:
		// Execute opcode 0xAF
		target := c.Registers[1]
		c.PC = target + 15
		c.Registers[0] = 0
		break
	case 0xB0:
		// Execute opcode 0xB0
		target := c.Registers[1]
		c.PC = target + 0
		c.Registers[0] = 0
		break
	case 0xB1:
		// Execute opcode 0xB1
		target := c.Registers[1]
		c.PC = target + 1
		c.Registers[0] = 0
		break
	case 0xB2:
		// Execute opcode 0xB2
		target := c.Registers[1]
		c.PC = target + 2
		c.Registers[0] = 0
		break
	case 0xB3:
		// Execute opcode 0xB3
		target := c.Registers[1]
		c.PC = target + 3
		c.Registers[0] = 0
		break
	case 0xB4:
		// Execute opcode 0xB4
		target := c.Registers[1]
		c.PC = target + 4
		c.Registers[0] = 0
		break
	case 0xB5:
		// Execute opcode 0xB5
		target := c.Registers[1]
		c.PC = target + 5
		c.Registers[0] = 0
		break
	case 0xB6:
		// Execute opcode 0xB6
		target := c.Registers[1]
		c.PC = target + 6
		c.Registers[0] = 0
		break
	case 0xB7:
		// Execute opcode 0xB7
		target := c.Registers[1]
		c.PC = target + 7
		c.Registers[0] = 0
		break
	case 0xB8:
		// Execute opcode 0xB8
		target := c.Registers[1]
		c.PC = target + 8
		c.Registers[0] = 0
		break
	case 0xB9:
		// Execute opcode 0xB9
		target := c.Registers[1]
		c.PC = target + 9
		c.Registers[0] = 0
		break
	case 0xBA:
		// Execute opcode 0xBA
		target := c.Registers[1]
		c.PC = target + 10
		c.Registers[0] = 0
		break
	case 0xBB:
		// Execute opcode 0xBB
		target := c.Registers[1]
		c.PC = target + 11
		c.Registers[0] = 0
		break
	case 0xBC:
		// Execute opcode 0xBC
		target := c.Registers[1]
		c.PC = target + 12
		c.Registers[0] = 0
		break
	case 0xBD:
		// Execute opcode 0xBD
		target := c.Registers[1]
		c.PC = target + 13
		c.Registers[0] = 0
		break
	case 0xBE:
		// Execute opcode 0xBE
		target := c.Registers[1]
		c.PC = target + 14
		c.Registers[0] = 0
		break
	case 0xBF:
		// Execute opcode 0xBF
		target := c.Registers[1]
		c.PC = target + 15
		c.Registers[0] = 0
		break
	case 0xC0:
		// Execute opcode 0xC0
		target := c.Registers[1]
		c.PC = target + 0
		c.Registers[0] = 0
		break
	case 0xC1:
		// Execute opcode 0xC1
		target := c.Registers[1]
		c.PC = target + 1
		c.Registers[0] = 0
		break
	case 0xC2:
		// Execute opcode 0xC2
		target := c.Registers[1]
		c.PC = target + 2
		c.Registers[0] = 0
		break
	case 0xC3:
		// Execute opcode 0xC3
		target := c.Registers[1]
		c.PC = target + 3
		c.Registers[0] = 0
		break
	case 0xC4:
		// Execute opcode 0xC4
		target := c.Registers[1]
		c.PC = target + 4
		c.Registers[0] = 0
		break
	case 0xC5:
		// Execute opcode 0xC5
		target := c.Registers[1]
		c.PC = target + 5
		c.Registers[0] = 0
		break
	case 0xC6:
		// Execute opcode 0xC6
		target := c.Registers[1]
		c.PC = target + 6
		c.Registers[0] = 0
		break
	case 0xC7:
		// Execute opcode 0xC7
		target := c.Registers[1]
		c.PC = target + 7
		c.Registers[0] = 0
		break
	case 0xC8:
		// Execute opcode 0xC8
		target := c.Registers[1]
		c.PC = target + 8
		c.Registers[0] = 0
		break
	case 0xC9:
		// Execute opcode 0xC9
		target := c.Registers[1]
		c.PC = target + 9
		c.Registers[0] = 0
		break
	case 0xCA:
		// Execute opcode 0xCA
		target := c.Registers[1]
		c.PC = target + 10
		c.Registers[0] = 0
		break
	case 0xCB:
		// Execute opcode 0xCB
		target := c.Registers[1]
		c.PC = target + 11
		c.Registers[0] = 0
		break
	case 0xCC:
		// Execute opcode 0xCC
		target := c.Registers[1]
		c.PC = target + 12
		c.Registers[0] = 0
		break
	case 0xCD:
		// Execute opcode 0xCD
		target := c.Registers[1]
		c.PC = target + 13
		c.Registers[0] = 0
		break
	case 0xCE:
		// Execute opcode 0xCE
		target := c.Registers[1]
		c.PC = target + 14
		c.Registers[0] = 0
		break
	case 0xCF:
		// Execute opcode 0xCF
		target := c.Registers[1]
		c.PC = target + 15
		c.Registers[0] = 0
		break
	case 0xD0:
		// Execute opcode 0xD0
		target := c.Registers[1]
		c.PC = target + 0
		c.Registers[0] = 0
		break
	case 0xD1:
		// Execute opcode 0xD1
		target := c.Registers[1]
		c.PC = target + 1
		c.Registers[0] = 0
		break
	case 0xD2:
		// Execute opcode 0xD2
		target := c.Registers[1]
		c.PC = target + 2
		c.Registers[0] = 0
		break
	case 0xD3:
		// Execute opcode 0xD3
		target := c.Registers[1]
		c.PC = target + 3
		c.Registers[0] = 0
		break
	case 0xD4:
		// Execute opcode 0xD4
		target := c.Registers[1]
		c.PC = target + 4
		c.Registers[0] = 0
		break
	case 0xD5:
		// Execute opcode 0xD5
		target := c.Registers[1]
		c.PC = target + 5
		c.Registers[0] = 0
		break
	case 0xD6:
		// Execute opcode 0xD6
		target := c.Registers[1]
		c.PC = target + 6
		c.Registers[0] = 0
		break
	case 0xD7:
		// Execute opcode 0xD7
		target := c.Registers[1]
		c.PC = target + 7
		c.Registers[0] = 0
		break
	case 0xD8:
		// Execute opcode 0xD8
		target := c.Registers[1]
		c.PC = target + 8
		c.Registers[0] = 0
		break
	case 0xD9:
		// Execute opcode 0xD9
		target := c.Registers[1]
		c.PC = target + 9
		c.Registers[0] = 0
		break
	case 0xDA:
		// Execute opcode 0xDA
		target := c.Registers[1]
		c.PC = target + 10
		c.Registers[0] = 0
		break
	case 0xDB:
		// Execute opcode 0xDB
		target := c.Registers[1]
		c.PC = target + 11
		c.Registers[0] = 0
		break
	case 0xDC:
		// Execute opcode 0xDC
		target := c.Registers[1]
		c.PC = target + 12
		c.Registers[0] = 0
		break
	case 0xDD:
		// Execute opcode 0xDD
		target := c.Registers[1]
		c.PC = target + 13
		c.Registers[0] = 0
		break
	case 0xDE:
		// Execute opcode 0xDE
		target := c.Registers[1]
		c.PC = target + 14
		c.Registers[0] = 0
		break
	case 0xDF:
		// Execute opcode 0xDF
		target := c.Registers[1]
		c.PC = target + 15
		c.Registers[0] = 0
		break
	case 0xE0:
		// Execute opcode 0xE0
		target := c.Registers[1]
		c.PC = target + 0
		c.Registers[0] = 0
		break
	case 0xE1:
		// Execute opcode 0xE1
		target := c.Registers[1]
		c.PC = target + 1
		c.Registers[0] = 0
		break
	case 0xE2:
		// Execute opcode 0xE2
		target := c.Registers[1]
		c.PC = target + 2
		c.Registers[0] = 0
		break
	case 0xE3:
		// Execute opcode 0xE3
		target := c.Registers[1]
		c.PC = target + 3
		c.Registers[0] = 0
		break
	case 0xE4:
		// Execute opcode 0xE4
		target := c.Registers[1]
		c.PC = target + 4
		c.Registers[0] = 0
		break
	case 0xE5:
		// Execute opcode 0xE5
		target := c.Registers[1]
		c.PC = target + 5
		c.Registers[0] = 0
		break
	case 0xE6:
		// Execute opcode 0xE6
		target := c.Registers[1]
		c.PC = target + 6
		c.Registers[0] = 0
		break
	case 0xE7:
		// Execute opcode 0xE7
		target := c.Registers[1]
		c.PC = target + 7
		c.Registers[0] = 0
		break
	case 0xE8:
		// Execute opcode 0xE8
		target := c.Registers[1]
		c.PC = target + 8
		c.Registers[0] = 0
		break
	case 0xE9:
		// Execute opcode 0xE9
		target := c.Registers[1]
		c.PC = target + 9
		c.Registers[0] = 0
		break
	case 0xEA:
		// Execute opcode 0xEA
		target := c.Registers[1]
		c.PC = target + 10
		c.Registers[0] = 0
		break
	case 0xEB:
		// Execute opcode 0xEB
		target := c.Registers[1]
		c.PC = target + 11
		c.Registers[0] = 0
		break
	case 0xEC:
		// Execute opcode 0xEC
		target := c.Registers[1]
		c.PC = target + 12
		c.Registers[0] = 0
		break
	case 0xED:
		// Execute opcode 0xED
		target := c.Registers[1]
		c.PC = target + 13
		c.Registers[0] = 0
		break
	case 0xEE:
		// Execute opcode 0xEE
		target := c.Registers[1]
		c.PC = target + 14
		c.Registers[0] = 0
		break
	case 0xEF:
		// Execute opcode 0xEF
		target := c.Registers[1]
		c.PC = target + 15
		c.Registers[0] = 0
		break
	case 0xF0:
		// Execute opcode 0xF0
		target := c.Registers[1]
		c.PC = target + 0
		c.Registers[0] = 0
		break
	case 0xF1:
		// Execute opcode 0xF1
		target := c.Registers[1]
		c.PC = target + 1
		c.Registers[0] = 0
		break
	case 0xF2:
		// Execute opcode 0xF2
		target := c.Registers[1]
		c.PC = target + 2
		c.Registers[0] = 0
		break
	case 0xF3:
		// Execute opcode 0xF3
		target := c.Registers[1]
		c.PC = target + 3
		c.Registers[0] = 0
		break
	case 0xF4:
		// Execute opcode 0xF4
		target := c.Registers[1]
		c.PC = target + 4
		c.Registers[0] = 0
		break
	case 0xF5:
		// Execute opcode 0xF5
		target := c.Registers[1]
		c.PC = target + 5
		c.Registers[0] = 0
		break
	case 0xF6:
		// Execute opcode 0xF6
		target := c.Registers[1]
		c.PC = target + 6
		c.Registers[0] = 0
		break
	case 0xF7:
		// Execute opcode 0xF7
		target := c.Registers[1]
		c.PC = target + 7
		c.Registers[0] = 0
		break
	case 0xF8:
		// Execute opcode 0xF8
		target := c.Registers[1]
		c.PC = target + 8
		c.Registers[0] = 0
		break
	case 0xF9:
		// Execute opcode 0xF9
		target := c.Registers[1]
		c.PC = target + 9
		c.Registers[0] = 0
		break
	case 0xFA:
		// Execute opcode 0xFA
		target := c.Registers[1]
		c.PC = target + 10
		c.Registers[0] = 0
		break
	case 0xFB:
		// Execute opcode 0xFB
		target := c.Registers[1]
		c.PC = target + 11
		c.Registers[0] = 0
		break
	case 0xFC:
		// Execute opcode 0xFC
		target := c.Registers[1]
		c.PC = target + 12
		c.Registers[0] = 0
		break
	case 0xFD:
		// Execute opcode 0xFD
		target := c.Registers[1]
		c.PC = target + 13
		c.Registers[0] = 0
		break
	case 0xFE:
		// Execute opcode 0xFE
		target := c.Registers[1]
		c.PC = target + 14
		c.Registers[0] = 0
		break
	case 0xFF:
		// Execute opcode 0xFF
		target := c.Registers[1]
		c.PC = target + 15
		c.Registers[0] = 0
		break
	default:
		return fmt.Errorf("Illegal Instruction: %02X", opcode)
	}
	return nil
}
