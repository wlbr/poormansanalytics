package main

import (
	"maps"
	"slices"
)

func (s *stats) get_top_projects(n int) []*projectStats {
	vals := slices.Collect(maps.Values(s.projects))
	slices.SortFunc(vals, func(a, b *projectStats) int {
		if a.count-b.count < 0 {
			return 1
		} else if a.count-b.count > 0 {
			return -1
		} else {
			return 0
		}
	})

	return vals[0:min(n, len(vals))]
}

func (s *stats) get_top_users(n int) []*userStats {
	vals := slices.Collect(maps.Values(s.users))
	slices.SortFunc(vals, func(a, b *userStats) int {
		if a.count-b.count < 0 {
			return 1
		} else if a.count-b.count > 0 {
			return -1
		} else {
			return 0
		}
	})

	return vals[0:min(n, len(vals))]
}

func (s *stats) get_top_models(n int) []*modelStats {
	vals := slices.Collect(maps.Values(s.models))
	slices.SortFunc(vals, func(a, b *modelStats) int {
		if a.count-b.count < 0 {
			return 1
		} else if a.count-b.count > 0 {
			return -1
		} else {
			return 0
		}
	})

	return vals[0:min(n, len(vals))]
}

func (s *stats) get_top_models_in_project(project_id uuid, n int) []*modelStats {
	ps := s.projects[uuid(project_id)]
	vals := slices.Collect(maps.Values(ps.models))
	slices.SortFunc(vals, func(a, b *modelStats) int {
		if a.count-b.count < 0 {
			return 1
		} else if a.count-b.count > 0 {
			return -1
		} else {
			return 0
		}
	})

	return vals[0:min(n, len(vals))]
}

func (s *stats) get_top_users_in_project(project_id uuid, n int) []*userStats {
	ps := s.projects[uuid(project_id)]
	vals := slices.Collect(maps.Values(ps.users))
	slices.SortFunc(vals, func(a, b *userStats) int {
		if a.count-b.count < 0 {
			return 1
		} else if a.count-b.count > 0 {
			return -1
		} else {
			return 0
		}
	})

	return vals[0:min(n, len(vals))]
}
