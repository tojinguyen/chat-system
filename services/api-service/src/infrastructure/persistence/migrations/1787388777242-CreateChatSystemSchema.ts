import { MigrationInterface, QueryRunner } from 'typeorm';

export class CreateChatSystemSchema1787388777242 implements MigrationInterface {
  name = 'CreateChatSystemSchema1787388777242';

  public async up(queryRunner: QueryRunner): Promise<void> {
    // 1. Thêm deletedAt vào users nếu chưa có
    await queryRunner.query(
      `ALTER TABLE "users" ADD COLUMN IF NOT EXISTS "deletedAt" TIMESTAMP WITH TIME ZONE`,
    );

    // 2. Tạo Enum & Bảng Friends
    await queryRunner.query(
      `CREATE TYPE "friend_status_enum" AS ENUM('pending', 'accepted', 'declined', 'blocked')`,
    );
    await queryRunner.query(`
      CREATE TABLE "friends" (
        "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
        "requesterId" uuid NOT NULL,
        "addresseeId" uuid NOT NULL,
        "status" "friend_status_enum" NOT NULL DEFAULT 'pending',
        "createdAt" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
        "updatedAt" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
        CONSTRAINT "UQ_friends_requester_addressee" UNIQUE ("requesterId", "addresseeId"),
        CONSTRAINT "PK_friends_id" PRIMARY KEY ("id"),
        CONSTRAINT "FK_friends_requester" FOREIGN KEY ("requesterId") REFERENCES "users"("id") ON DELETE CASCADE,
        CONSTRAINT "FK_friends_addressee" FOREIGN KEY ("addresseeId") REFERENCES "users"("id") ON DELETE CASCADE
      )
    `);
    await queryRunner.query(
      `CREATE INDEX "IDX_friends_requester" ON "friends" ("requesterId")`,
    );
    await queryRunner.query(
      `CREATE INDEX "IDX_friends_addressee" ON "friends" ("addresseeId")`,
    );

    // 3. Tạo Enum & Bảng Conversations
    await queryRunner.query(
      `CREATE TYPE "conversation_type_enum" AS ENUM('direct', 'group', 'clan', 'region', 'global')`,
    );
    await queryRunner.query(`
      CREATE TABLE "conversations" (
        "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
        "name" character varying(100),
        "type" "conversation_type_enum" NOT NULL DEFAULT 'direct',
        "iconUrl" character varying,
        "createdBy" uuid,
        "createdAt" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
        "updatedAt" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
        CONSTRAINT "PK_conversations_id" PRIMARY KEY ("id"),
        CONSTRAINT "FK_conversations_creator" FOREIGN KEY ("createdBy") REFERENCES "users"("id") ON DELETE SET NULL
      )
    `);

    // 4. Tạo Enum & Bảng Conversation Members
    await queryRunner.query(
      `CREATE TYPE "member_role_enum" AS ENUM('admin', 'member')`,
    );
    await queryRunner.query(`
      CREATE TABLE "conversation_members" (
        "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
        "conversationId" uuid NOT NULL,
        "userId" uuid NOT NULL,
        "role" "member_role_enum" NOT NULL DEFAULT 'member',
        "joinedAt" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
        CONSTRAINT "UQ_convo_members" UNIQUE ("conversationId", "userId"),
        CONSTRAINT "PK_conversation_members_id" PRIMARY KEY ("id"),
        CONSTRAINT "FK_members_conversation" FOREIGN KEY ("conversationId") REFERENCES "conversations"("id") ON DELETE CASCADE,
        CONSTRAINT "FK_members_user" FOREIGN KEY ("userId") REFERENCES "users"("id") ON DELETE CASCADE
      )
    `);
    await queryRunner.query(
      `CREATE INDEX "IDX_conversation_members_convo" ON "conversation_members" ("conversationId")`,
    );
    await queryRunner.query(
      `CREATE INDEX "IDX_conversation_members_user" ON "conversation_members" ("userId")`,
    );

    // 5. Bảng Conversation Read Receipts
    await queryRunner.query(`
      CREATE TABLE "conversation_read_receipts" (
        "conversationId" uuid NOT NULL,
        "userId" uuid NOT NULL,
        "lastMessageSeen" character varying(64) NOT NULL,
        "lastReadAt" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
        CONSTRAINT "PK_conversation_read_receipts" PRIMARY KEY ("conversationId", "userId"),
        CONSTRAINT "FK_receipts_conversation" FOREIGN KEY ("conversationId") REFERENCES "conversations"("id") ON DELETE CASCADE,
        CONSTRAINT "FK_receipts_user" FOREIGN KEY ("userId") REFERENCES "users"("id") ON DELETE CASCADE
      )
    `);
  }

  public async down(queryRunner: QueryRunner): Promise<void> {
    await queryRunner.query(`DROP TABLE "conversation_read_receipts"`);
    await queryRunner.query(`DROP TABLE "conversation_members"`);
    await queryRunner.query(`DROP TYPE "member_role_enum"`);
    await queryRunner.query(`DROP TABLE "conversations"`);
    await queryRunner.query(`DROP TYPE "conversation_type_enum"`);
    await queryRunner.query(`DROP TABLE "friends"`);
    await queryRunner.query(`DROP TYPE "friend_status_enum"`);
    await queryRunner.query(
      `ALTER TABLE "users" DROP COLUMN IF EXISTS "deletedAt"`,
    );
  }
}
