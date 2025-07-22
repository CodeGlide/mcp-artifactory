package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// StreetViewResponse represents the StreetViewResponse schema from the OpenAPI specification
type StreetViewResponse struct {
	Status string `json:"status"` // The `status` field within the Streetview Metadata response object contains the status of the request. The `status` field may contain the following values: - `OK` indicates that no errors occurred; a panorama is found and metadata is returned. - `INVALID_REQUEST` indicates that the request was malformed. - `NOT_FOUND` indicates that the address string provided in the `location` parameter could not be found. This may occur if a non-existent address is given. - `ZERO_RESULTS` indicates that no panorama could be found near the provided location. This may occur if a non-existent or invalid `pano` id is given. - `OVER_QUERY_LIMIT` indicates the requestor has exceeded quota. - `REQUEST_DENIED` indicates that your request was denied. This may occur if you did not [authorize](https://developers.google.com/maps/documentation/streetview/get-api-key) your request, or if the Street View Static API is not activated in the Google Cloud Console project containing your API key. - `UNKNOWN_ERROR` indicates that the request could not be processed due to a server error. This is often a temporary status. The request may succeed if you try again
	Copyright string `json:"copyright,omitempty"` // Contains the copyright notices associated with this panorama.
	Date string `json:"date,omitempty"` // A string indicating year and month that the panorama was captured.
	Location LatLngLiteral `json:"location,omitempty"` // An object describing a specific location with Latitude and Longitude in decimal degrees.
	Pano_id string `json:"pano_id,omitempty"` // A specific panorama ID. These are generally stable, though panoramas may change ID over time as imagery is refreshed.
}

// ErrorResponse represents the ErrorResponse schema from the OpenAPI specification
type ErrorResponse struct {
	ErrorField ErrorObject `json:"error"`
}

// DirectionsTransitDetails represents the DirectionsTransitDetails schema from the OpenAPI specification
type DirectionsTransitDetails struct {
	Trip_short_name string `json:"trip_short_name,omitempty"` // The text that appears in schedules and sign boards to identify a transit trip to passengers. The text should uniquely identify a trip within a service day. For example, "538" is the `trip_short_name` of the Amtrak train that leaves San Jose, CA at 15:10 on weekdays to Sacramento, CA.
	Arrival_stop DirectionsTransitStop `json:"arrival_stop,omitempty"`
	Departure_stop DirectionsTransitStop `json:"departure_stop,omitempty"`
	Headway int `json:"headway,omitempty"` // Specifies the expected number of seconds between departures from the same stop at this time. For example, with a `headway` value of 600, you would expect a ten minute wait if you should miss your bus.
	Headsign string `json:"headsign,omitempty"` // Specifies the direction in which to travel on this line, as it is marked on the vehicle or at the departure stop. This will often be the terminus station.
	Departure_time TimeZoneTextValueObject `json:"departure_time,omitempty"` // An object containing Unix time, a time zone, and its formatted text representation.
	Line DirectionsTransitLine `json:"line,omitempty"`
	Arrival_time TimeZoneTextValueObject `json:"arrival_time,omitempty"` // An object containing Unix time, a time zone, and its formatted text representation.
	Num_stops int `json:"num_stops,omitempty"` // The number of stops from the departure to the arrival stop. This includes the arrival stop, but not the departure stop. For example, if your directions involve leaving from Stop A, passing through stops B and C, and arriving at stop D, `num_stops` will return 3.
}

// PlaceAutocompleteStructuredFormat represents the PlaceAutocompleteStructuredFormat schema from the OpenAPI specification
type PlaceAutocompleteStructuredFormat struct {
	Main_text_matched_substrings []PlaceAutocompleteMatchedSubstring `json:"main_text_matched_substrings"` // Contains an array with `offset` value and `length`. These describe the location of the entered term in the prediction result text, so that the term can be highlighted if desired.
	Secondary_text string `json:"secondary_text,omitempty"` // Contains the secondary text of a prediction, usually the location of the place.
	Secondary_text_matched_substrings []PlaceAutocompleteMatchedSubstring `json:"secondary_text_matched_substrings,omitempty"` // Contains an array with `offset` value and `length`. These describe the location of the entered term in the prediction result text, so that the term can be highlighted if desired.
	Main_text string `json:"main_text"` // Contains the main text of a prediction, usually the name of the place.
}

// DirectionsTrafficSpeedEntry represents the DirectionsTrafficSpeedEntry schema from the OpenAPI specification
type DirectionsTrafficSpeedEntry struct {
	Speed_category string `json:"speed_category"` // The current traffic/speed conditions on this portion of a path.
	Offset_meters float64 `json:"offset_meters"` // The offset along the path (in meters) up to which this speed category is valid.
}

// ElevationResult represents the ElevationResult schema from the OpenAPI specification
type ElevationResult struct {
	Elevation float64 `json:"elevation"` // The elevation of the location in meters.
	Location LatLngLiteral `json:"location"` // An object describing a specific location with Latitude and Longitude in decimal degrees.
	Resolution float64 `json:"resolution,omitempty"` // The value indicating the maximum distance between data points from which the elevation was interpolated, in meters. This property will be missing if the resolution is not known. Note that elevation data becomes more coarse (larger resolution values) when multiple points are passed. To obtain the most accurate elevation value for a point, it should be queried independently.
}

// Geometry represents the Geometry schema from the OpenAPI specification
type Geometry struct {
	Location LatLngLiteral `json:"location"` // An object describing a specific location with Latitude and Longitude in decimal degrees.
	Viewport Bounds `json:"viewport"` // A rectangle in geographical coordinates from points at the southwest and northeast corners.
}

// DirectionsTransitStop represents the DirectionsTransitStop schema from the OpenAPI specification
type DirectionsTransitStop struct {
	Location LatLngLiteral `json:"location"` // An object describing a specific location with Latitude and Longitude in decimal degrees.
	Name string `json:"name"` // The name of the transit stop.
}

// ElevationResponse represents the ElevationResponse schema from the OpenAPI specification
type ElevationResponse struct {
	Status string `json:"status"` // Status codes returned by service. - `OK` indicating the API request was successful. - `DATA_NOT_AVAILABLE` indicating that there's no available data for the input locations. - `INVALID_REQUEST` indicating the API request was malformed. - `OVER_DAILY_LIMIT` indicating any of the following: - The API key is missing or invalid. - Billing has not been enabled on your account. - A self-imposed usage cap has been exceeded. - The provided method of payment is no longer valid (for example, a credit card has expired). - `OVER_QUERY_LIMIT` indicating the requestor has exceeded quota. - `REQUEST_DENIED` indicating the API did not complete the request. - `UNKNOWN_ERROR` indicating an unknown error.
	Error_message string `json:"error_message,omitempty"` // When the service returns a status code other than `OK`, there may be an additional `error_message` field within the response object. This field contains more detailed information about thereasons behind the given status code. This field is not always returned, and its content is subject to change.
	Results []ElevationResult `json:"results"`
}

// DistanceMatrixElement represents the DistanceMatrixElement schema from the OpenAPI specification
type DistanceMatrixElement struct {
	Status string `json:"status"` // - `OK` indicates the response contains a valid result. - `NOT_FOUND` indicates that the origin and/or destination of this pairing could not be geocoded. - `ZERO_RESULTS` indicates no route could be found between the origin and destination. - `MAX_ROUTE_LENGTH_EXCEEDED` indicates the requested route is too long and cannot be processed.
	Distance TextValueObject `json:"distance,omitempty"` // An object containing a numeric value and its formatted text representation.
	Duration TextValueObject `json:"duration,omitempty"` // An object containing a numeric value and its formatted text representation.
	Duration_in_traffic TextValueObject `json:"duration_in_traffic,omitempty"` // An object containing a numeric value and its formatted text representation.
	Fare Fare `json:"fare,omitempty"` // The total fare for the route. ``` { "currency" : "USD", "value" : 6, "text" : "$6.00" } ```
}

// GeocodingGeometry represents the GeocodingGeometry schema from the OpenAPI specification
type GeocodingGeometry struct {
	Location_type string `json:"location_type"` // Stores additional data about the specified location. The following values are currently supported: - "ROOFTOP" indicates that the returned result is a precise geocode for which we have location information accurate down to street address precision. - "RANGE_INTERPOLATED" indicates that the returned result reflects an approximation (usually on a road) interpolated between two precise points (such as intersections). Interpolated results are generally returned when rooftop geocodes are unavailable for a street address. - "GEOMETRIC_CENTER" indicates that the returned result is the geometric center of a result such as a polyline (for example, a street) or polygon (region). - "APPROXIMATE" indicates that the returned result is approximate.
	Viewport Bounds `json:"viewport"` // A rectangle in geographical coordinates from points at the southwest and northeast corners.
	Bounds Bounds `json:"bounds,omitempty"` // A rectangle in geographical coordinates from points at the southwest and northeast corners.
	Location LatLngLiteral `json:"location"` // An object describing a specific location with Latitude and Longitude in decimal degrees.
}

