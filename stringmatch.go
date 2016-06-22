package main
import (
	"fmt"
)

func main() {
	var num_queries int
	var query string
	var cur_query string
	var cur_count int
	
	fmt.Scanf("%v", &num_queries)
	fmt.Scanf("%v", &query)

	for num_queries > 0 {
		fmt.Scanf("%v", &cur_query)
		for i:=0; i < len(cur_query);i++ {
			if cur_query[i] == query[0] && (len(cur_query) - i) >= len(query) {
				char_count,j:=0,i

				for k:=0; k < len(query); k++ {
					if cur_query[j+k] == query[k] {
						char_count++
					}
				}
				
				if char_count == len(query) {
					cur_count++
					break
				}
			}
		}
		num_queries--
	}

	fmt.Println(cur_count)
}