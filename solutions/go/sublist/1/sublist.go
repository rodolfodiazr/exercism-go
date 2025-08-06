package sublist

// Relation type is defined in relations.go file.

func Sublist(l1, l2 []int) Relation {
	switch {
	case equal(l1, l2):
		return RelationEqual
	case isSublist(l1, l2):
		return RelationSublist
	case isSublist(l2, l1):
		return RelationSuperlist
	default:
		return RelationUnequal
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
    
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func isSublist(sub, full []int) bool {
	if len(sub) == 0 {
		return true
	}
    
	if len(sub) > len(full) {
		return false
	}
    
	for i := 0; i <= len(full)-len(sub); i++ {
		match := true
		for j := range sub {
			if full[i+j] != sub[j] {
				match = false
				break
			}
		}

		if match {
			return true
		}
	}
	return false
}