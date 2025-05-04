package util

import (
	"database/sql"
	"strconv"
	"time"
)

// NullInt16ToPtr converts a sql.NullInt16 to a pointer to int16.
func NullInt16ToPtr(nullInt sql.NullInt16) *int16 {
    if nullInt.Valid {
        return &nullInt.Int16
    }
    return nil
}

func NullStringToPtr(ns sql.NullString) *string {
    if ns.Valid {
        return &ns.String
    }
    return nil
}

func NullTimeToPtr(nt sql.NullTime) *time.Time {
    if nt.Valid {
        return &nt.Time
    }
    return nil
}

func TimeToNullTime(t time.Time) sql.NullTime {
    return sql.NullTime{Time: t, Valid: true}
}

func IntToNullString(value int) sql.NullString {
    return sql.NullString{
        String: strconv.Itoa(value),
        Valid:  true,
    }
}