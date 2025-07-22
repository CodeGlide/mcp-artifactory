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

func GeocodeHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["address"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("address=%v", val))
		}
		if val, ok := args["bounds"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("bounds=%v", val))
		}
		if val, ok := args["components"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("components=%v", val))
		}
		if val, ok := args["latlng"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("latlng=%v", val))
		}
		if val, ok := args["location_type"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("location_type=%v", val))
		}
		if val, ok := args["place_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("place_id=%v", val))
		}
		if val, ok := args["result_type"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("result_type=%v", val))
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
		url := fmt.Sprintf("%s/maps/api/geocode/json%s", cfg.BaseURL, queryString)
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
		var result models.GeocodingResponse
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

func CreateGeocodeTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_maps_api_geocode_json",
		mcp.WithDescription(""),
		mcp.WithString("address", mcp.Description("The street address or plus code that you want to geocode. Specify addresses in accordance with the format used by the national postal service of the country concerned. Additional address elements such as business names and unit, suite or floor numbers should be avoided. Street address elements should be delimited by spaces (shown here as url-escaped to `%20`):\n\n```\naddress=24%20Sussex%20Drive%20Ottawa%20ON\n```\n\nFormat plus codes as shown here (plus signs are url-escaped to `%2B` and spaces are url-escaped to `%20`):\n- global code is a 4 character area code and 6 character or longer local code (`849VCWC8+R9` is `849VCWC8%2BR9`).\n- compound code is a 6 character or longer local code with an explicit location (`CWC8+R9 Mountain View, CA, USA` is `CWC8%2BR9%20Mountain%20View%20CA%20USA`).\n\n<div class=\"note\">Note: At least one of `address` or `components` is required.</div>")),
		mcp.WithArray("bounds", mcp.Description("The bounding box of the viewport within which to bias geocode results more prominently. This parameter will only influence, not fully restrict, results from the geocoder.")),
		mcp.WithArray("components", mcp.Description("A components filter with elements separated by a pipe (|). The components filter is also accepted as an optional parameter if an address is provided. Each element in the components filter consists of a `component:value` pair, and fully restricts the results from the geocoder.\n\nThe components that can be filtered include:\n\n* `postal_code` matches `postal_code` and `postal_code_prefix`.\n* `country` matches a country name or a two letter ISO 3166-1 country code. The API follows the ISO standard for defining countries, and the filtering works best when using the corresponding ISO code of the country.\n  <aside class=\"note\"><strong>Note</strong>: If you receive unexpected results with a country code, verify that you are using a code which includes the countries, dependent territories, and special areas of geographical interest you intend. You can find code information at Wikipedia: List of ISO 3166 country codes or the ISO Online Browsing Platform.</aside>\n  \nThe following components may be used to influence results, but will not be enforced:\n\n* `route` matches the long or short name of a route.\n* `locality` matches against `locality` and `sublocality` types.\n* `administrative_area` matches all the `administrative_area` levels.\n  \nNotes about component filtering:\n\n* Do not repeat these component filters in requests, or the API will return `INVALID_REQUEST`: \n  * `country`\n  * `postal_code`\n  * `route`\n* If the request contains repeated component filters, the API evaluates those filters as an AND, not an OR.\n* Results are consistent with Google Maps, which occasionally yields unexpected `ZERO_RESULTS` responses. Using Place Autocomplete may provide better results in some use cases. To learn more, see [this FAQ](https://developers..google.com/maps/documentation/geocoding/faq#trbl_component_filtering).\n* For each address component, either specify it in the address parameter or in a components filter, but not both. Specifying the same values in both may result in `ZERO_RESULTS`.\n\n<div class=\"note\">Note: At least one of `address` or `components` is required.</div>")),
		mcp.WithString("latlng", mcp.Description("The street address that you want to geocode, in the format used by the national postal service of the country concerned. Additional address elements such as business names and unit, suite or floor numbers should be avoided.")),
		mcp.WithArray("location_type", mcp.Description("A filter of one or more location types, separated by a pipe (`|`). If the parameter contains multiple location types, the API returns all addresses that match any of the types. A note about processing: The `location_type` parameter does not restrict the search to the specified location type(s). Rather, the `location_type` acts as a post-search filter: the API fetches all results for the specified latlng, then discards those results that do not match the specified location type(s). The following values are supported:\n* `APPROXIMATE` returns only the addresses that are characterized as approximate.\n* `GEOMETRIC_CENTER` returns only geometric centers of a location such as a polyline (for example, a street) or polygon (region).\n* `RANGE_INTERPOLATED` returns only the addresses that reflect an approximation (usually on a road) interpolated between two precise points (such as intersections). An interpolated range generally indicates that rooftop geocodes are unavailable for a street address.\n* `ROOFTOP` returns only the addresses for which Google has location information accurate down to street address precision.")),
		mcp.WithString("place_id", mcp.Description("A textual identifier that uniquely identifies a place, returned from a [Place Search](https://developers.google.com/maps/documentation/places/web-service/search).\nFor more information about place IDs, see the [place ID overview](https://developers.google.com/maps/documentation/places/web-service/place-id).\n")),
		mcp.WithArray("result_type", mcp.Description("A filter of one or more address types, separated by a pipe (|). If the parameter contains multiple address types, the API returns all addresses that match any of the types. A note about processing: The `result_type` parameter does not restrict the search to the specified address type(s). Rather, the `result_type` acts as a post-search filter: the API fetches all results for the specified `latlng`, then discards those results that do not match the specified address type(s).The following values are supported:\n* `administrative_area_level_1` indicates a first-order civil entity below the country level. Within the United States, these administrative levels are states. Not all nations exhibit these administrative levels. In most cases, administrative_area_level_1   * `short` names will closely match ISO 3166-2 subdivisions and other widely circulated lists; however this is not guaranteed as our geocoding results are based on a variety of signals and location data.\n* `administrative_area_level_2` indicates a second-order civil entity below the country level. Within the United States, these administrative levels are counties. Not all nations exhibit these administrative levels.\n* `administrative_area_level_3` indicates a third-order civil entity below the country level. This type indicates a minor civil division. Not all nations exhibit these administrative levels.\n* `administrative_area_level_4` indicates a fourth-order civil entity below the country level. This type indicates a minor civil division. Not all nations exhibit these administrative levels.\n* `administrative_area_level_5` indicates a fifth-order civil entity below the country level. This type indicates a minor civil division. Not all nations exhibit these administrative levels.\n* `airport` indicates an airport.\n* `colloquial_area` indicates a commonly-used alternative name for the entity.\n* `country` indicates the national political entity, and is typically the highest order type returned by the Geocoder.\n* `intersection` indicates a major intersection, usually of two major roads.\n* `locality` indicates an incorporated city or town political entity.\n* `natural_feature` indicates a prominent natural feature.\n* `neighborhood` indicates a named neighborhood\n* `park` indicates a named park.\n* `plus_code` indicates an encoded location reference, derived from latitude and longitude. Plus codes can be used as a replacement for street addresses in places where they do not exist (where buildings are not numbered or streets are not named). See [https://plus.codes/](https://plus.codes/) for details.\n* `point_of_interest` indicates a named point of interest. Typically, these \"POI\"s are prominent local entities that don't easily fit in another category, such as \"Empire State Building\" or \"Eiffel Tower\".\n* `political` indicates a political entity. Usually, this type indicates a polygon of some civil administration.\n* `postal_code` indicates a postal code as used to address postal mail within the country.\n* `premise` indicates a named location, usually a building or collection of buildings with a common name\n* `route` indicates a named route (such as \"US 101\").\n* `street_address` indicates a precise street address.\n* `sublocality` indicates a first-order civil entity below a locality. For some locations may receive one of the additional types: `sublocality_level_1` to `sublocality_level_5`. Each sublocality level is a civil entity. Larger numbers indicate a smaller geographic area.\n* `subpremise` indicates a first-order entity below a named location, usually a singular building within a collection of buildings with a common name")),
		mcp.WithString("language", mcp.Description("The language in which to return results.\n\n* See the [list of supported languages](https://developers.google.com/maps/faq#languagesupport). Google often updates the supported languages, so this list may not be exhaustive.\n* If `language` is not supplied, the API attempts to use the preferred language as specified in the `Accept-Language` header.\n* The API does its best to provide a street address that is readable for both the user and locals. To achieve that goal, it returns street addresses in the local language, transliterated to a script readable by the user if necessary, observing the preferred language. All other addresses are returned in the preferred language. Address components are all returned in the same language, which is chosen from the first component.\n* If a name is not available in the preferred language, the API uses the closest match.\n* The preferred language has a small influence on the set of results that the API chooses to return, and the order in which they are returned. The geocoder interprets abbreviations differently depending on language, such as the abbreviations for street types, or synonyms that may be valid in one language but not in another. For example, _utca_ and _tér_ are synonyms for street in Hungarian.")),
		mcp.WithString("region", mcp.Description("The region code, specified as a [ccTLD (\"top-level domain\")](https://en.wikipedia.org/wiki/List_of_Internet_top-level_domains#Country_code_top-level_domains) two-character value. Most ccTLD codes are identical to ISO 3166-1 codes, with some notable exceptions. For example, the United Kingdom's ccTLD is \"uk\" (.co.uk) while its ISO 3166-1 code is \"gb\" (technically for the entity of \"The United Kingdom of Great Britain and Northern Ireland\").")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GeocodeHandler(cfg),
	}
}
