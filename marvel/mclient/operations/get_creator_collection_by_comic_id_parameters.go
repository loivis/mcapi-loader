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

// NewGetCreatorCollectionByComicIDParams creates a new GetCreatorCollectionByComicIDParams object
// with the default values initialized.
func NewGetCreatorCollectionByComicIDParams() *GetCreatorCollectionByComicIDParams {
	var ()
	return &GetCreatorCollectionByComicIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCreatorCollectionByComicIDParamsWithTimeout creates a new GetCreatorCollectionByComicIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCreatorCollectionByComicIDParamsWithTimeout(timeout time.Duration) *GetCreatorCollectionByComicIDParams {
	var ()
	return &GetCreatorCollectionByComicIDParams{

		timeout: timeout,
	}
}

// NewGetCreatorCollectionByComicIDParamsWithContext creates a new GetCreatorCollectionByComicIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCreatorCollectionByComicIDParamsWithContext(ctx context.Context) *GetCreatorCollectionByComicIDParams {
	var ()
	return &GetCreatorCollectionByComicIDParams{

		Context: ctx,
	}
}

// NewGetCreatorCollectionByComicIDParamsWithHTTPClient creates a new GetCreatorCollectionByComicIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCreatorCollectionByComicIDParamsWithHTTPClient(client *http.Client) *GetCreatorCollectionByComicIDParams {
	var ()
	return &GetCreatorCollectionByComicIDParams{
		HTTPClient: client,
	}
}

