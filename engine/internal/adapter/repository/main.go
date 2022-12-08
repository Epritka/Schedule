package repository

import (
	"engine/internal/core/interfaces"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

type repositoryManager struct {
	DB *pg.DB
	Tx *pg.Tx
}

func NewRepositoryManager(DB *pg.DB, Tx *pg.Tx) interfaces.RepositoryManager {
	return &repositoryManager{
		DB: DB,
		Tx: Tx,
	}
}

func (t *repositoryManager) BeginTran() (interfaces.RepositoryManager, error) {
	tx, err := t.DB.Begin()
	if err != nil {
		return nil, err
	}
	return NewRepositoryManager(t.DB, tx), nil
}

func (t *repositoryManager) CommitTran() error {
	defer t.clearTransaction()
	return t.Tx.Commit()
}

func (t *repositoryManager) RollbackTran() error {
	defer t.clearTransaction()
	return t.Tx.Rollback()
}

func (t *repositoryManager) Transaction(
	callback func(interfaces.RepositoryManager) error,
) error {
	tm, err := t.BeginTran()
	if err != nil {
		return err
	}

	err = callback(tm)

	if err != nil {
		err := tm.RollbackTran()
		if err != nil {
			return err
		}
		return err
	}

	return tm.CommitTran()
}

func (t *repositoryManager) clearTransaction() {
	t.Tx = nil
}

func (t *repositoryManager) getConnect() orm.DB {
	if t.Tx == nil {
		return t.DB
	}
	return t.Tx
}

func (t *repositoryManager) GetDayRepository() interfaces.DayRepository {
	return NewDayRepository(t.getConnect())
}
