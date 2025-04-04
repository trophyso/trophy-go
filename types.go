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

type StreakFrequency string

const (
	StreakFrequencyDaily   StreakFrequency = "daily"
	StreakFrequencyWeekly  StreakFrequency = "weekly"
	StreakFrequencyMonthly StreakFrequency = "monthly"
	StreakFrequencyYearly  StreakFrequency = "yearly"
)

func NewStreakFrequencyFromString(s string) (StreakFrequency, error) {
	switch s {
	case "daily":
		return StreakFrequencyDaily, nil
	case "weekly":
		return StreakFrequencyWeekly, nil
	case "monthly":
		return StreakFrequencyMonthly, nil
	case "yearly":
		return StreakFrequencyYearly, nil
	}
	var t StreakFrequency
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (s StreakFrequency) Ptr() *StreakFrequency {
	return &s
}

type StreakResponse struct {
	// The length of the user's current streak.
	Length int `json:"length" url:"length"`
	// The frequency of the streak.
	Frequency StreakFrequency `json:"frequency" url:"frequency"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (s *StreakResponse) GetLength() int {
	if s == nil {
		return 0
	}
	return s.Length
}

func (s *StreakResponse) GetFrequency() StreakFrequency {
	if s == nil {
		return ""
	}
	return s.Frequency
}

func (s *StreakResponse) GetExtraProperties() map[string]interface{} {
	return s.extraProperties
}

func (s *StreakResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler StreakResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = StreakResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	s.rawJSON = json.RawMessage(data)
	return nil
}

func (s *StreakResponse) String() string {
	if len(s.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(s.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(s); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", s)
}

// An object with editable user fields.
type UpdatedUser struct {
	// The user's email address. Required if subscribeToEmails is true.
	Email *string `json:"email,omitempty" url:"email,omitempty"`
	// The name to refer to the user by in emails.
	Name *string `json:"name,omitempty" url:"name,omitempty"`
	// The user's timezone (used for email scheduling).
	Tz *string `json:"tz,omitempty" url:"tz,omitempty"`
	// Whether the user should receive Trophy-powered emails. Cannot be false if an email is provided.
	SubscribeToEmails *bool `json:"subscribeToEmails,omitempty" url:"subscribeToEmails,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (u *UpdatedUser) GetEmail() *string {
	if u == nil {
		return nil
	}
	return u.Email
}

func (u *UpdatedUser) GetName() *string {
	if u == nil {
		return nil
	}
	return u.Name
}

func (u *UpdatedUser) GetTz() *string {
	if u == nil {
		return nil
	}
	return u.Tz
}

func (u *UpdatedUser) GetSubscribeToEmails() *bool {
	if u == nil {
		return nil
	}
	return u.SubscribeToEmails
}

func (u *UpdatedUser) GetExtraProperties() map[string]interface{} {
	return u.extraProperties
}

func (u *UpdatedUser) UnmarshalJSON(data []byte) error {
	type unmarshaler UpdatedUser
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UpdatedUser(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	u.rawJSON = json.RawMessage(data)
	return nil
}

func (u *UpdatedUser) String() string {
	if len(u.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(u.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(u); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", u)
}

// An object with editable user fields.
type UpsertedUser struct {
	// The user's email address. Required if subscribeToEmails is true.
	Email *string `json:"email,omitempty" url:"email,omitempty"`
	// The name to refer to the user by in emails.
	Name *string `json:"name,omitempty" url:"name,omitempty"`
	// The user's timezone (used for email scheduling).
	Tz *string `json:"tz,omitempty" url:"tz,omitempty"`
	// Whether the user should receive Trophy-powered emails. Cannot be false if an email is provided.
	SubscribeToEmails *bool `json:"subscribeToEmails,omitempty" url:"subscribeToEmails,omitempty"`
	// The ID of the user in your database. Must be a string.
	Id string `json:"id" url:"id"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (u *UpsertedUser) GetEmail() *string {
	if u == nil {
		return nil
	}
	return u.Email
}

func (u *UpsertedUser) GetName() *string {
	if u == nil {
		return nil
	}
	return u.Name
}

func (u *UpsertedUser) GetTz() *string {
	if u == nil {
		return nil
	}
	return u.Tz
}

func (u *UpsertedUser) GetSubscribeToEmails() *bool {
	if u == nil {
		return nil
	}
	return u.SubscribeToEmails
}

func (u *UpsertedUser) GetId() string {
	if u == nil {
		return ""
	}
	return u.Id
}

func (u *UpsertedUser) GetExtraProperties() map[string]interface{} {
	return u.extraProperties
}

func (u *UpsertedUser) UnmarshalJSON(data []byte) error {
	type unmarshaler UpsertedUser
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UpsertedUser(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	u.rawJSON = json.RawMessage(data)
	return nil
}

func (u *UpsertedUser) String() string {
	if len(u.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(u.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(u); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", u)
}
