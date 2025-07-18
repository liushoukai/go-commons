package dateutils

import (
	"fmt"
	"time"
)

func FormatYearMonth(timestamp int64) string {
	t := time.UnixMilli(timestamp)
	return fmt.Sprintf("%d%02d", t.Year(), t.Month())
}