// DirectionsTransitVehicle represents the DirectionsTransitVehicle schema from the OpenAPI specification
type DirectionsTransitVehicle struct {
	TypeField string `json:"type"` // The type of vehicle used. * `BUS` -	Bus. * `CABLE_CAR` -	A vehicle that operates on a cable, usually on the ground. Aerial cable cars may be of the type GONDOLA_LIFT. * `COMMUTER_TRAIN` -	Commuter rail. * `FERRY` -	Ferry. * `FUNICULAR` -	A vehicle that is pulled up a steep incline by a cable. A Funicular typically consists of two cars, with each car acting as a counterweight for the other. * `GONDOLA_LIFT` -	An aerial cable car. * `HEAVY_RAIL` -	Heavy rail. * `HIGH_SPEED_TRAIN` -	High speed train. * `INTERCITY_BUS` -	Intercity bus. * `LONG_DISTANCE_TRAIN` -	Long distance train. * `METRO_RAIL` -	Light rail transit. * `MONORAIL` -	Monorail. * `OTHER` -	All other vehicles will return this type. * `RAIL` -	Rail. * `SHARE_TAXI` -	Share taxi is a kind of bus with the ability to drop off and pick up passengers anywhere on its route. * `SUBWAY` -	Underground light rail. * `TRAM` -	Above ground light rail. * `TROLLEYBUS` -	Trolleybus.
	Icon string `json:"icon,omitempty"` // Contains the URL for an icon associated with this vehicle type.
	Local_icon string `json:"local_icon,omitempty"` // Contains the URL for the icon associated with this vehicle type, based on the local transport signage.
	Name string `json:"name"` // The name of this vehicle, capitalized.
}

// PlaceSpecialDay represents the PlaceSpecialDay schema from the OpenAPI specification
type PlaceSpecialDay struct {
	Date string `json:"date,omitempty"` // A date expressed in RFC3339 format in the local timezone for the place, for example 2010-12-31.
	Exceptional_hours bool `json:"exceptional_hours,omitempty"` // True if there are exceptional hours for this day. If `true`, this means that there is at least one exception for this day. Exceptions cause different values to occur in the subfields of `current_opening_hours` and `secondary_opening_hours` such as `periods`, `weekday_text`, `open_now`. The exceptions apply to the hours, and the hours are used to generate the other fields.
}

// Bounds represents the Bounds schema from the OpenAPI specification
type Bounds struct {
	Northeast LatLngLiteral `json:"northeast"` // An object describing a specific location with Latitude and Longitude in decimal degrees.
	Southwest LatLngLiteral `json:"southwest"` // An object describing a specific location with Latitude and Longitude in decimal degrees.
}

// TimeZoneResponse represents the TimeZoneResponse schema from the OpenAPI specification
type TimeZoneResponse struct {
	Timezonename string `json:"timeZoneName,omitempty"` // The long form name of the time zone. This field will be localized if the language parameter is set. eg. `Pacific Daylight Time` or `Australian Eastern Daylight Time`.
	Dstoffset float64 `json:"dstOffset,omitempty"` // The offset for daylight-savings time in seconds. This will be zero if the time zone is not in Daylight Savings Time during the specified `timestamp`.
	Errormessage string `json:"errorMessage,omitempty"` // Detailed information about the reasons behind the given status code. Included if status other than `Ok`.
	Rawoffset float64 `json:"rawOffset,omitempty"` // The offset from UTC (in seconds) for the given location. This does not take into effect daylight savings.
	Status string `json:"status"` // The `status` field within the Time Zone response object contains the status of the request. The `status` field may contain the following values: - `OK` indicates that the request was successful. - `INVALID_REQUEST` indicates that the request was malformed. - `OVER_DAILY_LIMIT` indicates any of the following: - The API key is missing or invalid. - Billing has not been enabled on your account. - A self-imposed usage cap has been exceeded. - The provided method of payment is no longer valid (for example, a credit card has expired). - `OVER_QUERY_LIMIT` indicates the requestor has exceeded quota. - `REQUEST_DENIED` indicates that the API did not complete the request. Confirm that the request was sent over HTTPS instead of HTTP. - `UNKNOWN_ERROR` indicates an unknown error. - `ZERO_RESULTS` indicates that no time zone data could be found for the specified position or time. Confirm that the request is for a location on land, and not over water.
	Timezoneid string `json:"timeZoneId,omitempty"` // a string containing the ID of the time zone, such as "America/Los_Angeles" or "Australia/Sydney". These IDs are defined by [Unicode Common Locale Data Repository (CLDR) project](http://cldr.unicode.org/), and currently available in file timezone.xml. When a timezone has several IDs, the canonical one is returned. In xml responses, this is the first alias of each timezone. For example, "Asia/Calcutta" is returned, not "Asia/Kolkata".
}

// ErrorObject represents the ErrorObject schema from the OpenAPI specification
type ErrorObject struct {
	Code float64 `json:"code"` // This is the same as the HTTP status of the response.
	Details []ErrorDetail `json:"details,omitempty"` // Additional details about the error.
	Errors []ErrorDetail `json:"errors"` // A list of errors which occurred. Each error contains an identifier for the type of error and a short description.
	Message string `json:"message"` // A short description of the error.
	Status string `json:"status,omitempty"` // A status code that indicates the error type.
}

// DirectionsStep represents the DirectionsStep schema from the OpenAPI specification
type DirectionsStep struct {
	Travel_mode string `json:"travel_mode"` // - `DRIVING` (default) indicates calculation using the road network. - `BICYCLING` requests calculation for bicycling via bicycle paths & preferred streets (where available). - `TRANSIT` requests calculation via public transit routes (where available). - `WALKING` requests calculation for walking via pedestrian paths & sidewalks (where available).
	Distance TextValueObject `json:"distance,omitempty"` // An object containing a numeric value and its formatted text representation.
	End_location LatLngLiteral `json:"end_location"` // An object describing a specific location with Latitude and Longitude in decimal degrees.
	Html_instructions string `json:"html_instructions"` // Contains formatted instructions for this step, presented as an HTML text string. This content is meant to be read as-is. Do not programmatically parse this display-only content.
	Start_location LatLngLiteral `json:"start_location"` // An object describing a specific location with Latitude and Longitude in decimal degrees.
	Duration TextValueObject `json:"duration"` // An object containing a numeric value and its formatted text representation.
	Maneuver string `json:"maneuver,omitempty"` // Contains the action to take for the current step (turn left, merge, straight, etc.). Values are subject to change, and new values may be introduced without prior notice.
	Polyline DirectionsPolyline `json:"polyline"` // [Polyline encoding](https://developers.google.com/maps/documentation/utilities/polylinealgorithm) is a lossy compression algorithm that allows you to store a series of coordinates as a single string. Point coordinates are encoded using signed values. If you only have a few static points, you may also wish to use the interactive polyline encoding utility. The encoding process converts a binary value into a series of character codes for ASCII characters using the familiar base64 encoding scheme: to ensure proper display of these characters, encoded values are summed with 63 (the ASCII character '?') before converting them into ASCII. The algorithm also checks for additional character codes for a given point by checking the least significant bit of each byte group; if this bit is set to 1, the point is not yet fully formed and additional data must follow. Additionally, to conserve space, points only include the offset from the previous point (except of course for the first point). All points are encoded in Base64 as signed integers, as latitudes and longitudes are signed values. The encoding format within a polyline needs to represent two coordinates representing latitude and longitude to a reasonable precision. Given a maximum longitude of +/- 180 degrees to a precision of 5 decimal places (180.00000 to -180.00000), this results in the need for a 32 bit signed binary integer value.
	Steps interface{} `json:"steps,omitempty"` // Contains detailed directions for walking or driving steps in transit directions. Substeps are only available when travel_mode is set to "transit". The inner steps array is of the same type as steps.
	Transit_details DirectionsTransitDetails `json:"transit_details,omitempty"` // Additional information that is not relevant for other modes of transportation.
}

// DirectionsViaWaypoint represents the DirectionsViaWaypoint schema from the OpenAPI specification
type DirectionsViaWaypoint struct {
	Step_interpolation float64 `json:"step_interpolation,omitempty"` // The position of the waypoint along the step's polyline, expressed as a ratio from 0 to 1.
	Location LatLngLiteral `json:"location,omitempty"` // An object describing a specific location with Latitude and Longitude in decimal degrees.
	Step_index int `json:"step_index,omitempty"` // The index of the step containing the waypoint.
}

// PlaceAutocompleteTerm represents the PlaceAutocompleteTerm schema from the OpenAPI specification
type PlaceAutocompleteTerm struct {
	Offset float64 `json:"offset"` // Defines the start position of this term in the description, measured in Unicode characters
	Value string `json:"value"` // The text of the term.
}

