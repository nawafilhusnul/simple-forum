ALTER TABLE refresh_tokens ADD COLUMN issued_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP AFTER refresh_token;