package store

import (
	"context"
	"spider/zhenai/parser"

	"gopkg.in/olivere/elastic"
)

// func ItemSaver(index string, etype string) (chan string, error) {
// 	client, err := elastic.NewClient(
// 		// Must turn off sniff in docker
// 		elastic.SetSniff(true))

// 	if err != nil {
// 		return nil, err
// 	}

// 	out := make(chan string)
// 	go func() {
// 		itemCount := 0
// 		for {
// 			item := <-out
// 			fmt.Printf("Item Saver: got item "+
// 				"#%d: %v", itemCount, item)
// 			log.Printf("Item Saver: got item "+
// 				"#%d: %v", itemCount, item)
// 			itemCount++

// 			err := Save(client, index, etype, item)
// 			if err != nil {
// 				log.Printf("Item Saver: error "+
// 					"saving item %v: %v",
// 					item, err)
// 			}
// 		}
// 	}()
// 	return out, nil
// }

func Save(client *elastic.Client, index string, etype string, item parser.Person) error {
	indexService := client.Index().
		Index(index).
		Type(etype).
		BodyJson(item)
	_, err := indexService.
		Do(context.Background())

	return err
}
