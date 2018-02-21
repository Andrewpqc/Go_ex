package main

import (
	"testing"
	"time"
)

func TestDecode(t *testing.T) {
	t.Parallel()
	post, err := decode("post.json")
	if err != nil {
		t.Error(err)
	}
	if post.Id != 1 {
		t.Errorf("wrong id ,want 1 but got %d", post.Id)
	}
	if post.Comments[1].Content != "have a good time" {
		t.Errorf("wrong comment[1].content ,want 'have a good time' but got %s",
			post.Comments[1].Content)
	}
}

func TestEncode(t *testing.T) {
	t.Parallel()
	t.Skip("skipping encode for now")
}

func TestLongRunningFunc(t *testing.T) {
	t.Parallel()
	if testing.Short() {
		t.Skip("skipping longrunningfunc for now")
	}
	time.Sleep(10 * time.Second)
}

func TestParallel_1(t *testing.T) {
	t.Parallel()
	time.Sleep(1 * time.Second)
}

func TestParallel_2(t *testing.T) {
	t.Parallel()
	time.Sleep(2 * time.Second)
}

func TestParallel_3(t *testing.T) {
	t.Parallel()
	time.Sleep(3 * time.Second)
}



func BenchmarkDecode(b *testing.B){
	for i:=0;i<b.N;i++{
		decode("post.json")
	}
}

func BenchmarkUnmarshal(b *testing.B){
	for i:=0;i<b.N;i++{
		unmarshal("post.json")
	}
}