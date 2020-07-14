-- we don't know how to generate root <with-no-name> (class Root) :(
create table roots
(
	account_id bigserial not null
		constraint roots_pkey
			primary key
);

alter table roots owner to postgres;

create table pods
(
	id bigserial not null
		constraint pods_pkey
			primary key,
	name text,
	datastore_id bigint,
	cluster_id bigint,
	subscription_type_id bigint,
	status text,
	ip_address text,
	created_at timestamp with time zone default now()
);

alter table pods owner to postgres;

create table subscription_types
(
	id bigserial not null
		constraint subscription_types_pkey
			primary key,
	ref_type_id bigint,
	expiry_time timestamp with time zone,
	tenant_id bigint
);

alter table subscription_types owner to postgres;

create table ref_types
(
	id bigserial not null
		constraint ref_types_pkey
			primary key,
	type_name text,
	vapp_template_name text
);

alter table ref_types owner to postgres;

create table vcenters
(
	id bigserial not null
		constraint vcenters_pkey
			primary key,
	ip_address text,
	user_name text,
	password text
);

alter table vcenters owner to postgres;

create table clusters
(
	id bigserial not null
		constraint clusters_pkey
			primary key,
	name text,
	vcenter_id bigint
);

alter table clusters owner to postgres;

create table datastores
(
	id bigserial not null
		constraint datastores_pkey
			primary key,
	name text,
	vcenter_id bigint
);

alter table datastores owner to postgres;

create table training_contents
(
	id bigserial not null
		constraint training_contents_pkey
			primary key,
	name text,
	content bytea,
	ref_type_id bigint
);

alter table training_contents owner to postgres;

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