// PlaceAutocompleteMatchedSubstring represents the PlaceAutocompleteMatchedSubstring schema from the OpenAPI specification
type PlaceAutocompleteMatchedSubstring struct {
	Length float64 `json:"length"` // Length of the matched substring in the prediction result text.
	Offset float64 `json:"offset"` // Start location of the matched substring in the prediction result text.
}

// PlaceAutocompletePrediction represents the PlaceAutocompletePrediction schema from the OpenAPI specification
type PlaceAutocompletePrediction struct {
	Matched_substrings []PlaceAutocompleteMatchedSubstring `json:"matched_substrings"` // A list of substrings that describe the location of the entered term in the prediction result text, so that the term can be highlighted if desired.
	Place_id string `json:"place_id,omitempty"` // A textual identifier that uniquely identifies a place. To retrieve information about the place, pass this identifier in the placeId field of a Places API request. For more information about place IDs, see the [Place IDs](https://developers.google.com/maps/documentation/places/web-service/place-id) overview.
	Reference string `json:"reference,omitempty"` // See place_id.
	Structured_formatting PlaceAutocompleteStructuredFormat `json:"structured_formatting"`
	Terms []PlaceAutocompleteTerm `json:"terms"` // Contains an array of terms identifying each section of the returned description (a section of the description is generally terminated with a comma). Each entry in the array has a `value` field, containing the text of the term, and an `offset` field, defining the start position of this term in the description, measured in Unicode characters.
	Types []string `json:"types,omitempty"` // Contains an array of types that apply to this place. For example: `[ "political", "locality" ]` or `[ "establishment", "geocode", "beauty_salon" ]`. The array can contain multiple values. Learn more about [Place types](https://developers.google.com/maps/documentation/places/web-service/supported_types).
	Description string `json:"description"` // Contains the human-readable name for the returned result. For `establishment` results, this is usually the business name. This content is meant to be read as-is. Do not programmatically parse the formatted address.
	Distance_meters int `json:"distance_meters,omitempty"` // The straight-line distance in meters from the origin. This field is only returned for requests made with an `origin`.
}

// GeocodingResult represents the GeocodingResult schema from the OpenAPI specification
type GeocodingResult struct {
	Formatted_address string `json:"formatted_address"` // The human-readable address of this location.
	Geometry GeocodingGeometry `json:"geometry"` // An object describing the location.
	Partial_match bool `json:"partial_match,omitempty"` // Indicates that the geocoder did not return an exact match for the original request, though it was able to match part of the requested address. You may wish to examine the original request for misspellings and/or an incomplete address. Partial matches most often occur for street addresses that do not exist within the locality you pass in the request. Partial matches may also be returned when a request matches two or more locations in the same locality.
	Place_id string `json:"place_id"` // A unique identifier that can be used with other Google APIs. For example, you can use the `place_id` in a Places API request to get details of a local business, such as phone number, opening hours, user reviews, and more. See the [place ID overview](https://developers.google.com/places/place-id).
	Plus_code PlusCode `json:"plus_code,omitempty"` // An encoded location reference, derived from latitude and longitude coordinates, that represents an area, 1/8000th of a degree by 1/8000th of a degree (about 14m x 14m at the equator) or smaller. Plus codes can be used as a replacement for street addresses in places where they do not exist (where buildings are not numbered or streets are not named).
	Postcode_localities []string `json:"postcode_localities,omitempty"` // An array denoting all the localities contained in a postal code. This is only present when the result is a postal code that contains multiple localities.
	Types []string `json:"types"` // The `types[]` array indicates the type of the returned result. This array contains a set of zero or more tags identifying the type of feature returned in the result. For example, a geocode of "Chicago" returns "locality" which indicates that "Chicago" is a city, and also returns "political" which indicates it is a political entity.
	Address_components []AddressComponent `json:"address_components"` // An array containing the separate components applicable to this address.
}

// DirectionsTransitLine represents the DirectionsTransitLine schema from the OpenAPI specification
type DirectionsTransitLine struct {
	Agencies []DirectionsTransitAgency `json:"agencies"` // The transit agency (or agencies) that operates this transit line.
	Color string `json:"color,omitempty"` // The color commonly used in signage for this line.
	Icon string `json:"icon,omitempty"` // Contains the URL for the icon associated with this line.
	Name string `json:"name"` // The full name of this transit line, e.g. "8 Avenue Local".
	Short_name string `json:"short_name,omitempty"` // The short name of this transit line. This will normally be a line number, such as "M7" or "355".
	Text_color string `json:"text_color,omitempty"` // The color commonly used in signage for this line.
	Url string `json:"url,omitempty"` // Contains the URL for this transit line as provided by the transit agency.
	Vehicle DirectionsTransitVehicle `json:"vehicle,omitempty"`
}

// GeolocationRequest represents the GeolocationRequest schema from the OpenAPI specification
type GeolocationRequest struct {
	Homemobilecountrycode int `json:"homeMobileCountryCode,omitempty"` // The cell tower's Mobile Country Code (MCC).
	Homemobilenetworkcode int `json:"homeMobileNetworkCode,omitempty"` // The cell tower's Mobile Network Code. This is the MNC for GSM and WCDMA; CDMA uses the System ID (SID).
	Radiotype string `json:"radioType,omitempty"` // The mobile radio type. Supported values are lte, gsm, cdma, and wcdma. While this field is optional, it should be included if a value is available, for more accurate results.
	Wifiaccesspoints []WiFiAccessPoint `json:"wifiAccessPoints,omitempty"` // An array of two or more WiFi access point objects.
	Carrier string `json:"carrier,omitempty"` // The carrier name.
	Celltowers []CellTower `json:"cellTowers,omitempty"` // The request body's cellTowers array contains zero or more cell tower objects.
	Considerip string `json:"considerIp,omitempty"` // Specifies whether to fall back to IP geolocation if wifi and cell tower signals are not available. Defaults to true. Set considerIp to false to disable fall back.
}

// SnapToRoadsResponse represents the SnapToRoadsResponse schema from the OpenAPI specification
type SnapToRoadsResponse struct {
	Warningmessage string `json:"warningMessage,omitempty"` // A string containing a user-visible warning.
	Snappedpoints []SnappedPoint `json:"snappedPoints,omitempty"` // An array of snapped points.
}

// NearestRoadsError represents the NearestRoadsError schema from the OpenAPI specification
type NearestRoadsError struct {
	Status string `json:"status"` // An error such as `INVALID_ARGUMENT`.
	Code float64 `json:"code"` // This is the same as the HTTP status of the response.
	Message string `json:"message"` // A short description of the error.
}

