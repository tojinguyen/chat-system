import { User } from './user.entity';

export enum FriendStatus {
  PENDING = 'pending',
  ACCEPTED = 'accepted',
  DECLINED = 'declined',
  BLOCKED = 'blocked',
}

export class Friend {
  id: string;
  requesterId: string;
  addresseeId: string;
  status: FriendStatus;
  createdAt: Date;
  updatedAt: Date;

  // Optional loaded relations
  requester?: User;
  addressee?: User;

  constructor(partial: Partial<Friend>) {
    Object.assign(this, partial);
  }
}
