###### logit

logit is a command-line daily logging tool.

1. logit command will create a markdown file with today's date as the title, if it doesn't already exist.

2. Then it will open the markdown file template in neovim to be modified by the user (me).

How will we handle the user's current project? Maybe it should be pre-populated if it exists. Maybe there's a
command that changes the current project... and if there is no current project, it should ask the user what
it's currently working on and set that up.
