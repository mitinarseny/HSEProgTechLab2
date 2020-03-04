package main

import (
	"encoding/json"
	"flag"
	"log"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"testing"
	"time"

	"github.com/emirpasic/gods/maps/treemap"
	"github.com/mitinarseny/HSEProgTechLab1/students"
	"github.com/mitinarseny/HSEProgTechLab2/search"
)

var data [][]students.Student

func Benchmark(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	for _, d := range data {
		el := d[len(d)-1]
		comparator := func(i int) bool {
			return d[i].FullName >= el.FullName
		}
		b.Logf("searching for FullName %q", el.FullName)
		b.Run(strconv.Itoa(len(d)), func(b *testing.B) {
			b.Run("SortAndBin", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					sort.Sort(students.Order(d, students.ByFullName))
					if ind := search.Bin(len(d), comparator); ind == len(d) || d[ind].FullName != el.FullName {
						b.Errorf("student with FullName %q was not found, but it is in the array", el.FullName)
					}
				}
			})
			b.Run("Bin", func(b *testing.B) {
				// now d is already sorted
				for i := 0; i < b.N; i++ {
					if ind := search.Bin(len(d), comparator); ind == len(d) || d[ind].FullName != el.FullName {
						b.Errorf("student with FullName %q was not found, but it is in the array", el.FullName)
					}
				}
			})
			b.Run("Full", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					if ind := search.Full(len(d), comparator); ind == len(d) || d[ind].FullName != el.FullName {
						b.Errorf("student with FullName %q was not found, but it is in the array", el.FullName)
					}
				}
			})
			b.Run("HashTableMap", func(b *testing.B) {
				m := make(map[string]students.Student, len(d))
				for _, s := range d {
					m[s.FullName] = s
				}
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					_, found := m[el.FullName]
					if !found {
						b.Errorf("student wih FullName %q was not found, but it is in the map", el.FullName)
					}
				}
			})
			b.Run("RedBlackTreeMap", func(b *testing.B) {
				m := treemap.NewWithStringComparator()
				for _, s := range d {
					m.Put(s.FullName, s)
				}
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					_, found := m.Get(el.FullName)
					if !found {
						b.Errorf("student wih FullName %q was not found, but it is in the map", el.FullName)
					}
				}
			})
		})
	}
}

func TestMain(m *testing.M) {
	flag.Parse()
	if flag.NArg() == 0 {
		log.Print("no data was provided")
		os.Exit(0)
	}

	data = make([][]students.Student, 0, flag.NArg())
	for _, fname := range flag.Args() {
		f, err := os.Open(fname)
		if err != nil {
			log.Print(err)
			continue
		}

		var s []students.Student
		if err := json.NewDecoder(f).Decode(&s); err != nil {
			log.Printf("unable to parse JSON: %s", err)
		}
		f.Close()
		data = append(data, s)
	}
	os.Exit(m.Run())
}
