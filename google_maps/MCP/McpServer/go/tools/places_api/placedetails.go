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

func PlacedetailsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["place_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("place_id=%v", val))
		}
		if val, ok := args["fields"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fields=%v", val))
		}
		if val, ok := args["sessiontoken"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("sessiontoken=%v", val))
		}
		if val, ok := args["language"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("language=%v", val))
		}
		if val, ok := args["region"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("region=%v", val))
		}
		if val, ok := args["reviews_sort"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("reviews_sort=%v", val))
		}
		if val, ok := args["reviews_no_translations"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("reviews_no_translations=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/maps/api/place/details/json%s", cfg.BaseURL, queryString)
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
		var result models.PlacesDetailsResponse
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

func CreatePlacedetailsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_maps_api_place_details_json",
		mcp.WithDescription(""),
		mcp.WithString("place_id", mcp.Required(), mcp.Description("A textual identifier that uniquely identifies a place, returned from a [Place Search](https://developers.google.com/maps/documentation/places/web-service/search).\nFor more information about place IDs, see the [place ID overview](https://developers.google.com/maps/documentation/places/web-service/place-id).\n")),
		mcp.WithArray("fields", mcp.Description("\n<div class=\"caution\"> Caution: Place Search requests and Place Details requests do not return the same fields. Place Search requests return a subset of the fields that are returned by Place Details requests. If the field you want is not returned by Place Search, you can use Place Search to get a <code>place_id</code>, then use that Place ID to make a Place Details request. For more information on the fields that are unavailable in a Place Search request, see <a href=\"https://developers.google.com/maps/documentation/places/web-service/place-data-fields#places-api-fields-support\">Places API fields support</a>.</div>\n\nUse the fields parameter to specify a comma-separated list of place data types to return. For example: `fields=formatted_address,name,geometry`. Use a forward slash when specifying compound values. For example: `opening_hours/open_now`.\n\nFields are divided into three billing categories: Basic, Contact, and Atmosphere. Basic fields are billed at base rate, and incur no additional charges. Contact and Atmosphere fields are billed at a higher rate. See the [pricing sheet](https://developers.google.com/maps/documentation/places/web-service/usage-and-billing/) for more information. Attributions, `html_attributions`, are always returned with every call, regardless of whether the field has been requested.\n\n**Basic**\n\nThe Basic category includes the following fields: `address_components`, `adr_address`, `business_status`, `formatted_address`, `geometry`, `icon`, `icon_mask_base_uri`, `icon_background_color`, `name`, `permanently_closed` ([deprecated](https://developers.google.com/maps/deprecations)), `photo`, `place_id`, `plus_code`, `type`, `url`, `utc_offset`, `vicinity`, `wheelchair_accessible_entrance`.\n\n**Contact**\n\nThe Contact category includes the following fields: `current_opening_hours`, `formatted_phone_number`, `international_phone_number`, `opening_hours`, `secondary_opening_hours`, `website`\n\n**Atmosphere**\n\nThe Atmosphere category includes the following fields: `curbside_pickup`, `delivery`, `dine_in`, `editorial_summary`, `price_level`, `rating`, `reservable`, `reviews`, `serves_beer`, `serves_breakfast`, `serves_brunch`, `serves_dinner`, `serves_lunch`, `serves_vegetarian_food`, `serves_wine`, `takeout`, `user_ratings_total`.\n")),
		mcp.WithString("sessiontoken", mcp.Description("A random string which identifies an autocomplete [session](https://developers.google.com/maps/documentation/places/web-service/details#session_tokens) for billing purposes.\n\nThe session begins when the user starts typing a query, and concludes when they select a place and a call to Place Details is made. Each session can have multiple queries, followed by one place selection. The API key(s) used for each request within a session must belong to the same Google Cloud Console project. Once a session has concluded, the token is no longer valid; your app must generate a fresh token for each session. If the `sessiontoken` parameter is omitted, or if you reuse a session token, the session is charged as if no session token was provided (each request is billed separately).\n\nWe recommend the following guidelines:\n\n- Use session tokens for all autocomplete sessions.\n- Generate a fresh token for each session. Using a version 4 UUID is recommended.\n- Ensure that the API key(s) used for all Place Autocomplete and Place Details requests within a session belong to the same Cloud Console project.\n- Be sure to pass a unique session token for each new session. Using the same token for more than one session will result in each request being billed individually.\n")),
		mcp.WithString("language", mcp.Description("The language in which to return results.\n\n* See the [list of supported languages](https://developers.google.com/maps/faq#languagesupport). Google often updates the supported languages, so this list may not be exhaustive.\n* If `language` is not supplied, the API attempts to use the preferred language as specified in the `Accept-Language` header.\n* The API does its best to provide a street address that is readable for both the user and locals. To achieve that goal, it returns street addresses in the local language, transliterated to a script readable by the user if necessary, observing the preferred language. All other addresses are returned in the preferred language. Address components are all returned in the same language, which is chosen from the first component.\n* If a name is not available in the preferred language, the API uses the closest match.\n* The preferred language has a small influence on the set of results that the API chooses to return, and the order in which they are returned. The geocoder interprets abbreviations differently depending on language, such as the abbreviations for street types, or synonyms that may be valid in one language but not in another. For example, _utca_ and _tér_ are synonyms for street in Hungarian.")),
		mcp.WithString("region", mcp.Description("The region code, specified as a [ccTLD (\"top-level domain\")](https://en.wikipedia.org/wiki/List_of_Internet_top-level_domains#Country_code_top-level_domains) two-character value. Most ccTLD codes are identical to ISO 3166-1 codes, with some notable exceptions. For example, the United Kingdom's ccTLD is \"uk\" (.co.uk) while its ISO 3166-1 code is \"gb\" (technically for the entity of \"The United Kingdom of Great Britain and Northern Ireland\").")),
		mcp.WithString("reviews_sort", mcp.Description("The sorting method to use when returning reviews. Can be set to `most_relevant` (default) or `newest`.\n\n- For `most_relevant` (default), reviews are sorted by relevance; the service will bias the results to return reviews originally written in the preferred language.\n- For `newest`, reviews are sorted in chronological order; the preferred language does not affect the sort order.\n\nGoogle recommends that you display how the reviews are being sorted to the end user.\n")),
		mcp.WithBoolean("reviews_no_translations", mcp.Description("\nSpecify `reviews_no_translations=true` to disable translation of reviews; specify `reviews_no_translations=false` to enable translation of reviews. Reviews are returned in their original language.\n\nIf omitted, or passed with no value, translation of reviews is enabled. If the `language` parameter was specified in the request, use the specified language as the preferred language for translation. If `language` is omitted, the API attempts to use the `Accept-Language` header as the preferred language.\n")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    PlacedetailsHandler(cfg),
	}
}