// Place represents the Place schema from the OpenAPI specification
type Place struct {
	Dine_in bool `json:"dine_in,omitempty"` // Specifies if the business supports indoor or outdoor seating options.
	Wheelchair_accessible_entrance bool `json:"wheelchair_accessible_entrance,omitempty"` // Specifies if the place has an entrance that is wheelchair-accessible.
	Serves_beer bool `json:"serves_beer,omitempty"` // Specifies if the place serves beer.
	User_ratings_total float64 `json:"user_ratings_total,omitempty"` // The total number of reviews, with or without text, for this place.
	Icon string `json:"icon,omitempty"` // Contains the URL of a suggested icon which may be displayed to the user when indicating this result on a map.
	Serves_breakfast bool `json:"serves_breakfast,omitempty"` // Specifies if the place serves breakfast.
	Utc_offset float64 `json:"utc_offset,omitempty"` // Contains the number of minutes this place’s current timezone is offset from UTC. For example, for places in Sydney, Australia during daylight saving time this would be 660 (+11 hours from UTC), and for places in California outside of daylight saving time this would be -480 (-8 hours from UTC).
	Delivery bool `json:"delivery,omitempty"` // Specifies if the business supports delivery.
	Reference string `json:"reference,omitempty"`
	Geometry Geometry `json:"geometry,omitempty"` // An object describing the location.
	Secondary_opening_hours []PlaceOpeningHours `json:"secondary_opening_hours,omitempty"` // Contains an array of entries for the next seven days including information about secondary hours of a business. Secondary hours are different from a business's main hours. For example, a restaurant can specify drive through hours or delivery hours as its secondary hours. This field populates the `type` subfield, which draws from a predefined list of opening hours types (such as `DRIVE_THROUGH`, `PICKUP`, or `TAKEOUT`) based on the types of the place. This field includes the `special_days` subfield of all hours, set for dates that have exceptional hours.
	Icon_mask_base_uri string `json:"icon_mask_base_uri,omitempty"` // Contains the URL of a recommended icon, minus the `.svg` or `.png` file type extension.
	Plus_code PlusCode `json:"plus_code,omitempty"` // An encoded location reference, derived from latitude and longitude coordinates, that represents an area, 1/8000th of a degree by 1/8000th of a degree (about 14m x 14m at the equator) or smaller. Plus codes can be used as a replacement for street addresses in places where they do not exist (where buildings are not numbered or streets are not named).
	Serves_brunch bool `json:"serves_brunch,omitempty"` // Specifies if the place serves brunch.
	Adr_address string `json:"adr_address,omitempty"` // A representation of the place's address in the [adr microformat](http://microformats.org/wiki/adr).
	Place_id string `json:"place_id,omitempty"` // A textual identifier that uniquely identifies a place. To retrieve information about the place, pass this identifier in the `place_id` field of a Places API request. For more information about place IDs, see the [place ID overview](https://developers.google.com/maps/documentation/places/web-service/place-id).
	Serves_dinner bool `json:"serves_dinner,omitempty"` // Specifies if the place serves dinner.
	Photos []PlacePhoto `json:"photos,omitempty"` // An array of photo objects, each containing a reference to an image. A request may return up to ten photos. More information about place photos and how you can use the images in your application can be found in the [Place Photos](https://developers.google.com/maps/documentation/places/web-service/photos) documentation.
	Serves_vegetarian_food bool `json:"serves_vegetarian_food,omitempty"` // Specifies if the place serves vegetarian food.
	Vicinity string `json:"vicinity,omitempty"` // For establishment (`types:["establishment", ...])` results only, the `vicinity` field contains a simplified address for the place, including the street name, street number, and locality, but not the province/state, postal code, or country. For all other results, the `vicinity` field contains the name of the narrowest political (`types:["political", ...]`) feature that is present in the address of the result. This content is meant to be read as-is. Do not programmatically parse the formatted address.
	Business_status string `json:"business_status,omitempty"` // Indicates the operational status of the place, if it is a business. If no data exists, `business_status` is not returned.
	Current_opening_hours PlaceOpeningHours `json:"current_opening_hours,omitempty"` // An object describing the opening hours of a place.
	Opening_hours PlaceOpeningHours `json:"opening_hours,omitempty"` // An object describing the opening hours of a place.
	Serves_lunch bool `json:"serves_lunch,omitempty"` // Specifies if the place serves lunch.
	Takeout bool `json:"takeout,omitempty"` // Specifies if the business supports takeout.
	Rating float64 `json:"rating,omitempty"` // Contains the place's rating, from 1.0 to 5.0, based on aggregated user reviews.
	Price_level float64 `json:"price_level,omitempty"` // The price level of the place, on a scale of 0 to 4. The exact amount indicated by a specific value will vary from region to region. Price levels are interpreted as follows: - 0 Free - 1 Inexpensive - 2 Moderate - 3 Expensive - 4 Very Expensive
	Editorial_summary PlaceEditorialSummary `json:"editorial_summary,omitempty"` // Contains a summary of the place. A summary is comprised of a textual overview, and also includes the language code for these if applicable. Summary text must be presented as-is and can not be modified or altered.
	Address_components []AddressComponent `json:"address_components,omitempty"` // An array containing the separate components applicable to this address.
	Name string `json:"name,omitempty"` // Contains the human-readable name for the returned result. For `establishment` results, this is usually the canonicalized business name.
	Permanently_closed bool `json:"permanently_closed,omitempty"` // Use `business_status` to get the operational status of businesses.
	Reviews []PlaceReview `json:"reviews,omitempty"` // A JSON array of up to five reviews. By default, the reviews are sorted in order of relevance. Use the `reviews_sort` request parameter to control sorting. - For `most_relevant` (default), reviews are sorted by relevance; the service will bias the results to return reviews originally written in the preferred language. - For `newest`, reviews are sorted in chronological order; the preferred language does not affect the sort order. Google recommends indicating to users whether results are ordered by `most_relevant` or `newest`.
	Icon_background_color string `json:"icon_background_color,omitempty"` // Contains the default HEX color code for the place's category.
	Website string `json:"website,omitempty"` // The authoritative website for this place, such as a business' homepage.
	Formatted_phone_number string `json:"formatted_phone_number,omitempty"` // Contains the place's phone number in its [local format](http://en.wikipedia.org/wiki/Local_conventions_for_writing_telephone_numbers).
	Reservable bool `json:"reservable,omitempty"` // Specifies if the place supports reservations.
	International_phone_number string `json:"international_phone_number,omitempty"` // Contains the place's phone number in international format. International format includes the country code, and is prefixed with the plus, +, sign. For example, the international_phone_number for Google's Sydney, Australia office is `+61 2 9374 4000`.
	Types []string `json:"types,omitempty"` // Contains an array of feature types describing the given result. See the list of [supported types](https://developers.google.com/maps/documentation/places/web-service/supported_types#table2).
	Formatted_address string `json:"formatted_address,omitempty"` // A string containing the human-readable address of this place. Often this address is equivalent to the postal address. Note that some countries, such as the United Kingdom, do not allow distribution of true postal addresses due to licensing restrictions. The formatted address is logically composed of one or more address components. For example, the address "111 8th Avenue, New York, NY" consists of the following components: "111" (the street number), "8th Avenue" (the route), "New York" (the city) and "NY" (the US state). Do not parse the formatted address programmatically. Instead you should use the individual address components, which the API response includes in addition to the formatted address field.
	Serves_wine bool `json:"serves_wine,omitempty"` // Specifies if the place serves wine.
	Curbside_pickup bool `json:"curbside_pickup,omitempty"` // Specifies if the business supports curbside pickup.
	Scope string `json:"scope,omitempty"`
	Url string `json:"url,omitempty"` // Contains the URL of the official Google page for this place. This will be the Google-owned page that contains the best available information about the place. Applications must link to or embed this page on any screen that shows detailed results about the place to the user.
}

// LatLngLiteral represents the LatLngLiteral schema from the OpenAPI specification
type LatLngLiteral struct {
	Lat float64 `json:"lat"` // Latitude in decimal degrees
	Lng float64 `json:"lng"` // Longitude in decimal degrees
}

// TextValueObject represents the TextValueObject schema from the OpenAPI specification
type TextValueObject struct {
	Text string `json:"text"` // String value.
	Value float64 `json:"value"` // Numeric value.
}

// NearestRoadsErrorResponse represents the NearestRoadsErrorResponse schema from the OpenAPI specification
type NearestRoadsErrorResponse struct {
	ErrorField NearestRoadsError `json:"error,omitempty"`
}

// LatitudeLongitudeLiteral represents the LatitudeLongitudeLiteral schema from the OpenAPI specification
type LatitudeLongitudeLiteral struct {
	Latitude float64 `json:"latitude"` // Latitude in decimal degrees
	Longitude float64 `json:"longitude"` // Longitude in decimal degrees
}

// PlaceOpeningHours represents the PlaceOpeningHours schema from the OpenAPI specification
type PlaceOpeningHours struct {
	TypeField string `json:"type,omitempty"` // A type string used to identify the type of secondary hours (for example, `DRIVE_THROUGH`, `HAPPY_HOUR`, `DELIVERY`, `TAKEOUT`, `KITCHEN`, `BREAKFAST`, `LUNCH`, `DINNER`, `BRUNCH`, `PICKUP`, `SENIOR_HOURS`). Set for `secondary_opening_hours` only.
	Weekday_text []string `json:"weekday_text,omitempty"` // An array of strings describing in human-readable text the hours of the place.
	Open_now bool `json:"open_now,omitempty"` // A boolean value indicating if the place is open at the current time.
	Periods []PlaceOpeningHoursPeriod `json:"periods,omitempty"` // An array of opening periods covering seven days, starting from Sunday, in chronological order.
	Special_days []PlaceSpecialDay `json:"special_days,omitempty"` // An array of up to seven entries corresponding to the next seven days.
}

// AddressComponent represents the AddressComponent schema from the OpenAPI specification
type AddressComponent struct {
	Long_name string `json:"long_name"` // The full text description or name of the address component as returned by the Geocoder.
	Short_name string `json:"short_name"` // An abbreviated textual name for the address component, if available. For example, an address component for the state of Alaska may have a long_name of "Alaska" and a short_name of "AK" using the 2-letter postal abbreviation.
	Types []string `json:"types"` // An array indicating the type of the address component. See the list of [supported types](https://developers.google.com/maps/documentation/places/web-service/supported_types).
}

