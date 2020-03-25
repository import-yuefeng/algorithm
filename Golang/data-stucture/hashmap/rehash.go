package main

func (h *Hashmap) rehash() {
	if h.inRehash {
		if h.newBucket == nil || len(h.newBucket) <= len(h.oldBucket) || len(h.newBucket) == 0 {
			h.newBucket = make([]*node, h.size*2)
		}
		curidx := h.rehashidx + 1
		if h.oldBucket[curidx] == nil {
			// pass
		} else {
			k := Hash(h.oldBucket[curidx].key)
			bucketidx := k % uint32(len(h.newBucket))
			newSite := h.newBucket[bucketidx]
			if newSite != nil {
				for newSite.Next != nil {
					newSite = newSite.Next
				}
				newSite = h.oldBucket[curidx]
				h.oldBucket[curidx] = nil
			} else {
				h.newBucket[bucketidx] = h.oldBucket[curidx]
				h.oldBucket[curidx] = nil
			}
		}
		h.rehashidx++
		if h.rehashidx == len(h.oldBucket)-1 {
			h.inRehash = false
			h.rehashidx = -1
			h.size *= 2
			h.oldBucket, h.newBucket = h.newBucket, h.oldBucket
		}
	} else {
		if h.size == 0 || h.used == 0 {
			return
		}
		curLoad := h.used / h.size
		if curLoad > h.load {
			h.inRehash = true
			h.rehash()
		}
	}
}
