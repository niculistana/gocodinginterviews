package main
import "fmt"

func main() {
    var num_levels int
    
    fmt.Scanf("%v", &num_levels)
    
    current_level := 0
    
    for current_level < num_levels {
        current_level++
        
        pounds := current_level
        spaces := num_levels-current_level
        
        for spaces > 0 {
            fmt.Print(" ")
            spaces--
        }
        
        for pounds > 0 {
            fmt.Print("#")
            pounds--
        }
        fmt.Println()
    }
}