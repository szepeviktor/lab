# Based on hub's autocompletion script at: https://github.com/github/hub
# If there is no git tab completion, but we have the _completion loader try to load it
if ! declare -F _git > /dev/null && declare -F _completion_loader > /dev/null; then
  _completion_loader git
fi

# Check that git tab completion is available
if declare -F _git > /dev/null; then
  # Duplicate and rename the 'list_all_commands' function
  eval "$(declare -f __git_list_all_commands | \
      sed 's/__git_list_all_commands ()/__git__list__all__commands_without_lab ()/')"

  # Wrap the 'list_all_commands' function with extra lab commands
  __git_list_all_commands() {
    cat <<-EOF
fork
mr
issue
merge-request
EOF
    __git__list__all__commands_without_lab
  }

  # Ensure cached commands are cleared
  __git_all_commands=""

  _git_issue ()
  {
       __gitcomp "create list"
  }

  _git_mr ()
  {
       __gitcomp "checkout create list"
  }

  complete -o bashdefault -o default -o nospace -F _git lab 2>/dev/null \
    || complete -o default -o nospace -F _git lab
fi
