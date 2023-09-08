// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// ExtrasCustomFieldChoiceSetsReadReader is a Reader for the ExtrasCustomFieldChoiceSetsRead structure.
type ExtrasCustomFieldChoiceSetsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasCustomFieldChoiceSetsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasCustomFieldChoiceSetsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasCustomFieldChoiceSetsReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasCustomFieldChoiceSetsReadOK creates a ExtrasCustomFieldChoiceSetsReadOK with default headers values
func NewExtrasCustomFieldChoiceSetsReadOK() *ExtrasCustomFieldChoiceSetsReadOK {
	return &ExtrasCustomFieldChoiceSetsReadOK{}
}

/*
ExtrasCustomFieldChoiceSetsReadOK describes a response with status code 200, with default header values.

ExtrasCustomFieldChoiceSetsReadOK extras custom field choice sets read o k
*/
type ExtrasCustomFieldChoiceSetsReadOK struct {
	Payload *models.CustomFieldChoiceSet
}

// IsSuccess returns true when this extras custom field choice sets read o k response has a 2xx status code
func (o *ExtrasCustomFieldChoiceSetsReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras custom field choice sets read o k response has a 3xx status code
func (o *ExtrasCustomFieldChoiceSetsReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras custom field choice sets read o k response has a 4xx status code
func (o *ExtrasCustomFieldChoiceSetsReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras custom field choice sets read o k response has a 5xx status code
func (o *ExtrasCustomFieldChoiceSetsReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras custom field choice sets read o k response a status code equal to that given
func (o *ExtrasCustomFieldChoiceSetsReadOK) IsCode(code int) bool {
	return code == 200
}

func (o *ExtrasCustomFieldChoiceSetsReadOK) Error() string {
	return fmt.Sprintf("[GET /extras/custom-field-choice-sets/{id}/][%d] extrasCustomFieldChoiceSetsReadOK  %+v", 200, o.Payload)
}

func (o *ExtrasCustomFieldChoiceSetsReadOK) String() string {
	return fmt.Sprintf("[GET /extras/custom-field-choice-sets/{id}/][%d] extrasCustomFieldChoiceSetsReadOK  %+v", 200, o.Payload)
}

func (o *ExtrasCustomFieldChoiceSetsReadOK) GetPayload() *models.CustomFieldChoiceSet {
	return o.Payload
}

func (o *ExtrasCustomFieldChoiceSetsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomFieldChoiceSet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasCustomFieldChoiceSetsReadDefault creates a ExtrasCustomFieldChoiceSetsReadDefault with default headers values
func NewExtrasCustomFieldChoiceSetsReadDefault(code int) *ExtrasCustomFieldChoiceSetsReadDefault {
	return &ExtrasCustomFieldChoiceSetsReadDefault{
		_statusCode: code,
	}
}

/*
ExtrasCustomFieldChoiceSetsReadDefault describes a response with status code -1, with default header values.

ExtrasCustomFieldChoiceSetsReadDefault extras custom field choice sets read default
*/
type ExtrasCustomFieldChoiceSetsReadDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the extras custom field choice sets read default response
func (o *ExtrasCustomFieldChoiceSetsReadDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this extras custom field choice sets read default response has a 2xx status code
func (o *ExtrasCustomFieldChoiceSetsReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this extras custom field choice sets read default response has a 3xx status code
func (o *ExtrasCustomFieldChoiceSetsReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this extras custom field choice sets read default response has a 4xx status code
func (o *ExtrasCustomFieldChoiceSetsReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this extras custom field choice sets read default response has a 5xx status code
func (o *ExtrasCustomFieldChoiceSetsReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this extras custom field choice sets read default response a status code equal to that given
func (o *ExtrasCustomFieldChoiceSetsReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ExtrasCustomFieldChoiceSetsReadDefault) Error() string {
	return fmt.Sprintf("[GET /extras/custom-field-choice-sets/{id}/][%d] extras_custom-field-choice-sets_read default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasCustomFieldChoiceSetsReadDefault) String() string {
	return fmt.Sprintf("[GET /extras/custom-field-choice-sets/{id}/][%d] extras_custom-field-choice-sets_read default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasCustomFieldChoiceSetsReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasCustomFieldChoiceSetsReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
