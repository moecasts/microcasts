// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package endpoint

import (
	endpoint "github.com/go-kit/kit/endpoint"
	service "github.com/moecasts/microcasts/novels/pkg/service"
)

// Endpoints collects all of the endpoints that compose a profile service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	BrowseEndpoint       endpoint.Endpoint
	ReadEndpoint         endpoint.Endpoint
	AddEndpoint          endpoint.Endpoint
	EditEndpoint         endpoint.Endpoint
	TrashEndpoint        endpoint.Endpoint
	RestoreEndpoint      endpoint.Endpoint
	DestroyEndpoint      endpoint.Endpoint
	BatchAddEndpoint     endpoint.Endpoint
	BatchEditEndpoint    endpoint.Endpoint
	BatchTrashEndpoint   endpoint.Endpoint
	BatchRestoreEndpoint endpoint.Endpoint
	BatchDestroyEndpoint endpoint.Endpoint
}

// New returns a Endpoints struct that wraps the provided service, and wires in all of the
// expected endpoint middlewares
func New(s service.NovelsService, mdw map[string][]endpoint.Middleware) Endpoints {
	eps := Endpoints{
		AddEndpoint:          MakeAddEndpoint(s),
		BatchAddEndpoint:     MakeBatchAddEndpoint(s),
		BatchDestroyEndpoint: MakeBatchDestroyEndpoint(s),
		BatchEditEndpoint:    MakeBatchEditEndpoint(s),
		BatchRestoreEndpoint: MakeBatchRestoreEndpoint(s),
		BatchTrashEndpoint:   MakeBatchTrashEndpoint(s),
		BrowseEndpoint:       MakeBrowseEndpoint(s),
		DestroyEndpoint:      MakeDestroyEndpoint(s),
		EditEndpoint:         MakeEditEndpoint(s),
		ReadEndpoint:         MakeReadEndpoint(s),
		RestoreEndpoint:      MakeRestoreEndpoint(s),
		TrashEndpoint:        MakeTrashEndpoint(s),
	}
	for _, m := range mdw["Browse"] {
		eps.BrowseEndpoint = m(eps.BrowseEndpoint)
	}
	for _, m := range mdw["Read"] {
		eps.ReadEndpoint = m(eps.ReadEndpoint)
	}
	for _, m := range mdw["Add"] {
		eps.AddEndpoint = m(eps.AddEndpoint)
	}
	for _, m := range mdw["Edit"] {
		eps.EditEndpoint = m(eps.EditEndpoint)
	}
	for _, m := range mdw["Trash"] {
		eps.TrashEndpoint = m(eps.TrashEndpoint)
	}
	for _, m := range mdw["Restore"] {
		eps.RestoreEndpoint = m(eps.RestoreEndpoint)
	}
	for _, m := range mdw["Destroy"] {
		eps.DestroyEndpoint = m(eps.DestroyEndpoint)
	}
	for _, m := range mdw["BatchAdd"] {
		eps.BatchAddEndpoint = m(eps.BatchAddEndpoint)
	}
	for _, m := range mdw["BatchEdit"] {
		eps.BatchEditEndpoint = m(eps.BatchEditEndpoint)
	}
	for _, m := range mdw["BatchTrash"] {
		eps.BatchTrashEndpoint = m(eps.BatchTrashEndpoint)
	}
	for _, m := range mdw["BatchRestore"] {
		eps.BatchRestoreEndpoint = m(eps.BatchRestoreEndpoint)
	}
	for _, m := range mdw["BatchDestroy"] {
		eps.BatchDestroyEndpoint = m(eps.BatchDestroyEndpoint)
	}
	return eps
}
