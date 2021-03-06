package mongodb

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/loivis/marvel-comics-api-data-loader/maco"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

var defaultTimeout time.Duration = 15 * time.Second

const (
	ColCharacters = "characters"
	ColComics     = "comics"
	ColCreators   = "creators"
	ColEvents     = "events"
	ColSeries     = "series"
	ColStories    = "stories"
)

type MongoDB struct {
	client   *mongo.Client
	database string
	timeout  time.Duration

	cacheMu  sync.Mutex
	cacheIDs map[string][]int // map of collection to all ids
}

func New(uri string, database string) (*MongoDB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	opts := options.Client().
		SetRetryWrites(true).
		SetWriteConcern(
			writeconcern.New(writeconcern.WMajority()),
		)
	client, err := mongo.Connect(ctx, opts.ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to mongodb: %v", err)
	}

	// set context to nil in order to get real error instead of early timeout
	err = client.Ping(nil, readpref.Primary())
	if err != nil {
		return nil, fmt.Errorf("failed to ping mongodb: %v", err)
	}

	return &MongoDB{
		client:   client,
		database: database,
		timeout:  defaultTimeout,
		cacheIDs: make(map[string][]int),
	}, nil
}

func (m *MongoDB) GetCount(ctx context.Context, collection string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), m.timeout)
	defer cancel()

	col := m.client.Database(m.database).Collection(collection)

	count, err := col.CountDocuments(ctx, bson.D{})

	return int(count), err
}

func (m *MongoDB) IncompleteIDs(ctx context.Context, collection string) ([]int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), m.timeout)
	defer cancel()

	col := m.client.Database(m.database).Collection(collection)

	cur, err := col.Find(ctx,
		bson.D{{Key: "intact", Value: false}},
		options.Find().SetProjection(bson.D{{Key: "id", Value: 1}}),
	)
	if err != nil {
		return nil, fmt.Errorf("error finding incomplete documents: %v", err)
	}

	var ids []int

	for cur.Next(ctx) {
		var elem struct{ ID int }
		err := cur.Decode(&elem)
		if err != nil {
			return nil, fmt.Errorf("error decoding document: %v", err)
		}

		ids = append(ids, elem.ID)
	}

	if err := cur.Err(); err != nil {
		return nil, fmt.Errorf("error decoding all documents: %v", err)
	}

	cur.Close(ctx)

	return ids, nil
}

func (m *MongoDB) SaveCharacters(ctx context.Context, chars []*maco.Character) error {
	ctx, cancel := context.WithTimeout(context.Background(), m.timeout)
	defer cancel()

	var docs []maco.Doc
	for _, char := range chars {
		docs = append(docs, char)
	}

	return m.saveMany(ctx, ColCharacters, docs)
}

func (m *MongoDB) SaveComics(ctx context.Context, comics []*maco.Comic) error {
	ctx, cancel := context.WithTimeout(context.Background(), m.timeout)
	defer cancel()

	var docs []maco.Doc
	for _, comic := range comics {
		docs = append(docs, comic)
	}

	return m.saveMany(ctx, ColComics, docs)
}

func (m *MongoDB) SaveCreators(ctx context.Context, creators []*maco.Creator) error {
	ctx, cancel := context.WithTimeout(context.Background(), m.timeout)
	defer cancel()

	var docs []maco.Doc
	for _, creator := range creators {
		docs = append(docs, creator)
	}

	return m.saveMany(ctx, ColCreators, docs)
}

func (m *MongoDB) SaveEvents(ctx context.Context, events []*maco.Event) error {
	ctx, cancel := context.WithTimeout(context.Background(), m.timeout)
	defer cancel()

	var docs []maco.Doc
	for _, event := range events {
		docs = append(docs, event)
	}

	return m.saveMany(ctx, ColEvents, docs)
}

func (m *MongoDB) SaveSeries(ctx context.Context, series []*maco.Series) error {
	ctx, cancel := context.WithTimeout(context.Background(), m.timeout)
	defer cancel()

	var docs []maco.Doc
	for _, s := range series {
		docs = append(docs, s)
	}

	return m.saveMany(ctx, ColSeries, docs)
}