// DirectionsResponse represents the DirectionsResponse schema from the OpenAPI specification
type DirectionsResponse struct {
	Status string `json:"status"` // The status field within the Directions response object contains the status of the request, and may contain debugging information to help you track down why the Directions service failed. The status field may contain the following values: - `OK` indicates the response contains a valid result. - `NOT_FOUND` indicates at least one of the locations specified in the request's origin, destination, or waypoints could not be geocoded. - `ZERO_RESULTS` indicates no route could be found between the origin and destination. - `MAX_WAYPOINTS_EXCEEDED` indicates that too many waypoints were provided in the request. For applications using the Directions API as a web service, or the directions service in the Maps JavaScript API, the maximum allowed number of waypoints is 25, plus the origin and destination. - `MAX_ROUTE_LENGTH_EXCEEDED` indicates the requested route is too long and cannot be processed. This error occurs when more complex directions are returned. Try reducing the number of waypoints, turns, or instructions. - `INVALID_REQUEST` indicates that the provided request was invalid. Common causes of this status include an invalid parameter or parameter value. - `OVER_DAILY_LIMIT` indicates any of the following: - The API key is missing or invalid. - Billing has not been enabled on your account. - A self-imposed usage cap has been exceeded. - The provided method of payment is no longer valid (for example, a credit card has expired). See the [Maps FAQ](https://developers.google.com/maps/faq#over-limit-key-error) to learn how to fix this. - `OVER_QUERY_LIMIT` indicates the service has received too many requests from your application within the allowed time period. - `REQUEST_DENIED` indicates that the service denied use of the directions service by your application. - `UNKNOWN_ERROR` indicates a directions request could not be processed due to a server error. The request may succeed if you try again.
	Available_travel_modes []string `json:"available_travel_modes,omitempty"` // Contains an array of available travel modes. This field is returned when a request specifies a travel mode and gets no results. The array contains the available travel modes in the countries of the given set of waypoints. This field is not returned if one or more of the waypoints are 'via waypoints'.
	Error_message string `json:"error_message,omitempty"` // When the service returns a status code other than `OK`, there may be an additional `error_message` field within the response object. This field contains more detailed information about the reasons behind the given status code. This field is not always returned, and its content is subject to change.
	Geocoded_waypoints []DirectionsGeocodedWaypoint `json:"geocoded_waypoints,omitempty"` // Contains an array with details about the geocoding of origin, destination and waypoints. Elements in the geocoded_waypoints array correspond, by their zero-based position, to the origin, the waypoints in the order they are specified, and the destination. These details will not be present for waypoints specified as textual latitude/longitude values if the service returns no results. This is because such waypoints are only reverse geocoded to obtain their representative address after a route has been found. An empty JSON object will occupy the corresponding places in the geocoded_waypoints array.
	Routes []DirectionsRoute `json:"routes"` // Contains an array of routes from the origin to the destination. Routes consist of nested Legs and Steps.
}

// PlaceEditorialSummary represents the PlaceEditorialSummary schema from the OpenAPI specification
type PlaceEditorialSummary struct {
	Language string `json:"language,omitempty"` // The language of the previous fields. May not always be present.
	Overview string `json:"overview,omitempty"` // A medium-length textual summary of the place.
}

// PlaceOpeningHoursPeriod represents the PlaceOpeningHoursPeriod schema from the OpenAPI specification
type PlaceOpeningHoursPeriod struct {
	CloseField PlaceOpeningHoursPeriodDetail `json:"close,omitempty"`
	Open PlaceOpeningHoursPeriodDetail `json:"open"`
}

// SnappedPoint represents the SnappedPoint schema from the OpenAPI specification
type SnappedPoint struct {
	Location LatitudeLongitudeLiteral `json:"location"` // An object describing a specific location with Latitude and Longitude in decimal degrees.
	Originalindex float64 `json:"originalIndex,omitempty"` // An integer that indicates the corresponding value in the original request. Each value in the request should map to a snapped value in the response. However, if you've set interpolate=true or if you're using nearest roads, then it's possible that the response will contain more coordinates than the request. Interpolated values will not have an `originalIndex`. These values are indexed from `0`, so a point with an originalIndex of `4` will be the snapped value of the 5th latitude/longitude passed to the path parameter. Nearest Roads points may contain several points for single coordinates with differing location or placeId.
	Placeid string `json:"placeId"` // A unique identifier for a place. All place IDs returned by the Roads API correspond to road segments.
}

// GeolocationResponse represents the GeolocationResponse schema from the OpenAPI specification
type GeolocationResponse struct {
	Accuracy float64 `json:"accuracy"` // The accuracy of the estimated location, in meters. This represents the radius of a circle around the given `location`. If your Geolocation response shows a very high value in the `accuracy` field, the service may be geolocating based on the request IP, instead of WiFi points or cell towers. This can happen if no cell towers or access points are valid or recognized. To confirm that this is the issue, set `considerIp` to `false` in your request. If the response is a `404`, you've confirmed that your `wifiAccessPoints` and `cellTowers` objects could not be geolocated.
	Location LatLngLiteral `json:"location"` // An object describing a specific location with Latitude and Longitude in decimal degrees.
}

// PlacePhoto represents the PlacePhoto schema from the OpenAPI specification
type PlacePhoto struct {
	Html_attributions []string `json:"html_attributions"` // The HTML attributions for the photo.
	Photo_reference string `json:"photo_reference"` // A string used to identify the photo when you perform a Photo request.
	Width float64 `json:"width"` // The width of the photo.
	Height float64 `json:"height"` // The height of the photo.
}

// FieldViolation represents the FieldViolation schema from the OpenAPI specification
type FieldViolation struct {
	Description string `json:"description"` // A short description of the error.
	Field string `json:"field"` // The name of the invalid field.
}

// PlacesDetailsResponse represents the PlacesDetailsResponse schema from the OpenAPI specification
type PlacesDetailsResponse struct {
	Html_attributions []string `json:"html_attributions"` // May contain a set of attributions about this listing which must be displayed to the user (some listings may not have attribution).
	Info_messages []string `json:"info_messages,omitempty"` // When the service returns additional information about the request specification, there may be an additional `info_messages` field within the response object. This field is only returned for successful requests. It may not always be returned, and its content is subject to change.
	Result Place `json:"result"` // Attributes describing a place. Not all attributes will be available for all place types.
	Status string `json:"status"` // Status codes returned by service. - `OK` indicating the API request was successful. - `ZERO_RESULTS` indicating that the referenced location, `place_id`, was valid but no longer refers to a valid result. This may occur if the establishment is no longer in business. - `NOT_FOUND` indicating that that the referenced location, `place_id`, was not found in the Places database. - `INVALID_REQUEST` indicating the API request was malformed. - `OVER_QUERY_LIMIT` indicating any of the following: - You have exceeded the QPS limits. - Billing has not been enabled on your account. - The monthly $200 credit, or a self-imposed usage cap, has been exceeded. - The provided method of payment is no longer valid (for example, a credit card has expired). See the [Maps FAQ](https://developers.google.com/maps/faq#over-limit-key-error) for more information about how to resolve this error. - `REQUEST_DENIED` indicating that your request was denied, generally because: - The request is missing an API key. - The `key` parameter is invalid. - `UNKNOWN_ERROR` indicating an unknown error.
}

// DirectionsPolyline represents the DirectionsPolyline schema from the OpenAPI specification
type DirectionsPolyline struct {
	Points string `json:"points"` // A single string representation of the polyline.
}

// PlusCode represents the PlusCode schema from the OpenAPI specification
type PlusCode struct {
	Compound_code string `json:"compound_code,omitempty"` // The `compound_code` is a 6 character or longer local code with an explicit location (`CWC8+R9, Mountain View, CA, USA`). Some APIs may return an empty string if the `compound_code` is not available.
	Global_code string `json:"global_code"` // The `global_code` is a 4 character area code and 6 character or longer local code (`849VCWC8+R9`).
}

// PlacesNearbySearchResponse represents the PlacesNearbySearchResponse schema from the OpenAPI specification
type PlacesNearbySearchResponse struct {
	Html_attributions []string `json:"html_attributions"` // May contain a set of attributions about this listing which must be displayed to the user (some listings may not have attribution).
	Info_messages []string `json:"info_messages,omitempty"` // When the service returns additional information about the request specification, there may be an additional `info_messages` field within the response object. This field is only returned for successful requests. It may not always be returned, and its content is subject to change.
	Next_page_token string `json:"next_page_token,omitempty"` // Contains a token that can be used to return up to 20 additional results. A next_page_token will not be returned if there are no additional results to display. The maximum number of results that can be returned is 60. There is a short delay between when a next_page_token is issued, and when it will become valid.
	Results []Place `json:"results"` // Contains an array of places. <div class="caution">Place Search requests return a subset of the fields that are returned by Place Details requests. If the field you want is not returned by Place Search, you can use Place Search to get a `place_id`, then use that Place ID to make a Place Details request.</div>
	Status string `json:"status"` // Status codes returned by service. - `OK` indicating the API request was successful. - `ZERO_RESULTS` indicating that the search was successful but returned no results. This may occur if the search was passed a `latlng` in a remote location. - `INVALID_REQUEST` indicating the API request was malformed, generally due to missing required query parameter (`location` or `radius`). - `OVER_QUERY_LIMIT` indicating any of the following: - You have exceeded the QPS limits. - Billing has not been enabled on your account. - The monthly $200 credit, or a self-imposed usage cap, has been exceeded. - The provided method of payment is no longer valid (for example, a credit card has expired). See the [Maps FAQ](https://developers.google.com/maps/faq#over-limit-key-error) for more information about how to resolve this error. - `REQUEST_DENIED` indicating that your request was denied, generally because: - The request is missing an API key. - The `key` parameter is invalid. - `UNKNOWN_ERROR` indicating an unknown error.
	Error_message string `json:"error_message,omitempty"` // When the service returns a status code other than `OK<`, there may be an additional `error_message` field within the response object. This field contains more detailed information about thereasons behind the given status code. This field is not always returned, and its content is subject to change.
}

