package main

import "fmt"


/*
 这个状态机的实现主要是把State抽象了出来。
其实不用抽象出来，直接构建一个结构体，里面包含数据，配置信息，然后按照配置信息去执行就可以，bus_svr就是这样实现的
 */
type State interface {
	SetValue(Doc)
	AddTask(func(Doc) (Doc, error))
	Execute() error
	Save() error
	GetValue() Doc
	IsFinal() bool
}

type FSM struct {
	stateMap map[int]State
}

type Doc struct {
	Title string
	Content string
}


type DocState struct {
	value Doc
	task func(Doc) (Doc, error)
	isFinal bool
}

func(s *DocState) AddTask(t func(Doc) (Doc, error)) {
	s.task = t
}

func(s *DocState) GetValue() Doc {
	return s.value
}

func(s *DocState) SetValue(doc Doc) {
	s.value = doc
}

func(s *DocState) Execute() error {
	res, err := s.task(s.value)
	if err != nil {
		return err
	}
	s.value = res
	fmt.Println(s.value)
	return nil
}

func(s *DocState) Save() error {
	//save db
	return nil
}

func(s *DocState) IsFinal() bool {
	return s.isFinal
}


func (f *FSM) Run() {

	status := 1

	var err error
	for err == nil {//Condition is always 'true' because 'err' is always 'nil'，这个地方goland的提示有问题
		if f.stateMap[status].IsFinal() {
			fmt.Println("final...")
			return
		}
		err = f.stateMap[status].Execute()
		if err != nil {
			fmt.Println("err:", err)
			return
		}
		err = f.stateMap[status].Save()
		if err != nil {
			return
		}
		res := f.stateMap[status].GetValue()
		status = status + 1
		f.stateMap[status].SetValue(res)
	}
}


func ModifyTitle(doc Doc) (Doc, error) {
	fmt.Println("modify title")
	doc.Title = "title test"
	return doc, nil
}

func ModifyContent(doc Doc) (Doc, error) {
	fmt.Println("modify content")
	if doc.Content == "" {
		return doc, fmt.Errorf("doc content empty")
	}
	doc.Content = "content test"
	return doc, nil
}

func main() {

	doc := Doc{}
	s1 := &DocState{
		value: doc,
		isFinal: false,
	}
	s1.AddTask(ModifyTitle)

	s2 := &DocState{
		isFinal: false,
	}
	s2.AddTask(ModifyContent)

	s3 := &DocState{
		isFinal: true,
	}


	fsm := FSM{
		stateMap: map[int]State{1:s1,2:s2,3:s3},
	}
	fsm.Run()

}



