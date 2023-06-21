// THIS FILE IS GENERATED BY https://github.com/fastenhealth/fasten-onprem/blob/main/backend/pkg/models/database/generate.go
// PLEASE DO NOT EDIT BY HAND

package database

import (
	"encoding/json"
	"fmt"
	goja "github.com/dop251/goja"
	models "github.com/fastenhealth/fastenhealth-onprem/backend/pkg/models"
	datatypes "gorm.io/datatypes"
	"time"
)

type FhirCoverage struct {
	models.ResourceBase
	// Covered party
	// https://hl7.org/fhir/r4/search.html#reference
	Beneficiary datatypes.JSON `gorm:"column:beneficiary;type:text;serializer:json" json:"beneficiary,omitempty"`
	// Coverage class (eg. plan, group)
	// https://hl7.org/fhir/r4/search.html#token
	ClassType datatypes.JSON `gorm:"column:classType;type:text;serializer:json" json:"classType,omitempty"`
	// Value of the class (eg. Plan number, group number)
	// https://hl7.org/fhir/r4/search.html#string
	ClassValue string `gorm:"column:classValue;type:text" json:"classValue,omitempty"`
	// Dependent number
	// https://hl7.org/fhir/r4/search.html#string
	Dependent string `gorm:"column:dependent;type:text" json:"dependent,omitempty"`
	// The primary identifier of the insured and the coverage
	// https://hl7.org/fhir/r4/search.html#token
	Identifier datatypes.JSON `gorm:"column:identifier;type:text;serializer:json" json:"identifier,omitempty"`
	// Language of the resource content
	// https://hl7.org/fhir/r4/search.html#token
	Language datatypes.JSON `gorm:"column:language;type:text;serializer:json" json:"language,omitempty"`
	// When the resource version last changed
	// https://hl7.org/fhir/r4/search.html#date
	LastUpdated time.Time `gorm:"column:lastUpdated;type:datetime" json:"lastUpdated,omitempty"`
	// The identity of the insurer or party paying for services
	// https://hl7.org/fhir/r4/search.html#reference
	Payor datatypes.JSON `gorm:"column:payor;type:text;serializer:json" json:"payor,omitempty"`
	// Reference to the policyholder
	// https://hl7.org/fhir/r4/search.html#reference
	PolicyHolder datatypes.JSON `gorm:"column:policyHolder;type:text;serializer:json" json:"policyHolder,omitempty"`
	// Profiles this resource claims to conform to
	// https://hl7.org/fhir/r4/search.html#reference
	Profile datatypes.JSON `gorm:"column:profile;type:text;serializer:json" json:"profile,omitempty"`
	// Identifies where the resource comes from
	// https://hl7.org/fhir/r4/search.html#uri
	SourceUri string `gorm:"column:sourceUri;type:text" json:"sourceUri,omitempty"`
	// The status of the Coverage
	// https://hl7.org/fhir/r4/search.html#token
	Status datatypes.JSON `gorm:"column:status;type:text;serializer:json" json:"status,omitempty"`
	// Reference to the subscriber
	// https://hl7.org/fhir/r4/search.html#reference
	Subscriber datatypes.JSON `gorm:"column:subscriber;type:text;serializer:json" json:"subscriber,omitempty"`
	// Tags applied to this resource
	// https://hl7.org/fhir/r4/search.html#token
	Tag datatypes.JSON `gorm:"column:tag;type:text;serializer:json" json:"tag,omitempty"`
	// Text search against the narrative
	// https://hl7.org/fhir/r4/search.html#string
	Text string `gorm:"column:text;type:text" json:"text,omitempty"`
	// A resource type filter
	// https://hl7.org/fhir/r4/search.html#special
	Type datatypes.JSON `gorm:"column:type;type:text;serializer:json" json:"type,omitempty"`
}

