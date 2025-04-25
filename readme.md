# logit

logit is a simple command-line tool for daily developer logging. It creates a markdown log file for the current day, remembers your current project, and helps you stay focused.

## Usage

```
logit             # Creates or opens today's log
logit -p NAME     # Sets the current project
logit -t          # Sets or updates today's tasks
logit -l          # Lists all existing logs
```

## Example

```
logit -p my-project
logit -t 
logit
```

This will set the current project, set two tasks for the day with and open today's prefilled log file

## Install

```
git clone https://github.com/mikevidotto/logit.git
cd logit
go build -o logit
mv logit /usr/local/bin/   # optional
```
