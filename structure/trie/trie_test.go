package trie

import (
    "reflect"
    "sort"
    "testing"
)

func Test_runeTrie_SearchByPrefix(t *testing.T) {
    tr := &runeTrie{root: &runeNode{make(map[rune]*runeNode), false}}
    tr.Insert("周杰伦")
    tr.Insert("周星驰")
    tr.Insert("周杰龙")
    tr.Insert("周杰")
    
    type args struct {
        prefix string
    }
    tests := []struct {
        name     string
        instance *runeTrie
        args     args
        want     []string
    }{
        {
            instance: tr,
            args:     args{prefix: "周"},
            want: []string{
                "周杰伦",
                "周杰龙",
                "周星驰",
                "周杰",
            },
        },
        {
            instance: tr,
            args:     args{prefix: "周杰"},
            want: []string{
                "周杰伦",
                "周杰龙",
                "周杰",
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t1 *testing.T) {
            got := tt.instance.SearchByPrefix(tt.args.prefix)
            sort.Strings(got)
            sort.Strings(tt.want)
            if !reflect.DeepEqual(got, tt.want) {
                t1.Errorf("SearchByPrefix() = %v, want %v", got, tt.want)
            }
        })
    }
}

func Test_runeTrie_Size(t *testing.T) {
    tr := &runeTrie{root: &runeNode{make(map[rune]*runeNode), false}}
    tr.Insert("周杰伦")
    tr.Insert("周星驰")
    tr.Insert("周杰龙")
    tr.Insert("周杰")
    
    tests := []struct {
        name string
        want int
    }{
        {
            want: 4,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t1 *testing.T) {
            t := tr
            if got := t.Size(); got != tt.want {
                t1.Errorf("Size() = %v, want %v", got, tt.want)
            }
        })
    }
}
