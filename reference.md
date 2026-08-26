# Reference
## Achievements
<details><summary><code>client.Achievements.All() -> []*trophygo.AchievementWithStatsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get all achievements and their completion stats.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &trophygo.AchievementsAllRequest{
        UserAttributes: trophygo.String(
            "plan-type:premium,region:us-east",
        ),
    }
client.Achievements.All(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userAttributes:** `*string` — Optional colon-delimited user attributes in the format attribute:value,attribute:value. Only achievements accessible to a user with the provided attributes will be returned.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Achievements.Complete(Key, request) -> *trophygo.AchievementCompletionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Mark an achievement as completed for a user.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &trophygo.AchievementsCompleteRequest{
        User: &trophygo.UpsertedUser{
            Email: trophygo.String(
                "user@example.com",
            ),
            Tz: trophygo.String(
                "Europe/London",
            ),
            Id: "user-id",
        },
    }
client.Achievements.Complete(
        context.TODO(),
        "finish-onboarding",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**key:** `string` — Unique reference of the achievement as set when created.
    
</dd>
</dl>

<dl>
<dd>

**user:** `*trophygo.UpsertedUser` — The user that completed the achievement.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Metrics
<details><summary><code>client.Metrics.Event(Key, request) -> *trophygo.EventResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Increment or decrement the value of a metric for a user.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &trophygo.MetricsEventRequest{
        IdempotencyKey: trophygo.String(
            "e4296e4b-8493-4bd1-9c30-5a1a9ac4d78f",
        ),
        User: &trophygo.UpsertedUser{
            Email: trophygo.String(
                "user@example.com",
            ),
            Tz: trophygo.String(
                "Europe/London",
            ),
            Attributes: map[string]string{
                "department": "engineering",
                "role": "developer",
            },
            Id: "18",
        },
        Value: 750,
        Attributes: map[string]string{
            "category": "writing",
            "source": "mobile-app",
        },
    }
client.Metrics.Event(
        context.TODO(),
        "words-written",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**key:** `string` — Unique reference of the metric as set when created.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `*string` — The idempotency key for the event.
    
</dd>
</dl>

<dl>
<dd>

**user:** `*trophygo.UpsertedUser` — The user that triggered the event.
    
</dd>
</dl>

<dl>
<dd>

**value:** `float64` — The value to add to the user's current total for the given metric.
    
</dd>
</dl>

<dl>
<dd>

**attributes:** `map[string]string` — Event attributes as key-value pairs. Keys must match existing event attributes set up in the Trophy dashboard.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Users
<details><summary><code>client.Users.Create(request) -> *trophygo.User</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a new user.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &trophygo.UpsertedUser{
        Id: "user-id",
    }
client.Users.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*trophygo.UpsertedUser` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.Get(Id) -> *trophygo.User</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a single user.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Users.Get(
        context.TODO(),
        "userId",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the user to get.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.Identify(Id, request) -> *trophygo.User</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Identify a user.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &trophygo.UpdatedUser{
        Email: trophygo.String(
            "user@example.com",
        ),
        Tz: trophygo.String(
            "Europe/London",
        ),
        Attributes: map[string]string{
            "department": "engineering",
            "role": "developer",
        },
    }
client.Users.Identify(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the user to identify.
    
</dd>
</dl>

<dl>
<dd>

**request:** `*trophygo.UpdatedUser` — The user object.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.Update(Id, request) -> *trophygo.User</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update a user.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &trophygo.UpdatedUser{
        Email: trophygo.String(
            "user@example.com",
        ),
        Tz: trophygo.String(
            "Europe/London",
        ),
        Attributes: map[string]string{
            "department": "engineering",
            "role": "developer",
        },
    }
client.Users.Update(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the user to update.
    
</dd>
</dl>

<dl>
<dd>

**request:** `*trophygo.UpdatedUser` — The user object.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.GetPreferences(Id) -> *trophygo.UserPreferencesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a user's notification preferences.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Users.GetPreferences(
        context.TODO(),
        "user-123",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The user's ID in your database.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.UpdatePreferences(Id, request) -> *trophygo.UserPreferencesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update a user's notification and streak preferences. Streak preferences other than `streak.enabled` require streak customization to be enabled in your Trophy dashboard settings.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &trophygo.UpdateUserPreferencesRequest{
        Notifications: &trophygo.NotificationPreferences{
            StreakReminder: []trophygo.NotificationChannel{
                trophygo.NotificationChannelEmail,
            },
        },
    }
client.Users.UpdatePreferences(
        context.TODO(),
        "user-123",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The user's ID in your database.
    
</dd>
</dl>

<dl>
<dd>

**notifications:** `*trophygo.NotificationPreferences` 
    
</dd>
</dl>

<dl>
<dd>

**streak:** `*trophygo.StreakPreferences` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.AllMetrics(Id) -> []*trophygo.MetricResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a single user's progress against all active metrics.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Users.AllMetrics(
        context.TODO(),
        "userId",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the user
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.SingleMetric(Id, Key) -> *trophygo.MetricResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a user's progress against a single active metric.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Users.SingleMetric(
        context.TODO(),
        "userId",
        "key",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the user.
    
</dd>
</dl>

<dl>
<dd>

**key:** `string` — Unique key of the metric.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.MetricEventSummary(Id, Key) -> []*trophygo.UsersMetricEventSummaryResponseItem</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a summary of metric events over time for a user.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &trophygo.UsersMetricEventSummaryRequest{
        Aggregation: trophygo.UsersMetricEventSummaryRequestAggregationDaily,
        StartDate: "2024-01-01",
        EndDate: "2024-01-31",
    }
client.Users.MetricEventSummary(
        context.TODO(),
        "userId",
        "words-written",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the user.
    
</dd>
</dl>

<dl>
<dd>

**key:** `string` — Unique key of the metric.
    
</dd>
</dl>

<dl>
<dd>

**aggregation:** `*trophygo.UsersMetricEventSummaryRequestAggregation` — The time period over which to aggregate the event data.
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — The start date for the data range in YYYY-MM-DD format. The startDate must be before the endDate, and the date range must not exceed 400 days.
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `string` — The end date for the data range in YYYY-MM-DD format. The endDate must be after the startDate, and the date range must not exceed 400 days.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.Achievements(Id) -> []*trophygo.UserAchievementWithStatsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a user's achievements.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &trophygo.UsersAchievementsRequest{}
client.Users.Achievements(
        context.TODO(),
        "userId",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the user.
    
</dd>
</dl>

<dl>
<dd>

**includeIncomplete:** `*string` — When set to 'true', returns both completed and incomplete achievements for the user. When omitted or set to any other value, returns only completed achievements.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.Streak(Id) -> *trophygo.StreakResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a user's streak data.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &trophygo.UsersStreakRequest{
        HistoryPeriods: trophygo.Int(
            1,
        ),
    }
client.Users.Streak(
        context.TODO(),
        "userId",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the user.
    
</dd>
</dl>

<dl>
<dd>

**historyPeriods:** `*int` — The number of past streak periods to include in the streakHistory field of the  response.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.Points(Id, Key) -> *trophygo.GetUserPointsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a user's points for a specific points system.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &trophygo.UsersPointsRequest{
        Awards: trophygo.Int(
            1,
        ),
    }
client.Users.Points(
        context.TODO(),
        "userId",
        "points-system-key",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the user.
    
</dd>
</dl>

<dl>
<dd>

**key:** `string` — Key of the points system.
    
</dd>
</dl>

<dl>
<dd>

**awards:** `*int` — The number of recent point awards to return.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.PointsBoosts(Id, Key) -> []*trophygo.PointsBoost</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get active points boosts for a user in a specific points system. Returns both global boosts the user is eligible for and user-specific boosts.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Users.PointsBoosts(
        context.TODO(),
        "userId",
        "points-system-key",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the user.
    
</dd>
</dl>

<dl>
<dd>

**key:** `string` — Key of the points system.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.PointsEventSummary(Id, Key) -> []*trophygo.UsersPointsEventSummaryResponseItem</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a summary of points awards over time for a user for a specific points system.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &trophygo.UsersPointsEventSummaryRequest{
        Aggregation: trophygo.UsersPointsEventSummaryRequestAggregationDaily,
        StartDate: "2024-01-01",
        EndDate: "2024-01-31",
    }
client.Users.PointsEventSummary(
        context.TODO(),
        "userId",
        "points-system-key",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the user.
    
</dd>
</dl>

<dl>
<dd>

**key:** `string` — Key of the points system.
    
</dd>
</dl>

<dl>
<dd>

**aggregation:** `*trophygo.UsersPointsEventSummaryRequestAggregation` — The time period over which to aggregate the event data.
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `string` — The start date for the data range in YYYY-MM-DD format. The startDate must be before the endDate, and the date range must not exceed 400 days.
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `string` — The end date for the data range in YYYY-MM-DD format. The endDate must be after the startDate, and the date range must not exceed 400 days.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.Leaderboard(Id, Key) -> *trophygo.UserLeaderboardResponseWithHistory</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a user's rank, value, and daily ranking history for a specific leaderboard.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &trophygo.UsersLeaderboardRequest{
        Run: trophygo.String(
            "2025-01-15",
        ),
        NumEvents: trophygo.Int(
            1,
        ),
    }
client.Users.Leaderboard(
        context.TODO(),
        "user-123",
        "weekly-words",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The user's ID in your database.
    
</dd>
</dl>

<dl>
<dd>

**key:** `string` — Unique key of the leaderboard as set when created.
    
</dd>
</dl>

<dl>
<dd>

**run:** `*string` — Specific run date in YYYY-MM-DD format. If not provided, returns the current run.
    
</dd>
</dl>

<dl>
<dd>

**numEvents:** `*int` — The number of days to return in the leaderboard history for the user.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.Wrapped(Id) -> *trophygo.WrappedResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a user's year-in-review wrapped data.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &trophygo.UsersWrappedRequest{
        Year: trophygo.Int(
            1,
        ),
    }
client.Users.Wrapped(
        context.TODO(),
        "user-123",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The user's ID in your database.
    
</dd>
</dl>

<dl>
<dd>

**year:** `*int` — The year to get wrapped data for. Defaults to the current year. Must be an integer between 1 and the current year.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Streaks
<details><summary><code>client.Streaks.List() -> trophygo.BulkStreakResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get the streak lengths of a list of users, ranked by streak length from longest to shortest.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &trophygo.StreaksListRequest{
        UserIds: []*string{
            trophygo.String(
                "userIds",
            ),
        },
    }
client.Streaks.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userIds:** `*string` — A list of up to 100 user IDs.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Points
<details><summary><code>client.Points.Summary(Key) -> trophygo.PointsSummaryResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a breakdown of the number of users with points in each range.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &trophygo.PointsSummaryRequest{
        UserAttributes: trophygo.String(
            "plan-type:premium,region:us-east",
        ),
    }
client.Points.Summary(
        context.TODO(),
        "points-system-key",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**key:** `string` — Key of the points system.
    
</dd>
</dl>

<dl>
<dd>

**userAttributes:** `*string` — Optional colon-delimited user attribute filters in the format attribute:value,attribute:value. Only users matching ALL specified attributes will be included in the points breakdown.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Points.System(Key) -> *trophygo.PointsSystemResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a points system with its triggers.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Points.System(
        context.TODO(),
        "points-system-key",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**key:** `string` — Key of the points system.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Points.Boosts(Key) -> []*trophygo.PointsBoost</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get all global boosts for a points system. Finished boosts are excluded by default.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &trophygo.PointsBoostsRequest{
        IncludeFinished: trophygo.Bool(
            true,
        ),
    }
client.Points.Boosts(
        context.TODO(),
        "points-system-key",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**key:** `string` — Key of the points system.
    
</dd>
</dl>

<dl>
<dd>

**includeFinished:** `*bool` — When set to 'true', boosts that have finished (past their end date) will be included in the response. By default, finished boosts are excluded.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Points.Levels(Key) -> []*trophygo.PointsLevel</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get all levels for a points system.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Points.Levels(
        context.TODO(),
        "points-system-key",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**key:** `string` — Key of the points system.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Points.LevelSummary(Key) -> trophygo.PointsLevelSummaryResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a breakdown of the number of users at each level in a points system.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Points.LevelSummary(
        context.TODO(),
        "points-system-key",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**key:** `string` — Key of the points system.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Leaderboards
<details><summary><code>client.Leaderboards.All() -> []*trophygo.LeaderboardsAllResponseItem</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get all leaderboards for your organization. Finished leaderboards are excluded by default.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &trophygo.LeaderboardsAllRequest{
        IncludeFinished: trophygo.Bool(
            true,
        ),
    }
client.Leaderboards.All(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**includeFinished:** `*bool` — When set to 'true', leaderboards with status 'finished' will be included in the response. By default, finished leaderboards are excluded.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Leaderboards.Get(Key) -> *trophygo.LeaderboardResponseWithRankings</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a specific leaderboard by its key.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &trophygo.LeaderboardsGetRequest{
        Offset: trophygo.Int(
            1,
        ),
        Limit: trophygo.Int(
            1,
        ),
        Run: trophygo.String(
            "2025-01-15",
        ),
        UserId: trophygo.String(
            "user-123",
        ),
        UserAttributes: trophygo.String(
            "city:London",
        ),
    }
client.Leaderboards.Get(
        context.TODO(),
        "weekly-words",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**key:** `string` — Unique key of the leaderboard as set when created.
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Number of rankings to skip for pagination.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Maximum number of rankings to return. Cannot be greater than the size of the leaderboard.
    
</dd>
</dl>

<dl>
<dd>

**run:** `*string` — Specific run date in YYYY-MM-DD format. If not provided, returns the current run.
    
</dd>
</dl>

<dl>
<dd>

**userId:** `*string` — When provided, offset is relative to this user's position on the leaderboard. If the user is not found in the leaderboard, returns empty rankings array.
    
</dd>
</dl>

<dl>
<dd>

**userAttributes:** `*string` — Attribute key and value to filter the rankings by, separated by a colon. For example, `city:London`. This parameter is required, and only valid for leaderboards with a breakdown attribute.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Admin Attributes
<details><summary><code>client.Admin.Attributes.List() -> trophygo.ListAttributesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List attributes.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &admin.AttributesListRequest{
        Limit: trophygo.Int(
            1,
        ),
        Skip: trophygo.Int(
            1,
        ),
    }
client.Admin.Attributes.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**limit:** `*int` — Number of records to return.
    
</dd>
</dl>

<dl>
<dd>

**skip:** `*int` — Number of records to skip from the start of the list.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Admin.Attributes.Create(request) -> *trophygo.CreateAttributesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create attributes.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := []*trophygo.CreateAttributeRequestItem{
        &trophygo.CreateAttributeRequestItem{
            Name: "Plan",
            Key: "plan",
            Type: trophygo.CreateAttributeRequestItemTypeUser,
        },
        &trophygo.CreateAttributeRequestItem{
            Name: "Device",
            Key: "device",
            Type: trophygo.CreateAttributeRequestItemTypeEvent,
        },
    }
client.Admin.Attributes.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `trophygo.CreateAttributesRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Admin.Attributes.Delete() -> *trophygo.DeleteAttributesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete attributes by ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &admin.AttributesDeleteRequest{
        Ids: []*string{
            trophygo.String(
                "550e8400-e29b-41d4-a716-446655440000",
            ),
            trophygo.String(
                "550e8400-e29b-41d4-a716-446655440001",
            ),
        },
    }
client.Admin.Attributes.Delete(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**ids:** `*string` — Attribute IDs to delete. Repeat the query param or provide a comma-separated list.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Admin.Attributes.Update(request) -> *trophygo.UpdateAttributesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update attributes by ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := []*trophygo.UpdateAttributeRequestItem{
        &trophygo.UpdateAttributeRequestItem{
            Id: "550e8400-e29b-41d4-a716-446655440000",
            Name: trophygo.String(
                "Subscription Plan",
            ),
        },
    }
client.Admin.Attributes.Update(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `trophygo.UpdateAttributesRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Admin.Attributes.Get(Id) -> *trophygo.AdminAttribute</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get an attribute by ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Admin.Attributes.Get(
        context.TODO(),
        "550e8400-e29b-41d4-a716-446655440000",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The UUID of the attribute to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Admin Metrics
<details><summary><code>client.Admin.Metrics.List() -> trophygo.ListMetricsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List metrics.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &admin.MetricsListRequest{
        Limit: trophygo.Int(
            1,
        ),
        Skip: trophygo.Int(
            1,
        ),
    }
client.Admin.Metrics.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**limit:** `*int` — Number of records to return.
    
</dd>
</dl>

<dl>
<dd>

**skip:** `*int` — Number of records to skip from the start of the list.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Admin.Metrics.Create(request) -> *trophygo.CreateMetricsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create metrics.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := []*trophygo.CreateMetricRequestItem{
        &trophygo.CreateMetricRequestItem{
            Name: "Invites Sent",
            Key: "invites-sent",
        },
        &trophygo.CreateMetricRequestItem{
            Name: "Revenue",
            Key: "revenue",
            UnitType: trophygo.CreateMetricRequestItemUnitTypeCurrency.Ptr(),
            Units: trophygo.String(
                "USD",
            ),
        },
    }
client.Admin.Metrics.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `trophygo.CreateMetricsRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Admin.Metrics.Delete() -> *trophygo.DeleteMetricsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete metrics by ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &admin.MetricsDeleteRequest{
        Ids: []*string{
            trophygo.String(
                "550e8400-e29b-41d4-a716-446655440000",
            ),
            trophygo.String(
                "550e8400-e29b-41d4-a716-446655440001",
            ),
        },
    }
client.Admin.Metrics.Delete(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**ids:** `*string` — Metric IDs to delete. Repeat the query param or provide a comma-separated list.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Admin.Metrics.Update(request) -> *trophygo.UpdateMetricsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update metrics by ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := []*trophygo.UpdateMetricRequestItem{
        &trophygo.UpdateMetricRequestItem{
            Id: "550e8400-e29b-41d4-a716-446655440000",
            Name: trophygo.String(
                "Invites Completed",
            ),
            Units: trophygo.String(
                "invites",
            ),
        },
        &trophygo.UpdateMetricRequestItem{
            Id: "550e8400-e29b-41d4-a716-446655440001",
            UnitType: trophygo.UpdateMetricRequestItemUnitTypeNumber.Ptr(),
            Units: trophygo.String(
                "dollars",
            ),
        },
    }
client.Admin.Metrics.Update(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `trophygo.UpdateMetricsRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Admin.Metrics.Get(Id) -> *trophygo.CreatedMetric</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a metric by ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Admin.Metrics.Get(
        context.TODO(),
        "550e8400-e29b-41d4-a716-446655440000",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The UUID of the metric to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Admin.Metrics.BatchEvents(request) -> *trophygo.BatchEventsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Submit up to 1,000 metric events for asynchronous processing.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := []*trophygo.BatchMetricEvent{
        &trophygo.BatchMetricEvent{
            Key: "words-written",
            User: &trophygo.BatchMetricEventUser{
                Id: "18",
                Email: trophygo.String(
                    "user@example.com",
                ),
                Tz: trophygo.String(
                    "Europe/London",
                ),
                Attributes: map[string]string{
                    "department": "engineering",
                    "role": "developer",
                },
            },
            Value: 750,
            Attributes: map[string]string{
                "category": "writing",
                "source": "mobile-app",
            },
            IdempotencyKey: trophygo.String(
                "e4296e4b-8493-4bd1-9c30-5a1a9ac4d78f",
            ),
        },
    }
client.Admin.Metrics.BatchEvents(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `[]*trophygo.BatchMetricEvent` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Admin Leaderboards
<details><summary><code>client.Admin.Leaderboards.List() -> trophygo.ListLeaderboardsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List leaderboards.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &admin.LeaderboardsListRequest{
        Limit: trophygo.Int(
            1,
        ),
        Skip: trophygo.Int(
            1,
        ),
    }
client.Admin.Leaderboards.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**limit:** `*int` — Number of records to return.
    
</dd>
</dl>

<dl>
<dd>

**skip:** `*int` — Number of records to skip from the start of the list.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Admin.Leaderboards.Create(request) -> *trophygo.CreateLeaderboardsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create leaderboards. Maximum 100 leaderboards per request.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := []*trophygo.CreateLeaderboardRequestItem{
        &trophygo.CreateLeaderboardRequestItem{
            Name: "Revenue Champions",
            Key: "revenue-champions",
            Status: trophygo.CreateLeaderboardRequestItemStatusInactive.Ptr(),
            RankBy: trophygo.CreateLeaderboardRequestItemRankByMetric,
            MetricId: trophygo.String(
                "550e8400-e29b-41d4-a716-446655440000",
            ),
            MaxParticipants: trophygo.Int(
                100,
            ),
            Start: trophygo.String(
                "2026-04-20",
            ),
            BreakdownAttributes: []string{
                "550e8400-e29b-41d4-a716-446655440010",
            },
            RunUnit: trophygo.CreateLeaderboardRequestItemRunUnitMonth.Ptr(),
            RunInterval: trophygo.Int(
                1,
            ),
        },
        &trophygo.CreateLeaderboardRequestItem{
            Name: "Streak Legends",
            Key: "streak-legends",
            Status: trophygo.CreateLeaderboardRequestItemStatusScheduled.Ptr(),
            RankBy: trophygo.CreateLeaderboardRequestItemRankByStreak,
            Start: trophygo.String(
                "2026-04-27",
            ),
        },
    }
client.Admin.Leaderboards.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `trophygo.CreateLeaderboardsRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Admin.Leaderboards.Delete() -> *trophygo.DeleteLeaderboardsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete leaderboards by ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &admin.LeaderboardsDeleteRequest{
        Ids: []*string{
            trophygo.String(
                "ids",
            ),
        },
    }
client.Admin.Leaderboards.Delete(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**ids:** `*string` — Leaderboard IDs to delete. Repeat the query param or provide a comma-separated list.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Admin.Leaderboards.Update(request) -> *trophygo.UpdateLeaderboardsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update leaderboards by ID. Updating `status` behaves the same as activating, scheduling, deactivating, or finishing a leaderboard in the dashboard.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := []*trophygo.UpdateLeaderboardRequestItem{
        &trophygo.UpdateLeaderboardRequestItem{
            Id: "550e8400-e29b-41d4-a716-446655440100",
            Name: trophygo.String(
                "Monthly Revenue Champions",
            ),
            Description: trophygo.String(
                "Ranked by monthly revenue",
            ),
            Status: trophygo.UpdateLeaderboardRequestItemStatusActive.Ptr(),
        },
        &trophygo.UpdateLeaderboardRequestItem{
            Id: "550e8400-e29b-41d4-a716-446655440101",
            Status: trophygo.UpdateLeaderboardRequestItemStatusFinished.Ptr(),
        },
    }
client.Admin.Leaderboards.Update(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `trophygo.UpdateLeaderboardsRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Admin.Leaderboards.Get(Id) -> *trophygo.AdminLeaderboard</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a leaderboard by ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Admin.Leaderboards.Get(
        context.TODO(),
        "550e8400-e29b-41d4-a716-446655440100",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The UUID of the leaderboard to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Admin Streaks
<details><summary><code>client.Admin.Streaks.Restore(request) -> *trophygo.RestoreStreaksResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Restore streaks for multiple users to the maximum previously achieved streak length found within the current restore window: the last 90 days for daily streaks, weekly periods starting with the week containing the start of the current calendar year for weekly streaks, and monthly periods starting at the beginning of the previous calendar year for monthly streaks.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &admin.RestoreStreaksRequest{
        Users: []*admin.RestoreStreaksRequestUsersItem{
            &admin.RestoreStreaksRequestUsersItem{
                Id: "user-123",
            },
            &admin.RestoreStreaksRequestUsersItem{
                Id: "user-456",
            },
        },
    }
client.Admin.Streaks.Restore(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**users:** `[]*admin.RestoreStreaksRequestUsersItem` — Array of users to restore streaks for. Maximum 100 users per request.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Admin ApplicationApiKeys
<details><summary><code>client.Admin.ApplicationApiKeys.Create(request) -> *trophygo.CreateApplicationKeysResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create application API keys scoped to specific users. Each key can only perform operations on behalf of the user it was created for.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := []*trophygo.CreateApplicationKeyRequestItem{
        &trophygo.CreateApplicationKeyRequestItem{
            UserId: "user_123",
        },
        &trophygo.CreateApplicationKeyRequestItem{
            UserId: "user_456",
        },
    }
client.Admin.ApplicationApiKeys.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `trophygo.CreateApplicationKeysRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Admin.ApplicationApiKeys.Delete() -> *trophygo.DeleteApplicationKeysResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete application API keys by ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &admin.ApplicationApiKeysDeleteRequest{
        Ids: []*string{
            trophygo.String(
                "550e8400-e29b-41d4-a716-446655440000",
            ),
        },
    }
client.Admin.ApplicationApiKeys.Delete(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**ids:** `*string` — Application API key IDs (UUIDs returned at creation time). Repeat the query param or provide a comma-separated list. Maximum 100 IDs per request.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Admin Tenants
<details><summary><code>client.Admin.Tenants.List() -> trophygo.ListTenantsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List tenants in the current environment.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &admin.TenantsListRequest{
        Limit: trophygo.Int(
            1,
        ),
        Skip: trophygo.Int(
            1,
        ),
    }
client.Admin.Tenants.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**limit:** `*int` — Number of records to return.
    
</dd>
</dl>

<dl>
<dd>

**skip:** `*int` — Number of records to skip from the start of the list.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Admin.Tenants.Create(request) -> *trophygo.CreateTenantsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create tenants.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := []*trophygo.CreateTenantRequestItem{
        &trophygo.CreateTenantRequestItem{
            CustomerId: "customer_12345",
            Name: "Acme Corp",
        },
        &trophygo.CreateTenantRequestItem{
            CustomerId: "customer_67890",
            Name: "Globex Inc",
        },
    }
client.Admin.Tenants.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `trophygo.CreateTenantsRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Admin.Tenants.Delete() -> *trophygo.DeleteTenantsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete tenants by ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &admin.TenantsDeleteRequest{
        Ids: []*string{
            trophygo.String(
                "550e8400-e29b-41d4-a716-446655440000",
            ),
            trophygo.String(
                "550e8400-e29b-41d4-a716-446655440001",
            ),
        },
    }
client.Admin.Tenants.Delete(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**ids:** `*string` — Tenant IDs to delete. Repeat the query param or provide a comma-separated list.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Admin.Tenants.Update(request) -> *trophygo.UpdateTenantsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update tenants by ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := []*trophygo.UpdateTenantRequestItem{
        &trophygo.UpdateTenantRequestItem{
            Id: "550e8400-e29b-41d4-a716-446655440000",
            Name: trophygo.String(
                "Acme Corporation",
            ),
        },
    }
client.Admin.Tenants.Update(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `trophygo.UpdateTenantsRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Admin.Tenants.Get(Id) -> *trophygo.AdminTenant</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a tenant by ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Admin.Tenants.Get(
        context.TODO(),
        "550e8400-e29b-41d4-a716-446655440000",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The UUID of the tenant to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Admin Points Systems
<details><summary><code>client.Admin.Points.Systems.List() -> trophygo.ListPointsSystemsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List points systems.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &points.SystemsListRequest{
        Limit: trophygo.Int(
            1,
        ),
        Skip: trophygo.Int(
            1,
        ),
    }
client.Admin.Points.Systems.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**limit:** `*int` — Number of records to return.
    
</dd>
</dl>

<dl>
<dd>

**skip:** `*int` — Number of records to skip from the start of the list.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Admin.Points.Systems.Create(request) -> *trophygo.CreatePointsSystemsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create points systems. Optionally include sub-entities (levels, boosts, triggers) in each system payload to create them alongside the system.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := []*trophygo.CreatePointsSystemRequestItem{
        &trophygo.CreatePointsSystemRequestItem{
            Name: "XP",
            Key: "xp",
            Description: trophygo.String(
                "Experience points",
            ),
            Levels: []*trophygo.CreatePointsLevelRequestItem{
                &trophygo.CreatePointsLevelRequestItem{
                    Name: "Bronze",
                    Key: "bronze",
                    Points: 100,
                },
                &trophygo.CreatePointsLevelRequestItem{
                    Name: "Silver",
                    Key: "silver",
                    Points: 500,
                },
            },
        },
    }
client.Admin.Points.Systems.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `trophygo.CreatePointsSystemsRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Admin.Points.Systems.Delete() -> *trophygo.DeletePointsSystemsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete points systems by ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &points.SystemsDeleteRequest{
        Ids: []*string{
            trophygo.String(
                "550e8400-e29b-41d4-a716-446655440000",
            ),
        },
    }
client.Admin.Points.Systems.Delete(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**ids:** `*string` — The IDs of the points systems to delete.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Admin.Points.Systems.Update(request) -> *trophygo.UpdatePointsSystemsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update points systems by ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := []*trophygo.UpdatePointsSystemRequestItem{
        &trophygo.UpdatePointsSystemRequestItem{
            Id: "550e8400-e29b-41d4-a716-446655440000",
            Name: trophygo.String(
                "New Name",
            ),
        },
    }
client.Admin.Points.Systems.Update(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `trophygo.UpdatePointsSystemsRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Admin.Points.Systems.Get(Id) -> *trophygo.AdminPointsSystem</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a points system by ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Admin.Points.Systems.Get(
        context.TODO(),
        "550e8400-e29b-41d4-a716-446655440000",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The ID of the points system.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Admin Points Boosts
<details><summary><code>client.Admin.Points.Boosts.List(SystemId) -> trophygo.ListPointsBoostsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List points boosts for a system.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &points.BoostsListRequest{
        Limit: trophygo.Int(
            1,
        ),
        Skip: trophygo.Int(
            1,
        ),
    }
client.Admin.Points.Boosts.List(
        context.TODO(),
        "550e8400-e29b-41d4-a716-446655440000",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**systemId:** `string` — The UUID of the points system.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Maximum number of results to return (1-100, default 10).
    
</dd>
</dl>

<dl>
<dd>

**skip:** `*int` — Number of results to skip for pagination (default 0).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Admin.Points.Boosts.Create(SystemId, request) -> *trophygo.CreatePointsBoostsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create points boosts.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := []*trophygo.CreatePointsBoostRequestItem{
        &trophygo.CreatePointsBoostRequestItem{
            UserId: trophygo.String(
                "user-123",
            ),
            Name: "Double XP Weekend",
            Start: "2024-01-01",
            End: trophygo.String(
                "2024-01-03",
            ),
            Multiplier: 2,
        },
    }
client.Admin.Points.Boosts.Create(
        context.TODO(),
        "550e8400-e29b-41d4-a716-446655440000",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**systemId:** `string` — The UUID of the points system.
    
</dd>
</dl>

<dl>
<dd>

**request:** `trophygo.CreatePointsBoostsRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Admin.Points.Boosts.Delete(SystemId) -> *trophygo.DeletePointsBoostsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete multiple points boosts by ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &points.BoostsDeleteRequest{
        Ids: []*string{
            trophygo.String(
                "ids",
            ),
        },
    }
client.Admin.Points.Boosts.Delete(
        context.TODO(),
        "550e8400-e29b-41d4-a716-446655440000",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**systemId:** `string` — The UUID of the points system.
    
</dd>
</dl>

<dl>
<dd>

**ids:** `*string` — A list of up to 100 boost IDs.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Admin.Points.Boosts.Update(SystemId, request) -> *trophygo.PatchPointsBoostsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update multiple points boosts.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := []*trophygo.PatchPointsBoostsRequestItem{
        &trophygo.PatchPointsBoostsRequestItem{
            Id: "550e8400-e29b-41d4-a716-446655440000",
            Name: trophygo.String(
                "Updated Boost Name",
            ),
            Multiplier: trophygo.Float64(
                3,
            ),
        },
    }
client.Admin.Points.Boosts.Update(
        context.TODO(),
        "550e8400-e29b-41d4-a716-446655440000",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**systemId:** `string` — The UUID of the points system.
    
</dd>
</dl>

<dl>
<dd>

**request:** `trophygo.PatchPointsBoostsRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Admin.Points.Boosts.Get(SystemId, Id) -> *trophygo.AdminPointsBoost</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a single points boost by ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Admin.Points.Boosts.Get(
        context.TODO(),
        "550e8400-e29b-41d4-a716-446655440000",
        "660f9500-f30c-42e5-b827-557766550001",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**systemId:** `string` — The UUID of the points system.
    
</dd>
</dl>

<dl>
<dd>

**id:** `string` — The UUID of the points boost.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Admin Points Levels
<details><summary><code>client.Admin.Points.Levels.List(SystemId) -> trophygo.ListPointsLevelsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List points levels for a system.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &points.LevelsListRequest{
        Limit: trophygo.Int(
            1,
        ),
        Skip: trophygo.Int(
            1,
        ),
    }
client.Admin.Points.Levels.List(
        context.TODO(),
        "550e8400-e29b-41d4-a716-446655440000",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**systemId:** `string` — The UUID of the points system.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of records to return.
    
</dd>
</dl>

<dl>
<dd>

**skip:** `*int` — Number of records to skip from the start of the list.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Admin.Points.Levels.Create(SystemId, request) -> *trophygo.CreatePointsLevelsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create points levels. Maximum 100 levels per request.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := []*trophygo.CreatePointsLevelRequestItem{
        &trophygo.CreatePointsLevelRequestItem{
            Name: "Bronze",
            Key: "bronze",
            Points: 100,
        },
    }
client.Admin.Points.Levels.Create(
        context.TODO(),
        "550e8400-e29b-41d4-a716-446655440000",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**systemId:** `string` — The UUID of the points system.
    
</dd>
</dl>

<dl>
<dd>

**request:** `trophygo.CreatePointsLevelsRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Admin.Points.Levels.Delete(SystemId) -> *trophygo.DeletePointsLevelsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete multiple points levels by ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &points.LevelsDeleteRequest{
        Ids: []*string{
            trophygo.String(
                "ids",
            ),
        },
    }
client.Admin.Points.Levels.Delete(
        context.TODO(),
        "550e8400-e29b-41d4-a716-446655440000",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**systemId:** `string` — The UUID of the points system.
    
</dd>
</dl>

<dl>
<dd>

**ids:** `*string` — Comma-separated list of level UUIDs to delete.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Admin.Points.Levels.Update(SystemId, request) -> *trophygo.PatchPointsLevelsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update multiple points levels. Each item must include an ID. `key` cannot be changed.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := []*trophygo.PatchPointsLevelsRequestItem{
        &trophygo.PatchPointsLevelsRequestItem{
            Id: "550e8400-e29b-41d4-a716-446655440000",
        },
    }
client.Admin.Points.Levels.Update(
        context.TODO(),
        "550e8400-e29b-41d4-a716-446655440000",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**systemId:** `string` — The UUID of the points system.
    
</dd>
</dl>

<dl>
<dd>

**request:** `trophygo.PatchPointsLevelsRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Admin.Points.Levels.Get(SystemId, Id) -> *trophygo.AdminPointsLevel</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a single points level by ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Admin.Points.Levels.Get(
        context.TODO(),
        "550e8400-e29b-41d4-a716-446655440000",
        "660f9500-f30c-42e5-b827-557766550001",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**systemId:** `string` — The UUID of the points system.
    
</dd>
</dl>

<dl>
<dd>

**id:** `string` — The UUID of the points level.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Admin Points Triggers
<details><summary><code>client.Admin.Points.Triggers.List(SystemId) -> trophygo.ListPointsTriggersResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List points triggers for a system.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &points.TriggersListRequest{
        Limit: trophygo.Int(
            1,
        ),
        Skip: trophygo.Int(
            1,
        ),
    }
client.Admin.Points.Triggers.List(
        context.TODO(),
        "550e8400-e29b-41d4-a716-446655440000",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**systemId:** `string` — The UUID of the points system.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Maximum number of results to return (1-100, default 10).
    
</dd>
</dl>

<dl>
<dd>

**skip:** `*int` — Number of results to skip for pagination (default 0).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Admin.Points.Triggers.Create(SystemId, request) -> *trophygo.CreatePointsTriggersResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create points triggers in bulk. Maximum 100 triggers per request.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := []*trophygo.CreatePointsTriggerRequestItem{
        &trophygo.CreatePointsTriggerRequestItem{
            Type: trophygo.CreatePointsTriggerRequestItemTypeMetric,
            Points: 10,
        },
    }
client.Admin.Points.Triggers.Create(
        context.TODO(),
        "550e8400-e29b-41d4-a716-446655440000",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**systemId:** `string` — The UUID of the points system.
    
</dd>
</dl>

<dl>
<dd>

**request:** `trophygo.CreatePointsTriggersRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Admin.Points.Triggers.Delete(SystemId) -> *trophygo.DeletePointsTriggersResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete points triggers by ID. Maximum 100 trigger IDs per request.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &points.TriggersDeleteRequest{
        Ids: []*string{
            trophygo.String(
                "550e8400-e29b-41d4-a716-446655440000",
            ),
        },
    }
client.Admin.Points.Triggers.Delete(
        context.TODO(),
        "550e8400-e29b-41d4-a716-446655440000",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**systemId:** `string` — The UUID of the points system.
    
</dd>
</dl>

<dl>
<dd>

**ids:** `*string` — Trigger IDs to delete. Can be repeated or comma-separated.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Admin.Points.Triggers.Update(SystemId, request) -> *trophygo.PatchPointsTriggersResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update points triggers in bulk. Maximum 100 triggers per request. Only provided fields are updated; omitted fields are preserved.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := []*trophygo.PatchPointsTriggersRequestItem{
        &trophygo.PatchPointsTriggersRequestItem{
            Id: "id",
        },
    }
client.Admin.Points.Triggers.Update(
        context.TODO(),
        "550e8400-e29b-41d4-a716-446655440000",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**systemId:** `string` — The UUID of the points system.
    
</dd>
</dl>

<dl>
<dd>

**request:** `trophygo.PatchPointsTriggersRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Admin.Points.Triggers.Get(SystemId, Id) -> *trophygo.AdminPointsTrigger</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a single points trigger by ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Admin.Points.Triggers.Get(
        context.TODO(),
        "550e8400-e29b-41d4-a716-446655440000",
        "660f9500-f30c-42e5-b827-557766550001",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**systemId:** `string` — The UUID of the points system.
    
</dd>
</dl>

<dl>
<dd>

**id:** `string` — The UUID of the points trigger.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Admin Streaks Freezes
<details><summary><code>client.Admin.Streaks.Freezes.Create(request) -> *trophygo.CreateStreakFreezesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create streak freezes for multiple users.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &streaks.CreateStreakFreezesRequest{
        Freezes: []*streaks.CreateStreakFreezesRequestFreezesItem{
            &streaks.CreateStreakFreezesRequestFreezesItem{
                UserId: "user-123",
            },
            &streaks.CreateStreakFreezesRequestFreezesItem{
                UserId: "user-456",
            },
            &streaks.CreateStreakFreezesRequestFreezesItem{
                UserId: "user-123",
            },
        },
    }
client.Admin.Streaks.Freezes.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**freezes:** `[]*streaks.CreateStreakFreezesRequestFreezesItem` — Array of freezes to create. Maximum 100 freezes per request.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Admin Streaks Pauses
<details><summary><code>client.Admin.Streaks.Pauses.Create(request) -> *trophygo.CreateStreakPausesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create streak pauses for multiple users. A pause covers a specific date range and maintains the user's streak length during that range instead of ending the streak. Start dates in the past are rejected.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &streaks.CreateStreakPausesRequest{
        Pauses: []*streaks.CreateStreakPausesRequestPausesItem{
            &streaks.CreateStreakPausesRequestPausesItem{
                UserId: "user-123",
                Start: "2026-08-20",
                End: "2026-08-27",
            },
            &streaks.CreateStreakPausesRequestPausesItem{
                UserId: "user-456",
                Start: "2026-09-01",
                End: "2026-09-07",
            },
        },
    }
client.Admin.Streaks.Pauses.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**pauses:** `[]*streaks.CreateStreakPausesRequestPausesItem` — Array of pauses to create. Maximum 100 pauses per request.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Admin.Streaks.Pauses.Delete() -> *trophygo.DeleteStreakPausesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete streak pauses by ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &streaks.PausesDeleteRequest{
        Ids: []*string{
            trophygo.String(
                "550e8400-e29b-41d4-a716-446655440000",
            ),
            trophygo.String(
                "550e8400-e29b-41d4-a716-446655440001",
            ),
        },
    }
client.Admin.Streaks.Pauses.Delete(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**ids:** `*string` — Streak pause IDs to delete. Repeat the query param or provide a comma-separated list.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

