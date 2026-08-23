import { ConversationMember } from './conversation-member.entity';
import { ReadReceipt } from './read-receipt.entity';

export enum ConversationType {
  DIRECT = 'direct',
  GROUP = 'group',
  CLAN = 'clan',
  REGION = 'region',
  GLOBAL = 'global',
}

export class Conversation {
  id: string;
  name?: string;
  type: ConversationType;
  iconUrl?: string;
  createdBy?: string;
  createdAt: Date;
  updatedAt: Date;

  // Optional loaded relations
  members?: ConversationMember[];
  readReceipts?: ReadReceipt[];

  constructor(partial: Partial<Conversation>) {
    Object.assign(this, partial);
  }
}
