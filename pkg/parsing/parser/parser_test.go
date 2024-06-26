package parser

import (
	"fmt"
	"testing"
)

func Test_parser_Run(t *testing.T) {

	t.Run("OK", func(t *testing.T) {
		parser := NewEventParser(Sources)
		results := parser.Run()
		for result := range results {
			fmt.Println(result)
		}
	})

}

// func Test_Import(t *testing.T) {
// 	file, err := os.Open(os.Getenv("PARSER_DATA"))
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	bytes, err := io.ReadAll(file)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	var data []Result
// 	err = json.Unmarshal(bytes, &data)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	db, err := pgxpool.New(context.Background(), os.Getenv("TEST_POSTGRES_URL"))
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	eventStorage := postgres.NewStorage(db)
// 	locationStorage := postgres.NewStorage(db)
//
// 	for _, i := range data {
// 		fmt.Printf("Creating a New Event: %s\n", i.Event.Name)
//
// 		var err error
// 		var place location.Location
// 		place, err = locationStorage.FindByName(i.Location.Name)
// 		if err != nil {
// 			if errors.Is(err, postgres.ErrLocationNotFound) {
// 				fmt.Printf("Location not found: name=%s\n", i.Location.Name)
//
// 				i.Location.Type = location.Unknown
// 				place, err = locationStorage.Save(i.Location)
// 				if err != nil {
// 					t.Fatal(err)
// 				}
// 				fmt.Printf("Created a New Location: id=%d\n", place.ID)
//
// 			} else {
// 				t.Fatal(err)
// 			}
// 		}
//
// 		i.Event.LocationID = place.ID
// 		i.Event.Price = make(event.Price)
//
// 		newEvent, err := eventStorage.Save(i.Event)
// 		if err != nil {
// 			t.Fatal(err)
// 		}
//
// 		fmt.Printf("Created a New Event [id:%d, name:%s]\n", newEvent.ID, newEvent.Name)
// 	}
// }
