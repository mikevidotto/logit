###### logit

logit is a command-line daily logging tool.

1. logit command will create a markdown file with today's date as the title, if it doesn't already exist.

2. Then it will open the markdown file template in neovim to be modified by the user (me).

3. FLAGS:
    - [x] -p flag will allow user to set their current project value in the config.
    - [x] -h flag will show usage and flags.
    - [x] -l flag will list all of the previous logs, sorted by date created.
    - [x] -t flag will show you the two tasks that you've set for today.

4. logs directory
    - [x] Remove /logs/ directory from project folder (and git repo) and place it somewhere else so you're not making logs public every time you push to git...

5. Refactor and Implement Project Folder Structure:
    - [ ] Move config manipulation functions into internal/config
    - [ ] Move log file creation and handling into internal/logger
    - [ ] Move utility functions into internal/utils
    - [ ] Move main.go into cmd/logit

       logit/
       ├── cmd/
       │   └── logit/
       │       └── main.go
       ├── internal/
       │   ├── config/
       │   │   └── config.go
       │   ├── logger/
       │   │   └── logger.go
       │   └── utils/
       │       └── utils.go
       ├── go.mod
       └── README.md

6. Make the sections of the template into separate components
    - [ ] Make a template for each section:
        - [ ] Daily Log 
        - [ ] Currently Working On 
        - [ ] Intent for Today
        - [ ] Tasks
        - [ ] Learning & Thoughts
    - [ ] Render the templates
        - templates should:
            - [ ] not overwrite changes made in the log file
                - [ ] parse log file for existing changes...
                - [ ] OR only append to templates
            - [ ] be rendered in order
            - [ ] have the correct data passed in
