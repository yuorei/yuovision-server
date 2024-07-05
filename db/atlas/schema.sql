-- Create "video" table
CREATE TABLE `video` (
 `id` varchar(255) NOT NULL,
 `video_url` varchar(255) NOT NULL,
 `thumbnail_image_url` varchar(255) NOT NULL,
 `title` varchar(255) NOT NULL,
 `description` text NULL,
 `created_at` timestamp NOT NULL,
 `updated_at` timestamp NOT NULL,
 `is_private` bool NOT NULL,
 `is_adult` bool NOT NULL,
 `is_ad` bool NOT NULL,
 `uploader_id` varchar(255) NOT NULL,
 `watch_count` int NOT NULL,
 `is_external_cutout` bool NOT NULL,
 PRIMARY KEY (`id`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "user" table
CREATE TABLE `user` (
 `id` varchar(255) NOT NULL,
 `name` varchar(255) NOT NULL,
 `profile_image_url` varchar(255) NOT NULL,
 PRIMARY KEY (`id`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "comment" table
CREATE TABLE `comment` (
 `id` varchar(255) NOT NULL,
 `video_id` varchar(255) NOT NULL,
 `text` text NOT NULL,
 `created_at` timestamp NOT NULL,
 `updated_at` timestamp NOT NULL,
 `user_id` varchar(255) NULL,
 PRIMARY KEY (`id`),
 INDEX `user_id` (`user_id`),
 INDEX `video_id` (`video_id`),
 CONSTRAINT `comment_ibfk_1` FOREIGN KEY (`video_id`) REFERENCES `video` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION,
 CONSTRAINT `comment_ibfk_2` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "history" table
CREATE TABLE `history` (
 `id` varchar(255) NOT NULL,
 `user_id` varchar(255) NOT NULL,
 `video_id` varchar(255) NOT NULL,
 `watched_at` timestamp NOT NULL,
 PRIMARY KEY (`id`),
 INDEX `user_id` (`user_id`),
 INDEX `video_id` (`video_id`),
 CONSTRAINT `history_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION,
 CONSTRAINT `history_ibfk_2` FOREIGN KEY (`video_id`) REFERENCES `video` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "like_dislike" table
CREATE TABLE `like_dislike` (
 `id` varchar(255) NOT NULL,
 `user_id` varchar(255) NOT NULL,
 `video_id` varchar(255) NULL,
 `comment_id` varchar(255) NULL,
 `is_like` bool NOT NULL,
 `created_at` timestamp NOT NULL,
 PRIMARY KEY (`id`),
 INDEX `comment_id` (`comment_id`),
 INDEX `user_id` (`user_id`),
 INDEX `video_id` (`video_id`),
 CONSTRAINT `like_dislike_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION,
 CONSTRAINT `like_dislike_ibfk_2` FOREIGN KEY (`video_id`) REFERENCES `video` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION,
 CONSTRAINT `like_dislike_ibfk_3` FOREIGN KEY (`comment_id`) REFERENCES `comment` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "playlist" table
CREATE TABLE `playlist` (
 `id` varchar(255) NOT NULL,
 `name` varchar(255) NOT NULL,
 `description` text NULL,
 `user_id` varchar(255) NOT NULL,
 `created_at` timestamp NOT NULL,
 PRIMARY KEY (`id`),
 INDEX `user_id` (`user_id`),
 CONSTRAINT `playlist_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "playlist_videos" table
CREATE TABLE `playlist_videos` (
 `playlist_id` varchar(255) NOT NULL,
 `video_id` varchar(255) NOT NULL,
 PRIMARY KEY (`playlist_id`, `video_id`),
 INDEX `video_id` (`video_id`),
 CONSTRAINT `playlist_videos_ibfk_1` FOREIGN KEY (`playlist_id`) REFERENCES `playlist` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION,
 CONSTRAINT `playlist_videos_ibfk_2` FOREIGN KEY (`video_id`) REFERENCES `video` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "report" table
CREATE TABLE `report` (
 `id` varchar(255) NOT NULL,
 `user_id` varchar(255) NOT NULL,
 `video_id` varchar(255) NULL,
 `comment_id` varchar(255) NULL,
 `reason` text NOT NULL,
 `created_at` timestamp NOT NULL,
 PRIMARY KEY (`id`),
 INDEX `comment_id` (`comment_id`),
 INDEX `user_id` (`user_id`),
 INDEX `video_id` (`video_id`),
 CONSTRAINT `report_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION,
 CONSTRAINT `report_ibfk_2` FOREIGN KEY (`video_id`) REFERENCES `video` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION,
 CONSTRAINT `report_ibfk_3` FOREIGN KEY (`comment_id`) REFERENCES `comment` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "subscription" table
CREATE TABLE `subscription` (
 `user_id` varchar(255) NOT NULL,
 `channel_id` varchar(255) NOT NULL,
 PRIMARY KEY (`user_id`, `channel_id`),
 INDEX `channel_id` (`channel_id`),
 CONSTRAINT `subscription_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION,
 CONSTRAINT `subscription_ibfk_2` FOREIGN KEY (`channel_id`) REFERENCES `user` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "category" table
CREATE TABLE `category` (
 `id` varchar(255) NOT NULL,
 `name` varchar(255) NOT NULL,
 PRIMARY KEY (`id`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "video_category" table
CREATE TABLE `video_category` (
 `video_id` varchar(255) NOT NULL,
 `category_id` varchar(255) NOT NULL,
 PRIMARY KEY (`video_id`, `category_id`),
 INDEX `category_id` (`category_id`),
 CONSTRAINT `video_category_ibfk_1` FOREIGN KEY (`video_id`) REFERENCES `video` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION,
 CONSTRAINT `video_category_ibfk_2` FOREIGN KEY (`category_id`) REFERENCES `category` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "tag" table
CREATE TABLE `tag` (
 `id` int NOT NULL AUTO_INCREMENT,
 `tag_name` varchar(255) NOT NULL,
 PRIMARY KEY (`id`),
 UNIQUE INDEX `tag_name` (`tag_name`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci AUTO_INCREMENT 3;
-- Create "video_tags" table
CREATE TABLE `video_tags` (
 `video_id` varchar(255) NOT NULL,
 `tag_id` int NOT NULL,
 PRIMARY KEY (`video_id`, `tag_id`),
 INDEX `tag_id` (`tag_id`),
 CONSTRAINT `video_tags_ibfk_1` FOREIGN KEY (`video_id`) REFERENCES `video` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION,
 CONSTRAINT `video_tags_ibfk_2` FOREIGN KEY (`tag_id`) REFERENCES `tag` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
