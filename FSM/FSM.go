package FSM

import (
	"fmt"
	"sync"
)

type FSMState string            //状态
type FSMEvent string            //事件
type FSMHandler func() FSMState //处理方法，并返回新的状态

//有限状态机
type FSM struct {
	mu       sync.Mutex                           //排它锁
	state    FSMState                             //当前状态
	handlers map[FSMState]map[FSMEvent]FSMHandler //每个状态都有多个事件，给个事件执行一个handler
}

//获取当前状态
func (f *FSM) getState() FSMState {
	return f.state
}

//设置当前状态
func (f *FSM) setState(newState FSMState) {
	f.state = newState
}

//添加事件处理方法
func (f *FSM) AddHandler(state FSMState, event FSMEvent, handler FSMHandler) *FSM {
	if _, ok := f.handlers[state]; !ok {
		f.handlers[state] = make(map[FSMEvent]FSMHandler)
	}

	if _, ok := f.handlers[state][event]; ok {
		fmt.Printf("[警告]状态（%s）事件（%s）已定义过", state, event)
	}
	f.handlers[state][event] = handler
	return f
}

//事件处理
func (f *FSM) Call(event FSMEvent) FSMState {
	f.mu.Lock()
	defer f.mu.Unlock()

	events := f.handlers[f.getState()]
	if events == nil {
		return f.getState()
	}

	if fn, ok := events[event]; ok {
		oldState := f.getState()
		f.setState(fn())
		newState := f.getState()
		fmt.Printf("状态从[%s]变成[%s]\n", oldState, newState)
	}

	return f.getState()
}

//实例化FSM
func CreateFSM(initState FSMState) *FSM {
	return &FSM{
		state:    initState,
		handlers: make(map[FSMState]map[FSMEvent]FSMHandler),
	}
}
