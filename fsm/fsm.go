package main

import "fmt"

type State interface {
	AddTask(func(int) (int, error))
	DoTask(int) (int, error)
	Save() error
	IsFinal() bool
	GetValue() int
}

type FSM struct {
	stateArr []State
}

type DefaultState struct {
	Value int
	task func(int) (int, error)
	isFinal bool
}

func(s *DefaultState) AddTask(t func(int) (int, error)) {
	s.task = t
}

func(s *DefaultState) GetValue() int {
	return s.Value
}

func(s *DefaultState) DoTask(value int) (int, error) {
	res, err := s.task(value)
	if err != nil {
		return 0, err
	}
	fmt.Println(res)
	return res, nil
}

func(s *DefaultState) Save() error {
	//save db
	return nil
}

func(s *DefaultState) IsFinal() bool {
	return s.isFinal
}


func (f *FSM) Run() {
	res := f.stateArr[0].GetValue()
	var err error
	for _, sta := range f.stateArr {
		res, err = sta.DoTask(res)
		if err != nil {
			return
		}
		err = sta.Save()
		if err != nil {
			return
		}
	}
}


func Add(d int) (int, error) {
	fmt.Println("add...")
	return d+1, nil
}

func Sub(d int) (int, error) {
	fmt.Println("sub...")
	return d-1, nil
}

func main() {

	s1 := DefaultState{
		Value: 2,
		isFinal: false,
	}
	s1.AddTask(Add)

	s2 := DefaultState{
		isFinal: false,
	}
	s2.AddTask(Sub)

	fsm := FSM{
		stateArr: []State{&s1, &s2},
	}
	fsm.Run()

}



