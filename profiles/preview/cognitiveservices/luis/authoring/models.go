// +build go1.9

// Copyright 2018 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package authoring

import original "github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v2.0/luis/authoring"

type AppsClient = original.AppsClient
type BaseClient = original.BaseClient
type ExamplesClient = original.ExamplesClient
type FeaturesClient = original.FeaturesClient
type ModelClient = original.ModelClient
type OperationStatusType = original.OperationStatusType

const (
	Failed  OperationStatusType = original.Failed
	FAILED  OperationStatusType = original.FAILED
	Success OperationStatusType = original.Success
)

type ReadableType = original.ReadableType

const (
	ReadableTypeClosedListEntityExtractor        ReadableType = original.ReadableTypeClosedListEntityExtractor
	ReadableTypeCompositeEntityExtractor         ReadableType = original.ReadableTypeCompositeEntityExtractor
	ReadableTypeEntityExtractor                  ReadableType = original.ReadableTypeEntityExtractor
	ReadableTypeHierarchicalChildEntityExtractor ReadableType = original.ReadableTypeHierarchicalChildEntityExtractor
	ReadableTypeHierarchicalEntityExtractor      ReadableType = original.ReadableTypeHierarchicalEntityExtractor
	ReadableTypeIntentClassifier                 ReadableType = original.ReadableTypeIntentClassifier
	ReadableTypePatternAnyEntityExtractor        ReadableType = original.ReadableTypePatternAnyEntityExtractor
	ReadableTypePrebuiltEntityExtractor          ReadableType = original.ReadableTypePrebuiltEntityExtractor
	ReadableTypeRegexEntityExtractor             ReadableType = original.ReadableTypeRegexEntityExtractor
)

type ReadableType1 = original.ReadableType1

const (
	ReadableType1ClosedListEntityExtractor        ReadableType1 = original.ReadableType1ClosedListEntityExtractor
	ReadableType1CompositeEntityExtractor         ReadableType1 = original.ReadableType1CompositeEntityExtractor
	ReadableType1EntityExtractor                  ReadableType1 = original.ReadableType1EntityExtractor
	ReadableType1HierarchicalChildEntityExtractor ReadableType1 = original.ReadableType1HierarchicalChildEntityExtractor
	ReadableType1HierarchicalEntityExtractor      ReadableType1 = original.ReadableType1HierarchicalEntityExtractor
	ReadableType1IntentClassifier                 ReadableType1 = original.ReadableType1IntentClassifier
	ReadableType1PatternAnyEntityExtractor        ReadableType1 = original.ReadableType1PatternAnyEntityExtractor
	ReadableType1PrebuiltEntityExtractor          ReadableType1 = original.ReadableType1PrebuiltEntityExtractor
	ReadableType1RegexEntityExtractor             ReadableType1 = original.ReadableType1RegexEntityExtractor
)

type ReadableType10 = original.ReadableType10

const (
	ReadableType10ClosedListEntityExtractor        ReadableType10 = original.ReadableType10ClosedListEntityExtractor
	ReadableType10CompositeEntityExtractor         ReadableType10 = original.ReadableType10CompositeEntityExtractor
	ReadableType10EntityExtractor                  ReadableType10 = original.ReadableType10EntityExtractor
	ReadableType10HierarchicalChildEntityExtractor ReadableType10 = original.ReadableType10HierarchicalChildEntityExtractor
	ReadableType10HierarchicalEntityExtractor      ReadableType10 = original.ReadableType10HierarchicalEntityExtractor
	ReadableType10IntentClassifier                 ReadableType10 = original.ReadableType10IntentClassifier
	ReadableType10PatternAnyEntityExtractor        ReadableType10 = original.ReadableType10PatternAnyEntityExtractor
	ReadableType10PrebuiltEntityExtractor          ReadableType10 = original.ReadableType10PrebuiltEntityExtractor
	ReadableType10RegexEntityExtractor             ReadableType10 = original.ReadableType10RegexEntityExtractor
)

type ReadableType2 = original.ReadableType2

