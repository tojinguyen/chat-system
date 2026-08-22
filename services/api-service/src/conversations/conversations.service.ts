import {
  BadRequestException,
  ForbiddenException,
  Inject,
  Injectable,
  NotFoundException,
} from '@nestjs/common';
import {
  Conversation,
  ConversationType,
} from 'src/domain/entities/conversation.entity';
import {
  ConversationMember,
  MemberRole,
} from 'src/domain/entities/conversation-member.entity';
import { ReadReceipt } from 'src/domain/entities/read-receipt.entity';
import {
  CONVERSATION_REPOSITORY,
  type IConversationRepository,
} from 'src/domain/repositories/conversation.repository.interface';
import { UserService } from 'src/users/users.service';
import { CreateGroupConversationDto } from './dto/create-group-conversation.dto';
import { AddMemberDto } from './dto/add-member.dto';

@Injectable()
export class ConversationsService {
  constructor(
    @Inject(CONVERSATION_REPOSITORY)
    private readonly convoRepo: IConversationRepository,
    private readonly userService: UserService,
  ) {}

  async getOrCreateDirectConversation(
    userAId: string,
    userBId: string,
  ): Promise<Conversation> {
    if (userAId === userBId) {
      throw new BadRequestException(
        'Không thể tạo cuộc hội thoại trực tiếp với chính mình',
      );
    }

    const partner = await this.userService.findById(userBId);
    if (!partner) {
      throw new NotFoundException('Không tìm thấy người dùng đối phương');
    }

    const existing = await this.convoRepo.findDirectConversation(
      userAId,
      userBId,
    );
    if (existing) {
      return existing;
    }

    return await this.convoRepo.create(
      {
        type: ConversationType.DIRECT,
        createdBy: userAId,
      },
      [
        { userId: userAId, role: MemberRole.MEMBER },
        { userId: userBId, role: MemberRole.MEMBER },
      ],
    );
  }

  async createGroupConversation(
    creatorId: string,
    dto: CreateGroupConversationDto,
  ): Promise<Conversation> {
    const uniqueMemberIds = Array.from(
      new Set(dto.memberIds.filter((id) => id !== creatorId)),
    );

    if (uniqueMemberIds.length === 0) {
      throw new BadRequestException('Nhóm phải có ít nhất 1 thành viên khác');
    }

    // Validate members existence
    for (const memberId of uniqueMemberIds) {
      const user = await this.userService.findById(memberId);
      if (!user) {
        throw new NotFoundException(
          `Không tìm thấy người dùng ID: ${memberId}`,
        );
      }
    }

    const membersToInsert = [
      { userId: creatorId, role: MemberRole.ADMIN },
      ...uniqueMemberIds.map((userId) => ({
        userId,
        role: MemberRole.MEMBER,
      })),
    ];

    return await this.convoRepo.create(
      {
        name: dto.name,
        type: dto.type || ConversationType.GROUP,
        iconUrl: dto.iconUrl,
        createdBy: creatorId,
      },
      membersToInsert,
    );
  }

  async getUserConversations(userId: string): Promise<Conversation[]> {
    return await this.convoRepo.findUserConversations(userId);
  }

  async getConversationById(
    conversationId: string,
    userId: string,
  ): Promise<Conversation> {
    const convo = await this.convoRepo.findById(conversationId);
    if (!convo) {
      throw new NotFoundException('Không tìm thấy cuộc hội thoại');
    }

    const isMember = convo.members?.some((m) => m.userId === userId);
    if (!isMember) {
      throw new ForbiddenException(
        'Bạn không phải là thành viên của cuộc hội thoại này',
      );
    }

    return convo;
  }

  async addMember(
    conversationId: string,
    requesterId: string,
    dto: AddMemberDto,
  ): Promise<ConversationMember> {
    const convo = await this.getConversationById(conversationId, requesterId);

    if (convo.type === ConversationType.DIRECT) {
      throw new BadRequestException(
        'Không thể thêm thành viên vào cuộc hội thoại 1-1. Hãy tạo nhóm mới.',
      );
    }

    const targetUser = await this.userService.findById(dto.userId);
    if (!targetUser) {
      throw new NotFoundException('Người dùng cần thêm không tồn tại');
    }

    return await this.convoRepo.addMember(
      conversationId,
      dto.userId,
      dto.role || MemberRole.MEMBER,
    );
  }

  async removeMember(
    conversationId: string,
    requesterId: string,
    targetUserId: string,
  ): Promise<{ success: boolean }> {
    const convo = await this.getConversationById(conversationId, requesterId);

    const requesterMember = convo.members?.find(
      (m) => m.userId === requesterId,
    );
    const targetMember = convo.members?.find((m) => m.userId === targetUserId);

    if (!targetMember) {
      throw new NotFoundException(
        'Thành viên không ở trong cuộc hội thoại này',
      );
    }

    // Nếu tự rời nhóm
    if (requesterId === targetUserId) {
      const success = await this.convoRepo.removeMember(
        conversationId,
        targetUserId,
      );
      return { success };
    }

    // Nếu kick người khác -> Phải là ADMIN
    if (requesterMember?.role !== MemberRole.ADMIN) {
      throw new ForbiddenException(
        'Chỉ quản trị viên mới có quyền xóa thành viên khỏi nhóm',
      );
    }

    const success = await this.convoRepo.removeMember(
      conversationId,
      targetUserId,
    );
    return { success };
  }

  async updateReadReceipt(
    conversationId: string,
    userId: string,
    lastMessageSeen: string,
  ): Promise<ReadReceipt> {
    // Validate user belongs to conversation
    await this.getConversationById(conversationId, userId);

    return await this.convoRepo.upsertReadReceipt(
      conversationId,
      userId,
      lastMessageSeen,
    );
  }
}
