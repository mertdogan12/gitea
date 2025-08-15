// Copyright 2022 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package convert

import (
	"context"

	repo_model "code.gitea.io/gitea/models/repo"
	api "code.gitea.io/gitea/modules/structs"
)

// ToPushMirror convert from repo_model.PushMirror and remoteAddress to api.TopicResponse
func ToPushMirror(ctx context.Context, pm *repo_model.PushMirror) (*api.PushMirror, error) {
	repo := pm.GetRepository(ctx)
	return &api.PushMirror{
		RepoName:       repo.Name,
		RemoteName:     pm.RemoteName,
		RemoteAddress:  pm.RemoteAddress,
		CreatedUnix:    pm.CreatedUnix.AsTime(),
		LastUpdateUnix: pm.LastUpdateUnix.AsTimePtr(),
		LastError:      pm.LastError,
		Interval:       pm.Interval.String(),
		SyncOnCommit:   pm.SyncOnCommit,
	}, nil
}

// ToPullMirror convert from repo_model.Mirror and remoteAddress to api.TopicResponse
func ToPullMirror(ctx context.Context, pm *repo_model.Mirror) (*api.PullMirror, error) {
	repo := pm.GetRepository(ctx)
	return &api.PullMirror{
		RepoName:       repo.Name,
		MirrorAddress:  pm.RemoteAddress,
		Interval:       pm.Interval.String(),
		NextUpdateUnix: pm.NextUpdateUnix.AsTimePtr(),
		EnablePrune:    pm.EnablePrune,
		LFS:            pm.LFS,
		LFSEndpoint:    pm.LFSEndpoint,
	}, nil
}
