package go_bitpay_client

import (
	uuid "github.com/nu7hatch/gouuid"
)

// This should be generated without including a separate lib
func guid() string {
	u, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	return u.String()
}
