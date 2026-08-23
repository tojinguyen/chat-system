import { ApiProperty } from '@nestjs/swagger';
import { Friend, FriendStatus } from 'src/domain/entities/friend.entity';
import { UserResponseDto } from 'src/users/dto/user-response.dto';

export class FriendResponseDto {
  @ApiProperty({ example: 'a1b2c3d4-e5f6-7890-abcd-ef1234567890' })
  id: string;

  @ApiProperty({ example: 'a1b2c3d4-e5f6-7890-abcd-ef1234567890' })
  requesterId: string;

  @ApiProperty({ example: 'b2c3d4e5-f6a7-8901-bcde-f12345678901' })
  addresseeId: string;

  @ApiProperty({ enum: FriendStatus, example: FriendStatus.PENDING })
  status: FriendStatus;

  @ApiProperty({ example: '2026-08-22T08:00:00.000Z' })
  createdAt: Date;

  @ApiProperty({ example: '2026-08-22T08:00:00.000Z' })
  updatedAt: Date;

  @ApiProperty({ type: () => UserResponseDto, required: false })
  requester?: UserResponseDto;

  @ApiProperty({ type: () => UserResponseDto, required: false })
  addressee?: UserResponseDto;

  constructor(friend: Friend) {
    this.id = friend.id;
    this.requesterId = friend.requesterId;
    this.addresseeId = friend.addresseeId;
    this.status = friend.status;
    this.createdAt = friend.createdAt;
    this.updatedAt = friend.updatedAt;
    this.requester = friend.requester
      ? new UserResponseDto(friend.requester)
      : undefined;
    this.addressee = friend.addressee
      ? new UserResponseDto(friend.addressee)
      : undefined;
  }
}
