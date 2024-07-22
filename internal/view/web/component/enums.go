package component

import "github.com/orsinium-labs/enum"

type (
	size             enum.Member[string]
	dropdownPosition enum.Member[string]
)

var (
	SizeSm = size{"sm"}
	SizeMd = size{"md"}
	SizeLg = size{"lg"}

	DropdownPositionTop    = dropdownPosition{"top"}
	DropdownPositionBottom = dropdownPosition{"bottom"}
	DropdownPositionLeft   = dropdownPosition{"left"}
	DropdownPositionRight  = dropdownPosition{"right"}
)
