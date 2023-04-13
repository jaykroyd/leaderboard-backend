package client

import "fmt"

var (
	ErrBuildingUrl      = fmt.Errorf("failed to build request url")
	ErrParsingUri       = fmt.Errorf("error parsing uri string into url")
	ErrDecodingResponse = fmt.Errorf("error decoding response")
	ErrStatusNotOk      = fmt.Errorf("status code is not ok")
)
