# The Monkey Go

This is an custom interpreter of the popular Monkey Language, made popualar by Thorsten Ball. This interpreter is written in Go.

## Specifications of the language

#### Supported Operations

- Addition (+)
- Subtraction (-)
- Multipliction (\*)
- Division (/)
- Exponentiation (\*\*)
- Remainder (%)
- Bracketing ((), {})
- Semicolon `;` and Comma `,`
- Basic Arithmetic Comparision Operators (`>`, `>=`, `<`, `<=`, `==`)
- Negation (`!`, `!=`)

#### Language Specifications

- Variable name's must begin with an alphabet and can contain alphabets, numbers and `-`
- Case-sensitive
- Whitespace agnostic i.e. whitespace between literals is ignored during the lexing phase
- `;` is used to denote the end of a statement
- Supports only `int` datatype for numbers
