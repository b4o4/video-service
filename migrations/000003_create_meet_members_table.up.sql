BEGIN;

CREATE TABLE meet_members (
   meet_id UUID NOT NULL,
   member_id UUID NOT NULL,

   CONSTRAINT fk_meet_members_meet_id
       FOREIGN KEY (meet_id)
           REFERENCES meets(id)
           ON DELETE RESTRICT
           ON UPDATE CASCADE,

   CONSTRAINT fk_meet_members_member_id
       FOREIGN KEY (member_id)
           REFERENCES users(id)
           ON DELETE RESTRICT
           ON UPDATE CASCADE
);

CREATE UNIQUE INDEX idx_meet_members_meet_id_member_id ON meet_members(meet_id, member_id);

COMMENT ON TABLE meet_members IS 'Участники встречи';
COMMIT;