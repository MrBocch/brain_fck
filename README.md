# brainfck interpreter

I idk how correct it is.

## Specs

- 30,000 array uint8 initialized to zero.

- Out of bounds check

## Custom

Added own operator for dereferencing value at certain address

`_<addr in decimal>`

example

```brainfuck
+++>
+++<
[- > + <] _1
```
Output: 6
