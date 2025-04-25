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

This will open today's log file pre-filled with the current date and project.

## Install

```
git clone https://github.com/mikevidotto/logit.git
cd logit
go build -o logit
mv logit /usr/local/bin/   # optional
```