const (
	ReadableType2ClosedListEntityExtractor        ReadableType2 = original.ReadableType2ClosedListEntityExtractor
	ReadableType2CompositeEntityExtractor         ReadableType2 = original.ReadableType2CompositeEntityExtractor
	ReadableType2EntityExtractor                  ReadableType2 = original.ReadableType2EntityExtractor
	ReadableType2HierarchicalChildEntityExtractor ReadableType2 = original.ReadableType2HierarchicalChildEntityExtractor
	ReadableType2HierarchicalEntityExtractor      ReadableType2 = original.ReadableType2HierarchicalEntityExtractor
	ReadableType2IntentClassifier                 ReadableType2 = original.ReadableType2IntentClassifier
	ReadableType2PatternAnyEntityExtractor        ReadableType2 = original.ReadableType2PatternAnyEntityExtractor
	ReadableType2PrebuiltEntityExtractor          ReadableType2 = original.ReadableType2PrebuiltEntityExtractor
	ReadableType2RegexEntityExtractor             ReadableType2 = original.ReadableType2RegexEntityExtractor
)

type ReadableType3 = original.ReadableType3

const (
	ReadableType3ClosedListEntityExtractor        ReadableType3 = original.ReadableType3ClosedListEntityExtractor
	ReadableType3CompositeEntityExtractor         ReadableType3 = original.ReadableType3CompositeEntityExtractor
	ReadableType3EntityExtractor                  ReadableType3 = original.ReadableType3EntityExtractor
	ReadableType3HierarchicalChildEntityExtractor ReadableType3 = original.ReadableType3HierarchicalChildEntityExtractor
	ReadableType3HierarchicalEntityExtractor      ReadableType3 = original.ReadableType3HierarchicalEntityExtractor
	ReadableType3IntentClassifier                 ReadableType3 = original.ReadableType3IntentClassifier
	ReadableType3PatternAnyEntityExtractor        ReadableType3 = original.ReadableType3PatternAnyEntityExtractor
	ReadableType3PrebuiltEntityExtractor          ReadableType3 = original.ReadableType3PrebuiltEntityExtractor
	ReadableType3RegexEntityExtractor             ReadableType3 = original.ReadableType3RegexEntityExtractor
)

type ReadableType4 = original.ReadableType4

const (
	ReadableType4ClosedListEntityExtractor        ReadableType4 = original.ReadableType4ClosedListEntityExtractor
	ReadableType4CompositeEntityExtractor         ReadableType4 = original.ReadableType4CompositeEntityExtractor
	ReadableType4EntityExtractor                  ReadableType4 = original.ReadableType4EntityExtractor
	ReadableType4HierarchicalChildEntityExtractor ReadableType4 = original.ReadableType4HierarchicalChildEntityExtractor
	ReadableType4HierarchicalEntityExtractor      ReadableType4 = original.ReadableType4HierarchicalEntityExtractor
	ReadableType4IntentClassifier                 ReadableType4 = original.ReadableType4IntentClassifier
	ReadableType4PatternAnyEntityExtractor        ReadableType4 = original.ReadableType4PatternAnyEntityExtractor
	ReadableType4PrebuiltEntityExtractor          ReadableType4 = original.ReadableType4PrebuiltEntityExtractor
	ReadableType4RegexEntityExtractor             ReadableType4 = original.ReadableType4RegexEntityExtractor
)

type ReadableType5 = original.ReadableType5

const (
	ReadableType5ClosedListEntityExtractor        ReadableType5 = original.ReadableType5ClosedListEntityExtractor
	ReadableType5CompositeEntityExtractor         ReadableType5 = original.ReadableType5CompositeEntityExtractor
	ReadableType5EntityExtractor                  ReadableType5 = original.ReadableType5EntityExtractor
	ReadableType5HierarchicalChildEntityExtractor ReadableType5 = original.ReadableType5HierarchicalChildEntityExtractor
	ReadableType5HierarchicalEntityExtractor      ReadableType5 = original.ReadableType5HierarchicalEntityExtractor
	ReadableType5IntentClassifier                 ReadableType5 = original.ReadableType5IntentClassifier
	ReadableType5PatternAnyEntityExtractor        ReadableType5 = original.ReadableType5PatternAnyEntityExtractor
	ReadableType5PrebuiltEntityExtractor          ReadableType5 = original.ReadableType5PrebuiltEntityExtractor
	ReadableType5RegexEntityExtractor             ReadableType5 = original.ReadableType5RegexEntityExtractor
)