func (m *MongoDB) SaveStories(ctx context.Context, stories []*maco.Story) error {
	ctx, cancel := context.WithTimeout(context.Background(), m.timeout)
	defer cancel()

	var docs []maco.Doc
	for _, s := range stories {
		docs = append(docs, s)
	}

	return m.saveMany(ctx, ColStories, docs)
}

func (m *MongoDB) saveMany(ctx context.Context, collection string, docs []maco.Doc) error {
	if len(docs) == 0 {
		log.Info().Msg("no docs to save")
		return nil
	}

	ids, err := m.getAllIds(ctx, collection)
	if err != nil {
		return err
	}

	newDocs := diff(ids, docs)

	if len(newDocs) == 0 {
		log.Info().Msg("no new docs to save")
		return nil
	}

	many := []interface{}{}
	for _, doc := range newDocs {
		many = append(many, doc)
	}

	col := m.client.Database(m.database).Collection(collection)
	_, err = col.InsertMany(ctx, many)
	if err != nil {
		return err
	}

	log.Info().Int("count", len(many)).Msg("saved new docs")

	m.cacheMu.Lock()
	for _, doc := range newDocs {
		m.cacheIDs[collection] = append(m.cacheIDs[collection], doc.Identify())
	}
	m.cacheMu.Unlock()
	log.Info().Int("before", len(ids)).Int("after", len(ids)+len(newDocs)).Msg("added new doc ids to cache")

	return nil
}

func (m *MongoDB) SaveOne(ctx context.Context, doc maco.Doc) error {
	var collection string
	switch doc.(type) {
	case *maco.Character:
		collection = ColCharacters
	case *maco.Comic:
		collection = ColComics
	case *maco.Creator:
		collection = ColCreators
	case *maco.Event:
		collection = ColEvents
	case *maco.Series:
		collection = ColSeries
	case *maco.Story:
		collection = ColStories
	default:
		return fmt.Errorf("unsupported type: %T", doc)
	}

	ctx, cancel := context.WithTimeout(context.Background(), m.timeout)
	defer cancel()

	col := m.client.Database(m.database).Collection(collection)

	id := doc.Identify()

	result, err := col.ReplaceOne(ctx, bson.D{{Key: "id", Value: id}}, doc)
	if err != nil {
		return err
	}

	log.Info().Interface("result", result).Int("id", id).Msgf("document %T(%d) replaced", doc, id)

	return nil
}

func (m *MongoDB) getAllIds(ctx context.Context, collection string) ([]int, error) {
	m.cacheMu.Lock()
	defer m.cacheMu.Unlock()

	if ids, ok := m.cacheIDs[collection]; ok {
		log.Info().Int("count", len(ids)).Msg("returned all ids from cache")
		return ids, nil
	}

	ids := []int{}

	col := m.client.Database(m.database).Collection(string(collection))

	cur, err := col.Find(ctx, bson.D{})
	if err != nil {
		return nil, fmt.Errorf("error finding documents: %v", err)
	}

	for cur.Next(ctx) {
		var elem struct{ ID int }

		err := cur.Decode(&elem)
		if err != nil {
			return nil, fmt.Errorf("error decoding document: %v", err)
		}
		ids = append(ids, elem.ID)
	}

	if err := cur.Err(); err != nil {

		return nil, fmt.Errorf("error from cursor: %v", err)
	}
	cur.Close(ctx)

	m.cacheIDs[collection] = ids

	log.Info().Int("count", len(ids)).Msg("returned all ids from mongodb query")

	return ids, nil
}

func diff(ids []int, docs []maco.Doc) []maco.Doc {
	m := make(map[int]struct{}, len(ids))
	for i := range ids {
		m[ids[i]] = struct{}{}
	}

	var r []maco.Doc
	seen := map[int]struct{}{}
	for i := range docs {
		if _, ok := m[docs[i].Identify()]; ok {
			continue
		}

		if _, ok := seen[docs[i].Identify()]; ok {
			continue
		}

		seen[docs[i].Identify()] = struct{}{}
		r = append(r, docs[i])
	}

	return r
}
