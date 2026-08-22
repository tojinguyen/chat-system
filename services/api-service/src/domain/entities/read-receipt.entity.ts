import { User } from './user.entity';

export class ReadReceipt {
  conversationId: string;
  userId: string;
  lastMessageSeen: string;
  lastReadAt: Date;

  // Optional loaded relation
  user?: User;

  constructor(partial: Partial<ReadReceipt>) {
    Object.assign(this, partial);
  }
}
