package repository

import (
	"golang-couchbase/pkg/api"

	"github.com/couchbase/gocb/v2"
)

type Storage interface {
	SearchAirport(searchKey string) (api.SearchAirportResponse, error)
}

type storage struct {
	cluster *gocb.Cluster
	bucket  *gocb.Bucket
}

func NewStorage(cluster *gocb.Cluster, bucket *gocb.Bucket) Storage {
	return &storage{cluster: cluster, bucket: bucket}
}
