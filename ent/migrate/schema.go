// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// SubscriptionsColumns holds the columns for the "subscriptions" table.
	SubscriptionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "tenant_id", Type: field.TypeString},
		{Name: "status", Type: field.TypeString, Default: "published"},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeString, Nullable: true},
		{Name: "updated_by", Type: field.TypeString, Nullable: true},
		{Name: "lookup_key", Type: field.TypeString, Nullable: true},
		{Name: "customer_id", Type: field.TypeString},
		{Name: "plan_id", Type: field.TypeString},
		{Name: "subscription_status", Type: field.TypeString, Default: "active"},
		{Name: "currency", Type: field.TypeString},
		{Name: "billing_anchor", Type: field.TypeTime},
		{Name: "start_date", Type: field.TypeTime},
		{Name: "end_date", Type: field.TypeTime, Nullable: true},
		{Name: "current_period_start", Type: field.TypeTime},
		{Name: "current_period_end", Type: field.TypeTime},
		{Name: "cancelled_at", Type: field.TypeTime, Nullable: true},
		{Name: "cancel_at", Type: field.TypeTime, Nullable: true},
		{Name: "cancel_at_period_end", Type: field.TypeBool, Default: false},
		{Name: "trial_start", Type: field.TypeTime, Nullable: true},
		{Name: "trial_end", Type: field.TypeTime, Nullable: true},
		{Name: "invoice_cadence", Type: field.TypeString},
		{Name: "billing_cadence", Type: field.TypeString},
		{Name: "billing_period", Type: field.TypeString},
		{Name: "billing_period_count", Type: field.TypeInt, Default: 1},
		{Name: "version", Type: field.TypeInt, Default: 1},
	}
	// SubscriptionsTable holds the schema information for the "subscriptions" table.
	SubscriptionsTable = &schema.Table{
		Name:       "subscriptions",
		Columns:    SubscriptionsColumns,
		PrimaryKey: []*schema.Column{SubscriptionsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "subscription_tenant_id_customer_id_status",
				Unique:  false,
				Columns: []*schema.Column{SubscriptionsColumns[1], SubscriptionsColumns[8], SubscriptionsColumns[2]},
			},
			{
				Name:    "subscription_tenant_id_plan_id_status",
				Unique:  false,
				Columns: []*schema.Column{SubscriptionsColumns[1], SubscriptionsColumns[9], SubscriptionsColumns[2]},
			},
			{
				Name:    "subscription_tenant_id_subscription_status_status",
				Unique:  false,
				Columns: []*schema.Column{SubscriptionsColumns[1], SubscriptionsColumns[10], SubscriptionsColumns[2]},
			},
			{
				Name:    "subscription_tenant_id_current_period_end_subscription_status_status",
				Unique:  false,
				Columns: []*schema.Column{SubscriptionsColumns[1], SubscriptionsColumns[16], SubscriptionsColumns[10], SubscriptionsColumns[2]},
			},
			{
				Name:    "subscription_tenant_id_lookup_key",
				Unique:  true,
				Columns: []*schema.Column{SubscriptionsColumns[1], SubscriptionsColumns[7]},
			},
		},
	}
	// WalletsColumns holds the columns for the "wallets" table.
	WalletsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "tenant_id", Type: field.TypeString},
		{Name: "customer_id", Type: field.TypeString},
		{Name: "currency", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "metadata", Type: field.TypeJSON, Nullable: true, SchemaType: map[string]string{"postgres": "jsonb"}},
		{Name: "balance", Type: field.TypeOther, SchemaType: map[string]string{"postgres": "numeric(20,9)"}},
		{Name: "wallet_status", Type: field.TypeString, Default: "active"},
		{Name: "status", Type: field.TypeString, Default: "published"},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeString, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeString, Nullable: true},
	}
	// WalletsTable holds the schema information for the "wallets" table.
	WalletsTable = &schema.Table{
		Name:       "wallets",
		Columns:    WalletsColumns,
		PrimaryKey: []*schema.Column{WalletsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "wallet_tenant_id_customer_id_status",
				Unique:  false,
				Columns: []*schema.Column{WalletsColumns[1], WalletsColumns[2], WalletsColumns[8]},
			},
			{
				Name:    "wallet_tenant_id_status_wallet_status",
				Unique:  false,
				Columns: []*schema.Column{WalletsColumns[1], WalletsColumns[8], WalletsColumns[7]},
			},
		},
	}
	// WalletTransactionsColumns holds the columns for the "wallet_transactions" table.
	WalletTransactionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "tenant_id", Type: field.TypeString},
		{Name: "wallet_id", Type: field.TypeString},
		{Name: "type", Type: field.TypeString, Default: "credit"},
		{Name: "amount", Type: field.TypeOther, SchemaType: map[string]string{"postgres": "numeric(20,9)"}},
		{Name: "balance_before", Type: field.TypeOther, SchemaType: map[string]string{"postgres": "numeric(20,9)"}},
		{Name: "balance_after", Type: field.TypeOther, SchemaType: map[string]string{"postgres": "numeric(20,9)"}},
		{Name: "reference_type", Type: field.TypeString, Nullable: true},
		{Name: "reference_id", Type: field.TypeString, Nullable: true},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "metadata", Type: field.TypeJSON, Nullable: true, SchemaType: map[string]string{"postgres": "jsonb"}},
		{Name: "transaction_status", Type: field.TypeString, Default: "pending"},
		{Name: "status", Type: field.TypeString, Default: "published"},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeString, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeString, Nullable: true},
	}
	// WalletTransactionsTable holds the schema information for the "wallet_transactions" table.
	WalletTransactionsTable = &schema.Table{
		Name:       "wallet_transactions",
		Columns:    WalletTransactionsColumns,
		PrimaryKey: []*schema.Column{WalletTransactionsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "wallettransaction_tenant_id_wallet_id_status",
				Unique:  false,
				Columns: []*schema.Column{WalletTransactionsColumns[1], WalletTransactionsColumns[2], WalletTransactionsColumns[12]},
			},
			{
				Name:    "wallettransaction_tenant_id_reference_type_reference_id_status",
				Unique:  false,
				Columns: []*schema.Column{WalletTransactionsColumns[1], WalletTransactionsColumns[7], WalletTransactionsColumns[8], WalletTransactionsColumns[12]},
			},
			{
				Name:    "wallettransaction_created_at",
				Unique:  false,
				Columns: []*schema.Column{WalletTransactionsColumns[13]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		SubscriptionsTable,
		WalletsTable,
		WalletTransactionsTable,
	}
)

func init() {
}
