/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package loader

import (
	"context"

	applicationapiv1alpha1 "github.com/redhat-appstudio/application-api/api/v1alpha1"
	"github.com/redhat-appstudio/integration-service/api/v1beta1"
	releasev1alpha1 "github.com/redhat-appstudio/release-service/api/v1alpha1"
	tektonv1beta1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type (
	contextKey int
	MockData   struct {
		ContextKey contextKey
		Err        error
		Resource   any
	}
	mockLoader struct {
		loader ObjectLoader
	}
)

const (
	ApplicationContextKey                      contextKey = iota
	ComponentContextKey                        contextKey = iota
	SnapshotContextKey                         contextKey = iota
	SnapshotEnvironmentBindingContextKey       contextKey = iota
	DeploymentTargetContextKey                 contextKey = iota
	DeploymentTargetClaimContextKey            contextKey = iota
	IntegrationTestScenarioContextKey          contextKey = iota
	TaskRunContextKey                          contextKey = iota
	ApplicationComponentsContextKey            contextKey = iota
	SnapshotComponentsContextKey               contextKey = iota
	EnvironmentContextKey                      contextKey = iota
	ReleaseContextKey                          contextKey = iota
	PipelineRunsContextKey                     contextKey = iota
	DeploymentTargetClassContextKey            contextKey = iota
	AllIntegrationTestScenariosContextKey      contextKey = iota
	RequiredIntegrationTestScenariosContextKey contextKey = iota
	AllSnapshotsContextKey                     contextKey = iota
	AutoReleasePlansContextKey                 contextKey = iota
)

func GetMockedContext(ctx context.Context, data []MockData) context.Context {
	for _, mockData := range data {
		ctx = context.WithValue(ctx, mockData.ContextKey, mockData)
	}

	return ctx
}

func NewMockLoader() ObjectLoader {
	return &mockLoader{
		loader: NewLoader(),
	}
}

// getMockedResourceAndErrorFromContext returns the mocked data found in the context passed as an argument. The data is
// to be found in the contextDataKey key. If not there, a panic will be raised.
func getMockedResourceAndErrorFromContext[T any](ctx context.Context, contextKey contextKey, _ T) (T, error) {
	var resource T
	var err error

	value := ctx.Value(contextKey)
	if value == nil {
		panic("Mocked data not found in the context")
	}

	data, _ := value.(MockData)

	if data.Resource != nil {
		resource = data.Resource.(T)
	}

	if data.Err != nil {
		err = data.Err
	}

	return resource, err
}

// GetAllEnvironments returns the resource and error passed as values of the context.
func (l *mockLoader) GetAllEnvironments(c client.Client, ctx context.Context, application *applicationapiv1alpha1.Application) (*[]applicationapiv1alpha1.Environment, error) {
	if ctx.Value(EnvironmentContextKey) == nil {
		return l.loader.GetAllEnvironments(c, ctx, application)
	}
	environment, err := getMockedResourceAndErrorFromContext(ctx, EnvironmentContextKey, &applicationapiv1alpha1.Environment{})
	return &[]applicationapiv1alpha1.Environment{*environment}, err
}

// GetReleasesWithSnapshot returns the resource and error passed as values of the context.
func (l *mockLoader) GetReleasesWithSnapshot(c client.Client, ctx context.Context, snapshot *applicationapiv1alpha1.Snapshot) (*[]releasev1alpha1.Release, error) {
	if ctx.Value(ReleaseContextKey) == nil {
		return l.loader.GetReleasesWithSnapshot(c, ctx, snapshot)
	}
	release, err := getMockedResourceAndErrorFromContext(ctx, ReleaseContextKey, &releasev1alpha1.Release{})
	return &[]releasev1alpha1.Release{*release}, err
}

// GetAllApplicationComponents returns the resource and error passed as values of the context.
func (l *mockLoader) GetAllApplicationComponents(c client.Client, ctx context.Context, application *applicationapiv1alpha1.Application) (*[]applicationapiv1alpha1.Component, error) {
	if ctx.Value(ApplicationComponentsContextKey) == nil {
		return l.loader.GetAllApplicationComponents(c, ctx, application)
	}
	components, err := getMockedResourceAndErrorFromContext(ctx, ApplicationComponentsContextKey, []applicationapiv1alpha1.Component{})
	return &components, err
}

// GetAllSnapshotComponents returns the resource and error passed as values of the context.
func (l *mockLoader) GetAllSnapshotComponents(c client.Client, ctx context.Context, snapshot *applicationapiv1alpha1.Snapshot) (*[]applicationapiv1alpha1.Component, error) {
	if ctx.Value(SnapshotComponentsContextKey) == nil {
		return l.loader.GetAllSnapshotComponents(c, ctx, snapshot)
	}
	components, err := getMockedResourceAndErrorFromContext(ctx, SnapshotComponentsContextKey, []applicationapiv1alpha1.Component{})
	return &components, err
}

// GetApplicationFromSnapshot returns the resource and error passed as values of the context.
func (l *mockLoader) GetApplicationFromSnapshot(c client.Client, ctx context.Context, snapshot *applicationapiv1alpha1.Snapshot) (*applicationapiv1alpha1.Application, error) {
	if ctx.Value(ApplicationContextKey) == nil {
		return l.loader.GetApplicationFromSnapshot(c, ctx, snapshot)
	}
	return getMockedResourceAndErrorFromContext(ctx, ApplicationContextKey, &applicationapiv1alpha1.Application{})
}

// GetComponentFromSnapshot returns the resource and error passed as values of the context.
func (l *mockLoader) GetComponentFromSnapshot(c client.Client, ctx context.Context, snapshot *applicationapiv1alpha1.Snapshot) (*applicationapiv1alpha1.Component, error) {
	if ctx.Value(ComponentContextKey) == nil {
		return l.loader.GetComponentFromSnapshot(c, ctx, snapshot)
	}
	return getMockedResourceAndErrorFromContext(ctx, ComponentContextKey, &applicationapiv1alpha1.Component{})
}

// GetComponentFromPipelineRun returns the resource and error passed as values of the context.
func (l *mockLoader) GetComponentFromPipelineRun(c client.Client, ctx context.Context, pipelineRun *tektonv1beta1.PipelineRun) (*applicationapiv1alpha1.Component, error) {
	if ctx.Value(ComponentContextKey) == nil {
		return l.loader.GetComponentFromPipelineRun(c, ctx, pipelineRun)
	}
	return getMockedResourceAndErrorFromContext(ctx, ComponentContextKey, &applicationapiv1alpha1.Component{})
}

// GetApplicationFromPipelineRun returns the resource and error passed as values of the context.
func (l *mockLoader) GetApplicationFromPipelineRun(c client.Client, ctx context.Context, pipelineRun *tektonv1beta1.PipelineRun) (*applicationapiv1alpha1.Application, error) {
	if ctx.Value(ApplicationContextKey) == nil {
		return l.loader.GetApplicationFromPipelineRun(c, ctx, pipelineRun)
	}
	return getMockedResourceAndErrorFromContext(ctx, ApplicationContextKey, &applicationapiv1alpha1.Application{})
}

// GetApplicationFromComponent returns the resource and error passed as values of the context.
func (l *mockLoader) GetApplicationFromComponent(c client.Client, ctx context.Context, component *applicationapiv1alpha1.Component) (*applicationapiv1alpha1.Application, error) {
	if ctx.Value(ApplicationContextKey) == nil {
		return l.loader.GetApplicationFromComponent(c, ctx, component)
	}
	return getMockedResourceAndErrorFromContext(ctx, ApplicationContextKey, &applicationapiv1alpha1.Application{})
}

// GetEnvironmentFromIntegrationPipelineRun returns the resource and error passed as values of the context.
func (l *mockLoader) GetEnvironmentFromIntegrationPipelineRun(c client.Client, ctx context.Context, pipelineRun *tektonv1beta1.PipelineRun) (*applicationapiv1alpha1.Environment, error) {
	if ctx.Value(EnvironmentContextKey) == nil {
		return l.loader.GetEnvironmentFromIntegrationPipelineRun(c, ctx, pipelineRun)
	}
	return getMockedResourceAndErrorFromContext(ctx, EnvironmentContextKey, &applicationapiv1alpha1.Environment{})
}

// GetSnapshotFromPipelineRun returns the resource and error passed as values of the context.
func (l *mockLoader) GetSnapshotFromPipelineRun(c client.Client, ctx context.Context, pipelineRun *tektonv1beta1.PipelineRun) (*applicationapiv1alpha1.Snapshot, error) {
	if ctx.Value(SnapshotContextKey) == nil {
		return l.loader.GetSnapshotFromPipelineRun(c, ctx, pipelineRun)
	}
	return getMockedResourceAndErrorFromContext(ctx, SnapshotContextKey, &applicationapiv1alpha1.Snapshot{})
}

// FindAvailableDeploymentTargetClass returns the resource and error passed as values of the context.
func (l *mockLoader) FindAvailableDeploymentTargetClass(c client.Client, ctx context.Context) (*applicationapiv1alpha1.DeploymentTargetClass, error) {
	if ctx.Value(DeploymentTargetClassContextKey) == nil {
		return l.loader.FindAvailableDeploymentTargetClass(c, ctx)
	}
	return getMockedResourceAndErrorFromContext(ctx, DeploymentTargetClassContextKey, &applicationapiv1alpha1.DeploymentTargetClass{})
}

// GetAllIntegrationTestScenariosForApplication returns the resource and error passed as values of the context.
func (l *mockLoader) GetAllIntegrationTestScenariosForApplication(c client.Client, ctx context.Context, application *applicationapiv1alpha1.Application) (*[]v1beta1.IntegrationTestScenario, error) {
	if ctx.Value(AllIntegrationTestScenariosContextKey) == nil {
		return l.loader.GetAllIntegrationTestScenariosForApplication(c, ctx, application)
	}
	integrationTestScenarios, err := getMockedResourceAndErrorFromContext(ctx, AllIntegrationTestScenariosContextKey, []v1beta1.IntegrationTestScenario{})
	return &integrationTestScenarios, err
}

// GetRequiredIntegrationTestScenariosForApplication returns the resource and error passed as values of the context.
func (l *mockLoader) GetRequiredIntegrationTestScenariosForApplication(c client.Client, ctx context.Context, application *applicationapiv1alpha1.Application) (*[]v1beta1.IntegrationTestScenario, error) {
	if ctx.Value(RequiredIntegrationTestScenariosContextKey) == nil {
		return l.loader.GetRequiredIntegrationTestScenariosForApplication(c, ctx, application)
	}
	integrationTestScenarios, err := getMockedResourceAndErrorFromContext(ctx, RequiredIntegrationTestScenariosContextKey, []v1beta1.IntegrationTestScenario{})
	return &integrationTestScenarios, err
}

// GetDeploymentTargetClaimForEnvironment returns the resource and error passed as values of the context.
func (l *mockLoader) GetDeploymentTargetClaimForEnvironment(c client.Client, ctx context.Context, environment *applicationapiv1alpha1.Environment) (*applicationapiv1alpha1.DeploymentTargetClaim, error) {
	if ctx.Value(DeploymentTargetClaimContextKey) == nil {
		return l.loader.GetDeploymentTargetClaimForEnvironment(c, ctx, environment)
	}
	return getMockedResourceAndErrorFromContext(ctx, DeploymentTargetClaimContextKey, &applicationapiv1alpha1.DeploymentTargetClaim{})
}

// GetDeploymentTargetForDeploymentTargetClaim returns the resource and error passed as values of the context.
func (l *mockLoader) GetDeploymentTargetForDeploymentTargetClaim(c client.Client, ctx context.Context, dtc *applicationapiv1alpha1.DeploymentTargetClaim) (*applicationapiv1alpha1.DeploymentTarget, error) {
	if ctx.Value(DeploymentTargetContextKey) == nil {
		return l.loader.GetDeploymentTargetForDeploymentTargetClaim(c, ctx, dtc)
	}
	return getMockedResourceAndErrorFromContext(ctx, DeploymentTargetContextKey, &applicationapiv1alpha1.DeploymentTarget{})
}

// FindExistingSnapshotEnvironmentBinding returns the resource and error passed as values of the context.
func (l *mockLoader) FindExistingSnapshotEnvironmentBinding(c client.Client, ctx context.Context, application *applicationapiv1alpha1.Application, environment *applicationapiv1alpha1.Environment) (*applicationapiv1alpha1.SnapshotEnvironmentBinding, error) {
	if ctx.Value(SnapshotEnvironmentBindingContextKey) == nil {
		return l.loader.FindExistingSnapshotEnvironmentBinding(c, ctx, application, environment)
	}
	return getMockedResourceAndErrorFromContext(ctx, SnapshotEnvironmentBindingContextKey, &applicationapiv1alpha1.SnapshotEnvironmentBinding{})
}

// GetAllPipelineRunsForSnapshotAndScenario returns the resource and error passed as values of the context.
func (l *mockLoader) GetAllPipelineRunsForSnapshotAndScenario(c client.Client, ctx context.Context, snapshot *applicationapiv1alpha1.Snapshot, integrationTestScenario *v1beta1.IntegrationTestScenario) (*[]tektonv1beta1.PipelineRun, error) {
	if ctx.Value(PipelineRunsContextKey) == nil {
		return l.loader.GetAllPipelineRunsForSnapshotAndScenario(c, ctx, snapshot, integrationTestScenario)
	}
	pipelineRuns, err := getMockedResourceAndErrorFromContext(ctx, PipelineRunsContextKey, []tektonv1beta1.PipelineRun{})
	return &pipelineRuns, err
}

// GetAllBuildPipelineRunsForComponent returns the resource and error passed as values of the context.
func (l *mockLoader) GetAllBuildPipelineRunsForComponent(c client.Client, ctx context.Context, component *applicationapiv1alpha1.Component) (*[]tektonv1beta1.PipelineRun, error) {
	if ctx.Value(PipelineRunsContextKey) == nil {
		return l.loader.GetAllBuildPipelineRunsForComponent(c, ctx, component)
	}
	pipelineRuns, err := getMockedResourceAndErrorFromContext(ctx, PipelineRunsContextKey, []tektonv1beta1.PipelineRun{})
	return &pipelineRuns, err
}

// GetAllSnapshots returns the resource and error passed as values of the context.
func (l *mockLoader) GetAllSnapshots(c client.Client, ctx context.Context, application *applicationapiv1alpha1.Application) (*[]applicationapiv1alpha1.Snapshot, error) {
	if ctx.Value(AllSnapshotsContextKey) == nil {
		return l.loader.GetAllSnapshots(c, ctx, application)
	}
	snapshots, err := getMockedResourceAndErrorFromContext(ctx, AllSnapshotsContextKey, []applicationapiv1alpha1.Snapshot{})
	return &snapshots, err
}

// GetAutoReleasePlansForApplication returns the resource and error passed as values of the context.
func (l *mockLoader) GetAutoReleasePlansForApplication(c client.Client, ctx context.Context, application *applicationapiv1alpha1.Application) (*[]releasev1alpha1.ReleasePlan, error) {
	if ctx.Value(AutoReleasePlansContextKey) == nil {
		return l.loader.GetAutoReleasePlansForApplication(c, ctx, application)
	}
	autoReleasePlans, err := getMockedResourceAndErrorFromContext(ctx, AutoReleasePlansContextKey, []releasev1alpha1.ReleasePlan{})
	return &autoReleasePlans, err
}
