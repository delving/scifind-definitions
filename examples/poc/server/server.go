/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a simple gRPC server that demonstrates how to use gRPC-Go libraries
// to perform unary, client streaming, server streaming and full duplex RPCs.
//
// It implements the route guide service whose definition can be found in routeguide/route_guide.proto.
package main

import (
	"context"
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/examples/data"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	pb "github.com/delving/vhu-definition/examples/poc/vhu/v1"
)

//go:embed objects.json
var exampleData []byte

var (
	tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile   = flag.String("cert_file", "", "The TLS cert file")
	keyFile    = flag.String("key_file", "", "The TLS key file")
	jsonDBFile = flag.String("json_db_file", "", "A json file containing a list of objects")
	port       = flag.Int("port", 50051, "The server port")
	host       = flag.String("host", "localhost", "The server hostname")
)

type objectRecordServer struct {
	pb.UnimplementedObjectRecordServiceServer
	objects []*pb.ObjectRecord // read-only after initialized
}

func (s *objectRecordServer) GetObject(ctx context.Context, r *pb.GetObjectRequest) (*pb.GetObjectResponse, error) {
	for _, obj := range s.objects {
		if obj.GetId() == r.GetId() {
			return &pb.GetObjectResponse{Object: obj}, nil
		}
	}
	return nil, status.Errorf(codes.NotFound, "record with id %q is not found", r.GetId())
}

func (s *objectRecordServer) ListObjects(ctx context.Context, r *pb.ListObjectsRequest) (*pb.ListObjectsResponse, error) {
	return &pb.ListObjectsResponse{
		Objects: s.objects,
	}, nil
}

func newServer() *objectRecordServer {
	s := &objectRecordServer{}
	s.loadRecords(*jsonDBFile)
	return s
}

// loadRecords loads records from a JSON file.
func (s *objectRecordServer) loadRecords(filePath string) {
	var data []byte
	if filePath != "" {
		var err error
		data, err = os.ReadFile(filePath)
		if err != nil {
			log.Fatalf("Failed to load default features: %v", err)
		}
	} else {
		data = exampleData
	}
	if err := json.Unmarshal(data, &s.objects); err != nil {
		log.Fatalf("Failed to load default features: %v", err)
	}
}

func main() {
	flag.Parse()
	serverDSN := fmt.Sprintf("%s:%d", *host, *port)
	lis, err := net.Listen("tcp", serverDSN)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	if *tls {
		if *certFile == "" {
			*certFile = data.Path("x509/server_cert.pem")
		}
		if *keyFile == "" {
			*keyFile = data.Path("x509/server_key.pem")
		}
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			log.Fatalf("Failed to generate credentials: %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	grpcServer := grpc.NewServer(opts...)
	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)
	objectServer := newServer()
	pb.RegisterObjectRecordServiceServer(grpcServer, objectServer)
	log.Printf("started grpc server with %d records on %s", len(objectServer.objects), serverDSN)
	grpcServer.Serve(lis)
}
