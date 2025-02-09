// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	core "go-mod-path/generated/go/core"
)

// Bad Request
type BadRequestError struct {
	*core.APIError
	Body *ErrorBody
}

func (b *BadRequestError) UnmarshalJSON(data []byte) error {
	var body *ErrorBody
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	b.StatusCode = 400
	b.Body = body
	return nil
}

func (b *BadRequestError) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.Body)
}

func (b *BadRequestError) Unwrap() error {
	return b.APIError
}

// Achievement Not Found
type NotFoundError struct {
	*core.APIError
	Body *ErrorBody
}

func (n *NotFoundError) UnmarshalJSON(data []byte) error {
	var body *ErrorBody
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	n.StatusCode = 404
	n.Body = body
	return nil
}

func (n *NotFoundError) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.Body)
}

func (n *NotFoundError) Unwrap() error {
	return n.APIError
}

// Unauthorized
type UnauthorizedError struct {
	*core.APIError
	Body *ErrorBody
}

func (u *UnauthorizedError) UnmarshalJSON(data []byte) error {
	var body *ErrorBody
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	u.StatusCode = 401
	u.Body = body
	return nil
}

func (u *UnauthorizedError) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.Body)
}

func (u *UnauthorizedError) Unwrap() error {
	return u.APIError
}

// Unprocessible Entity
type UnprocessableEntityError struct {
	*core.APIError
	Body *ErrorBody
}

func (u *UnprocessableEntityError) UnmarshalJSON(data []byte) error {
	var body *ErrorBody
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	u.StatusCode = 422
	u.Body = body
	return nil
}

func (u *UnprocessableEntityError) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.Body)
}

func (u *UnprocessableEntityError) Unwrap() error {
	return u.APIError
}
