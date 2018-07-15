/*
 *  Copyright (c) 2018, https://github.com/airwide-code
 *  All rights reserved.
 *
 *
 *
 */

package webpage

import (
	"testing"
	"fmt"
	"net/url"
)

func TestGetWebpageOgList(t *testing.T) {
	ogContents := GetWebpageOgList("https://github.com/airwide-code/airwide.datacenter", []string{"image", "site_name", "title", "description"})
	fmt.Println(ogContents)
}

func TestUrlParser(t *testing.T) {
	var (
		u *url.URL
		err error
	)

	u, err = url.Parse("aaaa")
	fmt.Println(u, err)
	u, err = url.Parse("https://github.com/airwide-code/airwide.datacenter")
	fmt.Println(u, err)
}

