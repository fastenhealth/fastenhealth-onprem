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

type FhirPatient struct {
	models.OriginBase
	// Whether the patient record is active
	// https://hl7.org/fhir/r4/search.html#token
	Active datatypes.JSON `gorm:"column:active;type:text;serializer:json" json:"active,omitempty"`
	/*
	   Multiple Resources:

	   * [Patient](patient.html): A server defined search that may match any of the string fields in the Address, including line, city, district, state, country, postalCode, and/or text
	   * [Person](person.html): A server defined search that may match any of the string fields in the Address, including line, city, district, state, country, postalCode, and/or text
	   * [Practitioner](practitioner.html): A server defined search that may match any of the string fields in the Address, including line, city, district, state, country, postalCode, and/or text
	   * [RelatedPerson](relatedperson.html): A server defined search that may match any of the string fields in the Address, including line, city, district, state, country, postalCode, and/or text
	*/
	// https://hl7.org/fhir/r4/search.html#string
	Address string `gorm:"column:address;type:text" json:"address,omitempty"`
	/*
	   Multiple Resources:

	   * [Patient](patient.html): A city specified in an address
	   * [Person](person.html): A city specified in an address
	   * [Practitioner](practitioner.html): A city specified in an address
	   * [RelatedPerson](relatedperson.html): A city specified in an address
	*/
	// https://hl7.org/fhir/r4/search.html#string
	AddressCity string `gorm:"column:addressCity;type:text" json:"addressCity,omitempty"`
	/*
	   Multiple Resources:

	   * [Patient](patient.html): A country specified in an address
	   * [Person](person.html): A country specified in an address
	   * [Practitioner](practitioner.html): A country specified in an address
	   * [RelatedPerson](relatedperson.html): A country specified in an address
	*/
	// https://hl7.org/fhir/r4/search.html#string
	AddressCountry string `gorm:"column:addressCountry;type:text" json:"addressCountry,omitempty"`
	/*
	   Multiple Resources:

	   * [Patient](patient.html): A postalCode specified in an address
	   * [Person](person.html): A postal code specified in an address
	   * [Practitioner](practitioner.html): A postalCode specified in an address
	   * [RelatedPerson](relatedperson.html): A postal code specified in an address
	*/
	// https://hl7.org/fhir/r4/search.html#string
	AddressPostalcode string `gorm:"column:addressPostalcode;type:text" json:"addressPostalcode,omitempty"`
	/*
	   Multiple Resources:

	   * [Patient](patient.html): A state specified in an address
	   * [Person](person.html): A state specified in an address
	   * [Practitioner](practitioner.html): A state specified in an address
	   * [RelatedPerson](relatedperson.html): A state specified in an address
	*/
	// https://hl7.org/fhir/r4/search.html#string
	AddressState string `gorm:"column:addressState;type:text" json:"addressState,omitempty"`
	/*
	   Multiple Resources:

	   * [Patient](patient.html): A use code specified in an address
	   * [Person](person.html): A use code specified in an address
	   * [Practitioner](practitioner.html): A use code specified in an address
	   * [RelatedPerson](relatedperson.html): A use code specified in an address
	*/
	// https://hl7.org/fhir/r4/search.html#token
	AddressUse datatypes.JSON `gorm:"column:addressUse;type:text;serializer:json" json:"addressUse,omitempty"`
	/*
	   Multiple Resources:

	   * [Patient](patient.html): The patient's date of birth
	   * [Person](person.html): The person's date of birth
	   * [RelatedPerson](relatedperson.html): The Related Person's date of birth
	*/
	// https://hl7.org/fhir/r4/search.html#date
	Birthdate time.Time `gorm:"column:birthdate;type:datetime" json:"birthdate,omitempty"`
	// The date of death has been provided and satisfies this search value
	// https://hl7.org/fhir/r4/search.html#date
	DeathDate time.Time `gorm:"column:deathDate;type:datetime" json:"deathDate,omitempty"`
	// This patient has been marked as deceased, or has a death date entered
	// https://hl7.org/fhir/r4/search.html#token
	Deceased datatypes.JSON `gorm:"column:deceased;type:text;serializer:json" json:"deceased,omitempty"`
	/*
	   Multiple Resources:

	   * [Patient](patient.html): A value in an email contact
	   * [Person](person.html): A value in an email contact
	   * [Practitioner](practitioner.html): A value in an email contact
	   * [PractitionerRole](practitionerrole.html): A value in an email contact
	   * [RelatedPerson](relatedperson.html): A value in an email contact
	*/
	// https://hl7.org/fhir/r4/search.html#token
	Email datatypes.JSON `gorm:"column:email;type:text;serializer:json" json:"email,omitempty"`
	/*
	   Multiple Resources:

	   * [Patient](patient.html): A portion of the family name of the patient
	   * [Practitioner](practitioner.html): A portion of the family name
	*/
	// https://hl7.org/fhir/r4/search.html#string
	Family string `gorm:"column:family;type:text" json:"family,omitempty"`
	/*
	   Multiple Resources:

	   * [Patient](patient.html): Gender of the patient
	   * [Person](person.html): The gender of the person
	   * [Practitioner](practitioner.html): Gender of the practitioner
	   * [RelatedPerson](relatedperson.html): Gender of the related person
	*/
	// https://hl7.org/fhir/r4/search.html#token
	Gender datatypes.JSON `gorm:"column:gender;type:text;serializer:json" json:"gender,omitempty"`
	// Patient's nominated general practitioner, not the organization that manages the record
	// https://hl7.org/fhir/r4/search.html#reference
	GeneralPractitioner datatypes.JSON `gorm:"column:generalPractitioner;type:text;serializer:json" json:"generalPractitioner,omitempty"`
	/*
	   Multiple Resources:

	   * [Patient](patient.html): A portion of the given name of the patient
	   * [Practitioner](practitioner.html): A portion of the given name
	*/
	// https://hl7.org/fhir/r4/search.html#string
	Given string `gorm:"column:given;type:text" json:"given,omitempty"`
	// A patient identifier
	// https://hl7.org/fhir/r4/search.html#token
	Identifier datatypes.JSON `gorm:"column:identifier;type:text;serializer:json" json:"identifier,omitempty"`
	// Language of the resource content
	// https://hl7.org/fhir/r4/search.html#token
	Language datatypes.JSON `gorm:"column:language;type:text;serializer:json" json:"language,omitempty"`
	// When the resource version last changed
	// https://hl7.org/fhir/r4/search.html#date
	LastUpdated time.Time `gorm:"column:lastUpdated;type:datetime" json:"lastUpdated,omitempty"`
	// All patients linked to the given patient
	// https://hl7.org/fhir/r4/search.html#reference
	Link datatypes.JSON `gorm:"column:link;type:text;serializer:json" json:"link,omitempty"`
	// A server defined search that may match any of the string fields in the HumanName, including family, give, prefix, suffix, suffix, and/or text
	// https://hl7.org/fhir/r4/search.html#string
	Name string `gorm:"column:name;type:text" json:"name,omitempty"`
	// The organization that is the custodian of the patient record
	// https://hl7.org/fhir/r4/search.html#reference
	Organization datatypes.JSON `gorm:"column:organization;type:text;serializer:json" json:"organization,omitempty"`
	/*
	   Multiple Resources:

	   * [Patient](patient.html): A value in a phone contact
	   * [Person](person.html): A value in a phone contact
	   * [Practitioner](practitioner.html): A value in a phone contact
	   * [PractitionerRole](practitionerrole.html): A value in a phone contact
	   * [RelatedPerson](relatedperson.html): A value in a phone contact
	*/
	// https://hl7.org/fhir/r4/search.html#token
	Phone datatypes.JSON `gorm:"column:phone;type:text;serializer:json" json:"phone,omitempty"`
	/*
	   Multiple Resources:

	   * [Patient](patient.html): A portion of either family or given name using some kind of phonetic matching algorithm
	   * [Person](person.html): A portion of name using some kind of phonetic matching algorithm
	   * [Practitioner](practitioner.html): A portion of either family or given name using some kind of phonetic matching algorithm
	   * [RelatedPerson](relatedperson.html): A portion of name using some kind of phonetic matching algorithm
	*/
	// https://hl7.org/fhir/r4/search.html#string
	Phonetic string `gorm:"column:phonetic;type:text" json:"phonetic,omitempty"`
	// Profiles this resource claims to conform to
	// https://hl7.org/fhir/r4/search.html#reference
	Profile datatypes.JSON `gorm:"column:profile;type:text;serializer:json" json:"profile,omitempty"`
	// The raw resource content in JSON format
	// https://hl7.org/fhir/r4/search.html#special
	RawResource datatypes.JSON `gorm:"column:rawResource;type:text;serializer:json" json:"rawResource,omitempty"`
	// Identifies where the resource comes from
	// https://hl7.org/fhir/r4/search.html#uri
	SourceUri string `gorm:"column:sourceUri;type:text" json:"sourceUri,omitempty"`
	// Tags applied to this resource
	// https://hl7.org/fhir/r4/search.html#token
	Tag datatypes.JSON `gorm:"column:tag;type:text;serializer:json" json:"tag,omitempty"`
	/*
	   Multiple Resources:

	   * [Patient](patient.html): The value in any kind of telecom details of the patient
	   * [Person](person.html): The value in any kind of contact
	   * [Practitioner](practitioner.html): The value in any kind of contact
	   * [PractitionerRole](practitionerrole.html): The value in any kind of contact
	   * [RelatedPerson](relatedperson.html): The value in any kind of contact
	*/
	// https://hl7.org/fhir/r4/search.html#token
	Telecom datatypes.JSON `gorm:"column:telecom;type:text;serializer:json" json:"telecom,omitempty"`
	// Text search against the narrative
	// https://hl7.org/fhir/r4/search.html#string
	Text string `gorm:"column:text;type:text" json:"text,omitempty"`
	// A resource type filter
	// https://hl7.org/fhir/r4/search.html#special
	Type datatypes.JSON `gorm:"column:type;type:text;serializer:json" json:"type,omitempty"`
}

