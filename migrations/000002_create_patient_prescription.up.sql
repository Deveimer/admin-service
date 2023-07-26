CREATE TABLE IF NOT EXISTS "patient_prescription" (
                                                     "id" SERIAL NOT NULL,
                                                     "patient_id" varchar NOT NULL,
                                                     "doctor_id" varchar NOT NULL,
                                                     "prescription_location" varchar NOT NULL,
                                                     "notes"  varchar,
                                                     "createdAt" TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                                                     "updatedAt" TIMESTAMP WITHOUT TIME ZONE,
                                                     "deletedAt" TIMESTAMP WITHOUT TIME ZONE,
                                                     PRIMARY KEY ("id")
    )