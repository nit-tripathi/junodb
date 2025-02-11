//
//  Copyright 2023 PayPal Inc.
//
//  Licensed to the Apache Software Foundation (ASF) under one or more
//  contributor license agreements.  See the NOTICE file distributed with
//  this work for additional information regarding copyright ownership.
//  The ASF licenses this file to You under the Apache License, Version 2.0
//  (the "License"); you may not use this file except in compliance with
//  the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.
//

package util

import (
	"fmt"
	"unicode"
)

func ToPrintableString(b []byte) string {
	sz := len(b)
	if sz == 0 {
		return ""
	}
	buf := make([]byte, sz)
	for i := 0; i < sz; i++ {
		if b[i] < 32 || b[i] > 126 {
			buf[i] = '.'
		} else {
			buf[i] = b[i]
		}
	}
	return string(buf)
}

func ToHexString(data []byte) string {
	return fmt.Sprintf("%X", data)
}

func ToPrintableAndHexString(data []byte) string {
	return fmt.Sprintf("%s [%X]", ToPrintableString(data), data)
}

func PrintBytesForTest(data []byte) {
	fmt.Print("{")
	szRawMsg := len(data)
	for i := 0; i < szRawMsg; i++ {
		if i != 0 {
			fmt.Print(", ")
		}
		if i%8 == 0 {
			fmt.Print("\n    ")
		}
		fmt.Printf("0x%02X", int((data[i])))
	}
	fmt.Print("\n}\n")
}

func HexDump(data []byte) {

	fmt.Printf("   %9X  ", 0)
	for i := 1; i < 16; i++ {
		fmt.Printf("%X  ", i)
	}

	for i := 0; i < 16; i++ {
		fmt.Printf("%X", i)
	}
	fmt.Print("\n")

	szData := len(data)
	start := 0
	end := 16
	for start < szData {
		if end > szData {
			end = szData
		}
		fmt.Printf("%09X ", start) //&data[start])
		for j := start; j < end; j++ {
			fmt.Printf("%02X ", data[j])
		}
		for j := (end - 1) % 16; j < 15; j++ {
			fmt.Print("   ")
		}

		fmt.Print(" ")
		for j := start; j < end; j++ {
			v := data[j]
			if unicode.IsPrint(rune(v)) {
				fmt.Printf("%c", v)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
		start += 16
		end += 16
	}
}
