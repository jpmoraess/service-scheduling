package persistence

import (
	"context"
	"fmt"
	"github.com/jpmoraess/service-scheduling/internal/infra/persistence/util"
	"time"

	"github.com/jpmoraess/service-scheduling/configs"
	"github.com/jpmoraess/service-scheduling/internal/domain/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PasswordResetData struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	AccountID primitive.ObjectID `bson:"accountID"`
	Token     string             `bson:"token"`
	Expiry    time.Time          `bson:"expiry"`
}

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
	passwordResetData, err := toPasswordResetData(entity)
	if err != nil {
		fmt.Println("error parse password reset data from entity")
		return err
	}
	res, err := p.coll.InsertOne(ctx, passwordResetData)
	if err != nil {
		return err
	}
	fmt.Println("password token inserted with id: ", res.InsertedID.(primitive.ObjectID).Hex())
	return nil
}

func (p *PasswordResetMongoRepository) FindByToken(ctx context.Context, token string) (*entity.PasswordReset, error) {
	var passwordResetData PasswordResetData
	if err := p.coll.FindOne(ctx, bson.M{"token": token}).Decode(&passwordResetData); err != nil {
		return nil, err
	}
	passwordReset, err := fromPasswordResetData(&passwordResetData)
	if err != nil {
		return nil, err
	}
	return passwordReset, nil
}

func toPasswordResetData(entity *entity.PasswordReset) (*PasswordResetData, error) {
	accountID, err := util.GetObjectID(entity.AccountID())
	if err != nil {
		return nil, err
	}
	return &PasswordResetData{
		AccountID: accountID,
		Token:     entity.Token(),
		Expiry:    entity.Expiry(),
	}, nil
}

func fromPasswordResetData(data *PasswordResetData) (*entity.PasswordReset, error) {
	passwordReset, err := entity.RestorePasswordReset(data.ID.Hex(), data.AccountID.Hex(), data.Token, data.Expiry)
	if err != nil {
		return nil, err
	}
	return passwordReset, nil
}
