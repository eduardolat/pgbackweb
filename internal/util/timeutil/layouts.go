package timeutil

import "time"

const (
	// LayoutRFC3339 is the layout for dates in the format YYYY-MM-DDTHH:MM:SSZ
	LayoutRFC3339 = time.RFC3339

	// LayoutRFC3339Nano is the layout for dates in the format YYYY-MM-DDTHH:MM:SS.sssssssssZ
	LayoutRFC3339Nano = time.RFC3339Nano

	// LayoutInputDateTimeLocal is the layout for dates in the format YYYY-MM-DDThh:mm
	//
	// As described in:
	// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/datetime-local#value
	//
	// Used in <input type="datetime-local"/>
	LayoutInputDateTimeLocal = "2006-01-02T15:04"

	// LayoutInputDate is the layout for dates in the format YYYY-MM-DD
	//
	// As described in:
	// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/date#value
	//
	// Used in <input type="date"/>
	LayoutInputDate = "2006-01-02"

	// LayoutSlashDDMMYYYY is the layout for dates in the format DD/MM/YYYY
	LayoutSlashDDMMYYYY = "02/01/2006"

	// LayoutSlashYYYYMMDD is the layout for dates in the format YYYY/MM/DD
	LayoutSlashYYYYMMDD = "2006/01/02"

	// LayoutDashDDMMYYYY is the layout for dates in the format DD-MM-YYYY
	LayoutDashDDMMYYYY = "02-01-2006"

	// LayoutDashYYYYMMDD is the layout for dates in the format YYYY-MM-DD
	LayoutDashYYYYMMDD = "2006-01-02"

	// LayoutYYYYMMDDHHMMSS is the layout for dates in the format YYYYMMDD-HHMMSS
	LayoutYYYYMMDDHHMMSS = "20060102-150405"
)
