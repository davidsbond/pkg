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

// Code generated by protoc-gen-go-apiclientgrpc. DO NOT EDIT.

package registryv1alpha1apiclientgrpc

import (
	context "context"
	v1alpha1 "github.com/bufbuild/buf/internal/gen/proto/go/buf/alpha/registry/v1alpha1"
	zap "go.uber.org/zap"
)

type authzService struct {
	logger          *zap.Logger
	client          v1alpha1.AuthzServiceClient
	contextModifier func(context.Context) context.Context
}

// UserCanAddUserOrganizationScopes returns whether the user is authorized
// to remove user scopes from an organization.
func (s *authzService) UserCanAddUserOrganizationScopes(ctx context.Context, organizationId string) (authorized bool, _ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	response, err := s.client.UserCanAddUserOrganizationScopes(
		ctx,
		&v1alpha1.UserCanAddUserOrganizationScopesRequest{
			OrganizationId: organizationId,
		},
	)
	if err != nil {
		return false, err
	}
	return response.Authorized, nil
}

// UserCanRemoveUserOrganizationScopes returns whether the user is authorized
// to remove user scopes from an organization.
func (s *authzService) UserCanRemoveUserOrganizationScopes(ctx context.Context, organizationId string) (authorized bool, _ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	response, err := s.client.UserCanRemoveUserOrganizationScopes(
		ctx,
		&v1alpha1.UserCanRemoveUserOrganizationScopesRequest{
			OrganizationId: organizationId,
		},
	)
	if err != nil {
		return false, err
	}
	return response.Authorized, nil
}

// UserCanCreateOrganizationRepository returns whether the user is authorized
// to create repositories in an organization.
func (s *authzService) UserCanCreateOrganizationRepository(ctx context.Context, organizationId string) (authorized bool, _ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	response, err := s.client.UserCanCreateOrganizationRepository(
		ctx,
		&v1alpha1.UserCanCreateOrganizationRepositoryRequest{
			OrganizationId: organizationId,
		},
	)
	if err != nil {
		return false, err
	}
	return response.Authorized, nil
}

// UserCanCreateOrganizationTeam returns whether the user is authorized
// to create teams in an organization.
func (s *authzService) UserCanCreateOrganizationTeam(ctx context.Context, organizationId string) (authorized bool, _ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	response, err := s.client.UserCanCreateOrganizationTeam(
		ctx,
		&v1alpha1.UserCanCreateOrganizationTeamRequest{
			OrganizationId: organizationId,
		},
	)
	if err != nil {
		return false, err
	}
	return response.Authorized, nil
}

// UserCanListOrganizationTeams returns whether the user is authorized
// to list teams in an organization.
func (s *authzService) UserCanListOrganizationTeams(ctx context.Context, organizationId string) (authorized bool, _ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	response, err := s.client.UserCanListOrganizationTeams(
		ctx,
		&v1alpha1.UserCanListOrganizationTeamsRequest{
			OrganizationId: organizationId,
		},
	)
	if err != nil {
		return false, err
	}
	return response.Authorized, nil
}

// UserCanSeeRepositorySettings returns whether the user is authorized
// to see repository settings.
func (s *authzService) UserCanSeeRepositorySettings(ctx context.Context, repositoryId string) (authorized bool, _ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	response, err := s.client.UserCanSeeRepositorySettings(
		ctx,
		&v1alpha1.UserCanSeeRepositorySettingsRequest{
			RepositoryId: repositoryId,
		},
	)
	if err != nil {
		return false, err
	}
	return response.Authorized, nil
}

// UserCanSeeOrganizationSettings returns whether the user is authorized
// to see organization settings.
func (s *authzService) UserCanSeeOrganizationSettings(ctx context.Context, organizationId string) (authorized bool, _ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	response, err := s.client.UserCanSeeOrganizationSettings(
		ctx,
		&v1alpha1.UserCanSeeOrganizationSettingsRequest{
			OrganizationId: organizationId,
		},
	)
	if err != nil {
		return false, err
	}
	return response.Authorized, nil
}

// UserCanReadPlugin returns whether the user has read access to the specified plugin.
func (s *authzService) UserCanReadPlugin(
	ctx context.Context,
	owner string,
	name string,
) (authorized bool, _ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	response, err := s.client.UserCanReadPlugin(
		ctx,
		&v1alpha1.UserCanReadPluginRequest{
			Owner: owner,
			Name:  name,
		},
	)
	if err != nil {
		return false, err
	}
	return response.Authorized, nil
}

// UserCanCreatePluginVersion returns whether the user is authorized
// to create a plugin version under the specified plugin.
func (s *authzService) UserCanCreatePluginVersion(
	ctx context.Context,
	owner string,
	name string,
) (authorized bool, _ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	response, err := s.client.UserCanCreatePluginVersion(
		ctx,
		&v1alpha1.UserCanCreatePluginVersionRequest{
			Owner: owner,
			Name:  name,
		},
	)
	if err != nil {
		return false, err
	}
	return response.Authorized, nil
}

// UserCanCreateTemplateVersion returns whether the user is authorized
// to create a template version under the specified template.
func (s *authzService) UserCanCreateTemplateVersion(
	ctx context.Context,
	owner string,
	name string,
) (authorized bool, _ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	response, err := s.client.UserCanCreateTemplateVersion(
		ctx,
		&v1alpha1.UserCanCreateTemplateVersionRequest{
			Owner: owner,
			Name:  name,
		},
	)
	if err != nil {
		return false, err
	}
	return response.Authorized, nil
}

// UserCanCreateOrganizationPlugin returns whether the user is authorized to create
// a plugin in an organization.
func (s *authzService) UserCanCreateOrganizationPlugin(ctx context.Context, organizationId string) (authorized bool, _ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	response, err := s.client.UserCanCreateOrganizationPlugin(
		ctx,
		&v1alpha1.UserCanCreateOrganizationPluginRequest{
			OrganizationId: organizationId,
		},
	)
	if err != nil {
		return false, err
	}
	return response.Authorized, nil
}

// UserCanCreateOrganizationPlugin returns whether the user is authorized to create
// a template in an organization.
func (s *authzService) UserCanCreateOrganizationTemplate(ctx context.Context, organizationId string) (authorized bool, _ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	response, err := s.client.UserCanCreateOrganizationTemplate(
		ctx,
		&v1alpha1.UserCanCreateOrganizationTemplateRequest{
			OrganizationId: organizationId,
		},
	)
	if err != nil {
		return false, err
	}
	return response.Authorized, nil
}

// UserCanSeePluginSettings returns whether the user is authorized
// to see plugin settings.
func (s *authzService) UserCanSeePluginSettings(
	ctx context.Context,
	owner string,
	name string,
) (authorized bool, _ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	response, err := s.client.UserCanSeePluginSettings(
		ctx,
		&v1alpha1.UserCanSeePluginSettingsRequest{
			Owner: owner,
			Name:  name,
		},
	)
	if err != nil {
		return false, err
	}
	return response.Authorized, nil
}

// UserCanSeeTemplateSettings returns whether the user is authorized
// to see template settings.
func (s *authzService) UserCanSeeTemplateSettings(
	ctx context.Context,
	owner string,
	name string,
) (authorized bool, _ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	response, err := s.client.UserCanSeeTemplateSettings(
		ctx,
		&v1alpha1.UserCanSeeTemplateSettingsRequest{
			Owner: owner,
			Name:  name,
		},
	)
	if err != nil {
		return false, err
	}
	return response.Authorized, nil
}
