package main

const (
    running int64 = iota
    finished
    wantInput
    haveOutput
)

type intCode struct {
    program      []int64
    program_copy []int64
    state        int64
    pc           int64
    nextPc       int64
    base         int64
    i, j, k      int64
}

func (ic *intCode) currentOp() int64 {
    return ic.program[ic.pc]
}

func (ic *intCode) reset() {
    ic.state = 0
    ic.pc = 0
    ic.nextPc = 0
    ic.base = 0
    ic.i = 0
    ic.j = 0
    ic.k = 0
    copy(ic.program, ic.program_copy)
}

func (ic *intCode) instructionLength() int64 {
    instructionLengths := map[int64]int64{
        1: 4, 2: 4, 3: 2, 4: 2, 5: 3, 6: 3, 7: 4, 8: 4, 9: 2, 99: 1}
    return instructionLengths[ic.currentOp()%100]
}

func (ic *intCode) getParams() {
    if ic.instructionLength() >= 2 {
        if ic.currentOp()/100%10 == 1 {
            ic.i = ic.pc + 1
        } else if ic.currentOp()/100%10 == 2 {
            ic.i = ic.program[ic.pc+1] + ic.base
        } else {
            ic.i = ic.program[ic.pc+1]
        }
    }
    if ic.instructionLength() >= 3 {
        if ic.currentOp()/1000%10 == 1 {
            ic.j = ic.pc + 2
        } else if ic.currentOp()/1000%10 == 2 {
            ic.j = ic.program[ic.pc+2] + ic.base
        } else {
            ic.j = ic.program[ic.pc+2]
        }
    }
    if ic.instructionLength() >= 4 {
        if ic.currentOp()/10000%10 == 1 {
            panic("intCode.getParams(): error")
        } else if ic.currentOp()/10000%10 == 2 {
            ic.k = ic.program[ic.pc+3] + ic.base
        } else {
            ic.k = ic.program[ic.pc+3]
        }
    }
}

func (ic *intCode) run() {
    for ic.state = running; ic.state == running; {
        ic.getParams()
        ic.nextPc = ic.pc + ic.instructionLength()
        switch ic.currentOp() % 100 {
        case 1:
            ic.program[ic.k] = ic.program[ic.i] + ic.program[ic.j]
        case 2:
            ic.program[ic.k] = ic.program[ic.i] * ic.program[ic.j]
        case 3:
            ic.state = wantInput
        case 4:
            ic.state = haveOutput
        case 5:
            if ic.program[ic.i] != 0 {
                ic.nextPc = ic.program[ic.j]
            }
        case 6:
            if ic.program[ic.i] == 0 {
                ic.nextPc = ic.program[ic.j]
            }
        case 7:
            if ic.program[ic.i] < ic.program[ic.j] {
                ic.program[ic.k] = 1
            } else {
                ic.program[ic.k] = 0
            }
        case 8:
            if ic.program[ic.i] == ic.program[ic.j] {
                ic.program[ic.k] = 1
            } else {
                ic.program[ic.k] = 0
            }
        case 9:
            ic.base += ic.program[ic.i]
        case 99:
            ic.state = finished
        default:
            panic("intCode.run(): error")
        }
        ic.pc = ic.nextPc
    }
}
