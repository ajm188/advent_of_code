package alu

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type Instruction interface {
	Execute(*ALU)
}

var (
	ErrInvalidInstruction = errors.New("invalid instruction")

	instructionRegexp = regexp.MustCompile(`^(?:(?P<input>inp) (?P<register>[wxyz]))|(?:(?P<name>add|mul|div|mod|eql) (?P<register1>[wxyz]) (?:(?P<register2>[wxyz])|(?P<val>-?\d+)))$`)
)

func ParseInstruction(s string) (Instruction, error) {
	m := instructionRegexp.FindStringSubmatch(s)
	if m == nil {
		return nil, fmt.Errorf("%w: does not match %s", ErrInvalidInstruction, instructionRegexp)
	}

	if m[instructionRegexp.SubexpIndex("input")] != "" {
		return &inputInstr{register: m[instructionRegexp.SubexpIndex("register")]}, nil
	}

	a, b, val := m[instructionRegexp.SubexpIndex("register1")], m[instructionRegexp.SubexpIndex("register2")], m[instructionRegexp.SubexpIndex("val")]
	var n *int64
	if val == "" && b == "" {
		panic("impossible; second operand must be register or number")
	}

	if val != "" {
		_n, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("%w: operand failed to parse as int %s", ErrInvalidInstruction, err)
		}

		n = &_n
	}

	switch m[instructionRegexp.SubexpIndex("name")] {
	case "add":
		return &addInstr{a: a, b: b, n: n}, nil
	case "mul":
		return &mulInstr{a: a, b: b, n: n}, nil
	case "div":
		return &divInstr{a: a, b: b, n: n}, nil
	case "mod":
		return &modInstr{a: a, b: b, n: n}, nil
	case "eql":
		return &eqlInstr{a: a, b: b, n: n}, nil
	default:
		panic("impossible; unknown instruction name " + m[instructionRegexp.SubexpIndex("name")])
	}
}

type inputInstr struct {
	register string
}

func (instr *inputInstr) Execute(alu *ALU) {
	if len(alu.Input) == 0 {
		panic("")
	}

	v := reflect.ValueOf(alu).Elem()

	field := v.FieldByName(strings.Title(instr.register))
	if field.Kind() == reflect.Invalid {
		panic("")
	}

	field.Set(reflect.ValueOf(alu.Input[0]))
	alu.Input = alu.Input[1:]
}

func execCommon(alu *ALU, aName, bName string, n *int64, f func(a, b int64) int64) {
	v := reflect.ValueOf(alu).Elem()
	a := v.FieldByName(strings.Title(aName))
	if a.Kind() == reflect.Invalid {
		panic("")
	}

	var b int64
	switch n {
	case nil:
		_b := v.FieldByName(strings.Title(bName))
		if _b.Kind() == reflect.Invalid {
			panic("")
		}

		b = _b.Int()
	default:
		b = *n
	}

	a.SetInt(f(a.Int(), b))
}

type addInstr struct {
	a, b string
	n    *int64
}

func (instr *addInstr) Execute(alu *ALU) {
	execCommon(alu, instr.a, instr.b, instr.n, func(a, b int64) int64 { return a + b })
}

type mulInstr struct {
	a, b string
	n    *int64
}

func (instr *mulInstr) Execute(alu *ALU) {
	execCommon(alu, instr.a, instr.b, instr.n, func(a, b int64) int64 { return a * b })
}

type divInstr struct {
	a, b string
	n    *int64
}

func (instr *divInstr) Execute(alu *ALU) {
	execCommon(alu, instr.a, instr.b, instr.n, func(a, b int64) int64 { return a / b })
}

type modInstr struct {
	a, b string
	n    *int64
}

func (instr *modInstr) Execute(alu *ALU) {
	execCommon(alu, instr.a, instr.b, instr.n, func(a, b int64) int64 { return a % b })
}

type eqlInstr struct {
	a, b string
	n    *int64
}

func (instr *eqlInstr) Execute(alu *ALU) {
	execCommon(alu, instr.a, instr.b, instr.n, func(a, b int64) int64 {
		if a == b {
			return 1
		}

		return 0
	})
}
