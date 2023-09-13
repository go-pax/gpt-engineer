package github

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-pax/gpt-engineer/database"
	"golang.org/x/oauth2"
	"io"
	"net/http"
	nurl "net/url"
	"path"
	"strings"

	"github.com/google/go-github/v39/github"
)

func init() {
	database.Register("github", &Github{})
}

var (
	ErrNoUserInfo  = fmt.Errorf("no username:token provided")
	ErrNoAccess    = fmt.Errorf("personal access token needed")
	ErrInvalidRepo = fmt.Errorf("invalid repo")
	ErrInvalidRef  = fmt.Errorf("invalid ref")
)

type Github struct {
	hasUser      bool
	config       *Config
	client       *github.Client
	options      *github.RepositoryContentGetOptions
	subDirectory string
}

type Config struct {
	Owner string
	Repo  string
	Path  string
	Ref   string
}

func (g *Github) Open(url string, subDirectory string) (database.Database, error) {
	u, err := nurl.Parse(url)
	if err != nil {
		return nil, err
	}

	// client defaults to http.DefaultClient
	var client *http.Client
	if u.User != nil {
		password, ok := u.User.Password()
		if !ok {
			return nil, ErrNoUserInfo
		}
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: password},
		)
		client = oauth2.NewClient(context.Background(), ts)
	}

	gn := &Github{
		hasUser:      u.User != nil,
		client:       github.NewClient(client),
		subDirectory: subDirectory,
		options:      &github.RepositoryContentGetOptions{Ref: u.Fragment},
	}

	gn.ensureFields()

	// set owner, repo and path in repo
	gn.config.Owner = u.Host
	pe := strings.Split(strings.Trim(u.Path, "/"), "/")
	if len(pe) < 1 {
		return nil, ErrInvalidRepo
	}
	gn.config.Repo = pe[0]
	if len(pe) > 1 {
		gn.config.Path = strings.Join(pe[1:], "/")
	}

	if subDirectory != "" {
		gn.config.Path = path.Join(gn.config.Path, subDirectory)
	}

	if err := gn.ensureRepo(); err != nil {
		return nil, err
	}

	return gn, nil
}

func (g *Github) ensureRepo() error {
	g.ensureFields()

	if _, _, err := g.client.Repositories.Get(
		context.Background(),
		g.config.Owner,
		g.config.Repo,
	); err != nil {
		// create repo
		visibility := "internal"
		if _, _, err = g.client.Repositories.Create(context.Background(), g.config.Owner, &github.Repository{
			Name:       &g.config.Repo,
			Visibility: &visibility,
		}); err != nil {
			return err
		}
	}
	return nil
}

func (g *Github) ensureFields() {
	if g.config == nil {
		g.config = &Config{}
	}
}

func (g *Github) isAuthenticated() bool {
	return g.hasUser
}

func (g *Github) Close() error {
	return nil
}

func (g *Github) Path() string {
	return g.config.Path
}

func (g *Github) Get(file string) (contents string, err error) {
	g.ensureFields()

	r, _, err := g.client.Repositories.DownloadContents(
		context.Background(),
		g.config.Owner,
		g.config.Repo,
		path.Join(g.config.Path, file),
		g.options,
	)

	if err != nil {
		var er *github.ErrorResponse
		if errors.As(err, &er) {
			if er.Response.StatusCode == 404 {
				return "", ErrInvalidRef
			}
		}
		return "", err
	}
	defer func() {
		if err := r.Close(); err != nil {
			fmt.Printf("failed to close reader, %v\n", err)
		}
	}()

	contents = ""
	data := make([]byte, 4096)
	for {
		n, err := r.Read(data)
		if err != nil {
			if errors.Is(err, io.EOF) {
				contents += string(data[:n])
				break
			}
			return "", err
		}
		contents += string(data[:n])
	}
	return contents, nil
}

func (g *Github) Set(file, contents string) error {
	if !g.isAuthenticated() {
		return ErrNoAccess
	}

	fileContent, _, _, _ := g.client.Repositories.GetContents(
		context.Background(),
		g.config.Owner,
		g.config.Repo,
		path.Join(g.config.Path, file),
		g.options,
	)

	options := &github.RepositoryContentFileOptions{
		Message: github.String(fmt.Sprintf("committing file, %s", file)),
		Content: []byte(contents),
		SHA:     nil,
		Branch:  &g.options.Ref,
	}

	fn := g.client.Repositories.CreateFile
	if fileContent != nil {
		fn = g.client.Repositories.UpdateFile
		options.SHA = fileContent.SHA
	}

	_, _, err := fn(
		context.Background(),
		g.config.Owner,
		g.config.Repo,
		path.Join(g.config.Path, file),
		options,
	)
	return err
}
