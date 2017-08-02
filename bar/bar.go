package bar

import "fmt"

func init() {
        fmt.Println("bar.  <-- here is issue")
}

func Bar() {
        fmt.Println("bar Bar.")
}
