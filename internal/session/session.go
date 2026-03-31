package session

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"time"
)

type State struct {
	EndTime     time.Time `json:"end_time"`
	DurationMin int       `json:"duration_min"`
	Sites       []string  `json:"sites"`
	Stopping    bool      `json:"stopping"`
	StopAt      time.Time `json:"stop_at,omitempty"`
}

var stateFile = filepath.Join(os.Getenv("APPDATA"), "focus", "session.json")
var sitesFile = filepath.Join(os.Getenv("APPDATA"), "focus", "sites.json")

func init() {
	dir := filepath.Dir(stateFile)
	_ = os.MkdirAll(dir, 0755)
}

func Load() (*State, error) {
	data, err := os.ReadFile(stateFile)
	if err != nil {
		return nil, errors.New("no session file")
	}
	var s State
	if err := json.Unmarshal(data, &s); err != nil {
		return nil, err
	}
	return &s, nil
}

func Save(s *State) error {
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(stateFile, data, 0644)
}

func Clear() {
	_ = os.Remove(stateFile)
}

func IsActive() bool {
	s, err := Load()
	if err != nil {
		return false
	}
	return s.Remaining() > 0 && !s.Stopping
}

func (s *State) Remaining() time.Duration {
	return time.Until(s.EndTime)
}

func (s *State) EndTimeMs() int64 {
	return s.EndTime.UnixMilli()
}

func ForceStop() error {
	s, err := Load()
	if err != nil {
		return errors.New("no active session")
	}
	_ = s
	Clear()
	return nil
}

func LoadSites() []string {
	data, err := os.ReadFile(sitesFile)
	if err != nil {
		return defaultSites()
	}
	var sites []string
	if err := json.Unmarshal(data, &sites); err != nil {
		return defaultSites()
	}
	return sites
}

func SaveSites(sites []string) error {
	data, err := json.MarshalIndent(sites, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(sitesFile, data, 0644)
}

func defaultSites() []string {
	return []string{
		"facebook.com",
		"instagram.com",
		"x.com",
		"tiktok.com",
		"reddit.com",
		"youtube.com",
		"twitter.com",
	}
}
