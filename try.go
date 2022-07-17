package main

import "strings"

func main() {
	tests := []struct {
		name     string
		choice1  string
		choice2  string
		expected string
	}{
		{
			name:     "chooses Bugatti over Ford",
			choice1:  "Bugatti Veyron",
			choice2:  "Ford Pinto",
			expected: "Bugatti Veyron is clearly the better choice.",
		}, {
			name:     "chooses Chery over Kia",
			choice1:  "Chery EQ",
			choice2:  "Kia Niro Elektro",
			expected: "Chery EQ is clearly the better choice.",
		}, {
			name:     "chooses Ford Focus over Ford Pinto",
			choice1:  "Ford Focus",
			choice2:  "Ford Pinto",
			expected: "Ford Focus is clearly the better choice.",
		}, {
			name:     "chooses 2018 over 2020",
			choice1:  "2018 Bergamont City",
			choice2:  "2020 Gazelle Medeo",
			expected: "2018 Bergamont City is clearly the better choice.",
		}, {
			name:     "chooses Bugatti over ford",
			choice1:  "Bugatti Veyron",
			choice2:  "ford",
			expected: "Bugatti Veyron is clearly the better choice.",
		},
	}

	for _, test := range tests {
		if strings.Compare(test.choice1, test.choice2) == 0 {
			println("The choice is: " + test.choice1)
		} else if strings.Compare(test.choice1, test.choice2) == -1 {
			println("The choice is: " + test.choice1)
		} else {
			println("The choice is: " + test.choice2)
		}
	}
}
