package jpush

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

const schedulePath = "https://api.jpush.cn/v3/schedules"

//ScheduleClient is
type ScheduleClient struct {
	AppKey       string
	MasterSecret string
}

//ScheduleCreateResponse is
type ScheduleCreateResponse struct {
	ScheduleID string `json:"schedule_id"`
	Name       string `json:"name"`
}

//SchedulePageResponse is
type SchedulePageResponse struct {
	TotalCount int                `json:"total_count"`
	TotalPages int                `json:"total_pages"`
	Page       int                `json:"page"`
	Schedules  []ScheduleResponse `json:"schedules"`
}

//CreateSchedule is
func (s *ScheduleClient) CreateSchedule(p *SchedulePayload) (*ScheduleCreateResponse, error) {
	body, err := json.Marshal(*p)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", schedulePath, bytes.NewReader(body))
	r, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var resp ScheduleCreateResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

//UpdateSchedule is
func (s *ScheduleClient) UpdateSchedule(p *ScheduleUpdateRequest, id string) (*SchedulePayload, error) {
	body, err := json.Marshal(*p)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PUT", schedulePath+"/"+id, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	r, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var resp SchedulePayload
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

//DeleteSchedule is
func (s *ScheduleClient) DeleteSchedule(id string) error {
	req, err := http.NewRequest("DELETE", schedulePath+"/"+id, nil)
	if err != nil {
		return err
	}
	_, err = s.doRequest(req)
	if err != nil {
		return err
	}
	return nil
}

//Schedules is
func (s *ScheduleClient) Schedules(page int) (*SchedulePageResponse, error) {
	req, err := http.NewRequest("GET", schedulePath+"?page="+strconv.Itoa(page), nil)
	if err != nil {
		return nil, err
	}
	r, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var resp SchedulePageResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

//Schedule is
func (s *ScheduleClient) Schedule(id string) (*ScheduleResponse, error) {
	req, err := http.NewRequest("GET", schedulePath+"/"+id, nil)
	if err != nil {
		return nil, err
	}
	r, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var resp ScheduleResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *ScheduleClient) doRequest(req *http.Request) ([]byte, error) {
	req.SetBasicAuth(s.AppKey, s.MasterSecret)
	req.Header.Add("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(string(result))
	}
	return result, nil
}

//NewScheduleClient is
func NewScheduleClient(appkey, secret string) *ScheduleClient {
	s := new(ScheduleClient)
	s.AppKey = appkey
	s.MasterSecret = secret
	return s
}
