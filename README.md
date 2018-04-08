# Graph Theory Project
## Introduction
My name is Colm Woodlock, student number : G00341460, and this is my attempt at the project for the graph theory module.
This project is worth 50% of the overall mark of the module. The project is written in Golang.

### Instructions to Downloading and Running this project
### Downloading
First to download the project you need to ensure you have both git and go installed on your machine. To do this follow the links below and you can download and install both.
https://golang.org/dl/
https://git-scm.com/downloads

Once you have both downloaded you can now download the project itself. To do this simply open a command line and navigate to a directory you wish to download the file to. Once you are in this directory simply enter the following commands: 
git clone https://github.com/cwoodlock/GraphTheoryProject
Once this command completes you will now have it cloned into your selected directory.

### Running
Once you have downloaded everything you need you can now run the project. Again this is very simple. All you have to do is ensure you are in the directory which contains the project and enter the command: 
go run main.go
This will run the project and you will be prompted by the command line menu.

## Problem statement
You must write a program in the Go programming language [2] that can
build a non-deterministic finite automaton (NFA) from a regular expression,
and can use the NFA to check if the regular expression matches any given
string of text. You must write the program from scratch and cannot use the
regexp package from the Go standard library nor any other external library.
A regular expression is a string containing a series of characters, some
of which may have a special meaning. For example, the three characters
“.”, “|”, and “∗
” have the special meanings “concatenate”, “or”, and “Kleene
star” respectively. So, 0.1 means a 0 followed by a 1, 0|1 means a 0 or a 1,
and 1∗ means any number of 1’s. These special characters must be used in
your submission. Other special characters you might consider allowing as input are brackets
“()” which can be used for grouping, “+” which means “at least one of”, and
“?” which means “zero or one of”. You might also decide to remove the
concatenation character, so that 1.0 becomes 10, with the concatenation
implicit.
You may initially restrict the non-special characters your program works
with to 0 and 1, if you wish. However, you should at least attempt to expand
these to all of the digits, and the characters a to z, and A to Z.
You are expected to be able to break this project into a number of smaller
tasks that are easier to solve, and to plug these together after they have been
completed.

## How I broke down the problem
From my reading of the problem there were a few main things that had to be done. One was the ability to change infix regular expressions into postfix regular expressions. Secondly was to be able to build an NFA(non-deterministic finite automaton) from a regular expression. Finally was the ability to check if the regular expression matches any given string of test by using the NFA. Using this as a guide I was able to break the project up into steps.

1. First did the video lectures online which gave a good baseline of code and understanding to work from
2. From there established a simple command line menu that let you select an option to change infix to postfix
3. After this it was working on an option to build NFA's from regular expression
4. Next was building an additional option to compare a user entered string to compare to the regukar expression.
5. Finally was extending the functionality of some of the methods to include other "wildcard" characters like + or ?

## Conclusion
Overall I found this project pretty challenging, but feel like I learned something from the project. The videos I found super useful , both this project and the last module with did with lecturer Ian McLoughlin. I found them easy to follow along, things were explained well and the abilty to go back and few things multiple times was a massive help. An issue I had at the start was I coulldnt get the packages in seperate files to link to the main, I fixed this by just including everything in main.go. Other than that I didn't have many issues that weren't solved by sitting down and thinking about the logic. Overall I felt like this was a good learning experience.

## References
