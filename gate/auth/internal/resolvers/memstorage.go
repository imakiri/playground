package resolvers

import (
	"errors"
	"github.com/imakiri/playground/core"
	"github.com/imakiri/playground/gate/auth/internal/types"
	"sync"
)

type MemStorage struct {
	data     map[types.Factor_Rand]types.ID
	data_rev map[types.ID]types.Factor_Rand
	rwmux    *sync.RWMutex
}

func (mem *MemStorage) Enrol(factors ...core.Factor) (bool, error) {
	var r_id = factors[len(factors)-2]
	var r_rand = factors[len(factors)-1]

	var id types.ID
	var rand types.Factor_Rand

	switch t := r_id.(type) {
	case types.ID:
		id = t
	default:
		return false, errors.New("factor mismatch: ID")
	}

	switch t := r_rand.(type) {
	case types.Factor_Rand:
		rand = t
	default:
		return false, errors.New("factor mismatch: Rand")
	}

	return mem.create(&rand, &id)
}

func (mem *MemStorage) Withdraw(factors ...core.Factor) (core.ID, error) {
	var r_id = factors[len(factors)-2]
	var r_rand = factors[len(factors)-1]

	var id *types.ID
	var rand *types.Factor_Rand

	switch t := r_id.(type) {
	case types.ID:
		id = &t
	}

	switch t := r_rand.(type) {
	case types.Factor_Rand:
		rand = &t
	}

	return mem.delete(rand, id)
}

func (mem *MemStorage) Verify(factors ...core.Factor) (core.ID, error) {
	var r_rand = factors[len(factors)-1]
	var rand types.Factor_Rand

	switch t := r_rand.(type) {
	case types.Factor_Rand:
		rand = t
	default:
		return false, errors.New("factor mismatch: Rand")
	}

	var id types.ID
	return mem.readID(&rand, &id)
}

func (mem *MemStorage) create(rand *types.Factor_Rand, id *types.ID) (bool, error) {
	if !core.IsNilSafe(rand, id) {
		return false, errors.New("nil factor")
	}

	mem.rwmux.Lock()
	defer mem.rwmux.Unlock()

	if _, ok := mem.data[*rand]; ok {
		return false, errors.New("given rand factor already exists")
	} else {
		mem.data[*rand] = *id
		mem.data_rev[*id] = *rand
		return true, nil
	}
}

func (mem *MemStorage) readID(rand *types.Factor_Rand, id *types.ID) (bool, error) {
	if !core.IsNilSafe(rand, id) {
		return false, errors.New("nil factor")
	}

	mem.rwmux.RLock()
	defer mem.rwmux.RUnlock()

	if v, ok := mem.data[*rand]; ok {
		*id = v
		return true, nil
	} else {
		return false, errors.New("no id exists for given rand factor")
	}
}

func (mem *MemStorage) delete(rand *types.Factor_Rand, id *types.ID) (bool, error) {
	var path = core.IsNilSafeEx(rand, id)

	if !path[0] && !path[1] {
		return false, errors.New("both factors are nil")
	}

	mem.rwmux.Lock()
	defer mem.rwmux.Unlock()

	if path[0] {
		if v, ok := mem.data[*rand]; ok {
			delete(mem.data_rev, v)
			delete(mem.data, *rand)
			return true, nil
		} else {
			return false, errors.New("nothing to delete")
		}
	}

	if path[1] {
		if v, ok := mem.data_rev[*id]; ok {
			delete(mem.data, v)
			delete(mem.data_rev, *rand)
			return true, nil
		} else {
			return false, errors.New("nothing to delete")
		}
	}

	return false, errors.New("delete internal error")
}
