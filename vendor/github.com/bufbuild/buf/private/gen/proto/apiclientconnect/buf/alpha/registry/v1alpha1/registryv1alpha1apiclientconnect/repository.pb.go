// Copyright 2020-2022 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-apiclientconnect. DO NOT EDIT.

package registryv1alpha1apiclientconnect

import (
	context "context"
	registryv1alpha1connect "github.com/bufbuild/buf/private/gen/proto/connect/buf/alpha/registry/v1alpha1/registryv1alpha1connect"
	v1alpha1 "github.com/bufbuild/buf/private/gen/proto/go/buf/alpha/registry/v1alpha1"
	connect_go "github.com/bufbuild/connect-go"
	zap "go.uber.org/zap"
)

type repositoryServiceClient struct {
	logger          *zap.Logger
	client          registryv1alpha1connect.RepositoryServiceClient
	contextModifier func(context.Context) context.Context
}

// GetRepository gets a repository by ID.
func (s *repositoryServiceClient) GetRepository(
	ctx context.Context,
	id string,
) (repository *v1alpha1.Repository, counts *v1alpha1.RepositoryCounts, _ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	response, err := s.client.GetRepository(
		ctx,
		connect_go.NewRequest(
			&v1alpha1.GetRepositoryRequest{
				Id: id,
			}),
	)
	if err != nil {
		return nil, nil, err
	}
	return response.Msg.Repository, response.Msg.Counts, nil
}

// GetRepositoryByFullName gets a repository by full name.
func (s *repositoryServiceClient) GetRepositoryByFullName(
	ctx context.Context,
	fullName string,
) (repository *v1alpha1.Repository, counts *v1alpha1.RepositoryCounts, _ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	response, err := s.client.GetRepositoryByFullName(
		ctx,
		connect_go.NewRequest(
			&v1alpha1.GetRepositoryByFullNameRequest{
				FullName: fullName,
			}),
	)
	if err != nil {
		return nil, nil, err
	}
	return response.Msg.Repository, response.Msg.Counts, nil
}

// ListRepositories lists all repositories.
func (s *repositoryServiceClient) ListRepositories(
	ctx context.Context,
	pageSize uint32,
	pageToken string,
	reverse bool,
) (repositories []*v1alpha1.Repository, nextPageToken string, _ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	response, err := s.client.ListRepositories(
		ctx,
		connect_go.NewRequest(
			&v1alpha1.ListRepositoriesRequest{
				PageSize:  pageSize,
				PageToken: pageToken,
				Reverse:   reverse,
			}),
	)
	if err != nil {
		return nil, "", err
	}
	return response.Msg.Repositories, response.Msg.NextPageToken, nil
}

// ListUserRepositories lists all repositories belonging to a user.
func (s *repositoryServiceClient) ListUserRepositories(
	ctx context.Context,
	userId string,
	pageSize uint32,
	pageToken string,
	reverse bool,
) (repositories []*v1alpha1.Repository, nextPageToken string, _ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	response, err := s.client.ListUserRepositories(
		ctx,
		connect_go.NewRequest(
			&v1alpha1.ListUserRepositoriesRequest{
				UserId:    userId,
				PageSize:  pageSize,
				PageToken: pageToken,
				Reverse:   reverse,
			}),
	)
	if err != nil {
		return nil, "", err
	}
	return response.Msg.Repositories, response.Msg.NextPageToken, nil
}

// ListRepositoriesUserCanAccess lists all repositories a user can access.
func (s *repositoryServiceClient) ListRepositoriesUserCanAccess(
	ctx context.Context,
	pageSize uint32,
	pageToken string,
	reverse bool,
) (repositories []*v1alpha1.Repository, nextPageToken string, _ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	response, err := s.client.ListRepositoriesUserCanAccess(
		ctx,
		connect_go.NewRequest(
			&v1alpha1.ListRepositoriesUserCanAccessRequest{
				PageSize:  pageSize,
				PageToken: pageToken,
				Reverse:   reverse,
			}),
	)
	if err != nil {
		return nil, "", err
	}
	return response.Msg.Repositories, response.Msg.NextPageToken, nil
}

// ListOrganizationRepositories lists all repositories for an organization.
func (s *repositoryServiceClient) ListOrganizationRepositories(
	ctx context.Context,
	organizationId string,
	pageSize uint32,
	pageToken string,
	reverse bool,
) (repositories []*v1alpha1.Repository, nextPageToken string, _ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	response, err := s.client.ListOrganizationRepositories(
		ctx,
		connect_go.NewRequest(
			&v1alpha1.ListOrganizationRepositoriesRequest{
				OrganizationId: organizationId,
				PageSize:       pageSize,
				PageToken:      pageToken,
				Reverse:        reverse,
			}),
	)
	if err != nil {
		return nil, "", err
	}
	return response.Msg.Repositories, response.Msg.NextPageToken, nil
}

// CreateRepositoryByFullName creates a new repository by full name.
func (s *repositoryServiceClient) CreateRepositoryByFullName(
	ctx context.Context,
	fullName string,
	visibility v1alpha1.Visibility,
) (repository *v1alpha1.Repository, _ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	response, err := s.client.CreateRepositoryByFullName(
		ctx,
		connect_go.NewRequest(
			&v1alpha1.CreateRepositoryByFullNameRequest{
				FullName:   fullName,
				Visibility: visibility,
			}),
	)
	if err != nil {
		return nil, err
	}
	return response.Msg.Repository, nil
}

// DeleteRepository deletes a repository.
func (s *repositoryServiceClient) DeleteRepository(ctx context.Context, id string) (_ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	_, err := s.client.DeleteRepository(
		ctx,
		connect_go.NewRequest(
			&v1alpha1.DeleteRepositoryRequest{
				Id: id,
			}),
	)
	if err != nil {
		return err
	}
	return nil
}

