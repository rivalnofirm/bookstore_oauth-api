package db

import (
	"github.com/rivalnofirm/bookstore_oauth-api/src/src/clients/cassandra"
	"github.com/rivalnofirm/bookstore_oauth-api/src/src/domain/access_token"
	"github.com/rivalnofirm/bookstore_oauth-api/src/src/utils/errors"
)

func NewRepsitory() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}

type dbRepository struct {
}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	session, err := cassandra.GetSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	//TODO Implement get acces token from cassandraDB.
	return nil, errors.NewInternalServerError("database connection not implemented yet")
}
