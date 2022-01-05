package main 

import (
    "testing"
    "sort"
    "fmt"
)

type SumTest struct {
    testValues []int
    expectedResult int
}

func TestSum(t *testing.T) {
    testVals := []SumTest {
        { testValues: []int{10, 20, 30}, expectedResult:  10},
        { testValues: []int{ -10, 0, -10 }, expectedResult:  -20},
        { testValues: []int{ -10, 0, -10 }, expectedResult:  -20},
    }
    for index, testVal := range testVals {
        t.Run(fmt.Sprintf("Sum #%v", index), func(subT *testing.T) {
            if (t.Failed()) {
                subT.SkipNow()
            }
            _, sum := sortAndTotal(testVal.testValues)
            if (sum != testVal.expectedResult) {
                subT.Fatalf("Expected %v, Got %v", testVal.expectedResult, sum)
            }
        })
    }
}

func TestSort(t *testing.T) {
    slices := [][]int {
        { 1, 279, 48, 12, 3 },
        { -10, 0, -10 },
        { 1, 2, 3, 4, 5, 6, 7 },
        { 1 },
    }
    for index, data := range slices {
        t.Run(fmt.Sprintf("Sort #%v", index), func(subT *testing.T) {
            sorted, _ := sortAndTotal(data)
            if (!sort.IntsAreSorted(sorted)) {
                subT.Fatalf("Unsorted data %v", sorted)
            }
        })
    }    
}
