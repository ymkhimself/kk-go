package main

import "sync"

func main() {
	mutex := sync.Mutex{}
	mutex.Lock()
	mutex.TryLock()
}
