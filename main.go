package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

var MEMORY = []int{0}


func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	now := time.Now()
	filename := os.Args[1]
	content_bytes, err := ioutil.ReadFile(filename)
	content := string(content_bytes)
	check(err)
	fmt.Print(content)
	ptr := 0
	fmt.Println(time.Since(now), "Start")
	run(&ptr, &content)
	fmt.Println(time.Since(now))
}

func run(ptr *int, content *string) {
	//fmt.Println(*content)
	jump := -1

	for i := 0; i < len(*content); i++ {
		if jump != -1 && i < jump {
			continue
		}

		current_char := (*content)[i]
		if current_char == '+' {
			MEMORY[*ptr] += 1
		} else if current_char == '-' {
			MEMORY[*ptr] -= 1
		} else if current_char == '>' {
			*ptr += 1
			if len(MEMORY) - 1 < * ptr {
				MEMORY = append(MEMORY, 0)
			}
		} else if current_char == '<' {
			*ptr -= 1
		} else if current_char == '.' {
			fmt.Print(string(MEMORY[*ptr]))
		} else if current_char == '[' {
			skiploop := 1
			end_index := i + 1
			for MEMORY[*ptr] > 0 {
				for skiploop != 0 {
					if (*content)[end_index] == '[' {
						skiploop += 1
					} else if (*content)[end_index] == ']' {
						skiploop -= 1
					}
					end_index += 1
				}
				new_content := string((*content)[i +1 : end_index - 1])
				run(ptr, &new_content)
			}
			jump = end_index - 1
		}
	}

}
