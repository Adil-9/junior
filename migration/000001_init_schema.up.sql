do
$$
begin
    execute 'ALTER DATABASE ' || current_database() || ' SET timezone = ''+06''';
end;
$$;


CREATE TABLE IF NOT EXISTS person (
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL,
	surname TEXT NOT NULL,
	patronymic TEXT NOT NULL,
	age INTEGER NOT NULL,
	gender TEXT NOT NULL,
	country TEXT NOT NULL
);