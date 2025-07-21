package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// Currencydc697bd5bc1b1000085ab3d0e5da0108 represents the Currencydc697bd5bc1b1000085ab3d0e5da0108 schema from the OpenAPI specification
type Currencydc697bd5bc1b1000085ab3d0e5da0108 struct {
	Descriptor string `json:"descriptor,omitempty"` // A description of the instance
	Href string `json:"href,omitempty"` // A link to the instance
	Id string `json:"id"` // wid / id / reference id
}

// PayComponentd4198c33f48d10000ec87ef21d4700f1 represents the PayComponentd4198c33f48d10000ec87ef21d4700f1 schema from the OpenAPI specification
type PayComponentd4198c33f48d10000ec87ef21d4700f1 struct {
	Code string `json:"code,omitempty"` // The payroll code of the pay component.
	Descriptor string `json:"descriptor,omitempty"` // A preview of the instance
	Id string `json:"id,omitempty"` // Id of the instance
}

// NextPeriodToProcessba5ffacfc4e210000c4f57abf59d0120 represents the NextPeriodToProcessba5ffacfc4e210000c4f57abf59d0120 schema from the OpenAPI specification
type NextPeriodToProcessba5ffacfc4e210000c4f57abf59d0120 struct {
	Descriptor string `json:"descriptor,omitempty"` // A description of the instance
	Href string `json:"href,omitempty"` // A link to the instance
	Id string `json:"id"` // wid / id / reference id
}

// ERRORMODELREFERENCE represents the ERRORMODELREFERENCE schema from the OpenAPI specification
type ERRORMODELREFERENCE struct {
	Error string `json:"error"` // A description of the error
	Field string `json:"field,omitempty"` // The field related to the error
	Path string `json:"path,omitempty"` // The path of the field related to the error
	Code string `json:"code,omitempty"` // The code that corresponds to the error message. Use the error code to drive programmatic error-handling behavior. Don't use error message strings for this purpose because they are subject to change
}

// Company78c55a3bc2cc10000a9d95f9824e00f9 represents the Company78c55a3bc2cc10000a9d95f9824e00f9 schema from the OpenAPI specification
type Company78c55a3bc2cc10000a9d95f9824e00f9 struct {
	Company string `json:"company,omitempty"` // The Reference ID to use for lookups within our Workday Web Services. For supervisory organizations, this is also the Organization_ID.
	Descriptor string `json:"descriptor,omitempty"` // A preview of the instance
	Fein string `json:"fein,omitempty"` // The FEIN for the US company.
	Id string `json:"id,omitempty"` // Id of the instance
}

// JobProfileJobViewfab4a151b24810000d511d61ee641262 represents the JobProfileJobViewfab4a151b24810000d511d61ee641262 schema from the OpenAPI specification
type JobProfileJobViewfab4a151b24810000d511d61ee641262 struct {
	Descriptor string `json:"descriptor,omitempty"` // A preview of the instance
	Id string `json:"id,omitempty"` // Id of the instance
}

// PayRunGroupSimpleViewba5ffacfc4e210000de956ae16b40122 represents the PayRunGroupSimpleViewba5ffacfc4e210000de956ae16b40122 schema from the OpenAPI specification
type PayRunGroupSimpleViewba5ffacfc4e210000de956ae16b40122 struct {
	Descriptor string `json:"descriptor,omitempty"` // A preview of the instance
	Id string `json:"id,omitempty"` // Id of the instance
}

// RunCategory02313d5912d210000a19cc6245a2010c represents the RunCategory02313d5912d210000a19cc6245a2010c schema from the OpenAPI specification
type RunCategory02313d5912d210000a19cc6245a2010c struct {
	Descriptor string `json:"descriptor,omitempty"` // A description of the instance
	Href string `json:"href,omitempty"` // A link to the instance
	Id string `json:"id"` // wid / id / reference id
}

// Typedc697bd5bc1b10001427adc082700133 represents the Typedc697bd5bc1b10001427adc082700133 schema from the OpenAPI specification
type Typedc697bd5bc1b10001427adc082700133 struct {
	Name string `json:"name,omitempty"` // The alternate ID of the related calculation for the pay component and pay component related calculation.
	Descriptor string `json:"descriptor,omitempty"` // A preview of the instance
	Id string `json:"id,omitempty"` // Id of the instance
}

