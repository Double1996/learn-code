package observer

import "testing"

func TestObserver_Notify(t *testing.T) {
	sub := &Subject{}
	sub.Register(Observer1{})
	sub.Register(Observer2{})
	sub.Notify("hi")
}