package local

import (
	"accounts/internal/domain/repository"
)

type localRepositoryRegistry struct {
	accountReadOnlyRepository  accountReadOnlyRepository
	accountWriteOnlyRepository accountWriteOnlyRepository
}

//NewLocalRepositoryRegistry builds localRepositoryRegistry with its dependencies
func NewLocalRepositoryRegistry() repository.Registry {
	return localRepositoryRegistry{
		accountReadOnlyRepository:  accountReadOnlyRepository{},
		accountWriteOnlyRepository: accountWriteOnlyRepository{},
	}
}

//AccountReadOnlyRepository returns the in-memory read only repository
func (r localRepositoryRegistry) AccountReadOnlyRepository() repository.AccountReadOnly {
	return r.accountReadOnlyRepository
}

//AccountWriteOnlyRepository returns the in-memory write only repository
func (r localRepositoryRegistry) AccountWriteOnlyRepository() repository.AccountWriteOnly {
	return r.accountWriteOnlyRepository
}

//DoInTransaction .
func (r localRepositoryRegistry) DoInTransaction(txFunc repository.InTransaction) (out interface{}, err error) {
	return txFunc(r)
}
