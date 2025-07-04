// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/trophyso/trophy-go/internal"
	time "time"
)

type PointsRange struct {
	// The start of the points range. Inclusive.
	From *float64 `json:"from,omitempty" url:"from,omitempty"`
	// The end of the points range. Inclusive.
	To *float64 `json:"to,omitempty" url:"to,omitempty"`
	// The number of users in this points range.
	Users *float64 `json:"users,omitempty" url:"users,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (p *PointsRange) GetFrom() *float64 {
	if p == nil {
		return nil
	}
	return p.From
}

func (p *PointsRange) GetTo() *float64 {
	if p == nil {
		return nil
	}
	return p.To
}

func (p *PointsRange) GetUsers() *float64 {
	if p == nil {
		return nil
	}
	return p.Users
}

func (p *PointsRange) GetExtraProperties() map[string]interface{} {
	return p.extraProperties
}

func (p *PointsRange) UnmarshalJSON(data []byte) error {
	type unmarshaler PointsRange
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PointsRange(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	p.rawJSON = json.RawMessage(data)
	return nil
}

func (p *PointsRange) String() string {
	if len(p.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(p.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(p); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", p)
}

// A list of eleven points ranges, with the first starting and ending at 0, and the remaining 10 being calculated as 10 equally sized ranges from 1 to the greatest number of points a user has, rounded up to the nearest power of 10.
type PointsSummaryResponse = []*PointsRange

type PointsTriggerResponse struct {
	// The unique ID of the trigger.
	Id *string `json:"id,omitempty" url:"id,omitempty"`
	// The type of trigger.
	Type *PointsTriggerResponseType `json:"type,omitempty" url:"type,omitempty"`
	// The points awarded by this trigger.
	Points *float64 `json:"points,omitempty" url:"points,omitempty"`
	// The status of the trigger.
	Status *PointsTriggerResponseStatus `json:"status,omitempty" url:"status,omitempty"`
	// The unique ID of the achievement associated with this trigger, if the trigger is an achievement.
	AchievementId *string `json:"achievementId,omitempty" url:"achievementId,omitempty"`
	// The unique ID of the metric associated with this trigger, if the trigger is a metric.
	MetricId *string `json:"metricId,omitempty" url:"metricId,omitempty"`
	// The amount that a user must increase the metric to earn the points, if the trigger is a metric.
	MetricThreshold *float64 `json:"metricThreshold,omitempty" url:"metricThreshold,omitempty"`
	// The number of consecutive streak periods that a user must complete to earn the points, if the trigger is a streak.
	StreakLengthThreshold *float64 `json:"streakLengthThreshold,omitempty" url:"streakLengthThreshold,omitempty"`
	// The name of the metric associated with this trigger, if the trigger is a metric.
	MetricName *string `json:"metricName,omitempty" url:"metricName,omitempty"`
	// The name of the achievement associated with this trigger, if the trigger is an achievement.
	AchievementName *string `json:"achievementName,omitempty" url:"achievementName,omitempty"`
	// The date and time the trigger was created, in ISO 8601 format.
	Created *time.Time `json:"created,omitempty" url:"created,omitempty"`
	// The date and time the trigger was last updated, in ISO 8601 format.
	Updated *time.Time `json:"updated,omitempty" url:"updated,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (p *PointsTriggerResponse) GetId() *string {
	if p == nil {
		return nil
	}
	return p.Id
}

func (p *PointsTriggerResponse) GetType() *PointsTriggerResponseType {
	if p == nil {
		return nil
	}
	return p.Type
}

func (p *PointsTriggerResponse) GetPoints() *float64 {
	if p == nil {
		return nil
	}
	return p.Points
}

func (p *PointsTriggerResponse) GetStatus() *PointsTriggerResponseStatus {
	if p == nil {
		return nil
	}
	return p.Status
}

