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

type FhirDocumentReference struct {
	models.ResourceBase
	// Who/what authenticated the document
	// https://hl7.org/fhir/r4/search.html#reference
	Authenticator datatypes.JSON `gorm:"column:authenticator;type:text;serializer:json" json:"authenticator,omitempty"`
	// Who and/or what authored the document
	// https://hl7.org/fhir/r4/search.html#reference
	Author datatypes.JSON `gorm:"column:author;type:text;serializer:json" json:"author,omitempty"`
	// Categorization of document
	// https://hl7.org/fhir/r4/search.html#token
	Category datatypes.JSON `gorm:"column:category;type:text;serializer:json" json:"category,omitempty"`
	// Mime type of the content, with charset etc.
	// https://hl7.org/fhir/r4/search.html#token
	Contenttype datatypes.JSON `gorm:"column:contenttype;type:text;serializer:json" json:"contenttype,omitempty"`
	// Organization which maintains the document
	// https://hl7.org/fhir/r4/search.html#reference
	Custodian datatypes.JSON `gorm:"column:custodian;type:text;serializer:json" json:"custodian,omitempty"`
	// When this document reference was created
	// https://hl7.org/fhir/r4/search.html#date
	Date *time.Time `gorm:"column:date;type:datetime" json:"date,omitempty"`
	// Human-readable description
	// https://hl7.org/fhir/r4/search.html#string
	Description datatypes.JSON `gorm:"column:description;type:text;serializer:json" json:"description,omitempty"`
	/*
	   Multiple Resources:

	   * [Composition](composition.html): Context of the Composition
	   * [DeviceRequest](devicerequest.html): Encounter during which request was created
	   * [DiagnosticReport](diagnosticreport.html): The Encounter when the order was made
	   * [DocumentReference](documentreference.html): Context of the document  content
	   * [Flag](flag.html): Alert relevant during encounter
	   * [List](list.html): Context in which list created
	   * [NutritionOrder](nutritionorder.html): Return nutrition orders with this encounter identifier
	   * [Observation](observation.html): Encounter related to the observation
	   * [Procedure](procedure.html): Encounter created as part of
	   * [RiskAssessment](riskassessment.html): Where was assessment performed?
	   * [ServiceRequest](servicerequest.html): An encounter in which this request is made
	   * [VisionPrescription](visionprescription.html): Return prescriptions with this encounter identifier
	*/
	// https://hl7.org/fhir/r4/search.html#reference
	Encounter datatypes.JSON `gorm:"column:encounter;type:text;serializer:json" json:"encounter,omitempty"`
	// Main clinical acts documented
	// https://hl7.org/fhir/r4/search.html#token
	Event datatypes.JSON `gorm:"column:event;type:text;serializer:json" json:"event,omitempty"`
	// Kind of facility where patient was seen
	// https://hl7.org/fhir/r4/search.html#token
	Facility datatypes.JSON `gorm:"column:facility;type:text;serializer:json" json:"facility,omitempty"`
	// Format/content rules for the document
	// https://hl7.org/fhir/r4/search.html#token
	Format datatypes.JSON `gorm:"column:format;type:text;serializer:json" json:"format,omitempty"`
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
	// Uri where the data can be found
	// https://hl7.org/fhir/r4/search.html#uri
	Location string `gorm:"column:location;type:text" json:"location,omitempty"`
	// Time of service that is being documented
	// https://hl7.org/fhir/r4/search.html#date
	Period *time.Time `gorm:"column:period;type:datetime" json:"period,omitempty"`
	// Profiles this resource claims to conform to
	// https://hl7.org/fhir/r4/search.html#reference
	Profile datatypes.JSON `gorm:"column:profile;type:text;serializer:json" json:"profile,omitempty"`
	// Related identifiers or resources
	// https://hl7.org/fhir/r4/search.html#reference
	Related datatypes.JSON `gorm:"column:related;type:text;serializer:json" json:"related,omitempty"`
	// Target of the relationship
	// https://hl7.org/fhir/r4/search.html#reference
	Relatesto datatypes.JSON `gorm:"column:relatesto;type:text;serializer:json" json:"relatesto,omitempty"`
	// replaces | transforms | signs | appends
	// https://hl7.org/fhir/r4/search.html#token
	Relation datatypes.JSON `gorm:"column:relation;type:text;serializer:json" json:"relation,omitempty"`
	// Document security-tags
	// https://hl7.org/fhir/r4/search.html#token
	SecurityLabel datatypes.JSON `gorm:"column:securityLabel;type:text;serializer:json" json:"securityLabel,omitempty"`
	// Additional details about where the content was created (e.g. clinical specialty)
	// https://hl7.org/fhir/r4/search.html#token
	Setting datatypes.JSON `gorm:"column:setting;type:text;serializer:json" json:"setting,omitempty"`
	// current | superseded | entered-in-error
	// https://hl7.org/fhir/r4/search.html#token
	Status datatypes.JSON `gorm:"column:status;type:text;serializer:json" json:"status,omitempty"`
	// Who/what is the subject of the document
	// https://hl7.org/fhir/r4/search.html#reference
	Subject datatypes.JSON `gorm:"column:subject;type:text;serializer:json" json:"subject,omitempty"`
	// Tags applied to this resource
	// https://hl7.org/fhir/r4/search.html#token
	Tag datatypes.JSON `gorm:"column:tag;type:text;serializer:json" json:"tag,omitempty"`
	// Text search against the narrative
	// https://hl7.org/fhir/r4/search.html#string
	Text datatypes.JSON `gorm:"column:text;type:text;serializer:json" json:"text,omitempty"`
	// A resource type filter
	// https://hl7.org/fhir/r4/search.html#special
	Type datatypes.JSON `gorm:"column:type;type:text;serializer:json" json:"type,omitempty"`
}