// CellTower represents the CellTower schema from the OpenAPI specification
type CellTower struct {
	Mobilecountrycode int `json:"mobileCountryCode"` // The cell tower's Mobile Country Code (MCC).
	Mobilenetworkcode int `json:"mobileNetworkCode"` // The cell tower's Mobile Network Code. This is the MNC for GSM and WCDMA; CDMA uses the System ID (SID).
	Signalstrength float64 `json:"signalStrength,omitempty"` // Radio signal strength measured in dBm.
	Timingadvance float64 `json:"timingAdvance,omitempty"` // The timing advance value.
	Age int `json:"age,omitempty"` // The number of milliseconds since this cell was primary. If age is 0, the cellId represents a current measurement.
	Cellid int `json:"cellId"` // Unique identifier of the cell. On GSM, this is the Cell ID (CID); CDMA networks use the Base Station ID (BID). WCDMA networks use the UTRAN/GERAN Cell Identity (UC-Id), which is a 32-bit value concatenating the Radio Network Controller (RNC) and Cell ID. Specifying only the 16-bit Cell ID value in WCDMA networks may return inaccurate results.
	Locationareacode int `json:"locationAreaCode"` // The Location Area Code (LAC) for GSM and WCDMA networks. The Network ID (NID) for CDMA networks.
}

// Fare represents the Fare schema from the OpenAPI specification
type Fare struct {
	Currency string `json:"currency"` // An [ISO 4217 currency code](https://en.wikipedia.org/wiki/ISO_4217) indicating the currency that the amount is expressed in.
	Text string `json:"text"` // The total fare amount, formatted in the requested language.
	Value float64 `json:"value"` // The total fare amount, in the currency specified.
}

// DirectionsLeg represents the DirectionsLeg schema from the OpenAPI specification
type DirectionsLeg struct {
	Distance TextValueObject `json:"distance,omitempty"` // An object containing a numeric value and its formatted text representation.
	Duration_in_traffic TextValueObject `json:"duration_in_traffic,omitempty"` // An object containing a numeric value and its formatted text representation.
	Start_location LatLngLiteral `json:"start_location"` // An object describing a specific location with Latitude and Longitude in decimal degrees.
	Steps []DirectionsStep `json:"steps"` // An array of steps denoting information about each separate step of the leg of the journey.
	Departure_time TimeZoneTextValueObject `json:"departure_time,omitempty"` // An object containing Unix time, a time zone, and its formatted text representation.
	Duration TextValueObject `json:"duration,omitempty"` // An object containing a numeric value and its formatted text representation.
	End_location LatLngLiteral `json:"end_location"` // An object describing a specific location with Latitude and Longitude in decimal degrees.
	Start_address string `json:"start_address"` // Contains the human-readable address (typically a street address) resulting from reverse geocoding the `start_location` of this leg. This content is meant to be read as-is. Do not programmatically parse the formatted address.
	End_address string `json:"end_address"` // Contains the human-readable address (typically a street address) from reverse geocoding the `end_location` of this leg. This content is meant to be read as-is. Do not programmatically parse the formatted address.
	Arrival_time TimeZoneTextValueObject `json:"arrival_time,omitempty"` // An object containing Unix time, a time zone, and its formatted text representation.
	Traffic_speed_entry []DirectionsTrafficSpeedEntry `json:"traffic_speed_entry"` // Information about traffic speed along the leg.
	Via_waypoint []DirectionsViaWaypoint `json:"via_waypoint"` // The locations of via waypoints along this leg.
}

// DistanceMatrixRow represents the DistanceMatrixRow schema from the OpenAPI specification
type DistanceMatrixRow struct {
	Elements []DistanceMatrixElement `json:"elements"` // When the Distance Matrix API returns results, it places them within a JSON rows array. Even if no results are returned (such as when the origins and/or destinations don't exist), it still returns an empty array. Rows are ordered according to the values in the origin parameter of the request. Each row corresponds to an origin, and each element within that row corresponds to a pairing of the origin with a destination value. Each row array contains one or more element entries, which in turn contain the information about a single origin-destination pairing.
}

// WiFiAccessPoint represents the WiFiAccessPoint schema from the OpenAPI specification
type WiFiAccessPoint struct {
	Age int `json:"age,omitempty"` // The number of milliseconds since this access point was detected.
	Channel int `json:"channel,omitempty"` // The channel over which the client is communication with the access point.
	Macaddress string `json:"macAddress"` // The MAC address of the WiFi node. It's typically called a BSS, BSSID or MAC address. Separators must be `:` (colon).
	Signalstrength int `json:"signalStrength,omitempty"` // The current signal strength measured in dBm.
	Signaltonoiseratio int `json:"signalToNoiseRatio,omitempty"` // The current signal to noise ratio measured in dB.
}

// PlacesQueryAutocompleteResponse represents the PlacesQueryAutocompleteResponse schema from the OpenAPI specification
type PlacesQueryAutocompleteResponse struct {
	Error_message string `json:"error_message,omitempty"` // When the service returns a status code other than `OK`, there may be an additional `error_message` field within the response object. This field contains more detailed information about thereasons behind the given status code. This field is not always returned, and its content is subject to change.
	Info_messages []string `json:"info_messages,omitempty"` // When the service returns additional information about the request specification, there may be an additional `info_messages` field within the response object. This field is only returned for successful requests. It may not always be returned, and its content is subject to change.
	Predictions []PlaceAutocompletePrediction `json:"predictions"` // Contains an array of predictions.
	Status string `json:"status"` // Status codes returned by service. - `OK` indicating the API request was successful. - `ZERO_RESULTS` indicating that the search was successful but returned no results. This may occur if the search was passed a bounds in a remote location. - `INVALID_REQUEST` indicating the API request was malformed, generally due to the missing `input` parameter. - `OVER_QUERY_LIMIT` indicating any of the following: - You have exceeded the QPS limits. - Billing has not been enabled on your account. - The monthly $200 credit, or a self-imposed usage cap, has been exceeded. - The provided method of payment is no longer valid (for example, a credit card has expired). See the [Maps FAQ](https://developers.google.com/maps/faq#over-limit-key-error) for more information about how to resolve this error. - `REQUEST_DENIED` indicating that your request was denied, generally because: - The request is missing an API key. - The `key` parameter is invalid. - `UNKNOWN_ERROR` indicating an unknown error.
}

// PlaceOpeningHoursPeriodDetail represents the PlaceOpeningHoursPeriodDetail schema from the OpenAPI specification
type PlaceOpeningHoursPeriodDetail struct {
	Date string `json:"date,omitempty"` // A date expressed in RFC3339 format in the local timezone for the place, for example 2010-12-31.
	Day float64 `json:"day"` // A number from 0–6, corresponding to the days of the week, starting on Sunday. For example, 2 means Tuesday.
	Time string `json:"time"` // May contain a time of day in 24-hour hhmm format. Values are in the range 0000–2359. The time will be reported in the place’s time zone.
	Truncated bool `json:"truncated,omitempty"` // True if a given period was truncated due to a seven-day cutoff, where the period starts before midnight on the date of the request and/or ends at or after midnight on the last day. This property indicates that the period for open or close can extend past this seven-day cutoff.
}

// NearestRoadsResponse represents the NearestRoadsResponse schema from the OpenAPI specification
type NearestRoadsResponse struct {
	Snappedpoints []SnappedPoint `json:"snappedPoints,omitempty"` // An array of snapped points. Sometimes containing several snapped points for the same point with differing placeId or location.
}

// TimeZoneTextValueObject represents the TimeZoneTextValueObject schema from the OpenAPI specification
type TimeZoneTextValueObject struct {
	Value float64 `json:"value"` // The time specified as Unix time, or seconds since midnight, January 1, 1970 UTC.
	Text string `json:"text"` // The time specified as a string in the time zone.
	Time_zone string `json:"time_zone"` // Contains the time zone. The value is the name of the time zone as defined in the [IANA Time Zone Database](http://www.iana.org/time-zones), e.g. "America/New_York".
}

