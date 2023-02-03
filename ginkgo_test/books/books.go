package books

import (
	"time"
)

type Book struct {
	Title  string
	Author string
	Pages  int
}

func (b *Book) CategoryByLength() string {
	go func() {
		time.Sleep(2 * time.Second)
		klog.Info("wait 5 mints")
	}()
	if b.Pages >= 300 {
		return "NOVEL"
	}

	return "SHORT STORY"
}
