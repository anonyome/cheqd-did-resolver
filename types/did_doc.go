package types

import (
	"encoding/json"

	did "github.com/cheqd/cheqd-node/x/did/types"
)

type DidDoc struct {
	Context              []string             `json:"@context,omitempty" example:"https://www.w3.org/ns/did/v1"`
	Id                   string               `json:"id,omitempty" example:"did:cheqd:testnet:55dbc8bf-fba3-4117-855c-1e0dc1d3bb47"`
	Controller           []string             `json:"controller,omitempty" example:"did:cheqd:testnet:55dbc8bf-fba3-4117-855c-1e0dc1d3bb47"`
	VerificationMethod   []VerificationMethod `json:"verificationMethod,omitempty"`
	Authentication       []string             `json:"authentication,omitempty" example:"did:cheqd:testnet:55dbc8bf-fba3-4117-855c-1e0dc1d3bb47#key-1"`
	AssertionMethod      []string             `json:"assertionMethod,omitempty"`
	CapabilityInvocation []string             `json:"capabilityInvocation,omitempty"`
	CapabilityDelegation []string             `json:"capability_delegation,omitempty"`
	KeyAgreement         []string             `json:"keyAgreement,omitempty"`
	Service              []Service            `json:"service,omitempty"`
	AlsoKnownAs          []string             `json:"alsoKnownAs,omitempty"`
}

type VerificationMethod struct {
	Context            []string          `json:"@context,omitempty"`
	Id                 string            `json:"id,omitempty"`
	Type               string            `json:"type,omitempty"`
	Controller         string            `json:"controller,omitempty"`
	PublicKeyJwk       map[string]string `json:"publicKeyJwk,omitempty"`
	PublicKeyMultibase string            `json:"publicKeyMultibase,omitempty"`
}

type VerificationMaterial struct {
	PublicKeyJwk       map[string]string `json:"publicKeyJwk,omitempty"`
	PublicKeyMultibase string            `json:"publicKeyMultibase,omitempty"`
}

type Service struct {
	Context         []string `json:"@context,omitempty"`
	Id              string   `json:"id,omitempty" example:"did:cheqd:testnet:55dbc8bf-fba3-4117-855c-1e0dc1d3bb47#service-1"`
	Type            string   `json:"type,omitempty" example:"did-communication"`
	ServiceEndpoint []string `json:"serviceEndpoint,omitempty" example:"https://example.com/endpoint/8377464"`
}

func NewDidDoc(protoDidDoc did.DidDoc) DidDoc {
	verificationMethods := []VerificationMethod{}
	for _, vm := range protoDidDoc.VerificationMethod {
		verificationMethods = append(verificationMethods, *NewVerificationMethod(vm))
	}

	services := []Service{}
	for _, s := range protoDidDoc.Service {
		services = append(services, *NewService(s))
	}

	return DidDoc{
		Id:                   protoDidDoc.Id,
		Controller:           protoDidDoc.Controller,
		VerificationMethod:   verificationMethods,
		Authentication:       protoDidDoc.Authentication,
		AssertionMethod:      protoDidDoc.AssertionMethod,
		CapabilityInvocation: protoDidDoc.CapabilityInvocation,
		CapabilityDelegation: protoDidDoc.CapabilityDelegation,
		KeyAgreement:         protoDidDoc.KeyAgreement,
		Service:              services,
		AlsoKnownAs:          protoDidDoc.AlsoKnownAs,
	}
}

func NewVerificationMethod(protoVerificationMethod *did.VerificationMethod) *VerificationMethod {
	var verificationMaterial *VerificationMaterial
	err := json.Unmarshal([]byte(protoVerificationMethod.VerificationMaterial), &verificationMaterial)
	if err != nil {
		println("Invalid verification material !!!")
		panic(err)
	}

	return &VerificationMethod{
		Id:                 protoVerificationMethod.Id,
		Type:               protoVerificationMethod.Type,
		Controller:         protoVerificationMethod.Controller,
		PublicKeyJwk:       verificationMaterial.PublicKeyJwk,
		PublicKeyMultibase: verificationMaterial.PublicKeyMultibase,
	}
}

func NewService(protoService *did.Service) *Service {
	return &Service{
		Id:              protoService.Id,
		Type:            protoService.Type,
		ServiceEndpoint: protoService.ServiceEndpoint,
	}
}

func (e *DidDoc) AddContext(newProtocol string) { e.Context = AddElemToSet(e.Context, newProtocol) }
func (e *DidDoc) RemoveContext()                { e.Context = nil }
func (e *DidDoc) GetBytes() []byte              { return []byte{} }

func (e *Service) AddContext(newProtocol string) { e.Context = AddElemToSet(e.Context, newProtocol) }
func (e *Service) RemoveContext()                { e.Context = nil }
func (e *Service) GetBytes() []byte              { return []byte{} }

func (e *VerificationMethod) AddContext(newProtocol string) {
	e.Context = AddElemToSet(e.Context, newProtocol)
}
func (e *VerificationMethod) RemoveContext()   { e.Context = nil }
func (e *VerificationMethod) GetBytes() []byte { return []byte{} }