/*GetCreatorCollectionByComicIDParams contains all the parameters to send to the API endpoint
for the get creator collection by comic Id operation typically these are written to a http.Request
*/
type GetCreatorCollectionByComicIDParams struct {

	/*Apikey
	  [Auth] public apikey

	*/
	Apikey string
	/*ComicID
	  The comic id.

	*/
	ComicID int32
	/*Comics
	  Return only creators who worked on in the specified comics (accepts a comma-separated list of ids).

	*/
	Comics []int32
	/*FirstName
	  Filter by creator first name (e.g. brian).

	*/
	FirstName *string
	/*FirstNameStartsWith
	  Filter by creator first names that match critera (e.g. B, St L).

	*/
	FirstNameStartsWith *string
	/*Hash
	  [Auth] md5 digest of concatenation of ts, private key, public key

	*/
	Hash string
	/*LastName
	  Filter by creator last name (e.g. Bendis).

	*/
	LastName *string
	/*LastNameStartsWith
	  Filter by creator last names that match critera (e.g. Ben).

	*/
	LastNameStartsWith *string
	/*Limit
	  Limit the result set to the specified number of resources.

	*/
	Limit *int32
	/*MiddleName
	  Filter by creator middle name (e.g. Michael).

	*/
	MiddleName *string
	/*MiddleNameStartsWith
	  Filter by creator middle names that match critera (e.g. Mi).

	*/
	MiddleNameStartsWith *string
	/*ModifiedSince
	  Return only creators which have been modified since the specified date.

	*/
	ModifiedSince *strfmt.Date
	/*NameStartsWith
	  Filter by creator names that match critera (e.g. B, St L).

	*/
	NameStartsWith *string
	/*Offset
	  Skip the specified number of resources in the result set.

	*/
	Offset *int32
	/*OrderBy
	  Order the result set by a field or fields. Add a "-" to the value sort in descending order. Multiple values are given priority in the order in which they are passed.

	*/
	OrderBy []string
	/*Series
	  Return only creators who worked on the specified series (accepts a comma-separated list of ids).

	*/
	Series []int32
	/*Stories
	  Return only creators who worked on the specified stories (accepts a comma-separated list of ids).

	*/
	Stories []int32
	/*Suffix
	  Filter by suffix or honorific (e.g. Jr., Sr.).

	*/
	Suffix *string
	/*Ts
	  [Auth] timestamp

	*/
	Ts string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get creator collection by comic Id params
func (o *GetCreatorCollectionByComicIDParams) WithTimeout(timeout time.Duration) *GetCreatorCollectionByComicIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get creator collection by comic Id params
func (o *GetCreatorCollectionByComicIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get creator collection by comic Id params
func (o *GetCreatorCollectionByComicIDParams) WithContext(ctx context.Context) *GetCreatorCollectionByComicIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get creator collection by comic Id params
func (o *GetCreatorCollectionByComicIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get creator collection by comic Id params
func (o *GetCreatorCollectionByComicIDParams) WithHTTPClient(client *http.Client) *GetCreatorCollectionByComicIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get creator collection by comic Id params
func (o *GetCreatorCollectionByComicIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApikey adds the apikey to the get creator collection by comic Id params
func (o *GetCreatorCollectionByComicIDParams) WithApikey(apikey string) *GetCreatorCollectionByComicIDParams {
	o.SetApikey(apikey)
	return o
}

// SetApikey adds the apikey to the get creator collection by comic Id params
func (o *GetCreatorCollectionByComicIDParams) SetApikey(apikey string) {
	o.Apikey = apikey
}

// WithComicID adds the comicID to the get creator collection by comic Id params
func (o *GetCreatorCollectionByComicIDParams) WithComicID(comicID int32) *GetCreatorCollectionByComicIDParams {
	o.SetComicID(comicID)
	return o
}

// SetComicID adds the comicId to the get creator collection by comic Id params
func (o *GetCreatorCollectionByComicIDParams) SetComicID(comicID int32) {
	o.ComicID = comicID
}

// WithComics adds the comics to the get creator collection by comic Id params
func (o *GetCreatorCollectionByComicIDParams) WithComics(comics []int32) *GetCreatorCollectionByComicIDParams {
	o.SetComics(comics)
	return o
}

// SetComics adds the comics to the get creator collection by comic Id params
func (o *GetCreatorCollectionByComicIDParams) SetComics(comics []int32) {
	o.Comics = comics
}

// WithFirstName adds the firstName to the get creator collection by comic Id params
func (o *GetCreatorCollectionByComicIDParams) WithFirstName(firstName *string) *GetCreatorCollectionByComicIDParams {
	o.SetFirstName(firstName)
	return o
}

// SetFirstName adds the firstName to the get creator collection by comic Id params
func (o *GetCreatorCollectionByComicIDParams) SetFirstName(firstName *string) {
	o.FirstName = firstName
}

// WithFirstNameStartsWith adds the firstNameStartsWith to the get creator collection by comic Id params
func (o *GetCreatorCollectionByComicIDParams) WithFirstNameStartsWith(firstNameStartsWith *string) *GetCreatorCollectionByComicIDParams {
	o.SetFirstNameStartsWith(firstNameStartsWith)
	return o
}

// SetFirstNameStartsWith adds the firstNameStartsWith to the get creator collection by comic Id params
func (o *GetCreatorCollectionByComicIDParams) SetFirstNameStartsWith(firstNameStartsWith *string) {
	o.FirstNameStartsWith = firstNameStartsWith
}

// WithHash adds the hash to the get creator collection by comic Id params
func (o *GetCreatorCollectionByComicIDParams) WithHash(hash string) *GetCreatorCollectionByComicIDParams {
	o.SetHash(hash)
	return o
}

// SetHash adds the hash to the get creator collection by comic Id params
func (o *GetCreatorCollectionByComicIDParams) SetHash(hash string) {
	o.Hash = hash
}

// WithLastName adds the lastName to the get creator collection by comic Id params
func (o *GetCreatorCollectionByComicIDParams) WithLastName(lastName *string) *GetCreatorCollectionByComicIDParams {
	o.SetLastName(lastName)
	return o
}

// SetLastName adds the lastName to the get creator collection by comic Id params
func (o *GetCreatorCollectionByComicIDParams) SetLastName(lastName *string) {
	o.LastName = lastName
}

// WithLastNameStartsWith adds the lastNameStartsWith to the get creator collection by comic Id params
func (o *GetCreatorCollectionByComicIDParams) WithLastNameStartsWith(lastNameStartsWith *string) *GetCreatorCollectionByComicIDParams {
	o.SetLastNameStartsWith(lastNameStartsWith)
	return o
}

// SetLastNameStartsWith adds the lastNameStartsWith to the get creator collection by comic Id params
func (o *GetCreatorCollectionByComicIDParams) SetLastNameStartsWith(lastNameStartsWith *string) {
	o.LastNameStartsWith = lastNameStartsWith
}

// WithLimit adds the limit to the get creator collection by comic Id params
func (o *GetCreatorCollectionByComicIDParams) WithLimit(limit *int32) *GetCreatorCollectionByComicIDParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get creator collection by comic Id params
func (o *GetCreatorCollectionByComicIDParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithMiddleName adds the middleName to the get creator collection by comic Id params
func (o *GetCreatorCollectionByComicIDParams) WithMiddleName(middleName *string) *GetCreatorCollectionByComicIDParams {
	o.SetMiddleName(middleName)
	return o
}

// SetMiddleName adds the middleName to the get creator collection by comic Id params
func (o *GetCreatorCollectionByComicIDParams) SetMiddleName(middleName *string) {
	o.MiddleName = middleName
}

// WithMiddleNameStartsWith adds the middleNameStartsWith to the get creator collection by comic Id params
func (o *GetCreatorCollectionByComicIDParams) WithMiddleNameStartsWith(middleNameStartsWith *string) *GetCreatorCollectionByComicIDParams {
	o.SetMiddleNameStartsWith(middleNameStartsWith)
	return o
}

// SetMiddleNameStartsWith adds the middleNameStartsWith to the get creator collection by comic Id params
func (o *GetCreatorCollectionByComicIDParams) SetMiddleNameStartsWith(middleNameStartsWith *string) {
	o.MiddleNameStartsWith = middleNameStartsWith
}

// WithModifiedSince adds the modifiedSince to the get creator collection by comic Id params
func (o *GetCreatorCollectionByComicIDParams) WithModifiedSince(modifiedSince *strfmt.Date) *GetCreatorCollectionByComicIDParams {
	o.SetModifiedSince(modifiedSince)
	return o
}

// SetModifiedSince adds the modifiedSince to the get creator collection by comic Id params
func (o *GetCreatorCollectionByComicIDParams) SetModifiedSince(modifiedSince *strfmt.Date) {
	o.ModifiedSince = modifiedSince
}

// WithNameStartsWith adds the nameStartsWith to the get creator collection by comic Id params
func (o *GetCreatorCollectionByComicIDParams) WithNameStartsWith(nameStartsWith *string) *GetCreatorCollectionByComicIDParams {
	o.SetNameStartsWith(nameStartsWith)
	return o
}

// SetNameStartsWith adds the nameStartsWith to the get creator collection by comic Id params
func (o *GetCreatorCollectionByComicIDParams) SetNameStartsWith(nameStartsWith *string) {
	o.NameStartsWith = nameStartsWith
}

// WithOffset adds the offset to the get creator collection by comic Id params
func (o *GetCreatorCollectionByComicIDParams) WithOffset(offset *int32) *GetCreatorCollectionByComicIDParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get creator collection by comic Id params
func (o *GetCreatorCollectionByComicIDParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithOrderBy adds the orderBy to the get creator collection by comic Id params
func (o *GetCreatorCollectionByComicIDParams) WithOrderBy(orderBy []string) *GetCreatorCollectionByComicIDParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the get creator collection by comic Id params
func (o *GetCreatorCollectionByComicIDParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithSeries adds the series to the get creator collection by comic Id params
func (o *GetCreatorCollectionByComicIDParams) WithSeries(series []int32) *GetCreatorCollectionByComicIDParams {
	o.SetSeries(series)
	return o
}

// SetSeries adds the series to the get creator collection by comic Id params
func (o *GetCreatorCollectionByComicIDParams) SetSeries(series []int32) {
	o.Series = series
}

// WithStories adds the stories to the get creator collection by comic Id params
func (o *GetCreatorCollectionByComicIDParams) WithStories(stories []int32) *GetCreatorCollectionByComicIDParams {
	o.SetStories(stories)
	return o
}

// SetStories adds the stories to the get creator collection by comic Id params
func (o *GetCreatorCollectionByComicIDParams) SetStories(stories []int32) {
	o.Stories = stories
}

// WithSuffix adds the suffix to the get creator collection by comic Id params
func (o *GetCreatorCollectionByComicIDParams) WithSuffix(suffix *string) *GetCreatorCollectionByComicIDParams {
	o.SetSuffix(suffix)
	return o
}

// SetSuffix adds the suffix to the get creator collection by comic Id params
func (o *GetCreatorCollectionByComicIDParams) SetSuffix(suffix *string) {
	o.Suffix = suffix
}

// WithTs adds the ts to the get creator collection by comic Id params
func (o *GetCreatorCollectionByComicIDParams) WithTs(ts string) *GetCreatorCollectionByComicIDParams {
	o.SetTs(ts)
	return o
}

// SetTs adds the ts to the get creator collection by comic Id params
func (o *GetCreatorCollectionByComicIDParams) SetTs(ts string) {
	o.Ts = ts
}

// WriteToRequest writes these params to a swagger request
func (o *GetCreatorCollectionByComicIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param comicId
	if err := r.SetPathParam("comicId", swag.FormatInt32(o.ComicID)); err != nil {
		return err
	}

	var valuesComics []string
	for _, v := range o.Comics {
		valuesComics = append(valuesComics, swag.FormatInt32(v))
	}

	joinedComics := swag.JoinByFormat(valuesComics, "")
	// query array param comics
	if err := r.SetQueryParam("comics", joinedComics...); err != nil {
		return err
	}

	if o.FirstName != nil {

		// query param firstName
		var qrFirstName string
		if o.FirstName != nil {
			qrFirstName = *o.FirstName
		}
		qFirstName := qrFirstName
		if qFirstName != "" {
			if err := r.SetQueryParam("firstName", qFirstName); err != nil {
				return err
			}
		}

	}

	if o.FirstNameStartsWith != nil {

		// query param firstNameStartsWith
		var qrFirstNameStartsWith string
		if o.FirstNameStartsWith != nil {
			qrFirstNameStartsWith = *o.FirstNameStartsWith
		}
		qFirstNameStartsWith := qrFirstNameStartsWith
		if qFirstNameStartsWith != "" {
			if err := r.SetQueryParam("firstNameStartsWith", qFirstNameStartsWith); err != nil {
				return err
			}
		}

	}

	// query param hash
	qrHash := o.Hash
	qHash := qrHash
	if qHash != "" {
		if err := r.SetQueryParam("hash", qHash); err != nil {
			return err
		}
	}

	if o.LastName != nil {

		// query param lastName
		var qrLastName string
		if o.LastName != nil {
			qrLastName = *o.LastName
		}
		qLastName := qrLastName
		if qLastName != "" {
			if err := r.SetQueryParam("lastName", qLastName); err != nil {
				return err
			}
		}

	}

	if o.LastNameStartsWith != nil {

		// query param lastNameStartsWith
		var qrLastNameStartsWith string
		if o.LastNameStartsWith != nil {
			qrLastNameStartsWith = *o.LastNameStartsWith
		}
		qLastNameStartsWith := qrLastNameStartsWith
		if qLastNameStartsWith != "" {
			if err := r.SetQueryParam("lastNameStartsWith", qLastNameStartsWith); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int32
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt32(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.MiddleName != nil {

		// query param middleName
		var qrMiddleName string
		if o.MiddleName != nil {
			qrMiddleName = *o.MiddleName
		}
		qMiddleName := qrMiddleName
		if qMiddleName != "" {
			if err := r.SetQueryParam("middleName", qMiddleName); err != nil {
				return err
			}
		}

	}

	if o.MiddleNameStartsWith != nil {

		// query param middleNameStartsWith
		var qrMiddleNameStartsWith string
		if o.MiddleNameStartsWith != nil {
			qrMiddleNameStartsWith = *o.MiddleNameStartsWith
		}
		qMiddleNameStartsWith := qrMiddleNameStartsWith
		if qMiddleNameStartsWith != "" {
			if err := r.SetQueryParam("middleNameStartsWith", qMiddleNameStartsWith); err != nil {
				return err
			}
		}

	}

	if o.ModifiedSince != nil {

		// query param modifiedSince
		var qrModifiedSince strfmt.Date
		if o.ModifiedSince != nil {
			qrModifiedSince = *o.ModifiedSince
		}
		qModifiedSince := qrModifiedSince.String()
		if qModifiedSince != "" {
			if err := r.SetQueryParam("modifiedSince", qModifiedSince); err != nil {
				return err
			}
		}

	}

	if o.NameStartsWith != nil {

		// query param nameStartsWith
		var qrNameStartsWith string
		if o.NameStartsWith != nil {
			qrNameStartsWith = *o.NameStartsWith
		}
		qNameStartsWith := qrNameStartsWith
		if qNameStartsWith != "" {
			if err := r.SetQueryParam("nameStartsWith", qNameStartsWith); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int32
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt32(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	valuesOrderBy := o.OrderBy

	joinedOrderBy := swag.JoinByFormat(valuesOrderBy, "")
	// query array param orderBy
	if err := r.SetQueryParam("orderBy", joinedOrderBy...); err != nil {
		return err
	}

	var valuesSeries []string
	for _, v := range o.Series {
		valuesSeries = append(valuesSeries, swag.FormatInt32(v))
	}

	joinedSeries := swag.JoinByFormat(valuesSeries, "")
	// query array param series
	if err := r.SetQueryParam("series", joinedSeries...); err != nil {
		return err
	}

	var valuesStories []string
	for _, v := range o.Stories {
		valuesStories = append(valuesStories, swag.FormatInt32(v))
	}

	joinedStories := swag.JoinByFormat(valuesStories, "")
	// query array param stories
	if err := r.SetQueryParam("stories", joinedStories...); err != nil {
		return err
	}

	if o.Suffix != nil {

		// query param suffix
		var qrSuffix string
		if o.Suffix != nil {
			qrSuffix = *o.Suffix
		}
		qSuffix := qrSuffix
		if qSuffix != "" {
			if err := r.SetQueryParam("suffix", qSuffix); err != nil {
				return err
			}
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