// CompanySuiRateDetails985b8306e95d100002daf466c84d0005 represents the CompanySuiRateDetails985b8306e95d100002daf466c84d0005 schema from the OpenAPI specification
type CompanySuiRateDetails985b8306e95d100002daf466c84d0005 struct {
	StartDate string `json:"startDate"` // The start date for company tax reporting.
	EndDate string `json:"endDate,omitempty"` // The end date for company tax reporting.
	TaxCode string `json:"taxCode"` // The deduction for company tax reporting.
	Ein string `json:"ein,omitempty"` // The payroll tax authority EIN field for company tax reporting.
	ApplicableRate string `json:"applicableRate,omitempty"` // The tax override rate for company tax reporting.
	CompanyInstance CompanyInstance0b4d146f73561000154ef82e7d1300fa `json:"companyInstance"` // The company object for company tax reporting.
	Exempt bool `json:"exempt,omitempty"` // If true, the SUI rate is exempt.
	StateInstance StateInstance0b4d146f73561000154ef7f3f09900f9 `json:"stateInstance"` // The payroll tax authority object for company tax reporting.
	Id string `json:"id,omitempty"` // Id of the instance
}

// CurrentPeriodInProgressba5ffacfc4e210000c4724393162011e represents the CurrentPeriodInProgressba5ffacfc4e210000c4724393162011e schema from the OpenAPI specification
type CurrentPeriodInProgressba5ffacfc4e210000c4724393162011e struct {
	Descriptor string `json:"descriptor,omitempty"` // A description of the instance
	Href string `json:"href,omitempty"` // A link to the instance
	Id string `json:"id"` // wid / id / reference id
}

// CompanyInstance0b4d146f73561000154ef82e7d1300fa represents the CompanyInstance0b4d146f73561000154ef82e7d1300fa schema from the OpenAPI specification
type CompanyInstance0b4d146f73561000154ef82e7d1300fa struct {
	Id string `json:"id,omitempty"` // Id of the instance
}

// PeriodScheduleba5ffacfc4e210000bfed14c65fb0118 represents the PeriodScheduleba5ffacfc4e210000bfed14c65fb0118 schema from the OpenAPI specification
type PeriodScheduleba5ffacfc4e210000bfed14c65fb0118 struct {
	Descriptor string `json:"descriptor,omitempty"` // A description of the instance
	Href string `json:"href,omitempty"` // A link to the instance
	Id string `json:"id"` // wid / id / reference id
}

// WorkerJobViewfab4a151b24810000e7ff8d0c7f0126d represents the WorkerJobViewfab4a151b24810000e7ff8d0c7f0126d schema from the OpenAPI specification
type WorkerJobViewfab4a151b24810000e7ff8d0c7f0126d struct {
	Id string `json:"id,omitempty"` // Id of the instance
	Descriptor string `json:"descriptor,omitempty"` // A preview of the instance
}

// InputLinedc697bd5bc1b10001050f53a976e0119 represents the InputLinedc697bd5bc1b10001050f53a976e0119 schema from the OpenAPI specification
type InputLinedc697bd5bc1b10001050f53a976e0119 struct {
	Value int `json:"value,omitempty"` // The value for the input line.
	Id string `json:"id,omitempty"` // Id of the instance
	Type Typedc697bd5bc1b10001427adc082700133 `json:"type,omitempty"` // The related calculation type for the input line.
}

// Country8ae53f6e8b8a100007a313ae869e0d65 represents the Country8ae53f6e8b8a100007a313ae869e0d65 schema from the OpenAPI specification
type Country8ae53f6e8b8a100007a313ae869e0d65 struct {
	Descriptor string `json:"descriptor,omitempty"` // A description of the instance
	Href string `json:"href,omitempty"` // A link to the instance
	Id string `json:"id"` // wid / id / reference id
}

// PayComponentdc697bd5bc1b100005f310c75c8e00f9 represents the PayComponentdc697bd5bc1b100005f310c75c8e00f9 schema from the OpenAPI specification
type PayComponentdc697bd5bc1b100005f310c75c8e00f9 struct {
	Descriptor string `json:"descriptor,omitempty"` // A preview of the instance
	Id string `json:"id,omitempty"` // Id of the instance
	Code string `json:"code,omitempty"` // The payroll code of the pay component.
}

