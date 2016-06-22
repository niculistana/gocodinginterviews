package main
import (
	"fmt"
)

func main() {
	var dict_count int
	var dict_item string

	var query_count int
	var query_item string
	
	fmt.Scanf("%v", &dict_count)
	dict_items := []string{}	

	for dict_count > 0 {
		fmt.Scanf("%s", &dict_item)
		dict_items = append(dict_items, dict_item)
		dict_count--
	}
	
	fmt.Scanf("%v", &query_count)
	query_items := []string{}	

	for query_count > 0 {
		fmt.Scanf("%s", &query_item)
		query_items = append(query_items, query_item)
		query_count--
	}

	var match_count int

	if (query_count <= dict_count) {
		for i:=0;i<len(query_items);i++ {
			for j:=0;j<len(dict_items);j++ {
				if query_items[i] == dict_items[j] {
					match_count++					
				}
			}
			fmt.Println(match_count)
			match_count = 0
		}
	}
}