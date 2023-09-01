package strings

import (
    "reflect"
    "testing"
)

func TestIndexes(t *testing.T) {
    type args struct {
        s      string
        substr string
    }
    tests := []struct {
        name string
        args args
        want []int
    }{
        {
            args: args{
                s:      "aabbaaccaa",
                substr: "aa",
            },
            want: []int{0, 4, 8},
        },
        {
            args: args{
                s:      "aaabbbaa",
                substr: "aa",
            },
            want: []int{0, 1, 6},
        },
        {
            args: args{
                s:      "周杰伦你好",
                substr: "杰伦",
            },
            want: []int{1},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := Indexes(tt.args.s, tt.args.substr); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("Indexes() = %v, want %v", got, tt.want)
            }
        })
    }
}
