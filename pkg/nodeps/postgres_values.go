package nodeps

// PostgresDefaultVersion is the default PostgreSQL version
const PostgresDefaultVersion = Postgres14

// ValidPostgresVersions is the versions of PostgreSQL that are valid
var ValidPostgresVersions = map[string]bool{
	Postgres15: true,
	Postgres14: true,
	Postgres13: true,
	Postgres12: true,
	Postgres11: true,
	Postgres10: true,
	Postgres9:  true,
}

// PostgreSQL Versions
const (
	Postgres15 = "15"
	Postgres14 = "14"
	Postgres13 = "13"
	Postgres12 = "12"
	Postgres11 = "11"
	Postgres10 = "10"
	Postgres9  = "9"
)

// PostgresConfigDir is in-container location of PostgreSQL config
const PostgresConfigDir = "/etc/postgresql"
