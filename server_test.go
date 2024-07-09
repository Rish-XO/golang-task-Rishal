package main

import (
    "context"
    "testing"

    pb "golang-task-Rishal/userpb"
)

func TestGetUser(t *testing.T) {
    s := &server{
        users: []pb.User{
            {Id: 1, Fname: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
            {Id: 2, Fname: "John", City: "NY", Phone: 9876543210, Height: 6.0, Married: false},
        },
    }
    req := &pb.UserIdRequest{Id: 1}
    res, err := s.GetUser(context.Background(), req)
    if err != nil {
        t.Fatalf("GetUser failed: %v", err)
    }
    if res.User.Id != 1 {
        t.Errorf("expected user ID 1, got %d", res.User.Id)
    }
}

func TestGetUsers(t *testing.T) {
    s := &server{
        users: []pb.User{
            {Id: 1, Fname: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
            {Id: 2, Fname: "John", City: "NY", Phone: 9876543210, Height: 6.0, Married: false},
        },
    }
    req := &pb.UserIdsRequest{Ids: []int32{1, 2}}
    res, err := s.GetUsers(context.Background(), req)
    if err != nil {
        t.Fatalf("GetUsers failed: %v", err)
    }
    if len(res.Users) != 2 {
        t.Errorf("expected 2 users, got %d", len(res.Users))
    }
}

func TestSearchUsers(t *testing.T) {
    s := &server{
        users: []pb.User{
            {Id: 1, Fname: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
            {Id: 2, Fname: "John", City: "NY", Phone: 9876543210, Height: 6.0, Married: false},
        },
    }
    req := &pb.SearchRequest{City: "LA"}
    res, err := s.SearchUsers(context.Background(), req)
    if err != nil {
        t.Fatalf("SearchUsers failed: %v", err)
    }
    if len(res.Users) != 1 {
        t.Errorf("expected 1 user, got %d", len(res.Users))
    }
}