// DirectionsGeocodedWaypoint represents the DirectionsGeocodedWaypoint schema from the OpenAPI specification
type DirectionsGeocodedWaypoint struct {
	Partial_match interface{} `json:"partial_match,omitempty"` // Indicates that the geocoder did not return an exact match for the original request, though it was able to match part of the requested address. You may wish to examine the original request for misspellings and/or an incomplete address. Partial matches most often occur for street addresses that do not exist within the locality you pass in the request. Partial matches may also be returned when a request matches two or more locations in the same locality. For example, "21 Henr St, Bristol, UK" will return a partial match for both Henry Street and Henrietta Street. Note that if a request includes a misspelled address component, the geocoding service may suggest an alternative address. Suggestions triggered in this way will also be marked as a partial match.
	Place_id string `json:"place_id,omitempty"` // A unique identifier that can be used with other Google APIs. See the [Place ID overview](https://developers.google.com/maps/documentation/places/web-service/place-id).
	Types []string `json:"types,omitempty"` // Indicates the address type of the geocoding result used for calculating directions. * `administrative_area_level_1` indicates a first-order civil entity below the country level. Within the United States, these administrative levels are states. Not all nations exhibit these administrative levels. In most cases, administrative_area_level_1 short names will closely match ISO 3166-2 subdivisions and other widely circulated lists; however this is not guaranteed as our geocoding results are based on a variety of signals and location data. * `administrative_area_level_2` indicates a second-order civil entity below the country level. Within the United States, these administrative levels are counties. Not all nations exhibit these administrative levels. * `administrative_area_level_3` indicates a third-order civil entity below the country level. This type indicates a minor civil division. Not all nations exhibit these administrative levels. * `administrative_area_level_4` indicates a fourth-order civil entity below the country level. This type indicates a minor civil division. Not all nations exhibit these administrative levels. * `administrative_area_level_5` indicates a fifth-order civil entity below the country level. This type indicates a minor civil division. Not all nations exhibit these administrative levels. * `airport` indicates an airport. * `colloquial_area` indicates a commonly-used alternative name for the entity. * `country` indicates the national political entity, and is typically the highest order type returned by the Geocoder. * `intersection` indicates a major intersection, usually of two major roads. * `locality` indicates an incorporated city or town political entity. * `natural_feature` indicates a prominent natural feature. * `neighborhood` indicates a named neighborhood * `park` indicates a named park. * `plus_code` indicates an encoded location reference, derived from latitude and longitude. Plus codes can be used as a replacement for street addresses in places where they do not exist (where buildings are not numbered or streets are not named). See [https://plus.codes](https://plus.codes/) for details. * `point_of_interest` indicates a named point of interest. Typically, these "POI"s are prominent local entities that don't easily fit in another category, such as "Empire State Building" or "Eiffel Tower". * `political` indicates a political entity. Usually, this type indicates a polygon of some civil administration. * `postal_code` indicates a postal code as used to address postal mail within the country. * `premise` indicates a named location, usually a building or collection of buildings with a common name * `route` indicates a named route (such as "US 101"). * `street_address` indicates a precise street address. * `sublocality` indicates a first-order civil entity below a locality. For some locations may receive one of the additional types: sublocality_level_1 to sublocality_level_5. Each sublocality level is a civil entity. Larger numbers indicate a smaller geographic area. * `subpremise` indicates a first-order entity below a named location, usually a singular building within a collection of buildings with a common name * `tourist_attraction` indicates a tourist attraction. * `train_station` indicates a train station. * `transit_station` indicates a transit station. An empty list of types indicates there are no known types for the particular address component, for example, Lieu-dit in France.
	Geocoder_status string `json:"geocoder_status,omitempty"` // Indicates the status code resulting from the geocoding operation. This field may contain the following values.
}

// DistanceMatrixResponse represents the DistanceMatrixResponse schema from the OpenAPI specification
type DistanceMatrixResponse struct {
	Destination_addresses []string `json:"destination_addresses"` // An array of addresses as returned by the API from your original request. As with `origin_addresses`, these are localized if appropriate. This content is meant to be read as-is. Do not programatically parse the formatted addresses.
	Error_message string `json:"error_message,omitempty"` // A string containing the human-readable text of any errors encountered while the request was being processed.
	Origin_addresses []string `json:"origin_addresses"` // An array of addresses as returned by the API from your original request. These are formatted by the geocoder and localized according to the language parameter passed with the request. This content is meant to be read as-is. Do not programatically parse the formatted addresses.
	Rows []DistanceMatrixRow `json:"rows"` // An array of elements, which in turn each contain a `status`, `duration`, and `distance` element.
	Status string `json:"status"` // Status codes returned by service. - `OK` indicates the response contains a valid result. - `INVALID_REQUEST` indicates that the provided request was invalid. - `MAX_ELEMENTS_EXCEEDED` indicates that the product of origins and destinations exceeds the per-query limit. - `MAX_DIMENSIONS_EXCEEDED` indicates that the number of origins or destinations exceeds the per-query limit. - `OVER_DAILY_LIMIT` indicates any of the following: - The API key is missing or invalid. - Billing has not been enabled on your account. - A self-imposed usage cap has been exceeded. - The provided method of payment is no longer valid (for example, a credit card has expired). - `OVER_QUERY_LIMIT` indicates the service has received too many requests from your application within the allowed time period. - `REQUEST_DENIED` indicates that the service denied use of the Distance Matrix service by your application. - `UNKNOWN_ERROR` indicates a Distance Matrix request could not be processed due to a server error. The request may succeed if you try again.
}

// PlacesFindPlaceFromTextResponse represents the PlacesFindPlaceFromTextResponse schema from the OpenAPI specification
type PlacesFindPlaceFromTextResponse struct {
	Info_messages []string `json:"info_messages,omitempty"` // When the service returns additional information about the request specification, there may be an additional `info_messages` field within the response object. This field is only returned for successful requests. It may not always be returned, and its content is subject to change.
	Status string `json:"status"` // Status codes returned by service. - `OK` indicating the API request was successful. - `ZERO_RESULTS` indicating that the search was successful but returned no results. This may occur if the search was passed a `latlng` in a remote location. - `INVALID_REQUEST` indicating the API request was malformed, generally due to missing required query parameter (`location` or `radius`). - `OVER_QUERY_LIMIT` indicating any of the following: - You have exceeded the QPS limits. - Billing has not been enabled on your account. - The monthly $200 credit, or a self-imposed usage cap, has been exceeded. - The provided method of payment is no longer valid (for example, a credit card has expired). See the [Maps FAQ](https://developers.google.com/maps/faq#over-limit-key-error) for more information about how to resolve this error. - `REQUEST_DENIED` indicating that your request was denied, generally because: - The request is missing an API key. - The `key` parameter is invalid. - `UNKNOWN_ERROR` indicating an unknown error.
	Candidates []Place `json:"candidates"` // Contains an array of Place candidates. <div class="caution">Place Search requests return a subset of the fields that are returned by Place Details requests. If the field you want is not returned by Place Search, you can use Place Search to get a place_id, then use that Place ID to make a Place Details request.</div>
	Error_message string `json:"error_message,omitempty"` // When the service returns a status code other than `OK<`, there may be an additional `error_message` field within the response object. This field contains more detailed information about thereasons behind the given status code. This field is not always returned, and its content is subject to change.
}

// GeocodingResponse represents the GeocodingResponse schema from the OpenAPI specification
type GeocodingResponse struct {
	Status string `json:"status"` // The `status` field within the Geocoding response object contains the status of the request, and may contain debugging information to help you track down why geocoding is not working. The "status" field may contain the following values: - `OK` indicates that no errors occurred; the address was successfully parsed and at least one geocode was returned. - `ZERO_RESULTS` indicates that the geocode was successful but returned no results. This may occur if the geocoder was passed a non-existent address or a `latlng` in a remote location. - `OVER_DAILY_LIMIT` indicates any of the following: - The API key is missing or invalid. - Billing has not been enabled on your account. - A self-imposed usage cap has been exceeded. - The provided method of payment is no longer valid (for example, a credit card has expired). - `OVER_QUERY_LIMIT` indicates that you are over your quota. - `REQUEST_DENIED` indicates that your request was denied. - `INVALID_REQUEST` generally indicates that the query (address, components, or latlng) is missing. - `UNKNOWN_ERROR` indicates that the request could not be processed due to a server error. The request may succeed if you try again.
	Error_message string `json:"error_message,omitempty"` // A short description of the error.
	Plus_code PlusCode `json:"plus_code,omitempty"` // An encoded location reference, derived from latitude and longitude coordinates, that represents an area, 1/8000th of a degree by 1/8000th of a degree (about 14m x 14m at the equator) or smaller. Plus codes can be used as a replacement for street addresses in places where they do not exist (where buildings are not numbered or streets are not named).
	Results []GeocodingResult `json:"results"`
}

