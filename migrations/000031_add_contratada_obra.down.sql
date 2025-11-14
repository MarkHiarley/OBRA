-- Migration 000031: Remove campo contratada da tabela obra

ALTER TABLE obra DROP COLUMN IF EXISTS contratada;
