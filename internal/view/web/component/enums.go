package component

import "github.com/orsinium-labs/enum"

type (
	size             enum.Member[string]
	color            enum.Member[string]
	dropdownPosition enum.Member[string]
	inputType        enum.Member[string]
)

var (
	SizeSm = size{"sm"}
	SizeMd = size{"md"}
	SizeLg = size{"lg"}

	ColorPrimary   = color{"primary"}
	ColorSecondary = color{"secondary"}
	ColorAccent    = color{"accent"}
	ColorNeutral   = color{"neutral"}
	ColorInfo      = color{"info"}
	ColorSuccess   = color{"success"}
	ColorWarning   = color{"warning"}
	ColorError     = color{"error"}

	DropdownPositionTop    = dropdownPosition{"top"}
	DropdownPositionBottom = dropdownPosition{"bottom"}
	DropdownPositionLeft   = dropdownPosition{"left"}
	DropdownPositionRight  = dropdownPosition{"right"}

	InputTypeText     = inputType{"text"}
	InputTypePassword = inputType{"password"}
	InputTypeEmail    = inputType{"email"}
	InputTypeNumber   = inputType{"number"}
	InputTypeTel      = inputType{"tel"}
	InputTypeUrl      = inputType{"url"}
)
