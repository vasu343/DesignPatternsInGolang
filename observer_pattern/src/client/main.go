package main

import (
	obs "client/observerpattern"
)

// client code for driving the observer pattern
func main() {
	subject := obs.SubjectModel{Marks:2,ObserverList:[]obs.Observer{}}
	observer1 := obs.ObserverModel{SubjectModel: &subject}
	observer2 := obs.ObserverModel2{SubjectModel: &subject}
	subject.AttachSubscriber(&observer1)
	subject.AttachSubscriber(&observer2)
	subject.SetState(23)
	subject.Notify()
	subject.DetachSubscriber(&observer1)
	subject.SetState(24)
	subject.Notify()
}