type ReadableType6 = original.ReadableType6

const (
	ReadableType6ClosedListEntityExtractor        ReadableType6 = original.ReadableType6ClosedListEntityExtractor
	ReadableType6CompositeEntityExtractor         ReadableType6 = original.ReadableType6CompositeEntityExtractor
	ReadableType6EntityExtractor                  ReadableType6 = original.ReadableType6EntityExtractor
	ReadableType6HierarchicalChildEntityExtractor ReadableType6 = original.ReadableType6HierarchicalChildEntityExtractor
	ReadableType6HierarchicalEntityExtractor      ReadableType6 = original.ReadableType6HierarchicalEntityExtractor
	ReadableType6IntentClassifier                 ReadableType6 = original.ReadableType6IntentClassifier
	ReadableType6PatternAnyEntityExtractor        ReadableType6 = original.ReadableType6PatternAnyEntityExtractor
	ReadableType6PrebuiltEntityExtractor          ReadableType6 = original.ReadableType6PrebuiltEntityExtractor
	ReadableType6RegexEntityExtractor             ReadableType6 = original.ReadableType6RegexEntityExtractor
)

type ReadableType7 = original.ReadableType7

const (
	ReadableType7ClosedListEntityExtractor        ReadableType7 = original.ReadableType7ClosedListEntityExtractor
	ReadableType7CompositeEntityExtractor         ReadableType7 = original.ReadableType7CompositeEntityExtractor
	ReadableType7EntityExtractor                  ReadableType7 = original.ReadableType7EntityExtractor
	ReadableType7HierarchicalChildEntityExtractor ReadableType7 = original.ReadableType7HierarchicalChildEntityExtractor
	ReadableType7HierarchicalEntityExtractor      ReadableType7 = original.ReadableType7HierarchicalEntityExtractor
	ReadableType7IntentClassifier                 ReadableType7 = original.ReadableType7IntentClassifier
	ReadableType7PatternAnyEntityExtractor        ReadableType7 = original.ReadableType7PatternAnyEntityExtractor
	ReadableType7PrebuiltEntityExtractor          ReadableType7 = original.ReadableType7PrebuiltEntityExtractor
	ReadableType7RegexEntityExtractor             ReadableType7 = original.ReadableType7RegexEntityExtractor
)

type ReadableType8 = original.ReadableType8

const (
	ReadableType8ClosedListEntityExtractor        ReadableType8 = original.ReadableType8ClosedListEntityExtractor
	ReadableType8CompositeEntityExtractor         ReadableType8 = original.ReadableType8CompositeEntityExtractor
	ReadableType8EntityExtractor                  ReadableType8 = original.ReadableType8EntityExtractor
	ReadableType8HierarchicalChildEntityExtractor ReadableType8 = original.ReadableType8HierarchicalChildEntityExtractor
	ReadableType8HierarchicalEntityExtractor      ReadableType8 = original.ReadableType8HierarchicalEntityExtractor
	ReadableType8IntentClassifier                 ReadableType8 = original.ReadableType8IntentClassifier
	ReadableType8PatternAnyEntityExtractor        ReadableType8 = original.ReadableType8PatternAnyEntityExtractor
	ReadableType8PrebuiltEntityExtractor          ReadableType8 = original.ReadableType8PrebuiltEntityExtractor
	ReadableType8RegexEntityExtractor             ReadableType8 = original.ReadableType8RegexEntityExtractor
)

type ReadableType9 = original.ReadableType9

