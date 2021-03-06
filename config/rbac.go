package config

import "sort"

// CreateRbacConfig return *[]KeyValuePair
func CreateRbacConfig() *RbacConfig {

	cfg := &RbacConfig{
		{
			Key:   "rbac.all",
			Value: 1,
		},
		{
			Key:   "rbac.edit",
			Value: 2,
		},
		{
			Key:   "ticket.all",
			Value: 4,
		},
		{
			Key:   "ticket.edit",
			Value: 8,
		},
		{
			Key:   "ticketType.all",
			Value: 16,
		},
		{
			Key:   "ticketType.edit",
			Value: 32,
		},
	}

	return cfg
}

// RbacConfig struct
type RbacConfig []KeyValuePair

// KeyValuePair struct
type KeyValuePair struct {
	Key   string
	Value int64
}

// Len return len
func (o *RbacConfig) Len() int { return len(*o) }

// Swap swap i, j
func (o *RbacConfig) Swap(i, j int) { (*o)[i], (*o)[j] = (*o)[j], (*o)[i] }

// Less compare i, j
func (o *RbacConfig) Less(i, j int) bool { return (*o)[i].Key < (*o)[j].Key }

// Search uses binary search to find and return the smallest index Value
func (o *RbacConfig) Search(key string) int64 {

	i := sort.Search(o.Len(), func(i int) bool { return (*o)[i].Key >= key })

	if i < o.Len() && (*o)[i].Key == key {
		return (*o)[i].Value
	}

	return 0
}