// DeleteRepositoryByFullName deletes a repository by full name.
func (s *repositoryServiceClient) DeleteRepositoryByFullName(ctx context.Context, fullName string) (_ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	_, err := s.client.DeleteRepositoryByFullName(
		ctx,
		connect_go.NewRequest(
			&v1alpha1.DeleteRepositoryByFullNameRequest{
				FullName: fullName,
			}),
	)
	if err != nil {
		return err
	}
	return nil
}

// DeprecateRepositoryByName deprecates the repository.
func (s *repositoryServiceClient) DeprecateRepositoryByName(
	ctx context.Context,
	ownerName string,
	repositoryName string,
	deprecationMessage string,
) (repository *v1alpha1.Repository, _ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	response, err := s.client.DeprecateRepositoryByName(
		ctx,
		connect_go.NewRequest(
			&v1alpha1.DeprecateRepositoryByNameRequest{
				OwnerName:          ownerName,
				RepositoryName:     repositoryName,
				DeprecationMessage: deprecationMessage,
			}),
	)
	if err != nil {
		return nil, err
	}
	return response.Msg.Repository, nil
}

// UndeprecateRepositoryByName makes the repository not deprecated and removes any deprecation_message.
func (s *repositoryServiceClient) UndeprecateRepositoryByName(
	ctx context.Context,
	ownerName string,
	repositoryName string,
) (repository *v1alpha1.Repository, _ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	response, err := s.client.UndeprecateRepositoryByName(
		ctx,
		connect_go.NewRequest(
			&v1alpha1.UndeprecateRepositoryByNameRequest{
				OwnerName:      ownerName,
				RepositoryName: repositoryName,
			}),
	)
	if err != nil {
		return nil, err
	}
	return response.Msg.Repository, nil
}

// GetRepositoriesByFullName gets repositories by full name. Response order is unspecified.
// Errors if any of the repositories don't exist or the caller does not have access to any of the repositories.
func (s *repositoryServiceClient) GetRepositoriesByFullName(ctx context.Context, fullNames []string) (repositories []*v1alpha1.Repository, _ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	response, err := s.client.GetRepositoriesByFullName(
		ctx,
		connect_go.NewRequest(
			&v1alpha1.GetRepositoriesByFullNameRequest{
				FullNames: fullNames,
			}),
	)
	if err != nil {
		return nil, err
	}
	return response.Msg.Repositories, nil
}

// SetRepositoryContributor sets the role of a user in the repository.
func (s *repositoryServiceClient) SetRepositoryContributor(
	ctx context.Context,
	repositoryId string,
	userId string,
	repositoryRole v1alpha1.RepositoryRole,
) (_ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	_, err := s.client.SetRepositoryContributor(
		ctx,
		connect_go.NewRequest(
			&v1alpha1.SetRepositoryContributorRequest{
				RepositoryId:   repositoryId,
				UserId:         userId,
				RepositoryRole: repositoryRole,
			}),
	)
	if err != nil {
		return err
	}
	return nil
}

// ListRepositoryContributors returns the list of contributors that has an explicit role against the repository.
// This does not include users who have implicit roles against the repository, unless they have also been
// assigned a role explicitly.
func (s *repositoryServiceClient) ListRepositoryContributors(
	ctx context.Context,
	repositoryId string,
	pageSize uint32,
	pageToken string,
	reverse bool,
) (users []*v1alpha1.RepositoryContributor, nextPageToken string, _ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	response, err := s.client.ListRepositoryContributors(
		ctx,
		connect_go.NewRequest(
			&v1alpha1.ListRepositoryContributorsRequest{
				RepositoryId: repositoryId,
				PageSize:     pageSize,
				PageToken:    pageToken,
				Reverse:      reverse,
			}),
	)
	if err != nil {
		return nil, "", err
	}
	return response.Msg.Users, response.Msg.NextPageToken, nil
}

// GetRepositoryContributor returns the contributor information of a user in a repository.
func (s *repositoryServiceClient) GetRepositoryContributor(
	ctx context.Context,
	repositoryId string,
	userId string,
) (user *v1alpha1.RepositoryContributor, _ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	response, err := s.client.GetRepositoryContributor(
		ctx,
		connect_go.NewRequest(
			&v1alpha1.GetRepositoryContributorRequest{
				RepositoryId: repositoryId,
				UserId:       userId,
			}),
	)
	if err != nil {
		return nil, err
	}
	return response.Msg.User, nil
}

// GetRepositorySettings gets the settings of a repository.
func (s *repositoryServiceClient) GetRepositorySettings(ctx context.Context, repositoryId string) (contributorsCount uint32, _ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	response, err := s.client.GetRepositorySettings(
		ctx,
		connect_go.NewRequest(
			&v1alpha1.GetRepositorySettingsRequest{
				RepositoryId: repositoryId,
			}),
	)
	if err != nil {
		return 0, err
	}
	return response.Msg.ContributorsCount, nil
}

// UpdateRepositorySettingsByName updates the settings of a repository.
func (s *repositoryServiceClient) UpdateRepositorySettingsByName(
	ctx context.Context,
	ownerName string,
	repositoryName string,
	visibility v1alpha1.Visibility,
) (_ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	_, err := s.client.UpdateRepositorySettingsByName(
		ctx,
		connect_go.NewRequest(
			&v1alpha1.UpdateRepositorySettingsByNameRequest{
				OwnerName:      ownerName,
				RepositoryName: repositoryName,
				Visibility:     visibility,
			}),
	)
	if err != nil {
		return err
	}
	return nil
}
