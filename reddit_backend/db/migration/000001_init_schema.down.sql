DROP TABLE IF EXISTS Viewer;
DROP TABLE IF EXISTS Form;
DROP TABLE IF EXISTS Post;
DROP TABLE IF EXISTS Account;
DROP TABLE IF EXISTS Session;
DROP TABLE IF EXISTS Verify_email CASCADE;
DROP TABLE IF EXISTS Reply;
DROP TABLE IF EXISTS Community;
DROP TABLE IF EXISTS Comment;

ALTER TABLE Account DROP COLUMN is_email_verified;