// INSTANCEMODELREFERENCE represents the INSTANCEMODELREFERENCE schema from the OpenAPI specification
type INSTANCEMODELREFERENCE struct {
	Descriptor string `json:"descriptor,omitempty"` // A description of the instance
	Href string `json:"href,omitempty"` // A link to the instance
	Id string `json:"id"` // wid / id / reference id
}

// Position95417070df52100038a792408e430185 represents the Position95417070df52100038a792408e430185 schema from the OpenAPI specification
type Position95417070df52100038a792408e430185 struct {
	Descriptor string `json:"descriptor,omitempty"` // A description of the instance
	Href string `json:"href,omitempty"` // A link to the instance
	Id string `json:"id"` // wid / id / reference id
}

// InputInterface14e354bf8650100010fc3cbd9c520123 represents the InputInterface14e354bf8650100010fc3cbd9c520123 schema from the OpenAPI specification
type InputInterface14e354bf8650100010fc3cbd9c520123 struct {
	Descriptor string `json:"descriptor,omitempty"` // A preview of the instance
	Id string `json:"id,omitempty"` // Id of the instance
	Name string `json:"name,omitempty"` // The alternate ID of the related calculation for the pay component and pay component related calculation.
}

// SupervisoryOrganizationfab4a151b24810000d13073c5d341257 represents the SupervisoryOrganizationfab4a151b24810000d13073c5d341257 schema from the OpenAPI specification
type SupervisoryOrganizationfab4a151b24810000d13073c5d341257 struct {
	Id string `json:"id,omitempty"` // Id of the instance
	Descriptor string `json:"descriptor,omitempty"` // A preview of the instance
}

// StateInstance0b4d146f73561000154ef7f3f09900f9 represents the StateInstance0b4d146f73561000154ef7f3f09900f9 schema from the OpenAPI specification
type StateInstance0b4d146f73561000154ef7f3f09900f9 struct {
	Id string `json:"id,omitempty"` // Id of the instance
}

// Companyc15dfab2b81e100008cee72ff399011e represents the Companyc15dfab2b81e100008cee72ff399011e schema from the OpenAPI specification
type Companyc15dfab2b81e100008cee72ff399011e struct {
	Id string `json:"id,omitempty"` // Id of the instance
}

// Workerfab4a151b24810000e66f769304c126b represents the Workerfab4a151b24810000e66f769304c126b schema from the OpenAPI specification
type Workerfab4a151b24810000e66f769304c126b struct {
	Id string `json:"id,omitempty"` // Id of the instance
	Descriptor string `json:"descriptor,omitempty"` // A preview of the instance
}

// FirstProcessingPeriodba5ffacfc4e210000e3c99501f9f012a represents the FirstProcessingPeriodba5ffacfc4e210000e3c99501f9f012a schema from the OpenAPI specification
type FirstProcessingPeriodba5ffacfc4e210000e3c99501f9f012a struct {
	Descriptor string `json:"descriptor,omitempty"` // A description of the instance
	Href string `json:"href,omitempty"` // A link to the instance
	Id string `json:"id"` // wid / id / reference id
}

// SupervisoryOrganizationJobViewfab4a151b24810000d29746fb7e21259 represents the SupervisoryOrganizationJobViewfab4a151b24810000d29746fb7e21259 schema from the OpenAPI specification
type SupervisoryOrganizationJobViewfab4a151b24810000d29746fb7e21259 struct {
	Descriptor string `json:"descriptor,omitempty"` // A preview of the instance
	Id string `json:"id,omitempty"` // Id of the instance
}

