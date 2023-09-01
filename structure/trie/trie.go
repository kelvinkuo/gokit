package trie

import (
    "unicode/utf8"
)

type Trie interface {
    // Insert inserts a word into the trie and returns true
    // if the word already exists.
    Insert(word string) bool
    // Contains returns true if an exact match of the word is found, otherwise false.
    Contains(word string) bool
    // SearchByPrefix finds and returns words by prefix.
    SearchByPrefix(prefix string) []string
    // Size returns the number of keys in the tree.
    Size() int
}

type runeNode struct {
    children map[rune]*runeNode
    last     bool
}

// runeTrie is a rune-wise trie implementation.
type runeTrie struct {
    root *runeNode
    size int
}

// NewRuneTrie creates a rune-wise new trie.
func NewRuneTrie() Trie {
    return &runeTrie{root: &runeNode{make(map[rune]*runeNode), false}}
}

// Insert inserts a word into the trie and returns true
// if the word already exists.
func (t *runeTrie) Insert(word string) bool {
    exists := true
    current := t.root
    for _, letter := range word {
        n, ok := current.children[letter]
        if !ok {
            n = &runeNode{make(map[rune]*runeNode), false}
            current.children[letter] = n
        }
        
        current = n
    }
    
    if !current.last {
        t.size++
        exists = false
    }
    current.last = true
    
    return exists
}

// nodeByPrefix returns the node which matches the prefix word and the last letter of prefix,
// if prefix matches no trie tree path then return nil, 0
func (t *runeTrie) nodeByPrefix(prefix string) (*runeNode, rune) {
    current := t.root
    var r rune
    for _, letter := range prefix {
        n, ok := current.children[letter]
        if !ok {
            return nil, 0
        }
        
        current = n
        r = letter
    }
    
    return current, r
}

// search
// inputs: node and the rune in the node and a prefix
// search all paths from the currentNode to leaf nodes.
// returns the paths and add the input prefix in front.
func search(currentNode *runeNode, currentRune rune, prefix []rune) []string {
    var words []string
    if currentNode == nil {
        return words
    }
    
    newPrefix := append(prefix, currentRune)
    if currentNode.last {
        words = append(words, string(newPrefix))
    }
    
    for letter, node := range currentNode.children {
        newWords := search(node, letter, newPrefix)
        words = append(words, newWords...)
    }
    
    return words
}

// Contains returns true if an exact match of the word is found, otherwise false.
func (t *runeTrie) Contains(word string) bool {
    n, _ := t.nodeByPrefix(word)
    
    return n != nil && n.last
}

// SearchByPrefix finds and returns all words by prefix.
func (t *runeTrie) SearchByPrefix(prefix string) []string {
    node, r := t.nodeByPrefix(prefix)
    
    return search(node, r, []rune(trimLastChar(prefix)))
}

// Size returns the size of distinct words in the trie tree.
func (t *runeTrie) Size() int {
    return t.size
}

func trimLastChar(s string) string {
    r, size := utf8.DecodeLastRuneInString(s)
    if r == utf8.RuneError && (size == 0 || size == 1) {
        size = 0
    }
    return s[:len(s)-size]
}
