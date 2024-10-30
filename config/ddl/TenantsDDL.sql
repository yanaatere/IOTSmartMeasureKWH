CREATE TABLE public.tenants (
                                tenant_id bigserial NOT NULL,
                                tenant_name varchar(255) NOT NULL,
                                address text NULL,
                                created_at timestamptz NULL,
                                updated_at timestamptz NULL,
                                CONSTRAINT tenants_pkey PRIMARY KEY (tenant_id)
);

-- Create the trigger function
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at := NOW(); -- Set updated_at to the current timestamp
RETURN NEW; -- Return the updated row
END;
$$ LANGUAGE plpgsql;

-- Create the trigger
CREATE TRIGGER update_updated_at
    BEFORE UPDATE ON public.tenants
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();
