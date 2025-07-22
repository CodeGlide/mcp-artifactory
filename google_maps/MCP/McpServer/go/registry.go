package main

import (
	"github.com/google-maps-platform/mcp-server/config"
	"github.com/google-maps-platform/mcp-server/models"
	tools_places_api "github.com/google-maps-platform/mcp-server/tools/places_api"
	tools_roads_api "github.com/google-maps-platform/mcp-server/tools/roads_api"
	tools_time_zone_api "github.com/google-maps-platform/mcp-server/tools/time_zone_api"
	tools_geolocation_api "github.com/google-maps-platform/mcp-server/tools/geolocation_api"
	tools_distance_matrix_api "github.com/google-maps-platform/mcp-server/tools/distance_matrix_api"
	tools_street_view_api "github.com/google-maps-platform/mcp-server/tools/street_view_api"
	tools_elevation_api "github.com/google-maps-platform/mcp-server/tools/elevation_api"
	tools_geocoding_api "github.com/google-maps-platform/mcp-server/tools/geocoding_api"
	tools_directions_api "github.com/google-maps-platform/mcp-server/tools/directions_api"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_places_api.CreateNearbysearchTool(cfg),
		tools_roads_api.CreateNearestroadsTool(cfg),
		tools_places_api.CreateAutocompleteTool(cfg),
		tools_places_api.CreateTextsearchTool(cfg),
		tools_time_zone_api.CreateTimezoneTool(cfg),
		tools_places_api.CreateFindplacefromtextTool(cfg),
		tools_geolocation_api.CreateGeolocateTool(cfg),
		tools_places_api.CreatePlacedetailsTool(cfg),
		tools_distance_matrix_api.CreateDistancematrixTool(cfg),
		tools_street_view_api.CreateStreetviewmetadataTool(cfg),
		tools_places_api.CreateQueryautocompleteTool(cfg),
		tools_elevation_api.CreateElevationTool(cfg),
		tools_geocoding_api.CreateGeocodeTool(cfg),
		tools_directions_api.CreateDirectionsTool(cfg),
		tools_roads_api.CreateSnaptoroadsTool(cfg),
	}
}
