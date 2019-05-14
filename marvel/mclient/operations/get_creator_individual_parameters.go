// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetCreatorIndividualParams creates a new GetCreatorIndividualParams object
// with the default values initialized.
func NewGetCreatorIndividualParams() *GetCreatorIndividualParams {
	var ()
	return &GetCreatorIndividualParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCreatorIndividualParamsWithTimeout creates a new GetCreatorIndividualParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCreatorIndividualParamsWithTimeout(timeout time.Duration) *GetCreatorIndividualParams {
	var ()
	return &GetCreatorIndividualParams{

		timeout: timeout,
	}
}

// NewGetCreatorIndividualParamsWithContext creates a new GetCreatorIndividualParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCreatorIndividualParamsWithContext(ctx context.Context) *GetCreatorIndividualParams {
	var ()
	return &GetCreatorIndividualParams{

		Context: ctx,
	}
}

// NewGetCreatorIndividualParamsWithHTTPClient creates a new GetCreatorIndividualParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCreatorIndividualParamsWithHTTPClient(client *http.Client) *GetCreatorIndividualParams {
	var ()
	return &GetCreatorIndividualParams{
		HTTPClient: client,
	}
}

/*GetCreatorIndividualParams contains all the parameters to send to the API endpoint
for the get creator individual operation typically these are written to a http.Request
*/
type GetCreatorIndividualParams struct {

	/*Apikey
	  [Auth] public apikey

	*/
	Apikey string
	/*CreatorID
	  A single creator id.

	*/
	CreatorID int32
	/*Hash
	  [Auth] md5 digest of concatenation of ts, private key, public key

	*/
	Hash string
	/*Ts
	  [Auth] timestamp

	*/
	Ts string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get creator individual params
func (o *GetCreatorIndividualParams) WithTimeout(timeout time.Duration) *GetCreatorIndividualParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get creator individual params
func (o *GetCreatorIndividualParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get creator individual params
func (o *GetCreatorIndividualParams) WithContext(ctx context.Context) *GetCreatorIndividualParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get creator individual params
func (o *GetCreatorIndividualParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get creator individual params
func (o *GetCreatorIndividualParams) WithHTTPClient(client *http.Client) *GetCreatorIndividualParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get creator individual params
func (o *GetCreatorIndividualParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApikey adds the apikey to the get creator individual params
func (o *GetCreatorIndividualParams) WithApikey(apikey string) *GetCreatorIndividualParams {
	o.SetApikey(apikey)
	return o
}

// SetApikey adds the apikey to the get creator individual params
func (o *GetCreatorIndividualParams) SetApikey(apikey string) {
	o.Apikey = apikey
}

// WithCreatorID adds the creatorID to the get creator individual params
func (o *GetCreatorIndividualParams) WithCreatorID(creatorID int32) *GetCreatorIndividualParams {
	o.SetCreatorID(creatorID)
	return o
}

// SetCreatorID adds the creatorId to the get creator individual params
func (o *GetCreatorIndividualParams) SetCreatorID(creatorID int32) {
	o.CreatorID = creatorID
}

// WithHash adds the hash to the get creator individual params
func (o *GetCreatorIndividualParams) WithHash(hash string) *GetCreatorIndividualParams {
	o.SetHash(hash)
	return o
}

// SetHash adds the hash to the get creator individual params
func (o *GetCreatorIndividualParams) SetHash(hash string) {
	o.Hash = hash
}

// WithTs adds the ts to the get creator individual params
func (o *GetCreatorIndividualParams) WithTs(ts string) *GetCreatorIndividualParams {
	o.SetTs(ts)
	return o
}

// SetTs adds the ts to the get creator individual params
func (o *GetCreatorIndividualParams) SetTs(ts string) {
	o.Ts = ts
}

// WriteToRequest writes these params to a swagger request
func (o *GetCreatorIndividualParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param apikey
	qrApikey := o.Apikey
	qApikey := qrApikey
	if qApikey != "" {
		if err := r.SetQueryParam("apikey", qApikey); err != nil {
			return err
		}
	}

	// path param creatorId
	if err := r.SetPathParam("creatorId", swag.FormatInt32(o.CreatorID)); err != nil {
		return err
	}

	// query param hash
	qrHash := o.Hash
	qHash := qrHash
	if qHash != "" {
		if err := r.SetQueryParam("hash", qHash); err != nil {
			return err
		}
	}

	// query param ts
	qrTs := o.Ts
	qTs := qrTs
	if qTs != "" {
		if err := r.SetQueryParam("ts", qTs); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}