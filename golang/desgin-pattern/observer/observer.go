package observer

import "fmt"

type ISubject interface {
	Register(observer IObserver)
	Remove(observer IObserver)
	Notify(msg string)
}

// 观察者
type IObserver interface {
	Update(msg string)
}

// subject 被观察者，依赖的对象
type Subject struct {
	observers []IObserver
}


// Register 注册 观察者
func (sub *Subject) Register(observer IObserver) {
	sub.observers = append(sub.observers, observer)
}


// Remove 观察者
func (sub *Subject) Remove(observer IObserver) {
	for i, ob := range sub.observers {
				if ob == observer {
					sub.observers = append(sub.observers[:i], sub.observers[i+1:]...)
				}
			}
}

// Notify 通知
func (sub *Subject) Notify(msg string) {
	for _, o := range sub.observers {
		o.Update(msg)
	}
}


// Observer1 Observer1
type Observer1 struct{}

// Update 实现观察者接口
func (Observer1) Update(msg string) {
	fmt.Printf("Observer1: %s", msg)
}

// Observer2 Observer2
type Observer2 struct{}

// Update 实现观察者接口
func (Observer2) Update(msg string) {
	fmt.Printf("Observer2: %s", msg)
}

