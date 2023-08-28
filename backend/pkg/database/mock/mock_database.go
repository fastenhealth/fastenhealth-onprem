// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package mock_database is a generated GoMock package.
package mock_database

import (
	context "context"
	reflect "reflect"

	pkg "github.com/fastenhealth/fasten-onprem/backend/pkg"
	models "github.com/fastenhealth/fasten-onprem/backend/pkg/models"
	models0 "github.com/fastenhealth/fasten-sources/clients/models"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockDatabaseRepository is a mock of DatabaseRepository interface.
type MockDatabaseRepository struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseRepositoryMockRecorder
}

// MockDatabaseRepositoryMockRecorder is the mock recorder for MockDatabaseRepository.
type MockDatabaseRepositoryMockRecorder struct {
	mock *MockDatabaseRepository
}

// NewMockDatabaseRepository creates a new mock instance.
func NewMockDatabaseRepository(ctrl *gomock.Controller) *MockDatabaseRepository {
	mock := &MockDatabaseRepository{ctrl: ctrl}
	mock.recorder = &MockDatabaseRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDatabaseRepository) EXPECT() *MockDatabaseRepositoryMockRecorder {
	return m.recorder
}

// AddResourceAssociation mocks base method.
func (m *MockDatabaseRepository) AddResourceAssociation(ctx context.Context, source *models.SourceCredential, resourceType, resourceId string, relatedSource *models.SourceCredential, relatedResourceType, relatedResourceId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddResourceAssociation", ctx, source, resourceType, resourceId, relatedSource, relatedResourceType, relatedResourceId)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddResourceAssociation indicates an expected call of AddResourceAssociation.
func (mr *MockDatabaseRepositoryMockRecorder) AddResourceAssociation(ctx, source, resourceType, resourceId, relatedSource, relatedResourceType, relatedResourceId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddResourceAssociation", reflect.TypeOf((*MockDatabaseRepository)(nil).AddResourceAssociation), ctx, source, resourceType, resourceId, relatedSource, relatedResourceType, relatedResourceId)
}

// AddResourceComposition mocks base method.
func (m *MockDatabaseRepository) AddResourceComposition(ctx context.Context, compositionTitle string, resources []*models.ResourceBase) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddResourceComposition", ctx, compositionTitle, resources)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddResourceComposition indicates an expected call of AddResourceComposition.
func (mr *MockDatabaseRepositoryMockRecorder) AddResourceComposition(ctx, compositionTitle, resources interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddResourceComposition", reflect.TypeOf((*MockDatabaseRepository)(nil).AddResourceComposition), ctx, compositionTitle, resources)
}

// Close mocks base method.
func (m *MockDatabaseRepository) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockDatabaseRepositoryMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockDatabaseRepository)(nil).Close))
}

// CreateGlossaryEntry mocks base method.
func (m *MockDatabaseRepository) CreateGlossaryEntry(ctx context.Context, glossaryEntry *models.Glossary) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateGlossaryEntry", ctx, glossaryEntry)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateGlossaryEntry indicates an expected call of CreateGlossaryEntry.
func (mr *MockDatabaseRepositoryMockRecorder) CreateGlossaryEntry(ctx, glossaryEntry interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGlossaryEntry", reflect.TypeOf((*MockDatabaseRepository)(nil).CreateGlossaryEntry), ctx, glossaryEntry)
}

// CreateSource mocks base method.
func (m *MockDatabaseRepository) CreateSource(arg0 context.Context, arg1 *models.SourceCredential) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSource", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateSource indicates an expected call of CreateSource.
func (mr *MockDatabaseRepositoryMockRecorder) CreateSource(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSource", reflect.TypeOf((*MockDatabaseRepository)(nil).CreateSource), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockDatabaseRepository) CreateUser(arg0 context.Context, arg1 *models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockDatabaseRepositoryMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockDatabaseRepository)(nil).CreateUser), arg0, arg1)
}

