package go_oura

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func TestGetActivity(t *testing.T) {
	tt := []struct {
		name           string
		documentId     string
		mockResponse   string
		expectedOutput DailyActivity
		expectErr      bool
	}{
		{name: "Valid Response", documentId: "1", mockResponse: `{"id":"45173cbe-ef26-430f-adc4-c4a1424b45ab","class_5_min":"111122321111111111111111111111111111233211111111111111113433322233232322223332230343333333333222222222222233222211111112323332200332333222323333323222322332232222222223222222222221123112111122222233232222222222221333332211111111111111111111111111111111111111123211111111111111111111111111","score":96,"active_calories":286,"average_met_minutes":1.3125,"contributors":{"meet_daily_targets":100,"move_every_hour":95,"recovery_time":97,"stay_active":81,"training_frequency":100,"training_volume":100},"equivalent_walking_distance":5195,"high_activity_met_minutes":13,"high_activity_time":120,"inactivity_alerts":1,"low_activity_met_minutes":156,"low_activity_time":13620,"medium_activity_met_minutes":50,"medium_activity_time":1020,"met":{"interval":60,"items":[0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,1.1,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,1.4,1.8,1.9,1.7,1.2,1.2,1.3,2.5,1.7,1.9,1.6,2.1,1.5,1.2,1.2,1.1,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,1.1,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,1.1,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,1.1,0.9,0.9,0.9,0.9,0.9,1.5,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,1.1,3.1,4.4,3.3,4,3.2,1.6,1.2,1.2,1.2,1.4,1.5,1.2,1.2,1.1,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,1.1,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,1.1,0.9,1.1,0.9,1.1,1.4,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,1.1,0.9,0.9,0.9,0.9,0.9,0.9,0.9,1,0.9,0.9,1.1,1.3,0.9,0.9,0.9,0.9,0.9,2.1,7,7.7,3.7,2.2,3,1.3,1.2,2.6,5.2,2.1,2.1,1.7,0.9,1.1,1.5,1.5,0.1,2,2.2,1.4,1.4,1.3,1.3,1.4,1.2,1.2,1.2,1.2,1.3,1.2,1.2,1.4,1.6,1.2,1.6,2,1.5,1.8,2.8,2.4,5.7,2.1,1.3,1.2,1.2,1.2,1.1,1,1.2,1,2.8,2.3,1,0.9,1.1,0.9,0.9,0.9,1.6,3,1.9,1.3,1.3,1.2,1.2,1.2,1.3,1.2,1.2,1.3,1.2,1.2,1.2,1.2,1.2,1.2,1.2,1.2,1.3,1.2,1.1,1.2,1.1,1.3,2.5,2.2,2.3,2.4,2.3,3.7,2.8,2.8,1.9,2.6,3,1.8,1.3,1.3,1.3,1.6,2.3,1.6,1.3,1.4,1.4,1.6,2.3,1.3,1.2,2.5,0.9,1.4,1.2,2.2,0.1,0.1,0.1,0.1,0.1,0.1,0.1,0.1,3.3,5,7.9,5,3.3,3.3,5.5,3.5,3.3,4.3,1.8,2.2,2.3,1.4,2.5,2.8,2.6,1.9,3.6,2.4,3.2,2.3,1.7,4,3.5,2.2,1.4,1.3,1.7,1.6,4,2.7,1.5,3.3,1.4,1.3,1.3,1.2,1.9,2.2,3.3,3.6,2.1,1.9,2.2,2.5,2.9,1.3,1.9,1.7,1.7,2.5,3,3.1,3,3.9,3.6,1.5,1.3,1.2,1.4,1.3,1.4,1.4,1.3,1.4,1.2,1.3,1.3,1.3,1.4,1.2,1.2,1.3,1.4,1.2,1.1,1.3,1.2,1.2,1.5,1.2,1.2,1.3,1.2,1.2,1.2,1.2,1.5,1.7,1.6,1.4,1.4,1.2,1.6,1.4,1.2,1.2,1.2,1.1,1.2,1.2,1.6,1.2,1.2,1.2,1.2,1.1,1.2,1.2,1.2,1.2,1.2,1.3,1.3,1.3,1.3,1.4,1.4,1.3,1.2,1.2,1.2,1.2,1.3,1.6,2.3,3.1,2.3,1.3,2.5,1.6,1.2,1.2,1.1,0.9,0.9,1.1,1.4,1.5,1.2,1.2,1.4,1.2,1.2,1.2,0.9,1.5,1.3,1.2,1.1,1.5,1.4,0.9,0.9,1,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,1.1,0.9,1,0.9,0.9,0.9,1.2,1,0.9,0.9,1,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,1.4,1,0.9,0.9,1,0.9,1,1,2.1,2,1.7,1.6,1.4,1.3,1.4,1.3,1.2,1.2,1.1,1.3,2.5,1,2.2,1.5,2.2,2.1,1.4,4,1.8,1.5,2.1,1.7,3,0.9,1.9,1.4,1.3,1.4,1.2,1.3,1.6,2,0.1,0.1,0.1,0.1,0.1,0.1,0.1,0.1,0.1,0.1,0.1,1.8,1.4,1.3,2.4,5.5,2.2,1.9,1.7,2.4,1.6,1.4,1.4,1.3,1.5,2.8,2.5,1.4,1.3,1.8,1.7,2,2.1,1.5,1.9,1.2,1.2,1.7,1.9,1.8,1.7,1.2,1.5,1.5,1.2,2.1,1.3,1.6,1.3,1.3,1.3,1.5,2.5,1.3,1.6,1.4,1.6,1.3,1.7,1.3,2.6,2.2,1.5,1.5,1.4,1.5,1.6,1.2,1.8,1.6,1.5,1.4,1.5,2.8,2.6,2.3,2,1.7,1.6,2,1.7,1.7,1.8,1.2,1.2,1.4,2.2,3.1,1.3,1.2,1.6,1.4,1.5,1.4,1.6,1.6,1.7,1.9,1.6,1.8,1.7,1.6,1.4,1.6,1.4,1.7,1.5,1.7,1.7,1.7,1.9,1.3,1.2,1.5,1.2,1.2,1.5,2.3,1.5,3.9,1.3,1.4,1.5,1.3,1.5,2,1.3,1.3,1.7,1.6,1.4,1.5,1.7,1.6,1.7,1.6,3.1,1.5,1.5,1.5,1.9,1.6,1.7,1.9,1.4,1.5,1.4,1.5,1.6,1.5,1.5,2.2,1.9,1.9,1.3,1.8,1.6,1.2,1.3,1.5,1.3,1.3,1.2,1.2,1.3,1.2,1.3,1.2,1.3,1.2,1.4,1.5,1.7,1.5,1.2,1.3,2.2,1.6,1.5,1.6,1.3,1.3,1.4,1.3,1.3,1.7,1.3,1.6,1.2,1.6,1.3,1.4,1.2,1.5,1.5,1.3,1.4,1.2,1.2,1.2,1.2,1.6,2,2.7,1.4,1.4,1.4,1.6,1.4,1.4,1.6,1.2,1.3,1.4,1.6,1.3,1.3,1.4,1.3,1.6,1.4,1.3,1.3,1.3,1.4,1.4,1.4,1.3,1.3,1.5,1.3,1.2,1.3,1.3,1.4,1.3,1.3,1.3,1.2,1.3,1.2,1.3,1.3,1.2,1.3,1.3,1.5,1.4,1.6,1.4,1.2,1.3,1.2,1.2,1.5,1.3,1.2,1.2,1.2,1.1,1.1,0.9,0.9,1,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,1.1,1.3,1.3,1.8,2.4,1.3,1.8,1.6,1.2,1.2,1.1,1,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,1.1,1,1.2,1.2,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,1.1,0.9,1.1,0.9,0.9,0.9,1.1,0.9,0.9,0.9,0.9,0.9,1,1.1,2.3,1.3,1.5,1.5,1.4,1.3,1.3,1.3,1.2,1.4,1.3,1.6,1.3,1.3,1.3,1.3,1.3,1.3,1.3,1.2,1.2,1.3,1.3,1.4,1.4,1.6,1.9,1.4,1.5,1.8,2.4,2.2,4.8,2.5,2.2,1.9,1.6,1.2,1.3,1.7,1.3,1.8,2.4,1.4,1.8,1.4,1.3,1.2,1.2,1.8,1.3,1.2,1.3,1.4,1.4,1.5,1.3,1.3,1.3,1.1,1.2,1.2,1.1,1.2,1.3,1.2,1.2,1.2,1.2,0.9,1.2,1.1,1.2,1.2,1.2,1.2,1.2,1.2,1.2,1.2,1.2,1.2,1.2,1.2,1.6,1.3,1.4,1.4,1.2,1.2,1.2,1.3,1.2,1.2,1.2,1.3,1.2,1.2,1.3,1.2,1.2,1.2,1.2,1.2,1.2,0.9,0.9,1,1,0.9,0.9,0.9,1.1,1.5,1.3,1.9,2.2,2.9,0.9,3,3,2.8,4.1,1.9,2.1,2,3.3,2.8,2.4,2.8,2.3,1.3,3.7,2.1,1.4,1.3,1.2,1.2,1.2,1.2,1.3,1.2,1.2,1.2,1.1,1.2,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,1.1,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,1.3,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,2.3,1.4,0.9,2.1,1.6,1.2,1,1.5,1.1,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,1.2,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,1.3,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,1.2,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,1.1,0.9,0.9,0.9,0.9,0.9,1.1,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9,0.9],"timestamp":"2024-01-01T04:00:00.000-05:00"},"meters_to_target":-500,"non_wear_time":1200,"resting_time":40200,"sedentary_met_minutes":9,"sedentary_time":30240,"steps":5797,"target_calories":250,"target_meters":6000,"total_calories":2382,"day":"2024-01-01","timestamp":"2024-01-01T04:00:00-05:00"}`,
			expectedOutput: DailyActivity{
				ID:                "45173cbe-ef26-430f-adc4-c4a1424b45ab",
				Class5Min:         "111122321111111111111111111111111111233211111111111111113433322233232322223332230343333333333222222222222233222211111112323332200332333222323333323222322332232222222223222222222221123112111122222233232222222222221333332211111111111111111111111111111111111111123211111111111111111111111111",
				Score:             96,
				ActiveCalories:    286,
				AverageMetMinutes: 1.3125,
				Contributors: Contributor{
					MeetDailyTargets:  100,
					MoveEveryHour:     95,
					RecoveryTime:      97,
					StayActive:        81,
					TrainingFrequency: 100,
					TrainingVolume:    100,
				},
				EquivalentWalkingDistance: 5195,
				HighActivityMetMinutes:    13,
				HighActivityTime:          120,
				InactivityAlerts:          1,
				LowActivityMetMinutes:     156,
				LowActivityTime:           13620,
				MediumActivityMetMinutes:  50,
				MediumActivityTime:        1020,
				Met: Met{
					Interval: 60,
					Items:    []float64{0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 1.1, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 1.4, 1.8, 1.9, 1.7, 1.2, 1.2, 1.3, 2.5, 1.7, 1.9, 1.6, 2.1, 1.5, 1.2, 1.2, 1.1, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 1.1, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 1.1, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 1.1, 0.9, 0.9, 0.9, 0.9, 0.9, 1.5, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 1.1, 3.1, 4.4, 3.3, 4, 3.2, 1.6, 1.2, 1.2, 1.2, 1.4, 1.5, 1.2, 1.2, 1.1, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 1.1, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 1.1, 0.9, 1.1, 0.9, 1.1, 1.4, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 1.1, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 1, 0.9, 0.9, 1.1, 1.3, 0.9, 0.9, 0.9, 0.9, 0.9, 2.1, 7, 7.7, 3.7, 2.2, 3, 1.3, 1.2, 2.6, 5.2, 2.1, 2.1, 1.7, 0.9, 1.1, 1.5, 1.5, 0.1, 2, 2.2, 1.4, 1.4, 1.3, 1.3, 1.4, 1.2, 1.2, 1.2, 1.2, 1.3, 1.2, 1.2, 1.4, 1.6, 1.2, 1.6, 2, 1.5, 1.8, 2.8, 2.4, 5.7, 2.1, 1.3, 1.2, 1.2, 1.2, 1.1, 1, 1.2, 1, 2.8, 2.3, 1, 0.9, 1.1, 0.9, 0.9, 0.9, 1.6, 3, 1.9, 1.3, 1.3, 1.2, 1.2, 1.2, 1.3, 1.2, 1.2, 1.3, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.3, 1.2, 1.1, 1.2, 1.1, 1.3, 2.5, 2.2, 2.3, 2.4, 2.3, 3.7, 2.8, 2.8, 1.9, 2.6, 3, 1.8, 1.3, 1.3, 1.3, 1.6, 2.3, 1.6, 1.3, 1.4, 1.4, 1.6, 2.3, 1.3, 1.2, 2.5, 0.9, 1.4, 1.2, 2.2, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 3.3, 5, 7.9, 5, 3.3, 3.3, 5.5, 3.5, 3.3, 4.3, 1.8, 2.2, 2.3, 1.4, 2.5, 2.8, 2.6, 1.9, 3.6, 2.4, 3.2, 2.3, 1.7, 4, 3.5, 2.2, 1.4, 1.3, 1.7, 1.6, 4, 2.7, 1.5, 3.3, 1.4, 1.3, 1.3, 1.2, 1.9, 2.2, 3.3, 3.6, 2.1, 1.9, 2.2, 2.5, 2.9, 1.3, 1.9, 1.7, 1.7, 2.5, 3, 3.1, 3, 3.9, 3.6, 1.5, 1.3, 1.2, 1.4, 1.3, 1.4, 1.4, 1.3, 1.4, 1.2, 1.3, 1.3, 1.3, 1.4, 1.2, 1.2, 1.3, 1.4, 1.2, 1.1, 1.3, 1.2, 1.2, 1.5, 1.2, 1.2, 1.3, 1.2, 1.2, 1.2, 1.2, 1.5, 1.7, 1.6, 1.4, 1.4, 1.2, 1.6, 1.4, 1.2, 1.2, 1.2, 1.1, 1.2, 1.2, 1.6, 1.2, 1.2, 1.2, 1.2, 1.1, 1.2, 1.2, 1.2, 1.2, 1.2, 1.3, 1.3, 1.3, 1.3, 1.4, 1.4, 1.3, 1.2, 1.2, 1.2, 1.2, 1.3, 1.6, 2.3, 3.1, 2.3, 1.3, 2.5, 1.6, 1.2, 1.2, 1.1, 0.9, 0.9, 1.1, 1.4, 1.5, 1.2, 1.2, 1.4, 1.2, 1.2, 1.2, 0.9, 1.5, 1.3, 1.2, 1.1, 1.5, 1.4, 0.9, 0.9, 1, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 1.1, 0.9, 1, 0.9, 0.9, 0.9, 1.2, 1, 0.9, 0.9, 1, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 1.4, 1, 0.9, 0.9, 1, 0.9, 1, 1, 2.1, 2, 1.7, 1.6, 1.4, 1.3, 1.4, 1.3, 1.2, 1.2, 1.1, 1.3, 2.5, 1, 2.2, 1.5, 2.2, 2.1, 1.4, 4, 1.8, 1.5, 2.1, 1.7, 3, 0.9, 1.9, 1.4, 1.3, 1.4, 1.2, 1.3, 1.6, 2, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 1.8, 1.4, 1.3, 2.4, 5.5, 2.2, 1.9, 1.7, 2.4, 1.6, 1.4, 1.4, 1.3, 1.5, 2.8, 2.5, 1.4, 1.3, 1.8, 1.7, 2, 2.1, 1.5, 1.9, 1.2, 1.2, 1.7, 1.9, 1.8, 1.7, 1.2, 1.5, 1.5, 1.2, 2.1, 1.3, 1.6, 1.3, 1.3, 1.3, 1.5, 2.5, 1.3, 1.6, 1.4, 1.6, 1.3, 1.7, 1.3, 2.6, 2.2, 1.5, 1.5, 1.4, 1.5, 1.6, 1.2, 1.8, 1.6, 1.5, 1.4, 1.5, 2.8, 2.6, 2.3, 2, 1.7, 1.6, 2, 1.7, 1.7, 1.8, 1.2, 1.2, 1.4, 2.2, 3.1, 1.3, 1.2, 1.6, 1.4, 1.5, 1.4, 1.6, 1.6, 1.7, 1.9, 1.6, 1.8, 1.7, 1.6, 1.4, 1.6, 1.4, 1.7, 1.5, 1.7, 1.7, 1.7, 1.9, 1.3, 1.2, 1.5, 1.2, 1.2, 1.5, 2.3, 1.5, 3.9, 1.3, 1.4, 1.5, 1.3, 1.5, 2, 1.3, 1.3, 1.7, 1.6, 1.4, 1.5, 1.7, 1.6, 1.7, 1.6, 3.1, 1.5, 1.5, 1.5, 1.9, 1.6, 1.7, 1.9, 1.4, 1.5, 1.4, 1.5, 1.6, 1.5, 1.5, 2.2, 1.9, 1.9, 1.3, 1.8, 1.6, 1.2, 1.3, 1.5, 1.3, 1.3, 1.2, 1.2, 1.3, 1.2, 1.3, 1.2, 1.3, 1.2, 1.4, 1.5, 1.7, 1.5, 1.2, 1.3, 2.2, 1.6, 1.5, 1.6, 1.3, 1.3, 1.4, 1.3, 1.3, 1.7, 1.3, 1.6, 1.2, 1.6, 1.3, 1.4, 1.2, 1.5, 1.5, 1.3, 1.4, 1.2, 1.2, 1.2, 1.2, 1.6, 2, 2.7, 1.4, 1.4, 1.4, 1.6, 1.4, 1.4, 1.6, 1.2, 1.3, 1.4, 1.6, 1.3, 1.3, 1.4, 1.3, 1.6, 1.4, 1.3, 1.3, 1.3, 1.4, 1.4, 1.4, 1.3, 1.3, 1.5, 1.3, 1.2, 1.3, 1.3, 1.4, 1.3, 1.3, 1.3, 1.2, 1.3, 1.2, 1.3, 1.3, 1.2, 1.3, 1.3, 1.5, 1.4, 1.6, 1.4, 1.2, 1.3, 1.2, 1.2, 1.5, 1.3, 1.2, 1.2, 1.2, 1.1, 1.1, 0.9, 0.9, 1, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 1.1, 1.3, 1.3, 1.8, 2.4, 1.3, 1.8, 1.6, 1.2, 1.2, 1.1, 1, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 1.1, 1, 1.2, 1.2, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 1.1, 0.9, 1.1, 0.9, 0.9, 0.9, 1.1, 0.9, 0.9, 0.9, 0.9, 0.9, 1, 1.1, 2.3, 1.3, 1.5, 1.5, 1.4, 1.3, 1.3, 1.3, 1.2, 1.4, 1.3, 1.6, 1.3, 1.3, 1.3, 1.3, 1.3, 1.3, 1.3, 1.2, 1.2, 1.3, 1.3, 1.4, 1.4, 1.6, 1.9, 1.4, 1.5, 1.8, 2.4, 2.2, 4.8, 2.5, 2.2, 1.9, 1.6, 1.2, 1.3, 1.7, 1.3, 1.8, 2.4, 1.4, 1.8, 1.4, 1.3, 1.2, 1.2, 1.8, 1.3, 1.2, 1.3, 1.4, 1.4, 1.5, 1.3, 1.3, 1.3, 1.1, 1.2, 1.2, 1.1, 1.2, 1.3, 1.2, 1.2, 1.2, 1.2, 0.9, 1.2, 1.1, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.6, 1.3, 1.4, 1.4, 1.2, 1.2, 1.2, 1.3, 1.2, 1.2, 1.2, 1.3, 1.2, 1.2, 1.3, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 0.9, 0.9, 1, 1, 0.9, 0.9, 0.9, 1.1, 1.5, 1.3, 1.9, 2.2, 2.9, 0.9, 3, 3, 2.8, 4.1, 1.9, 2.1, 2, 3.3, 2.8, 2.4, 2.8, 2.3, 1.3, 3.7, 2.1, 1.4, 1.3, 1.2, 1.2, 1.2, 1.2, 1.3, 1.2, 1.2, 1.2, 1.1, 1.2, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 1.1, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 1.3, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 2.3, 1.4, 0.9, 2.1, 1.6, 1.2, 1, 1.5, 1.1, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 1.2, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 1.3, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 1.2, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 1.1, 0.9, 0.9, 0.9, 0.9, 0.9, 1.1, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9}, // Place the entire items array from the JSON here
					Timestamp: func() time.Time {
						layout := "2006-01-02T15:04:05Z07:00"
						t, _ := time.Parse(layout, "2024-01-01T04:00:00.000-05:00")
						return t
					}(),
				},
				MetersToTarget:      -500,
				NonWearTime:         1200,
				RestingTime:         40200,
				SedentaryMetMinutes: 9,
				SedentaryTime:       30240,
				Steps:               5797,
				TargetCalories:      250,
				TargetMeters:        6000,
				TotalCalories:       2382,
				Day: func() Date {
					layout := "2006-01-02"
					t, _ := time.Parse(layout, "2024-01-01")
					return Date{Time: t}
				}(),
				Timestamp: func() time.Time {
					layout := "2006-01-02T15:04:05Z07:00"
					t, _ := time.Parse(layout, "2024-01-01T04:00:00-05:00")
					return t
				}(),
			}, expectErr: false},
		{
			name:           "Invalid Response",
			documentId:     "2",
			mockResponse:   `{"message": "invalid"}`,
			expectedOutput: DailyActivity{},
			expectErr:      true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				_, err := rw.Write([]byte(tc.mockResponse))
				if err != nil {
					http.Error(rw, err.Error(), http.StatusInternalServerError)
					return
				}
			}))

			client := NewClientWithUrlAndHttp("", server.URL, server.Client())

			activity, err := client.GetActivity(tc.documentId)
			if tc.expectErr {
				if err == nil {
					t.Errorf("Expected error, got nil")
				}

				var ouraErr *OuraError
				if !errors.As(err, &ouraErr) {
					t.Errorf("expected an OuraError but got a different error: %v", err)
				}

				return
			} else if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			if !reflect.DeepEqual(activity, tc.expectedOutput) {
				t.Errorf("Expected %v, got %v", tc.expectedOutput, activity)
			}
		})
	}
}
