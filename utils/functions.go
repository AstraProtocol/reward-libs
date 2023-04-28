package util

import (
	"fmt"
	"strconv"
	"strings"
)

const IMPORT_ID_PREFIX string = "import_id:"

func RequestIdToImportId(requestId string) int {
	parsed, err := strconv.Atoi(strings.TrimPrefix(requestId, IMPORT_ID_PREFIX))
	if err != nil {
		return 0
	}
	return parsed
}

func ImportIdToRequestId(importId int) string {
	return fmt.Sprintf("%s%d", IMPORT_ID_PREFIX, importId)
}