func (s *FhirDocumentReference) GetSearchParameters() map[string]string {
	searchParameters := map[string]string{
		"authenticator":        "reference",
		"author":               "reference",
		"category":             "token",
		"contenttype":          "token",
		"custodian":            "reference",
		"date":                 "date",
		"description":          "string",
		"encounter":            "reference",
		"event":                "token",
		"facility":             "token",
		"format":               "token",
		"id":                   "keyword",
		"identifier":           "token",
		"language":             "token",
		"lastUpdated":          "date",
		"location":             "uri",
		"period":               "date",
		"profile":              "reference",
		"related":              "reference",
		"relatesto":            "reference",
		"relation":             "token",
		"securityLabel":        "token",
		"setting":              "token",
		"source_id":            "keyword",
		"source_resource_id":   "keyword",
		"source_resource_type": "keyword",
		"source_uri":           "keyword",
		"status":               "token",
		"subject":              "reference",
		"tag":                  "token",
		"text":                 "string",
		"type":                 "special",
	}
	return searchParameters
}
func (s *FhirDocumentReference) PopulateAndExtractSearchParameters(resourceRaw json.RawMessage) error {
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
	// extracting Authenticator
	authenticatorResult, err := vm.RunString(` 
							AuthenticatorResult = window.fhirpath.evaluate(fhirResource, 'DocumentReference.authenticator')
						
							if(AuthenticatorResult.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(AuthenticatorResult)
							}
						 `)
	if err == nil && authenticatorResult.String() != "undefined" {
		s.Authenticator = []byte(authenticatorResult.String())
	}
	// extracting Author
	authorResult, err := vm.RunString(` 
							AuthorResult = window.fhirpath.evaluate(fhirResource, 'DocumentReference.author')
						
							if(AuthorResult.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(AuthorResult)
							}
						 `)
	if err == nil && authorResult.String() != "undefined" {
		s.Author = []byte(authorResult.String())
	}
	// extracting Category
	categoryResult, err := vm.RunString(` 
							CategoryResult = window.fhirpath.evaluate(fhirResource, 'DocumentReference.category')
							CategoryProcessed = CategoryResult.reduce((accumulator, currentValue) => {
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
						
				
							if(CategoryProcessed.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(CategoryProcessed)
							}
						 `)
	if err == nil && categoryResult.String() != "undefined" {
		s.Category = []byte(categoryResult.String())
	}
	// extracting Contenttype
	contenttypeResult, err := vm.RunString(` 
							ContenttypeResult = window.fhirpath.evaluate(fhirResource, 'DocumentReference.content.attachment.contentType')
							ContenttypeProcessed = ContenttypeResult.reduce((accumulator, currentValue) => {
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
						
				
							if(ContenttypeProcessed.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(ContenttypeProcessed)
							}
						 `)
	if err == nil && contenttypeResult.String() != "undefined" {
		s.Contenttype = []byte(contenttypeResult.String())
	}
	// extracting Custodian
	custodianResult, err := vm.RunString(` 
							CustodianResult = window.fhirpath.evaluate(fhirResource, 'DocumentReference.custodian')
						
							if(CustodianResult.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(CustodianResult)
							}
						 `)
	if err == nil && custodianResult.String() != "undefined" {
		s.Custodian = []byte(custodianResult.String())
	}
	// extracting Date
	dateResult, err := vm.RunString("window.fhirpath.evaluate(fhirResource, 'DocumentReference.date')[0]")
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
	// extracting Description
	descriptionResult, err := vm.RunString(` 
							DescriptionResult = window.fhirpath.evaluate(fhirResource, 'DocumentReference.description')
							DescriptionProcessed = DescriptionResult.reduce((accumulator, currentValue) => {
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
						
							if(DescriptionProcessed.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(DescriptionProcessed)
							}
						 `)
	if err == nil && descriptionResult.String() != "undefined" {
		s.Description = []byte(descriptionResult.String())
	}
	// extracting Encounter
	encounterResult, err := vm.RunString(` 
							EncounterResult = window.fhirpath.evaluate(fhirResource, 'Composition.encounter | DeviceRequest.encounter | DiagnosticReport.encounter | DocumentReference.context.encounter.where(resolve() is Encounter) | Flag.encounter | List.encounter | NutritionOrder.encounter | Observation.encounter | Procedure.encounter | RiskAssessment.encounter | ServiceRequest.encounter | VisionPrescription.encounter')
						
							if(EncounterResult.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(EncounterResult)
							}
						 `)
	if err == nil && encounterResult.String() != "undefined" {
		s.Encounter = []byte(encounterResult.String())
	}
	// extracting Event
	eventResult, err := vm.RunString(` 
							EventResult = window.fhirpath.evaluate(fhirResource, 'DocumentReference.context.event')
							EventProcessed = EventResult.reduce((accumulator, currentValue) => {
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
						
				
							if(EventProcessed.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(EventProcessed)
							}
						 `)
	if err == nil && eventResult.String() != "undefined" {
		s.Event = []byte(eventResult.String())
	}
	// extracting Facility
	facilityResult, err := vm.RunString(` 
							FacilityResult = window.fhirpath.evaluate(fhirResource, 'DocumentReference.context.facilityType')
							FacilityProcessed = FacilityResult.reduce((accumulator, currentValue) => {
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
						
				
							if(FacilityProcessed.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(FacilityProcessed)
							}
						 `)
	if err == nil && facilityResult.String() != "undefined" {
		s.Facility = []byte(facilityResult.String())
	}
	// extracting Format
	formatResult, err := vm.RunString(` 
							FormatResult = window.fhirpath.evaluate(fhirResource, 'DocumentReference.content.format')
							FormatProcessed = FormatResult.reduce((accumulator, currentValue) => {
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
						
				
							if(FormatProcessed.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(FormatProcessed)
							}
						 `)
	if err == nil && formatResult.String() != "undefined" {
		s.Format = []byte(formatResult.String())
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
	locationResult, err := vm.RunString("window.fhirpath.evaluate(fhirResource, 'DocumentReference.content.attachment.url')[0]")
	if err == nil && locationResult.String() != "undefined" {
		s.Location = locationResult.String()
	}
	// extracting Period
	periodResult, err := vm.RunString("window.fhirpath.evaluate(fhirResource, 'DocumentReference.context.period')[0]")
	if err == nil && periodResult.String() != "undefined" {
		t, err := time.Parse(time.RFC3339, periodResult.String())
		if err == nil {
			s.Period = &t
		} else if err != nil {
			d, err := time.Parse("2006-01-02", periodResult.String())
			if err == nil {
				s.Period = &d
			}
		}
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
	// extracting Related
	relatedResult, err := vm.RunString(` 
							RelatedResult = window.fhirpath.evaluate(fhirResource, 'DocumentReference.context.related')
						
							if(RelatedResult.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(RelatedResult)
							}
						 `)
	if err == nil && relatedResult.String() != "undefined" {
		s.Related = []byte(relatedResult.String())
	}
	// extracting Relatesto
	relatestoResult, err := vm.RunString(` 
							RelatestoResult = window.fhirpath.evaluate(fhirResource, 'DocumentReference.relatesTo.target')
						
							if(RelatestoResult.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(RelatestoResult)
							}
						 `)
	if err == nil && relatestoResult.String() != "undefined" {
		s.Relatesto = []byte(relatestoResult.String())
	}
	// extracting Relation
	relationResult, err := vm.RunString(` 
							RelationResult = window.fhirpath.evaluate(fhirResource, 'DocumentReference.relatesTo.code')
							RelationProcessed = RelationResult.reduce((accumulator, currentValue) => {
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
						
				
							if(RelationProcessed.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(RelationProcessed)
							}
						 `)
	if err == nil && relationResult.String() != "undefined" {
		s.Relation = []byte(relationResult.String())
	}
	// extracting SecurityLabel
	securityLabelResult, err := vm.RunString(` 
							SecurityLabelResult = window.fhirpath.evaluate(fhirResource, 'DocumentReference.securityLabel')
							SecurityLabelProcessed = SecurityLabelResult.reduce((accumulator, currentValue) => {
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
						
				
							if(SecurityLabelProcessed.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(SecurityLabelProcessed)
							}
						 `)
	if err == nil && securityLabelResult.String() != "undefined" {
		s.SecurityLabel = []byte(securityLabelResult.String())
	}
	// extracting Setting
	settingResult, err := vm.RunString(` 
							SettingResult = window.fhirpath.evaluate(fhirResource, 'DocumentReference.context.practiceSetting')
							SettingProcessed = SettingResult.reduce((accumulator, currentValue) => {
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
						
				
							if(SettingProcessed.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(SettingProcessed)
							}
						 `)
	if err == nil && settingResult.String() != "undefined" {
		s.Setting = []byte(settingResult.String())
	}
	// extracting Status
	statusResult, err := vm.RunString(` 
							StatusResult = window.fhirpath.evaluate(fhirResource, 'DocumentReference.status')
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
	// extracting Subject
	subjectResult, err := vm.RunString(` 
							SubjectResult = window.fhirpath.evaluate(fhirResource, 'DocumentReference.subject')
						
							if(SubjectResult.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(SubjectResult)
							}
						 `)
	if err == nil && subjectResult.String() != "undefined" {
		s.Subject = []byte(subjectResult.String())
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
	return nil
}

// TableName overrides the table name from fhir_observations (pluralized) to `fhir_observation`. https://gorm.io/docs/conventions.html#TableName
func (s *FhirDocumentReference) TableName() string {
	return "fhir_document_reference"
}
