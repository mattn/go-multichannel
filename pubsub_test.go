package pubsub_test

import (
	"github.com/mattn/go-pubsub"
	"testing"
)

func TestInt(t *testing.T) {
	done := make(chan int)
	ps := pubsub.New()
	ps.Sub(func(i int) {
		done <- i
	})
	ps.Pub(1)
	i := <-done
	if i != 1 {
		t.Fatalf("Expected %v, but %d:", 1, i)
	}
}

func TestString(t *testing.T) {
	done := make(chan string)
	ps := pubsub.New()
	ps.Sub(func(s string) {
		done <- s
	})
	ps.Pub("hello world")
	s := <-done
	if s != "hello world" {
		t.Fatalf("Expected %v, but %d:", "hello world", s)
	}
}

type F struct {
	m string
}

func TestStruct(t *testing.T) {
	done := make(chan *F)
	ps := pubsub.New()
	ps.Sub(func(f *F) {
		done <- f
	})
	ps.Pub(&F{"hello world"})
	f := <-done
	if f.m != "hello world" {
		t.Fatalf("Expected %v, but %d:", "hello world", f.m)
	}
}

func TestOnly(t *testing.T) {
	doneInt := make(chan int)
	doneF := make(chan *F)
	ps := pubsub.New()
	ps.Sub(func(i int) {
		doneInt <- i
	})
	ps.Sub(func(f *F) {
		doneF <- f
	})
	ps.Pub(&F{"hello world"})
	ps.Pub(2)
	i := <-doneInt
	f := <-doneF
	if f.m != "hello world" {
		t.Fatalf("Expected %v, but %d:", "hello world", f.m)
	}
	if i != 2 {
		t.Fatalf("Expected %v, but %d:", 2, f.m)
	}
}
