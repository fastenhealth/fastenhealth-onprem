// THIS FILE IS GENERATED BY https://github.com/fastenhealth/fasten-onprem/blob/main/backend/pkg/models/database/generate.go
// PLEASE DO NOT EDIT BY HAND

package database

import (
	"encoding/json"
	"fmt"
	goja "github.com/dop251/goja"
	models "github.com/fastenhealth/fasten-onprem/backend/pkg/models"
	datatypes "gorm.io/datatypes"
	"time"
)

type FhirMedicationStatement struct {
	models.ResourceBase
	// Returns statements of this category of medicationstatement
	// https://hl7.org/fhir/r4/search.html#token
	Category datatypes.JSON `gorm:"column:category;type:text;serializer:json" json:"category,omitempty"`
	/*
	   Multiple Resources:

	   * [AllergyIntolerance](allergyintolerance.html): Code that identifies the allergy or intolerance
	   * [Condition](condition.html): Code for the condition
	   * [DeviceRequest](devicerequest.html): Code for what is being requested/ordered
	   * [DiagnosticReport](diagnosticreport.html): The code for the report, as opposed to codes for the atomic results, which are the names on the observation resource referred to from the result
	   * [FamilyMemberHistory](familymemberhistory.html): A search by a condition code
	   * [List](list.html): What the purpose of this list is
	   * [Medication](medication.html): Returns medications for a specific code
	   * [MedicationAdministration](medicationadministration.html): Return administrations of this medication code
	   * [MedicationDispense](medicationdispense.html): Returns dispenses of this medicine code
	   * [MedicationRequest](medicationrequest.html): Return prescriptions of this medication code
	   * [MedicationStatement](medicationstatement.html): Return statements of this medication code
	   * [Observation](observation.html): The code of the observation type
	   * [Procedure](procedure.html): A code to identify a  procedure
	   * [ServiceRequest](servicerequest.html): What is being requested/ordered
	*/
	// https://hl7.org/fhir/r4/search.html#token
	Code datatypes.JSON `gorm:"column:code;type:text;serializer:json" json:"code,omitempty"`
	// Returns statements for a specific context (episode or episode of Care).
	// https://hl7.org/fhir/r4/search.html#reference
	Context datatypes.JSON `gorm:"column:context;type:text;serializer:json" json:"context,omitempty"`
	// Date when patient was taking (or not taking) the medication
	// https://hl7.org/fhir/r4/search.html#date
	Effective *time.Time `gorm:"column:effective;type:datetime" json:"effective,omitempty"`
	/*
	   Multiple Resources:

	   * [AllergyIntolerance](allergyintolerance.html): External ids for this item
	   * [CarePlan](careplan.html): External Ids for this plan
	   * [CareTeam](careteam.html): External Ids for this team
	   * [Composition](composition.html): Version-independent identifier for the Composition
	   * [Condition](condition.html): A unique identifier of the condition record
	   * [Consent](consent.html): Identifier for this record (external references)
	   * [DetectedIssue](detectedissue.html): Unique id for the detected issue
	   * [DeviceRequest](devicerequest.html): Business identifier for request/order
	   * [DiagnosticReport](diagnosticreport.html): An identifier for the report
	   * [DocumentManifest](documentmanifest.html): Unique Identifier for the set of documents
	   * [DocumentReference](documentreference.html): Master Version Specific Identifier
	   * [Encounter](encounter.html): Identifier(s) by which this encounter is known
	   * [EpisodeOfCare](episodeofcare.html): Business Identifier(s) relevant for this EpisodeOfCare
	   * [FamilyMemberHistory](familymemberhistory.html): A search by a record identifier
	   * [Goal](goal.html): External Ids for this goal
	   * [ImagingStudy](imagingstudy.html): Identifiers for the Study, such as DICOM Study Instance UID and Accession number
	   * [Immunization](immunization.html): Business identifier
	   * [List](list.html): Business identifier
	   * [MedicationAdministration](medicationadministration.html): Return administrations with this external identifier
	   * [MedicationDispense](medicationdispense.html): Returns dispenses with this external identifier
	   * [MedicationRequest](medicationrequest.html): Return prescriptions with this external identifier
	   * [MedicationStatement](medicationstatement.html): Return statements with this external identifier
	   * [NutritionOrder](nutritionorder.html): Return nutrition orders with this external identifier
	   * [Observation](observation.html): The unique id for a particular observation
	   * [Procedure](procedure.html): A unique identifier for a procedure
	   * [RiskAssessment](riskassessment.html): Unique identifier for the assessment
	   * [ServiceRequest](servicerequest.html): Identifiers assigned to this order
	   * [SupplyDelivery](supplydelivery.html): External identifier
	   * [SupplyRequest](supplyrequest.html): Business Identifier for SupplyRequest
	   * [VisionPrescription](visionprescription.html): Return prescriptions with this external identifier
	*/
	// https://hl7.org/fhir/r4/search.html#token
	Identifier datatypes.JSON `gorm:"column:identifier;type:text;serializer:json" json:"identifier,omitempty"`
	// Language of the resource content
	// https://hl7.org/fhir/r4/search.html#token
	Language datatypes.JSON `gorm:"column:language;type:text;serializer:json" json:"language,omitempty"`
	/*
	   Multiple Resources:

	   * [MedicationAdministration](medicationadministration.html): Return administrations of this medication resource
	   * [MedicationDispense](medicationdispense.html): Returns dispenses of this medicine resource
	   * [MedicationRequest](medicationrequest.html): Return prescriptions for this medication reference
	   * [MedicationStatement](medicationstatement.html): Return statements of this medication reference
	*/
	// https://hl7.org/fhir/r4/search.html#reference
	Medication datatypes.JSON `gorm:"column:medication;type:text;serializer:json" json:"medication,omitempty"`
	// When the resource version last changed
	// https://hl7.org/fhir/r4/search.html#date
	MetaLastUpdated *time.Time `gorm:"column:metaLastUpdated;type:datetime" json:"metaLastUpdated,omitempty"`
	// Profiles this resource claims to conform to
	// https://hl7.org/fhir/r4/search.html#reference
	MetaProfile datatypes.JSON `gorm:"column:metaProfile;type:text;serializer:json" json:"metaProfile,omitempty"`
	// Tags applied to this resource
	// https://hl7.org/fhir/r4/search.html#token
	MetaTag datatypes.JSON `gorm:"column:metaTag;type:text;serializer:json" json:"metaTag,omitempty"`
	// Tags applied to this resource
	// https://hl7.org/fhir/r4/search.html#keyword
	MetaVersionId string `gorm:"column:metaVersionId;type:text" json:"metaVersionId,omitempty"`
	// Returns statements that are part of another event.
	// https://hl7.org/fhir/r4/search.html#reference
	PartOf datatypes.JSON `gorm:"column:partOf;type:text;serializer:json" json:"partOf,omitempty"`
	// Who or where the information in the statement came from
	// https://hl7.org/fhir/r4/search.html#reference
	Source datatypes.JSON `gorm:"column:source;type:text;serializer:json" json:"source,omitempty"`
	/*
	   Multiple Resources:

	   * [MedicationAdministration](medicationadministration.html): MedicationAdministration event status (for example one of active/paused/completed/nullified)
	   * [MedicationDispense](medicationdispense.html): Returns dispenses with a specified dispense status
	   * [MedicationRequest](medicationrequest.html): Status of the prescription
	   * [MedicationStatement](medicationstatement.html): Return statements that match the given status
	*/
	// https://hl7.org/fhir/r4/search.html#token
	Status datatypes.JSON `gorm:"column:status;type:text;serializer:json" json:"status,omitempty"`
	// The identity of a patient, animal or group to list statements for
	// https://hl7.org/fhir/r4/search.html#reference
	Subject datatypes.JSON `gorm:"column:subject;type:text;serializer:json" json:"subject,omitempty"`
	// Text search against the narrative
	// https://hl7.org/fhir/r4/search.html#string
	Text datatypes.JSON `gorm:"column:text;type:text;serializer:json" json:"text,omitempty"`
	// A resource type filter
	// https://hl7.org/fhir/r4/search.html#special
	Type datatypes.JSON `gorm:"column:type;type:text;serializer:json" json:"type,omitempty"`
}

