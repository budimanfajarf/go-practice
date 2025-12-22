// https://youtu.be/Dd_NgYVOdLk?si=ld32mWM4eCUB_xhL
// https://gist.github.com/JyotinderSingh/d2bd0096e146aa3083442ceb48eab6b4
package edit_distance

func MinDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word2)+1)
	for i := range dp {
		dp[i] = make([]int, len(word1)+1)
	}

	// fmt.Println("dp #1", dp)

	/*
	 * Going downwards in the first column (for "" as the first string),
	 * we're performing the insert operation
	 * And the number of inserts to turn an empty string to a
	 * string with length L is L
	 */
	for i := range dp {
		dp[i][0] = i
	}

	// fmt.Println("dp #2", dp)

	/*
	 * Going towards the right in the first row (for "" as the second string)
	 * we're performing the delete operation
	 * And the number of deletes to turn a string of length L into an empty
	 * string "" is L.
	 */

	for i := 0; i < len(dp[0]); i++ {
		dp[0][i] = i
	}

	// fmt.Println("dp #3", dp)

	for row := 1; row < len(dp); row++ {
		for col := 1; col < len(dp[0]); col++ {
			// fmt.Print("compare ", string(word1[col-1]), " & ", string(word2[row-1]))
			if word1[col-1] == word2[row-1] {
				/*
				 * If the characters match, we remove the character from
				 * both the strings, and find the answer to the subproblem
				 * formed by the remaining words (which is located at [row - 1][col - 1])
				 */
				dp[row][col] = dp[row-1][col-1]

				// fmt.Print(" => dp[", row, "][", col, "] = ", dp[row][col])
				// fmt.Print("; match, dp[", row-1, "][", col-1, "] = ", dp[row-1][col-1], "\n")
			} else {
				/*
				 * In case of a mismatch, we see which operation out of
				 * replacement, insertion, and deletion takes the minimum steps
				 * to convert the word1 to word2.
				 */
				dp[row][col] = min(dp[row-1][col-1], dp[row-1][col], dp[row][col-1]) + 1

				// fmt.Print(" => dp[", row, "][", col, "] = ", dp[row][col])
				// fmt.Print("; mismatch, 1 + min(replace dp[", row-1, "][", col-1, "] = ", dp[row-1][col-1])
				// fmt.Print(", insert dp[", row-1, "][", col, "] = ", dp[row-1][col])
				// fmt.Print(", delete dp[", row, "][", col-1, "] = ", dp[row][col-1], ")\n")
			}
		}
	}

	// fmt.Println("dp #4", dp)

	// The last element of the matrix contains the edit distance for the original problem.
	return dp[len(dp)-1][len(dp[0])-1]
}
