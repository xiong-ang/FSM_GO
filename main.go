package main

import (
	"FSM_GO/FSM"
	"fmt"
)

var (
	Poweroff        = FSM.FSMState("关闭")
	FirstGear       = FSM.FSMState("一档")
	SecondGear      = FSM.FSMState("二挡")
	ThirdGear       = FSM.FSMState("三挡")
	PoweroffEvent   = FSM.FSMEvent("按下关闭")
	FirstGearEvent  = FSM.FSMEvent("按下一档")
	SecondGearEvent = FSM.FSMEvent("按下二挡")
	ThirdGearEvent  = FSM.FSMEvent("按下三挡")
)

//电风扇
type ElectricFan struct {
	*FSM.FSM
}

func (e *ElectricFan) Init() {
	//状态Handlers
	PoweroffHandler := func() FSM.FSMState {
		fmt.Println("电风扇已关闭")
		return Poweroff
	}
	FirstGearHandler := func() FSM.FSMState {
		fmt.Println("电风扇开启一档")
		return FirstGear
	}
	SecondGearHandler := func() FSM.FSMState {
		fmt.Println("电风扇开启二挡")
		return SecondGear
	}
	ThirdGearHandler := func() FSM.FSMState {
		fmt.Println("电风扇开启三挡")
		return ThirdGear
	}

	//关闭状态
	e.AddHandler(Poweroff, PoweroffEvent, PoweroffHandler)
	e.AddHandler(Poweroff, FirstGearEvent, FirstGearHandler)
	e.AddHandler(Poweroff, SecondGearEvent, SecondGearHandler)
	e.AddHandler(Poweroff, ThirdGearEvent, ThirdGearHandler)

	//一档状态
	e.AddHandler(FirstGear, PoweroffEvent, PoweroffHandler)
	e.AddHandler(FirstGear, FirstGearEvent, FirstGearHandler)
	e.AddHandler(FirstGear, SecondGearEvent, SecondGearHandler)
	e.AddHandler(FirstGear, ThirdGearEvent, ThirdGearHandler)

	//二挡状态
	e.AddHandler(SecondGear, PoweroffEvent, PoweroffHandler)
	e.AddHandler(SecondGear, FirstGearEvent, FirstGearHandler)
	e.AddHandler(SecondGear, SecondGearEvent, SecondGearHandler)
	e.AddHandler(SecondGear, ThirdGearEvent, ThirdGearHandler)

	//三挡状态
	e.AddHandler(ThirdGear, PoweroffEvent, PoweroffHandler)
	e.AddHandler(ThirdGear, FirstGearEvent, FirstGearHandler)
	e.AddHandler(ThirdGear, SecondGearEvent, SecondGearHandler)
	e.AddHandler(ThirdGear, ThirdGearEvent, ThirdGearHandler)
}

func main() {
	efan := &ElectricFan{
		FSM: FSM.CreateFSM(Poweroff),
	}
	efan.Init()

	efan.Call(ThirdGearEvent)
	efan.Call(FirstGearEvent)
	efan.Call(SecondGearEvent)
	efan.Call(PoweroffEvent)
}
