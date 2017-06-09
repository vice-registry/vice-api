package persistence

import (
	"log"

	gocb "gopkg.in/couchbase/gocb.v1"
	"omi-gitlab.e-technik.uni-ulm.de/vice/vice-api/models"
)

// CreateDeployment creates the provided deployment
func CreateDeployment(item *models.Deployment) (*models.Deployment, error) {
	id := generateID(defaultIDLength)
	item.ID = id
	bucket, err := couchbaseCredentials.Cluster.OpenBucket("vice-deployments", couchbaseCredentials.Password)
	if err != nil {
		log.Printf("Error in persistence CreateDeployment: cannot open bucket %s: %s", "vice-deployments", err)
		return nil, err
	}
	defer bucket.Close()
	_, err = bucket.Insert(id, item, 0)
	if err != nil {
		log.Printf("Error in persistence CreateDeployment: cannot create item %+v in bucket %s: %s", item, "vice-deployments", err)
		return nil, err
	}
	return item, nil
}

// UpdateDeployment updates the provided deployment
func UpdateDeployment(item *models.Deployment) (*models.Deployment, error) {
	err := updateItem(item, item.ID, "vice-deployments")
	if err != nil {
		return nil, err
	}
	return item, nil
}

// DeleteDeployment returns a single deployment by id
func DeleteDeployment(id string) error {
	err := deleteItem(id, "vice-deployment")
	if err != nil {
		return err
	}
	return nil
}

// GetDeployment returns a single deployment by id
func GetDeployment(id string) (*models.Deployment, error) {
	var item models.Deployment
	err := getItem(&item, id, "vice-deployment")
	if err != nil {
		return nil, err
	}
	return &item, nil
}

// GetDeployments returns an array of deployments of the authenticated user
func GetDeployments() ([]*models.Deployment, error) {
	query := gocb.NewN1qlQuery("SELECT id, userid, imageid, environmentId, environmentReference FROM `vice-environments` AS environments;")
	bucket, err := couchbaseCredentials.Cluster.OpenBucket("vice-deployment", couchbaseCredentials.Password)
	if err != nil {
		log.Printf("Error in persistence GetDeployments: cannot open bucket %s: %s", "vice-deployments", err)
		return nil, err
	}
	rows, err := bucket.ExecuteN1qlQuery(query, []interface{}{})
	if err != nil {
		log.Printf("Error in persistence GetDeployments: cannot run query on bucket %s: %s", "vice-deployments", err)
		return nil, err
	}
	var items []*models.Deployment
	var item models.Deployment
	for rows.Next(&item) {
		copy := new(models.Deployment)
		*copy = item
		items = append(items, copy)
	}
	return items, nil
}
