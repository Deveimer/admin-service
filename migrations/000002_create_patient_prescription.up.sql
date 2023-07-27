CREATE TABLE IF NOT EXISTS "patient_prescription" (
                                                     "id" SERIAL NOT NULL,
                                                     "patient_id" varchar NOT NULL,
                                                     "doctor_id" varchar NOT NULL,
                                                     "prescription_location" varchar NOT NULL,
                                                     "notes"  varchar,
                                                     "created_at" TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                                                     "updated_at" TIMESTAMP WITHOUT TIME ZONE,
                                                     "deleted_at" TIMESTAMP WITHOUT TIME ZONE DEFAULT NULL,
                                                     PRIMARY KEY ("id")
    )