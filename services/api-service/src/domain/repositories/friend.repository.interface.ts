import { Friend, FriendStatus } from '../entities/friend.entity';

export const FRIEND_REPOSITORY = 'FRIEND_REPOSITORY';

export interface IFriendRepository {
  findById(id: string): Promise<Friend | null>;
  findByUsers(userA: string, userB: string): Promise<Friend | null>;
  findPendingRequests(userId: string): Promise<Friend[]>;
  findFriendsList(userId: string): Promise<Friend[]>;
  create(
    friend: Omit<Friend, 'id' | 'createdAt' | 'updatedAt'>,
  ): Promise<Friend>;
  updateStatus(id: string, status: FriendStatus): Promise<Friend | null>;
  delete(id: string): Promise<boolean>;
}
