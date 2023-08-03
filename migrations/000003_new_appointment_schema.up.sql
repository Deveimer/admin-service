CREATE TYPE IF NOT EXISTS appointment_status AS ENUM ('CANCELLED', 'COMPLETED', 'PENDING', 'SCHEDULED');
CREATE TABLE IF NOT EXISTS "appointments" (
                                                     "id" SERIAL NOT NULL,
                                                     "doctor_id" varchar NOT NULL,
                                                     "patient_id" varchar NOT NULL,
                                                     "status" appointment_status NOT NULL,
                                                     "description" varchar NULL,
                                                     "start_time" TIME NOT NULL,
                                                     "end_time" TIME NOT NULL,
                                                     "appointment_date" DATE NOT NULL,
                                                     "created_by" varchar NOT NULL,
                                                     "reason" varchar NULL,
                                                     PRIMARY KEY ("id")
    )