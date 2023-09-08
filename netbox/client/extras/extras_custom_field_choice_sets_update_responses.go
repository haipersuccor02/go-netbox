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

// ExtrasCustomFieldChoiceSetsUpdateReader is a Reader for the ExtrasCustomFieldChoiceSetsUpdate structure.
type ExtrasCustomFieldChoiceSetsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasCustomFieldChoiceSetsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasCustomFieldChoiceSetsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasCustomFieldChoiceSetsUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasCustomFieldChoiceSetsUpdateOK creates a ExtrasCustomFieldChoiceSetsUpdateOK with default headers values
func NewExtrasCustomFieldChoiceSetsUpdateOK() *ExtrasCustomFieldChoiceSetsUpdateOK {
	return &ExtrasCustomFieldChoiceSetsUpdateOK{}
}

/*
ExtrasCustomFieldChoiceSetsUpdateOK describes a response with status code 200, with default header values.

ExtrasCustomFieldChoiceSetsUpdateOK extras custom field choice sets update o k
*/
type ExtrasCustomFieldChoiceSetsUpdateOK struct {
	Payload *models.CustomFieldChoiceSet
}

// IsSuccess returns true when this extras custom field choice sets update o k response has a 2xx status code
func (o *ExtrasCustomFieldChoiceSetsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras custom field choice sets update o k response has a 3xx status code
func (o *ExtrasCustomFieldChoiceSetsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras custom field choice sets update o k response has a 4xx status code
func (o *ExtrasCustomFieldChoiceSetsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras custom field choice sets update o k response has a 5xx status code
func (o *ExtrasCustomFieldChoiceSetsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras custom field choice sets update o k response a status code equal to that given
func (o *ExtrasCustomFieldChoiceSetsUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *ExtrasCustomFieldChoiceSetsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /extras/custom-field-choice-sets/{id}/][%d] extrasCustomFieldChoiceSetsUpdateOK  %+v", 200, o.Payload)
}

func (o *ExtrasCustomFieldChoiceSetsUpdateOK) String() string {
	return fmt.Sprintf("[PUT /extras/custom-field-choice-sets/{id}/][%d] extrasCustomFieldChoiceSetsUpdateOK  %+v", 200, o.Payload)
}

func (o *ExtrasCustomFieldChoiceSetsUpdateOK) GetPayload() *models.CustomFieldChoiceSet {
	return o.Payload
}

func (o *ExtrasCustomFieldChoiceSetsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomFieldChoiceSet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasCustomFieldChoiceSetsUpdateDefault creates a ExtrasCustomFieldChoiceSetsUpdateDefault with default headers values
func NewExtrasCustomFieldChoiceSetsUpdateDefault(code int) *ExtrasCustomFieldChoiceSetsUpdateDefault {
	return &ExtrasCustomFieldChoiceSetsUpdateDefault{
		_statusCode: code,
	}
}

/*
ExtrasCustomFieldChoiceSetsUpdateDefault describes a response with status code -1, with default header values.

ExtrasCustomFieldChoiceSetsUpdateDefault extras custom field choice sets update default
*/
type ExtrasCustomFieldChoiceSetsUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the extras custom field choice sets update default response
func (o *ExtrasCustomFieldChoiceSetsUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this extras custom field choice sets update default response has a 2xx status code
func (o *ExtrasCustomFieldChoiceSetsUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this extras custom field choice sets update default response has a 3xx status code
func (o *ExtrasCustomFieldChoiceSetsUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this extras custom field choice sets update default response has a 4xx status code
func (o *ExtrasCustomFieldChoiceSetsUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this extras custom field choice sets update default response has a 5xx status code
func (o *ExtrasCustomFieldChoiceSetsUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this extras custom field choice sets update default response a status code equal to that given
func (o *ExtrasCustomFieldChoiceSetsUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ExtrasCustomFieldChoiceSetsUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /extras/custom-field-choice-sets/{id}/][%d] extras_custom-field-choice-sets_update default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasCustomFieldChoiceSetsUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /extras/custom-field-choice-sets/{id}/][%d] extras_custom-field-choice-sets_update default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasCustomFieldChoiceSetsUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasCustomFieldChoiceSetsUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