const (
	ReadableType9ClosedListEntityExtractor        ReadableType9 = original.ReadableType9ClosedListEntityExtractor
	ReadableType9CompositeEntityExtractor         ReadableType9 = original.ReadableType9CompositeEntityExtractor
	ReadableType9EntityExtractor                  ReadableType9 = original.ReadableType9EntityExtractor
	ReadableType9HierarchicalChildEntityExtractor ReadableType9 = original.ReadableType9HierarchicalChildEntityExtractor
	ReadableType9HierarchicalEntityExtractor      ReadableType9 = original.ReadableType9HierarchicalEntityExtractor
	ReadableType9IntentClassifier                 ReadableType9 = original.ReadableType9IntentClassifier
	ReadableType9PatternAnyEntityExtractor        ReadableType9 = original.ReadableType9PatternAnyEntityExtractor
	ReadableType9PrebuiltEntityExtractor          ReadableType9 = original.ReadableType9PrebuiltEntityExtractor
	ReadableType9RegexEntityExtractor             ReadableType9 = original.ReadableType9RegexEntityExtractor
)

type Status = original.Status

const (
	StatusFail       Status = original.StatusFail
	StatusInProgress Status = original.StatusInProgress
	StatusQueued     Status = original.StatusQueued
	StatusSuccess    Status = original.StatusSuccess
	StatusUpToDate   Status = original.StatusUpToDate
)

type Status1 = original.Status1

const (
	Status1Fail       Status1 = original.Status1Fail
	Status1InProgress Status1 = original.Status1InProgress
	Status1Queued     Status1 = original.Status1Queued
	Status1Success    Status1 = original.Status1Success
	Status1UpToDate   Status1 = original.Status1UpToDate
)

type TrainingStatus = original.TrainingStatus

const (
	InProgress    TrainingStatus = original.InProgress
	NeedsTraining TrainingStatus = original.NeedsTraining
	Trained       TrainingStatus = original.Trained
)

