package day08

import (
	"aoc2020/util"
	"strconv"
	"strings"
)

type machine struct {
	Accumulator  int
	Pos          int
	Executed     map[int]struct{}
	Instructions []instruction
	Handlers     map[string]handler
	RepeatHook   func() bool
	HandlerHook  func(instruction) handler
}

type instruction struct {
	op  string
	arg int
}

type handler func(instruction)

func NewMachine(instructionFile string) *machine {
	m := &machine{}
	m.Executed = map[int]struct{}{}

	handlers := map[string]handler{}

	handlers["acc"] = m.acc
	handlers["jmp"] = m.jmp
	handlers["nop"] = m.nop
	m.Handlers = handlers

	// Default Repeat Func
	m.RepeatHook = func() bool { return false }

	// Default Handler Decision
	m.HandlerHook = func(i instruction) handler { return m.Handlers[i.op] }

	m.parseFile(instructionFile)

	return m
}

func (m *machine) Run() {
	for m.Next() {
	}
}

func (m *machine) Next() bool {
	if _, ok := m.Executed[m.Pos]; ok {
		if !m.RepeatHook() {
			return false
		}
	}

	m.Executed[m.Pos] = struct{}{}

	if m.Pos >= len(m.Instructions) {
		return false
	}

	i := m.Instructions[m.Pos]

	handler := m.HandlerHook(i)
	handler(i)

	return true
}

func (m *machine) Reset() {
	m.Executed = map[int]struct{}{}
	m.Pos = 0
	m.Accumulator = 0
}

func (m *machine) acc(i instruction) {
	m.Accumulator += i.arg
	m.Pos++
}

func (m *machine) jmp(i instruction) {
	m.Pos += i.arg
}

func (m *machine) nop(i instruction) {
	m.Pos++
}

func (m *machine) parseFile(inputfile string) {
	data, err := util.ReadFileToStringSlice(inputfile)
	util.Check(err)

	instructions := []instruction{}

	for _, line := range data {
		lineS := strings.Split(line, " ")
		i := instruction{}
		i.op = strings.TrimSpace(lineS[0])
		i.arg, _ = strconv.Atoi(lineS[1])
		instructions = append(instructions, i)
	}

	m.Instructions = instructions
}
