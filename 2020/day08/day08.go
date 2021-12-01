package day08

func part1(inputfile string) int {
	m := NewMachine(inputfile)
	m.Run()

	return m.Accumulator
}

func part2(inputfile string) int {
	m := NewMachine(inputfile)

	swapMe := indexOfNextJmpNop(m.Instructions, -1)

	m.RepeatHook = func() bool {
		swapMe = indexOfNextJmpNop(m.Instructions, swapMe)
		m.Reset()
		return true
	}

	m.HandlerHook = func(i instruction) handler {
		if m.Pos == swapMe {
			if i.op == "jmp" {
				return m.Handlers["nop"]
			}

			if i.op == "nop" {
				return m.Handlers["jmp"]
			}
		}

		return m.Handlers[i.op]
	}

	m.Run()

	return m.Accumulator
}

func indexOfNextJmpNop(ins []instruction, cur int) int {
	for i := cur + 1; i < len(ins); i++ {
		in := ins[i]
		if in.op == "jmp" {
			return i
		}
		if in.op == "nop" && in.arg != 0 {
			return i
		}
	}

	return -1
}
