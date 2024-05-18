package persistence

import (
	"context"
	"fmt"

	"github.com/jpmoraess/service-scheduling/configs"
	"github.com/jpmoraess/service-scheduling/internal/domain/entity"
	"github.com/jpmoraess/service-scheduling/internal/infra/persistence/mapper"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PasswordResetMongoRepository struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewPasswordResetMongoRepository(client *mongo.Client) *PasswordResetMongoRepository {
	return &PasswordResetMongoRepository{
		client: client,
		coll:   client.Database(configs.DATABASE_NAME).Collection("password_reset"),
	}
}

func (p *PasswordResetMongoRepository) Save(ctx context.Context, entity *entity.PasswordReset) error {
	passwordResetData, err := mapper.ToPasswordResetData(entity)
	if err != nil {
		fmt.Println("error parse password reset data from entity")
		return err
	}
	res, err := p.coll.InsertOne(ctx, passwordResetData)
	if err != nil {
		return err
	}
	fmt.Println("password token inserted: ", res.InsertedID.(primitive.ObjectID).Hex())
	return nil
}
