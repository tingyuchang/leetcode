package queue


type queue interface {
	Add(interface{})
	Remove() (interface{})
}

type list struct {
	Values []interface{}
}

func (l *list) Add(n interface{})  {
	l.Values = append(l.Values, n)
}

func (l *list) Remove() interface{} {
	n := l.Values[len(l.Values)-1]
	l.Values = l.Values[:len(l.Values)-1]
	return n
}