/*
Problem Statement

Shashank likes strings in which consecutive characters are different. For example, he likes ABABA, while he doesn't like ABAA. Given a string containing characters A and B only, he wants to change it into a string he likes. To do this, he is allowed to delete the characters in the string.

Your task is to find the minimum number of required deletions.

Input Format
The first line contains an integer T i.e. the number of test cases.
Next T lines contain a string each.

Output Format
Print minimum number of required steps for each test case.

Constraints

1≤T≤10
1≤lengthofString≤105

Sample Input

5
AAAA
BBBBB
ABABABAB
BABABA
AAABBB

Sample Output

3
4
0
0
4

Explanation

AAAA⟹A, 3 deletions
BBBBB⟹B, 4 deletions
ABABABAB⟹ABABABAB, 0 deletions
BABABA⟹BABABA, 0 deletions
AAABBB⟹AB, 4 deletions
*/
package main

import "fmt"

func main() {
	var n, mind int
	var s string

	fmt.Scanf("%v", &n)

	for n > 0 {
		fmt.Scanf("%s", &s)

		mind = 0

		for i:=1; i<len(s); i++ {
			if s[i] == s[i-1] {
				mind++
			}
		}

		fmt.Println(mind)
		n--
	}	
}