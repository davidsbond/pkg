// Copyright 2020-2021 Buf Technologies, Inc.
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

// Code generated by protoc-gen-go-apiclienttwirp. DO NOT EDIT.

package registryv1alpha1apiclienttwirp

import (
	context "context"
	v1alpha1 "github.com/bufbuild/buf/internal/gen/proto/go/buf/alpha/registry/v1alpha1"
	zap "go.uber.org/zap"
)

type repositoryService struct {
	logger          *zap.Logger
	client          v1alpha1.RepositoryService
	contextModifier func(context.Context) context.Context
}

// GetRepository gets a repository by ID.
func (s *repositoryService) GetRepository(ctx context.Context, id string) (repository *v1alpha1.Repository, _ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	response, err := s.client.GetRepository(
		ctx,
		&v1alpha1.GetRepositoryRequest{
			Id: id,
		},
	)
	if err != nil {
		return nil, err
	}
	return response.Repository, nil
}

// GetRepositoryByFullName gets a repository by full name.
func (s *repositoryService) GetRepositoryByFullName(ctx context.Context, fullName string) (repository *v1alpha1.Repository, _ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	response, err := s.client.GetRepositoryByFullName(
		ctx,
		&v1alpha1.GetRepositoryByFullNameRequest{
			FullName: fullName,
		},
	)
	if err != nil {
		return nil, err
	}
	return response.Repository, nil
}

// ListRepositories lists all repositories.
func (s *repositoryService) ListRepositories(
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
		&v1alpha1.ListRepositoriesRequest{
			PageSize:  pageSize,
			PageToken: pageToken,
			Reverse:   reverse,
		},
	)
	if err != nil {
		return nil, "", err
	}
	return response.Repositories, response.NextPageToken, nil
}

// ListUserRepositories lists all repositories belonging to a user.
func (s *repositoryService) ListUserRepositories(
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
		&v1alpha1.ListUserRepositoriesRequest{
			UserId:    userId,
			PageSize:  pageSize,
			PageToken: pageToken,
			Reverse:   reverse,
		},
	)
	if err != nil {
		return nil, "", err
	}
	return response.Repositories, response.NextPageToken, nil
}

// ListUserRepositories lists all repositories a user can access.
func (s *repositoryService) ListRepositoriesUserCanAccess(
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
		&v1alpha1.ListRepositoriesUserCanAccessRequest{
			PageSize:  pageSize,
			PageToken: pageToken,
			Reverse:   reverse,
		},
	)
	if err != nil {
		return nil, "", err
	}
	return response.Repositories, response.NextPageToken, nil
}

// ListOrganizationRepositories lists all repositories for an organization.
func (s *repositoryService) ListOrganizationRepositories(
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
		&v1alpha1.ListOrganizationRepositoriesRequest{
			OrganizationId: organizationId,
			PageSize:       pageSize,
			PageToken:      pageToken,
			Reverse:        reverse,
		},
	)
	if err != nil {
		return nil, "", err
	}
	return response.Repositories, response.NextPageToken, nil
}

// CreateRepositoryByFullName creates a new repository by full name.
func (s *repositoryService) CreateRepositoryByFullName(
	ctx context.Context,
	fullName string,
	visibility v1alpha1.Visibility,
) (repository *v1alpha1.Repository, _ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	response, err := s.client.CreateRepositoryByFullName(
		ctx,
		&v1alpha1.CreateRepositoryByFullNameRequest{
			FullName:   fullName,
			Visibility: visibility,
		},
	)
	if err != nil {
		return nil, err
	}
	return response.Repository, nil
}

// UpdateRepositoryName updates a repository's name.
func (s *repositoryService) UpdateRepositoryName(
	ctx context.Context,
	id string,
	newName string,
) (repository *v1alpha1.Repository, _ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	response, err := s.client.UpdateRepositoryName(
		ctx,
		&v1alpha1.UpdateRepositoryNameRequest{
			Id:      id,
			NewName: newName,
		},
	)
	if err != nil {
		return nil, err
	}
	return response.Repository, nil
}

// UpdateRepositoryNameByFullName updates a repository's name by full name.
func (s *repositoryService) UpdateRepositoryNameByFullName(
	ctx context.Context,
	fullName string,
	newName string,
) (repository *v1alpha1.Repository, _ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	response, err := s.client.UpdateRepositoryNameByFullName(
		ctx,
		&v1alpha1.UpdateRepositoryNameByFullNameRequest{
			FullName: fullName,
			NewName:  newName,
		},
	)
	if err != nil {
		return nil, err
	}
	return response.Repository, nil
}

// UpdateRepositoryVisibility updates a repository's visibility.
func (s *repositoryService) UpdateRepositoryVisibility(
	ctx context.Context,
	id string,
	newVisibility v1alpha1.Visibility,
) (repository *v1alpha1.Repository, _ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	response, err := s.client.UpdateRepositoryVisibility(
		ctx,
		&v1alpha1.UpdateRepositoryVisibilityRequest{
			Id:            id,
			NewVisibility: newVisibility,
		},
	)
	if err != nil {
		return nil, err
	}
	return response.Repository, nil
}

// UpdateRepositoryVisibilityByName updates a repository's visibility by name.
func (s *repositoryService) UpdateRepositoryVisibilityByName(
	ctx context.Context,
	ownerName string,
	repositoryName string,
	newVisibility v1alpha1.Visibility,
) (repository *v1alpha1.Repository, _ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	response, err := s.client.UpdateRepositoryVisibilityByName(
		ctx,
		&v1alpha1.UpdateRepositoryVisibilityByNameRequest{
			OwnerName:      ownerName,
			RepositoryName: repositoryName,
			NewVisibility:  newVisibility,
		},
	)
	if err != nil {
		return nil, err
	}
	return response.Repository, nil
}

// DeleteRepository deletes a repository.
func (s *repositoryService) DeleteRepository(ctx context.Context, id string) (_ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	_, err := s.client.DeleteRepository(
		ctx,
		&v1alpha1.DeleteRepositoryRequest{
			Id: id,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

// DeleteRepositoryByFullName deletes a repository by full name.
func (s *repositoryService) DeleteRepositoryByFullName(ctx context.Context, fullName string) (_ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	_, err := s.client.DeleteRepositoryByFullName(
		ctx,
		&v1alpha1.DeleteRepositoryByFullNameRequest{
			FullName: fullName,
		},
	)
	if err != nil {
		return err
	}
	return nil
}
