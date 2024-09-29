package v1

import (
	"context"
	"fmt"
	"github.com/mokaz111/candy-server/biz/dal/mongo"
	v1 "github.com/mokaz111/candy-server/hertz_gen/api/v1"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// NewClusterService 创建一个新的 ClusterService 实例
func NewClusterService() *ClusterService {
	return &ClusterService{}
}

// ClusterService 实现了 v1.ClusterService 接口
type ClusterService struct{}

// CreateCluster 创建集群
func (s *ClusterService) CreateCluster(ctx context.Context, req *v1.CreateClusterRequest) (*v1.CreateClusterResponse, error) {

	cluster := &v1.Cluster{
		Name:         req.Name,
		ApiServerUrl: req.ApiServerUrl,
		AuthToken:    req.AuthToken,
		CreatedAt:    timestamppb.Now(),
		UpdatedAt:    timestamppb.Now(),
	}

	collection := mongo.MGDB.Collection("clusters")
	result, err := collection.InsertOne(ctx, cluster)
	if err != nil {
		return nil, err
	}

	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, fmt.Errorf("failed to convert inserted ID to ObjectID")
	}

	cluster.Id = insertedID.Hex()

	return &v1.CreateClusterResponse{
		Id:           cluster.Id,
		Name:         cluster.Name,
		ApiServerUrl: cluster.ApiServerUrl,
		AuthToken:    cluster.AuthToken,
		CreatedAt:    cluster.CreatedAt,
		UpdatedAt:    cluster.UpdatedAt,
	}, nil
}

// GetClusters 获取所有集群
func (s *ClusterService) GetClusters(ctx context.Context, _ *v1.GetClustersRequest) (*v1.GetClustersResponse, error) {
	var clusters []*v1.Cluster
	collection := mongo.MGDB.Collection("clusters")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &clusters); err != nil {
		return nil, err
	}
	return &v1.GetClustersResponse{Clusters: clusters}, nil
}

// UpdateCluster 更新集群
func (s *ClusterService) UpdateCluster(ctx context.Context, req *v1.UpdateClusterRequest) (*v1.UpdateClusterResponse, error) {
	objID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objID}
	update := bson.M{
		"$set": bson.M{
			"name":           req.Name,
			"api_server_url": req.ApiServerUrl,
			"auth_token":     req.AuthToken,
			"updated_at":     timestamppb.Now(),
		},
	}

	collection := mongo.MGDB.Collection("clusters")
	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateClusterResponse{}, nil
}

// DeleteCluster 删除集群
func (s *ClusterService) DeleteCluster(ctx context.Context, req *v1.DeleteClusterRequest) (*v1.DeleteClusterResponse, error) {
	objID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objID}
	collection := mongo.MGDB.Collection("clusters")
	_, err = collection.DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteClusterResponse{}, nil
}
