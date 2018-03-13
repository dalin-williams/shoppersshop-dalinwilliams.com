package etc

import (
	"database/sql"
	"fmt"
)

var (
	NoRows  = ExpectRows(0)
	OneRow  = ExpectRows(1)
	TwoRows = ExpectRows(2)
)

type ResultCheck func(sql.Result, error) error

func ExpectRows(n int) ResultCheck {
	return func(res sql.Result, err error) error {
		if err != nil {
			return err
		}

		numRows, err := res.RowsAffected()
		if err != nil {
			return fmt.Errorf("error reading num rows affected: %v", err)
		}

		if numRows != int64(n) {
			return fmt.Errorf("expected %d rows affected, saw %d rows affected", n, numRows)
		}

		return nil
	}
}
