-- 000019_add_endereco_pessoa.up.sql
-- Adiciona campos de endereço na tabela pessoa (idempotente)
DO $$
BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM information_schema.columns WHERE table_name='pessoa' AND column_name='endereco_rua'
  ) THEN
    ALTER TABLE pessoa ADD COLUMN endereco_rua varchar;
  END IF;

  IF NOT EXISTS (
    SELECT 1 FROM information_schema.columns WHERE table_name='pessoa' AND column_name='endereco_numero'
  ) THEN
    ALTER TABLE pessoa ADD COLUMN endereco_numero varchar;
  END IF;

  IF NOT EXISTS (
    SELECT 1 FROM information_schema.columns WHERE table_name='pessoa' AND column_name='endereco_complemento'
  ) THEN
    ALTER TABLE pessoa ADD COLUMN endereco_complemento varchar;
  END IF;

  IF NOT EXISTS (
    SELECT 1 FROM information_schema.columns WHERE table_name='pessoa' AND column_name='endereco_bairro'
  ) THEN
    ALTER TABLE pessoa ADD COLUMN endereco_bairro varchar;
  END IF;

  IF NOT EXISTS (
    SELECT 1 FROM information_schema.columns WHERE table_name='pessoa' AND column_name='endereco_cidade'
  ) THEN
    ALTER TABLE pessoa ADD COLUMN endereco_cidade varchar;
  END IF;

  IF NOT EXISTS (
    SELECT 1 FROM information_schema.columns WHERE table_name='pessoa' AND column_name='endereco_estado'
  ) THEN
    ALTER TABLE pessoa ADD COLUMN endereco_estado varchar;
  END IF;

  IF NOT EXISTS (
    SELECT 1 FROM information_schema.columns WHERE table_name='pessoa' AND column_name='endereco_cep'
  ) THEN
    ALTER TABLE pessoa ADD COLUMN endereco_cep varchar;
  END IF;
END$$;
