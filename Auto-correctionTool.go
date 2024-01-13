package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Punctuations(Text string) string {
	Letters := []rune(Text)
	var FormatedText []rune
	for a := 0; a < len(Letters)-1; a++ {
		if (Letters[a] == '.' || Letters[a] == ',' || Letters[a] == '!' || Letters[a] == '?' || Letters[a] == ':' || Letters[a] == ';') && Letters[a+1] != ' ' {
			FormatedText = append(FormatedText, Letters[a])
			if Letters[a+1] == '.' || Letters[a+1] == ',' || Letters[a+1] == '!' || Letters[a+1] == '?' || Letters[a+1] == ':' || Letters[a+1] == ';' {
				b := a + 1
				for Letters[b] == '.' || Letters[b] == ',' || Letters[b] == '!' || Letters[b] == '?' || Letters[b] == ':' || Letters[b] == ';' {
					if b == len(Letters)-1 {
						break
					}
					FormatedText = append(FormatedText, Letters[b])
					b++
				}
				a = b - 1
			}
			if Letters[a+1] == ' ' {
				continue
			}
			FormatedText = append(FormatedText, ' ')
		} else if Letters[a] != ' ' && (Letters[a+1] == '.' || Letters[a+1] == ',' || Letters[a+1] == '!' || Letters[a+1] == '?' || Letters[a+1] == ':' || Letters[a+1] == ';') {
			word := string([]rune{Letters[a-3], Letters[a-2], Letters[a-1], Letters[a]})
			if word == " (up" || word == "(low" || word == "(cap" {
				FormatedText = append(FormatedText, Letters[a])
			} else {
				FormatedText = append(FormatedText, Letters[a])
				FormatedText = append(FormatedText, ' ')
			}
		} else if Letters[a] == '\'' {
			if a == 0 {
				FormatedText = append(FormatedText, '\'')
				c := a + 1
				for Letters[c] != '\'' {
					if FormatedText[len(FormatedText)-1] == '\'' && Letters[c] == ' ' {
						c++
						continue
					} else if Letters[c] == ' ' && Letters[c+1] != '\'' {
						FormatedText = append(FormatedText, Letters[c])
						c++
					} else if Letters[c] == ' ' && Letters[c+1] == '\'' {
						c++
						continue
					} else {
						FormatedText = append(FormatedText, Letters[c])
						c++
					}
				}
				a = c
				FormatedText = append(FormatedText, '\'')
				if a != (len(Letters) - 1) {
					if Letters[a+2] != ' ' {
						FormatedText = append(FormatedText, ' ')
					}
				}
			}else if Letters[a-1] == ' ' {
				FormatedText = append(FormatedText, '\'')
				c := a + 1
			    for Letters[c] != '\'' {
				    if FormatedText[len(FormatedText)-1] == '\'' && Letters[c] == ' ' {
						c++
						continue
					} else if Letters[c] == ' ' && Letters[c+1] != '\'' {
						FormatedText = append(FormatedText, Letters[c])
						c++
					} else if Letters[c] == ' ' && Letters[c+1] == '\'' {
						c++
						continue
					} else {
						FormatedText = append(FormatedText, Letters[c])
						c++
					}
				}
				a = c
				FormatedText = append(FormatedText, '\'')
				if a != (len(Letters) - 1) {
					if Letters[a+2] != ' ' {
						FormatedText = append(FormatedText, ' ')
					}
				}
			}
		} else {
			FormatedText = append(FormatedText, Letters[a])
		}
	}
	if Letters[(len(Letters)-1)] != '\'' {
		FormatedText = append(FormatedText, Letters[(len(Letters)-1)])
	}
	return string(FormatedText)
}

func FinalFormat(Text string) string {
	Letters := []rune(Text)
	var FormatedText []rune
	for a := 0; a < len(Letters)-1; a++ {
		if Letters[a] == ' ' && (Letters[a+1] == '.' || Letters[a+1] == ',' || Letters[a+1] == '!' || Letters[a+1] == '?' || Letters[a+1] == ':' || Letters[a+1] == ';') {
			continue
		} else if Letters[a] <= '9' && Letters[a] >= '0' && Letters[a+1] == ')' {
			if Letters[a-1] != ' ' {
				FormatedText = DeleteChar(FormatedText, a-1)
			} else {
				if a+1 == len(Letters)-1 {
					return string(FormatedText)
				} else {
					a++
					continue
				}
			}
		} else if Letters[a] == ' ' && Letters[a+1] == ' ' {
			continue
		} else {
			FormatedText = append(FormatedText, Letters[a])
		}
	}
	FormatedText = append(FormatedText, Letters[(len(Letters)-1)])
	return string(FormatedText)
}

