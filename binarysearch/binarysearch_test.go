 package main

 import (
    "testing"
 )

 func TestBinarySearch(t *testing.T) {
     array := []int{1, 3, 4, 5, 6}
     index , err := BinarySearch(array,3)
     if err != nil {
         t.Error(err)
     }
     t.Log(index)
 }