type ApplicationCreateObject = original.ApplicationCreateObject
type ApplicationInfoResponse = original.ApplicationInfoResponse
type ApplicationPublishObject = original.ApplicationPublishObject
type ApplicationSettings = original.ApplicationSettings
type ApplicationSettingUpdateObject = original.ApplicationSettingUpdateObject
type ApplicationUpdateObject = original.ApplicationUpdateObject
type AppVersionSettingObject = original.AppVersionSettingObject
type AvailableCulture = original.AvailableCulture
type AvailablePrebuiltEntityModel = original.AvailablePrebuiltEntityModel
type BatchLabelExample = original.BatchLabelExample
type ChildEntity = original.ChildEntity
type ClosedList = original.ClosedList
type ClosedListEntityExtractor = original.ClosedListEntityExtractor
type ClosedListModelCreateObject = original.ClosedListModelCreateObject
type ClosedListModelPatchObject = original.ClosedListModelPatchObject
type ClosedListModelUpdateObject = original.ClosedListModelUpdateObject
type CollaboratorsArray = original.CollaboratorsArray
type CompositeChildModelCreateObject = original.CompositeChildModelCreateObject
type CompositeEntityExtractor = original.CompositeEntityExtractor
type CompositeEntityModel = original.CompositeEntityModel
type CustomPrebuiltModel = original.CustomPrebuiltModel
type EndpointInfo = original.EndpointInfo
type EnqueueTrainingResponse = original.EnqueueTrainingResponse
type EntitiesSuggestionExample = original.EntitiesSuggestionExample
type EntityExtractor = original.EntityExtractor
type EntityLabel = original.EntityLabel
type EntityLabelObject = original.EntityLabelObject
type EntityModelInfo = original.EntityModelInfo
type EntityPrediction = original.EntityPrediction
type EntityRole = original.EntityRole
type EntityRoleCreateObject = original.EntityRoleCreateObject
type EntityRoleUpdateObject = original.EntityRoleUpdateObject
type ErrorResponse = original.ErrorResponse
type ExampleLabelObject = original.ExampleLabelObject
type ExplicitListItem = original.ExplicitListItem
type ExplicitListItemCreateObject = original.ExplicitListItemCreateObject
type ExplicitListItemUpdateObject = original.ExplicitListItemUpdateObject
type FeatureInfoObject = original.FeatureInfoObject
type FeaturesResponseObject = original.FeaturesResponseObject
type HierarchicalChildEntity = original.HierarchicalChildEntity
type HierarchicalChildModelCreateObject = original.HierarchicalChildModelCreateObject
type HierarchicalChildModelUpdateObject = original.HierarchicalChildModelUpdateObject
type HierarchicalEntityExtractor = original.HierarchicalEntityExtractor
type HierarchicalEntityModel = original.HierarchicalEntityModel
type HierarchicalModel = original.HierarchicalModel
type IntentClassifier = original.IntentClassifier
type IntentPrediction = original.IntentPrediction
type IntentsSuggestionExample = original.IntentsSuggestionExample
type JSONEntity = original.JSONEntity
type JSONModelFeature = original.JSONModelFeature
type JSONRegexFeature = original.JSONRegexFeature
type JSONUtterance = original.JSONUtterance
type LabeledUtterance = original.LabeledUtterance
type LabelExampleResponse = original.LabelExampleResponse
type LabelTextObject = original.LabelTextObject
type LuisApp = original.LuisApp
type ModelCreateObject = original.ModelCreateObject
type ModelInfo = original.ModelInfo
type ModelInfoResponse = original.ModelInfoResponse
type ModelTrainingDetails = original.ModelTrainingDetails
type ModelTrainingInfo = original.ModelTrainingInfo
type ModelUpdateObject = original.ModelUpdateObject
type OperationError = original.OperationError
type OperationStatus = original.OperationStatus
type PatternAny = original.PatternAny
type PatternAnyEntityExtractor = original.PatternAnyEntityExtractor
type PatternAnyModelCreateObject = original.PatternAnyModelCreateObject
type PatternAnyModelUpdateObject = original.PatternAnyModelUpdateObject
type PatternCreateObject = original.PatternCreateObject
type PatternFeatureInfo = original.PatternFeatureInfo
type PatternRule = original.PatternRule
type PatternRuleCreateObject = original.PatternRuleCreateObject
type PatternRuleInfo = original.PatternRuleInfo
type PatternRuleUpdateObject = original.PatternRuleUpdateObject
type PatternUpdateObject = original.PatternUpdateObject
type PersonalAssistantsResponse = original.PersonalAssistantsResponse
type PhraselistCreateObject = original.PhraselistCreateObject
type PhraseListFeatureInfo = original.PhraseListFeatureInfo
type PhraselistUpdateObject = original.PhraselistUpdateObject
type PrebuiltDomain = original.PrebuiltDomain
type PrebuiltDomainCreateBaseObject = original.PrebuiltDomainCreateBaseObject
type PrebuiltDomainCreateObject = original.PrebuiltDomainCreateObject
type PrebuiltDomainItem = original.PrebuiltDomainItem
type PrebuiltDomainModelCreateObject = original.PrebuiltDomainModelCreateObject
type PrebuiltDomainObject = original.PrebuiltDomainObject
type PrebuiltEntity = original.PrebuiltEntity
type PrebuiltEntityExtractor = original.PrebuiltEntityExtractor
type ProductionOrStagingEndpointInfo = original.ProductionOrStagingEndpointInfo
type PublishSettings = original.PublishSettings
type PublishSettingUpdateObject = original.PublishSettingUpdateObject
type RegexEntity = original.RegexEntity
type RegexEntityExtractor = original.RegexEntityExtractor
type RegexModelCreateObject = original.RegexModelCreateObject
type RegexModelUpdateObject = original.RegexModelUpdateObject
type SetObject = original.SetObject
type SubClosedList = original.SubClosedList
type SubClosedListResponse = original.SubClosedListResponse
type TaskUpdateObject = original.TaskUpdateObject
type UserAccessList = original.UserAccessList
type UserCollaborator = original.UserCollaborator
type VersionInfo = original.VersionInfo
type WordListBaseUpdateObject = original.WordListBaseUpdateObject
type WordListObject = original.WordListObject
type PatternClient = original.PatternClient
type PermissionsClient = original.PermissionsClient
type SettingsClient = original.SettingsClient
type TrainClient = original.TrainClient
type VersionsClient = original.VersionsClient

