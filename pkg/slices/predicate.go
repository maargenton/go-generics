package slices

func Any[T any](s []T, p func(v T) bool) bool {
	for _, v := range s {
		if p(v) {
			return true
		}
	}
	return false
}

func All[T any](s []T, p func(v T) bool) bool {
	for _, v := range s {
		if !p(v) {
			return false
		}
	}
	return true
}

func None[T any](s []T, p func(v T) bool) bool {
	for _, v := range s {
		if p(v) {
			return false
		}
	}
	return true
}

func One[T any](s []T, p func(v T) bool) bool {
	var found = false
	for _, v := range s {
		if p(v) {
			if !found {
				found = true
			} else {
				return false
			}
		}
	}
	return found
}
