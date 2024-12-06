# Advent of Code

This repository contains my solutions to the [Advent of Code](https://adventofcode.com/) challenges.

Solutions are written in Go and use utility functions from https://codeberg.org/derat/advent-of-code.

## Setup helper cli

First you need to get your session cookie from the browser. You can do this by inspecting the network requests in the browser and copying the value of the `session` cookie.

Then you can run the following command to setup the helper cli:

```shell
./init.sh <session-cookie>
```

This will store the session cookie in a file called `~/.advent-of-code-session` and register the helper cli in your `PATH`.
