package main

import (
	"context"

	"github.com/go-playground/log/v8"
	"github.com/google/go-github/github"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	cLog := new(CustomHandler)
	log.AddHandler(cLog, log.AllLevels...)

}

// GetGithubRepositories per page the given username
func (a *App) GetGithubRepositories(username string, page int) []*github.Repository {
	client := github.NewClient(nil)

	repos, _, err := client.Repositories.List(context.Background(), username, &github.RepositoryListOptions{
		Sort: "updated",
		ListOptions: github.ListOptions{
			PerPage: 100,
			Page:    page,
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Info(repos)

	return repos
}