func (s *FhirCoverage) GetSearchParameters() map[string]string {
	searchParameters := map[string]string{
		"beneficiary":  "reference",
		"classType":    "token",
		"classValue":   "string",
		"dependent":    "string",
		"identifier":   "token",
		"language":     "token",
		"lastUpdated":  "date",
		"payor":        "reference",
		"policyHolder": "reference",
		"profile":      "reference",
		"sourceUri":    "uri",
		"status":       "token",
		"subscriber":   "reference",
		"tag":          "token",
		"text":         "string",
		"type":         "special",
	}
	return searchParameters
}
func (s *FhirCoverage) PopulateAndExtractSearchParameters(resourceRaw json.RawMessage) error {
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
	// extracting Beneficiary
	beneficiaryResult, err := vm.RunString("JSON.stringify(window.fhirpath.evaluate(fhirResource, 'Coverage.beneficiary'))")
	if err == nil && beneficiaryResult.String() != "undefined" {
		s.Beneficiary = []byte(beneficiaryResult.String())
	}
	// extracting ClassType
	classTypeResult, err := vm.RunString(` 
							ClassTypeResult = window.fhirpath.evaluate(fhirResource, 'Coverage.class.type')
							ClassTypeProcessed = ClassTypeResult.reduce((accumulator, currentValue) => {
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
						
				
							JSON.stringify(ClassTypeProcessed)
						 `)
	if err == nil && classTypeResult.String() != "undefined" {
		s.ClassType = []byte(classTypeResult.String())
	}
	// extracting ClassValue
	classValueResult, err := vm.RunString("window.fhirpath.evaluate(fhirResource, 'Coverage.class.value')[0]")
	if err == nil && classValueResult.String() != "undefined" {
		s.ClassValue = classValueResult.String()
	}
	// extracting Dependent
	dependentResult, err := vm.RunString("window.fhirpath.evaluate(fhirResource, 'Coverage.dependent')[0]")
	if err == nil && dependentResult.String() != "undefined" {
		s.Dependent = dependentResult.String()
	}
	// extracting Identifier
	identifierResult, err := vm.RunString(` 
							IdentifierResult = window.fhirpath.evaluate(fhirResource, 'Coverage.identifier')
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
						
				
							JSON.stringify(IdentifierProcessed)
						 `)
	if err == nil && identifierResult.String() != "undefined" {
		s.Identifier = []byte(identifierResult.String())
	}
	// extracting Language
	languageResult, err := vm.RunString(` 
							LanguageResult = window.fhirpath.evaluate(fhirResource, 'Resource.language')
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
						
				
							JSON.stringify(LanguageProcessed)
						 `)
	if err == nil && languageResult.String() != "undefined" {
		s.Language = []byte(languageResult.String())
	}
	// extracting LastUpdated
	lastUpdatedResult, err := vm.RunString("window.fhirpath.evaluate(fhirResource, 'Resource.meta.lastUpdated')[0]")
	if err == nil && lastUpdatedResult.String() != "undefined" {
		t, err := time.Parse(time.RFC3339, lastUpdatedResult.String())
		if err == nil {
			s.LastUpdated = t
		}
	}
	// extracting Payor
	payorResult, err := vm.RunString("JSON.stringify(window.fhirpath.evaluate(fhirResource, 'Coverage.payor'))")
	if err == nil && payorResult.String() != "undefined" {
		s.Payor = []byte(payorResult.String())
	}
	// extracting PolicyHolder
	policyHolderResult, err := vm.RunString("JSON.stringify(window.fhirpath.evaluate(fhirResource, 'Coverage.policyHolder'))")
	if err == nil && policyHolderResult.String() != "undefined" {
		s.PolicyHolder = []byte(policyHolderResult.String())
	}
	// extracting Profile
	profileResult, err := vm.RunString("JSON.stringify(window.fhirpath.evaluate(fhirResource, 'Resource.meta.profile'))")
	if err == nil && profileResult.String() != "undefined" {
		s.Profile = []byte(profileResult.String())
	}
	// extracting SourceUri
	sourceUriResult, err := vm.RunString("window.fhirpath.evaluate(fhirResource, 'Resource.meta.source')[0]")
	if err == nil && sourceUriResult.String() != "undefined" {
		s.SourceUri = sourceUriResult.String()
	}
	// extracting Status
	statusResult, err := vm.RunString(` 
							StatusResult = window.fhirpath.evaluate(fhirResource, 'Coverage.status')
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
						
				
							JSON.stringify(StatusProcessed)
						 `)
	if err == nil && statusResult.String() != "undefined" {
		s.Status = []byte(statusResult.String())
	}
	// extracting Subscriber
	subscriberResult, err := vm.RunString("JSON.stringify(window.fhirpath.evaluate(fhirResource, 'Coverage.subscriber'))")
	if err == nil && subscriberResult.String() != "undefined" {
		s.Subscriber = []byte(subscriberResult.String())
	}
	// extracting Tag
	tagResult, err := vm.RunString(` 
							TagResult = window.fhirpath.evaluate(fhirResource, 'Resource.meta.tag')
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
						
				
							JSON.stringify(TagProcessed)
						 `)
	if err == nil && tagResult.String() != "undefined" {
		s.Tag = []byte(tagResult.String())
	}
	return nil
}

// TableName overrides the table name from fhir_observations (pluralized) to `fhir_observation`. https://gorm.io/docs/conventions.html#TableName
func (s *FhirCoverage) TableName() string {
	return "fhir_coverage"
}
