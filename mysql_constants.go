package binlog

// These offset constants are based on v4 events
// sadly, golang doesn't support constant arrays (because of slices I think)
var BINLOG_MAGIC = [4]byte{0xfe, 0x62, 0x69, 0x6e}

const MAGIC_BYTES_LENGTH  int = 4
const EVENT_HEADER_LENGTH int = 19

// This is only int64 to save time on casting
// (hint: it probably shouldn't be int64)
const (
	EVENT_TYPE_OFFSET      int64 = 4
	EVENT_SERVER_ID_OFFSET       = 5
	EVENT_LEN_OFFSET             = 9
	EVENT_NEXT_OFFSET            = 13
	EVENT_FLAGS_OFFSET           = 17
	EVENT_EXTRA_OFFSET           = 19
)

type MysqlBinlogEventType byte

const (
	UNKOWN_EVENT MysqlBinlogEventType = iota
	START_EVENT_V3
	QUERY_EVENT
	STOP_EVENT
	ROTATE_EVENT
	INTVAR_EVENT
	LOAD_EVENT
	SLAVE_EVENT
	CREATE_FILE_EVENT
	APPEND_BLOCK_EVENT
	EXEC_LOAD_EVENT
	DELETE_FILE_EVENT
	NEW_LOAD_EVENT
	RAID_EVENT
	USER_VAR_EVENT
	FORMAT_DESCRIPTION_EVENT
	XID_EVENT
	BEGIN_LOAD_QUERY_EVENT
	EXECUTE_LOAD_QUERY_EVENT
	TABLE_MAP_EVENT
	WRITE_ROWS_EVENTv0
	UPDATE_ROWS_EVENTv0
	DELETE_ROWS_EVENTv0
	WRITE_ROWS_EVENTv1
	UPDATE_ROWS_EVENTv1
	DELETE_ROWS_EVENTv1
	INCIDENT_EVENT
	HEARTBEAT_EVENT
	IGNORABLE_EVENT
	ROWS_QUERY_EVENT
	WRITE_ROWS_EVENTv2
	UPDATE_ROWS_EVENTv2
	DELETE_ROWS_EVENTv2
	GTID_EVENT
	ANONYMOUS_GTID_EVENT
	PREVIOUS_GTIDS_EVENT
)

func (mbe MysqlBinlogEventType) String() string {
	switch mbe {
	case UNKOWN_EVENT:
		return "UNKOWN_EVENT"
	case START_EVENT_V3:
		return "START_EVENT_V3"
	case QUERY_EVENT:
		return "QUERY_EVENT"
	case STOP_EVENT:
		return "STOP_EVENT"
	case ROTATE_EVENT:
		return "ROTATE_EVENT"
	case INTVAR_EVENT:
		return "INTVAR_EVENT"
	case LOAD_EVENT:
		return "LOAD_EVENT"
	case SLAVE_EVENT:
		return "SLAVE_EVENT"
	case CREATE_FILE_EVENT:
		return "CREATE_FILE_EVENT"
	case APPEND_BLOCK_EVENT:
		return "APPEND_BLOCK_EVENT"
	case EXEC_LOAD_EVENT:
		return "EXEC_LOAD_EVENT"
	case DELETE_FILE_EVENT:
		return "DELETE_FILE_EVENT"
	case NEW_LOAD_EVENT:
		return "NEW_LOAD_EVENT"
	case RAID_EVENT:
		return "RAID_EVENT"
	case USER_VAR_EVENT:
		return "USER_VAR_EVENT"
	case FORMAT_DESCRIPTION_EVENT:
		return "FORMAT_DESCRIPTION_EVENT"
	case XID_EVENT:
		return "XID_EVENT"
	case BEGIN_LOAD_QUERY_EVENT:
		return "BEGIN_LOAD_QUERY_EVENT"
	case EXECUTE_LOAD_QUERY_EVENT:
		return "EXECUTE_LOAD_QUERY_EVENT"
	case TABLE_MAP_EVENT:
		return "TABLE_MAP_EVENT"
	case WRITE_ROWS_EVENTv0:
		return "WRITE_ROWS_EVENTv0"
	case UPDATE_ROWS_EVENTv0:
		return "UPDATE_ROWS_EVENTv0"
	case DELETE_ROWS_EVENTv0:
		return "DELETE_ROWS_EVENTv0"
	case WRITE_ROWS_EVENTv1:
		return "WRITE_ROWS_EVENTv1"
	case UPDATE_ROWS_EVENTv1:
		return "UPDATE_ROWS_EVENTv1"
	case DELETE_ROWS_EVENTv1:
		return "DELETE_ROWS_EVENTv1"
	case INCIDENT_EVENT:
		return "INCIDENT_EVENT"
	case HEARTBEAT_EVENT:
		return "HEARTBEAT_EVENT"
	case IGNORABLE_EVENT:
		return "IGNORABLE_EVENT"
	case ROWS_QUERY_EVENT:
		return "ROWS_QUERY_EVENT"
	case WRITE_ROWS_EVENTv2:
		return "WRITE_ROWS_EVENTv2"
	case UPDATE_ROWS_EVENTv2:
		return "UPDATE_ROWS_EVENTv2"
	case DELETE_ROWS_EVENTv2:
		return "DELETE_ROWS_EVENTv2"
	case GTID_EVENT:
		return "GTID_EVENT"
	case ANONYMOUS_GTID_EVENT:
		return "ANONYMOUS_GTID_EVENT"
	case PREVIOUS_GTIDS_EVENT:
		return "PREVIOUS_GTIDS_EVENT"
	}

	return "invalid"
}

