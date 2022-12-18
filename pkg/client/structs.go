package client

type ReadPropertyRequests struct {
	Requests []*ReadPropertyRequest `json:"requests"`
}

type ReadPropertyRequest struct {
	Reference *PropertyReference     `json:"reference"`
	Params    map[string]interface{} `json:"params"`
}

type PropertyReferenceValuePairs struct {
	Pairs []*PropertyReferenceValuePair `json:"pairs"`
}

type PropertyReferenceValuePair struct {
	Reference *PropertyReference `json:"reference"`
	Value     *PropertyValue     `json:"value"`
}

type ExecutionRequest struct {
	Operation   Operation                   `json:"operation"`
	Environment string                      `json:"environment"`
	Requests    []*PropertyExecutionRequest `json:"requests"`
}

type ExecutionResponse struct {
	Responses []*PropertyExecutionResponse `json:"responses"`
}

type PropertyExecutionRequest struct {
	Property *PropertyReference     `json:"property"`
	Value    *PropertyValue         `json:"value"`
	Params   map[string]interface{} `json:"params"`
}

type PropertyExecutionResponse struct {
	Success bool           `json:"success"`
	Value   *PropertyValue `json:"value"`
}

type PropertyReference struct {
	Namespace []string `json:"namespace"`
	Key       string   `json:"key"`
	Version   string   `json:"version"`
}

type PropertyValue struct {
	ID      string      `json:"id"`
	Data    interface{} `json:"data"`
	Rule    *Rule       `json:"rule"`
	Segment *Segment    `json:"segment"`
}

type Segment struct {
	Path []*SegmentPathComponent `json:"path"`
}

type SegmentPathComponent struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

type Rule struct {
	Expression string `json:"expression"`
}

type Operation string

const (
	OperationCreatePropertyValue Operation = "CREATE_PROPERTY_VALUE"
	OperationReadPropertyValue   Operation = "READ_PROPERTY_VALUE"
	OperationUpdatePropertyValue Operation = "UPDATE_PROPERTY_VALUE"
	OperationDeletePropertyValue Operation = "DELETE_PROPERTY_VALUE"
)

type ErrorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