// SuiRatesSummaryb2c5e6f505c3100018bb11b51393015b represents the SuiRatesSummaryb2c5e6f505c3100018bb11b51393015b schema from the OpenAPI specification
type SuiRatesSummaryb2c5e6f505c3100018bb11b51393015b struct {
	EinType string `json:"einType,omitempty"` // The EIN type. Valid values: SUI EIN, STATE EIN, FEIN.
	Exempt bool `json:"exempt,omitempty"` // If true, the SUI rate is exempt.
	Id string `json:"id,omitempty"` // Unique identifier for company SUI rate.
	StartDate string `json:"startDate,omitempty"` // The start date value of the row for company SUI rate.
	ApplicableRate string `json:"applicableRate,omitempty"` // The applicable rate for the company SUI.
	RateType string `json:"rateType,omitempty"` // The rate type. Valid values: OR for override rate, DR for default rate.
	EndDate string `json:"endDate,omitempty"` // The end date value of the row for company SUI rate.
	StateInstance StateInstance78c55a3bc2cc10000d3989d389220115 `json:"stateInstance,omitempty"` // The state instance which includes the ID, descriptor, and state reference ID.
	CompanyInstance CompanyInstance78c55a3bc2cc100009b2cec7cb7e00f7 `json:"companyInstance,omitempty"` // The company instance which includes the ID, descriptor, and company reference ID.
	TaxCode string `json:"taxCode,omitempty"` // The payroll tax code. As of v1, the default value is W_SUIER.
	Ein string `json:"ein,omitempty"` // The EIN value for the company SUI rate.
}

// RunCategoryba5ffacfc4e210000bee6ddb72e00116 represents the RunCategoryba5ffacfc4e210000bee6ddb72e00116 schema from the OpenAPI specification
type RunCategoryba5ffacfc4e210000bee6ddb72e00116 struct {
	Href string `json:"href,omitempty"` // A link to the instance
	Id string `json:"id"` // wid / id / reference id
	Descriptor string `json:"descriptor,omitempty"` // A description of the instance
}

// PayGroupDetailViewe5811059356910001433d1f0fa2e2256 represents the PayGroupDetailViewe5811059356910001433d1f0fa2e2256 schema from the OpenAPI specification
type PayGroupDetailViewe5811059356910001433d1f0fa2e2256 struct {
	PayGroup PayGroupca4f510c6a39100016f4913a34214f4a `json:"payGroup,omitempty"` // Name of the pay group organization.
	PayRunGroup []PayRunGroupSimpleViewba5ffacfc4e210000de956ae16b40122 `json:"payRunGroup,omitempty"` // Contains one or more pay groups that are processed or run at the same time.
	RunCategory RunCategoryba5ffacfc4e210000bee6ddb72e00116 `json:"runCategory,omitempty"` // The run category.
	LastPeriodCompleted LastPeriodCompletedba5ffacfc4e210000c384329175c011c `json:"lastPeriodCompleted,omitempty"` // The prior period processed including period begin date, period end date, and frequency.
	Id string `json:"id,omitempty"` // Id of the instance
	NextPeriodToProcess NextPeriodToProcessba5ffacfc4e210000c4f57abf59d0120 `json:"nextPeriodToProcess,omitempty"` // The next period to process, including period begin date, period end date, and frequency.
	PeriodSchedule PeriodScheduleba5ffacfc4e210000bfed14c65fb0118 `json:"periodSchedule,omitempty"` // The period schedule for the pay group/run category combination.
	CurrentPeriodInProgress CurrentPeriodInProgressba5ffacfc4e210000c4724393162011e `json:"currentPeriodInProgress,omitempty"` // The current period in progress including the period begin date, period end date, and frequency.
	Descriptor string `json:"descriptor,omitempty"` // A preview of the instance
	FirstProcessingPeriod FirstProcessingPeriodba5ffacfc4e210000e3c99501f9f012a `json:"firstProcessingPeriod,omitempty"` // The first processing period for the pay group detail.
}

// LastPeriodCompletedba5ffacfc4e210000c384329175c011c represents the LastPeriodCompletedba5ffacfc4e210000c384329175c011c schema from the OpenAPI specification
type LastPeriodCompletedba5ffacfc4e210000c384329175c011c struct {
	Descriptor string `json:"descriptor,omitempty"` // A description of the instance
	Href string `json:"href,omitempty"` // A link to the instance
	Id string `json:"id"` // wid / id / reference id
}

// Statec15dfab2b81e10000960777a67000121 represents the Statec15dfab2b81e10000960777a67000121 schema from the OpenAPI specification
type Statec15dfab2b81e10000960777a67000121 struct {
	Id string `json:"id,omitempty"` // Id of the instance
}

