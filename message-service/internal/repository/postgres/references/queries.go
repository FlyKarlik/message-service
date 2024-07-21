package references

// message queries
const (
	AddMessageQuery         = "INSERT INTO message (content) VALUES ($1) RETURNING id"
	GetMessageQuery         = "SELECT id,content,status,created_at,processed_at FROM message WHERE id=$1"
	GetAllMessageQuery      = "SELECT id,content,status,created_at,processed_at FROM message"
	GetAllProcessedMsgQuery = "SELECT id,content,status,created_at,processed_at FROM message WHERE status=0"
	UpdateStatusMsgQuery    = "UPDATE message SET status=0,processed_at=$1 WHERE id=$2"
)

const (
	GetStatsQuery = "SELECT processed_count,last_processed_message,last_update FROM statistics"
)