// FindResourceAssociationsByTypeAndId mocks base method.
func (m *MockDatabaseRepository) FindResourceAssociationsByTypeAndId(ctx context.Context, source *models.SourceCredential, resourceType, resourceId string) ([]models.RelatedResource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindResourceAssociationsByTypeAndId", ctx, source, resourceType, resourceId)
	ret0, _ := ret[0].([]models.RelatedResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindResourceAssociationsByTypeAndId indicates an expected call of FindResourceAssociationsByTypeAndId.
func (mr *MockDatabaseRepositoryMockRecorder) FindResourceAssociationsByTypeAndId(ctx, source, resourceType, resourceId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindResourceAssociationsByTypeAndId", reflect.TypeOf((*MockDatabaseRepository)(nil).FindResourceAssociationsByTypeAndId), ctx, source, resourceType, resourceId)
}

// GetCurrentUser mocks base method.
func (m *MockDatabaseRepository) GetCurrentUser(ctx context.Context) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentUser", ctx)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrentUser indicates an expected call of GetCurrentUser.
func (mr *MockDatabaseRepositoryMockRecorder) GetCurrentUser(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentUser", reflect.TypeOf((*MockDatabaseRepository)(nil).GetCurrentUser), ctx)
}

// GetFlattenedResourceGraph mocks base method.
func (m *MockDatabaseRepository) GetFlattenedResourceGraph(ctx context.Context, graphType pkg.ResourceGraphType) (map[string][]*models.ResourceBase, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFlattenedResourceGraph", ctx, graphType)
	ret0, _ := ret[0].(map[string][]*models.ResourceBase)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFlattenedResourceGraph indicates an expected call of GetFlattenedResourceGraph.
func (mr *MockDatabaseRepositoryMockRecorder) GetFlattenedResourceGraph(ctx, graphType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFlattenedResourceGraph", reflect.TypeOf((*MockDatabaseRepository)(nil).GetFlattenedResourceGraph), ctx, graphType)
}

// GetGlossaryEntry mocks base method.
func (m *MockDatabaseRepository) GetGlossaryEntry(ctx context.Context, code, codeSystem string) (*models.Glossary, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGlossaryEntry", ctx, code, codeSystem)
	ret0, _ := ret[0].(*models.Glossary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGlossaryEntry indicates an expected call of GetGlossaryEntry.
func (mr *MockDatabaseRepositoryMockRecorder) GetGlossaryEntry(ctx, code, codeSystem interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGlossaryEntry", reflect.TypeOf((*MockDatabaseRepository)(nil).GetGlossaryEntry), ctx, code, codeSystem)
}

// GetPatientForSources mocks base method.
func (m *MockDatabaseRepository) GetPatientForSources(ctx context.Context) ([]models.ResourceBase, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPatientForSources", ctx)
	ret0, _ := ret[0].([]models.ResourceBase)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPatientForSources indicates an expected call of GetPatientForSources.
func (mr *MockDatabaseRepositoryMockRecorder) GetPatientForSources(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPatientForSources", reflect.TypeOf((*MockDatabaseRepository)(nil).GetPatientForSources), ctx)
}

// GetResourceByResourceTypeAndId mocks base method.
func (m *MockDatabaseRepository) GetResourceByResourceTypeAndId(arg0 context.Context, arg1, arg2 string) (*models.ResourceBase, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResourceByResourceTypeAndId", arg0, arg1, arg2)
	ret0, _ := ret[0].(*models.ResourceBase)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResourceByResourceTypeAndId indicates an expected call of GetResourceByResourceTypeAndId.
func (mr *MockDatabaseRepositoryMockRecorder) GetResourceByResourceTypeAndId(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResourceByResourceTypeAndId", reflect.TypeOf((*MockDatabaseRepository)(nil).GetResourceByResourceTypeAndId), arg0, arg1, arg2)
}

// GetResourceBySourceId mocks base method.
func (m *MockDatabaseRepository) GetResourceBySourceId(arg0 context.Context, arg1, arg2 string) (*models.ResourceBase, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResourceBySourceId", arg0, arg1, arg2)
	ret0, _ := ret[0].(*models.ResourceBase)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResourceBySourceId indicates an expected call of GetResourceBySourceId.
func (mr *MockDatabaseRepositoryMockRecorder) GetResourceBySourceId(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResourceBySourceId", reflect.TypeOf((*MockDatabaseRepository)(nil).GetResourceBySourceId), arg0, arg1, arg2)
}

// GetSource mocks base method.
func (m *MockDatabaseRepository) GetSource(arg0 context.Context, arg1 string) (*models.SourceCredential, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSource", arg0, arg1)
	ret0, _ := ret[0].(*models.SourceCredential)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSource indicates an expected call of GetSource.
func (mr *MockDatabaseRepositoryMockRecorder) GetSource(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSource", reflect.TypeOf((*MockDatabaseRepository)(nil).GetSource), arg0, arg1)
}

// GetSourceSummary mocks base method.
func (m *MockDatabaseRepository) GetSourceSummary(arg0 context.Context, arg1 string) (*models.SourceSummary, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSourceSummary", arg0, arg1)
	ret0, _ := ret[0].(*models.SourceSummary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSourceSummary indicates an expected call of GetSourceSummary.
func (mr *MockDatabaseRepositoryMockRecorder) GetSourceSummary(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSourceSummary", reflect.TypeOf((*MockDatabaseRepository)(nil).GetSourceSummary), arg0, arg1)
}

// GetSources mocks base method.
func (m *MockDatabaseRepository) GetSources(arg0 context.Context) ([]models.SourceCredential, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSources", arg0)
	ret0, _ := ret[0].([]models.SourceCredential)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSources indicates an expected call of GetSources.
func (mr *MockDatabaseRepositoryMockRecorder) GetSources(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSources", reflect.TypeOf((*MockDatabaseRepository)(nil).GetSources), arg0)
}

// GetSummary mocks base method.
func (m *MockDatabaseRepository) GetSummary(ctx context.Context) (*models.Summary, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSummary", ctx)
	ret0, _ := ret[0].(*models.Summary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSummary indicates an expected call of GetSummary.
func (mr *MockDatabaseRepositoryMockRecorder) GetSummary(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSummary", reflect.TypeOf((*MockDatabaseRepository)(nil).GetSummary), ctx)
}

// GetUserByUsername mocks base method.
func (m *MockDatabaseRepository) GetUserByUsername(arg0 context.Context, arg1 string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByUsername", arg0, arg1)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByUsername indicates an expected call of GetUserByUsername.
func (mr *MockDatabaseRepositoryMockRecorder) GetUserByUsername(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByUsername", reflect.TypeOf((*MockDatabaseRepository)(nil).GetUserByUsername), arg0, arg1)
}

// ListResources mocks base method.
func (m *MockDatabaseRepository) ListResources(arg0 context.Context, arg1 models.ListResourceQueryOptions) ([]models.ResourceBase, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListResources", arg0, arg1)
	ret0, _ := ret[0].([]models.ResourceBase)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListResources indicates an expected call of ListResources.
func (mr *MockDatabaseRepositoryMockRecorder) ListResources(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListResources", reflect.TypeOf((*MockDatabaseRepository)(nil).ListResources), arg0, arg1)
}

// LoadUserSettings mocks base method.
func (m *MockDatabaseRepository) LoadUserSettings(ctx context.Context) (*models.UserSettings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadUserSettings", ctx)
	ret0, _ := ret[0].(*models.UserSettings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadUserSettings indicates an expected call of LoadUserSettings.
func (mr *MockDatabaseRepositoryMockRecorder) LoadUserSettings(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadUserSettings", reflect.TypeOf((*MockDatabaseRepository)(nil).LoadUserSettings), ctx)
}

// Migrate mocks base method.
func (m *MockDatabaseRepository) Migrate() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Migrate")
	ret0, _ := ret[0].(error)
	return ret0
}

// Migrate indicates an expected call of Migrate.
func (mr *MockDatabaseRepositoryMockRecorder) Migrate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Migrate", reflect.TypeOf((*MockDatabaseRepository)(nil).Migrate))
}

// PopulateDefaultUserSettings mocks base method.
func (m *MockDatabaseRepository) PopulateDefaultUserSettings(ctx context.Context, userId uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PopulateDefaultUserSettings", ctx, userId)
	ret0, _ := ret[0].(error)
	return ret0
}

// PopulateDefaultUserSettings indicates an expected call of PopulateDefaultUserSettings.
func (mr *MockDatabaseRepositoryMockRecorder) PopulateDefaultUserSettings(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PopulateDefaultUserSettings", reflect.TypeOf((*MockDatabaseRepository)(nil).PopulateDefaultUserSettings), ctx, userId)
}

// QueryResources mocks base method.
func (m *MockDatabaseRepository) QueryResources(ctx context.Context, query models.QueryResource) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryResources", ctx, query)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryResources indicates an expected call of QueryResources.
func (mr *MockDatabaseRepositoryMockRecorder) QueryResources(ctx, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryResources", reflect.TypeOf((*MockDatabaseRepository)(nil).QueryResources), ctx, query)
}

// RemoveResourceAssociation mocks base method.
func (m *MockDatabaseRepository) RemoveResourceAssociation(ctx context.Context, source *models.SourceCredential, resourceType, resourceId string, relatedSource *models.SourceCredential, relatedResourceType, relatedResourceId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveResourceAssociation", ctx, source, resourceType, resourceId, relatedSource, relatedResourceType, relatedResourceId)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveResourceAssociation indicates an expected call of RemoveResourceAssociation.
func (mr *MockDatabaseRepositoryMockRecorder) RemoveResourceAssociation(ctx, source, resourceType, resourceId, relatedSource, relatedResourceType, relatedResourceId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveResourceAssociation", reflect.TypeOf((*MockDatabaseRepository)(nil).RemoveResourceAssociation), ctx, source, resourceType, resourceId, relatedSource, relatedResourceType, relatedResourceId)
}

// SaveUserSettings mocks base method.
func (m *MockDatabaseRepository) SaveUserSettings(arg0 context.Context, arg1 *models.UserSettings) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveUserSettings", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveUserSettings indicates an expected call of SaveUserSettings.
func (mr *MockDatabaseRepositoryMockRecorder) SaveUserSettings(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveUserSettings", reflect.TypeOf((*MockDatabaseRepository)(nil).SaveUserSettings), arg0, arg1)
}

// UpdateSource mocks base method.
func (m *MockDatabaseRepository) UpdateSource(ctx context.Context, sourceCreds *models.SourceCredential) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSource", ctx, sourceCreds)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSource indicates an expected call of UpdateSource.
func (mr *MockDatabaseRepositoryMockRecorder) UpdateSource(ctx, sourceCreds interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSource", reflect.TypeOf((*MockDatabaseRepository)(nil).UpdateSource), ctx, sourceCreds)
}

// UpsertRawResource mocks base method.
func (m *MockDatabaseRepository) UpsertRawResource(ctx context.Context, sourceCredentials models0.SourceCredential, rawResource models0.RawResourceFhir) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertRawResource", ctx, sourceCredentials, rawResource)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpsertRawResource indicates an expected call of UpsertRawResource.
func (mr *MockDatabaseRepositoryMockRecorder) UpsertRawResource(ctx, sourceCredentials, rawResource interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertRawResource", reflect.TypeOf((*MockDatabaseRepository)(nil).UpsertRawResource), ctx, sourceCredentials, rawResource)
}
