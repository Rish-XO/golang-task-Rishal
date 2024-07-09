package main

import (
    "context"
    "errors"

    pb "golang-task-Rishal/userpb"
)

type Server struct {
    pb.UnimplementedUserServiceServer
    users []pb.User
}

func (s *Server) GetUser(ctx context.Context, req *pb.UserIdRequest) (*pb.UserResponse, error) {
    for _, user := range s.users {
        if user.Id == req.Id {
            return &pb.UserResponse{User: &user}, nil
        }
    }
    return nil, errors.New("user not found")
}

func (s *Server) GetUsers(ctx context.Context, req *pb.UserIdsRequest) (*pb.UsersResponse, error) {
    var users []*pb.User
    for _, id := range req.Ids {
        for _, user := range s.users {
            if user.Id == id {
                users = append(users, &user)
            }
        }
    }
    return &pb.UsersResponse{Users: users}, nil
}

func (s *Server) SearchUsers(ctx context.Context, req *pb.SearchRequest) (*pb.UsersResponse, error) {
    var users []*pb.User
    for _, user := range s.users {
        if (req.City == "" || user.City == req.City) &&
            (req.Phone == 0 || user.Phone == req.Phone) &&
            (!req.Married || user.Married == req.Married) {
            users = append(users, &user)
        }
    }
    return &pb.UsersResponse{Users: users}, nil
}
