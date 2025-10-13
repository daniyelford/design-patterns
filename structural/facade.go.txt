package main

import "fmt"

// Subsystems
type Light struct{}
func (l Light) On() { fmt.Println("Lights ON") }
func (l Light) Off() { fmt.Println("Lights OFF") }

type AC struct{}
func (a AC) On() { fmt.Println("AC ON") }
func (a AC) Off() { fmt.Println("AC OFF") }

type Security struct{}
func (s Security) Arm() { fmt.Println("Security Armed") }
func (s Security) Disarm() { fmt.Println("Security Disarmed") }

// Facade
type SmartHomeFacade struct {
	light    Light
	ac       AC
	security Security
}

func (f SmartHomeFacade) LeaveHome() {
	f.light.Off()
	f.ac.Off()
	f.security.Arm()
}

func (f SmartHomeFacade) ArriveHome() {
	f.security.Disarm()
	f.light.On()
	f.ac.On()
}

func main() {
	home := SmartHomeFacade{}
	home.LeaveHome()
	home.ArriveHome()
}
