package openapi3

import (
	"context"

	"github.com/colin-z/kin-openapi/jsoninfo"
)

// RequestBody is specified by OpenAPI/Swagger 3.0 standard.
type RequestBody struct {
	ExtensionProps
	Description string  `json:"description,omitempty" yaml:"description,omitempty"`
	Required    bool    `json:"required,omitempty" yaml:"required,omitempty"`
	Content     Content `json:"content,omitempty" yaml:"content,omitempty"`
}

func NewRequestBody() *RequestBody {
	return &RequestBody{}
}

func (requestBody *RequestBody) WithDescription(value string) *RequestBody {
	requestBody.Description = value
	return requestBody
}

func (requestBody *RequestBody) WithRequired(value bool) *RequestBody {
	requestBody.Required = value
	return requestBody
}

func (requestBody *RequestBody) WithContent(content Content) *RequestBody {
	requestBody.Content = content
	return requestBody
}

func (requestBody *RequestBody) WithJSONSchemaRef(value *SchemaRef) *RequestBody {
	requestBody.Content = NewContentWithJSONSchemaRef(value)
	return requestBody
}

func (requestBody *RequestBody) WithJSONSchema(value *Schema) *RequestBody {
	requestBody.Content = NewContentWithJSONSchema(value)
	return requestBody
}

func (requestBody *RequestBody) GetMediaType(mediaType string) *MediaType {
	m := requestBody.Content
	if m == nil {
		return nil
	}
	return m[mediaType]
}

func (requestBody *RequestBody) MarshalJSON() ([]byte, error) {
	return jsoninfo.MarshalStrictStruct(requestBody)
}

func (requestBody *RequestBody) UnmarshalJSON(data []byte) error {
	return jsoninfo.UnmarshalStrictStruct(data, requestBody)
}

func (requestBody *RequestBody) Validate(c context.Context) error {
	if v := requestBody.Content; v != nil {
		if err := v.Validate(c); err != nil {
			return err
		}
	}
	return nil
}
