package main

import (
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (s *stats) record_request(user_id uuid, project_id uuid, model_id string) {
	if ms, ok := s.models[model_id]; ok {
		ms.count++
	} else {
		s.models[model_id] = &modelStats{model_id, 0}
	}

	if us, ok := s.users[user_id]; ok {
		us.count++
	} else {
		s.users[user_id] = &userStats{user_id, 0}
	}

	ps, ok := s.projects[project_id]
	if ok {
		ps.count++
	} else {
		ps = &projectStats{project_id, 0, make(map[string]*modelStats), make(map[uuid]*userStats)}
		s.projects[project_id] = ps
	}

	// add scoped references
	if projectsuser, ok := ps.users[user_id]; ok {
		projectsuser.count++
	} else {
		ps.users[user_id] = &userStats{user_id, 0}
	}

	if projectsmodel, ok := ps.models[model_id]; ok {
		projectsmodel.count++
	} else {
		ps.models[model_id] = &modelStats{model_id, 0}
	}

}

func multiplyData(n int, data []string) []string {
	var tm []string
	for i := 0; i < n; i++ {
		for _, m := range data {
			tm = append(tm, fmt.Sprintf("%s-%d", m, i))
		}
	}
	return tm
}

func (s *stats) makeTestData(n, requests int) (modelid string, userid uuid, projectid uuid) {
	testmodels := []string{"Llama", "gpt-oss", "gemini", "midjourney", "Claude", "Mistral", "DeepSeek", "Grok", "Qwen"}
	testusers := []uuid{"James", "Mary", "Robert", "Patricia", "John", "Jennifer", "Michael", "Linda", "David", "Elizabeth", "William", "Barbara", "Richard", "Susan", "Joseph", "Jessica", "Thomas", "Sarah", "Christopher", "Karen"}
	testprojects := []uuid{"Burj Khalifa", "The Great Wall of China", "Eiffel Tower", "Taj Mahal", "International Space Station", "Panama Canal", "The Apollo Program", "Sydney Opera House", "Colosseum", "Guggenheim Museum Bilbao", "Channel Tunnel", "Three Gorges Dam", "Kubernetes", "React", "VS Code", "Fallingwater", "Sagrada Familia", "Golden Gate Bridge", "TensorFlow", "The Manhattan Project"}

	testmodels = multiplyData(n, testmodels)
	testusers = multiplyData(n, testusers)
	testprojects = multiplyData(n, testprojects)

	lenm := len(testmodels)
	lenu := len(testusers)
	lenp := len(testprojects)

	for i := 0; i <= requests; i++ {
		m := rand.Intn(lenm)
		u := rand.Intn(lenu)
		p := rand.Intn(lenp)

		s.record_request(testusers[u], testprojects[p], testmodels[m])
	}
	return testmodels[0], testusers[0], testprojects[0]
}

func main() {
	limitrequests := 10 * 1000000
	tops := 900
	s := NewStats()
	timeStart := time.Now()
	_, _, p := s.makeTestData(100, limitrequests)
	timeCreated := time.Now()

	//dummy to ensure jit
	for _, v := range s.get_top_projects(1) {
		fmt.Println(v)
	}

	timeCreated2 := time.Now()
	for _, project := range s.get_top_projects(tops) {
		fmt.Println(project)
	}
	timeTopProjects := time.Now()

	for _, user := range s.get_top_users(tops) {
		fmt.Println(user)
	}
	timeTopUsers := time.Now()

	for _, model := range s.get_top_models(tops) {
		fmt.Println(model)
	}
	timeTopModels := time.Now()

	fmt.Println("-----Top Models in ", p)
	for _, model := range s.get_top_models_in_project(p, tops) {
		fmt.Println(model)
	}
	timeTopModelsInProject := time.Now()

	fmt.Println("-----Top User in ", p)
	for _, users := range s.get_top_users_in_project(p, tops) {
		fmt.Println(users)
	}
	timeTopUsersInProject := time.Now()

	prn := message.NewPrinter(language.English)
	prn.Printf("\n-----Timings - %d Projects  %d Models  %d Users  - %d Requests \n", len(s.projects), len(s.models), len(s.users), limitrequests)
	fmt.Println("Create time:", timeCreated.Sub(timeStart))
	fmt.Println("TopProjects time:", timeTopProjects.Sub(timeCreated2))
	fmt.Println("TopUsers time:", timeTopUsers.Sub(timeTopProjects))
	fmt.Println("ToModels time:", timeTopModels.Sub(timeTopProjects))
	fmt.Println("ToModelsInProject time:", timeTopModelsInProject.Sub(timeTopModels))
	fmt.Println("TopUsersInProject time:", timeTopUsersInProject.Sub(timeTopModelsInProject))
}