func ConvertHex(Hex string) string {
	Dec, err := strconv.ParseInt(Hex, 16, 64)
	if err != nil {
		fmt.Println("Error: Please enter a valid number before hex")
		return "Error"
	} else {
		DecStr := strconv.FormatInt(Dec, 10)
		return DecStr
	}
}

func DeleteString(array []string, index int) []string {
	FinalArray := array[0:index]
	for a := index + 1; a < len(array); a++ {
		FinalArray = append(FinalArray, array[a])
	}
	return FinalArray
}

func DeleteChar(array []rune, index int) []rune {
	FinalArray := array[0:index]
	for a := index + 1; a < len(array); a++ {
		FinalArray = append(FinalArray, array[a])
	}
	return FinalArray
}

func ConvertBin(Bin string) string {
	Dec, err := strconv.ParseInt(Bin, 2, 64)
	if err != nil {
		fmt.Println("Error: Please enter a valid number before bin")
		return "Error"
	} else {
		DecStr := strconv.FormatInt(Dec, 10)
		return DecStr
	}
}

func UpperCase(word string) string {
	Letters := []rune(word)
	for i := 0; i < len(Letters); i++ {
		if Letters[i] >= 'a' && Letters[i] <= 'z' {
			Letters[i] = Letters[i] - 32
		}
	}
	return string(Letters)
}

func LowerCase(word string) string {
	Letters := []rune(word)
	for i := 0; i < len(Letters); i++ {
		if Letters[i] >= 'A' && Letters[i] <= 'Z' {
			Letters[i] = Letters[i] + 32
		}
	}
	return string(Letters)
}

func capitalize(word string) string {
	Letters := []rune(word)
	if Letters[0] >= 'a' && Letters[0] <= 'z' {
		Letters[0] = Letters[0] - 32
	}
	return string(Letters)
}

func CheckPunctuations(Text string) bool {
	word := []rune(Text)
	for i := 0; i < len(word); i++ {
		if (word[i] >= 'A' && word[i] <= 'Z') || (word[i] >= 'a' && word[i] <= 'z') {
			return false
		}
	}
	return true
}

