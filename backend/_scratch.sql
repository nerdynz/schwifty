 SELECT b.task,
    b.linked_from_title,
    b.project_id,
    b.board_ulid,
    b.sort,
    b.site_id,
    b.date_modified
   FROM ( SELECT a.task,
            p.name AS linked_from_title,
            a.project_id,
            a.board_ulid,
            a.sort,
            a.site_id,
            a.date_modified
           FROM ( SELECT milestone.title AS task,
                    ''::text AS linked_from_title,
                    milestone.project_id,
                    0 AS board_id,
                    1 AS sort,
                    milestone.site_id,
                    milestone.date_modified
                   FROM milestone
                UNION ALL
                 SELECT milestone_item.title AS task,
                    ''::text AS linked_from_title,
                    milestone_item.project_id,
                    0 AS board_id,
                    1 AS sort,
                    milestone_item.site_id,
                    milestone_item.date_modified
                   FROM milestone_item) a
             JOIN project p ON a.project_id = p.project_id
        UNION ALL
         SELECT task.title AS task,
            board.title AS linked_from_title,
            0 AS project_id,
            task.ulid as board,
            2 AS sort,
            task.site_id,
            task.date_modified
           FROM task
             JOIN board ON task.board_id = board.board_id
        UNION ALL
         SELECT time_entry.task,
            time_entry.title AS linked_from_title,
            0 AS project_id,
            0 AS board_ulid,
            3 AS sort,
            time_entry.site_id,
            max(time_entry.date_modified) AS date_modified
           FROM time_entry
          GROUP BY time_entry.task, time_entry.title, time_entry.site_id) b
  WHERE b.task <> ''::text AND b.linked_from_title <> ''::text
  ORDER BY b.sort, b.date_modified;