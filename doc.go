// Copyright 2018 Origoss Solutions. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package bmstruct provides structures and methods to manage byte slices with
embedded values.

Introduction

The following byte slice of 8 bytes contains an uint16 (2 bytes) starting at the
offset 3:

  0:  01000001
  1:  00010000
  2:  01000111
  3: |01001001|
  4: |10010101|
  5:  01011001
  6:  01010011
  7:  00110101

With the package bmstruct this data structure can be managed easily. First,
let's create a bmstruct.Template that describes the location of the uint16
value inside the byte slice:

  t := NewTemplate(8, Uint16Field("value"), 3)

This template can be applied on any byte slice with length of 8. The result is a
bmstruct.Struct.

  b := make([]byte, 8)
  s := t.New(b)

bmstruct.Struct can be easily manipulated with the Lookup and Update methods:

  v := s.Lookup("value") // returns bmstruct.Value
  fmt.Println("v: ", v.Uint16()) // shall print v: 0

The Update method can be used for changing the byte slice behind the Struct.

  s.Update("value", Uint16(42))

The byte slice will look like this:

  0:  00000000
  1:  00000000
  2:  00000000
  3: |00101010|
  4: |00000000|
  5:  00000000
  6:  00000000
  7:  00000000

*/
package bmstruct
