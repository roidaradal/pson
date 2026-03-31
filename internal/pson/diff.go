package pson

import (
	"fmt"

	"github.com/zeroibot/fn/lang"
	"github.com/zeroibot/fn/str"
)

// Compare each line and find differences
func rawDiff(lines1, lines2 []string) error {
	count1, count2 := len(lines1), len(lines2)
	maxLines := max(count1, count2)
	diffCount := 0
	for i := range maxLines {
		if i >= count1 {
			diffCount += count2 - i
			extra2 := str.Red(fmt.Sprintf("%d extra lines in file2:", count2-i))
			fmt.Printf("%s line %d - %d\n", extra2, i+1, count2)
			break
		}
		if i >= count2 {
			diffCount += count1 - i
			extra1 := str.Red(fmt.Sprintf("%d extra lines in file1:", count1-i))
			fmt.Printf("%s line %d - %d\n", extra1, i+1, count1)
			break
		}
		if lines1[i] == lines2[i] {
			continue // skip if same line
		}
		diffCount += 1
		label := fmt.Sprintf("Line %d:", i+1)
		fmt.Printf("[%d] %s \n", diffCount, str.Red(label))
		line1, line2 := getLineDiff(lines1[i], lines2[i])
		fmt.Printf("\t- %s\n", line1)
		fmt.Printf("\t- %s\n", line2)
	}
	fmt.Printf("%sRawDiff: %d", lang.Ternary(diffCount > 0, "\n", ""), diffCount)
	return nil
}

// Get line difference
func getLineDiff(line1, line2 string) (string, string) {
	previewLength := 10
	len1, len2 := len(line1), len(line2)

	for i := range max(len1, len2) {
		col := i + 1
		limit1 := min(i+previewLength, len1)
		limit2 := min(i+previewLength, len2)
		tail1, tail2 := "", ""
		valid := true

		if i >= len1 {
			valid = false
			tail2 = line2[i:limit2]
		}
		if i >= len2 {
			valid = false
			tail1 = line1[i:limit1]
		}
		if valid {
			if line1[i] == line2[i] {
				continue // skip if equal
			}
			tail1 = line1[i:limit1]
			tail2 = line2[i:limit2]
		}

		tail1 = fmt.Sprintf("Col %d: %s", col, tail1)
		tail2 = fmt.Sprintf("Col %d: %s", col, tail2)
		return tail1, tail2
	}
	return "", ""
}
