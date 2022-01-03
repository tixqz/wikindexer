package main

import (
    "time"
)

type Schedule struct {
    Sleep time.Duration 
    Cancel chan
    WorkersNum int8
}

func NewSchedule(sleep int, cancel chan, workersNum) *Schedule {
    return &Schedule{
        Sleep: sleep,
        Cancel: cancel,
        WorkersNum: workersNum,
    }
}

func (s *Schedule) Start(c *Client) {
    for i := range s.WorkersNum {
        go func(){}()
    }
}
