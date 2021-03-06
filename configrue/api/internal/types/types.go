// Code generated by goctl. DO NOT EDIT.
package types

type Environment struct {
	ID          int64  `json:"id"`
	Keyword     string `json:"keyword"`
	Name        string `json:"name"`
	Drive       string `json:"drive"`
	Config      string `json:"config"`
	Prefix      string `json:"prefix"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
	Operator    string `json:"operator"`
	OperatorID  int64  `json:"operator_id"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type GetEnvConfigRequest struct {
	Env string `form:"env"`
}

type GetEnvConfigResponse struct {
	Type     string `json:"type"`
	Host     string `json:"host"`
	Prefix   string `json:"prefix"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type GetEnvironmentResponse struct {
	List []Environment `json:"list"`
}

type AddEnvironmentRequest struct {
	Keyword     string `json:"keyword"`
	Name        string `json:"name"`
	Drive       string `json:"drive"`
	Config      string `json:"config"`
	Prefix      string `json:"prefix"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
}

type UpdateEnvironmentRequest struct {
	ID          int64  `json:"id"`
	Keyword     string `json:"keyword,optional"`
	Name        string `json:"name,optional"`
	Drive       string `json:"drive,optional"`
	Config      string `json:"config,optional"`
	Prefix      string `json:"prefix,optional"`
	Description string `json:"description,optional"`
	Status      *bool  `json:"status,optional"`
}

type DeleteEnvironmentRequest struct {
	ID int64 `json:"id"`
}

type Service struct {
	ID          int64  `json:"id"`
	Keyword     string `json:"keyword"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Operator    string `json:"operator"`
	OperatorID  int64  `json:"operator_id"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type PageServiceRequest struct {
	Page    int64  `form:"page"`
	Count   int64  `form:"count"`
	Keyword string `form:"keyword,optional"`
}

type PageServiceResponse struct {
	List  []Service `json:"list"`
	Total int64     `json:"total"`
}

type AllServiceResponse struct {
	List []Service `json:"list"`
}

type AddServiceRequest struct {
	Keyword     string `json:"keyword"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateServiceRequest struct {
	ID          int64  `json:"id"`
	Keyword     string `json:"keyword,optional"`
	Name        string `json:"name,optional"`
	Description string `json:"description,optional"`
}

type DeleteServiceRequest struct {
	ID int64 `json:"id"`
}

type Configure struct {
	ID          int64  `json:"id"`
	ServiceId   string `json:"service_id"`
	Template    string `json:"template"`
	Version     string `json:"version"`
	IsUse       bool   `json:"is_use"`
	Description string `json:"description"`
	Operator    string `json:"operator"`
	OperatorID  int64  `json:"operator_id"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type GetConfigureRequest struct {
	ID int64 `form:"id"`
}

type GetParseConfigureRequest struct {
	ID  int64  `form:"id"`
	Env string `form:"env"`
}

type SyncConfigureRequest struct {
	ID  int64  `json:"id"`
	Env string `json:"env"`
}

type ListConfigureRequest struct {
	ServiceId string `form:"service_id"`
	Version   string `form:"version,optional"`
	IsUse     *bool  `form:"is_use,optional"`
}

type ListConfigureResponse struct {
	List []Configure `json:"list"`
}

type AddConfigureRequest struct {
	ServiceId   int64  `json:"service_id"`
	ServiceName string `json:"service_name"`
	Template    string `json:"template"`
	Description string `json:"description"`
}

type UpdateConfigureRequest struct {
	ID          int64  `json:"id"`
	ServiceId   int64  `json:"service_id"`
	ServiceName string `json:"service_name"`
	IsUse       bool   `json:"is_use"`
}

type DeleteConfigureRequest struct {
	ID int64 `json:"id"`
}

type ConfigureLog struct {
	ID          int64  `json:"id"`
	ServiceName string `json:"service_name"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Operator    string `json:"operator"`
	OperatorID  int64  `json:"operator_id"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type GetConfigureLogRequest struct {
	ServiceName string `form:"service_name,optional"`
	Page        int64  `form:"page"`
	Count       int64  `form:"count,range=[0:50]"`
}

type GetConfigureLogResponse struct {
	List  []ConfigureLog `json:"list"`
	Total int64          `json:"total"`
}

type ConfigureField struct {
	ID          int64  `json:"id"`
	Type        string `json:"type"`
	Field       string `json:"field"`
	Config      string `json:"config"`
	Description string `json:"description"`
	Operator    string `json:"operator"`
	OperatorID  int64  `json:"operator_id"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type PageConfigureFieldRequest struct {
	Page  int64  `form:"page"`
	Count int64  `form:"count,range=[0:50]"`
	Field string `form:"field,optional"`
	Type  string `form:"type,optional"`
}

type PageConfigureFieldResponse struct {
	Total int64            `json:"total"`
	List  []ConfigureField `json:"list"`
}

type AddConfigureFieldRequest struct {
	Type        string `json:"type"`
	Field       string `json:"field"`
	Config      string `json:"config,optional"`
	Description string `json:"description"`
}

type UpdateConfigureFieldRequest struct {
	ID          int64  `json:"id"`
	Type        string `json:"type"`
	Field       string `json:"field"`
	Config      string `json:"config"`
	Description string `json:"description"`
}

type DeleteConfigureFieldRequest struct {
	ID int64 `json:"id"`
}

type ServiceField struct {
	ID          int64  `json:"id"`
	ServiceId   int64  `json:"service_id"`
	Field       string `json:"field"`
	Config      string `json:"config"`
	Description string `json:"description"`
	Operator    string `json:"operator"`
	OperatorID  int64  `json:"operator_id"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type PageServiceFieldRequest struct {
	Page      int64  `form:"page"`
	Count     int64  `form:"count,range=[0:50]"`
	Field     string `form:"field,optional"`
	ServiceId string `form:"service_id,optional"`
}

type PageServiceFieldResponse struct {
	Total int64          `json:"total"`
	List  []ServiceField `json:"list"`
}

type AddServiceFieldRequest struct {
	Field       string `json:"field"`
	ServiceId   int64  `json:"service_id"`
	Config      string `json:"config,optional"`
	Description string `json:"description"`
}

type UpdateServiceFieldRequest struct {
	ID          int64  `json:"id"`
	ServiceId   int64  `json:"service_id"`
	Field       string `json:"field"`
	Config      string `json:"config"`
	Description string `json:"description"`
}

type DeleteServiceFieldRequest struct {
	ID int64 `json:"id"`
}
