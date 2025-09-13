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

type printFunction func(a ...any) (n int, err error)

var red = color.New(color.FgRed).SprintFunc()

// Utils is the type used to instantiate this module. Any variable of this type will have access
// to all the methods with the reciever *Utils
type Utils struct {
	SpaceBeforeText   bool
	PlaySoundFunc     func(string)
	ValidInputSound   string
	InvalidInputSound string
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

// Display string s with printFunction pF using Utils SpaceBeforeText settings
func (u *Utils) Dialogue(s string, pF printFunction) {

	if u.SpaceBeforeText && s[0] != ' ' {
		s = " " + s
	}

	pF(s)

}

// Display string q on the screen and get a true or false answer
func (u *Utils) GetYesOrNo(q string) bool {
	reader := bufio.NewReader(os.Stdin)

	answerLineMsg := "-> "
	invalidInputMsg := red("Invalid input!")

	for {
		u.Dialogue(q, fmt.Println)
		u.Dialogue(answerLineMsg, fmt.Print)

		userInput, err := reader.ReadString('\n')
		if err != nil {
			errStr := fmt.Sprintf("Error reading input: %v", err)
			u.Dialogue(errStr, fmt.Println)
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
			u.Dialogue(invalidInputMsg, fmt.Println)
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
	invalidInputMsg := red("Please enter a whole number")

	for {
		u.Dialogue(q, fmt.Println)
		u.Dialogue(answerLineMsg, fmt.Print)

		userInput, err := reader.ReadString('\n')
		if err != nil {
			errStr := fmt.Sprintf("Error reading input: %v", err)
			u.Dialogue(errStr, fmt.Println)
			continue
		}
		userInput = strings.TrimSpace(userInput)

		num, err := strconv.Atoi(userInput)
		if err != nil {
			u.Dialogue(invalidInputMsg, fmt.Println)
			if u.PlaySoundFunc != nil && u.InvalidInputSound != "" {
				u.PlaySoundFunc(u.InvalidInputSound)
			}
			continue
		}

		if u.PlaySoundFunc != nil && u.ValidInputSound != "" {
			u.PlaySoundFunc(u.ValidInputSound)
		}
		return num
	}
}

// Display string q on the screen and get a string of minimum minLength characters
func (u *Utils) GetString(q string, minLength int) string {
	reader := bufio.NewReader(os.Stdin)

	answerLineMsg := "-> "
	invalidWordMsg := red("The word can't have any digit!")
	invalidLengthMsg := red("The word must contain atleast %d letters!", minLength)

	for {
		u.Dialogue(q, fmt.Println)
		u.Dialogue(answerLineMsg, fmt.Print)

		userInput, err := reader.ReadString('\n')
		if err != nil {
			errStr := fmt.Sprintf("Error reading input: %v", err)
			u.Dialogue(errStr, fmt.Println)
			continue
		}
		userInput = strings.TrimSpace(userInput)

		if u.HasDigit(userInput) != -1 {
			u.Dialogue(invalidWordMsg, fmt.Println)
			if u.PlaySoundFunc != nil && u.InvalidInputSound != "" {
				u.PlaySoundFunc(u.InvalidInputSound)
			}
			continue
		}

		if minLength > 0 && len(userInput) < minLength {
			u.Dialogue(invalidLengthMsg, fmt.Println)
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
