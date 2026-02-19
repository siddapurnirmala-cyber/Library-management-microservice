-- Migration to add role column to users table
ALTER TABLE users ADD COLUMN IF NOT EXISTS role VARCHAR(20) DEFAULT 'MEMBER';

-- Update existing users to have MEMBER role if they don't have one
UPDATE users SET role = 'MEMBER' WHERE role IS NULL;
