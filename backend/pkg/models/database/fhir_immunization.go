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

type FhirImmunization struct {
	models.ResourceBase
	/*
	   Multiple Resources:

	   * [AllergyIntolerance](allergyintolerance.html): Date first version of the resource instance was recorded
	   * [CarePlan](careplan.html): Time period plan covers
	   * [CareTeam](careteam.html): Time period team covers
	   * [ClinicalImpression](clinicalimpression.html): When the assessment was documented
	   * [Composition](composition.html): Composition editing time
	   * [Consent](consent.html): When this Consent was created or indexed
	   * [DiagnosticReport](diagnosticreport.html): The clinically relevant time of the report
	   * [Encounter](encounter.html): A date within the period the Encounter lasted
	   * [EpisodeOfCare](episodeofcare.html): The provided date search value falls within the episode of care's period
	   * [FamilyMemberHistory](familymemberhistory.html): When history was recorded or last updated
	   * [Flag](flag.html): Time period when flag is active
	   * [Immunization](immunization.html): Vaccination  (non)-Administration Date
	   * [List](list.html): When the list was prepared
	   * [Observation](observation.html): Obtained date/time. If the obtained element is a period, a date that falls in the period
	   * [Procedure](procedure.html): When the procedure was performed
	   * [RiskAssessment](riskassessment.html): When was assessment made?
	   * [SupplyRequest](supplyrequest.html): When the request was made
	*/
	// https://hl7.org/fhir/r4/search.html#date
	Date *time.Time `gorm:"column:date;type:datetime" json:"date,omitempty"`
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
	// When the resource version last changed
	// https://hl7.org/fhir/r4/search.html#date
	LastUpdated *time.Time `gorm:"column:lastUpdated;type:datetime" json:"lastUpdated,omitempty"`
	// The service delivery location or facility in which the vaccine was / was to be administered
	// https://hl7.org/fhir/r4/search.html#reference
	Location datatypes.JSON `gorm:"column:location;type:text;serializer:json" json:"location,omitempty"`
	// Vaccine Lot Number
	// https://hl7.org/fhir/r4/search.html#string
	LotNumber datatypes.JSON `gorm:"column:lotNumber;type:text;serializer:json" json:"lotNumber,omitempty"`
	// Vaccine Manufacturer
	// https://hl7.org/fhir/r4/search.html#reference
	Manufacturer datatypes.JSON `gorm:"column:manufacturer;type:text;serializer:json" json:"manufacturer,omitempty"`
	// The practitioner or organization who played a role in the vaccination
	// https://hl7.org/fhir/r4/search.html#reference
	Performer datatypes.JSON `gorm:"column:performer;type:text;serializer:json" json:"performer,omitempty"`
	// Profiles this resource claims to conform to
	// https://hl7.org/fhir/r4/search.html#reference
	Profile datatypes.JSON `gorm:"column:profile;type:text;serializer:json" json:"profile,omitempty"`
	// Additional information on reaction
	// https://hl7.org/fhir/r4/search.html#reference
	Reaction datatypes.JSON `gorm:"column:reaction;type:text;serializer:json" json:"reaction,omitempty"`
	// When reaction started
	// https://hl7.org/fhir/r4/search.html#date
	ReactionDate *time.Time `gorm:"column:reactionDate;type:datetime" json:"reactionDate,omitempty"`
	// Reason why the vaccine was administered
	// https://hl7.org/fhir/r4/search.html#token
	ReasonCode datatypes.JSON `gorm:"column:reasonCode;type:text;serializer:json" json:"reasonCode,omitempty"`
	// Why immunization occurred
	// https://hl7.org/fhir/r4/search.html#reference
	ReasonReference datatypes.JSON `gorm:"column:reasonReference;type:text;serializer:json" json:"reasonReference,omitempty"`
	// The series being followed by the provider
	// https://hl7.org/fhir/r4/search.html#string
	Series datatypes.JSON `gorm:"column:series;type:text;serializer:json" json:"series,omitempty"`
	// Immunization event status
	// https://hl7.org/fhir/r4/search.html#token
	Status datatypes.JSON `gorm:"column:status;type:text;serializer:json" json:"status,omitempty"`
	// Reason why the vaccine was not administered
	// https://hl7.org/fhir/r4/search.html#token
	StatusReason datatypes.JSON `gorm:"column:statusReason;type:text;serializer:json" json:"statusReason,omitempty"`
	// Tags applied to this resource
	// https://hl7.org/fhir/r4/search.html#token
	Tag datatypes.JSON `gorm:"column:tag;type:text;serializer:json" json:"tag,omitempty"`
	// The target disease the dose is being administered against
	// https://hl7.org/fhir/r4/search.html#token
	TargetDisease datatypes.JSON `gorm:"column:targetDisease;type:text;serializer:json" json:"targetDisease,omitempty"`
	// Text search against the narrative
	// https://hl7.org/fhir/r4/search.html#string
	Text datatypes.JSON `gorm:"column:text;type:text;serializer:json" json:"text,omitempty"`
	// A resource type filter
	// https://hl7.org/fhir/r4/search.html#special
	Type datatypes.JSON `gorm:"column:type;type:text;serializer:json" json:"type,omitempty"`
	// Vaccine Product Administered
	// https://hl7.org/fhir/r4/search.html#token
	VaccineCode datatypes.JSON `gorm:"column:vaccineCode;type:text;serializer:json" json:"vaccineCode,omitempty"`
}

