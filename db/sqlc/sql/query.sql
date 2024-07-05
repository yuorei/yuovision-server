-- name: GetUser :one
SELECT * FROM user WHERE id = ? LIMIT 1;

-- name: GetUserSubscribeChannelsID :many
SELECT u.id FROM user AS u JOIN subscription AS s ON u.id = s.channel_id WHERE s.user_id = ?;

-- name: CreatetUser :execresult
INSERT INTO user (id, name, profile_image_url) VALUES (?, ?, ?);

-- name: SubscribeChannel :execresult
INSERT INTO subscription (user_id, channel_id) VALUES (?, ?);

-- name: UnSubscribeChannel :execresult
DELETE FROM subscription WHERE user_id = ? AND channel_id = ?;

-- name: GetUserSubscriptionID :one
SELECT u.id FROM user AS u JOIN subscription AS s ON u.id = s.channel_id WHERE s.user_id = ? AND s.channel_id = ? LIMIT 1;

-- name: GetVideo :one
SELECT * FROM video WHERE id = ? LIMIT 1;

-- name: GetPublicAndNonAdultNonAdVideos :many
SELECT * FROM video WHERE is_private   = false AND is_adult = false AND is_ad = false;

-- name: GetPublicAndNonAdByUploaderID :many
SELECT * FROM video WHERE is_private   = false AND is_ad = false AND uploader_id = ?;

-- name: GetVideoComments :many
SELECT c.* , u.name  FROM comment c INNER JOIN user u ON c.user_id = u.id WHERE video_id = ?;

-- name: GetVideoLikes :many
SELECT * FROM like_dislike WHERE video_id = ? AND is_like = true;

-- name: GetVideoDislikes :many
SELECT * FROM like_dislike WHERE video_id = ? AND is_like = false;

-- name: GetAllVideosTags :many
SELECT
    v.id AS video_id,
    t.id AS tag_id,
    t.tag_name
FROM
    video v
    INNER JOIN video_tags vt ON v.id = vt.video_id
    INNER JOIN tag t ON vt.tag_id = t.id
WHERE
    v.is_adult = false
    AND v.is_ad = false
    AND v.is_private = false;

-- name: GetAllVideosTagsByUserID :many
SELECT
    v.id AS video_id,
    t.id AS tag_id,
    t.tag_name
FROM
    video v
    INNER JOIN video_tags vt ON v.id = vt.video_id
    INNER JOIN tag t ON vt.tag_id = t.id
WHERE
    v.uploader_id = ?
    AND v.is_adult = false
    AND v.is_ad = false
    AND v.is_private = false;

-- name: GetVideoTags :many
SELECT t.id, t.tag_name FROM tag AS t JOIN video_tags AS vt ON t.id = vt.tag_id WHERE vt.video_id = ?;

-- name: CreateVideo :execresult
INSERT INTO video (id, title, description, video_url, thumbnail_image_url, is_private,is_external_cutout , is_adult, is_ad, uploader_id, created_at,updated_at,watch_count) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: CreateVideoTags :execresult
INSERT INTO video_tags (video_id, tag_id) VALUES (?, ?);

-- name: CreateTags :execresult
INSERT INTO tag (tag_name) VALUES (?);

-- name: CreateComment :execresult
INSERT INTO comment (id, video_id, text, user_id, created_at,updated_at) VALUES (?, ?, ?, ?, ?, ?);
