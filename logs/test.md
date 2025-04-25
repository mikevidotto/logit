# Daily Log — {{Date}}

## Currently Working On
logit

## Intent for Today
What’s the one thing I want to move forward today?

- [ ] Make progress on logit 

## Tasks
- [x] logit: add flag to update "current project" in config file 
- [ ] logit: add functionality to populate placeholder data

## Learning & Thoughts
- What did I learn today?

1. How to create a custom command in the terminal!!!
    - You must use "Go build". This creates an executable file. 
    - Put the executable in your PATH location: ".../go/bin/" 
    - Now you can enter the command from any directory.

2. How to add flags to your Go command-line tool using the flags package.
    - First you declare the flag(s). 
    - Then you use flags.Parse() to parse the flags. 
    - You can check if a flag was passed in by using if *<flag> {...}

3. How to check if a file exists in Go (and how to compare errors)
    - Use the os.Stat method with the path to the file.
    - Use errors package to check for ErrNotExist error

    Code example:

		_, err := os.Stat(path) 
        errors.Is(err, os.ErrNotExist) {
            // file does not exist.
        }


- Notes, ideas, or distractions