func (s *FhirMedicationStatement) GetSearchParameters() map[string]string {
	searchParameters := map[string]string{
		"category":             "token",
		"code":                 "token",
		"context":              "reference",
		"effective":            "date",
		"id":                   "keyword",
		"identifier":           "token",
		"language":             "token",
		"medication":           "reference",
		"metaLastUpdated":      "date",
		"metaProfile":          "reference",
		"metaTag":              "token",
		"metaVersionId":        "keyword",
		"partOf":               "reference",
		"sort_date":            "date",
		"source":               "reference",
		"source_id":            "keyword",
		"source_resource_id":   "keyword",
		"source_resource_type": "keyword",
		"source_uri":           "keyword",
		"status":               "token",
		"subject":              "reference",
		"text":                 "string",
		"type":                 "special",
	}
	return searchParameters
}
func (s *FhirMedicationStatement) PopulateAndExtractSearchParameters(resourceRaw json.RawMessage) error {
	s.ResourceRaw = datatypes.JSON(resourceRaw)
	// unmarshal the raw resource (bytes) into a map
	var resourceRawMap map[string]interface{}
	err := json.Unmarshal(resourceRaw, &resourceRawMap)
	if err != nil {
		return err
	}
	if len(fhirPathJs) == 0 {
		return fmt.Errorf("fhirPathJs script is empty")
	}
	vm := goja.New()
	// setup the global window object
	vm.Set("window", vm.NewObject())
	// set the global FHIR Resource object
	vm.Set("fhirResource", resourceRawMap)
	// compile the fhirpath library
	fhirPathJsProgram, err := goja.Compile("fhirpath.min.js", fhirPathJs, true)
	if err != nil {
		return err
	}
	// compile the searchParametersExtractor library
	searchParametersExtractorJsProgram, err := goja.Compile("searchParameterExtractor.js", searchParameterExtractorJs, true)
	if err != nil {
		return err
	}
	// add the fhirpath library in the goja vm
	_, err = vm.RunProgram(fhirPathJsProgram)
	if err != nil {
		return err
	}
	// add the searchParametersExtractor library in the goja vm
	_, err = vm.RunProgram(searchParametersExtractorJsProgram)
	if err != nil {
		return err
	}
	// execute the fhirpath expression for each search parameter
	// extracting Category
	categoryResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'MedicationStatement.category')")
	if err == nil && categoryResult.String() != "undefined" {
		s.Category = []byte(categoryResult.String())
	}
	// extracting Code
	codeResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'AllergyIntolerance.code | AllergyIntolerance.reaction.substance | Condition.code | (DeviceRequest.codeCodeableConcept) | DiagnosticReport.code | FamilyMemberHistory.condition.code | List.code | Medication.code | (MedicationAdministration.medicationCodeableConcept) | (MedicationDispense.medicationCodeableConcept) | (MedicationRequest.medicationCodeableConcept) | (MedicationStatement.medicationCodeableConcept) | Observation.code | Procedure.code | ServiceRequest.code')")
	if err == nil && codeResult.String() != "undefined" {
		s.Code = []byte(codeResult.String())
	}
	// extracting Context
	contextResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'MedicationStatement.context')")
	if err == nil && contextResult.String() != "undefined" {
		s.Context = []byte(contextResult.String())
	}
	// extracting Effective
	effectiveResult, err := vm.RunString("extractSimpleSearchParameters(fhirResource, 'MedicationStatement.effective')")
	if err == nil && effectiveResult.String() != "undefined" {
		t, err := time.Parse(time.RFC3339, effectiveResult.String())
		if err == nil {
			s.Effective = &t
		} else if err != nil {
			d, err := time.Parse("2006-01-02", effectiveResult.String())
			if err == nil {
				s.Effective = &d
			}
		}
	}
	// extracting Identifier
	identifierResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'AllergyIntolerance.identifier | CarePlan.identifier | CareTeam.identifier | Composition.identifier | Condition.identifier | Consent.identifier | DetectedIssue.identifier | DeviceRequest.identifier | DiagnosticReport.identifier | DocumentManifest.masterIdentifier | DocumentManifest.identifier | DocumentReference.masterIdentifier | DocumentReference.identifier | Encounter.identifier | EpisodeOfCare.identifier | FamilyMemberHistory.identifier | Goal.identifier | ImagingStudy.identifier | Immunization.identifier | List.identifier | MedicationAdministration.identifier | MedicationDispense.identifier | MedicationRequest.identifier | MedicationStatement.identifier | NutritionOrder.identifier | Observation.identifier | Procedure.identifier | RiskAssessment.identifier | ServiceRequest.identifier | SupplyDelivery.identifier | SupplyRequest.identifier | VisionPrescription.identifier')")
	if err == nil && identifierResult.String() != "undefined" {
		s.Identifier = []byte(identifierResult.String())
	}
	// extracting Language
	languageResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'language')")
	if err == nil && languageResult.String() != "undefined" {
		s.Language = []byte(languageResult.String())
	}
	// extracting Medication
	medicationResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, '(MedicationAdministration.medicationReference) | (MedicationDispense.medicationReference) | (MedicationRequest.medicationReference) | (MedicationStatement.medicationReference)')")
	if err == nil && medicationResult.String() != "undefined" {
		s.Medication = []byte(medicationResult.String())
	}
	// extracting MetaLastUpdated
	metaLastUpdatedResult, err := vm.RunString("extractSimpleSearchParameters(fhirResource, 'meta.lastUpdated')")
	if err == nil && metaLastUpdatedResult.String() != "undefined" {
		t, err := time.Parse(time.RFC3339, metaLastUpdatedResult.String())
		if err == nil {
			s.MetaLastUpdated = &t
		} else if err != nil {
			d, err := time.Parse("2006-01-02", metaLastUpdatedResult.String())
			if err == nil {
				s.MetaLastUpdated = &d
			}
		}
	}
	// extracting MetaProfile
	metaProfileResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'meta.profile')")
	if err == nil && metaProfileResult.String() != "undefined" {
		s.MetaProfile = []byte(metaProfileResult.String())
	}
	// extracting MetaTag
	metaTagResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'meta.tag')")
	if err == nil && metaTagResult.String() != "undefined" {
		s.MetaTag = []byte(metaTagResult.String())
	}
	// extracting MetaVersionId
	metaVersionIdResult, err := vm.RunString("extractSimpleSearchParameters(fhirResource, 'meta.versionId')")
	if err == nil && metaVersionIdResult.String() != "undefined" {
		s.MetaVersionId = metaVersionIdResult.String()
	}
	// extracting PartOf
	partOfResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'MedicationStatement.partOf')")
	if err == nil && partOfResult.String() != "undefined" {
		s.PartOf = []byte(partOfResult.String())
	}
	// extracting Source
	sourceResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'MedicationStatement.informationSource')")
	if err == nil && sourceResult.String() != "undefined" {
		s.Source = []byte(sourceResult.String())
	}
	// extracting Status
	statusResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'MedicationAdministration.status | MedicationDispense.status | MedicationRequest.status | MedicationStatement.status')")
	if err == nil && statusResult.String() != "undefined" {
		s.Status = []byte(statusResult.String())
	}
	// extracting Subject
	subjectResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'MedicationStatement.subject')")
	if err == nil && subjectResult.String() != "undefined" {
		s.Subject = []byte(subjectResult.String())
	}
	return nil
}

// TableName overrides the table name from fhir_observations (pluralized) to `fhir_observation`. https://gorm.io/docs/conventions.html#TableName
func (s *FhirMedicationStatement) TableName() string {
	return "fhir_medication_statement"
}
