package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// Get raw data from cache or from server
// compare data to state
// state represents what was last alerted on

const (
	DefaultStateFile = ".state/leader_state.json"
)

// LeaderState represents all the 'stars' that have been alerted on already
type LeaderState struct {
	OldBoard  *Leaderboard
	oldRaw    []byte
	stateFile string
}

func (l *LeaderState) GetStateFile() string {
	file := DefaultStateFile
	if len(l.stateFile) > 0 {
		file = l.stateFile
	}

	return file
}

func (l *LeaderState) SetStateFile(file string) {
	l.stateFile = file
}

// LoadState reads state state file into memory, afterwards can run
// GenerateDiff to calculate differences
func (l *LeaderState) LoadState() error {
	file := l.GetStateFile()

	raw, err := os.ReadFile(file)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	l.oldRaw = raw

	board := new(Leaderboard)
	err = json.Unmarshal(raw, board)
	if err != nil {
		return err
	}
	l.OldBoard = board

	return nil
}

// SaveState will update the state file with the new data obtained from API/cache/etc
func (l *LeaderState) SaveState(board Leaderboard) error {
	file := l.GetStateFile()

	raw, err := json.MarshalIndent(board, "", "  ")
	if err != nil {
		return err
	}

	err = os.MkdirAll(filepath.Dir(file), 0755)
	if err != nil {
		return err
	}

	err = os.WriteFile(file, raw, 0766)

	return err
}

// GenerateDiff will generate a leaderboard that only contains the 'stuff that differs'
// between the old board and the new board
func (l *LeaderState) GenerateDiff(updatedBoard Leaderboard) Leaderboard {
	diff := Leaderboard{}

	if l.OldBoard == nil {
		// if the old board doesn't exist, the whole thing is the diff
		return updatedBoard
	}

	if l.OldBoard.Members == nil || len(l.OldBoard.Members) == 0 {
		// the old board (state) has zero members, so the diff should include 'everything'
		diff.Members = updatedBoard.Members
	} else {
		diff.Members = map[string]Member{}

		for memberID, member := range updatedBoard.Members {
			oMembers := l.OldBoard.Members
			if _, ok := oMembers[memberID]; !ok {
				// A new member has appeared, add the whole member to the diff
				diff.Members[memberID] = member
				continue
			}

			for day, stars := range member.CompletionDayLevel {
				// for each day, compare if its in old and new

				oStars, ok := l.OldBoard.Members[memberID].CompletionDayLevel[day]
				if !ok {
					// The entire day is new, add it all
					if _, ok := diff.Members[memberID]; !ok {
						diff.Members[memberID] = Member{
							ID:                 member.ID,
							Name:               member.Name,
							CompletionDayLevel: map[string]map[string]StarTimestamp{},
						}
					}
					diff.Members[memberID].CompletionDayLevel[day] = stars
					continue
				}

				if len(oStars) != len(stars) {
					if _, ok := diff.Members[memberID]; !ok {
						diff.Members[memberID] = Member{
							ID:                 member.ID,
							Name:               member.Name,
							CompletionDayLevel: map[string]map[string]StarTimestamp{},
						}
					}
					diff.Members[memberID].CompletionDayLevel[day] = stars
				}
			}
		}
	}

	return diff
}
