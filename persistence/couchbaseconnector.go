package persistence

import (
	"fmt"
	"math/rand"

	"gopkg.in/couchbase/gocb.v1"
	"omi-gitlab.e-technik.uni-ulm.de/vice/vice-api/models"
)

const idLetterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
const defaultIDLength int = 6

func generateID(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = idLetterBytes[rand.Intn(len(idLetterBytes))]
	}
	return string(b)
}

func getItem(cbCluster *gocb.Cluster, item interface{}, id string, bucketName string) {
	bucket, _ := cbCluster.OpenBucket(bucketName, "")
	defer bucket.Close()
	bucket.Get(id, item)
}

func updateItem(cbCluster *gocb.Cluster, item interface{}, id string, bucketName string) {
	bucket, _ := cbCluster.OpenBucket(bucketName, "")
	defer bucket.Close()
	bucket.Replace(id, item, 0, 0)
}

func deleteItem(cbCluster *gocb.Cluster, id string, bucketName string) {
	bucket, _ := cbCluster.OpenBucket(bucketName, "")
	defer bucket.Close()
	bucket.Remove(id, 0)
}

func CreateEnvironment(cbCluster *gocb.Cluster, item *models.Environment) *models.Environment {
	id := generateID(defaultIDLength)
	item.ID = id
	bucket, _ := cbCluster.OpenBucket("vice-environments", "")
	defer bucket.Close()
	bucket.Insert(id, item, 0)
	return item
}

func CreateUser(cbCluster *gocb.Cluster, item *models.User) *models.User {
	id := generateID(defaultIDLength)
	item.ID = id
	bucket, _ := cbCluster.OpenBucket("vice-users", "")
	defer bucket.Close()
	bucket.Insert(id, item, 0)
	return item
}

func CreateImage(cbCluster *gocb.Cluster, item *models.Image) *models.Image {
	id := generateID(defaultIDLength)
	item.ID = id
	bucket, _ := cbCluster.OpenBucket("vice-images", "")
	defer bucket.Close()
	bucket.Insert(id, item, 0)
	return item
}

func UpdateEnvironment(cbCluster *gocb.Cluster, item models.Environment) models.Environment {
	updateItem(cbCluster, item, item.ID, "vice-environments")
	return item
}

func UpdateUser(cbCluster *gocb.Cluster, item models.User) models.User {
	updateItem(cbCluster, item, item.ID, "vice-users")
	return item
}

func UpdateImage(cbCluster *gocb.Cluster, item models.Image) models.Image {
	updateItem(cbCluster, item, item.ID, "vice-images")
	return item
}

func DeleteEnvironment(cbCluster *gocb.Cluster, id string) {
	deleteItem(cbCluster, id, "vice-environments")
}

func DeleteUser(cbCluster *gocb.Cluster, id string) {
	deleteItem(cbCluster, id, "vice-users")
}

func DeleteImage(cbCluster *gocb.Cluster, id string) {
	deleteItem(cbCluster, id, "vice-images")
}

func GetEnvironment(cbCluster *gocb.Cluster, id string) *models.Environment {
	var item models.Environment
	getItem(cbCluster, &item, id, "vice-environments")
	return &item
}

func GetEnvironments(cbCluster *gocb.Cluster) []*models.Environment {
	query := gocb.NewN1qlQuery("SELECT id, credentials, managementlayer, runtimetechnology, userid FROM `vice-environments` AS environments;")
	bucket, _ := cbCluster.OpenBucket("vice-environments", "")
	rows, _ := bucket.ExecuteN1qlQuery(query, []interface{}{})
	var items []*models.Environment
	var item models.Environment
	for rows.Next(&item) {
		copy := new(models.Environment)
		*copy = item
		items = append(items, copy)
	}
	return items
}

func GetUser(cbCluster *gocb.Cluster, id string) *models.User {
	var item models.User
	getItem(cbCluster, &item, id, "vice-users")
	return &item
}

func GetUserByName(cbCluster *gocb.Cluster, name string) models.User {
	bucket, _ := cbCluster.OpenBucket("vice-users", "")
	query := gocb.NewN1qlQuery("SELECT `id`, `username`, `password` FROM `vice-users` AS users WHERE `username`=$1;")
	var params []interface{}
	params = append(params, name)

	rows, err := bucket.ExecuteN1qlQuery(query, params)
	if err != nil {
		fmt.Println("ERROR EXECUTING N1QL QUERY:", err)
	}
	var item models.User
	rows.One(&item)
	return item
}

func GetImage(cbCluster *gocb.Cluster, id string) *models.Image {
	var item models.Image
	getItem(cbCluster, &item, id, "vice-images")
	return &item
}

func GetImages(cbCluster *gocb.Cluster) []*models.Image {
	query := gocb.NewN1qlQuery("SELECT `id`, `content-type`, `image-type`, `originEnvironment`, `userid` FROM `vice-images` AS images;")
	bucket, _ := cbCluster.OpenBucket("vice-images", "")
	rows, _ := bucket.ExecuteN1qlQuery(query, []interface{}{})
	var items []*models.Image
	var item models.Image
	for rows.Next(&item) {
		copy := new(models.Image)
		*copy = item
		items = append(items, copy)
	}
	return items
}
