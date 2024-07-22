package component

import "github.com/orsinium-labs/enum"

type size enum.Member[string]

var (
	SizeSm = size{"sm"}
	SizeMd = size{"md"}
	SizeLg = size{"lg"}
)
