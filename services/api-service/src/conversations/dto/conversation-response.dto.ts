import { ApiProperty, ApiPropertyOptional } from '@nestjs/swagger';
import {
  Conversation,
  ConversationType,
} from 'src/domain/entities/conversation.entity';
import { MemberRole } from 'src/domain/entities/conversation-member.entity';
import { UserResponseDto } from 'src/users/dto/user-response.dto';

export class MemberResponseDto {
  @ApiProperty({ example: 'a1b2c3d4-e5f6-7890-abcd-ef1234567890' })
  userId: string;

  @ApiProperty({ enum: MemberRole, example: MemberRole.MEMBER })
  role: MemberRole;

  @ApiProperty({ example: '2026-08-22T08:00:00.000Z' })
  joinedAt: Date;

  @ApiPropertyOptional({ type: () => UserResponseDto })
  user?: UserResponseDto;
}

export class ReadReceiptResponseDto {
  @ApiProperty({ example: 'a1b2c3d4-e5f6-7890-abcd-ef1234567890' })
  userId: string;

  @ApiProperty({ example: '018dc5e1-cb00-7adb-6f4f-8d556994111c' })
  lastMessageSeen: string;

  @ApiProperty({ example: '2026-08-22T08:00:00.000Z' })
  lastReadAt: Date;
}

export class ConversationResponseDto {
  @ApiProperty({ example: 'a1b2c3d4-e5f6-7890-abcd-ef1234567890' })
  id: string;

  @ApiPropertyOptional({ example: 'Backend Team Chat' })
  name?: string;

  @ApiProperty({ enum: ConversationType, example: ConversationType.DIRECT })
  type: ConversationType;

  @ApiPropertyOptional({ example: 'https://example.com/icon.png' })
  iconUrl?: string;

  @ApiPropertyOptional({ example: 'a1b2c3d4-e5f6-7890-abcd-ef1234567890' })
  createdBy?: string;

  @ApiProperty({ example: '2026-08-22T08:00:00.000Z' })
  createdAt: Date;

  @ApiProperty({ example: '2026-08-22T08:00:00.000Z' })
  updatedAt: Date;

  @ApiProperty({ type: () => [MemberResponseDto], required: false })
  members?: MemberResponseDto[];

  @ApiProperty({ type: () => [ReadReceiptResponseDto], required: false })
  readReceipts?: ReadReceiptResponseDto[];

  constructor(convo: Conversation) {
    this.id = convo.id;
    this.name = convo.name;
    this.type = convo.type;
    this.iconUrl = convo.iconUrl;
    this.createdBy = convo.createdBy;
    this.createdAt = convo.createdAt;
    this.updatedAt = convo.updatedAt;

    this.members = convo.members?.map((m) => ({
      userId: m.userId,
      role: m.role,
      joinedAt: m.joinedAt,
      user: m.user ? new UserResponseDto(m.user) : undefined,
    }));

    this.readReceipts = convo.readReceipts?.map((r) => ({
      userId: r.userId,
      lastMessageSeen: r.lastMessageSeen,
      lastReadAt: r.lastReadAt,
    }));
  }
}
