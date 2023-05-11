ALTER TABLE barber DROP CONSTRAINT barber_name_id_idx;
CREATE index barber_name_id_idx on barber (name_id);