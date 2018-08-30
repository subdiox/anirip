package crunchyroll

import "github.com/subdiox/anirip/common"

// Season contains season metadata and child episodes
type Season struct {
	Title    string
	Number   int
	Length   int
	Episodes []Episode
}

// GetNumber returns the season number that will be used for folder naming
func (s *Season) GetNumber() int {
	return s.Number
}

// GetEpisodes copies the episodes on the Season and returns them as an
// common.Episodes
func (s *Season) GetEpisodes() common.Episodes {
	episodes := []common.Episode{}
	for i := 0; i < len(s.Episodes); i++ {
		episodes = append(episodes, &s.Episodes[i])
	}
	return episodes
}
