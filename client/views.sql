 SELECT board.title,
    board.client_id,
    board.site_id,
    board.date_created,
    board.date_modified,
    board.is_active,
    board.sort_position,
    board.ulid,
    board.color,
    ( SELECT count(task.task_id)::integer AS count
           FROM task
          WHERE task.board_ulid::text = board.ulid::text AND task.status <> 'Done'::text) AS task_count
   FROM board
  WHERE board.is_active = true;



   SELECT task.task_id,
    task.title,
    task.board_id,
    task.site_id,
    task.date_created,
    task.date_modified,
    task.notes,
    task.sort_position,
    task.status,
    task.agenda_date,
    task.percent,
    task.ulid,
    task.board_ulid,
        CASE
            WHEN task.status = 'Todo'::text THEN 0
            WHEN task.status = 'In Progress'::text THEN 1
            WHEN task.status = 'Needs Feedback'::text THEN 2
            WHEN task.status = 'Done'::text THEN 3
            WHEN task.status = 'On Hold'::text THEN 4
            ELSE NULL::integer
        END AS status_sort
   FROM task;


-- job_total
    SELECT count(job.job_id) AS count,
    job.site_id
   FROM job
  GROUP BY job.site_id;

-- project_invoice
 SELECT p.project_id,
    p.client_id,
    p.date_created,
    p.date_modified,
    p.deadline,
    p.name,
    p.rate,
    p.site_id,
    ( SELECT sum(mi.hours) AS sum
           FROM milestone_item mi
          WHERE mi.project_id = p.project_id) AS total_hours,
    1 AS invoice_percentage,
    1 AS to_invoice_this_time,
    phi.note AS to_note_this_time,
    phi.invoice_id
   FROM project p
     JOIN project_has_invoice phi ON p.project_id = phi.project_id;




-- title_time_autocomplete
      SELECT b.ulid,
    b.title,
    b.project_id,
    b.client_id,
    b.sort,
    b.category,
    b.site_id,
    b.date_modified
   FROM ( SELECT a.ulid,
            a.title,
            max(a.project_id) AS project_id,
            a.client_id,
            min(a.sort) AS sort,
            a.category,
            a.site_id,
            max(a.date_modified) AS date_modified
           FROM ( SELECT project.ulid,
                    project.name AS title,
                    project.project_id,
                    project.client_id,
                    1 AS sort,
                    'Projects'::text AS category,
                    project.site_id,
                    project.date_modified
                   FROM project
                  WHERE project.client_id > 0
                UNION ALL
                 SELECT board.ulid,
                    board.title,
                    0 AS project_id,
                    board.client_id,
                    4 AS sort,
                    'Boards'::text AS category,
                    board.site_id,
                    board.date_modified
                   FROM board
                  WHERE board.client_id > 0
                UNION ALL
                 SELECT client.ulid,
                    client.client_name AS title,
                    0 AS project_id,
                    client.client_id,
                    3 AS sort,
                    'Clients'::text AS category,
                    client.site_id,
                    client.date_modified AS max
                   FROM client) a
          GROUP BY a.ulid, a.title, a.client_id, a.category, a.site_id) b
  ORDER BY b.sort, b.date_modified;