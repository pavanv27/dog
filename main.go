/*Package dog is a temporary package created to compete the ninja level 12 task*/
package main

import "fmt"

func main() {
	var hy int
	fmt.Println("Enter Human years to convert to Dog years :")
	fmt.Scanf("%d", &hy)
	fmt.Println("The dog years corresponding to human years", hy, "is", Years(hy))
}

//Years function converts Human years to Dog years
func Years(hy int) int {
	return hy * 7
}