func (p *PointsTriggerResponse) GetAchievementId() *string {
	if p == nil {
		return nil
	}
	return p.AchievementId
}

func (p *PointsTriggerResponse) GetMetricId() *string {
	if p == nil {
		return nil
	}
	return p.MetricId
}

func (p *PointsTriggerResponse) GetMetricThreshold() *float64 {
	if p == nil {
		return nil
	}
	return p.MetricThreshold
}

func (p *PointsTriggerResponse) GetStreakLengthThreshold() *float64 {
	if p == nil {
		return nil
	}
	return p.StreakLengthThreshold
}

func (p *PointsTriggerResponse) GetMetricName() *string {
	if p == nil {
		return nil
	}
	return p.MetricName
}

func (p *PointsTriggerResponse) GetAchievementName() *string {
	if p == nil {
		return nil
	}
	return p.AchievementName
}

func (p *PointsTriggerResponse) GetCreated() *time.Time {
	if p == nil {
		return nil
	}
	return p.Created
}

func (p *PointsTriggerResponse) GetUpdated() *time.Time {
	if p == nil {
		return nil
	}
	return p.Updated
}

func (p *PointsTriggerResponse) GetExtraProperties() map[string]interface{} {
	return p.extraProperties
}

func (p *PointsTriggerResponse) UnmarshalJSON(data []byte) error {
	type embed PointsTriggerResponse
	var unmarshaler = struct {
		embed
		Created *internal.DateTime `json:"created,omitempty"`
		Updated *internal.DateTime `json:"updated,omitempty"`
	}{
		embed: embed(*p),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*p = PointsTriggerResponse(unmarshaler.embed)
	p.Created = unmarshaler.Created.TimePtr()
	p.Updated = unmarshaler.Updated.TimePtr()
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	p.rawJSON = json.RawMessage(data)
	return nil
}

func (p *PointsTriggerResponse) MarshalJSON() ([]byte, error) {
	type embed PointsTriggerResponse
	var marshaler = struct {
		embed
		Created *internal.DateTime `json:"created,omitempty"`
		Updated *internal.DateTime `json:"updated,omitempty"`
	}{
		embed:   embed(*p),
		Created: internal.NewOptionalDateTime(p.Created),
		Updated: internal.NewOptionalDateTime(p.Updated),
	}
	return json.Marshal(marshaler)
}

func (p *PointsTriggerResponse) String() string {
	if len(p.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(p.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(p); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", p)
}

// The status of the trigger.
type PointsTriggerResponseStatus string

const (
	PointsTriggerResponseStatusActive   PointsTriggerResponseStatus = "active"
	PointsTriggerResponseStatusArchived PointsTriggerResponseStatus = "archived"
)

func NewPointsTriggerResponseStatusFromString(s string) (PointsTriggerResponseStatus, error) {
	switch s {
	case "active":
		return PointsTriggerResponseStatusActive, nil
	case "archived":
		return PointsTriggerResponseStatusArchived, nil
	}
	var t PointsTriggerResponseStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (p PointsTriggerResponseStatus) Ptr() *PointsTriggerResponseStatus {
	return &p
}

// The type of trigger.
type PointsTriggerResponseType string

const (
	PointsTriggerResponseTypeMetric      PointsTriggerResponseType = "metric"
	PointsTriggerResponseTypeAchievement PointsTriggerResponseType = "achievement"
	PointsTriggerResponseTypeStreak      PointsTriggerResponseType = "streak"
)

func NewPointsTriggerResponseTypeFromString(s string) (PointsTriggerResponseType, error) {
	switch s {
	case "metric":
		return PointsTriggerResponseTypeMetric, nil
	case "achievement":
		return PointsTriggerResponseTypeAchievement, nil
	case "streak":
		return PointsTriggerResponseTypeStreak, nil
	}
	var t PointsTriggerResponseType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (p PointsTriggerResponseType) Ptr() *PointsTriggerResponseType {
	return &p
}
