import { User } from './user.entity';

export enum MemberRole {
  ADMIN = 'admin',
  MEMBER = 'member',
}

export class ConversationMember {
  id: string;
  conversationId: string;
  userId: string;
  role: MemberRole;
  joinedAt: Date;

  // Optional loaded relation
  user?: User;

  constructor(partial: Partial<ConversationMember>) {
    Object.assign(this, partial);
  }
}
