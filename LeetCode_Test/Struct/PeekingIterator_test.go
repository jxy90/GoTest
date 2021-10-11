package Struct

type Iterator struct {
}

func (this *Iterator) hasNext() bool {
	// Returns true if the iteration has more elements.
	return false
}

func (this *Iterator) next() int {
	// Returns the next element in the iteration.
	return 0
}

type PeekingIterator struct {
	iter     *Iterator
	_hasNext bool
	_next    int
}

func ConstructorPeekingIterator(iter *Iterator) *PeekingIterator {
	return &PeekingIterator{
		iter:     iter,
		_hasNext: iter.hasNext(),
		_next:    iter.next(),
	}
}

func (this *PeekingIterator) hasNext() bool {
	return this._hasNext
}

func (this *PeekingIterator) next() int {
	if !this._hasNext {
		return 0
	}
	ret := this._next
	this._hasNext = this.iter.hasNext()
	if this._hasNext {
		this._next = this.iter.next()
	}
	return ret
}

func (this *PeekingIterator) peek() int {
	return this._next
}
