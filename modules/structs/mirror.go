// Copyright 2021 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package structs

import "time"

// CreatePushMirrorOption represents need information to create a push mirror of a repository.
type CreatePushMirrorOption struct {
	RemoteAddress  string `json:"remote_address"`
	RemoteUsername string `json:"remote_username"`
	RemotePassword string `json:"remote_password"`
	Interval       string `json:"interval"`
	SyncOnCommit   bool   `json:"sync_on_commit"`
}

// ChangePullMirrorOption represents need information to change a pull mirror of a repository.
type ChangePullMirrorOption struct {
	MirrorAddress  string `json:"mirror_address"`
	MirrorUsername string `json:"mirror_username"`
	MirrorPassword string `json:"mirror_password"`
	EnablePrune    bool   `json:"enable_prune"`
	Interval       string `json:"interval"`
	LFS            bool   `json:"lfs"`
	LFSEndpoint    string `json:"lfs_endpoint"`
}

// PushMirror represents information of a push mirror
// swagger:model
type PushMirror struct {
	RepoName      string `json:"repo_name"`
	RemoteName    string `json:"remote_name"`
	RemoteAddress string `json:"remote_address"`
	// swagger:strfmt date-time
	CreatedUnix time.Time `json:"created"`
	// swagger:strfmt date-time
	LastUpdateUnix *time.Time `json:"last_update"`
	LastError      string     `json:"last_error"`
	Interval       string     `json:"interval"`
	SyncOnCommit   bool       `json:"sync_on_commit"`
}

// PullMirror represents information of a pull mirror
// swagger:model
type PullMirror struct {
	RepoName      string `json:"repo_name"`
	MirrorAddress string `json:"mirror_address"`
	Interval      string `json:"interval"`
	EnablePrune   bool   `json:"enable_prune"`

	// swagger:strfmt date-time
	NextUpdateUnix *time.Time `json:"next_update"`

	LFS         bool   `json:"lfs"`
	LFSEndpoint string `json:"lfs_endpoint"`
}
