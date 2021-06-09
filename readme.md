This project is designed to work pretty well exclusively with pgx and postgresql.  The module itself will setup 
connections for testing/development, but everything for actual data access exists in a sub path.  This also means that
tests are viable, and CCP is likely to break that down the lines, but hey, that's part of the fun. 

As go is classless, the system expects that you'll pass in a direct pgxpool connection to it, which should be
pre-resolved to the SDE data export in psql.  Ideally, nothing else will access this, and it will be a read-only
connection.  If the env var MEMCACHE_SERVER is set, it will attempt to use memcache as a long-term data store
but it is designed to handle data storage in-memory using a centrally accessable KV in memory store in the application
itself.

The connections here are read only, and self-detach once they're done running in their various sub functions.  The main
file is used for testing/example implementations and contains some envvars, documented below.  It does expect to have
some access to a caching layer, it'll use an in-memory LRU by default.