create table tenants
(
	id bigserial not null
		constraint tenants_pkey
			primary key,
	name text,
	description text,
	email text,
	root_account_id bigint,
	is_activated boolean,
	activate_later timestamp with time zone,
	created_at timestamp with time zone default now()
);

alter table tenants owner to postgres;