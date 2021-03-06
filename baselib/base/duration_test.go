/*
 *  Copyright (c) 2018, https://github.com/airwide-code/airwide.datacenter
 *  All rights reserved.
 *
 *
 *
 */

package base

import (
	"testing"
	"github.com/BurntSushi/toml"
	"github.com/stretchr/testify/assert"
)

func TestParseDuration(t *testing.T)  {
	var d Duration
	testTable := []string{"2s","10m","2h45m","10us","100ns","20h"}
	for _, v := range testTable {
		d.UnmarshalText([]byte(v))
		t.Log(d)
	}
}

func TestUnmarshalText(t *testing.T) {
	var testDurationSimple = `
		second = "2s"
		interval = "10m"
		mix = "2h45m"
		us = "10us"
		ns = "100ns"
		h = "20h"
		`
	type Config struct {
		Second      Duration
		Minute      Duration `toml:"interval"`
		Mix         Duration
		Microsecond Duration `toml:"us"`
		Nanosecond  Duration `toml:"ns"`
		Hour        Duration `toml:"h"`
	}

	var result Config
	_, err := toml.Decode(testDurationSimple, &result)
	if err != nil {
		t.Fatal(err)
	}

	//t.Log(result)

	expected := Config{
		Second:Duration(2000000000),
		Minute:Duration(600000000000),
		Mix:Duration(9900000000000),
		Microsecond:Duration(10000),
		Nanosecond:Duration(100),
		Hour:Duration(72000000000000),
	}

	assert.Equal(t, expected, result)
}

func BenchmarkUnmarshalText(b *testing.B) {
	var d Duration

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		d.UnmarshalText([]byte(`2h45m`))
	}
}
