ALTER TABLE pessoa
    DROP COLUMN IF EXISTS nome_fantasia,
    DROP COLUMN IF EXISTS eh_pessoa_juridica,
    DROP COLUMN IF EXISTS funcao,
    DROP COLUMN IF EXISTS endereco_rua,
    DROP COLUMN IF EXISTS endereco_numero,
    DROP COLUMN IF EXISTS endereco_complemento,
    DROP COLUMN IF EXISTS endereco_bairro,
    DROP COLUMN IF EXISTS endereco_cidade,
    DROP COLUMN IF EXISTS endereco_estado,
    DROP COLUMN IF EXISTS endereco_cep,
    DROP COLUMN IF EXISTS observacao;
