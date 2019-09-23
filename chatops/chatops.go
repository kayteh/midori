package chatops

import (
	"context"
	"net/http"

	"github.com/bradleyfalzon/ghinstallation"
	"github.com/die-net/lrucache"
	"github.com/google/go-github/github"
	"github.com/gregjones/httpcache"
	"golang.org/x/oauth2"
)

type ChatOpsProvider struct {
	ghClient   *github.Client
	cfg        *ChatOpsConfig
	httpClient *http.Client
}

type ChatOpsConfig struct {
	AppID       int
	AppKeyPath  string
	AccessToken string
}

var (
	httpLRU = lrucache.New(1e+8, 60*60) // 100MB, 1H
)

func NewChatOps(cfg *ChatOpsConfig) (*ChatOpsProvider, error) {
	tokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: cfg.AccessToken},
	)
	tokenClient := oauth2.NewClient(context.Background(), tokenSource)
	tokenClient.Transport = &httpcache.Transport{
		Cache: httpLRU,
	}

	client := github.NewClient(tokenClient)

	_, _, err := client.APIMeta(context.Background())
	if err != nil {
		return nil, err
	}

	return &ChatOpsProvider{
		ghClient:   client,
		cfg:        cfg,
		httpClient: tokenClient,
	}, nil
}

func (co *ChatOpsProvider) getInstallationID(owner string, repo string) (int, error) {
	installation, _, err := co.ghClient.Apps.FindRepositoryInstallation(context.Background(), owner, repo)
	if err != nil {
		return 0, err
	}

	return int(installation.GetID()), nil
}

type InstallationClient struct {
	co         *ChatOpsProvider
	instClient *github.Client
	ghClient   *github.Client
}

func (co *ChatOpsProvider) NewInstallationClient(owner string, repo string) (*InstallationClient, error) {
	instID, err := co.getInstallationID(owner, repo)
	if err != nil {
		return nil, err
	}

	tport, err := ghinstallation.NewKeyFromFile(&httpcache.Transport{
		Cache: httpLRU,
	}, co.cfg.AppID, instID, co.cfg.AppKeyPath)
	if err != nil {
		return nil, err
	}

	client := github.NewClient(&http.Client{Transport: tport})

	return &InstallationClient{
		co:         co,
		instClient: client,
		ghClient:   co.ghClient,
	}, nil
}
