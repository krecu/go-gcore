package gcore

import (
	"context"
	"fmt"
	"net/http"
)

var (
	OriginGroupsURL = "/originGroups"
	OriginGroupURL  = "/originGroups/%d"
)

// OriginGroupsService handles communication with the origin group related methods
// of the G-Core CDN API.
type OriginGroupsService service

// Origin represents G-Core's origin.
type Origin struct {
	ID      int    `json:"id,omitempty"`
	Enabled bool   `json:"enabled,omitempty"`
	Backup  bool   `json:"backup"`
	Source  string `json:"source"`
}

// Origin represents G-Core's origin group.
type OriginGroup struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	UseNext   bool     `json:"useNext"`
	OriginIDs []Origin `json:"origin_ids,omitempty"`
	Origins   []Origin `json:"origins"`
}

// UpdateOriginGroupBody represents request body for origin group update.
type UpdateOriginGroupBody struct {
	Name    string   `json:"name"`
	UseNext bool     `json:"useNext"`
	Origins []Origin `json:"origins"`
}

// CreateOriginGroupBody represents request body for origin group create.
type CreateOriginGroupBody struct {
	Name    string   `json:"name"`
	UseNext bool     `json:"useNext"`
	Origins []Origin `json:"origins"`
}

// List method returns list information about Origins Groups and Origin Sources.
func (s *OriginGroupsService) List(ctx context.Context) ([]*OriginGroup, *http.Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, OriginGroupsURL, nil)
	if err != nil {
		return nil, nil, err
	}

	originGroups := make([]*OriginGroup, 0)

	resp, err := s.client.Do(req, &originGroups)
	if err != nil {
		return nil, resp, err
	}

	return originGroups, resp, nil
}

// Get method returns origin group info for given ID.
func (s *OriginGroupsService) Get(ctx context.Context, originGroupID int) (*OriginGroup, *http.Response, error) {
	req, err := s.client.NewRequest(ctx,
		http.MethodGet,
		fmt.Sprintf(OriginGroupURL, originGroupID), nil)
	if err != nil {
		return nil, nil, err
	}

	originGroup := &OriginGroup{}

	resp, err := s.client.Do(req, originGroup)
	if err != nil {
		return nil, resp, err
	}

	return originGroup, resp, nil
}

// Create method creates origin group.
func (s *OriginGroupsService) Create(ctx context.Context, body *CreateOriginGroupBody) (*OriginGroup, *http.Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPost, OriginGroupsURL, body)
	if err != nil {
		return nil, nil, err
	}

	originGroup := &OriginGroup{}

	resp, err := s.client.Do(req, originGroup)
	if err != nil {
		return nil, resp, err
	}

	return originGroup, resp, nil
}

// Update method updates origin group info.
func (s *OriginGroupsService) Update(ctx context.Context, originGroupID int, body *UpdateOriginGroupBody) (*OriginGroup, *http.Response, error) {
	req, err := s.client.NewRequest(ctx,
		http.MethodPut,
		fmt.Sprintf(OriginGroupURL, originGroupID), body)
	if err != nil {
		return nil, nil, err
	}

	originGroup := &OriginGroup{}

	resp, err := s.client.Do(req, originGroup)
	if err != nil {
		return nil, resp, err
	}

	return originGroup, resp, nil
}

// Delete method deletes origin group.
func (s *OriginGroupsService) Delete(ctx context.Context, originGroupID int) (*http.Response, error) {
	req, err := s.client.NewRequest(ctx,
		http.MethodDelete,
		fmt.Sprintf(OriginGroupURL, originGroupID), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
