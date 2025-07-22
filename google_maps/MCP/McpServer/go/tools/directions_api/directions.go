package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/google-maps-platform/mcp-server/config"
	"github.com/google-maps-platform/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func DirectionsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["arrival_time"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("arrival_time=%v", val))
		}
		if val, ok := args["departure_time"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("departure_time=%v", val))
		}
		if val, ok := args["alternatives"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("alternatives=%v", val))
		}
		if val, ok := args["avoid"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("avoid=%v", val))
		}
		if val, ok := args["destination"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("destination=%v", val))
		}
		if val, ok := args["origin"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("origin=%v", val))
		}
		if val, ok := args["units"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("units=%v", val))
		}
		if val, ok := args["waypoints"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("waypoints=%v", val))
		}
		if val, ok := args["language"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("language=%v", val))
		}
		if val, ok := args["mode"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("mode=%v", val))
		}
		if val, ok := args["region"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("region=%v", val))
		}
		if val, ok := args["traffic_model"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("traffic_model=%v", val))
		}
		if val, ok := args["transit_mode"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("transit_mode=%v", val))
		}
		if val, ok := args["transit_routing_preference"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("transit_routing_preference=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/maps/api/directions/json%s", cfg.BaseURL, queryString)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication header based on auth type
		if cfg.APIKey != "" {
			req.Header.Set("key", cfg.APIKey)
		}
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.DirectionsResponse
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateDirectionsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_maps_api_directions_json",
		mcp.WithDescription(""),
		mcp.WithString("arrival_time", mcp.Description("Specifies the desired time of arrival for transit directions, in seconds since midnight, January 1, 1970 UTC. You can specify either `departure_time` or `arrival_time`, but not both. Note that `arrival_time` must be specified as an integer.\n")),
		mcp.WithString("departure_time", mcp.Description("Specifies the desired time of departure. You can specify the time as an integer in seconds since midnight, January 1, 1970 UTC. If a `departure_time` later than 9999-12-31T23:59:59.999999999Z is specified, the API will fall back the `departure_time` to 9999-12-31T23:59:59.999999999Z. Alternatively, you can specify a value of now, which sets the departure time to the current time (correct to the nearest second). The departure time may be specified in two cases:\n* For requests where the travel mode is transit: You can optionally specify one of `departure_time` or `arrival_time`. If neither time is specified, the `departure_time` defaults to now (that is, the departure time defaults to the current time).\n* For requests where the travel mode is driving: You can specify the `departure_time` to receive a route and trip duration (response field: duration_in_traffic) that take traffic conditions into account. The `departure_time` must be set to the current time or some time in the future. It cannot be in the past.\n\n<div class=\"note\">Note: If departure time is not specified, choice of route and duration are based on road network and average time-independent traffic conditions. Results for a given request may vary over time due to changes in the road network, updated average traffic conditions, and the distributed nature of the service. Results may also vary between nearly-equivalent routes at any time or frequency.</div>\n<div class=\"note\">Note: Distance Matrix requests specifying `departure_time` when `mode=driving` are limited to a maximum of 100 elements per request. The number of origins times the number of destinations defines the number of elements.</div>\n")),
		mcp.WithBoolean("alternatives", mcp.Description("If set to `true`, specifies that the Directions service may provide more than one route alternative in the response. Note that providing route alternatives may increase the response time from the server. This is only available for requests without intermediate waypoints. For more information, see the [guide to waypoints](https://developers.google.com/maps/documentation/directions/get-directions#Waypoints).\n")),
		mcp.WithString("avoid", mcp.Description("Indicates that the calculated route(s) should avoid the indicated features. This parameter supports the following arguments:\n* `tolls` indicates that the calculated route should avoid toll roads/bridges.\n* `highways` indicates that the calculated route should avoid highways.\n* `ferries` indicates that the calculated route should avoid ferries.\n* `indoor` indicates that the calculated route should avoid indoor steps for walking and transit directions.\n\nIt's possible to request a route that avoids any combination of tolls, highways and ferries by passing multiple restrictions to the avoid parameter. For example: \n\n```\navoid=tolls|highways|ferries.\n```\n")),
		mcp.WithString("destination", mcp.Required(), mcp.Description("The place ID, address, or textual latitude/longitude value to which you wish to calculate directions. The options for the destination parameter are the same as for the origin parameter.")),
		mcp.WithString("origin", mcp.Required(), mcp.Description("The place ID, address, or textual latitude/longitude value from which you wish to calculate directions.\n* Place IDs must be prefixed with `place_id:`. You can retrieve place IDs from the Geocoding API and the Places API (including Place Autocomplete). For an example using place IDs from Place Autocomplete, see [Place Autocomplete and Directions](https://developers.google.com/maps/documentation/javascript/examples/places-autocomplete-directions). For more about place IDs, see the [Place ID overview](https://developers.google.com/maps/documentation/places/web-service/place-id).\n  \n  ```\n  origin=place_id:ChIJ3S-JXmauEmsRUcIaWtf4MzE\n  ```\n  \n* If you pass an address, the Directions service geocodes the string and converts it to a latitude/longitude coordinate to calculate directions. This coordinate may be different from that returned by the Geocoding API, for example a building entrance rather than its center.\n  \n  ```\n  origin=24+Sussex+Drive+Ottawa+ON\n  ```\n  \n  Using place IDs is preferred over using addresses or latitude/longitude coordinates. Using coordinates will always result in the point being snapped to the road nearest to those coordinates - which may not be an access point to the property, or even a road that will quickly or safely lead to the destination.\n* If you pass coordinates, the point will snap to the nearest road. Passing a place ID is preferred. If you do pass coordinates, ensure that no space exists between the latitude and longitude values.\n  \n  ```\n  origin=41.43206,-81.38992\n  ```\n\n* Plus codes must be formatted as a global code or a compound code. Format plus codes as shown here (plus signs are url-escaped to `%2B` and spaces are url-escaped to `%20`). \n  \n  * **Global code** is a 4 character area code and 6 character or longer local code (849VCWC8+R9 is `849VCWC8%2BR9`). \n  * **Compound code** is a 6 character or longer local code with an explicit location (CWC8+R9 Mountain View, CA, USA is `CWC8%2BR9%20Mountain%20View%20CA%20USA`).\n\n<div class=\"note\">Note: For efficiency and accuracy, use place ID's when possible. These ID's are uniquely explicit like a lat/lng value pair and provide geocoding benefits for routing such as access points and traffic variables. Unlike an address, ID's do not require the service to perform a search or an intermediate request for place details; therefore, performance is better.</div>\n")),
		mcp.WithString("units", mcp.Description("Specifies the unit system to use when displaying results.\n\nDirections results contain text within distance fields that may be displayed to the user to indicate the distance of a particular \"step\" of the route. By default, this text uses the unit system of the origin's country or region.\n\nFor example, a route from \"Chicago, IL\" to \"Toronto, ONT\" will display results in miles, while the reverse route will display results in kilometers. You may override this unit system by setting one explicitly within the request's units parameter, passing one of the following values:\n\n* `metric` specifies usage of the metric system. Textual distances are returned using kilometers and meters.\n* `imperial` specifies usage of the Imperial (English) system. Textual distances are returned using miles and feet.\n\n<div class=\"note\">Note: this unit system setting only affects the text displayed within distance fields. The distance fields also contain values which are always expressed in meters.</div>\n")),
		mcp.WithString("waypoints", mcp.Description("<div class=\"caution\">Caution: Requests using more than 10 waypoints (between 11 and 25), or waypoint optimization, are billed at a higher rate. <a href=\"https://developers.google.com/maps/billing-and-pricing/pricing#directions-advanced\">Learn more about billing</a> for Google Maps Platform products.</div>\n\nSpecifies an array of intermediate locations to include along the route between the origin and destination points as pass through or stopover locations. Waypoints alter a route by directing it through the specified location(s). The API supports waypoints for these travel modes: driving, walking and bicycling; not transit.   You can supply one or more locations separated by the pipe character (`|` or `%7C`), in the form of a place ID, an address, or latitude/longitude coordinates. By default, the Directions service calculates a route using the waypoints in the order they are given. The precedence for parsing the value of the waypoint is place ID, latitude/longitude coordinates, then address.\n* If you pass a place ID, you must prefix it with `place_id:`. You can retrieve place IDs from the Geocoding API and the Places API (including Place Autocomplete). For an example using place IDs from Place Autocomplete, see [Place Autocomplete and Directions](/maps/documentation/javascript/examples/places-autocomplete-directions). For more about place IDs, see the [Place ID overview](/maps/documentation/places/web-service/place-id).\n  <div class=\"note\">For efficiency and accuracy, use place ID's when possible. These ID's are uniquely explicit like a lat/lng value pair and provide geocoding benefits for routing such as access points and traffic variables. Unlike an address, ID's do not require the service to perform a search or an intermediate request for place details; therefore, performance is better.</div>\n* If you pass latitude/longitude coordinates, the values go directly to the front-end server to calculate directions without geocoding. The points are snapped to roads and might not provide the accuracy your app needs. Use coordinates when you are confident the values truly specify the points your app needs for routing without regard to possible access points or additional geocoding details. Ensure that a comma (`%2C`) and not a space (`%20`) separates the latitude and longitude values.\n* If you pass an address, the Directions service will geocode the string and convert it into latitude/longitude coordinates to calculate directions. If the address value is ambiguous, the value might evoke a search to disambiguate from similar addresses. For example, \"1st Street\" could be a complete value or a partial value for \"1st street NE\" or \"1st St SE\". This result may be different from that returned by the Geocoding API. You can avoid possible misinterpretations using place IDs.\n* Alternatively, you can supply an encoded set of points using the [Encoded Polyline Algorithm](https://developers.google.com/maps/documentation/utilities/polylinealgorithm). You will find an encoded set is useful for a large number of waypoints, because the URL is significantly shorter. All web services have a URL limit of 8192 characters.\n  * Encoded polylines must be prefixed with `enc:` and followed by a colon (`:`). For example: `waypoints=enc:gfo}EtohhU:`.\n  * You can also include multiple encoded polylines, separated by the pipe character (`|`). For example, `waypoints=via:enc:wc~oAwquwMdlTxiKtqLyiK:|enc:c~vnAamswMvlTor@tjGi}L:| via:enc:udymA{~bxM:`\n\n##### Influence routes with stopover and pass through points\n\nFor each waypoint in the request, the directions response appends an entry to the `legs` array to provide the details for stopovers on that leg of the journey.\n\nIf you'd like to influence the route using waypoints without adding a stopover, add the prefix `via:` to the waypoint. Waypoints prefixed with `via:` will not add an entry to the `legs` array, but will route the journey through the waypoint.\n\nThe following URL modifies the previous request such that the journey is routed through Lexington without stopping:\n\n```\nhttps://maps.googleapis.com/maps/api/directions/json?\norigin=Boston,MA&destination=Concord,MA\n&waypoints=Charlestown,MA|via:Lexington,MA  \n```\n\nThe `via:` prefix is most effective when creating routes in response to the user dragging the waypoints on the map. Doing so allows the user to see how the final route may look in real-time and helps ensure that waypoints are placed in locations that are accessible to the Directions API.\n\n<div class=\"caution\">Caution: Using the `via:` prefix to avoid stopovers results in directions that are strict in their interpretation of the waypoint. This interpretation may result in severe detours on the route or `ZERO_RESULTS` in the response status code if the Directions API is unable to create directions through that point.</div>\n\n\n##### Optimize your waypoints\n\nBy default, the Directions service calculates a route through the provided waypoints in their given order. Optionally, you may pass `optimize:true` as the first argument within the waypoints parameter to allow the Directions service to optimize the provided route by rearranging the waypoints in a more efficient order. (This optimization is an application of the traveling salesperson problem.) Travel time is the primary factor which is optimized, but other factors such as distance, number of turns and many more may be taken into account when deciding which route is the most efficient. All waypoints must be stopovers for the Directions service to optimize their route.\n\nIf you instruct the Directions service to optimize the order of its waypoints, their order will be returned in the `waypoint_order` field within the routes object. The `waypoint_order` field returns values which are zero-based.\n\nThe following example calculates a road journey from Adelaide, South Australia to each of South Australia's main wine regions using route optimization.\n\n```\nhttps://maps.googleapis.com/maps/api/directions/json?\norigin=Adelaide,SA&destination=Adelaide,SA\n&waypoints=optimize:true|Barossa+Valley,SA|Clare,SA|Connawarra,SA|McLaren+Vale,SA\n```\n\nInspection of the calculated route will indicate that calculation uses waypoints in the following waypoint order:\n\n```\n\"waypoint_order\": [ 3, 2, 0, 1 ]\n```\n\n<div class=\"caution\">Caution: Requests using waypoint optimization are billed at a higher rate. <a href=\"https://developers.google.com/maps/billing-and-pricing/pricing#directions-advanced\">Learn more about how Google Maps Platform products are billed.</a></div>\n")),
		mcp.WithString("language", mcp.Description("The language in which to return results.\n\n* See the [list of supported languages](https://developers.google.com/maps/faq#languagesupport). Google often updates the supported languages, so this list may not be exhaustive.\n* If `language` is not supplied, the API attempts to use the preferred language as specified in the `Accept-Language` header.\n* The API does its best to provide a street address that is readable for both the user and locals. To achieve that goal, it returns street addresses in the local language, transliterated to a script readable by the user if necessary, observing the preferred language. All other addresses are returned in the preferred language. Address components are all returned in the same language, which is chosen from the first component.\n* If a name is not available in the preferred language, the API uses the closest match.\n* The preferred language has a small influence on the set of results that the API chooses to return, and the order in which they are returned. The geocoder interprets abbreviations differently depending on language, such as the abbreviations for street types, or synonyms that may be valid in one language but not in another. For example, _utca_ and _tér_ are synonyms for street in Hungarian.")),
		mcp.WithString("mode", mcp.Description("For the calculation of distances and directions, you may specify the transportation mode to use. By default, `DRIVING` mode is used. By default, directions are calculated as driving directions. The following travel modes are supported:\n\n* `driving` (default) indicates standard driving directions or distance using the road network.\n* `walking` requests walking directions or distance via pedestrian paths & sidewalks (where available).\n* `bicycling` requests bicycling directions or distance via bicycle paths & preferred streets (where available).\n* `transit` requests directions or distance via public transit routes (where available). Transit trips are available for up to 7 days in the past or 100 days in the future. If you set the mode to transit, you can optionally specify either a `departure_time` or an `arrival_time`. If neither time is specified, the `departure_time` defaults to now (that is, the departure time defaults to the current time). You can also optionally include a `transit_mode` and/or a `transit_routing_preference`. \n\n<div class=\"note\">Note: Both walking and bicycling directions may sometimes not include clear pedestrian or bicycling paths, so these directions will return warnings in the returned result which you must display to the user.</div>\n")),
		mcp.WithString("region", mcp.Description("The region code, specified as a [ccTLD (\"top-level domain\")](https://en.wikipedia.org/wiki/List_of_Internet_top-level_domains#Country_code_top-level_domains) two-character value. Most ccTLD codes are identical to ISO 3166-1 codes, with some notable exceptions. For example, the United Kingdom's ccTLD is \"uk\" (.co.uk) while its ISO 3166-1 code is \"gb\" (technically for the entity of \"The United Kingdom of Great Britain and Northern Ireland\").")),
		mcp.WithString("traffic_model", mcp.Description("Specifies the assumptions to use when calculating time in traffic. This setting affects the value returned in the duration_in_traffic field in the response, which contains the predicted time in traffic based on historical averages. The `traffic_model` parameter may only be specified for driving directions where the request includes a `departure_time`. The available values for this parameter are:\n* `best_guess` (default) indicates that the returned duration_in_traffic should be the best estimate of travel time given what is known about both historical traffic conditions and live traffic. Live traffic becomes more important the closer the `departure_time` is to now.\n* `pessimistic` indicates that the returned duration_in_traffic should be longer than the actual travel time on most days, though occasional days with particularly bad traffic conditions may exceed this value.\n* `optimistic` indicates that the returned duration_in_traffic should be shorter than the actual travel time on most days, though occasional days with particularly good traffic conditions may be faster than this value.\nThe default value of `best_guess` will give the most useful predictions for the vast majority of use cases. It is possible the `best_guess` travel time prediction may be shorter than `optimistic`, or alternatively, longer than `pessimistic`, due to the way the `best_guess` prediction model integrates live traffic information.\n")),
		mcp.WithString("transit_mode", mcp.Description("Specifies one or more preferred modes of transit. This parameter may only be specified for transit directions. The parameter supports the following arguments:\n* `bus` indicates that the calculated route should prefer travel by bus.\n* `subway` indicates that the calculated route should prefer travel by subway.\n* `train` indicates that the calculated route should prefer travel by train.\n* `tram` indicates that the calculated route should prefer travel by tram and light rail.\n* `rail` indicates that the calculated route should prefer travel by train, tram, light rail, and subway. This is equivalent to `transit_mode=train|tram|subway`.\n")),
		mcp.WithString("transit_routing_preference", mcp.Description("Specifies preferences for transit routes. Using this parameter, you can bias the options returned, rather than accepting the default best route chosen by the API. This parameter may only be specified for transit directions. The parameter supports the following arguments:\n* `less_walking` indicates that the calculated route should prefer limited amounts of walking.\n* `fewer_transfers` indicates that the calculated route should prefer a limited number of transfers.\n")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    DirectionsHandler(cfg),
	}
}
