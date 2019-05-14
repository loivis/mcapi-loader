// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/loivis/mcapi-loader/marvel/models"
)

// GetEventSeriesCollectionReader is a Reader for the GetEventSeriesCollection structure.
type GetEventSeriesCollectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEventSeriesCollectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetEventSeriesCollectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetEventSeriesCollectionOK creates a GetEventSeriesCollectionOK with default headers values
func NewGetEventSeriesCollectionOK() *GetEventSeriesCollectionOK {
	return &GetEventSeriesCollectionOK{}
}

/*GetEventSeriesCollectionOK handles this case with default header values.

No response was specified
*/
type GetEventSeriesCollectionOK struct {
	Payload *models.SeriesDataWrapper
}

func (o *GetEventSeriesCollectionOK) Error() string {
	return fmt.Sprintf("[GET /v1/public/events/{eventId}/series][%d] getEventSeriesCollectionOK  %+v", 200, o.Payload)
}

func (o *GetEventSeriesCollectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SeriesDataWrapper)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}