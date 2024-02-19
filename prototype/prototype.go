package prototype

func Apply[K comparable, T any, V any](fn map[K]func(T) V, dataMap map[K]T) map[K]V {
	rmap := make(map[K]V)
	for key, data := range dataMap {
		rmap[key] = fn[key](data)
	}
	return rmap
}

func Bind[K comparable, T any, V any](fn map[K]func(T) V) map[K]V {
	return func(dataMap map[K]T) map[K]V {
		Apply[K, T, V](fn, dataMap)
	}
}

func Map[K comparable, T any, V any](fn func(V) V, dataMap map[K]V) map[K]V {
	rmap := make(map[K]V)
	for key, data := range dataMap {
		rmap[key] = fn(data)
	}
	return rmap
}

func Every[K comparable, V any](dataMap map[K]V, fn func(V) bool) bool {
	for _, value := range dataMap {
		if !fn(value) {
			return false
		}
	}
	return true
}

func Some[K comparable, V any](dataMap map[K]V, fn func(V) bool) bool {
	for _, value := range dataMap {
		if fn(value) {
			return true
		}
	}
	return false
}

func Filter[K comparable, V any](dataMap map[K]V, fn func(V) bool) map[K]V {
	rmap := make(map[K]V)
	for key, value := range dataMap {
		if fn(value) {
			rmap[key] = value
		}
	}
	return rmap
}

func Reduce[K comparable, V any, R any](dataMap map[K]V, fn func(R, V) R, value R) R {
	rv := value
	for _, value := range dataMap {
		rv = fn(rv, value)
	}
	return rv
}
