package main

type observable interface {
	registerObserver(obs observer)
	unRegisterObserver(obs observer)
	notifyAll()
}

type DataSubject struct {
	observers []DataListener
	field     string
}

func (ds *DataSubject) ChangeItem(data string) {
	ds.field = data
	ds.notifyAll()
}

func (ds *DataSubject) registerObserver(o DataListener) {
	ds.observers = append(ds.observers, o)
}

func (ds *DataSubject) unRegisterObserver(o DataListener) {
	var newObservers []DataListener
	for _, observer := range ds.observers {
		if o.Name != observer.Name {
			newObservers = append(newObservers, observer)
		}
	}
	ds.observers = newObservers
}

func (ds *DataSubject) notifyAll() {
	for _, observer := range ds.observers {
		observer.onUpdate(ds.field)
	}
}
