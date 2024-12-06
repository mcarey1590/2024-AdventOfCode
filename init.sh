#!/bin/bash

source ~/.zshrc

cookie_session=$1

# Check if the cookie is valid
if [ -z "$cookie_session" ]; then
  echo "Please provide a session cookie"
  exit 1
fi

# Write the cookie to a file
echo $cookie_session > ~/.advent-of-code-session

# Give permissions to the script
chmod +x ./advent.sh

if ! type advent > /dev/null; then
  echo "Adding advent to .zshrc"

  input_var='"$1"'
  output_var= "$PWD/advent.sh \"$@\""
  output_cd='cd "\$(' + $PWD + '/advent.sh "\$@")"'

  # Update zshrc to include the following script
  cat << EOF >> ~/.zshrc

# AdventOfCode
advent() {
  case $input_var in
    -h|--help|check|checkall|help|input|run|save|stdin|web|next|prev)
      $output_var
      ;;
    *)
      ${output_cd}
      ;;
  esac
}
# End AdventOfCode
EOF

fi





