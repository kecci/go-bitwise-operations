# Go Bitwise Operations
Bitwise operations are operations that manipulate individual bits of binary data at a low-level, rather than operating on the entire data. These operations treat each bit of a binary number as a separate entity and perform logical or arithmetic operations on them.

These bitwise operators are commonly used in various applications such as manipulating flags, performing efficient bit-level calculations, encoding/decoding data, or optimizing memory usage. They provide fine-grained control over the individual bits of binary data.

## Bitwise
A bit wise operator perform logical or arithmetic operations on them.
- AND
- OR
- XOR
- NOT

## Bitshift
A bit shift operator is a low-level operator that works on the individual bits of an integer. It takes two operands. One is the integer whose bits we want to shift. The other represents the number of shifts. The result of a bit shift is a new integer. Bit-shift operators are easy to use and execute quickly on modern multicore CPU systems. We use them to reduce memory usage and speed up execution.
- LeftShift
- RightShift
- LeftCircularShift
- RightCircularShift

## Run Example
Command: `go run main.go`

Output Sample:
```go
Original number: 1
LeftShift number: 2147483648 (Shifts: 31)
RightShift number: 0 (Shifts: 31)
LeftCircularShift number: 10000000000000000000000000000001 (Shifts: 31 Bits: 31)
RightCircularShift number: 00000001 (Shifts: 31 Bits: 31)
```

## Source
- https://www.baeldung.com/cs/bitwise-shift-operators
- https://www.techtarget.com/whatis/definition/bitwise#:~:text=Bitwise%20operators%20are%20characters%20that,the%20manipulation%20of%20individual%20bits.
- https://www.w3schools.com/js/js_bitwise.asp