func (s *FhirImmunization) GetSearchParameters() map[string]string {
	searchParameters := map[string]string{
		"date":                 "date",
		"id":                   "keyword",
		"identifier":           "token",
		"language":             "token",
		"lastUpdated":          "date",
		"location":             "reference",
		"lotNumber":            "string",
		"manufacturer":         "reference",
		"performer":            "reference",
		"profile":              "reference",
		"reaction":             "reference",
		"reactionDate":         "date",
		"reasonCode":           "token",
		"reasonReference":      "reference",
		"series":               "string",
		"source_id":            "keyword",
		"source_resource_id":   "keyword",
		"source_resource_type": "keyword",
		"source_uri":           "keyword",
		"status":               "token",
		"statusReason":         "token",
		"tag":                  "token",
		"targetDisease":        "token",
		"text":                 "string",
		"type":                 "special",
		"vaccineCode":          "token",
	}
	return searchParameters
}
func (s *FhirImmunization) PopulateAndExtractSearchParameters(resourceRaw json.RawMessage) error {
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
	// add the fhirpath library in the goja vm
	_, err = vm.RunProgram(fhirPathJsProgram)
	if err != nil {
		return err
	}
	// execute the fhirpath expression for each search parameter
	// extracting Date
	dateResult, err := vm.RunString("window.fhirpath.evaluate(fhirResource, 'AllergyIntolerance.recordedDate | CarePlan.period | CareTeam.period | ClinicalImpression.date | Composition.date | Consent.dateTime | DiagnosticReport.effective | Encounter.period | EpisodeOfCare.period | FamilyMemberHistory.date | Flag.period | (Immunization.occurrenceDateTime) | List.date | Observation.effective | Procedure.performed | (RiskAssessment.occurrenceDateTime) | SupplyRequest.authoredOn')[0]")
	if err == nil && dateResult.String() != "undefined" {
		t, err := time.Parse(time.RFC3339, dateResult.String())
		if err == nil {
			s.Date = &t
		} else if err != nil {
			d, err := time.Parse("2006-01-02", dateResult.String())
			if err == nil {
				s.Date = &d
			}
		}
	}
	// extracting Identifier
	identifierResult, err := vm.RunString(` 
							IdentifierResult = window.fhirpath.evaluate(fhirResource, 'AllergyIntolerance.identifier | CarePlan.identifier | CareTeam.identifier | Composition.identifier | Condition.identifier | Consent.identifier | DetectedIssue.identifier | DeviceRequest.identifier | DiagnosticReport.identifier | DocumentManifest.masterIdentifier | DocumentManifest.identifier | DocumentReference.masterIdentifier | DocumentReference.identifier | Encounter.identifier | EpisodeOfCare.identifier | FamilyMemberHistory.identifier | Goal.identifier | ImagingStudy.identifier | Immunization.identifier | List.identifier | MedicationAdministration.identifier | MedicationDispense.identifier | MedicationRequest.identifier | MedicationStatement.identifier | NutritionOrder.identifier | Observation.identifier | Procedure.identifier | RiskAssessment.identifier | ServiceRequest.identifier | SupplyDelivery.identifier | SupplyRequest.identifier | VisionPrescription.identifier')
							IdentifierProcessed = IdentifierResult.reduce((accumulator, currentValue) => {
								if (currentValue.coding) {
									//CodeableConcept
									currentValue.coding.map((coding) => {
										accumulator.push({
											"code": coding.code,	
											"system": coding.system,
											"text": currentValue.text
										})
									})
								} else if (currentValue.value) {
									//ContactPoint, Identifier
									accumulator.push({
										"code": currentValue.value,
										"system": currentValue.system,
										"text": currentValue.type?.text
									})
								} else if (currentValue.code) {
									//Coding
									accumulator.push({
										"code": currentValue.code,
										"system": currentValue.system,
										"text": currentValue.display
									})
								} else if ((typeof currentValue === 'string') || (typeof currentValue === 'boolean')) {
									//string, boolean
									accumulator.push({
										"code": currentValue,
									})
								}
								return accumulator
							}, [])
						
				
							if(IdentifierProcessed.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(IdentifierProcessed)
							}
						 `)
	if err == nil && identifierResult.String() != "undefined" {
		s.Identifier = []byte(identifierResult.String())
	}
	// extracting Language
	languageResult, err := vm.RunString(` 
							LanguageResult = window.fhirpath.evaluate(fhirResource, 'language')
							LanguageProcessed = LanguageResult.reduce((accumulator, currentValue) => {
								if (currentValue.coding) {
									//CodeableConcept
									currentValue.coding.map((coding) => {
										accumulator.push({
											"code": coding.code,	
											"system": coding.system,
											"text": currentValue.text
										})
									})
								} else if (currentValue.value) {
									//ContactPoint, Identifier
									accumulator.push({
										"code": currentValue.value,
										"system": currentValue.system,
										"text": currentValue.type?.text
									})
								} else if (currentValue.code) {
									//Coding
									accumulator.push({
										"code": currentValue.code,
										"system": currentValue.system,
										"text": currentValue.display
									})
								} else if ((typeof currentValue === 'string') || (typeof currentValue === 'boolean')) {
									//string, boolean
									accumulator.push({
										"code": currentValue,
									})
								}
								return accumulator
							}, [])
						
				
							if(LanguageProcessed.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(LanguageProcessed)
							}
						 `)
	if err == nil && languageResult.String() != "undefined" {
		s.Language = []byte(languageResult.String())
	}
	// extracting LastUpdated
	lastUpdatedResult, err := vm.RunString("window.fhirpath.evaluate(fhirResource, 'meta.lastUpdated')[0]")
	if err == nil && lastUpdatedResult.String() != "undefined" {
		t, err := time.Parse(time.RFC3339, lastUpdatedResult.String())
		if err == nil {
			s.LastUpdated = &t
		} else if err != nil {
			d, err := time.Parse("2006-01-02", lastUpdatedResult.String())
			if err == nil {
				s.LastUpdated = &d
			}
		}
	}
	// extracting Location
	locationResult, err := vm.RunString(` 
							LocationResult = window.fhirpath.evaluate(fhirResource, 'Immunization.location')
						
							if(LocationResult.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(LocationResult)
							}
						 `)
	if err == nil && locationResult.String() != "undefined" {
		s.Location = []byte(locationResult.String())
	}
	// extracting LotNumber
	lotNumberResult, err := vm.RunString(` 
							LotNumberResult = window.fhirpath.evaluate(fhirResource, 'Immunization.lotNumber')
							LotNumberProcessed = LotNumberResult.reduce((accumulator, currentValue) => {
								if (typeof currentValue === 'string') {
									//basic string
									accumulator.push(currentValue)
								} else if (currentValue.family  || currentValue.given) {
									//HumanName http://hl7.org/fhir/R4/datatypes.html#HumanName
									var humanNameParts = []
									if (currentValue.prefix) {
										humanNameParts = humanNameParts.concat(currentValue.prefix)
									}
									if (currentValue.given) {	
										humanNameParts = humanNameParts.concat(currentValue.given)
									}	
									if (currentValue.family) {	
										humanNameParts.push(currentValue.family)	
									}	
									if (currentValue.suffix) {	
										humanNameParts = humanNameParts.concat(currentValue.suffix)	
									}
									accumulator.push(humanNameParts.join(" "))
								} else if (currentValue.city || currentValue.state || currentValue.country || currentValue.postalCode) {
									//Address http://hl7.org/fhir/R4/datatypes.html#Address
									var addressParts = []		
									if (currentValue.line) {
										addressParts = addressParts.concat(currentValue.line)
									}
									if (currentValue.city) {
										addressParts.push(currentValue.city)
									}	
									if (currentValue.state) {	
										addressParts.push(currentValue.state)
									}	
									if (currentValue.postalCode) {
										addressParts.push(currentValue.postalCode)
									}	
									if (currentValue.country) {
										addressParts.push(currentValue.country)	
									}	
									accumulator.push(addressParts.join(" "))
								} else {
									//string, boolean
									accumulator.push(currentValue)
								}
								return accumulator
							}, [])
						
							if(LotNumberProcessed.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(LotNumberProcessed)
							}
						 `)
	if err == nil && lotNumberResult.String() != "undefined" {
		s.LotNumber = []byte(lotNumberResult.String())
	}
	// extracting Manufacturer
	manufacturerResult, err := vm.RunString(` 
							ManufacturerResult = window.fhirpath.evaluate(fhirResource, 'Immunization.manufacturer')
						
							if(ManufacturerResult.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(ManufacturerResult)
							}
						 `)
	if err == nil && manufacturerResult.String() != "undefined" {
		s.Manufacturer = []byte(manufacturerResult.String())
	}
	// extracting Performer
	performerResult, err := vm.RunString(` 
							PerformerResult = window.fhirpath.evaluate(fhirResource, 'Immunization.performer.actor')
						
							if(PerformerResult.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(PerformerResult)
							}
						 `)
	if err == nil && performerResult.String() != "undefined" {
		s.Performer = []byte(performerResult.String())
	}
	// extracting Profile
	profileResult, err := vm.RunString(` 
							ProfileResult = window.fhirpath.evaluate(fhirResource, 'meta.profile')
						
							if(ProfileResult.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(ProfileResult)
							}
						 `)
	if err == nil && profileResult.String() != "undefined" {
		s.Profile = []byte(profileResult.String())
	}
	// extracting Reaction
	reactionResult, err := vm.RunString(` 
							ReactionResult = window.fhirpath.evaluate(fhirResource, 'Immunization.reaction.detail')
						
							if(ReactionResult.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(ReactionResult)
							}
						 `)
	if err == nil && reactionResult.String() != "undefined" {
		s.Reaction = []byte(reactionResult.String())
	}
	// extracting ReactionDate
	reactionDateResult, err := vm.RunString("window.fhirpath.evaluate(fhirResource, 'Immunization.reaction.date')[0]")
	if err == nil && reactionDateResult.String() != "undefined" {
		t, err := time.Parse(time.RFC3339, reactionDateResult.String())
		if err == nil {
			s.ReactionDate = &t
		} else if err != nil {
			d, err := time.Parse("2006-01-02", reactionDateResult.String())
			if err == nil {
				s.ReactionDate = &d
			}
		}
	}
	// extracting ReasonCode
	reasonCodeResult, err := vm.RunString(` 
							ReasonCodeResult = window.fhirpath.evaluate(fhirResource, 'Immunization.reasonCode')
							ReasonCodeProcessed = ReasonCodeResult.reduce((accumulator, currentValue) => {
								if (currentValue.coding) {
									//CodeableConcept
									currentValue.coding.map((coding) => {
										accumulator.push({
											"code": coding.code,	
											"system": coding.system,
											"text": currentValue.text
										})
									})
								} else if (currentValue.value) {
									//ContactPoint, Identifier
									accumulator.push({
										"code": currentValue.value,
										"system": currentValue.system,
										"text": currentValue.type?.text
									})
								} else if (currentValue.code) {
									//Coding
									accumulator.push({
										"code": currentValue.code,
										"system": currentValue.system,
										"text": currentValue.display
									})
								} else if ((typeof currentValue === 'string') || (typeof currentValue === 'boolean')) {
									//string, boolean
									accumulator.push({
										"code": currentValue,
									})
								}
								return accumulator
							}, [])
						
				
							if(ReasonCodeProcessed.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(ReasonCodeProcessed)
							}
						 `)
	if err == nil && reasonCodeResult.String() != "undefined" {
		s.ReasonCode = []byte(reasonCodeResult.String())
	}
	// extracting ReasonReference
	reasonReferenceResult, err := vm.RunString(` 
							ReasonReferenceResult = window.fhirpath.evaluate(fhirResource, 'Immunization.reasonReference')
						
							if(ReasonReferenceResult.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(ReasonReferenceResult)
							}
						 `)
	if err == nil && reasonReferenceResult.String() != "undefined" {
		s.ReasonReference = []byte(reasonReferenceResult.String())
	}
	// extracting Series
	seriesResult, err := vm.RunString(` 
							SeriesResult = window.fhirpath.evaluate(fhirResource, 'Immunization.protocolApplied.series')
							SeriesProcessed = SeriesResult.reduce((accumulator, currentValue) => {
								if (typeof currentValue === 'string') {
									//basic string
									accumulator.push(currentValue)
								} else if (currentValue.family  || currentValue.given) {
									//HumanName http://hl7.org/fhir/R4/datatypes.html#HumanName
									var humanNameParts = []
									if (currentValue.prefix) {
										humanNameParts = humanNameParts.concat(currentValue.prefix)
									}
									if (currentValue.given) {	
										humanNameParts = humanNameParts.concat(currentValue.given)
									}	
									if (currentValue.family) {	
										humanNameParts.push(currentValue.family)	
									}	
									if (currentValue.suffix) {	
										humanNameParts = humanNameParts.concat(currentValue.suffix)	
									}
									accumulator.push(humanNameParts.join(" "))
								} else if (currentValue.city || currentValue.state || currentValue.country || currentValue.postalCode) {
									//Address http://hl7.org/fhir/R4/datatypes.html#Address
									var addressParts = []		
									if (currentValue.line) {
										addressParts = addressParts.concat(currentValue.line)
									}
									if (currentValue.city) {
										addressParts.push(currentValue.city)
									}	
									if (currentValue.state) {	
										addressParts.push(currentValue.state)
									}	
									if (currentValue.postalCode) {
										addressParts.push(currentValue.postalCode)
									}	
									if (currentValue.country) {
										addressParts.push(currentValue.country)	
									}	
									accumulator.push(addressParts.join(" "))
								} else {
									//string, boolean
									accumulator.push(currentValue)
								}
								return accumulator
							}, [])
						
							if(SeriesProcessed.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(SeriesProcessed)
							}
						 `)
	if err == nil && seriesResult.String() != "undefined" {
		s.Series = []byte(seriesResult.String())
	}
	// extracting Status
	statusResult, err := vm.RunString(` 
							StatusResult = window.fhirpath.evaluate(fhirResource, 'Immunization.status')
							StatusProcessed = StatusResult.reduce((accumulator, currentValue) => {
								if (currentValue.coding) {
									//CodeableConcept
									currentValue.coding.map((coding) => {
										accumulator.push({
											"code": coding.code,	
											"system": coding.system,
											"text": currentValue.text
										})
									})
								} else if (currentValue.value) {
									//ContactPoint, Identifier
									accumulator.push({
										"code": currentValue.value,
										"system": currentValue.system,
										"text": currentValue.type?.text
									})
								} else if (currentValue.code) {
									//Coding
									accumulator.push({
										"code": currentValue.code,
										"system": currentValue.system,
										"text": currentValue.display
									})
								} else if ((typeof currentValue === 'string') || (typeof currentValue === 'boolean')) {
									//string, boolean
									accumulator.push({
										"code": currentValue,
									})
								}
								return accumulator
							}, [])
						
				
							if(StatusProcessed.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(StatusProcessed)
							}
						 `)
	if err == nil && statusResult.String() != "undefined" {
		s.Status = []byte(statusResult.String())
	}
	// extracting StatusReason
	statusReasonResult, err := vm.RunString(` 
							StatusReasonResult = window.fhirpath.evaluate(fhirResource, 'Immunization.statusReason')
							StatusReasonProcessed = StatusReasonResult.reduce((accumulator, currentValue) => {
								if (currentValue.coding) {
									//CodeableConcept
									currentValue.coding.map((coding) => {
										accumulator.push({
											"code": coding.code,	
											"system": coding.system,
											"text": currentValue.text
										})
									})
								} else if (currentValue.value) {
									//ContactPoint, Identifier
									accumulator.push({
										"code": currentValue.value,
										"system": currentValue.system,
										"text": currentValue.type?.text
									})
								} else if (currentValue.code) {
									//Coding
									accumulator.push({
										"code": currentValue.code,
										"system": currentValue.system,
										"text": currentValue.display
									})
								} else if ((typeof currentValue === 'string') || (typeof currentValue === 'boolean')) {
									//string, boolean
									accumulator.push({
										"code": currentValue,
									})
								}
								return accumulator
							}, [])
						
				
							if(StatusReasonProcessed.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(StatusReasonProcessed)
							}
						 `)
	if err == nil && statusReasonResult.String() != "undefined" {
		s.StatusReason = []byte(statusReasonResult.String())
	}
	// extracting Tag
	tagResult, err := vm.RunString(` 
							TagResult = window.fhirpath.evaluate(fhirResource, 'meta.tag')
							TagProcessed = TagResult.reduce((accumulator, currentValue) => {
								if (currentValue.coding) {
									//CodeableConcept
									currentValue.coding.map((coding) => {
										accumulator.push({
											"code": coding.code,	
											"system": coding.system,
											"text": currentValue.text
										})
									})
								} else if (currentValue.value) {
									//ContactPoint, Identifier
									accumulator.push({
										"code": currentValue.value,
										"system": currentValue.system,
										"text": currentValue.type?.text
									})
								} else if (currentValue.code) {
									//Coding
									accumulator.push({
										"code": currentValue.code,
										"system": currentValue.system,
										"text": currentValue.display
									})
								} else if ((typeof currentValue === 'string') || (typeof currentValue === 'boolean')) {
									//string, boolean
									accumulator.push({
										"code": currentValue,
									})
								}
								return accumulator
							}, [])
						
				
							if(TagProcessed.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(TagProcessed)
							}
						 `)
	if err == nil && tagResult.String() != "undefined" {
		s.Tag = []byte(tagResult.String())
	}
	// extracting TargetDisease
	targetDiseaseResult, err := vm.RunString(` 
							TargetDiseaseResult = window.fhirpath.evaluate(fhirResource, 'Immunization.protocolApplied.targetDisease')
							TargetDiseaseProcessed = TargetDiseaseResult.reduce((accumulator, currentValue) => {
								if (currentValue.coding) {
									//CodeableConcept
									currentValue.coding.map((coding) => {
										accumulator.push({
											"code": coding.code,	
											"system": coding.system,
											"text": currentValue.text
										})
									})
								} else if (currentValue.value) {
									//ContactPoint, Identifier
									accumulator.push({
										"code": currentValue.value,
										"system": currentValue.system,
										"text": currentValue.type?.text
									})
								} else if (currentValue.code) {
									//Coding
									accumulator.push({
										"code": currentValue.code,
										"system": currentValue.system,
										"text": currentValue.display
									})
								} else if ((typeof currentValue === 'string') || (typeof currentValue === 'boolean')) {
									//string, boolean
									accumulator.push({
										"code": currentValue,
									})
								}
								return accumulator
							}, [])
						
				
							if(TargetDiseaseProcessed.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(TargetDiseaseProcessed)
							}
						 `)
	if err == nil && targetDiseaseResult.String() != "undefined" {
		s.TargetDisease = []byte(targetDiseaseResult.String())
	}
	// extracting VaccineCode
	vaccineCodeResult, err := vm.RunString(` 
							VaccineCodeResult = window.fhirpath.evaluate(fhirResource, 'Immunization.vaccineCode')
							VaccineCodeProcessed = VaccineCodeResult.reduce((accumulator, currentValue) => {
								if (currentValue.coding) {
									//CodeableConcept
									currentValue.coding.map((coding) => {
										accumulator.push({
											"code": coding.code,	
											"system": coding.system,
											"text": currentValue.text
										})
									})
								} else if (currentValue.value) {
									//ContactPoint, Identifier
									accumulator.push({
										"code": currentValue.value,
										"system": currentValue.system,
										"text": currentValue.type?.text
									})
								} else if (currentValue.code) {
									//Coding
									accumulator.push({
										"code": currentValue.code,
										"system": currentValue.system,
										"text": currentValue.display
									})
								} else if ((typeof currentValue === 'string') || (typeof currentValue === 'boolean')) {
									//string, boolean
									accumulator.push({
										"code": currentValue,
									})
								}
								return accumulator
							}, [])
						
				
							if(VaccineCodeProcessed.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(VaccineCodeProcessed)
							}
						 `)
	if err == nil && vaccineCodeResult.String() != "undefined" {
		s.VaccineCode = []byte(vaccineCodeResult.String())
	}
	return nil
}

// TableName overrides the table name from fhir_observations (pluralized) to `fhir_observation`. https://gorm.io/docs/conventions.html#TableName
func (s *FhirImmunization) TableName() string {
	return "fhir_immunization"
}
