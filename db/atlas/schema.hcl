table "category" {
  schema = schema.yuovision
  column "id" {
    null = false
    type = varchar(255)
  }
  column "name" {
    null = false
    type = varchar(255)
  }
  primary_key {
    columns = [column.id]
  }
}
table "comment" {
  schema = schema.yuovision
  column "id" {
    null = false
    type = varchar(255)
  }
  column "video_id" {
    null = false
    type = varchar(255)
  }
  column "text" {
    null = false
    type = text
  }
  column "created_at" {
    null = false
    type = timestamp
  }
  column "updated_at" {
    null = false
    type = timestamp
  }
  column "user_id" {
    null = true
    type = varchar(255)
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "comment_ibfk_1" {
    columns     = [column.video_id]
    ref_columns = [table.video.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  foreign_key "comment_ibfk_2" {
    columns     = [column.user_id]
    ref_columns = [table.user.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  index "user_id" {
    columns = [column.user_id]
  }
  index "video_id" {
    columns = [column.video_id]
  }
}
table "history" {
  schema = schema.yuovision
  column "id" {
    null = false
    type = varchar(255)
  }
  column "user_id" {
    null = false
    type = varchar(255)
  }
  column "video_id" {
    null = false
    type = varchar(255)
  }
  column "watched_at" {
    null = false
    type = timestamp
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "history_ibfk_1" {
    columns     = [column.user_id]
    ref_columns = [table.user.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  foreign_key "history_ibfk_2" {
    columns     = [column.video_id]
    ref_columns = [table.video.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  index "user_id" {
    columns = [column.user_id]
  }
  index "video_id" {
    columns = [column.video_id]
  }
}
table "like_dislike" {
  schema = schema.yuovision
  column "id" {
    null = false
    type = varchar(255)
  }
  column "user_id" {
    null = false
    type = varchar(255)
  }
  column "video_id" {
    null = true
    type = varchar(255)
  }
  column "comment_id" {
    null = true
    type = varchar(255)
  }
  column "is_like" {
    null = false
    type = bool
  }
  column "created_at" {
    null = false
    type = timestamp
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "like_dislike_ibfk_1" {
    columns     = [column.user_id]
    ref_columns = [table.user.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  foreign_key "like_dislike_ibfk_2" {
    columns     = [column.video_id]
    ref_columns = [table.video.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  foreign_key "like_dislike_ibfk_3" {
    columns     = [column.comment_id]
    ref_columns = [table.comment.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  index "comment_id" {
    columns = [column.comment_id]
  }
  index "user_id" {
    columns = [column.user_id]
  }
  index "video_id" {
    columns = [column.video_id]
  }
}
table "playlist" {
  schema = schema.yuovision
  column "id" {
    null = false
    type = varchar(255)
  }
  column "name" {
    null = false
    type = varchar(255)
  }
  column "description" {
    null = true
    type = text
  }
  column "user_id" {
    null = false
    type = varchar(255)
  }
  column "created_at" {
    null = false
    type = timestamp
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "playlist_ibfk_1" {
    columns     = [column.user_id]
    ref_columns = [table.user.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  index "user_id" {
    columns = [column.user_id]
  }
}
table "playlist_videos" {
  schema = schema.yuovision
  column "playlist_id" {
    null = false
    type = varchar(255)
  }
  column "video_id" {
    null = false
    type = varchar(255)
  }
  primary_key {
    columns = [column.playlist_id, column.video_id]
  }
  foreign_key "playlist_videos_ibfk_1" {
    columns     = [column.playlist_id]
    ref_columns = [table.playlist.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  foreign_key "playlist_videos_ibfk_2" {
    columns     = [column.video_id]
    ref_columns = [table.video.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  index "video_id" {
    columns = [column.video_id]
  }
}
table "report" {
  schema = schema.yuovision
  column "id" {
    null = false
    type = varchar(255)
  }
  column "user_id" {
    null = false
    type = varchar(255)
  }
  column "video_id" {
    null = true
    type = varchar(255)
  }
  column "comment_id" {
    null = true
    type = varchar(255)
  }
  column "reason" {
    null = false
    type = text
  }
  column "created_at" {
    null = false
    type = timestamp
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "report_ibfk_1" {
    columns     = [column.user_id]
    ref_columns = [table.user.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  foreign_key "report_ibfk_2" {
    columns     = [column.video_id]
    ref_columns = [table.video.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  foreign_key "report_ibfk_3" {
    columns     = [column.comment_id]
    ref_columns = [table.comment.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  index "comment_id" {
    columns = [column.comment_id]
  }
  index "user_id" {
    columns = [column.user_id]
  }
  index "video_id" {
    columns = [column.video_id]
  }
}
table "subscription" {
  schema = schema.yuovision
  column "user_id" {
    null = false
    type = varchar(255)
  }
  column "channel_id" {
    null = false
    type = varchar(255)
  }
  primary_key {
    columns = [column.user_id, column.channel_id]
  }
  foreign_key "subscription_ibfk_1" {
    columns     = [column.user_id]
    ref_columns = [table.user.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  foreign_key "subscription_ibfk_2" {
    columns     = [column.channel_id]
    ref_columns = [table.user.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  index "channel_id" {
    columns = [column.channel_id]
  }
}
table "tag" {
  schema = schema.yuovision
  column "id" {
    null           = false
    type           = int
    auto_increment = true
  }
  column "tag_name" {
    null = false
    type = varchar(255)
  }
  primary_key {
    columns = [column.id]
  }
  index "tag_name" {
    unique  = true
    columns = [column.tag_name]
  }
}
table "user" {
  schema = schema.yuovision
  column "id" {
    null = false
    type = varchar(255)
  }
  column "name" {
    null = false
    type = varchar(255)
  }
  column "profile_image_url" {
    null = false
    type = varchar(255)
  }
  primary_key {
    columns = [column.id]
  }
}
table "video" {
  schema = schema.yuovision
  column "id" {
    null = false
    type = varchar(255)
  }
  column "video_url" {
    null = false
    type = varchar(255)
  }
  column "thumbnail_image_url" {
    null = false
    type = varchar(255)
  }
  column "title" {
    null = false
    type = varchar(255)
  }
  column "description" {
    null = true
    type = text
  }
  column "created_at" {
    null = false
    type = timestamp
  }
  column "updated_at" {
    null = false
    type = timestamp
  }
  column "is_private" {
    null = false
    type = bool
  }
  column "is_adult" {
    null = false
    type = bool
  }
  column "is_ad" {
    null = false
    type = bool
  }
  column "uploader_id" {
    null = false
    type = varchar(255)
  }
  column "watch_count" {
    null = false
    type = int
  }
  column "is_external_cutout" {
    null = false
    type = bool
  }
  primary_key {
    columns = [column.id]
  }
}
table "video_category" {
  schema = schema.yuovision
  column "video_id" {
    null = false
    type = varchar(255)
  }
  column "category_id" {
    null = false
    type = varchar(255)
  }
  primary_key {
    columns = [column.video_id, column.category_id]
  }
  foreign_key "video_category_ibfk_1" {
    columns     = [column.video_id]
    ref_columns = [table.video.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  foreign_key "video_category_ibfk_2" {
    columns     = [column.category_id]
    ref_columns = [table.category.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  index "category_id" {
    columns = [column.category_id]
  }
}
table "video_tags" {
  schema = schema.yuovision
  column "video_id" {
    null = false
    type = varchar(255)
  }
  column "tag_id" {
    null = false
    type = int
  }
  primary_key {
    columns = [column.video_id, column.tag_id]
  }
  foreign_key "video_tags_ibfk_1" {
    columns     = [column.video_id]
    ref_columns = [table.video.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  foreign_key "video_tags_ibfk_2" {
    columns     = [column.tag_id]
    ref_columns = [table.tag.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  index "tag_id" {
    columns = [column.tag_id]
  }
}
schema "yuovision" {
  charset = "utf8mb4"
  collate = "utf8mb4_0900_ai_ci"
}
