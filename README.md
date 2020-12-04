# Advent of Code

This is my take on the [2020 Advent of Code](https://adventofcode.com/2020) using the GO language.


## Quickstart

1. Checkout the project, then download dependencies:
    ```
    make init
    ```
   
1. Run, giving the door to open and the parts to run as parameters:
    ```
    make run -- -door 1 -parts 1,2
    ```
    or
    ```
    make build
    ./adventofcode.out -door 1 -parts 1,2
    ```
   `door` must be an integer between 1 and 24 (obviously), `parts` is a list of integers separated by commas.
   
You can also run unit tests:
```
make test
```
