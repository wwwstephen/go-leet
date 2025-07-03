package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
        return strs[0]
    }
    // Step 1 sort strings
    sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) < len(strs[j])
	})

	// Step 2: Get shortest string as reference
	shortest := strs[0]

	// Step 3: Compare characters across all strings
	prefix := ""
	for i := 0; i < len(shortest); i++ {
		char := shortest[i]
		for _, word := range strs {
			if word[i] != char {
				fmt.Println("Longest common prefix:", prefix)
				return prefix
			}
		}
		prefix += string(char)
	}
    return prefix
}
