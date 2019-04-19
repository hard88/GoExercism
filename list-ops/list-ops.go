package listops

type binFunc func(x, y int) int

type predFunc func(n int) bool

type unaryFunc func(n int) int



type IntList []int


func (list IntList) Foldr (f binFunc, initual int) int {
	for i := range list {
		//initual = f(initual, list[len(list)-i-1])
		initual = f(list[len(list)-i-1], initual)
	}
	return initual
}

func (list IntList) Foldl (f binFunc, initual int) int {
	for i := range list {
		initual = f(initual, list[i])
	}
	return initual
}

func (list IntList) Filter (f predFunc) IntList {

	l := make(IntList, 0)
	for i := range list{
		if f(list[i]) {
			l = append(l, list[i])
		}
	}

	return l
}

func (list IntList) Length () int {
	return len(list)
}


func (list IntList) Map (f unaryFunc) IntList {
	l := make(IntList, list.Length())
	for i, v := range list {
		l[i] = f(v)
	}
	return l
}


func (list IntList) Reverse () IntList {
	for i, j := 0, len(list) - 1; i < j; i, j = i+1, j -1 {
		list[i], list[j] = list[j], list[i]
	}
	return list
}

func (list IntList) Append (intList IntList) IntList {
	l := make(IntList, 0)
	l = append(l, list...)
	l = append(l, intList...)
	return l
}


func (list IntList) Concat (intLList []IntList) IntList {
	l := make(IntList, 0)
	l = l.Append(list)
	for _, i := range intLList{
		l = l .Append(i)
	}
	return l
}








