package utils

import "container/list"

type Cache struct {
	Size      int
	evictList *list.List
}
