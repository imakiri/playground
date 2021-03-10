package data

//type MemStorage struct {
//	data     map[internal.Key_Random]internal.ID
//	data_rev map[internal.ID]internal.Key_Random
//	rwmux    *sync.RWMutex
//}
//
//func (mem *MemStorage) Enrol(factors ...core.Factor) (bool, error) {
//	var r_id = factors[len(factors)-2]
//	var r_rand = factors[len(factors)-1]
//
//	var id internal.ID
//	var rand internal.Key_Random
//
//	switch t := r_id.(type) {
//	case internal.ID:
//		id = t
//	default:
//		return false, errors.New("factor mismatch: ID")
//	}
//
//	switch t := r_rand.(type) {
//	case internal.Key_Random:
//		rand = t
//	default:
//		return false, errors.New("factor mismatch: Rand")
//	}
//
//	return mem.create(&rand, &id)
//}
//
//func (mem *MemStorage) Withdraw(factors ...core.Factor) (bool, error) {
//	var r_id = factors[len(factors)-2]
//	var r_rand = factors[len(factors)-1]
//
//	var id *internal.ID
//	var rand *internal.Key_Random
//
//	switch t := r_id.(type) {
//	case internal.ID:
//		id = &t
//	}
//
//	switch t := r_rand.(type) {
//	case internal.Key_Random:
//		rand = &t
//	}
//
//	return mem.delete(rand, id)
//}
//
//func (mem *MemStorage) Verify(factors ...core.Factor) (bool, error) {
//	var r_rand = factors[len(factors)-1]
//	var rand internal.Key_Random
//
//	switch t := r_rand.(type) {
//	case internal.Key_Random:
//		rand = t
//	default:
//		return false, errors.New("factor mismatch: Rand")
//	}
//
//	var id internal.ID
//	return mem.readID(&rand, &id)
//}
//
//func (mem *MemStorage) create(rand *internal.Key_Random, id *internal.ID) (bool, error) {
//	if !core.IsNilSafe(rand, id) {
//		return false, errors.New("nil factor")
//	}
//
//	mem.rwmux.Lock()
//	defer mem.rwmux.Unlock()
//
//	if _, ok := mem.data[*rand]; ok {
//		return false, errors.New("given rand factor already exists")
//	} else {
//		mem.data[*rand] = *id
//		mem.data_rev[*id] = *rand
//		return true, nil
//	}
//}
//
//func (mem *MemStorage) readID(rand *internal.Key_Random, id *internal.ID) (bool, error) {
//	if !core.IsNilSafe(rand, id) {
//		return false, errors.New("nil factor")
//	}
//
//	mem.rwmux.RLock()
//	defer mem.rwmux.RUnlock()
//
//	if v, ok := mem.data[*rand]; ok {
//		*id = v
//		return true, nil
//	} else {
//		return false, errors.New("no id exists for given rand factor")
//	}
//}
//
//func (mem *MemStorage) delete(rand *internal.Key_Random, id *internal.ID) (bool, error) {
//	var path = core.IsNilSafeEx(rand, id)
//
//	if !path[0] && !path[1] {
//		return false, errors.New("both factors are nil")
//	}
//
//	mem.rwmux.Lock()
//	defer mem.rwmux.Unlock()
//
//	if path[0] {
//		if v, ok := mem.data[*rand]; ok {
//			delete(mem.data_rev, v)
//			delete(mem.data, *rand)
//			return true, nil
//		} else {
//			return false, errors.New("nothing to delete")
//		}
//	}
//
//	if path[1] {
//		if v, ok := mem.data_rev[*id]; ok {
//			delete(mem.data, v)
//			delete(mem.data_rev, *id)
//			return true, nil
//		} else {
//			return false, errors.New("nothing to delete")
//		}
//	}
//
//	return false, errors.New("delete internal error")
//}
