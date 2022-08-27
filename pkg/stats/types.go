package stats

type Stats struct {
	Type            string  `json:"message_type"`
	FilesNew        int     `json:"files_new" push:"files_new"`
	FilesChanged    int     `json:"files_changed" push:"files_changed"`
	FilesUnmodified int     `json:"files_unmodified" push:"files_unmodified"`
	DirsNew         int     `json:"dirs_new" push:"dirs_new"`
	DirsChanged     int     `json:"dirs_changed" push:"dirs_changed"`
	DirsUnmodified  int     `json:"dirs_unmodified" push:"dirs_unmodified"`
	DataBlobs       int     `json:"data_blobs" push:"data_blobs"`
	TreeBlobs       int     `json:"tree_blobs" push:"tree_blobs"`
	DataAdded       int     `json:"data_added" push:"data_added"`
	FilesProcessed  int     `json:"total_files_processed" push:"total_files_processed"`
	BytesProcessed  int     `json:"total_bytes_processed" push:"total_bytes_processed"`
	Duration        float64 `json:"total_duration" push:"duration"`
	SnapshotID      string  `json:"snapshot_id"`
}
