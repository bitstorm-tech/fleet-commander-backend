package arango

import (
	"github.com/pkg/errors"
)

var (
	EdgeHasResources = "hasResources"
)

type edge struct {
	From string `json:"_from"`
	To   string `json:"_to"`
}

func CreateEdge(from Persistable, to Persistable, collectionName string) error {
	graph, err := getGraph()
	if err != nil {
		return errors.Wrap(err, "error while getting graph")
	}

	edges, _, err := graph.EdgeCollection(nil, collectionName)
	if err != nil {
		return errors.Wrapf(err, "error while getting graph collection '%s'", collectionName)
	}

	edge := edge{
		From: from.id(),
		To:   to.id(),
	}

	if _, err := edges.CreateDocument(nil, edge); err != nil {
		return errors.Wrapf(err, "error while creating edge (%s -> %s) in collection '%s'", edge.From, edge.To, collectionName)
	}

	return nil
}

//func RemoveEdge(from Persistable, to Persistable, collectionName string) error {
//	graph, err := getGraph()
//	if err != nil {
//		return errors.Wrap(err, "error while getting graph")
//	}
//
//	edges, _, err := graph.EdgeCollection(nil, collectionName)
//	if err != nil {
//		return errors.Wrapf(err, "error while removing edge: %+v", persistable)
//	}
//
//	if _, err := edges.RemoveDocument(nil, persistable.Key()); err != nil {
//		return errors.Wrapf(err, "error while removing edge: %+v", persistable)
//	}
//
//	return nil
//}
