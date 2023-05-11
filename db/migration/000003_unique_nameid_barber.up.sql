DROP INDEX barber_name_id_idx;
CREATE unique index barber_name_id_idx on barber (name_id);