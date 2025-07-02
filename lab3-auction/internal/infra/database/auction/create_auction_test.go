package auction

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/andrevfarias/go-expert/lab3-auction/configuration/database/mongodb"
	"github.com/andrevfarias/go-expert/lab3-auction/internal/entity/auction_entity"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
)

type CreateAuctionTestSuite struct {
	suite.Suite
	dbMongo    *mongo.Database
	repository auction_entity.AuctionRepositoryInterface
}

func (s *CreateAuctionTestSuite) SetupTest() {
	os.Setenv("MONGODB_URL", "mongodb://admin:admin@mongodb:27017/auctions?authSource=admin")
	os.Setenv("MONGODB_DB", "test_auction")
	os.Setenv("AUCTION_INTERVAL", "1s")

	ctx := context.Background()
	dbMongo, err := mongodb.NewMongoDBConnection(ctx)
	s.dbMongo = dbMongo
	s.Require().NoError(err)

	s.repository = NewAuctionRepository(s.dbMongo)
}

func (s *CreateAuctionTestSuite) TearDownTest() {
	s.dbMongo.Drop(context.Background())
}

func TestCreateAuctionTestSuite(t *testing.T) {
	suite.Run(t, new(CreateAuctionTestSuite))
}

func (s *CreateAuctionTestSuite) TestCreateAuction_ShouldCloseAuctionAfterInterval() {
	// Arrange
	auction := &auction_entity.Auction{
		Id:          "test-id",
		ProductName: "Test Product",
		Category:    "Test Category",
		Description: "Test Description",
		Condition:   auction_entity.New,
		Status:      auction_entity.Active,
		Timestamp:   time.Now(),
	}

	// Act
	err := s.repository.CreateAuction(context.Background(), auction)
	s.Require().Nil(err)

	// Assert
	auctionEntity, err := s.repository.FindAuctionById(context.Background(), auction.Id)
	s.Require().Nil(err)
	s.Require().Equal(auction_entity.Active, auctionEntity.Status)

	time.Sleep(2 * time.Second)

	auctionEntity, err = s.repository.FindAuctionById(context.Background(), auction.Id)
	s.Require().Nil(err)
	s.Require().Equal(auction_entity.Completed, auctionEntity.Status)
}
