package clients

import (
	msgraph "github.com/yaegashi/msgraph.go/v1.0"
)

type Client struct {
	MailFolders *msgraph.UserMailFoldersCollectionRequestBuilder
}

func NewClient(b msgraph.BaseRequestBuilder) *Client {
	b.SetURL(b.URL() + "/me")
	userClient := msgraph.UserRequestBuilder{b}
	return &Client{
		MailFolders: userClient.MailFolders(),
	}
}
