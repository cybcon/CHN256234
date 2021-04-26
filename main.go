// Solution idea from: https://stackoverflow.com/questions/55201821/merging-overlapping-intervals-using-double-for-loop
package main

import (
    "fmt"
    "sort"
)

type Interval struct {
    Lo, Hi int
}

func merge(ivs []Interval) []Interval {
    m := append([]Interval(nil), ivs...)
    if len(m) <= 1 {
        return m
    }

    sort.Slice(m,
        func(i, j int) bool {
            if m[i].Lo < m[j].Lo {
                return true
            }
            if m[i].Lo == m[j].Lo && m[i].Hi < m[j].Hi {
                return true
            }
            return false
        },
    )

    j := 0
    for i := 1; i < len(m); i++ {
        if m[j].Hi >= m[i].Lo {
            if m[j].Hi < m[i].Hi {
                m[j].Hi = m[i].Hi
            }
        } else {
            j++
            m[j] = m[i]
        }

    }
    return append([]Interval(nil), m[:j+1]...)
}