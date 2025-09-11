package tgutils

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"unicode"

	"github.com/fatih/color"
)

// Utils is the type used to instantiate this module. Any variable of this type will have access
// to all the methods with the reciever *Utils
type Utils struct {
	SpaceBeforeText   bool
	ColoredText       bool
	PlaySoundFunc     func(string)
	ValidInputSound   string
	InvalidInputSound string
}

// textConfig unites all the text messages of a get method into one struct
type textConfig struct {
	answerLineText    string
	invalidInputText  string
	invalidLengthText string
	invalidWordText   string
}

// Get a random non-negative number within the range [min, max)
func (u *Utils) RangedRandom(min, max int) int {
	return rand.Intn(max-min) + min
}

// Check if element of type string is in slice of type []string
// Returns the index if found or -1 if not found
func (u *Utils) StrInSlice(slice []string, element string) int {
	for i, s := range slice {
		if s == element {
			return i
		}
	}
	return -1
}

// Display a slice in a neat way
// comaMode will also add comas
func (u *Utils) DisplaySlice(sl interface{}, commaMode bool) {
	v := reflect.ValueOf(sl)

	if v.Kind() != reflect.Slice {
		fmt.Println("DisplaySlice can only display slices!")
		return
	}

	fmt.Print(" ")
	for i := 0; i < v.Len(); i++ {
		fmt.Printf("%v", v.Index(i).Interface())
		if i != v.Len()-1 && commaMode {
			fmt.Print(", ")
		}
	}
	fmt.Println()
}

// Check if a string has a digit in it
// Returns the index of the first digit, or -1 if there are no digits
func (u *Utils) HasDigit(s string) int {
	for i, c := range s {
		if unicode.IsDigit(c) {
			return i
		}
	}
	return -1
}

// Clear the terminal screen
func (u *Utils) ClearScreen() {
	if strings.Contains(runtime.GOOS, "windows") {
		// windows
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		// linux or mac
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

// Display string q on the screen and get a true or false answer
func (u *Utils) GetYesOrNo(q string) bool {
	reader := bufio.NewReader(os.Stdin)

	answerLineMsg := "-> "
	invalidInputMsg := "Invalid input!"

	tc := textConfig{
		answerLineText:   u.getMethodHelper(answerLineMsg, false),
		invalidInputText: u.getMethodHelper(invalidInputMsg, true),
	}

	for {
		fmt.Println(q)
		fmt.Print(tc.answerLineText)

		userInput, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		userInput = strings.TrimSpace(userInput)
		userInput = strings.ToLower(userInput)

		switch userInput {
		case "y":
			if u.PlaySoundFunc != nil && u.ValidInputSound != "" {
				u.PlaySoundFunc(u.ValidInputSound)
			}
			return true
		case "n":
			if u.PlaySoundFunc != nil && u.ValidInputSound != "" {
				u.PlaySoundFunc(u.ValidInputSound)
			}
			return false
		default:
			fmt.Println(tc.invalidInputText)
			if u.PlaySoundFunc != nil && u.InvalidInputSound != "" {
				u.PlaySoundFunc(u.InvalidInputSound)
			}
		}
	}

}

// Display string q on the screen and get a number of type int
func (u *Utils) GetNumber(q string) int {
	reader := bufio.NewReader(os.Stdin)

	answerLineMsg := "-> "
	invalidInputMsg := "Please enter a whole number"

	tc := textConfig{
		answerLineText:   u.getMethodHelper(answerLineMsg, false),
		invalidInputText: u.getMethodHelper(invalidInputMsg, true),
	}

	for {
		fmt.Println(q)
		fmt.Print(tc.answerLineText)

		userInput, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		userInput = strings.TrimSpace(userInput)

		num, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println(tc.invalidInputText)
			if u.PlaySoundFunc != nil && u.InvalidInputSound != "" {
				u.PlaySoundFunc(u.InvalidInputSound)
			}
			continue
		} else {
			if u.PlaySoundFunc != nil && u.ValidInputSound != "" {
				u.PlaySoundFunc(u.ValidInputSound)
			}
			return num
		}
	}
}

// Display string q on the screen and get a string of minimum minLength characters
func (u *Utils) GetString(q string, minLength int) string {
	reader := bufio.NewReader(os.Stdin)

	answerLineMsg := "-> "
	invalidWordMsg := "The word can't have any digit!"
	invalidLengthMsg := fmt.Sprintf("The word must contain atleast %d letters!", minLength)

	tc := textConfig{
		answerLineText:    u.getMethodHelper(answerLineMsg, false),
		invalidWordText:   u.getMethodHelper(invalidWordMsg, true),
		invalidLengthText: u.getMethodHelper(invalidLengthMsg, true),
	}

	for {
		fmt.Println(q)
		fmt.Print(tc.answerLineText)

		userInput, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		userInput = strings.TrimSpace(userInput)

		if u.HasDigit(userInput) != -1 {
			fmt.Println(tc.invalidWordText)
			if u.PlaySoundFunc != nil && u.InvalidInputSound != "" {
				u.PlaySoundFunc(u.InvalidInputSound)
			}
			continue
		}

		if minLength > 0 && len(userInput) < minLength {
			fmt.Println(tc.invalidLengthText)
			if u.PlaySoundFunc != nil && u.InvalidInputSound != "" {
				u.PlaySoundFunc(u.InvalidInputSound)
			}
			continue
		}

		if u.PlaySoundFunc != nil && u.ValidInputSound != "" {
			u.PlaySoundFunc(u.ValidInputSound)
		}
		return userInput
	}
}

// Helper functions for Get methods
func (u *Utils) getMethodHelper(str string, colorable bool) string {
	var red = color.New(color.FgRed).SprintFunc()
	if str == "" {
		return str
	}

	if u.SpaceBeforeText && str[0] != ' ' {
		str = " " + str
	}
	if u.ColoredText && colorable {
		str = red(str)
	}
	return str
}
