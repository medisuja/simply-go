package main
 
import (
    "fmt"
    "strings"
)

func wordCount(str string) map[string]int {

    charList := strings.Split(str, "")

    counts := make(map[string]int)

    for _, chr := range charList {
        _, ok := counts[chr]
        if ok {
            counts[chr] += 1
        } else {
            counts[chr] = 1
        }
    }

    return counts
}
 
func main() {
    strLine := "LORD"

    for index,element := range wordCount(strLine){
        fmt.Println(index,"=",element)
    }
}
