package tests

import (
	"errors"
	"github.com/austinmoody/go-oura"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func TestGetSleepTime(t *testing.T) {
	tt := []struct {
		name           string
		documentId     string
		mockResponse   string
		expectedOutput go_oura.SleepTime
		expectErr      bool
	}{
		{
			name:         "Valid_SleepTime_Response_Without_Optimal",
			documentId:   "1",
			mockResponse: `{"id":"bb1044c6-6d85-406b-9bcd-0ce7dd438608","day":"2024-01-12","optimal_bedtime":null,"recommendation":"earlier_bedtime","status":"only_recommended_found"}`,
			expectedOutput: go_oura.SleepTime{
				ID: "bb1044c6-6d85-406b-9bcd-0ce7dd438608",
				Day: func() go_oura.Date {
					layout := "2006-01-02"
					t, _ := time.Parse(layout, "2024-01-12")
					return go_oura.Date{Time: t}
				}(),
				OptimalBedtime: nil,
				Recommendation: "earlier_bedtime",
				Status:         "only_recommended_found",
			},
		},
		{
			name:         "Valid_SleepTime_Response_With_Optimal",
			documentId:   "valid-with-optimal",
			mockResponse: `{"id":"bb1044c6-6d85-406b-9bcd-0ce7dd438608","day":"2024-01-12","optimal_bedtime":{"day_tz":1,"end_offset":2,"start_offset":3},"recommendation":"earlier_bedtime","status":"only_recommended_found"}`,
			expectedOutput: go_oura.SleepTime{
				ID: "bb1044c6-6d85-406b-9bcd-0ce7dd438608",
				Day: func() go_oura.Date {
					layout := "2006-01-02"
					t, _ := time.Parse(layout, "2024-01-12")
					return go_oura.Date{Time: t}
				}(),
				OptimalBedtime: &go_oura.OptimalBedtime{
					DayTz:       1,
					EndOffset:   2,
					StartOffset: 3,
				},
				Recommendation: "earlier_bedtime",
				Status:         "only_recommended_found",
			},
			expectErr: false,
		},
		{
			name:           "Invalid_SleepTime_Response",
			documentId:     "2",
			mockResponse:   `{"message": "invalid"}`,
			expectedOutput: go_oura.SleepTime{},
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

			client := go_oura.NewClientWithUrlAndHttp("", server.URL, server.Client())

			activity, err := client.GetSleepTime(tc.documentId)
			if tc.expectErr {
				if err == nil {
					t.Errorf("Expected error, got nil")
				}

				var ouraErr *go_oura.OuraError
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

//func TestGetSleepDocuments(t *testing.T) {
//	tt := []struct {
//		name           string
//		startTime      time.Time
//		endTime        time.Time
//		mockResponse   string
//		expectedOutput go_oura.Sleeps
//		expectErr      bool
//	}{
//		{
//			name:         "Valid_SleepDocuments_Response",
//			startTime:    time.Now().Add(-1 * time.Hour),
//			endTime:      time.Now().Add(-2 * time.Hour),
//			mockResponse: `{"data":[{"id":"4eaa0e18-3464-49cc-961a-2ffd5f8ea98e","contributors":{"deep_sleep":63,"efficiency":93,"latency":64,"rem_sleep":95,"restfulness":72,"timing":94,"total_sleep":90},"day":"2024-01-07","score":83,"timestamp":"2024-01-07T00:00:00+00:00"}],"next_token":null}`,
//			expectedOutput: go_oura.Sleeps{
//				Items: []go_oura.Sleep{
//					{
//						ID: "4eaa0e18-3464-49cc-961a-2ffd5f8ea98e",
//						Day: func() go_oura.Date {
//							layout := "2006-01-02"
//							t, _ := time.Parse(layout, "2024-01-07")
//							return go_oura.Date{Time: t}
//						}(),
//						Score: 83,
//						Timestamp: func() time.Time {
//							layout := "2006-01-02T15:04:05Z07:00"
//							t, _ := time.Parse(layout, "2024-01-07T00:00:00+00:00")
//							return t
//						}(),
//						Contributors: go_oura.SleepContributors{
//							DeepSleep:   63,
//							Efficiency:  93,
//							Latency:     64,
//							RemSleep:    95,
//							Restfulness: 72,
//							Timing:      94,
//							TotalSleep:  90,
//						},
//					},
//				},
//			},
//			expectErr: false,
//		}, {
//			name:           "Invalid_SleepDocuments_Response",
//			startTime:      time.Now(),
//			endTime:        time.Now(),
//			mockResponse:   `{"message": "invalid"}`,
//			expectedOutput: go_oura.Sleeps{},
//			expectErr:      true,
//		},
//	}
//
//	for _, tc := range tt {
//		t.Run(tc.name, func(t *testing.T) {
//			server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
//				_, err := rw.Write([]byte(tc.mockResponse))
//				if err != nil {
//					http.Error(rw, err.Error(), http.StatusInternalServerError)
//					return
//				}
//			}))
//
//			client := go_oura.NewClientWithUrlAndHttp("", server.URL, server.Client())
//
//			activity, err := client.GetSleeps(tc.startTime, tc.endTime, nil)
//			if tc.expectErr {
//				if err == nil {
//					t.Errorf("Expected error, got nil")
//				}
//
//				var ouraErr *go_oura.OuraError
//				if !errors.As(err, &ouraErr) {
//					t.Errorf("expected an OuraError but got a different error: %v", err)
//				}
//
//				return
//			} else if err != nil {
//				t.Errorf("Unexpected error: %v", err)
//				return
//			}
//
//			if !reflect.DeepEqual(activity, tc.expectedOutput) {
//				t.Errorf("Expected %v, got %v", tc.expectedOutput, activity)
//			}
//		})
//	}
//}