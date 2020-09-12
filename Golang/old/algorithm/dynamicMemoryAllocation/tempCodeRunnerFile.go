func (mt *MemoryTable) RecoveryByAddress(startAddress int) error {
// 	res := FindMemoryBlockByStartAddree(startAddress)
// 	if res != nil {
// 		/*
// 			1. 上下都已分配 res.Front.IsUsed && res.Next.IsUsed (true)
// 			2. 上分配, 下未分配 res.Front.IsUsed && !res.Next.IsUsed (true)
// 			3. 下分配, 上未分配 !res.Front.IsUsed && res.Next.IsUsed (true)
// 			4. 上下均未分配 !res.Front.IsUsed && !res.Next.IsUsed (true)
// 		**/

// 		if res.Front.IsUsed && res.Next.IsUsed {
// 			res.IsUsed = false
// 		} else if res.Front.IsUsed && !res.Next.IsUsed {
// 			appendMemory := res.Next
// 			res.Size += appendMemory.Size
// 			res.Next = appendMemory.Next

// 		} else if !res.Front.IsUsed && res.Next.IsUsed {

// 		} else {

// 		}

// 	} else {
// 		errors.New("not found!")
// 	}
// 	return nil
// }