type MysqlType byte

const (
	MYSQL_TYPE_DECIMAL MysqlType = iota
	MYSQL_TYPE_TINY
	MYSQL_TYPE_SHORT
	MYSQL_TYPE_LONG
	MYSQL_TYPE_FLOAT
	MYSQL_TYPE_DOUBLE
	MYSQL_TYPE_NULL
	MYSQL_TYPE_TIMESTAMP
	MYSQL_TYPE_LONGLONG
	MYSQL_TYPE_INT24
	MYSQL_TYPE_DATE
	MYSQL_TYPE_TIME
	MYSQL_TYPE_DATETIME
	MYSQL_TYPE_YEAR
	MYSQL_TYPE_NEWDATE // Does not appear in binlog
	MYSQL_TYPE_VARCHAR
	MYSQL_TYPE_BIT
	MYSQL_TYPE_TIMESTAMP_V2
	MYSQL_TYPE_DATETIME_V2
	MYSQL_TYPE_TIME_V2
)

const (
	MYSQL_TYPE_NEWDECIMAL  MysqlType = 246 + iota
	MYSQL_TYPE_ENUM                  // Does not appear in binlog
	MYSQL_TYPE_SET                   // Does not appear in binlog
	MYSQL_TYPE_TINY_BLOB             // Does not appear in binlog
	MYSQL_TYPE_MEDIUM_BLOB           // Does not appear in binlog
	MYSQL_TYPE_LONG_BLOB             // Does not appear in binlog
	MYSQL_TYPE_BLOB
	MYSQL_TYPE_VAR_STRING
	MYSQL_TYPE_STRING
	MYSQL_TYPE_GEOMETRY
)

func (mt MysqlType) String() string {
	switch mt {
	case MYSQL_TYPE_DECIMAL:
		return "MYSQL_TYPE_DECIMAL"
	case MYSQL_TYPE_TINY:
		return "MYSQL_TYPE_TINY"
	case MYSQL_TYPE_SHORT:
		return "MYSQL_TYPE_SHORT"
	case MYSQL_TYPE_LONG:
		return "MYSQL_TYPE_LONG"
	case MYSQL_TYPE_FLOAT:
		return "MYSQL_TYPE_FLOAT"
	case MYSQL_TYPE_DOUBLE:
		return "MYSQL_TYPE_DOUBLE"
	case MYSQL_TYPE_NULL:
		return "MYSQL_TYPE_NULL"
	case MYSQL_TYPE_TIMESTAMP:
		return "MYSQL_TYPE_TIMESTAMP"
	case MYSQL_TYPE_LONGLONG:
		return "MYSQL_TYPE_LONGLONG"
	case MYSQL_TYPE_INT24:
		return "MYSQL_TYPE_INT24"
	case MYSQL_TYPE_DATE:
		return "MYSQL_TYPE_DATE"
	case MYSQL_TYPE_TIME:
		return "MYSQL_TYPE_TIME"
	case MYSQL_TYPE_DATETIME:
		return "MYSQL_TYPE_DATETIME"
	case MYSQL_TYPE_YEAR:
		return "MYSQL_TYPE_YEAR"
	case MYSQL_TYPE_NEWDATE:
		return "MYSQL_TYPE_NEWDATE"
	case MYSQL_TYPE_VARCHAR:
		return "MYSQL_TYPE_VARCHAR"
	case MYSQL_TYPE_BIT:
		return "MYSQL_TYPE_BIT"
	case MYSQL_TYPE_TIMESTAMP_V2:
		return "MYSQL_TYPE_TIMESTAMP_V2"
	case MYSQL_TYPE_DATETIME_V2:
		return "MYSQL_TYPE_DATETIME_V2"
	case MYSQL_TYPE_TIME_V2:
		return "MYSQL_TYPE_TIME_V2"
	case MYSQL_TYPE_NEWDECIMAL:
		return "MYSQL_TYPE_NEWDECIMAL"
	case MYSQL_TYPE_ENUM:
		return "MYSQL_TYPE_ENUM"
	case MYSQL_TYPE_SET:
		return "MYSQL_TYPE_SET"
	case MYSQL_TYPE_TINY_BLOB:
		return "MYSQL_TYPE_TINY_BLOB"
	case MYSQL_TYPE_MEDIUM_BLOB:
		return "MYSQL_TYPE_MEDIUM_BLOB"
	case MYSQL_TYPE_LONG_BLOB:
		return "MYSQL_TYPE_LONG_BLOB"
	case MYSQL_TYPE_BLOB:
		return "MYSQL_TYPE_BLOB"
	case MYSQL_TYPE_VAR_STRING:
		return "MYSQL_TYPE_VAR_STRING"
	case MYSQL_TYPE_STRING:
		return "MYSQL_TYPE_STRING"
	case MYSQL_TYPE_GEOMETRY:
		return "MYSQL_TYPE_GEOMETRY"
	}

	return "invalid"
}

func checkBinlogMagic(magic []byte) bool {
	if len(magic) != len(BINLOG_MAGIC) {
		return false
	}

	for i, b := range magic {
		if b != BINLOG_MAGIC[i] {
			return false
		}
	}

	return true
}
