-- users
-- auto generated users with hash in order to seed other values
INSERT INTO users (name, password_hash) VALUES
  ('alice', '$2a$10$Mnc38ppwrgoPYv8u9g6fNuZqoC2SFOA0yHZ.CpZp3P6vVD34ntmsq'), -- password: alice123
  ('bob', '$2a$10$h0/S2AmA0U01y7TNEnyBK.IIvWHizuKmirU/IIDklItPmJLUW0Xee'); -- password: bob123

-- topics
INSERT INTO topic (title, user_id) VALUES
  ('General', 1),
  ('Announcements', 2);

-- posts
INSERT INTO post (title, content, user_id, topic_id) VALUES
  ('Welcome', 'First post', 1, 1);

-- comments
INSERT INTO comment (content, user_id, post_id, parent_comment_id) VALUES
  ('Nice!', 2, 1, NULL);
