package repository

//InTransaction is a function definition that will be used in DoInTransaction as a callback function
type InTransaction func(repoRegistry Registry) (interface{}, error)

//Registry will handle all repositories registries and exposes the DoInTransaction function to
//enable transactions management
type Registry interface {
	AccountReadOnlyRepository() AccountReadOnly
	AccountWriteOnlyRepository() AccountWriteOnly
	DoInTransaction(txFunc InTransaction) (out interface{}, err error)
}
