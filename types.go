// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/trophyso/trophy-go/internal"
	time "time"
)

type AchievementResponse struct {
	// The unique ID of the achievement.
	Id string `json:"id" url:"id"`
	// The name of this achievement.
	Name *string `json:"name,omitempty" url:"name,omitempty"`
	// The URL of the badge image for the achievement, if one has been uploaded.
	BadgeUrl *string `json:"badgeUrl,omitempty" url:"badgeUrl,omitempty"`
	// The ID of the metric associated with this achievement, if any.
	MetricId *string `json:"metricId,omitempty" url:"metricId,omitempty"`
	// The value of the metric required to complete the achievement, if this achievement is associated with a metric.
	MetricValue *float64 `json:"metricValue,omitempty" url:"metricValue,omitempty"`
	// The name of the metric associated with this achievement, if any.
	MetricName *string `json:"metricName,omitempty" url:"metricName,omitempty"`
	// The key used to reference this achievement in the API.
	Key *string `json:"key,omitempty" url:"key,omitempty"`
	// The date and time the achievement was completed, in ISO 8601 format.
	AchievedAt *time.Time `json:"achievedAt,omitempty" url:"achievedAt,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (a *AchievementResponse) GetId() string {
	if a == nil {
		return ""
	}
	return a.Id
}

func (a *AchievementResponse) GetName() *string {
	if a == nil {
		return nil
	}
	return a.Name
}

func (a *AchievementResponse) GetBadgeUrl() *string {
	if a == nil {
		return nil
	}
	return a.BadgeUrl
}

func (a *AchievementResponse) GetMetricId() *string {
	if a == nil {
		return nil
	}
	return a.MetricId
}

func (a *AchievementResponse) GetMetricValue() *float64 {
	if a == nil {
		return nil
	}
	return a.MetricValue
}

func (a *AchievementResponse) GetMetricName() *string {
	if a == nil {
		return nil
	}
	return a.MetricName
}

func (a *AchievementResponse) GetKey() *string {
	if a == nil {
		return nil
	}
	return a.Key
}

func (a *AchievementResponse) GetAchievedAt() *time.Time {
	if a == nil {
		return nil
	}
	return a.AchievedAt
}

func (a *AchievementResponse) GetExtraProperties() map[string]interface{} {
	return a.extraProperties
}

func (a *AchievementResponse) UnmarshalJSON(data []byte) error {
	type embed AchievementResponse
	var unmarshaler = struct {
		embed
		AchievedAt *internal.DateTime `json:"achievedAt,omitempty"`
	}{
		embed: embed(*a),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*a = AchievementResponse(unmarshaler.embed)
	a.AchievedAt = unmarshaler.AchievedAt.TimePtr()
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	a.rawJSON = json.RawMessage(data)
	return nil
}

func (a *AchievementResponse) MarshalJSON() ([]byte, error) {
	type embed AchievementResponse
	var marshaler = struct {
		embed
		AchievedAt *internal.DateTime `json:"achievedAt,omitempty"`
	}{
		embed:      embed(*a),
		AchievedAt: internal.NewOptionalDateTime(a.AchievedAt),
	}
	return json.Marshal(marshaler)
}

func (a *AchievementResponse) String() string {
	if len(a.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(a.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(a); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", a)
}

type ErrorBody struct {
	Error string `json:"error" url:"error"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (e *ErrorBody) GetError() string {
	if e == nil {
		return ""
	}
	return e.Error
}

func (e *ErrorBody) GetExtraProperties() map[string]interface{} {
	return e.extraProperties
}

func (e *ErrorBody) UnmarshalJSON(data []byte) error {
	type unmarshaler ErrorBody
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*e = ErrorBody(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *e)
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties
	e.rawJSON = json.RawMessage(data)
	return nil
}

func (e *ErrorBody) String() string {
	if len(e.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(e.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(e); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", e)
}

// The user that triggered the event.
type EventRequestUser struct {
	// The ID of the user in your database. Must be a string.
	Id string `json:"id" url:"id"`
	// The user's email address.
	Email *string `json:"email,omitempty" url:"email,omitempty"`
	// The name to refer to the user by in emails.
	Name *string `json:"name,omitempty" url:"name,omitempty"`
	// The user's timezone (used for email scheduling).
	Tz *string `json:"tz,omitempty" url:"tz,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (e *EventRequestUser) GetId() string {
	if e == nil {
		return ""
	}
	return e.Id
}

func (e *EventRequestUser) GetEmail() *string {
	if e == nil {
		return nil
	}
	return e.Email
}

func (e *EventRequestUser) GetName() *string {
	if e == nil {
		return nil
	}
	return e.Name
}

func (e *EventRequestUser) GetTz() *string {
	if e == nil {
		return nil
	}
	return e.Tz
}

func (e *EventRequestUser) GetExtraProperties() map[string]interface{} {
	return e.extraProperties
}

func (e *EventRequestUser) UnmarshalJSON(data []byte) error {
	type unmarshaler EventRequestUser
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*e = EventRequestUser(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *e)
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties
	e.rawJSON = json.RawMessage(data)
	return nil
}

func (e *EventRequestUser) String() string {
	if len(e.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(e.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(e); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", e)
}
