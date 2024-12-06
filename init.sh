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

brew install coreutils

if ! type advent > /dev/null; then
  echo "Adding advent to .zshrc"

  # Update zshrc to include the following script
  cat << EOF >> ~/.zshrc

# AdventOfCode
advent() {
  case "\$1" in
    -h|--help|check|checkall|help|input|run|save|stdin|web)
      $PWD/advent.sh "\$@"
      ;;
    *)
      cd "\$($PWD/advent.sh "\$@")"
      ;;
  esac
}
# End AdventOfCode
EOF

fi





