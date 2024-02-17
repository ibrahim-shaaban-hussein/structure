// pkg/shared/logger.go

package shared

import (
    "log"
)

// LogError logs an error.
func LogError(err error) {
    if err != nil {
        log.Println("Error:", err)
    }
}

