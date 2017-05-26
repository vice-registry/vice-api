package main

import (
	"fmt"

	"omi-gitlab.e-technik.uni-ulm.de/vice/vice-api/models"

	"gopkg.in/couchbase/gocb.v1"
)

/*func main() {
	initViceCouchbase()
}*/

func InitViceCouchbase() {
	cluster, _ := gocb.Connect("couchbase://localhost")

	clusterManager := cluster.Manager("", "")
	createBucket(cluster, clusterManager, "vice-users")
	createBucket(cluster, clusterManager, "vice-environments")
	createBucket(cluster, clusterManager, "vice-images")

	createAdminUser(cluster)
}

func createBucket(cluster *gocb.Cluster, clusterManager *gocb.ClusterManager, bucketname string) {
	settings := gocb.BucketSettings{}
	settings.Name = bucketname
	clusterManager.InsertBucket(&settings)

	bucket, _ := cluster.OpenBucket(bucketname, "")
	defer bucket.Close()
	bucketManager := bucket.Manager("", "")
	bucketManager.CreatePrimaryIndex("", true, false)
}

func createAdminUser(cluster *gocb.Cluster) {
	bucket, _ := cluster.OpenBucket("vice-users", "")
	var admin models.User
	admin.ID = "admin"
	admin.Username = "admin"
	admin.Password = "admin"
	bucket.Insert(admin.ID, admin, 0)
}

func TestingN1ql() {

	cluster, _ := gocb.Connect("couchbase://localhost")
	bucket, _ := cluster.OpenBucket("vice-users", "")

	// Setup a new query with a placeholder
	myQuery := gocb.NewN1qlQuery("SELECT * FROM `vice-users` AS users WHERE username=$1")

	// Setup an array for parameters
	var myParams []interface{}
	myParams = append(myParams, "user")

	// Execute Query
	rows, err := bucket.ExecuteN1qlQuery(myQuery, myParams)
	if err != nil {
		fmt.Println("ERROR EXECUTING N1QL QUERY:", err)
	}

	// Iterate through rows and print output
	var row interface{}
	for rows.Next(&row) {
		fmt.Printf("Results: %+v\n", row)
	}

}

func Testing() {

	cluster, _ := gocb.Connect("couchbase://localhost")
	bucket, _ := cluster.OpenBucket("vice-environments", "")

	/*creds := models.Credentials{
		Endpoint: "bla",
		Password: "blu",
		Username: "blo",
	}

	runtime := models.RuntimeTechnology{
		Software: "kvm",
		Type:     models.RuntimeTechnologyTypeVirtualmachine,
		Version:  "",
	}

	management := models.ManagementLayer{
		Software: "OpenStack",
		Type:     models.ManagementLayerTypeCloudcomputing,
		Version:  "kilo",
	}

	bucket.Insert("envtest",
	models.Environment{
		ID:                "envtest",
		Userid:            "admin",
		Credentials:       &creds,
		Runtimetechnology: &runtime,
		Managementlayer:   &management,
	}, 0)*/

	// Get the value back
	var item models.Environment
	bucket.Get("nfgDsc", &item)
	fmt.Printf("Item: %s\n", item)

	bucket.Get("envtest", &item)
	fmt.Printf("Item: %s\n", item)

}
