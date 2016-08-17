package main

import "fmt"

func main() {
		/*
		fmt - format package
		PringT fprint format where Pringl prints line
		%d for dicimal format
		%b for binary format
		%x for hexadecimal fromat

		42 - 101010 - 2a
		42 - 101010 - 0x2a
		42 - 101010 - 0X2A
		42       101010          0X2A

		*/
		fmt.Printf("%d - %b - %x \n", 42, 42, 42)
		fmt.Printf("%d - %b - %#x \n", 42, 42, 42)
		fmt.Printf("%d - %b - %#X \n", 42, 42, 42)
		fmt.Printf("%d \t %b \t %#X \n", 42, 42, 42)
}
