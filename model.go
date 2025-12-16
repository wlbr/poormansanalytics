package main

import "fmt"

type uuid = string

type modelStats struct {
	ID    string
	count int
}

func (m modelStats) String() string {
	return fmt.Sprintf("Model: %s  Count: %d", m.ID, m.count)
}

type projectStats struct {
	ID     uuid
	count  int
	models map[string]*modelStats
	users  map[uuid]*userStats
}

func (p projectStats) String() string {
	return fmt.Sprintf("Project: %s  Count: %d", p.ID, p.count)
}

type userStats struct {
	ID    uuid
	count int
}

func (u userStats) String() string {
	return fmt.Sprintf("User: %s  Count: %d", u.ID, u.count)
}

type stats struct {
	users    map[uuid]*userStats
	models   map[string]*modelStats
	projects map[uuid]*projectStats
}

func NewStats() *stats {
	return &stats{make(map[uuid]*userStats), make(map[string]*modelStats), make(map[uuid]*projectStats)}
}
