package mysql



type InTransaction func(*Repository) error

func PerformTransaction(db *Repository, fn InTransaction) error {

	// transaction begin
	tx := &Repository{
		DB: db.Begin(),
	}
	if tx.Error != nil {
		return tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// run
	err := fn(tx)

	// rollback (when excepiton)
	if err != nil {
		xerr := tx.Rollback().Error
		if xerr != nil {
			return xerr
		}
		return err
	}

	// commit (actually affects)
	if err = tx.Commit().Error; err != nil {
		return err
	}
	return nil
}