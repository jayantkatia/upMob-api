// Code generated by sqlc. DO NOT EDIT.

package db

import ()

type Device struct {
	DeviceName string `json:"device_name"`
	// "-" implies not specified, saved as :- Expected launch: Month 20xx
	Expected  string `json:"expected"`
	Price     int64  `json:"price"`
	ImgUrl    string `json:"img_url"`
	SourceUrl string `json:"source_url"`
	// 0 implies not specified
	SpecScore int32 `json:"spec_score"`
}
