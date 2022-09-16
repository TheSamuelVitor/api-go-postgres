create table equipes (
	id_equipe bigserial primary key,
	nome_equipe varchar(100) not null,
  createdAt timestamp default now() not null,
  updatedAt timestamp default now() not null
);

create table membros (
	id_membro bigserial primary key not null,
	nome_membro varchar(100) not null,
	funcao varchar(50),
	id_equipe bigint,
  createdAt timestamp default now() not null,
  updatedAt timestamp default now() not null,
	foreign key(id_equipe) references equipes(id_equipe) on delete cascade
);

create table projetos (
	id_projeto bigserial primary key not null,
	nome_projeto varchar(100),
	descricao varchar(100),
	id_equipe bigint,
  createdAt timestamp default now() not null,
  updatedAt timestamp default now() not null,
	foreign key(id_equipe) references equipes(id_equipe) on delete cascade
);

create table tarefas (
	id_tarefa bigserial primary key not null,
	nome_tarefa varchar(50),
	descricao varchar(100),
	id_projeto bigint,
	id_membro bigint,
  createdAt timestamp default now() not null,
  updatedAt timestamp default now() not null,
	foreign key(id_projeto) references projetos(id_projeto) on delete cascade,
	foreign key(id_membro) references membros(id_membro) on delete cascade
)