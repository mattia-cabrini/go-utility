// Copyright (c) 2023 Mattia Cabrini
// SPDX-License-Identifier: MIT

package utility

import "database/sql"

func CommitTx(tx *sql.Tx, err error) (_ error) {
	if err == nil {
		err = tx.Commit()
	} else {
		tx.Rollback()
	}

	return AppendError(err)
}
