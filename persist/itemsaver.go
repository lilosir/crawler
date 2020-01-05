package persist

import "log"

//ItemSaver save the items from every request
func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item $%d: %v", itemCount, item)
			itemCount++
		}
	}()
	return out
}