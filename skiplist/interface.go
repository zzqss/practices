package main

type SkipLister interface {
	Insert(key Key) Node
	Get(key Key) Node
	Delete(ket Key) Node
}

type IntKey struct {
	Key int
}

func NewIntKey(key int) *IntKey {
	return &IntKey{key}
}

func (this *IntKey) GetNum() int {
	return this.Key
}
func (this *IntKey) Compare(key Key) int8 {
	return (int8)(this.Key - key.GetNum())
}
