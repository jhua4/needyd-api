package routing

import (
	"context"
	"log"
	h "needyd/helpers"
	m "needyd/models"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type JobsHandler struct {
	collection *mongo.Collection
}

func (jobsHandler *JobsHandler) index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	h.Respond(w, "200", 200)
}

func (jobsHandler *JobsHandler) getJobs(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	keys, ok := r.URL.Query()["fromDate"]
	if !ok || len(keys) == 0 || len(keys[0]) < 1 {
		h.Respond(w, "Missing fromDate.", 400)
		return
	}
	fromDateStr := keys[0]

	keys, ok = r.URL.Query()["toDate"]
	if !ok || len(keys) == 0 || len(keys[0]) < 1 {
		h.Respond(w, "Missing toDate.", 400)
		return
	}
	toDateStr := keys[0]

	fromDate, err := time.Parse(time.RFC3339, fromDateStr)
	if err != nil {
		h.Respond(w, "Error parsing fromDate.\n"+err.Error(), 400)
		log.Println("Error while parsing date :", err.Error())
		return
	}

	toDate, err := time.Parse(time.RFC3339, toDateStr)
	if err != nil {
		h.Respond(w, "Error parsing toDate.\n"+err.Error(), 400)
		log.Println("Error while parsing date :", err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	jobs := []*m.Job{}
	cur, err := jobsHandler.collection.Find(ctx, bson.M{"posted": bson.M{"$gt": fromDate, "$lt": toDate}}, nil)
	if err != nil {
		h.Respond(w, err.Error(), 500)
		log.Println(err.Error())
		return
	}

	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var elem m.Job
		err := cur.Decode(&elem)
		if err != nil {
			h.Respond(w, err.Error(), 500)
			log.Println(err.Error())
			return
		}
		jobs = append(jobs, &elem)
	}

	if err := cur.Err(); err != nil {
		h.Respond(w, err.Error(), 500)
		log.Println(err.Error())
		return
	}

	h.Respond(w, jobs, 200)
}
