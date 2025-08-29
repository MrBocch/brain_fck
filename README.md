# brainfck interpreter

Theres prob a bug in '[] loops because rot13 wont work

## Specs

- 30,000 array uint8 initialized to zero.

- Out of bounds check

- Wont tell you if you have a unclosed loop

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
