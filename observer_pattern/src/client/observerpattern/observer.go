package observerpattern

import (
	"fmt"
	"log"
)

// Abstract Publisher
type Subject interface {
	AttachSubscriber(Observer)
	DetachSubscriber(Observer)
	GetState() int
	SetState(int)
	Notify()
}

// Abstract Observer
type Observer interface {
	SetSubject(Subject)
	Update()
}

// Concrete Publisher 1
type SubjectModel struct {
	Marks        int
	ObserverList []Observer
}

// Concrete Observer 1
type ObserverModel struct {
	SubjectModel Subject
	Marks        int
}

// Concrete Observer 2
type ObserverModel2 struct {
	SubjectModel Subject
	Marks        int
}

func remove(list []Observer,ind int) []Observer {
	list[ind] = list[len(list)-1]
	list[len(list) - 1] = nil
	return list[:len(list)-1]
}

// Publisher Functionalities
func (s *SubjectModel) AttachSubscriber(o Observer) {
	o.SetSubject(s)
	s.ObserverList = append(s.ObserverList, o)
}
func (s *SubjectModel) Notify() {
	for _, o := range s.ObserverList {
		o.Update()
	}
}
func (s *SubjectModel) SetState(m int) {
	s.Marks = m
}
func (s *SubjectModel) GetState() int {
	return s.Marks
}
func (s *SubjectModel) DetachSubscriber(o Observer) {
	o.SetSubject(nil)
	for i,_ := range s.ObserverList {
		if o == s.ObserverList[i] {
			s.ObserverList = remove(s.ObserverList,i)
			break
		}
	}
}

// Observer Functionalities
func (o *ObserverModel) SetSubject(s Subject) {
	o.SubjectModel = s
}
func (o *ObserverModel) Update() {
	if o == nil {
		log.Print("O1 is nil")
	} else {
		m := o.SubjectModel.GetState()
		fmt.Printf("Observer Model 1 :- %d\n", m)
	}
}

func (o *ObserverModel2) SetSubject(s Subject) {
	o.SubjectModel = s
}
func (o *ObserverModel2) Update() {
	if o == nil {
		log.Print("O2 is nil")
	} else {
		m := o.SubjectModel.GetState()
		fmt.Printf("Observer Model 2 :- %d\n", m)
	}
}
