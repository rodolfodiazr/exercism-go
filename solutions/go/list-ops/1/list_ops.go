package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
    length := s.Length()
    acc := initial
    for i := 0; i < length; i++ {
        acc = fn(acc, s[i])
    }
	return acc
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	length := s.Length()
    acc := initial
	for i := length-1; i >= 0; i-- {
		acc = fn(s[i], acc)
	}
	return acc
}

func (s IntList) Filter(fn func(int) bool) IntList {
    out := IntList{}
    for i := range s {
        if fn(s[i]) {
            out = out.Append(IntList{s[i]})
        }
    }
	return out
}

func (s IntList) Length() int {
	count := 0
    for range s {
        count++
	}
    return count
}

func (s IntList) Map(fn func(int) int) IntList {
    for i := range s {
        s[i] = fn(s[i])
    }
	return s
}

func (s IntList) Reverse() IntList {
	length := s.Length()
	out := make(IntList, length)
	for i := range s {
		out[length-1-i] = s[i]
	}
	return out
}

func (s IntList) Append(lst IntList) IntList {
    initialLength := s.Length()
    out := make(IntList, initialLength + lst.Length())
    for i := range s {
        out[i] = s[i]
    }
    for i := range lst {
        out[initialLength+i] = lst[i]
    }
	return out
}

func (s IntList) Concat(lists []IntList) IntList {
    for i := range lists {
        s = s.Append(lists[i])
    }
	return s
}
