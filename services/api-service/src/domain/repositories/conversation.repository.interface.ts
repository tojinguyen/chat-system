import { Conversation } from '../entities/conversation.entity';
import {
  ConversationMember,
  MemberRole,
} from '../entities/conversation-member.entity';
import { ReadReceipt } from '../entities/read-receipt.entity';

export const CONVERSATION_REPOSITORY = 'CONVERSATION_REPOSITORY';

export interface IConversationRepository {
  findById(id: string): Promise<Conversation | null>;
  findDirectConversation(
    userA: string,
    userB: string,
  ): Promise<Conversation | null>;
  findUserConversations(userId: string): Promise<Conversation[]>;
  create(
    conversation: Omit<
      Conversation,
      'id' | 'createdAt' | 'updatedAt' | 'members' | 'readReceipts'
    >,
    memberUserIds: { userId: string; role: MemberRole }[],
  ): Promise<Conversation>;
  addMember(
    conversationId: string,
    userId: string,
    role: MemberRole,
  ): Promise<ConversationMember>;
  removeMember(conversationId: string, userId: string): Promise<boolean>;
  findMember(
    conversationId: string,
    userId: string,
  ): Promise<ConversationMember | null>;
  upsertReadReceipt(
    conversationId: string,
    userId: string,
    lastMessageSeen: string,
  ): Promise<ReadReceipt>;
  getReadReceipt(
    conversationId: string,
    userId: string,
  ): Promise<ReadReceipt | null>;
}
