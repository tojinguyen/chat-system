import {
  BadRequestException,
  ConflictException,
  ForbiddenException,
  Inject,
  Injectable,
  NotFoundException,
} from '@nestjs/common';
import { Friend, FriendStatus } from 'src/domain/entities/friend.entity';
import {
  FRIEND_REPOSITORY,
  type IFriendRepository,
} from 'src/domain/repositories/friend.repository.interface';
import { UserService } from 'src/users/users.service';

@Injectable()
export class FriendsService {
  constructor(
    @Inject(FRIEND_REPOSITORY)
    private readonly friendRepository: IFriendRepository,
    private readonly userService: UserService,
  ) {}

  async sendFriendRequest(
    requesterId: string,
    targetUserId: string,
  ): Promise<Friend> {
    if (requesterId === targetUserId) {
      throw new BadRequestException(
        'Không thể gửi lời mời kết bạn cho chính mình',
      );
    }

    const targetUser = await this.userService.findById(targetUserId);
    if (!targetUser) {
      throw new NotFoundException('Người dùng mục tiêu không tồn tại');
    }

    const existing = await this.friendRepository.findByUsers(
      requesterId,
      targetUserId,
    );

    if (existing) {
      if (existing.status === FriendStatus.ACCEPTED) {
        throw new ConflictException('Hai bạn đã là bạn bè của nhau');
      }
      if (existing.status === FriendStatus.BLOCKED) {
        throw new ForbiddenException('Không thể gửi yêu cầu kết bạn');
      }
      if (existing.status === FriendStatus.PENDING) {
        if (existing.requesterId === requesterId) {
          throw new ConflictException('Bạn đã gửi lời mời kết bạn trước đó');
        } else {
          // Người kia đã gửi trước -> Tự động accept luôn
          const accepted = await this.friendRepository.updateStatus(
            existing.id,
            FriendStatus.ACCEPTED,
          );
          return accepted!;
        }
      }
      // Nếu trạng thái là DECLINED -> cập nhật lại thành PENDING
      const renewed = await this.friendRepository.updateStatus(
        existing.id,
        FriendStatus.PENDING,
      );
      return renewed!;
    }

    return await this.friendRepository.create({
      requesterId,
      addresseeId: targetUserId,
      status: FriendStatus.PENDING,
    });
  }

  async acceptFriendRequest(
    userId: string,
    requestId: string,
  ): Promise<Friend> {
    const friendReq = await this.friendRepository.findById(requestId);
    if (!friendReq) {
      throw new NotFoundException('Không tìm thấy yêu cầu kết bạn');
    }

    if (friendReq.addresseeId !== userId) {
      throw new ForbiddenException(
        'Bạn không có quyền chấp nhận lời mời kết bạn này',
      );
    }

    if (friendReq.status === FriendStatus.ACCEPTED) {
      return friendReq;
    }

    const updated = await this.friendRepository.updateStatus(
      requestId,
      FriendStatus.ACCEPTED,
    );
    return updated!;
  }

  async declineFriendRequest(
    userId: string,
    requestId: string,
  ): Promise<Friend> {
    const friendReq = await this.friendRepository.findById(requestId);
    if (!friendReq) {
      throw new NotFoundException('Không tìm thấy yêu cầu kết bạn');
    }

    if (friendReq.addresseeId !== userId) {
      throw new ForbiddenException(
        'Bạn không có quyền từ chối lời mời kết bạn này',
      );
    }

    const updated = await this.friendRepository.updateStatus(
      requestId,
      FriendStatus.DECLINED,
    );
    return updated!;
  }

  async blockUser(userId: string, targetUserId: string): Promise<Friend> {
    if (userId === targetUserId) {
      throw new BadRequestException('Không thể tự chặn chính mình');
    }

    const targetUser = await this.userService.findById(targetUserId);
    if (!targetUser) {
      throw new NotFoundException('Người dùng mục tiêu không tồn tại');
    }

    const existing = await this.friendRepository.findByUsers(
      userId,
      targetUserId,
    );

    if (existing) {
      const updated = await this.friendRepository.updateStatus(
        existing.id,
        FriendStatus.BLOCKED,
      );
      return updated!;
    }

    return await this.friendRepository.create({
      requesterId: userId,
      addresseeId: targetUserId,
      status: FriendStatus.BLOCKED,
    });
  }

  async unfriendOrCancel(
    userId: string,
    relationshipId: string,
  ): Promise<boolean> {
    const friend = await this.friendRepository.findById(relationshipId);
    if (!friend) {
      throw new NotFoundException('Không tìm thấy quan hệ bạn bè');
    }

    if (friend.requesterId !== userId && friend.addresseeId !== userId) {
      throw new ForbiddenException('Bạn không có quyền thao tác trên mục này');
    }

    return await this.friendRepository.delete(relationshipId);
  }

  async getPendingRequests(userId: string): Promise<Friend[]> {
    return await this.friendRepository.findPendingRequests(userId);
  }

  async getFriendsList(userId: string): Promise<Friend[]> {
    return await this.friendRepository.findFriendsList(userId);
  }
}
