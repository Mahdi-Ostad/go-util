package dbutil

import (
	"context"
	"database/sql"
	"time"
)

type LoggingStmt struct {
	*sql.Stmt
	Log   DatabaseLogger
	Query string
}

func (s *LoggingStmt) Exec(args ...any) (sql.Result, error) {
	start := time.Now()
	res, err := s.Stmt.Exec(args...)
	err = addErrorLine(s.Query, err)
	s.Log.QueryTiming(context.Background(), "Exec", "STMT", args, -1, time.Since(start), err)
	return res, err
}