// Worker95417070df5210000dc1bb7f05b8012c represents the Worker95417070df5210000dc1bb7f05b8012c schema from the OpenAPI specification
type Worker95417070df5210000dc1bb7f05b8012c struct {
	Descriptor string `json:"descriptor,omitempty"` // A preview of the instance
	Id string `json:"id,omitempty"` // Id of the instance
}

// CompanyInstance78c55a3bc2cc100009b2cec7cb7e00f7 represents the CompanyInstance78c55a3bc2cc100009b2cec7cb7e00f7 schema from the OpenAPI specification
type CompanyInstance78c55a3bc2cc100009b2cec7cb7e00f7 struct {
	Id string `json:"id,omitempty"` // Id of the instance
	Company string `json:"company,omitempty"` // The Reference ID to use for lookups within our Workday Web Services. For supervisory organizations, this is also the Organization_ID.
	Descriptor string `json:"descriptor,omitempty"` // A preview of the instance
	Fein string `json:"fein,omitempty"` // The FEIN for the US company.
}

// PayGroupViewa35ef5a5cdfe1000117cd514daa600f3 represents the PayGroupViewa35ef5a5cdfe1000117cd514daa600f3 schema from the OpenAPI specification
type PayGroupViewa35ef5a5cdfe1000117cd514daa600f3 struct {
	Country Country8ae53f6e8b8a100007a313ae869e0d65 `json:"country,omitempty"` // Contains the country or territory.
	Descriptor string `json:"descriptor,omitempty"` // A preview of the instance
	Id string `json:"id,omitempty"` // Id of the instance
	PayGroupDetails []PayGroupDetailMiniView02313d5912d210000a19cc195ef60105 `json:"payGroupDetails,omitempty"` // Contains the pay group detail.
}

// LocationDataJobView6d3eb084b4c410002b5fa13f0c9d0056 represents the LocationDataJobView6d3eb084b4c410002b5fa13f0c9d0056 schema from the OpenAPI specification
type LocationDataJobView6d3eb084b4c410002b5fa13f0c9d0056 struct {
	Country Countryaf21e47ff7c01000092ecf4d7ad30270 `json:"country,omitempty"` // Returns the country from the primary address for the location.
	Descriptor string `json:"descriptor,omitempty"` // A preview of the instance
	Id string `json:"id,omitempty"` // Id of the instance
}

// State78c55a3bc2cc10000d85e9bcf66a011d represents the State78c55a3bc2cc10000d85e9bcf66a011d schema from the OpenAPI specification
type State78c55a3bc2cc10000d85e9bcf66a011d struct {
	Descriptor string `json:"descriptor,omitempty"` // A preview of the instance
	Id string `json:"id,omitempty"` // Id of the instance
	PayrollStateAuthorityTaxCode string `json:"payrollStateAuthorityTaxCode,omitempty"` // The Payroll Authority Tax Code for a Payroll Tax Authority.
}

// Workerdc697bd5bc1b100005f310b2da6e00f7 represents the Workerdc697bd5bc1b100005f310b2da6e00f7 schema from the OpenAPI specification
type Workerdc697bd5bc1b100005f310b2da6e00f7 struct {
	Descriptor string `json:"descriptor,omitempty"` // A preview of the instance
	Id string `json:"id,omitempty"` // Id of the instance
}

