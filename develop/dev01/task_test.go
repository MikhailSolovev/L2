package main

import "testing"

func TestGetPreciseTime(t *testing.T) {
	tests := []struct {
		name    string
		isValid bool
		host    string
	}{
		{
			name:    "valid",
			isValid: true,
			host:    "0.beevik-ntp.pool.ntp.org",
		},
		{
			name:    "not valid",
			isValid: false,
			host:    "0beevik-ntp.pool.ntp.org",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.isValid == true {
				_, err := GetPreciseTime(test.host)
				if err != nil {
					t.Fatal(err)
				}
			} else {
				_, err := GetPreciseTime(test.host)
				if err == nil {
					t.Fatal()
				}
			}
		})
	}
}
