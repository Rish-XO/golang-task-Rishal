package main

import (
    "log"
    "net"

    "google.golang.org/grpc"
    pb "golang-task-Rishal/userpb"
)

func main() {
    users := []pb.User{
        {Id: 1, Fname: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
        {Id: 2, Fname: "John", City: "NY", Phone: 9876543210, Height: 6.0, Married: false},
        // Add more users as needed
    }

    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    s := grpc.NewServer()
    pb.RegisterUserServiceServer(s, &server{users: users})
    log.Printf("server listening at %v", lis.Addr())

    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