// PayrollInputViewdc697bd5bc1b100005f3109b925b00f6 represents the PayrollInputViewdc697bd5bc1b100005f3109b925b00f6 schema from the OpenAPI specification
type PayrollInputViewdc697bd5bc1b100005f3109b925b00f6 struct {
	Ongoing bool `json:"ongoing,omitempty"` // If true, the payroll input is ongoing.
	InputDetails []InputLinedc697bd5bc1b10001050f53a976e0119 `json:"inputDetails,omitempty"` // The details for this payroll input.
	StartDate string `json:"startDate,omitempty"` // The start date before which this input does not apply.
	FieldEditability string `json:"fieldEditability,omitempty"` // The editability status indicating the fields that can be updated in the payroll input request. Possible values: all, none, endDateOnly
	Worker Workerdc697bd5bc1b100005f310b2da6e00f7 `json:"worker,omitempty"` // The worker for this payroll input.
	Position Position95417070df52100038a792408e430185 `json:"position,omitempty"` // The worker's position the payroll input applies to if Multi Job Payroll is used.
	Currency Currencydc697bd5bc1b1000085ab3d0e5da0108 `json:"currency,omitempty"` // The currency for the payroll input. If no currency exists, the system assumes the Pay Group currency. The Pay Group currency is derived from the default currency for the Pay Group country.
	Comment string `json:"comment,omitempty"` // The text comment for this input.
	RunCategories []RunCategorydc697bd5bc1b10001036377098ea0116 `json:"runCategories,omitempty"` // The run category for the payroll input.
	Worktags []Worktagdc697bd5bc1b1000105a22eb21b8011b `json:"worktags,omitempty"` // The worktags associated with the payroll input.
	Adjustment bool `json:"adjustment,omitempty"` // If true, the input is for an adjustment as opposed to an override.
	PayComponent PayComponentdc697bd5bc1b100005f310c75c8e00f9 `json:"payComponent,omitempty"` // The pay component for this payroll input.
	Id string `json:"id,omitempty"` // Id of the instance
	Descriptor string `json:"descriptor,omitempty"` // A preview of the instance
	EndDate string `json:"endDate,omitempty"` // The end date after which this input does not apply.
}

// JobData1bfa8925c50510000ae4479925510026 represents the JobData1bfa8925c50510000ae4479925510026 schema from the OpenAPI specification
type JobData1bfa8925c50510000ae4479925510026 struct {
	Worker Workerfab4a151b24810000e66f769304c126b `json:"worker,omitempty"` // The position management worker filling the position. If the the position is overlapped, return the worker for the overlap position.
	BusinessTitle string `json:"businessTitle,omitempty"` // The business title for the position.
	Id string `json:"id,omitempty"` // Id of the instance
	JobProfile JobProfile6fb921f8a11d10001d5268980bbb0097 `json:"jobProfile,omitempty"` // The job profile for the position.
	JobType JobTypeb8ac205877fd100005c55b45c6400057 `json:"jobType,omitempty"` // The position's job type which can be one of the following values: Primary, Additional Job, or International Assignment. If the position is not effective, no value will be returned. Primary and additional job are effective dated due to the Switch Primary Job business process. An international assignment cannot be switched to a be a primary job.
	Descriptor string `json:"descriptor,omitempty"` // A preview of the instance
	NextPayPeriodStartDate string `json:"nextPayPeriodStartDate,omitempty"` // The next pay period start date for the job.
	SupervisoryOrganization SupervisoryOrganizationfab4a151b24810000d13073c5d341257 `json:"supervisoryOrganization,omitempty"` // The supervisory organization for the position.
	Location Location6d3eb084b4c410002b617efb943f0059 `json:"location,omitempty"` // The location associated with the filled position or position restrictions. When associated with position restrictions, the location returned is in this order: the location on the effective job requisition, hiring restrictions, or primary location of the supervisory organization.
}

// Countryaf21e47ff7c01000092ecf4d7ad30270 represents the Countryaf21e47ff7c01000092ecf4d7ad30270 schema from the OpenAPI specification
type Countryaf21e47ff7c01000092ecf4d7ad30270 struct {
	Descriptor string `json:"descriptor,omitempty"` // A preview of the instance
}

// StateInstance78c55a3bc2cc10000d3989d389220115 represents the StateInstance78c55a3bc2cc10000d3989d389220115 schema from the OpenAPI specification
type StateInstance78c55a3bc2cc10000d3989d389220115 struct {
	Id string `json:"id,omitempty"` // Id of the instance
	PayrollStateAuthorityTaxCode string `json:"payrollStateAuthorityTaxCode,omitempty"` // The Payroll Authority Tax Code for a Payroll Tax Authority.
	Descriptor string `json:"descriptor,omitempty"` // A preview of the instance
}

// VALIDATIONERRORMODELREFERENCE represents the VALIDATIONERRORMODELREFERENCE schema from the OpenAPI specification
type VALIDATIONERRORMODELREFERENCE struct {
	Error string `json:"error"` // A description of the error
	Errors []ERRORMODELREFERENCE `json:"errors,omitempty"` // An array of validation errors
}

