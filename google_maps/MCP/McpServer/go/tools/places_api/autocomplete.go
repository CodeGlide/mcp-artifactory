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

func AutocompleteHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["input"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("input=%v", val))
		}
		if val, ok := args["sessiontoken"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("sessiontoken=%v", val))
		}
		if val, ok := args["components"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("components=%v", val))
		}
		if val, ok := args["strictbounds"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("strictbounds=%v", val))
		}
		if val, ok := args["offset"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("offset=%v", val))
		}
		if val, ok := args["origin"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("origin=%v", val))
		}
		if val, ok := args["location"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("location=%v", val))
		}
		if val, ok := args["locationbias"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("locationbias=%v", val))
		}
		if val, ok := args["locationrestriction"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("locationrestriction=%v", val))
		}
		if val, ok := args["radius"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("radius=%v", val))
		}
		if val, ok := args["types"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("types=%v", val))
		}
		if val, ok := args["language"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("language=%v", val))
		}
		if val, ok := args["region"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("region=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/maps/api/place/autocomplete/json%s", cfg.BaseURL, queryString)
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
		var result models.PlacesAutocompleteResponse
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

func CreateAutocompleteTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_maps_api_place_autocomplete_json",
		mcp.WithDescription(""),
		mcp.WithString("input", mcp.Required(), mcp.Description("The text string on which to search. The Place Autocomplete service will return candidate matches based on this string and order results based on their perceived relevance.\n")),
		mcp.WithString("sessiontoken", mcp.Description("A random string which identifies an autocomplete [session](https://developers.google.com/maps/documentation/places/web-service/details#session_tokens) for billing purposes.\n\nThe session begins when the user starts typing a query, and concludes when they select a place and a call to Place Details is made. Each session can have multiple queries, followed by one place selection. The API key(s) used for each request within a session must belong to the same Google Cloud Console project. Once a session has concluded, the token is no longer valid; your app must generate a fresh token for each session. If the `sessiontoken` parameter is omitted, or if you reuse a session token, the session is charged as if no session token was provided (each request is billed separately).\n\nWe recommend the following guidelines:\n\n- Use session tokens for all autocomplete sessions.\n- Generate a fresh token for each session. Using a version 4 UUID is recommended.\n- Ensure that the API key(s) used for all Place Autocomplete and Place Details requests within a session belong to the same Cloud Console project.\n- Be sure to pass a unique session token for each new session. Using the same token for more than one session will result in each request being billed individually.\n")),
		mcp.WithString("components", mcp.Description("A grouping of places to which you would like to restrict your results. Currently, you can use components to filter by up to 5 countries. Countries must be passed as a two character, ISO 3166-1 Alpha-2 compatible country code. For example: `components=country:fr` would restrict your results to places within France. Multiple countries must be passed as multiple `country:XX` filters, with the pipe character `|` as a separator. For example: `components=country:us|country:pr|country:vi|country:gu|country:mp` would restrict your results to places within the United States and its unincorporated organized territories.\n<div class=\"note\"><strong>Note:</strong> If you receive unexpected results with a country code, verify that you are using a code which includes the countries, dependent territories, and special areas of geographical interest you intend.  You can find code information at <a href=\"https://en.wikipedia.org/wiki/List_of_ISO_3166_country_codes\" target=\"blank\" class=\"external\">Wikipedia: List of ISO 3166 country codes</a> or the <a href=\"https://www.iso.org/obp/ui/#search\" target=\"blank\" class=\"external\">ISO Online Browsing Platform</a>.</div>\n")),
		mcp.WithBoolean("strictbounds", mcp.Description("Returns only those places that are strictly within the region defined by `location` and `radius`. This is a restriction, rather than a bias, meaning that results outside this region will not be returned even if they match the user input.\n")),
		mcp.WithString("offset", mcp.Description("The position, in the input term, of the last character that the service uses to match predictions. For example, if the input is `Google` and the offset is 3, the service will match on `Goo`. The string determined by the offset is matched against the first word in the input term only. For example, if the input term is `Google abc` and the offset is 3, the service will attempt to match against `Goo abc`. If no offset is supplied, the service will use the whole term. The offset should generally be set to the position of the text caret.\n")),
		mcp.WithString("origin", mcp.Description("The origin point from which to calculate straight-line distance to the destination (returned as `distance_meters`). If this value is omitted, straight-line distance will not be returned. Must be specified as `latitude,longitude`.\n")),
		mcp.WithString("location", mcp.Description("The point around which to retrieve place information. This must be specified as `latitude,longitude`. The `radius` parameter must also be provided when specifying a location. If `radius` is not provided, the `location` parameter is ignored.\n\n<div class=\"note\">When using the Text Search API, the `location` parameter may be overriden if the `query` contains an explicit location such as `Market in Barcelona`.</div>\n")),
		mcp.WithString("locationbias", mcp.Description("Prefer results in a specified area, by specifying either a radius plus lat/lng, or two lat/lng pairs representing the points of a rectangle. If this parameter is not specified, the API uses IP address biasing by default.\n- IP bias: Instructs the API to use IP address biasing. Pass the string `ipbias` (this option has no additional parameters).\n- Circular: A string specifying radius in meters, plus lat/lng in decimal degrees. Use the following format: `circle:radius@lat,lng`.\n- Rectangular: A string specifying two lat/lng pairs in decimal degrees, representing the south/west and north/east points of a rectangle. Use the following format:`rectangle:south,west|north,east`. Note that east/west values are wrapped to the range -180, 180, and north/south values are clamped to the range -90, 90.\n")),
		mcp.WithString("locationrestriction", mcp.Description("Restrict results to a specified area, by specifying either a radius plus lat/lng, or two lat/lng pairs representing the points of a rectangle.\n- Circular: A string specifying radius in meters, plus lat/lng in decimal degrees. Use the following format: `circle:radius@lat,lng`.\n- Rectangular: A string specifying two lat/lng pairs in decimal degrees, representing the south/west and north/east points of a rectangle. Use the following format:`rectangle:south,west|north,east`. Note that east/west values are wrapped to the range -180, 180, and north/south values are clamped to the range -90, 90.\n")),
		mcp.WithString("radius", mcp.Required(), mcp.Description("Defines the distance (in meters) within which to return place results. You may bias results to a specified circle by passing a `location` and a `radius` parameter. Doing so instructs the Places service to _prefer_ showing results within that circle; results outside of the defined area may still be displayed.\n\nThe radius will automatically be clamped to a maximum value depending on the type of search and other parameters.\n\n* Autocomplete: 50,000 meters\n* Nearby Search: \n  * with `keyword` or `name`: 50,000 meters\n  * without `keyword` or `name`\n    * Up to 50,000 meters, adjusted dynamically based on area density, independent of `rankby` parameter.\n    * When using `rankby=distance`, the radius parameter will not be accepted, and will result in an `INVALID_REQUEST`.\n* Query Autocomplete: 50,000 meters\n* Text Search: 50,000 meters\n")),
		mcp.WithString("types", mcp.Description("You can restrict results from a Place Autocomplete request to be of a certain type by passing the `types` parameter. This parameter specifies a type or a type collection, as listed in [Place Types](/maps/documentation/places/web-service/supported_types). If nothing is specified, all types are returned.\n\nFor the value of the `types` parameter you can specify either:\n* Up to five values from [Table 1](/maps/documentation/places/web-service/supported_types#table1) or [Table 2](/maps/documentation/places/web-service/supported_types#table2). For multiple values, separate each value with a `|` (vertical bar). For example:\n\n  `types=book_store|cafe`\n\n* Any supported filter in [Table 3](/maps/documentation/places/web-service/supported_types#table3). You can safely mix the `geocode` and `establishment` types. You cannot mix type collections (`address`, `(cities)` or `(regions)`) with any other type, or an error occurs.\n\nThe request will be rejected with an `INVALID_REQUEST` error if:\n\n* More than five types are specified.\n* Any unrecognized types are present.\n* Any types from in [Table 1](/maps/documentation/places/web-service/supported_types#table1) or [Table 2](/maps/documentation/places/web-service/supported_types#table2) are mixed with any of the filters in [Table 3](/maps/documentation/places/web-service/supported_types#table3).\n")),
		mcp.WithString("language", mcp.Description("The language in which to return results.\n\n* See the [list of supported languages](https://developers.google.com/maps/faq#languagesupport). Google often updates the supported languages, so this list may not be exhaustive.\n* If `language` is not supplied, the API attempts to use the preferred language as specified in the `Accept-Language` header.\n* The API does its best to provide a street address that is readable for both the user and locals. To achieve that goal, it returns street addresses in the local language, transliterated to a script readable by the user if necessary, observing the preferred language. All other addresses are returned in the preferred language. Address components are all returned in the same language, which is chosen from the first component.\n* If a name is not available in the preferred language, the API uses the closest match.\n* The preferred language has a small influence on the set of results that the API chooses to return, and the order in which they are returned. The geocoder interprets abbreviations differently depending on language, such as the abbreviations for street types, or synonyms that may be valid in one language but not in another. For example, _utca_ and _tér_ are synonyms for street in Hungarian.")),
		mcp.WithString("region", mcp.Description("The region code, specified as a [ccTLD (\"top-level domain\")](https://en.wikipedia.org/wiki/List_of_Internet_top-level_domains#Country_code_top-level_domains) two-character value. Most ccTLD codes are identical to ISO 3166-1 codes, with some notable exceptions. For example, the United Kingdom's ccTLD is \"uk\" (.co.uk) while its ISO 3166-1 code is \"gb\" (technically for the entity of \"The United Kingdom of Great Britain and Northern Ireland\").")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    AutocompleteHandler(cfg),
	}
}
