## read.md

### Project: go-reloaded

Welcome to the go-reloaded project! This is my first Project and I'm thrilled to be part of this journey. This project gives us an incredible opportunity to showcase our Go programming skills and take them to the next level.

### Objectives

The primary objective of this project is to develop a powerful text completion, editing, and auto-correction tool using Go. We'll leverage our existing functions from our previous repository to create this tool and demonstrate our expertise.


### Introduction

Our project will be implemented in Go, adhering to the best practices and coding conventions of the language. Let's make sure our code is well-structured and maintainable.

The text completion, editing, and auto-correction tool we'll build accepts two arguments: the name of an input file containing text that requires modifications and the name of an output file where the modified text will be saved. Our program will perform the following modifications:

- Whenever we encounter (hex), we'll replace the preceding word with its decimal equivalent. We can assume that the word will always be a hexadecimal number. For example, if the text is "1E (hex) files were added," we should transform it to "30 files were added."

- Similarly, whenever we encounter (bin), we'll replace the preceding word with its decimal equivalent. Here also, we can assume that the word will always be a binary number. For instance, if the text is "It has been 10 (bin) years," we should transform it to "It has been 2 years."

- To convert the preceding word to uppercase, we'll use (up). If we encounter this tag, we'll modify the word accordingly. For example, "Ready, set, go (up)!" should become "Ready, set, GO!"

- Conversely, to convert the preceding word to lowercase, we'll use (low). This tag allows us to transform the word accordingly. For instance, "I should stop SHOUTING (low)" should be converted to "I should stop shouting."

- If we want to capitalize the preceding word, we'll use (cap). This tag will help us achieve the desired capitalization. For example, "Welcome to the Brooklyn bridge (cap)" should be changed to "Welcome to the Brooklyn Bridge."

- In addition to (low), (up), and (cap), we can specify a number after the comma to indicate the number of words to be converted. For example, (up, 2) should convert the two preceding words to uppercase. So, if the text is "This is so exciting (up, 2)," it should be transformed to "This is SO EXCITING."

- We should ensure that punctuation marks such as ., ,, !, ?, :, and ; are placed immediately after the preceding word, followed by a space before the next word. For example, "I was sitting over there, and then BAMM!!" should be corrected to "I was sitting over there, and then BAMM!!".

- However, if we encounter groups of punctuation marks like ... or !?, we should maintain the original formatting. For instance, "I was thinking ... You were right" should remain as "I was thinking... You were right."

- The punctuation mark ' should always appear with another instance of ' and be placed directly adjacent to the enclosed word, without any spaces. For example, "I am exactly how they describe me: 'awesome'" should be written as "I am exactly how they describe me: 'awesome'".

- If there are multiple words between the two ' marks, we should place the marks next to the corresponding words. For example, "As Elton John said: 'I am the most well-known homosexual in the world'" should be formatted as "As Elton John said: 'I am the most well-known homosexual in the world'".

- Lastly, we'll make the necessary changes to transform instances of the indefinite article "a" into "an" if the following word begins with a vowel (a, e, i, o, u) or an "h." For example, "There it was. A amazing rock!" should be converted to "There it was. An amazing rock!"

## USAGE

Write your text on the Sample.txt file then write the below command in the terminal:

```
go run . Sample.txt result.txt

```

The reasult will show up on the result.txt file :)

## AUTHORS

* Fatema Alawadhi
