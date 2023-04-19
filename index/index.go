package index

type Key interface {
	int
}

type Value interface {
	any
}

type Operations[K Key, V Value] interface {
	Insert(k K, v V) bool
	Delete(k K) bool
	Search(k K) V
}
