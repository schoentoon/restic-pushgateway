package stats

import (
	"encoding/json"
	"io"
)

//{"message_type":"summary","files_new":0,"files_changed":11,"files_unmodified":6926,"dirs_new":0,"dirs_changed":15,"dirs_unmodified":614,"data_blobs":8,"tree_blobs":15,
//"data_added":683689,"total_files_processed":6937,"total_bytes_processed":10189155310,"total_duration":1.125009677,"snapshot_id":"216ebdee"}

func ParseOutput(r io.Reader) (*Stats, error) {
	decoder := json.NewDecoder(r)

	for {
		var stats Stats
		err := decoder.Decode(&stats)
		if err != nil {
			return nil, err
		}

		// we really only care about the summary at the end, so with anything that isn't the summary we just continue
		if stats.Type == "summary" {
			return &stats, nil
		}
	}
}
