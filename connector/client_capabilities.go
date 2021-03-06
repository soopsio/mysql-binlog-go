package connector

const (
	CAPABILITIES_LONG_PASSWORD uint32 = 1 << iota
	CAPABILITIES_FOUND_ROWS
	CAPABILITIES_LONG_FLAG
	CAPABILITIES_CONNECT_WITH_DB
	CAPABILITIES_NO_SCHEMA
	CAPABILITIES_COMPRESS
	CAPABILITIES_ODBC
	CAPABILITIES_LOCAL_FILES
	CAPABILITIES_IGNORE_SPACE
	CAPABILITIES_PROTOCOL_41
	CAPABILITIES_INTERACTIVE
	CAPABILITIES_SSL
	CAPABILITIES_IGNORE_SIGPIPE
	CAPABILITIES_TRANSACTIONS
	CAPABILITIES_RESERVED
	CAPABILITIES_SECURE_CONNECTION
	CAPABILITIES_MULTI_STATEMENTS
	CAPABILITIES_MULTI_RESULTS
	CAPABILITIES_PS_MULTI_RESULTS
	CAPABILITIES_PLUGIN_AUTH
	CAPABILITIES_CONNECT_ATTRS
	CAPABILITIES_PLUGIN_AUTH_LENENC_CLIENT_DATA
)

const (
	CAPABILITIES_SSL_VERIFY_SERVER_CERT uint32 = 1 << (iota + 30)
	CAPABILITIES_REMEMBER_OPTIONS
)
