package j_team_invitation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPIJTeamInvitationRemoveIDReader is a Reader for the PostRemoteAPIJTeamInvitationRemoveID structure.
type PostRemoteAPIJTeamInvitationRemoveIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJTeamInvitationRemoveIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJTeamInvitationRemoveIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJTeamInvitationRemoveIDOK creates a PostRemoteAPIJTeamInvitationRemoveIDOK with default headers values
func NewPostRemoteAPIJTeamInvitationRemoveIDOK() *PostRemoteAPIJTeamInvitationRemoveIDOK {
	return &PostRemoteAPIJTeamInvitationRemoveIDOK{}
}

/*PostRemoteAPIJTeamInvitationRemoveIDOK handles this case with default header values.

OK
*/
type PostRemoteAPIJTeamInvitationRemoveIDOK struct {
	Payload PostRemoteAPIJTeamInvitationRemoveIDOKBody
}

func (o *PostRemoteAPIJTeamInvitationRemoveIDOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JTeamInvitation.remove/{id}][%d] postRemoteApiJTeamInvitationRemoveIdOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJTeamInvitationRemoveIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostRemoteAPIJTeamInvitationRemoveIDOKBody post remote API j team invitation remove ID o k body
swagger:model PostRemoteAPIJTeamInvitationRemoveIDOKBody
*/
type PostRemoteAPIJTeamInvitationRemoveIDOKBody struct {
	models.JTeamInvitation

	models.DefaultResponse
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostRemoteAPIJTeamInvitationRemoveIDOKBody) UnmarshalJSON(raw []byte) error {

	var postRemoteAPIJTeamInvitationRemoveIDOKBodyAO0 models.JTeamInvitation
	if err := swag.ReadJSON(raw, &postRemoteAPIJTeamInvitationRemoveIDOKBodyAO0); err != nil {
		return err
	}
	o.JTeamInvitation = postRemoteAPIJTeamInvitationRemoveIDOKBodyAO0

	var postRemoteAPIJTeamInvitationRemoveIDOKBodyAO1 models.DefaultResponse
	if err := swag.ReadJSON(raw, &postRemoteAPIJTeamInvitationRemoveIDOKBodyAO1); err != nil {
		return err
	}
	o.DefaultResponse = postRemoteAPIJTeamInvitationRemoveIDOKBodyAO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostRemoteAPIJTeamInvitationRemoveIDOKBody) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	postRemoteAPIJTeamInvitationRemoveIDOKBodyAO0, err := swag.WriteJSON(o.JTeamInvitation)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postRemoteAPIJTeamInvitationRemoveIDOKBodyAO0)

	postRemoteAPIJTeamInvitationRemoveIDOKBodyAO1, err := swag.WriteJSON(o.DefaultResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postRemoteAPIJTeamInvitationRemoveIDOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post remote API j team invitation remove ID o k body
func (o *PostRemoteAPIJTeamInvitationRemoveIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.JTeamInvitation.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.DefaultResponse.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
