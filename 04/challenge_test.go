package main

import "testing"

func Test_findSleepTime(t *testing.T) {
	type parameters struct {
		rawLogs []string
	}
	tests := []struct {
		name               string
		param              parameters
		wantmostAsleep     int
		wantfrequentMinute int
	}{
		{
			"example 1-1",
			parameters{[]string{
				"[1518-11-01 00:00] Guard #10 begins shift",
				"[1518-11-01 00:05] falls asleep",
				"[1518-11-01 00:25] wakes up",
				"[1518-11-01 00:30] falls asleep",
				"[1518-11-01 00:55] wakes up",
				"[1518-11-01 23:58] Guard #99 begins shift",
				"[1518-11-02 00:40] falls asleep",
				"[1518-11-02 00:50] wakes up",
				"[1518-11-03 00:05] Guard #10 begins shift",
				"[1518-11-03 00:24] falls asleep",
				"[1518-11-03 00:29] wakes up",
				"[1518-11-04 00:02] Guard #99 begins shift",
				"[1518-11-04 00:36] falls asleep",
				"[1518-11-04 00:46] wakes up",
				"[1518-11-05 00:03] Guard #99 begins shift",
				"[1518-11-05 00:45] falls asleep",
				"[1518-11-05 00:55] wakes up",
			}},
			240,
			4455,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotmostAsleep, gotfrequentMinute := findSleepTime(tt.param.rawLogs)
			if gotmostAsleep != tt.wantmostAsleep {
				t.Errorf("findSleepTime() mostAsleep: got %v, wanted %v", gotmostAsleep, tt.wantmostAsleep)
			}
			if gotfrequentMinute != tt.wantfrequentMinute {
				t.Errorf("findSleepTime() frequentMinute: got %v, wanted %v", gotfrequentMinute, tt.wantfrequentMinute)
			}
		})
	}
}
