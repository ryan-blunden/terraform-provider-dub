// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/ryan-blunden/terraform-provider-dub/internal/sdk/models/shared"
	"net/http"
)

// CreateTagColor - The color of the tag. If not provided, a random color will be used from the list: red, yellow, green, blue, purple, pink, brown.
type CreateTagColor string

const (
	CreateTagColorRed    CreateTagColor = "red"
	CreateTagColorYellow CreateTagColor = "yellow"
	CreateTagColorGreen  CreateTagColor = "green"
	CreateTagColorBlue   CreateTagColor = "blue"
	CreateTagColorPurple CreateTagColor = "purple"
	CreateTagColorPink   CreateTagColor = "pink"
	CreateTagColorBrown  CreateTagColor = "brown"
)

func (e CreateTagColor) ToPointer() *CreateTagColor {
	return &e
}
func (e *CreateTagColor) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "red":
		fallthrough
	case "yellow":
		fallthrough
	case "green":
		fallthrough
	case "blue":
		fallthrough
	case "purple":
		fallthrough
	case "pink":
		fallthrough
	case "brown":
		*e = CreateTagColor(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateTagColor: %v", v)
	}
}

type CreateTagRequest struct {
	// The name of the tag to create.
	Name *string `json:"name,omitempty"`
	// The color of the tag. If not provided, a random color will be used from the list: red, yellow, green, blue, purple, pink, brown.
	Color *CreateTagColor `json:"color,omitempty"`
	// The name of the tag to create.
	//
	// Deprecated: This will be removed in a future release, please migrate away from it as soon as possible.
	Tag *string `json:"tag,omitempty"`
}

func (o *CreateTagRequest) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *CreateTagRequest) GetColor() *CreateTagColor {
	if o == nil {
		return nil
	}
	return o.Color
}

func (o *CreateTagRequest) GetTag() *string {
	if o == nil {
		return nil
	}
	return o.Tag
}

type CreateTagResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The created tag
	TagSchema *shared.TagSchema
	// The server cannot or will not process the request due to something that is perceived to be a client error (e.g., malformed request syntax, invalid request message framing, or deceptive request routing).
	FourHundred *shared.FourHundred
	// Although the HTTP standard specifies "unauthorized", semantically this response means "unauthenticated". That is, the client must authenticate itself to get the requested response.
	FourHundredAndOne *shared.FourHundredAndOne
	// The client does not have access rights to the content; that is, it is unauthorized, so the server is refusing to give the requested resource. Unlike 401 Unauthorized, the client's identity is known to the server.
	FourHundredAndThree *shared.FourHundredAndThree
	// The server cannot find the requested resource.
	FourHundredAndFour *shared.FourHundredAndFour
	// This response is sent when a request conflicts with the current state of the server.
	FourHundredAndNine *shared.FourHundredAndNine
	// This response is sent when the requested content has been permanently deleted from server, with no forwarding address.
	FourHundredAndTen *shared.FourHundredAndTen
	// The request was well-formed but was unable to be followed due to semantic errors.
	FourHundredAndTwentyTwo *shared.FourHundredAndTwentyTwo
	// The user has sent too many requests in a given amount of time ("rate limiting")
	FourHundredAndTwentyNine *shared.FourHundredAndTwentyNine
	// The server has encountered a situation it does not know how to handle.
	FiveHundred *shared.FiveHundred
}

func (o *CreateTagResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateTagResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateTagResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateTagResponse) GetTagSchema() *shared.TagSchema {
	if o == nil {
		return nil
	}
	return o.TagSchema
}

func (o *CreateTagResponse) GetFourHundred() *shared.FourHundred {
	if o == nil {
		return nil
	}
	return o.FourHundred
}

func (o *CreateTagResponse) GetFourHundredAndOne() *shared.FourHundredAndOne {
	if o == nil {
		return nil
	}
	return o.FourHundredAndOne
}

func (o *CreateTagResponse) GetFourHundredAndThree() *shared.FourHundredAndThree {
	if o == nil {
		return nil
	}
	return o.FourHundredAndThree
}

func (o *CreateTagResponse) GetFourHundredAndFour() *shared.FourHundredAndFour {
	if o == nil {
		return nil
	}
	return o.FourHundredAndFour
}

func (o *CreateTagResponse) GetFourHundredAndNine() *shared.FourHundredAndNine {
	if o == nil {
		return nil
	}
	return o.FourHundredAndNine
}

func (o *CreateTagResponse) GetFourHundredAndTen() *shared.FourHundredAndTen {
	if o == nil {
		return nil
	}
	return o.FourHundredAndTen
}

func (o *CreateTagResponse) GetFourHundredAndTwentyTwo() *shared.FourHundredAndTwentyTwo {
	if o == nil {
		return nil
	}
	return o.FourHundredAndTwentyTwo
}

func (o *CreateTagResponse) GetFourHundredAndTwentyNine() *shared.FourHundredAndTwentyNine {
	if o == nil {
		return nil
	}
	return o.FourHundredAndTwentyNine
}

func (o *CreateTagResponse) GetFiveHundred() *shared.FiveHundred {
	if o == nil {
		return nil
	}
	return o.FiveHundred
}
