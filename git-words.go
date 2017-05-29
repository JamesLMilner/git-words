package main

import (
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"fmt"

	"github.com/speedata/gogit"
)

func main() {

	wordcount := make(map[string]int)

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	repository, err := gogit.OpenRepository(filepath.Join(wd, ".git"))
	if err != nil {
		log.Fatal(err)
	}

	ref, err := repository.LookupReference("HEAD")
	if err != nil {
		log.Fatal(err)
	}

	ci, err := repository.LookupCommit(ref.Oid)
	parent := ci.Parent(0)

	for parent != nil {

		msg := parent.CommitMessage
		//fmt.Println(msg)
		ci, err := repository.LookupCommit(parent.Oid)
		if err != nil {
			//fmt.Print(ci.CommitMessage)
		}
		parent = ci.Parent(0)

		// Word processing

		replacer := strings.NewReplacer(",", "", ".", "", ";", "", "!", "", ":", "", "-", "")
		msg = replacer.Replace(msg)
		words := strings.Fields(msg)
		for _, word := range words {
			wordcount[word]++
			//fmt.Println(i, " => ", word)
		}

	}

	pl := rankByWordCount(wordcount)
	for _, v := range pl {
		if v.Value > 1 {
			fmt.Println(v.Key, v.Value)
		}
	}

}

func rankByWordCount(wordFrequencies map[string]int) PairList {
	pl := make(PairList, len(wordFrequencies))
	i := 0
	for k, v := range wordFrequencies {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
