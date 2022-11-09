drop table public.catalog;

create table public.catalog (
                                id serial primary key,
                                name varchar(100),
                                created timestamp default now(),
                                updated timestamp default now()
);

insert into public.catalog (name) values ('potato');
insert into public.catalog (name) values ('tomato');
insert into public.catalog (name) values ('apple');

select * from public.catalog;

