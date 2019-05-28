package config

import (
	"bytes"
	"database/sql"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/prince1809/vchat-server/mlog"
	"github.com/prince1809/vchat-server/model"

	// Load the MySQL driver
	_ "github.com/go-sql-driver/mysql"
	// Load the Postgres driver
	_ "github.com/lib/pq"
)

// DatabaseStore is a config store backed by a database.
type DatabaseStore struct {
	commonStore

	originalDsn    string
	driverName     string
	dataSourceName string
	db             *sqlx.DB
}

// NewDatabaseStore creates a new instance of a config store backed by the given database.
func NewDatabaseStore(dsn string) (ds *DatabaseStore, err error) {
	driverName, dataSourceName, err := parseDSN(dsn)
	if err != nil {
		return nil, errors.Wrap(err, "invalid DSN")
	}
	db, err := sqlx.Open(driverName, dataSourceName)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to connect to %s database", driverName)
	}

	ds = &DatabaseStore{
		driverName:     driverName,
		originalDsn:    dsn,
		dataSourceName: dataSourceName,
		db:             db,
	}

	if err = initializeConfigurationsTable(ds.db); err != nil {
		return nil, errors.Wrap(err, "failed to initialize")
	}

	if err = ds.Load(); err != nil {
		return nil, errors.Wrap(err, "failed to load")
	}
	return ds, nil
}

// initializeConfigurationsTable ensures the requisite tables in place to form the backing store.
func initializeConfigurationsTable(db *sqlx.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS Configurations (
			Id VARCHAR(26) PRIMARY KEY,
			value TEXT NOT NULL,
			CreateAt BIGINT NOT NULL,
			Active BOOLEAN NULL UNIQUE
		)
	`)
	if err != nil {
		return errors.Wrap(err, "failed to create Configurations table")
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS ConfigurationFiles (
			Name VARCHAR(64) PRIMARY KEY,
			Data TEXT NOT NULL,
			CreatAt BIGINT NOT NULL,
			UpdateAt BIGINT NOT NULL
		)
	`)

	if err != nil {
		return errors.Wrap(err, "failed to create ConfigurationFiles table")
	}
	return nil
}

// parseDSN splits up a connection string into a driver name and data source name.
//
// For example:
// mysql://mmuser:mostest@dockerhost:3306/mattermost_test
// returns
// driverName = mysql
// dataSourceName = mmuser:mostest@dockerhost:3306/mattermost_test
//
// By contrast, Postgres DSN is returned unmodified
func parseDSN(dsn string) (string, string, error) {
	// Treat the DSN as the URL that it is.
	u, err := url.Parse(dsn)
	if err != nil {
		return "", "", errors.Wrap(err, "failed to parse DSN as URL")
	}

	scheme := u.Scheme
	switch scheme {
	case "mysql":
		// Strip off the mysql:// for the dsn with which to connect
		u.Scheme = ""
		dsn = strings.TrimPrefix(u.String(), "//")

	case "postgres":
	// No changes required
	default:
		return "", "", errors.Wrapf(err, "unsupported scheme %s", scheme)
	}

	return scheme, dsn, nil
}

// persist writes the configuration to the configured database.
func (ds *DatabaseStore) persist(cfg *model.Config) error {
	b, err := marshalConfig(cfg)
	if err != nil {
		return errors.Wrap(err, "failed to serialize")
	}

	id := model.NewId()
	value := string(b)
	createAt := model.GetMillis()

	tx, err := ds.db.Beginx()
	if err != nil {
		return errors.Wrap(err, "failed to begin transaction")
	}
	defer func() {
		// Rollback after commit just returns sql.ErrTxDone.
		if err := tx.Rollback(); err != nil && err != sql.ErrTxDone {
			mlog.Error("Failed to rollback configuration transactions", mlog.Err(err))
		}
	}()

	params := map[string]interface{}{
		"id":        id,
		"value":     value,
		"create_at": createAt,
		"key":       "ConfigurationId",
	}

	// Skip the persist altogether if we'are effectively writing the same configuration.
	var oldValue []byte
	row := ds.db.QueryRow("SELECT Value FROM Configurations WHERE Active")
	if err := row.Scan(&oldValue); err != nil && err != sql.ErrNoRows {
		return errors.Wrap(err, "failed to query active configuration")
	}
	if bytes.Equal(oldValue, b) {
		return nil
	}

	if _, err := tx.Exec("UPDATE Configurations SET Active = NULL WHERE  Active"); err != nil {
		return errors.Wrap(err, "failed to deactivate current configuration")
	}

	if _, err := tx.NamedExec("INSERT INTO Configurations (Id, Value, CreateAt, Active) VALUES (:id, :value, :create_at, TRUE)", params); err != nil {
		return errors.Wrap(err, "failed to record new configuration")
	}

	if err := tx.Commit(); err != nil {
		return errors.Wrap(err, "failed to commit transaction")
	}

	return nil
}

// Load updates the current configuration from the backing store.
func (ds *DatabaseStore) Load() (err error) {
	var needsSave bool
	var configurationData []byte

	row := ds.db.QueryRow("SELECT Value FROM Configurations WHERE Active")
	if err = row.Scan(&configurationData); err != nil && err != sql.ErrNoRows {
		return errors.Wrap(err, "failed to query active configuration")
	}
	// Initialize from the default config if no active configuration could be found.
	if len(configurationData) == 0 {

	}

	return ds.commonStore.load(ioutil.NopCloser(bytes.NewReader(configurationData)), needsSave, ds.commonStore.validate, ds.persist)
}

func (ds *DatabaseStore) Close() error {
	ds.configLock.Lock()
	defer ds.configLock.Unlock()

	return ds.db.Close()
}
