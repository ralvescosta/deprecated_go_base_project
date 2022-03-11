CREATE TABLE feiras (
  id serial NOT NULL,
  long INT NOT NULL,
  lat INT NOT NULL,
  setcens VARCHAR NOT NULL,
  areap VARCHAR NOT NULL,
  coddist INT NOT NULL,
  distrito VARCHAR NOT NULL,
  codsubpref INT NOT NULL,
  subpref VARCHAR NOT NULL,
  regiao5 VARCHAR NOT NULL,
  regiao8 VARCHAR NOT NULL,
  nome_feira VARCHAR NOT NULL,
  registro VARCHAR NOT NULL,
  lagradouro VARCHAR NOT NULL,
  numero VARCHAR NOT NULL,
  bairro VARCHAR NOT NULL,
  referencia VARCHAR NOT NULL,
  criado_em timestamptz NOT NULL,
  atualizado_em timestamptz NOT NULL,
  deletado_em timestamptz NOT NULL,
  CONSTRAINT feiras_pkey PRIMARY KEY (id)
)