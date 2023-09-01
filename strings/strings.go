package strings

import (
    "strings"
)

// Indexes return all the indexes occurs in the s
// if none sub occur in the s returns nil
func Indexes(s string, substr string) []int {
    var indexes []int
    cursor := 0
    for {
        i := strings.Index(s, substr)
        if i == -1 {
            break
        } else {
            indexes = append(indexes, i+cursor)
        }
        
        s = s[i+1:]
        cursor += i + 1
    }
    
    if len(indexes) <= 0 {
        return nil
    }
    
    return indexes
}
