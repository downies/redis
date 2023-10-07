package main

import "testing"

func Test_Example(t *testing.T) {
	tt := []struct {
		expr string
	}{
		{
			"$-1\r\n",
		},
		{
			"*1\r\n$4\r\nping\r\n",
		},
		{
			"*2\r\n$4\r\necho\r\n$5\r\nhello world\r\n",
		},
		{
			"*2\r\n$3\r\nget\r\n$3\r\nkey\r\n",
		},
		{
			"+OK\r\n",
		},
		{
			"-Error message\r\n",
		},
		{
			"$0\r\n\r\n",
		},
		{
			"\"+hello world\\r\\n",
		},
	}
}
