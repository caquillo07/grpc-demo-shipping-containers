package main

import (
	"fmt"
	pb "github.com/caquillo07/shipper/consignment-service/proto/consignment"

	vesselProto "github.com/caquillo07/shipper/vessel-service/proto/vessel"
	"github.com/micro/go-micro"
	"golang.org/x/net/context"
	"log"
)

type Repository interface {
	Create(consignment *pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

type ConsignmentRepository struct {
	consignments []*pb.Consignment
}

func (repo *ConsignmentRepository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	return consignment, nil
}

func (repo *ConsignmentRepository) GetAll() []*pb.Consignment {
	return repo.consignments
}

type service struct {
	repo         Repository
	vesselClient vesselProto.VesselServiceClient
}

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {

	// Here we call a client instance of our vessel service with our consignment weight,
	// and the amount of containers as the capacity value
	vesselResponse, err := s.vesselClient.FindAvailable(context.Background(), &vesselProto.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})
	log.Printf("Found vessel: %s \n", vesselResponse.Vessel.Name)
	if err != nil {
		return err
	}

	// We set the VesselId as the vessel we got back from our
	// vessel service
	req.VesselId = vesselResponse.Vessel.Id

	// Save our consignment
	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}

	// Return matching the `Response` message we created in our
	// protobuf definition.
	res.Created = true
	res.Consignment = consignment
	return nil
}

func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	consignments := s.repo.GetAll()
	res.Consignments = consignments
	return nil
}

func main() {
	repo := &ConsignmentRepository{}

	// create a new service. optionally include some options here
	srv := micro.NewService(
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)

	vesselService := vesselProto.NewVesselServiceClient("go.micro.srv.vessel", srv.Client())

	// init will parse the command line flags
	srv.Init()

	// register handler
	pb.RegisterShippingServiceHandler(srv.Server(), &service{repo: repo, vesselClient: vesselService})

	// run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
