package main

import "strings"

func Parse(s string) []map[string]string {
	r := []map[string]string{}
	ls := strings.Split(s, "\n")
	for _, l := range ls {
		r = append(r, ParseLine(l))
	}
	return r
}

func ParseLine(s string) map[string]string {
	r := map[string]string{}
	columns := strings.Split(s, "\t")
	for _, column := range columns {
		kv := strings.Split(column, ":")
		r[kv[0]] = kv[1]
	}
	return r
}
