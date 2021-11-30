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

// Code generated by protoc-gen-go-api. DO NOT EDIT.

package registryv1alpha1api

import (
	context "context"
)

// AuthzService supplies authorization helpers.
type AuthzService interface {
	// UserCanAddUserOrganizationScopes returns whether the user is authorized
	// to remove user scopes from an organization.
	UserCanAddUserOrganizationScopes(ctx context.Context, organizationId string) (authorized bool, err error)
	// UserCanRemoveUserOrganizationScopes returns whether the user is authorized
	// to remove user scopes from an organization.
	UserCanRemoveUserOrganizationScopes(ctx context.Context, organizationId string) (authorized bool, err error)
	// UserCanCreateOrganizationRepository returns whether the user is authorized
	// to create repositories in an organization.
	UserCanCreateOrganizationRepository(ctx context.Context, organizationId string) (authorized bool, err error)
	// UserCanCreateOrganizationTeam returns whether the user is authorized
	// to create teams in an organization.
	UserCanCreateOrganizationTeam(ctx context.Context, organizationId string) (authorized bool, err error)
	// UserCanListOrganizationTeams returns whether the user is authorized
	// to list teams in an organization.
	UserCanListOrganizationTeams(ctx context.Context, organizationId string) (authorized bool, err error)
	// UserCanSeeRepositorySettings returns whether the user is authorized
	// to see repository settings.
	UserCanSeeRepositorySettings(ctx context.Context, repositoryId string) (authorized bool, err error)
	// UserCanSeeOrganizationSettings returns whether the user is authorized
	// to see organization settings.
	UserCanSeeOrganizationSettings(ctx context.Context, organizationId string) (authorized bool, err error)
	// UserCanReadPlugin returns whether the user has read access to the specified plugin.
	UserCanReadPlugin(
		ctx context.Context,
		owner string,
		name string,
	) (authorized bool, err error)
	// UserCanCreatePluginVersion returns whether the user is authorized
	// to create a plugin version under the specified plugin.
	UserCanCreatePluginVersion(
		ctx context.Context,
		owner string,
		name string,
	) (authorized bool, err error)
	// UserCanCreateTemplateVersion returns whether the user is authorized
	// to create a template version under the specified template.
	UserCanCreateTemplateVersion(
		ctx context.Context,
		owner string,
		name string,
	) (authorized bool, err error)
	// UserCanCreateOrganizationPlugin returns whether the user is authorized to create
	// a plugin in an organization.
	UserCanCreateOrganizationPlugin(ctx context.Context, organizationId string) (authorized bool, err error)
	// UserCanCreateOrganizationPlugin returns whether the user is authorized to create
	// a template in an organization.
	UserCanCreateOrganizationTemplate(ctx context.Context, organizationId string) (authorized bool, err error)
	// UserCanSeePluginSettings returns whether the user is authorized
	// to see plugin settings.
	UserCanSeePluginSettings(
		ctx context.Context,
		owner string,
		name string,
	) (authorized bool, err error)
	// UserCanSeeTemplateSettings returns whether the user is authorized
	// to see template settings.
	UserCanSeeTemplateSettings(
		ctx context.Context,
		owner string,
		name string,
	) (authorized bool, err error)
}
