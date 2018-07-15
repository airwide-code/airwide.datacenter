/*
 *  Copyright (c) 2018, https://github.com/airwide-code/airwide.datacenter
 *  All rights reserved.
 *
 *
 *
 */

package main

import (
	"github.com/airwide-code/airwide.datacenter/mtproto"
	"fmt"
	"strconv"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println(" ./schema.tl.crc32_tool xxx [...]")
		os.Exit(0)
	}

	for i := 1; i < len(os.Args); i++ {
		n, err := strconv.ParseInt(os.Args[i], 0, 64)
		if err != nil {
			fmt.Println(os.Args[i], " conv error: ", err)
		} else {
			if crc32, ok := mtproto.TLConstructor_name[int32(n)]; !ok {
				fmt.Printf("[%d, 0x%x] ==> %s\n", int32(n), uint32(n), "CRC32_UNKNOWN")
			} else {
				fmt.Printf("[%d, 0x%x] ==> %s\n", int32(n), uint32(n), crc32)
			}
		}
	}
}
