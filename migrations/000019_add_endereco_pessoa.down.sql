-- 000019_add_endereco_pessoa.down.sql
-- Reverte a adição dos campos de endereço na tabela pessoa removendo colunas (idempotente)
DO $$
BEGIN
  IF EXISTS (
    SELECT 1 FROM information_schema.columns WHERE table_name='pessoa' AND column_name='endereco_rua'
  ) THEN
    ALTER TABLE pessoa DROP COLUMN IF EXISTS endereco_rua;
  END IF;

  IF EXISTS (
    SELECT 1 FROM information_schema.columns WHERE table_name='pessoa' AND column_name='endereco_numero'
  ) THEN
    ALTER TABLE pessoa DROP COLUMN IF EXISTS endereco_numero;
  END IF;

  IF EXISTS (
    SELECT 1 FROM information_schema.columns WHERE table_name='pessoa' AND column_name='endereco_complemento'
  ) THEN
    ALTER TABLE pessoa DROP COLUMN IF EXISTS endereco_complemento;
  END IF;

  IF EXISTS (
    SELECT 1 FROM information_schema.columns WHERE table_name='pessoa' AND column_name='endereco_bairro'
  ) THEN
    ALTER TABLE pessoa DROP COLUMN IF EXISTS endereco_bairro;
  END IF;

  IF EXISTS (
    SELECT 1 FROM information_schema.columns WHERE table_name='pessoa' AND column_name='endereco_cidade'
  ) THEN
    ALTER TABLE pessoa DROP COLUMN IF EXISTS endereco_cidade;
  END IF;

  IF EXISTS (
    SELECT 1 FROM information_schema.columns WHERE table_name='pessoa' AND column_name='endereco_estado'
  ) THEN
    ALTER TABLE pessoa DROP COLUMN IF EXISTS endereco_estado;
  END IF;

  IF EXISTS (
    SELECT 1 FROM information_schema.columns WHERE table_name='pessoa' AND column_name='endereco_cep'
  ) THEN
    ALTER TABLE pessoa DROP COLUMN IF EXISTS endereco_cep;
  END IF;
END$$;
