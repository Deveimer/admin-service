CREATE TYPE IF NOT EXISTS status_type AS ENUM ('OPEN', 'CANCELLED', 'SCHEDULED');
CREATE TABLE IF NOT EXISTS "doctor_opd_schedule" (
                                                 "id" SERIAL NOT NULL,
                                                 "doctor_id" varchar NOT NULL,
                                                 "opd_status" status_type NOT NULL,
                                                 "opd_start_date" DATE NOT NULL,
                                                 "opd_end_date" DATE  NOT NULL,
                                                 "opd_start_time" TIME NOT NULL,
                                                 "opd_end_time" TIME NOT NULL,
                                                 "opd_cancel_reason" varchar NULL,
                                                 PRIMARY KEY ("id")
    )