// PayGroupSimpleViewca4f510c6a39100016f0c406e4584f47 represents the PayGroupSimpleViewca4f510c6a39100016f0c406e4584f47 schema from the OpenAPI specification
type PayGroupSimpleViewca4f510c6a39100016f0c406e4584f47 struct {
	Id string `json:"id,omitempty"` // Id of the instance
	Descriptor string `json:"descriptor,omitempty"` // A preview of the instance
}

// JobTypeb8ac205877fd100005c55b45c6400057 represents the JobTypeb8ac205877fd100005c55b45c6400057 schema from the OpenAPI specification
type JobTypeb8ac205877fd100005c55b45c6400057 struct {
	Descriptor string `json:"descriptor,omitempty"` // A preview of the instance
}

// Worktagdc697bd5bc1b1000105a22eb21b8011b represents the Worktagdc697bd5bc1b1000105a22eb21b8011b schema from the OpenAPI specification
type Worktagdc697bd5bc1b1000105a22eb21b8011b struct {
	Descriptor string `json:"descriptor,omitempty"` // A preview of the instance
	Id string `json:"id,omitempty"` // Id of the instance
}

// RunCategorydc697bd5bc1b10001036377098ea0116 represents the RunCategorydc697bd5bc1b10001036377098ea0116 schema from the OpenAPI specification
type RunCategorydc697bd5bc1b10001036377098ea0116 struct {
	Descriptor string `json:"descriptor,omitempty"` // A preview of the instance
	Id string `json:"id,omitempty"` // Id of the instance
}

// Location6d3eb084b4c410002b617efb943f0059 represents the Location6d3eb084b4c410002b617efb943f0059 schema from the OpenAPI specification
type Location6d3eb084b4c410002b617efb943f0059 struct {
	Country Countryaf21e47ff7c01000092ecf4d7ad30270 `json:"country,omitempty"` // Returns the country from the primary address for the location.
	Descriptor string `json:"descriptor,omitempty"` // A preview of the instance
	Id string `json:"id,omitempty"` // Id of the instance
}

// JobProfile6fb921f8a11d10001d5268980bbb0097 represents the JobProfile6fb921f8a11d10001d5268980bbb0097 schema from the OpenAPI specification
type JobProfile6fb921f8a11d10001d5268980bbb0097 struct {
	Id string `json:"id,omitempty"` // Id of the instance
	Descriptor string `json:"descriptor,omitempty"` // A preview of the instance
}

// CountryDataForLocationDataaf21e47ff7c01000092e0e4f9a22026e represents the CountryDataForLocationDataaf21e47ff7c01000092e0e4f9a22026e schema from the OpenAPI specification
type CountryDataForLocationDataaf21e47ff7c01000092e0e4f9a22026e struct {
	Descriptor string `json:"descriptor,omitempty"` // A preview of the instance
}

// MULTIPLEINSTANCEMODELREFERENCE represents the MULTIPLEINSTANCEMODELREFERENCE schema from the OpenAPI specification
type MULTIPLEINSTANCEMODELREFERENCE struct {
	Data []INSTANCEMODELREFERENCE `json:"data,omitempty"`
	Total int `json:"total,omitempty"`
}

// PayGroupca4f510c6a39100016f4913a34214f4a represents the PayGroupca4f510c6a39100016f4913a34214f4a schema from the OpenAPI specification
type PayGroupca4f510c6a39100016f4913a34214f4a struct {
	Descriptor string `json:"descriptor,omitempty"` // A preview of the instance
	Id string `json:"id,omitempty"` // Id of the instance
}

// JobTypeDatab8ac205877fd100005c3b50b74a30055 represents the JobTypeDatab8ac205877fd100005c3b50b74a30055 schema from the OpenAPI specification
type JobTypeDatab8ac205877fd100005c3b50b74a30055 struct {
	Descriptor string `json:"descriptor,omitempty"` // A preview of the instance
}

// PayGroupDetailMiniView02313d5912d210000a19cc195ef60105 represents the PayGroupDetailMiniView02313d5912d210000a19cc195ef60105 schema from the OpenAPI specification
type PayGroupDetailMiniView02313d5912d210000a19cc195ef60105 struct {
	Descriptor string `json:"descriptor,omitempty"` // A preview of the instance
	Id string `json:"id,omitempty"` // Id of the instance
	RunCategory RunCategory02313d5912d210000a19cc6245a2010c `json:"runCategory,omitempty"` // The run category.
}
