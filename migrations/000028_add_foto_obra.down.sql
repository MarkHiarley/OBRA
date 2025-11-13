-- Remove foto column from obra table
ALTER TABLE obra 
DROP COLUMN IF EXISTS foto;
