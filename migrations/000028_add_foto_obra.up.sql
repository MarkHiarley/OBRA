-- Add foto column to obra table
ALTER TABLE obra 
ADD COLUMN IF NOT EXISTS foto TEXT;
