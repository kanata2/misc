package main

import "testing"

var (
	data = `key1:ok	key2:wow...	key3:unko!
key4:ok	key5:uuunko	key6:aaaa`
)

func TestParse(t *testing.T) {
	d := Parse(data)
}
