// https://www.codechef.com/problems/PLAYPIAN
package main

import "fmt"

func main() {
    var numOfCases int
    var data string

    _, _ = fmt.Scan(&numOfCases)

    for i := 0; i < numOfCases; i++ {
        _, _ = fmt.Scanln(&data)

        tampered := false
        for j := 0; j < len(data); j += 2 {
            if data[j] == data[j+1] {
                tampered = true
                break
            }
        }

        if tampered {
            fmt.Println("no")
        } else {
            fmt.Println("yes")
        }
    }
}
