package entity

type Config struct {
	SUPABASE_HOST     string
	SUPABASE_USER     string
	SUPABASE_PASSWORD string
	SUPABASE_DB       string
	SUPABASE_PORT     string

	R2_BUCKET_NAME      string
	R2_PUBLIC_URL       string
	R2_TOKEN            string
	R2_ACCOUNT_ID       string
	R2_ACCESSKEY_ID     string
	R2_ACCESSKEY_SECRET string
	R2_S3_API           string

	HASH_COST string

	ACCESS_TOKEN_SECRET  string
	REFRESH_TOKEN_SECRET string
}
