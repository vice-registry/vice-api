package persistence

import (
	"log"

	gocb "gopkg.in/couchbase/gocb.v1"
	"github.com/vice-registry/vice-api/models"
)

// CreateImage creates the provided image
func CreateImage(item *models.Image) (*models.Image, error) {
	id := generateID(defaultIDLength)
	item.ID = id
	bucket, err := couchbaseCredentials.Cluster.OpenBucket("vice-images", couchbaseCredentials.Password)
	if err != nil {
		log.Printf("Error in persistence CreateImage: cannot open bucket %s: %s", "vice-images", err)
		return nil, err
	}
	defer bucket.Close()
	_, err = bucket.Insert(id, item, 0)
	if err != nil {
		log.Printf("Error in persistence CreateImage: cannot create item %+v in bucket %s: %s", item, "vice-images", err)
		return nil, err
	}
	return item, nil
}

// UpdateImage updates the provided image
func UpdateImage(item *models.Image) (*models.Image, error) {
	err := updateItem(item, item.ID, "vice-images")
	if err != nil {
		return nil, err
	}
	return item, nil
}

// DeleteImage returns a single image by id
func DeleteImage(id string) error {
	err := deleteItem(id, "vice-images")
	if err != nil {
		return err
	}
	return nil
}

// GetImage returns a single image by id
func GetImage(id string) (*models.Image, error) {
	var item models.Image
	err := getItem(&item, id, "vice-images")
	if err != nil {
		return nil, err
	}
	return &item, nil
}

// GetImages returns an array of images of the authenticated user
func GetImages() ([]*models.Image, error) {
	var items []*models.Image
	query := gocb.NewN1qlQuery("SELECT `id`, `content-type`, `image-type`, `originEnvironment`, `userid` FROM `vice-images` AS images;")
	bucket, err := couchbaseCredentials.Cluster.OpenBucket("vice-images", couchbaseCredentials.Password)
	if err != nil {
		log.Printf("Error in persistence GetImages: cannot open bucket %s: %s", "vice-images", err)
		return nil, err
	}
	rows, err := bucket.ExecuteN1qlQuery(query, []interface{}{})
	if err != nil {
		log.Printf("Error in persistence GetImages: cannot run query on bucket %s: %s", "vice-images", err)
		return nil, err
	}
	if err != nil {
		log.Printf("Failed to execute query to couchbase bucket vice-images: %s", err)
	}
	var item models.Image
	for rows.Next(&item) {
		copy := new(models.Image)
		*copy = item
		items = append(items, copy)
	}
	return items, nil
}
