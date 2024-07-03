## ASCII ART
* ASCII-art is a program written in Go that takes a string as input and outputs a graphic representation of the string using ASCII characters. 

## Features ##
* Input Handling :-  Accepts strings containing numbers, letters, spaces, special characters, and a newline character.
* ASCII Representation:- Converts the input string into a graphical representation using ASCII characters.
* Customizable Templates:- Provides different graphical templates such as "shadow," "standard," and "thinkertoy" for ASCII representation.
* Command-Line Interface-: Accessible via the command line for easy input and output.
* Error Handling:- Implements robust error handling to ensure reliable performance.
* We also included our own bannerfile i.e lean.txt

## Instructions

To clone this repository: use the command below on your terminal:
```
git clone https://learn.zone01kisumu.ke/git/vandisi/ascii-art-fs.git
```

## Banner Format
* Each character in the ASCII representation has a height of 8 lines.
* Characters are separated by a newline character \n.
* Incase of an empty string such as "\n" it takes a height of 1 line.

The program includes several banner formats for graphical representation using ASCII characters:

* shadow
* standard
* thinkertoy
* lean

>  The program automatically takes the banner format of `standard`, but if the user wants to use other graphical representation, the banner format desired should be the third arguement in lower case as listed above.

## Usage
- To run the program, use the command below;
```bash
go run . [STRING] [BANNER]
```
### Running Tests
To run unit tests, navigate to the project directory and run:
```bash
go test -v
```