func (s *FhirPatient) SetOriginBase(originBase models.OriginBase) {
	s.OriginBase = originBase
}
func (s *FhirPatient) GetSearchParameters() map[string]string {
	searchParameters := map[string]string{
		"active":              "token",
		"address":             "string",
		"addressCity":         "string",
		"addressCountry":      "string",
		"addressPostalcode":   "string",
		"addressState":        "string",
		"addressUse":          "token",
		"birthdate":           "date",
		"deathDate":           "date",
		"deceased":            "token",
		"email":               "token",
		"family":              "string",
		"gender":              "token",
		"generalPractitioner": "reference",
		"given":               "string",
		"identifier":          "token",
		"language":            "token",
		"lastUpdated":         "date",
		"link":                "reference",
		"name":                "string",
		"organization":        "reference",
		"phone":               "token",
		"phonetic":            "string",
		"profile":             "reference",
		"rawResource":         "special",
		"sourceUri":           "uri",
		"tag":                 "token",
		"telecom":             "token",
		"text":                "string",
		"type":                "special",
	}
	return searchParameters
}
func (s *FhirPatient) PopulateAndExtractSearchParameters(rawResource json.RawMessage) error {
	s.RawResource = datatypes.JSON(rawResource)
	// unmarshal the raw resource (bytes) into a map
	var rawResourceMap map[string]interface{}
	err := json.Unmarshal(rawResource, &rawResourceMap)
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
	vm.Set("fhirResource", rawResourceMap)
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
	// extracting Active
	activeResult, err := vm.RunString("JSON.stringify(window.fhirpath.evaluate(fhirResource, 'Patient.active'))")
	if err == nil && activeResult.String() != "undefined" {
		s.Active = []byte(activeResult.String())
	}
	// extracting Address
	addressResult, err := vm.RunString("window.fhirpath.evaluate(fhirResource, 'Patient.address | Person.address | Practitioner.address | RelatedPerson.address')[0]")
	if err == nil && addressResult.String() != "undefined" {
		s.Address = addressResult.String()
	}
	// extracting AddressCity
	addressCityResult, err := vm.RunString("window.fhirpath.evaluate(fhirResource, 'Patient.address.city | Person.address.city | Practitioner.address.city | RelatedPerson.address.city')[0]")
	if err == nil && addressCityResult.String() != "undefined" {
		s.AddressCity = addressCityResult.String()
	}
	// extracting AddressCountry
	addressCountryResult, err := vm.RunString("window.fhirpath.evaluate(fhirResource, 'Patient.address.country | Person.address.country | Practitioner.address.country | RelatedPerson.address.country')[0]")
	if err == nil && addressCountryResult.String() != "undefined" {
		s.AddressCountry = addressCountryResult.String()
	}
	// extracting AddressPostalcode
	addressPostalcodeResult, err := vm.RunString("window.fhirpath.evaluate(fhirResource, 'Patient.address.postalCode | Person.address.postalCode | Practitioner.address.postalCode | RelatedPerson.address.postalCode')[0]")
	if err == nil && addressPostalcodeResult.String() != "undefined" {
		s.AddressPostalcode = addressPostalcodeResult.String()
	}
	// extracting AddressState
	addressStateResult, err := vm.RunString("window.fhirpath.evaluate(fhirResource, 'Patient.address.state | Person.address.state | Practitioner.address.state | RelatedPerson.address.state')[0]")
	if err == nil && addressStateResult.String() != "undefined" {
		s.AddressState = addressStateResult.String()
	}
	// extracting AddressUse
	addressUseResult, err := vm.RunString("JSON.stringify(window.fhirpath.evaluate(fhirResource, 'Patient.address.use | Person.address.use | Practitioner.address.use | RelatedPerson.address.use'))")
	if err == nil && addressUseResult.String() != "undefined" {
		s.AddressUse = []byte(addressUseResult.String())
	}
	// extracting Birthdate
	birthdateResult, err := vm.RunString("window.fhirpath.evaluate(fhirResource, 'Patient.birthDate | Person.birthDate | RelatedPerson.birthDate')[0]")
	if err == nil && birthdateResult.String() != "undefined" {
		t, err := time.Parse(time.RFC3339, birthdateResult.String())
		if err == nil {
			s.Birthdate = t
		}
	}
	// extracting DeathDate
	deathDateResult, err := vm.RunString("window.fhirpath.evaluate(fhirResource, '(Patient.deceasedDateTime)')[0]")
	if err == nil && deathDateResult.String() != "undefined" {
		t, err := time.Parse(time.RFC3339, deathDateResult.String())
		if err == nil {
			s.DeathDate = t
		}
	}
	// extracting Deceased
	deceasedResult, err := vm.RunString("JSON.stringify(window.fhirpath.evaluate(fhirResource, 'Patient.deceased.exists() and Patient.deceased != false'))")
	if err == nil && deceasedResult.String() != "undefined" {
		s.Deceased = []byte(deceasedResult.String())
	}
	// extracting Email
	emailResult, err := vm.RunString("JSON.stringify(window.fhirpath.evaluate(fhirResource, 'Patient.telecom.where(system='email') | Person.telecom.where(system='email') | Practitioner.telecom.where(system='email') | PractitionerRole.telecom.where(system='email') | RelatedPerson.telecom.where(system='email')'))")
	if err == nil && emailResult.String() != "undefined" {
		s.Email = []byte(emailResult.String())
	}
	// extracting Family
	familyResult, err := vm.RunString("window.fhirpath.evaluate(fhirResource, 'Patient.name.family | Practitioner.name.family')[0]")
	if err == nil && familyResult.String() != "undefined" {
		s.Family = familyResult.String()
	}
	// extracting Gender
	genderResult, err := vm.RunString("JSON.stringify(window.fhirpath.evaluate(fhirResource, 'Patient.gender | Person.gender | Practitioner.gender | RelatedPerson.gender'))")
	if err == nil && genderResult.String() != "undefined" {
		s.Gender = []byte(genderResult.String())
	}
	// extracting GeneralPractitioner
	generalPractitionerResult, err := vm.RunString("JSON.stringify(window.fhirpath.evaluate(fhirResource, 'Patient.generalPractitioner'))")
	if err == nil && generalPractitionerResult.String() != "undefined" {
		s.GeneralPractitioner = []byte(generalPractitionerResult.String())
	}
	// extracting Given
	givenResult, err := vm.RunString("window.fhirpath.evaluate(fhirResource, 'Patient.name.given | Practitioner.name.given')[0]")
	if err == nil && givenResult.String() != "undefined" {
		s.Given = givenResult.String()
	}
	// extracting Identifier
	identifierResult, err := vm.RunString("JSON.stringify(window.fhirpath.evaluate(fhirResource, 'Patient.identifier'))")
	if err == nil && identifierResult.String() != "undefined" {
		s.Identifier = []byte(identifierResult.String())
	}
	// extracting Language
	languageResult, err := vm.RunString("JSON.stringify(window.fhirpath.evaluate(fhirResource, 'Resource.language'))")
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
	// extracting Link
	linkResult, err := vm.RunString("JSON.stringify(window.fhirpath.evaluate(fhirResource, 'Patient.link.other'))")
	if err == nil && linkResult.String() != "undefined" {
		s.Link = []byte(linkResult.String())
	}
	// extracting Name
	nameResult, err := vm.RunString("window.fhirpath.evaluate(fhirResource, 'Patient.name')[0]")
	if err == nil && nameResult.String() != "undefined" {
		s.Name = nameResult.String()
	}
	// extracting Organization
	organizationResult, err := vm.RunString("JSON.stringify(window.fhirpath.evaluate(fhirResource, 'Patient.managingOrganization'))")
	if err == nil && organizationResult.String() != "undefined" {
		s.Organization = []byte(organizationResult.String())
	}
	// extracting Phone
	phoneResult, err := vm.RunString("JSON.stringify(window.fhirpath.evaluate(fhirResource, 'Patient.telecom.where(system='phone') | Person.telecom.where(system='phone') | Practitioner.telecom.where(system='phone') | PractitionerRole.telecom.where(system='phone') | RelatedPerson.telecom.where(system='phone')'))")
	if err == nil && phoneResult.String() != "undefined" {
		s.Phone = []byte(phoneResult.String())
	}
	// extracting Phonetic
	phoneticResult, err := vm.RunString("window.fhirpath.evaluate(fhirResource, 'Patient.name | Person.name | Practitioner.name | RelatedPerson.name')[0]")
	if err == nil && phoneticResult.String() != "undefined" {
		s.Phonetic = phoneticResult.String()
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
	// extracting Tag
	tagResult, err := vm.RunString("JSON.stringify(window.fhirpath.evaluate(fhirResource, 'Resource.meta.tag'))")
	if err == nil && tagResult.String() != "undefined" {
		s.Tag = []byte(tagResult.String())
	}
	// extracting Telecom
	telecomResult, err := vm.RunString("JSON.stringify(window.fhirpath.evaluate(fhirResource, 'Patient.telecom | Person.telecom | Practitioner.telecom | PractitionerRole.telecom | RelatedPerson.telecom'))")
	if err == nil && telecomResult.String() != "undefined" {
		s.Telecom = []byte(telecomResult.String())
	}
	return nil
}

// TableName overrides the table name from fhir_observations (pluralized) to `fhir_observation`. https://gorm.io/docs/conventions.html#TableName
func (s *FhirPatient) TableName() string {
	return "fhir_patient"
}
