package appengine

import (
	"cloud.google.com/go/datastore"
	"context"
	"fmt"
	"log"
	"net/http"
)

// Controller는 저장소 및 기타 상태를 저장한다.
type Controller struct {
	Store *datastore.Client
}

func (c *Controller) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "invalid method", http.StatusMethodNotAllowed)
		return
	}

	ctx := context.Background()

	// 새 메시지를 저장한다.
	r.ParseForm()
	if message := r.FormValue("message"); message != "" {
		if err := c.storeMessage(ctx, message); err != nil {
			log.Printf("could not store message: %v", err)
			http.Error(w, "could not store message", http.StatusInternalServerError)
			return
		}
	}

	// 현재 메시지를 가져와 보여준다.
	fmt.Fprintf(w, "Message:")
	messages, err := c.queryMessages(ctx, 10)
	if err != nil {
		log.Printf("could not get messages: %v", err)
		http.Error(w, "could not get messages", http.StatusInternalServerError)
		return
	}

	for _, message := range messages {
		fmt.Fprintf(w, message.Message)
	}
}
