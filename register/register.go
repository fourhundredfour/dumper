package register

import (
	"github.com/fourhundredfour/dumper/tree"
	"sync"
)

type Register struct {
	mu    sync.Mutex
	trees map[string]*tree.Tree
}

func (register *Register) generateMapHash(treeInstance *tree.Tree) string {
	return treeInstance.Table + "_" + treeInstance.Column + "_" + treeInstance.Value
}

func (register *Register) Add(treeInstance *tree.Tree) {
	register.mu.Lock()
	defer register.mu.Unlock()
	register.trees[register.generateMapHash(treeInstance)] = treeInstance
}

func (register *Register) Exists(treeInstance *tree.Tree) bool {
	register.mu.Lock()
	defer register.mu.Unlock()
	_, ok := register.trees[register.generateMapHash(treeInstance)]
	return ok
}

func New() *Register {
	return &Register{
		trees: make(map[string]*tree.Tree),
	}
}