func NewAppsClient(endpoint string, ocpApimSubscriptionKey string) AppsClient {
	return original.NewAppsClient(endpoint, ocpApimSubscriptionKey)
}
func New(endpoint string, ocpApimSubscriptionKey string) BaseClient {
	return original.New(endpoint, ocpApimSubscriptionKey)
}
func NewWithoutDefaults(endpoint string, ocpApimSubscriptionKey string) BaseClient {
	return original.NewWithoutDefaults(endpoint, ocpApimSubscriptionKey)
}
func NewExamplesClient(endpoint string, ocpApimSubscriptionKey string) ExamplesClient {
	return original.NewExamplesClient(endpoint, ocpApimSubscriptionKey)
}
func NewFeaturesClient(endpoint string, ocpApimSubscriptionKey string) FeaturesClient {
	return original.NewFeaturesClient(endpoint, ocpApimSubscriptionKey)
}
func NewModelClient(endpoint string, ocpApimSubscriptionKey string) ModelClient {
	return original.NewModelClient(endpoint, ocpApimSubscriptionKey)
}
func PossibleOperationStatusTypeValues() []OperationStatusType {
	return original.PossibleOperationStatusTypeValues()
}
func PossibleReadableTypeValues() []ReadableType {
	return original.PossibleReadableTypeValues()
}
func PossibleReadableType1Values() []ReadableType1 {
	return original.PossibleReadableType1Values()
}
func PossibleReadableType10Values() []ReadableType10 {
	return original.PossibleReadableType10Values()
}
func PossibleReadableType2Values() []ReadableType2 {
	return original.PossibleReadableType2Values()
}
func PossibleReadableType3Values() []ReadableType3 {
	return original.PossibleReadableType3Values()
}
func PossibleReadableType4Values() []ReadableType4 {
	return original.PossibleReadableType4Values()
}
func PossibleReadableType5Values() []ReadableType5 {
	return original.PossibleReadableType5Values()
}
func PossibleReadableType6Values() []ReadableType6 {
	return original.PossibleReadableType6Values()
}
func PossibleReadableType7Values() []ReadableType7 {
	return original.PossibleReadableType7Values()
}
func PossibleReadableType8Values() []ReadableType8 {
	return original.PossibleReadableType8Values()
}
func PossibleReadableType9Values() []ReadableType9 {
	return original.PossibleReadableType9Values()
}
func PossibleStatusValues() []Status {
	return original.PossibleStatusValues()
}
func PossibleStatus1Values() []Status1 {
	return original.PossibleStatus1Values()
}
func PossibleTrainingStatusValues() []TrainingStatus {
	return original.PossibleTrainingStatusValues()
}
func NewPatternClient(endpoint string, ocpApimSubscriptionKey string) PatternClient {
	return original.NewPatternClient(endpoint, ocpApimSubscriptionKey)
}
func NewPermissionsClient(endpoint string, ocpApimSubscriptionKey string) PermissionsClient {
	return original.NewPermissionsClient(endpoint, ocpApimSubscriptionKey)
}
func NewSettingsClient(endpoint string, ocpApimSubscriptionKey string) SettingsClient {
	return original.NewSettingsClient(endpoint, ocpApimSubscriptionKey)
}
func NewTrainClient(endpoint string, ocpApimSubscriptionKey string) TrainClient {
	return original.NewTrainClient(endpoint, ocpApimSubscriptionKey)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
func NewVersionsClient(endpoint string, ocpApimSubscriptionKey string) VersionsClient {
	return original.NewVersionsClient(endpoint, ocpApimSubscriptionKey)
}
