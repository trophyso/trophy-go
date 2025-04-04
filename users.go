// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/trophyso/trophy-go/internal"
	time "time"
)

type MetricResponse struct {
	// The unique ID of the metric.
	Id string `json:"id" url:"id"`
	// The unique key of the metric.
	Key string `json:"key" url:"key"`
	// The name of the metric.
	Name string `json:"name" url:"name"`
	// The emoji to represent the metric.
	Emoji string `json:"emoji" url:"emoji"`
	// The frequency of the streak.
	StreakFrequency StreakFrequency `json:"streakFrequency" url:"streakFrequency"`
	// The status of the metric.
	Status MetricStatus `json:"status" url:"status"`
	// The user's current total for the metric.
	Current float64 `json:"current" url:"current"`
	// A list of the metric's achievements and the user's progress towards each.
	Achievements []*AchievementResponse `json:"achievements,omitempty" url:"achievements,omitempty"`
	// The user's current streak for the metric, if the metric has streaks enabled.
	CurrentStreak *StreakResponse `json:"currentStreak,omitempty" url:"currentStreak,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (m *MetricResponse) GetId() string {
	if m == nil {
		return ""
	}
	return m.Id
}

func (m *MetricResponse) GetKey() string {
	if m == nil {
		return ""
	}
	return m.Key
}

func (m *MetricResponse) GetName() string {
	if m == nil {
		return ""
	}
	return m.Name
}

func (m *MetricResponse) GetEmoji() string {
	if m == nil {
		return ""
	}
	return m.Emoji
}

func (m *MetricResponse) GetStreakFrequency() StreakFrequency {
	if m == nil {
		return ""
	}
	return m.StreakFrequency
}

func (m *MetricResponse) GetStatus() MetricStatus {
	if m == nil {
		return ""
	}
	return m.Status
}

func (m *MetricResponse) GetCurrent() float64 {
	if m == nil {
		return 0
	}
	return m.Current
}

func (m *MetricResponse) GetAchievements() []*AchievementResponse {
	if m == nil {
		return nil
	}
	return m.Achievements
}

func (m *MetricResponse) GetCurrentStreak() *StreakResponse {
	if m == nil {
		return nil
	}
	return m.CurrentStreak
}

func (m *MetricResponse) GetExtraProperties() map[string]interface{} {
	return m.extraProperties
}

func (m *MetricResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler MetricResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*m = MetricResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *m)
	if err != nil {
		return err
	}
	m.extraProperties = extraProperties
	m.rawJSON = json.RawMessage(data)
	return nil
}

func (m *MetricResponse) String() string {
	if len(m.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(m.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(m); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", m)
}

// The status of the achievement.
type MetricStatus string

const (
	MetricStatusArchived MetricStatus = "archived"
	MetricStatusActive   MetricStatus = "active"
)

func NewMetricStatusFromString(s string) (MetricStatus, error) {
	switch s {
	case "archived":
		return MetricStatusArchived, nil
	case "active":
		return MetricStatusActive, nil
	}
	var t MetricStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (m MetricStatus) Ptr() *MetricStatus {
	return &m
}

// A user of your application.
type User struct {
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
	// Whether the user is in the control group, meaning they do not receive emails or other communications from Trophy.
	Control *bool `json:"control,omitempty" url:"control,omitempty"`
	// The date and time the user was created, in ISO 8601 format.
	Created *time.Time `json:"created,omitempty" url:"created,omitempty"`
	// The date and time the user was last updated, in ISO 8601 format.
	Updated *time.Time `json:"updated,omitempty" url:"updated,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (u *User) GetEmail() *string {
	if u == nil {
		return nil
	}
	return u.Email
}

func (u *User) GetName() *string {
	if u == nil {
		return nil
	}
	return u.Name
}

func (u *User) GetTz() *string {
	if u == nil {
		return nil
	}
	return u.Tz
}

func (u *User) GetSubscribeToEmails() *bool {
	if u == nil {
		return nil
	}
	return u.SubscribeToEmails
}

func (u *User) GetId() string {
	if u == nil {
		return ""
	}
	return u.Id
}

func (u *User) GetControl() *bool {
	if u == nil {
		return nil
	}
	return u.Control
}

func (u *User) GetCreated() *time.Time {
	if u == nil {
		return nil
	}
	return u.Created
}

func (u *User) GetUpdated() *time.Time {
	if u == nil {
		return nil
	}
	return u.Updated
}

func (u *User) GetExtraProperties() map[string]interface{} {
	return u.extraProperties
}

func (u *User) UnmarshalJSON(data []byte) error {
	type embed User
	var unmarshaler = struct {
		embed
		Created *internal.DateTime `json:"created,omitempty"`
		Updated *internal.DateTime `json:"updated,omitempty"`
	}{
		embed: embed(*u),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*u = User(unmarshaler.embed)
	u.Created = unmarshaler.Created.TimePtr()
	u.Updated = unmarshaler.Updated.TimePtr()
	extraProperties, err := internal.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	u.rawJSON = json.RawMessage(data)
	return nil
}

func (u *User) MarshalJSON() ([]byte, error) {
	type embed User
	var marshaler = struct {
		embed
		Created *internal.DateTime `json:"created,omitempty"`
		Updated *internal.DateTime `json:"updated,omitempty"`
	}{
		embed:   embed(*u),
		Created: internal.NewOptionalDateTime(u.Created),
		Updated: internal.NewOptionalDateTime(u.Updated),
	}
	return json.Marshal(marshaler)
}

func (u *User) String() string {
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