// DirectionsRoute represents the DirectionsRoute schema from the OpenAPI specification
type DirectionsRoute struct {
	Fare Fare `json:"fare,omitempty"` // The total fare for the route. ``` { "currency" : "USD", "value" : 6, "text" : "$6.00" } ```
	Legs []DirectionsLeg `json:"legs"` // An array which contains information about a leg of the route, between two locations within the given route. A separate leg will be present for each waypoint or destination specified. (A route with no waypoints will contain exactly one leg within the legs array.) Each leg consists of a series of steps.
	Overview_polyline DirectionsPolyline `json:"overview_polyline"` // [Polyline encoding](https://developers.google.com/maps/documentation/utilities/polylinealgorithm) is a lossy compression algorithm that allows you to store a series of coordinates as a single string. Point coordinates are encoded using signed values. If you only have a few static points, you may also wish to use the interactive polyline encoding utility. The encoding process converts a binary value into a series of character codes for ASCII characters using the familiar base64 encoding scheme: to ensure proper display of these characters, encoded values are summed with 63 (the ASCII character '?') before converting them into ASCII. The algorithm also checks for additional character codes for a given point by checking the least significant bit of each byte group; if this bit is set to 1, the point is not yet fully formed and additional data must follow. Additionally, to conserve space, points only include the offset from the previous point (except of course for the first point). All points are encoded in Base64 as signed integers, as latitudes and longitudes are signed values. The encoding format within a polyline needs to represent two coordinates representing latitude and longitude to a reasonable precision. Given a maximum longitude of +/- 180 degrees to a precision of 5 decimal places (180.00000 to -180.00000), this results in the need for a 32 bit signed binary integer value.
	Summary string `json:"summary"` // Contains a short textual description for the route, suitable for naming and disambiguating the route from alternatives.
	Warnings []string `json:"warnings"` // Contains an array of warnings to be displayed when showing these directions. You must handle and display these warnings yourself.
	Waypoint_order []int `json:"waypoint_order"` // An array indicating the order of any waypoints in the calculated route. This waypoints may be reordered if the request was passed optimize:true within its waypoints parameter.
	Bounds Bounds `json:"bounds"` // A rectangle in geographical coordinates from points at the southwest and northeast corners.
	Copyrights string `json:"copyrights"` // Contains the copyright notices to be displayed for this route. You must handle and display this information yourself. This content is meant to be read as-is. Do not programmatically parse this display-only content.
}

// ErrorDetail represents the ErrorDetail schema from the OpenAPI specification
type ErrorDetail struct {
	TypeField string `json:"@type,omitempty"` // The type of error.
	Domain string `json:"domain,omitempty"` // The domain in which the error occurred.
	Fieldviolations []FieldViolation `json:"fieldViolations,omitempty"` // A list of field violations.
	Message string `json:"message,omitempty"` // A short description of the error.
	Metadata map[string]interface{} `json:"metadata,omitempty"` // Additional metadata about the error.
	Reason string `json:"reason,omitempty"` // A reason for the error.
}

// PlaceReview represents the PlaceReview schema from the OpenAPI specification
type PlaceReview struct {
	Relative_time_description string `json:"relative_time_description"` // The time that the review was submitted in text, relative to the current time.
	Time float64 `json:"time"` // The time that the review was submitted, measured in the number of seconds since since midnight, January 1, 1970 UTC.
	Original_language string `json:"original_language,omitempty"` // An IETF language code indicating the original language of the review. If the review has been translated, then `original_language` != `language`. This field contains the main language tag only, and not the secondary tag indicating country or region. For example, all the English reviews are tagged as 'en', and not 'en-AU' or 'en-UK' and so on. This field is empty if there is only a rating with no review text.
	Translated bool `json:"translated,omitempty"` // A boolean value indicating if the review was translated from the original language it was written in. If a review has been translated, corresponding to a value of true, Google recommends that you indicate this to your users. For example, you can add the following string, “Translated by Google”, to the review.
	Language string `json:"language,omitempty"` // An IETF language code indicating the language of the returned review. This field contains the main language tag only, and not the secondary tag indicating country or region. For example, all the English reviews are tagged as 'en', and not 'en-AU' or 'en-UK' and so on. This field is empty if there is only a rating with no review text.
	Profile_photo_url string `json:"profile_photo_url,omitempty"` // The URL to the user's profile photo, if available.
	Text string `json:"text,omitempty"` // The user's review. When reviewing a location with Google Places, text reviews are considered optional. Therefore, this field may be empty. Note that this field may include simple HTML markup. For example, the entity reference `&amp;` may represent an ampersand character.
	Author_name string `json:"author_name"` // The name of the user who submitted the review. Anonymous reviews are attributed to "A Google user".
	Author_url string `json:"author_url,omitempty"` // The URL to the user's Google Maps Local Guides profile, if available.
	Rating float64 `json:"rating"` // The user's overall rating for this place. This is a whole number, ranging from 1 to 5.
}

// PlacesTextSearchResponse represents the PlacesTextSearchResponse schema from the OpenAPI specification
type PlacesTextSearchResponse struct {
	Info_messages []string `json:"info_messages,omitempty"` // When the service returns additional information about the request specification, there may be an additional `info_messages` field within the response object. This field is only returned for successful requests. It may not always be returned, and its content is subject to change.
	Next_page_token string `json:"next_page_token,omitempty"` // Contains a token that can be used to return up to 20 additional results. A next_page_token will not be returned if there are no additional results to display. The maximum number of results that can be returned is 60. There is a short delay between when a next_page_token is issued, and when it will become valid.
	Results []Place `json:"results"` // Contains an array of places. <div class="caution">Place Search requests return a subset of the fields that are returned by Place Details requests. If the field you want is not returned by Place Search, you can use Place Search to get a `place_id`, then use that Place ID to make a Place Details request.</div>
	Status string `json:"status"` // Status codes returned by service. - `OK` indicating the API request was successful. - `ZERO_RESULTS` indicating that the search was successful but returned no results. This may occur if the search was passed a `latlng` in a remote location. - `INVALID_REQUEST` indicating the API request was malformed, generally due to missing required query parameter (`location` or `radius`). - `OVER_QUERY_LIMIT` indicating any of the following: - You have exceeded the QPS limits. - Billing has not been enabled on your account. - The monthly $200 credit, or a self-imposed usage cap, has been exceeded. - The provided method of payment is no longer valid (for example, a credit card has expired). See the [Maps FAQ](https://developers.google.com/maps/faq#over-limit-key-error) for more information about how to resolve this error. - `REQUEST_DENIED` indicating that your request was denied, generally because: - The request is missing an API key. - The `key` parameter is invalid. - `UNKNOWN_ERROR` indicating an unknown error.
	Error_message string `json:"error_message,omitempty"` // When the service returns a status code other than `OK<`, there may be an additional `error_message` field within the response object. This field contains more detailed information about thereasons behind the given status code. This field is not always returned, and its content is subject to change.
	Html_attributions []string `json:"html_attributions"` // May contain a set of attributions about this listing which must be displayed to the user (some listings may not have attribution).
}

// DirectionsTransitAgency represents the DirectionsTransitAgency schema from the OpenAPI specification
type DirectionsTransitAgency struct {
	Name string `json:"name,omitempty"` // The name of this transit agency.
	Phone string `json:"phone,omitempty"` // The transit agency's phone number.
	Url string `json:"url,omitempty"` // The transit agency's URL.
}

// PlacesAutocompleteResponse represents the PlacesAutocompleteResponse schema from the OpenAPI specification
type PlacesAutocompleteResponse struct {
	Error_message string `json:"error_message,omitempty"` // When the service returns a status code other than `OK<`, there may be an additional `error_message` field within the response object. This field contains more detailed information about thereasons behind the given status code. This field is not always returned, and its content is subject to change.
	Info_messages []string `json:"info_messages,omitempty"` // When the service returns additional information about the request specification, there may be an additional `info_messages` field within the response object. This field is only returned for successful requests. It may not always be returned, and its content is subject to change.
	Predictions []PlaceAutocompletePrediction `json:"predictions"` // Contains an array of predictions.
	Status string `json:"status"` // Status codes returned by service. - `OK` indicating the API request was successful. - `ZERO_RESULTS` indicating that the search was successful but returned no results. This may occur if the search was passed a bounds in a remote location. - `INVALID_REQUEST` indicating the API request was malformed, generally due to the missing `input` parameter. - `OVER_QUERY_LIMIT` indicating any of the following: - You have exceeded the QPS limits. - Billing has not been enabled on your account. - The monthly $200 credit, or a self-imposed usage cap, has been exceeded. - The provided method of payment is no longer valid (for example, a credit card has expired). See the [Maps FAQ](https://developers.google.com/maps/faq#over-limit-key-error) for more information about how to resolve this error. - `REQUEST_DENIED` indicating that your request was denied, generally because: - The request is missing an API key. - The `key` parameter is invalid. - `UNKNOWN_ERROR` indicating an unknown error.
}