func main() {
	input := os.Args
	if len(input) < 3 {
		fmt.Println("Please provide two file names")
	} else if len(input) > 3 {
		fmt.Println("Please provide two file names only")
	} else {
		file, err := os.Open(input[1])
		if err != nil {
			fmt.Println("Error while opening file:", err.Error())
		} else {
			scanner := bufio.NewScanner(file)
			var Text string
			for scanner.Scan() {
				Text = scanner.Text() + Text
			}
			file.Close()
			FormatedText := Punctuations(Text)
			words := strings.Split(FormatedText, " ")
			WordsBefore := 0
			for i := 0; i < len(words); i++ {
				if words[i] != "." && words[i] != "," && words[i] != "!" && words[i] != "?" && words[i] != ":" && words[i] != ";" {
					WordsBefore++
				}
				if words[i] == "(hex)" {
					if len(words) == 1 {
						fmt.Println("Please write a valid number")
						words[i] = "Error"
					} else {
						words[i-1] = ConvertHex(words[i-1])
						words = DeleteString(words, i)
					}
					continue
				}
				if words[i] == "(bin)" {
					if len(words) == 1 {
						fmt.Println("Please write a valid number")
						words[i] = "Error"
					} else {
						words[i-1] = ConvertBin(words[i-1])
						words = DeleteString(words, i)
					}
					continue
				}
				if words[i] == "(up)" {
					if len(words) == 1 {
						fmt.Println("Please write a valid text")
						words[i] = "Error"
					} else {
						var check bool
						skip := 0
						for a := 1; a <= (1 + skip); a++ {
							check = CheckPunctuations(words[i-a])
							if check == false {
								words[i-a] = UpperCase(words[i-a])
							} else {
								skip++
							}
						}
						words = DeleteString(words, i)
					}
					continue
				}
				if words[i] == "(low)" {
					if len(words) == 1 {
						fmt.Println("Please write a valid text")
						words[i] = "Error"
					} else {
						var check bool
						skip := 0
						for a := 1; a <= (1 + skip); a++ {
							check = CheckPunctuations(words[i-a])
							if check == false {
								words[i-a] = LowerCase(words[i-a])
							} else {
								skip++
							}
						}
						words = DeleteString(words, i)
					}
					continue
				}
				if words[i] == "(cap)" {
					if len(words) == 1 {
						fmt.Println("Please write a valid text")
						words[i] = "Error"
					} else {
						var check bool
						skip := 0
						for a := 1; a <= (1 + skip); a++ {
							check = CheckPunctuations(words[i-a])
							if check == false {
								words[i-a] = capitalize(words[i-a])
							} else {
								skip++
							}
						}
						words = DeleteString(words, i)
					}
					continue
				}
				if words[i] == "(cap," {
					var check bool
					skip := 0
					numrune := []rune(words[i+1])
					Numplace := 1
					num := 0
					for i := (len(numrune) - 2); i >= 0; i-- {
						num = (int(numrune[i]-48) * Numplace) + num
						Numplace = Numplace * 10
					}
					if num >= WordsBefore {
						fmt.Println("Please ensure that there are as many words as the number mentioned")
						words[i] = "Error"
					} else {
						for a := 1; a <= (num + skip); a++ {
							check = CheckPunctuations(words[i-a])
							if check == false {
								words[i-a] = capitalize(words[i-a])
							} else {
								skip++
							}
						}
						words = DeleteString(words, i)
					}
					continue
				}
				if words[i] == "(up," {
					var check bool
					skip := 0
					numrune := []rune(words[i+1])
					Numplace := 1
					num := 0
					for i := (len(numrune) - 2); i >= 0; i-- {
						num = (int(numrune[i]-48) * Numplace) + num
						Numplace = Numplace * 10
					}
					if num >= WordsBefore {
						fmt.Println("Please ensure that there are as many words as the number mentioned")
						words[i] = "Error"
					} else {
						for a := 1; a <= (num + skip); a++ {
							check = CheckPunctuations(words[i-a])
							if check == false {
								words[i-a] = UpperCase(words[i-a])
							} else {
								skip++
							}
						}
						words = DeleteString(words, i)
					}
					continue
				}
				if words[i] == "(low," {
					var check bool
					skip := 0
					numrune := []rune(words[i+1])
					Numplace := 1
					num := 0
					for i := (len(numrune) - 2); i >= 0; i-- {
						num = (int(numrune[i]-48) * Numplace) + num
						Numplace = Numplace * 10
					}
					if num >= WordsBefore {
						fmt.Println("Please ensure that there are as many words as the number mentioned")
						words[i] = "Error"
					} else {
						for a := 1; a <= (num + skip); a++ {
							check = CheckPunctuations(words[i-a])
							if check == false {
								words[i-a] = LowerCase(words[i-a])
							} else {
								skip++
							}
						}
						words = DeleteString(words, i)
					}
					continue
				}
				if words[i] == "a" || words[i] == "A" {
					if i == len(words)-1 {
						continue
					} else {
						word := []rune(words[i+1])
						if word[0] == 'a' || word[0] == 'e' || word[0] == 'i' || word[0] == 'o' || word[0] == 'u' || word[0] == 'h' {
							if words[i] == "a" {
								words[i] = "an"
							}
							if words[i] == "A" {
								words[i] = "An"
							}
						}
					}
					continue
				}
			}
			FinalText := strings.Join(words, " ")
			FinalText = FinalFormat(FinalText)
			FinalText = FinalFormat(FinalText)
			file, err := os.Create(input[2])
			if err != nil {
				fmt.Println(err)
				return
			}
			defer file.Close()
			writer := bufio.NewWriter(file)

			// Write the text to the file
			_, err = writer.WriteString(FinalText)
			if err != nil {
				fmt.Println(err)
				return
			}

			// Flush the writer to ensure that all data has been written to the file
			err = writer.Flush()
			if err != nil {
				fmt.Println(err)
				return
			}

		}
	}
}
