package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"sort"
	"sync"

	"github.com/jeschu/autobahn/openapi"
	"golang.org/x/net/context"
)

func main() {

	cfg := openapi.NewConfiguration()
	cfg.Debug = false
	client := openapi.NewAPIClient(cfg)

	api := client.DefaultApi

	ctx := context.Background()

	roads, response, err := api.ListAutobahnen(ctx).Execute()
	if err != nil {
		log.Fatal(err)
	}
	if response.StatusCode != 200 {
		log.Fatalf("ListAutobahnen -> %s\n", response.Status)
	}

	infos := make(map[string]*AutobahnInfo, 0)
	wg := sync.WaitGroup{}
	wg.Add(len(roads.Roads))
	for _, road := range roads.Roads {
		autobahnInfo := &AutobahnInfo{Name: road}
		infos[road] = autobahnInfo
		go func(road string, ai *AutobahnInfo) {
			defer wg.Done()
			road = url.QueryEscape(road)

			stations, response, err := api.ListChargingStations(ctx, road).Execute()
			if err != nil {
				log.Printf("unable to ListChargingStations: %+v", err)
			}
			if response.StatusCode != http.StatusOK {
				log.Printf("ListChargingStations -> %s", response.Status)
			}
			ai.ChargingStations = stations.ElectricChargingStation

			closures, response, err := api.ListClosures(ctx, road).Execute()
			if err != nil {
				log.Printf("unable to ListClosures: %+v", err)
			}
			if response.StatusCode != http.StatusOK {
				log.Printf("ListClosures -> %s", response.Status)
			}
			ai.Closures = closures.Closure

			webcams, response, err := api.ListWebcams(ctx, road).Execute()
			if err != nil {
				log.Printf("unable to ListWebcams: %+v", err)
			}
			if response.StatusCode != http.StatusOK {
				log.Printf("ListWebcams -> %s", response.Status)
			}
			ai.Webcams = webcams.Webcam

			parkingLorries, response, err := api.ListParkingLorries(ctx, road).Execute()
			if err != nil {
				log.Printf("unable to ListParkingLorries: %+v", err)
			}
			if response.StatusCode != http.StatusOK {
				log.Printf("ListParkingLorries -> %s", response.Status)
			}
			ai.ParkingLorries = parkingLorries.ParkingLorry

			roadworks, response, err := api.ListRoadworks(ctx, road).Execute()
			if err != nil {
				log.Printf("unable to ListRoadworks: %+v", err)
			}
			if response.StatusCode != http.StatusOK {
				log.Printf("ListRoadworks -> %s", response.Status)
			}
			ai.Roadworks = roadworks.Roadworks

			warnings, response, err := api.ListWarnings(ctx, road).Execute()
			if err != nil {
				log.Printf("unable to ListWarnings: %+v", err)
			}
			if response.StatusCode != http.StatusOK {
				log.Printf("ListWarnings -> %s", response.Status)
			}
			ai.Warnings = warnings.Warning

		}(road, autobahnInfo)
	}
	wg.Wait()

	autobahnen := make([]string, 0)
	for _, info := range infos {
		autobahnen = append(autobahnen, info.Name)
	}
	sort.Strings(autobahnen)
	for _, autobahn := range autobahnen {
		info := infos[autobahn]
		roadworks := info.Roadworks
		if len(roadworks) > 0 {
			fmt.Printf("%s:", info.Name)
			for _, roadwork := range roadworks {
				if len(roadwork.Description) > 0 {
					fmt.Printf("\n  -> %s", roadwork.Description[0])
					if len(roadwork.Description) > 1 {
						for _, description := range roadwork.Description[1:] {
							fmt.Printf("\n     %s", description)
						}
					} else {
						fmt.Println()
					}
					fmt.Println()
				}
			}
		}
	}
	fmt.Println()
}

type AutobahnInfo struct {
	Name             string
	ChargingStations []openapi.ElectricChargingStation
	Closures         []openapi.Closure
	ParkingLorries   []openapi.ParkingLorry
	Roadworks        []openapi.Roadwork
	Warnings         []openapi.Warning
	Webcams          []openapi.Webcam
}

func (ai *AutobahnInfo) PrintInfo() string {
	return fmt.Sprintf(
		"%s[%d Ladestationen]",
		ai.Name,
		len(ai.ChargingStations),
	)
}
