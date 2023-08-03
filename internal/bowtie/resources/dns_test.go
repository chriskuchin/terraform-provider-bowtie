package resources

import (
	"testing"

	"github.com/chriskuchin/terraform-provider-bowtie/internal/bowtie/client"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func Test_mergeServerDetails(t *testing.T) {
	type args struct {
		serverList    []types.String
		serverDetails []dnsServersResourceModel
	}
	tests := []struct {
		name string
		args args
		want []client.Server
	}{
		{
			name: "rename",
			args: args{
				serverList: []types.String{
					types.StringValue("1.1.1.1"),
				},
				serverDetails: []dnsServersResourceModel{
					{
						ID:    types.StringValue("3c95739e-ec9e-40ea-8dca-e03f224ebb6b"),
						Addr:  types.StringValue("8.8.8.8"),
						Order: types.Int64Value(0),
					},
				},
			},
			want: []client.Server{
				{
					ID:    "3c95739e-ec9e-40ea-8dca-e03f224ebb6b",
					Addr:  "1.1.1.1",
					Order: 0,
				},
			},
		},
		{
			name: "remove",
			args: args{
				serverList: []types.String{
					types.StringValue("1.1.1.1"),
				},
				serverDetails: []dnsServersResourceModel{
					{
						ID:    types.StringValue("3c95739e-ec9e-40ea-8dca-e03f224ebb6b"),
						Addr:  types.StringValue("8.8.8.8"),
						Order: types.Int64Value(0),
					},
					{
						ID:    types.StringValue("4963983a-6f79-488d-b449-95288ead80b1"),
						Addr:  types.StringValue("4.4.8.8"),
						Order: types.Int64Value(1),
					},
				},
			},
			want: []client.Server{
				{
					ID:    "3c95739e-ec9e-40ea-8dca-e03f224ebb6b",
					Addr:  "1.1.1.1",
					Order: 0,
				},
			},
		},
		{
			name: "same",
			args: args{
				serverList: []types.String{
					types.StringValue("8.8.8.8"),
					types.StringValue("4.4.8.8"),
				},
				serverDetails: []dnsServersResourceModel{
					{
						ID:    types.StringValue("3c95739e-ec9e-40ea-8dca-e03f224ebb6b"),
						Addr:  types.StringValue("8.8.8.8"),
						Order: types.Int64Value(0),
					},
					{
						ID:    types.StringValue("4963983a-6f79-488d-b449-95288ead80b1"),
						Addr:  types.StringValue("4.4.8.8"),
						Order: types.Int64Value(1),
					},
				},
			},
			want: []client.Server{
				{
					ID:    "3c95739e-ec9e-40ea-8dca-e03f224ebb6b",
					Addr:  "8.8.8.8",
					Order: 0,
				},
				{
					ID:    "4963983a-6f79-488d-b449-95288ead80b1",
					Addr:  "4.4.8.8",
					Order: 1,
				},
			},
		},
		{
			name: "add",
			args: args{
				serverList: []types.String{
					types.StringValue("8.8.8.8"),
					types.StringValue("4.4.8.8"),
					types.StringValue("1.1.1.1"),
				},
				serverDetails: []dnsServersResourceModel{
					{
						ID:    types.StringValue("3c95739e-ec9e-40ea-8dca-e03f224ebb6b"),
						Addr:  types.StringValue("8.8.8.8"),
						Order: types.Int64Value(0),
					},
					{
						ID:    types.StringValue("4963983a-6f79-488d-b449-95288ead80b1"),
						Addr:  types.StringValue("4.4.8.8"),
						Order: types.Int64Value(1),
					},
				},
			},
			want: []client.Server{
				{
					ID:    "3c95739e-ec9e-40ea-8dca-e03f224ebb6b",
					Addr:  "8.8.8.8",
					Order: 0,
				},
				{
					ID:    "4963983a-6f79-488d-b449-95288ead80b1",
					Addr:  "4.4.8.8",
					Order: 1,
				},
				{
					Addr:  "1.1.1.1",
					Order: 2,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeServerDetails(tt.args.serverList, tt.args.serverDetails)

			if len(got) != len(tt.want) {
				t.Errorf("len(mergeServerDetails()) = %d, want %d", len(got), len(tt.want))
			}

			for index, val := range got {
				want := tt.want[index]
				_, err := uuid.Parse(val.ID)
				if err != nil {
					t.Errorf("mergeServerDetails() = %+v, want non empty uuid", err)
				}

				if want.ID != "" && want.ID != val.ID {
					t.Errorf("mergeServerDetails = %s, want %s", val.ID, want.ID)
				}

				if val.Addr != want.Addr {
					t.Errorf("mergeServerDetails() = %s, want %s", val.Addr, want.Addr)
				}

				if val.Order != want.Order {
					t.Errorf("mergeServerDetails() = %d, want %d", val.Order, want.Order)
				}
			}
		})
	}
}

func Test_mergeExcludeDNSDetails(t *testing.T) {
	type args struct {
		excludeList    []types.String
		excludeDetails []dnsExcludeResourceModel
	}
	tests := []struct {
		name string
		args args
		want []client.DNSExclude
	}{
		{
			name: "rename",
			args: args{
				excludeList: []types.String{
					types.StringValue("net.example.com"),
				},
				excludeDetails: []dnsExcludeResourceModel{
					{
						ID:    types.StringValue("c1cf6ef8-3234-41c2-a404-fa5fd71c854b"),
						Name:  types.StringValue("wrong.example.com"),
						Order: types.Int64Value(0),
					},
				},
			},
			want: []client.DNSExclude{
				{
					ID:    "c1cf6ef8-3234-41c2-a404-fa5fd71c854b",
					Name:  "net.example.com",
					Order: 0,
				},
			},
		},
		{
			name: "add",
			args: args{
				excludeList: []types.String{
					types.StringValue("net.example.com"),
					types.StringValue("other.example.com"),
				},
				excludeDetails: []dnsExcludeResourceModel{
					{
						ID:    types.StringValue("c1cf6ef8-3234-41c2-a404-fa5fd71c854b"),
						Name:  types.StringValue("wrong.example.com"),
						Order: types.Int64Value(0),
					},
				},
			},
			want: []client.DNSExclude{
				{
					ID:    "c1cf6ef8-3234-41c2-a404-fa5fd71c854b",
					Name:  "net.example.com",
					Order: 0,
				},
				{
					Name:  "other.example.com",
					Order: 1,
				},
			},
		},
		{
			name: "remove",
			args: args{
				excludeList: []types.String{
					types.StringValue("net.example.com"),
				},
				excludeDetails: []dnsExcludeResourceModel{
					{
						ID:    types.StringValue("c1cf6ef8-3234-41c2-a404-fa5fd71c854b"),
						Name:  types.StringValue("wrong.example.com"),
						Order: types.Int64Value(0),
					},
					{
						ID:    types.StringValue("255e1b92-8e06-4f98-bd3c-548cb722e0b6"),
						Name:  types.StringValue("other.example.com"),
						Order: types.Int64Value(1),
					},
				},
			},
			want: []client.DNSExclude{
				{
					ID:    "c1cf6ef8-3234-41c2-a404-fa5fd71c854b",
					Name:  "net.example.com",
					Order: 0,
				},
			},
		},
		{
			name: "same",
			args: args{
				excludeList: []types.String{
					types.StringValue("net.example.com"),
					types.StringValue("other.example.com"),
				},
				excludeDetails: []dnsExcludeResourceModel{
					{
						ID:    types.StringValue("c1cf6ef8-3234-41c2-a404-fa5fd71c854b"),
						Name:  types.StringValue("wrong.example.com"),
						Order: types.Int64Value(0),
					},
					{
						ID:    types.StringValue("255e1b92-8e06-4f98-bd3c-548cb722e0b6"),
						Name:  types.StringValue("other.example.com"),
						Order: types.Int64Value(1),
					},
				},
			},
			want: []client.DNSExclude{
				{
					ID:    "c1cf6ef8-3234-41c2-a404-fa5fd71c854b",
					Name:  "net.example.com",
					Order: 0,
				},
				{
					ID:    "255e1b92-8e06-4f98-bd3c-548cb722e0b6",
					Name:  "other.example.com",
					Order: 1,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeExcludeDNSDetails(tt.args.excludeList, tt.args.excludeDetails)
			if len(got) != len(tt.want) {
				t.Errorf("len(mergeExcludeDNSDetails()) = %d, want %d", len(got), len(tt.want))
			}

			for index, val := range got {
				want := tt.want[index]
				_, err := uuid.Parse(val.ID)
				if err != nil {
					t.Errorf("mergeExcludeDNSDetails() = %+v, want non empty uuid", err)
				}

				if want.ID != "" && want.ID != val.ID {
					t.Errorf("mergeExcludeDNSDetails = %s, want %s", val.ID, want.ID)
				}

				if val.Name != want.Name {
					t.Errorf("mergeExcludeDNSDetails() = %s, want %s", val.Name, want.Name)
				}

				if val.Order != want.Order {
					t.Errorf("mergeExcludeDNSDetails() = %d, want %d", val.Order, want.Order)
				}
			}
		})
	}
}
