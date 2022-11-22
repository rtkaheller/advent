package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
	//"sort"
	"strconv"
	"strings"
)

const start = 99999999999999

type Vars struct {
	I, W, X, Y, Z int
}

type InpPair struct {
	Pos, Value int
}

func (v *Vars) Set(r rune, val int) {
	switch r {
	case 'i':
		v.I = val
	case 'w':
		v.W = val
	case 'x':
		v.X = val
	case 'y':
		v.Y = val
	case 'z':
		v.Z = val
	}
}

func (v Vars) Get(r rune) int {
	switch r {
	case 'i':
		return v.I
	case 'w':
		return v.W
	case 'x':
		return v.X
	case 'y':
		return v.Y
	case 'z':
		return v.Z
	}
	return 0
}

type Node struct {
	Inst   *Instruction
	State  int
	Set    bool
	Index  int
	Input  map[rune]*Node
	User   []*Node
	Inps   map[int]*Node
	Values []*BackPropVal
}

type NodeList []*Node

func (l NodeList) Len() int {
	return len(l)
}

func (l NodeList) Less(i, j int) bool {
	return l[i].Index < l[j].Index
}
func (l NodeList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (n *Node) Unset() {
	n.Set = false
	for _, v := range n.User {
		v.Unset()
	}
}

func (n *Node) Propogate(state Globe) {
	if n.Inst == nil {
		n.State = 0
		return
	}
	fmt.Printf("Banana: %T\n", (*n.Inst))
	fmt.Println(n.Index)
	//fmt.Println(state)
	//fmt.Println((*n.Inst))
	//fmt.Println((*n.Inst).Inputs())
	//fmt.Println((*n.Inst).Output())
	fmt.Println(n.User)
	for k, v := range n.Input {
		if _, ok := state[k]; !ok {
			state[k] = v.State
		}
	}
	//fmt.Println(state)
	(*n.Inst).Do(&state)
	//fmt.Println(state)
	n.State = state[(*n.Inst).Output()]
	for _, v := range n.User {
		fmt.Println(v.Index)
		v.Propogate(map[rune]int{(*n.Inst).Output(): state[(*n.Inst).Output()]})
	}
}

func (n *Node) UnsafeCalc() {
	if n.Set {
		return
	}
	var state Vars
	for k, v := range n.Input {
		state.Set(k, v.State)
	}
	if n.Inst != nil {
		n.State = (*n.Inst).DoVars(state)
	} else {
		n.State = 0
	}
	n.Set = true
}

func (n *Node) InitCalcVars() {
	if n.Set {
		return
	}
	var state Vars
	for k, v := range n.Input {
		if !v.Set {
			v.InitCalc()
		}
		state.Set(k, v.State)
	}
	if n.Inst != nil {
		n.State = (*n.Inst).DoVars(state)
	} else {
		n.State = 0
	}
	n.Set = true
}

func (n *Node) InitCalc() {
	if n.Set {
		return
	}
	state := make(Globe)
	for k, v := range n.Input {
		if !v.Set {
			v.InitCalc()
		}
		state[k] = v.State
	}
	if n.Inst != nil {
		(*n.Inst).Do(&state)
		n.State = state[(*n.Inst).Output()]
	} else {
		n.State = 0
	}
	n.Set = true
}

type Instruction interface {
	Do(*Globe)
	DoVars(Vars) int
	Inputs() []rune
	Output() rune
}

type Inp struct {
	ResultVar rune
	Pos       int
}

func (i Inp) Do(state *Globe) {
	(*state)[i.ResultVar] = (*state)['i']
}
func (i Inp) DoVars(state Vars) int {
	return state.I
}
func (i Inp) Inputs() []rune {
	return []rune{}
}
func (i Inp) Output() rune {
	return i.ResultVar
}

type Add struct {
	ResultVar rune
	ValVar    rune
	Value     int
}

func (i Add) Do(state *Globe) {
	if i.ValVar != 0 {
		(*state)[i.ResultVar] += (*state)[i.ValVar]
	} else {
		(*state)[i.ResultVar] += i.Value
	}
}
func (i Add) DoVars(state Vars) int {
	if i.ValVar != 0 {
		return state.Get(i.ResultVar) + state.Get(i.ValVar)
	} else {
		return state.Get(i.ResultVar) + i.Value
	}
}
func (i Add) Inputs() []rune {
	if i.ValVar != 0 {
		return []rune{i.ResultVar, i.ValVar}
	} else {
		return []rune{i.ResultVar}
	}
}
func (i Add) Output() rune {
	return i.ResultVar
}

type Mul struct {
	ResultVar rune
	ValVar    rune
	Value     int
}

func (i Mul) DoVars(state Vars) int {
	if i.ValVar != 0 {
		return state.Get(i.ResultVar) * state.Get(i.ValVar)
	} else {
		return state.Get(i.ResultVar) * i.Value
	}
}
func (i Mul) Do(state *Globe) {
	if i.ValVar != 0 {
		(*state)[i.ResultVar] *= (*state)[i.ValVar]
	} else {
		(*state)[i.ResultVar] *= i.Value
	}
}
func (i Mul) Inputs() []rune {
	if i.ValVar != 0 {
		return []rune{i.ResultVar, i.ValVar}
	} else {
		return []rune{i.ResultVar}
	}
}
func (i Mul) Output() rune {
	return i.ResultVar
}

type Div struct {
	ResultVar rune
	ValVar    rune
	Value     int
}

func (i Div) DoVars(state Vars) int {
	if i.ValVar != 0 {
		return state.Get(i.ResultVar) / state.Get(i.ValVar)
	} else {
		return state.Get(i.ResultVar) / i.Value
	}
}
func (i Div) Do(state *Globe) {
	if i.ValVar != 0 {
		(*state)[i.ResultVar] /= (*state)[i.ValVar]
	} else {
		(*state)[i.ResultVar] /= i.Value
	}
}
func (i Div) Inputs() []rune {
	if i.ValVar != 0 {
		return []rune{i.ResultVar, i.ValVar}
	} else {
		return []rune{i.ResultVar}
	}
}
func (i Div) Output() rune {
	return i.ResultVar
}

type Mod struct {
	ResultVar rune
	ValVar    rune
	Value     int
}

func (i Mod) DoVars(state Vars) int {
	if i.ValVar != 0 {
		return state.Get(i.ResultVar) % state.Get(i.ValVar)
	} else {
		return state.Get(i.ResultVar) % i.Value
	}
}
func (i Mod) Do(state *Globe) {
	if i.ValVar != 0 {
		//fmt.Printf("%v %% %v = %v\n", (*state)[i.ResultVar], (*state)[i.ValVar], (*state)[i.ResultVar] % (*state)[i.ValVar])
		(*state)[i.ResultVar] %= (*state)[i.ValVar]
	} else {
		//fmt.Printf("%v %% %v = %v\n", (*state)[i.ResultVar], i.Value, (*state)[i.ResultVar] % i.Value)
		(*state)[i.ResultVar] %= i.Value
	}
}
func (i Mod) Inputs() []rune {
	if i.ValVar != 0 {
		return []rune{i.ResultVar, i.ValVar}
	} else {
		return []rune{i.ResultVar}
	}
}
func (i Mod) Output() rune {
	return i.ResultVar
}

type Eql struct {
	ResultVar rune
	ValVar    rune
	Value     int
}

func (i Eql) DoVars(state Vars) int {
	var comp int
	if i.ValVar != 0 {
		comp = state.Get(i.ValVar)
	} else {
		comp = i.Value
	}
	if state.Get(i.ResultVar) == comp {
		return 1
	} else {
		return 0
	}
}
func (i Eql) Inputs() []rune {
	if i.ValVar != 0 {
		return []rune{i.ResultVar, i.ValVar}
	} else {
		return []rune{i.ResultVar}
	}
}
func (i Eql) Output() rune {
	return i.ResultVar
}
func (i Eql) Do(state *Globe) {
	var comp int
	if i.ValVar != 0 {
		comp = (*state)[i.ValVar]
	} else {
		comp = i.Value
	}
	if (*state)[i.ResultVar] == comp {
		(*state)[i.ResultVar] = 1
	} else {
		(*state)[i.ResultVar] = 0
	}
}

func Digits(attempt int) []int {
	var d []int
	for attempt > 0 {
		d = append(d, attempt%10)
		attempt /= 10
	}
	return d
}

type Globe map[rune]int

func (g Globe) Repr() string {
	r := ""
	for _, k := range []rune{'i', 'x', 'y', 'z', 'w'} {
		r += fmt.Sprintf("%v=%v ", string(k), g[k])
	}
	return r
}

type MemoKey struct {
	Globe string
	Digs  int
}

var memo map[MemoKey]int

func Calc(root Node, inputs []*Node, digs int, inpMap map[*Node]NodeList) int {
	if digs == 0 {
		return 0
	}
	for i := 9; i > 0; i-- {
		if digs > 7 {
			fmt.Printf("%v%v%v\n", strings.Repeat("x", 14-digs), i, strings.Repeat(".", digs-1))
		}
		for _, node := range inpMap[inputs[14-digs]] {
			node.Set = false
		}
		inputs[14-digs].Set = true
		inputs[14-digs].State = i
		//for _, node := range inpMap[inputs[14-digs]] {
		//  node.UnsafeCalc()
		//}

		root.InitCalcVars()
		//root.InitCalc()
		if root.Input['z'].State == 0 {
			return i*int(math.Pow(10, float64(digs-1))) + (int(math.Pow(10, float64(digs-1))) - 1)
		}
		r := Calc(root, inputs, digs-1, inpMap)
		if root.Input['z'].State == 0 {
			return i*int(math.Pow(10, float64(digs-1))) + r
		}
	}
	return 0
}

func TryVal(globe Globe, instructs []Instruction, digs int) int {
	if val, ok := memo[MemoKey{Globe: globe.Repr(), Digs: digs}]; ok {
		return val
	}
	for i := 9; i > 0; i-- {
		if digs > 10 {
			fmt.Printf("%v%v%v\n", strings.Repeat("x", 14-digs), i, strings.Repeat(".", digs-1))
		}
		newGlobe := make(Globe)
		for k, v := range globe {
			newGlobe[k] = v
		}

		newGlobe['i'] = i
		instructs[0].Do(&newGlobe)
		for t, inst := range instructs[1:] {
			fmt.Println(t, i, newGlobe.Repr())
			found := false
			switch v := inst.(type) {
			case Inp:
				_ = v
				found = true
				TryVal(newGlobe, instructs[t+1:], digs-1)
				//if r != 0 {
				//  fmt.Println(digs, i, r, i*int(math.Pow(10, float64(digs-1))), i*int(math.Pow(10, float64(digs-1))) + r)
				//  ret := i*int(math.Pow(10, float64(digs-1))) + r
				//  memo[MemoKey{Globe: globe.Repr(), Digs: digs}] = ret
				//  return ret
				//}
				//break
			}
			if found {
				break
			}
			inst.Do(&newGlobe)
		}
		if newGlobe['z'] == 0 {
			memo[MemoKey{Globe: globe.Repr(), Digs: digs}] = i
			return i
		}
	}
	memo[MemoKey{Globe: globe.Repr(), Digs: digs}] = 0
	return 0
}

type InfPoint struct {
	N      *Node
	Values []int
}

func Inflections(n Node) map[int]*InfPoint {
	points := make(map[int]*InfPoint)
	if n.Inst == nil {
		for _, v := range n.Input {
			for node, point := range Inflections(*v) {
				points[node] = point
			}
		}
	} else {
		switch v := (*n.Inst).(type) {
		case Inp:
			points[n.Index] = &InfPoint{N: &n}
			for i := 1; i < 10; i++ {
				points[n.Index].Values = append(points[n.Index].Values, i)
			}
		case Eql:
			points[n.Index] = &InfPoint{N: &n}
			for i := 0; i < 2; i++ {
				points[n.Index].Values = append(points[n.Index].Values, i)
			}
		case Mod:
			if v.ValVar == 0 {
				points[n.Index] = &InfPoint{N: &n}
				for i := 0; i < v.Value; i++ {
					points[n.Index].Values = append(points[n.Index].Values, i)
				}
			}
		default:
			for _, v := range n.Input {
				for node, point := range Inflections(*v) {
					points[node] = point
				}
			}
		}
	}
	return points
}

func FindVal(b *BackPropVal) map[int]int {
	results := make(map[int]int)
	for _, parent := range b.ParentValue {
		for k, v := range FindVal(parent) {
			if val, ok := results[k]; !ok || val > results[k] {
				results[k] = v
			}
		}
	}
	if len(results) > 0 {
		return results
	}

	switch v := (*b.Me.Inst).(type) {
	case Inp:
		return map[int]int{v.Pos: b.Value}
	}
	return map[int]int{}
}

type BackPropVal struct {
	ParentValue []*BackPropVal
	Parent      []*Node
	Me          *Node
	Value       int
}

func BackProp(n *Node) []*BackPropVal {
	if len(n.Values) > 0 {
		return n.Values
	} else {
		fmt.Println(n.Index)
	}
	var points []*BackPropVal
	if n.Inst == nil {
		// Only root
		return BackProp(n.Input['z'])
	} else {
		switch v := (*n.Inst).(type) {
		case Inp:
			for i := 9; i > 0; i-- {
				points = append(points, &BackPropVal{Value: i, Me: n})
			}
			n.Values = points
			return points
		case Eql:
			for i := 0; i < 2; i++ {
				points = append(points, &BackPropVal{Value: i, Me: n})
			}
			n.Values = points
			return points
		case Mod:
			if v.ValVar == 0 {
				for i := 0; i < v.Value; i++ {
					points = append(points, &BackPropVal{Value: i, Me: n})
				}
				n.Values = points
				return points
			}
		case Mul:
			if v.ValVar == 0 && v.Value == 0 {
				n.Values = []*BackPropVal{&BackPropVal{Value: 0, Me: n}}
				return n.Values
			}
		}
	}

	var vars Vars
	if len(n.Input) < 1 {
		n.Values = []*BackPropVal{&BackPropVal{Value: 0, Me: n}}
		return n.Values
	}
	found := make(map[int]bool)
	for _, val := range BackProp(n.Input[(*n.Inst).Output()]) {
		vars.Set((*n.Inst).Output(), val.Value)
		if len((*n.Inst).Inputs()) == 2 {
			for _, val2 := range BackProp(n.Input[(*n.Inst).Inputs()[1]]) {
				vars.Set((*n.Inst).Inputs()[1], val2.Value)
				point := BackPropVal{Value: (*n.Inst).DoVars(vars), Me: n}
				point.ParentValue = []*BackPropVal{val, val2}
				point.Parent = []*Node{n.Input[(*n.Inst).Output()], n.Input[(*n.Inst).Inputs()[1]]}
				points = append(points, &point)
				found[point.Value] = true
			}
		} else {
			point := BackPropVal{Value: (*n.Inst).DoVars(vars), Me: n}
			point.ParentValue = []*BackPropVal{val}
			point.Parent = []*Node{n.Input[(*n.Inst).Output()]}
			points = append(points, &point)
		}
	}
	n.Values = points
	return points
}

type Phase struct {
	Input     []rune
	Instructs []Instruction
}

func (p Phase) Do(inputs map[int]int, pos int) map[int]int {
	results := make(map[int]int)
	for i := 9; i > 0; i-- {
		for input, model := range inputs {
			globe := Globe{'i': i, 'z': input}
			for _, inst := range p.Instructs {
				inst.Do(&globe)
			}
			if val, ok := results[globe['z']]; !ok || model+int(math.Pow(10, float64(13-pos)))*i < val {
				results[globe['z']] = model + int(math.Pow(10, float64(13-pos)))*i
			}
		}
	}
	return results
}

func Phases(instructions []Instruction) []Phase {
	var phases []Phase
	inits := map[rune]bool{instructions[0].Output(): true}
	used := make(map[rune]bool)

	newPhase := Phase{Instructs: []Instruction{instructions[0]}}
	for _, inst := range instructions[1:] {
		switch v := inst.(type) {
		case Inp:
			for k, _ := range used {
				if !inits[k] {
					newPhase.Input = append(newPhase.Input, k)
				}
			}
			phases = append(phases, newPhase)
			newPhase = Phase{}
			inits = map[rune]bool{inst.Output(): true}
			used = make(map[rune]bool)
		case Mul:
			if v.ValVar == 0 {
				if v.Value == 1 {
					// Mul x 1 is a noop
					continue
				} else if v.Value == 0 {
					if !used[v.ResultVar] {
						// Mul x 0 is reset var
						inits[v.ResultVar] = true
					}
				}
			}
		case Div:
			if v.ValVar == 0 && v.Value == 1 {
				// Div x 1 is a noop
				continue
			}
		case Add:
			if v.ValVar == 0 && v.Value == 0 {
				// Add x 0 is a noop
				continue
			}
		}
		for _, r := range inst.Inputs() {
			used[r] = true
		}
		newPhase.Instructs = append(newPhase.Instructs, inst)
	}
	for k, _ := range used {
		if !inits[k] {
			newPhase.Input = append(newPhase.Input, k)
		}
	}
	phases = append(phases, newPhase)
	return phases
}

func main() {
	memo = make(map[MemoKey]int)
	contents, err := ioutil.ReadFile("input.txt")
	//contents, err := ioutil.ReadFile("fake.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	var instructs []Instruction
	pos := 14
	for _, line := range bytes.Split(contents, []byte("\n")) {
		if len(line) == 0 {
			break
		}
		inst := bytes.Split(line, []byte(" "))
		switch string(inst[0]) {
		case "inp":
			var newInst = Inp{ResultVar: []rune(string(inst[1]))[0], Pos: pos}
			instructs = append(instructs, newInst)
			pos--
		case "add":
			var newInst = Add{ResultVar: []rune(string(inst[1]))[0]}
			if val, err := strconv.Atoi(string(inst[2])); err == nil {
				newInst.Value = val
			} else {
				newInst.ValVar = []rune(string(inst[2]))[0]
			}
			instructs = append(instructs, newInst)
		case "mul":
			var newInst = Mul{ResultVar: []rune(string(inst[1]))[0]}
			if val, err := strconv.Atoi(string(inst[2])); err == nil {
				newInst.Value = val
			} else {
				newInst.ValVar = []rune(string(inst[2]))[0]
			}
			instructs = append(instructs, newInst)
		case "div":
			var newInst = Div{ResultVar: []rune(string(inst[1]))[0]}
			if val, err := strconv.Atoi(string(inst[2])); err == nil {
				newInst.Value = val
			} else {
				newInst.ValVar = []rune(string(inst[2]))[0]
			}
			instructs = append(instructs, newInst)
		case "mod":
			var newInst = Mod{ResultVar: []rune(string(inst[1]))[0]}
			if val, err := strconv.Atoi(string(inst[2])); err == nil {
				newInst.Value = val
			} else {
				newInst.ValVar = []rune(string(inst[2]))[0]
			}
			instructs = append(instructs, newInst)
		case "eql":
			var newInst = Eql{ResultVar: []rune(string(inst[1]))[0]}
			if val, err := strconv.Atoi(string(inst[2])); err == nil {
				newInst.Value = val
			} else {
				newInst.ValVar = []rune(string(inst[2]))[0]
			}
			instructs = append(instructs, newInst)
		}
	}
	lastSetter := make(map[rune]*Node)
	c := 0
	var inputs []*Node
	for i, inst := range instructs {
		var hasInputs = true
		n := Node{Inst: &instructs[i], Input: make(map[rune]*Node), Index: c, Inps: make(map[int]*Node)}
		fmt.Println(c)
		c += 1
		switch v := inst.(type) {
		case Inp:
			inputs = append(inputs, &n)
			n.State = 9
			n.Set = true
			n.Inps = map[int]*Node{n.Index: &n}
		case Mul:
			// Multiply by zero resets the state
			if v.ValVar == 0 && v.Value == 0 {
				hasInputs = false
			}
		}
		if hasInputs {
			for _, in := range inst.Inputs() {
				if lastSetter[in] != nil {
					n.Input[in] = lastSetter[in]
					lastSetter[in].User = append(lastSetter[in].User, &n)
					for k, v := range lastSetter[in].Inps {
						n.Inps[k] = v
					}
				}
			}
		}
		lastSetter[inst.Output()] = &n
	}
	var root Node
	root.Input = map[rune]*Node{
		'z': lastSetter['z'],
	}
	//root.InitCalc()
	//for _, inf := range Inflections(*root.Input['z']) {
	//  fmt.Printf("%T\n", (*inf.N.Inst))
	//}
	//fmt.Println(len(Inflections(*root.Input['z'])))
	//var stack = []*Node{&root}
	//inpMap := make(map[*Node]map[*Node]bool)
	//vals := make(map[int]int)
	//nodes := make(map[int]*Node)
	//for ;; {
	//  if len(stack) == 0 {
	//    break
	//  }
	//  cur := stack[len(stack)-1]
	//  stack = stack[:len(stack)-1]
	//  vals[cur.Index] = len(cur.Inps)
	//  nodes[cur.Index] = cur
	//  for _, in := range cur.Input {
	//    stack = append(stack, in)
	//  }
	//  for _, inp := range cur.Inps {
	//    if _, ok := inpMap[inp]; !ok {
	//      inpMap[inp] = make(map[*Node]bool)
	//    }
	//    inpMap[inp][cur] = true
	//  }
	//}
	//inpList := make(map[*Node]NodeList)
	//for k := range inpMap {
	//  for n := range inpMap[k] {
	//    inpList[k] = append(inpList[k], n)
	//  }
	//  fmt.Println(len(inpList[k]))
	//  sort.Sort(inpList[k])
	//}
	////for _, inf := range Inflections(root) {
	////  fmt.Printf("%T\n", (*inf.N.Inst))
	////  switch v := (*inf.N.Inst).(type) {
	////  case Eql:
	////    _ = v
	////    fmt.Println(inf.N.Inps)
	////    for r, _ := range inf.N.Input {
	////      for _, tnf := range Inflections(*inf.N.Input[r]) {
	////        fmt.Printf("    %T\n", (*tnf.N.Inst))
	////      }
	////      fmt.Println(string(r), inf.N.Input[r].Index, inf.N.Index, len(Inflections(*inf.N.Input[r])))
	////    }
	////  }
	////}
	////BackProp(root)
	////for i := 0; i < len(vals); i++  {
	////  fmt.Println(i, vals[i])
	////}
	//for i, n := range Inflections(root) {
	//  fmt.Println(nodes[i].Index, len(n.N.Inps), len(n.Values))
	//}
	//fmt.Println(len(nodes))
	//fmt.Println(Calc(root, inputs, 14, inpList))
	//fmt.Println(len(BackProp(&root)))
	//for _, v := range BackProp(&root) {
	//  fmt.Println(FindVal(v), v.Value)
	//  if v.Value == 0 {
	//    fmt.Println("Found 0")
	//  }
	//}
	//fmt.Println(TryVal(make(Globe), instructs, 14))
	carryover := map[int]int{0: 0}
	for i, phase := range Phases(instructs) {
		carryover = phase.Do(carryover, i)
		fmt.Println(i, phase.Input, len(carryover))
	}
	fmt.Println(carryover[0])
}
