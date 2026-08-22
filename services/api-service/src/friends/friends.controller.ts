import {
  Controller,
  Delete,
  Get,
  HttpCode,
  HttpStatus,
  Param,
  Patch,
  Post,
  UseGuards,
} from '@nestjs/common';
import {
  ApiBearerAuth,
  ApiOperation,
  ApiResponse,
  ApiTags,
} from '@nestjs/swagger';
import { FriendsService } from './friends.service';
import { FriendResponseDto } from './dto/friend-response.dto';
import { JwtAuthGuard } from 'src/auth/guards/jwt-auth.guard';
import { CurrentUser } from 'src/common/decorators/current-user.decorator';
import { ResponseMessage } from 'src/common/decorators/response-message.decorator';

@ApiTags('Friends')
@ApiBearerAuth()
@UseGuards(JwtAuthGuard)
@Controller('friends')
export class FriendsController {
  constructor(private readonly friendsService: FriendsService) {}

  @Get()
  @HttpCode(HttpStatus.OK)
  @ResponseMessage('Lấy danh sách bạn bè thành công')
  @ApiOperation({ summary: 'Lấy danh sách bạn bè hiện tại' })
  @ApiResponse({ status: 200, type: [FriendResponseDto] })
  async getFriendsList(
    @CurrentUser('id') userId: string,
  ): Promise<FriendResponseDto[]> {
    const list = await this.friendsService.getFriendsList(userId);
    return list.map((f) => new FriendResponseDto(f));
  }

  @Get('requests')
  @HttpCode(HttpStatus.OK)
  @ResponseMessage('Lấy danh sách lời mời kết bạn thành công')
  @ApiOperation({ summary: 'Lấy danh sách lời mời kết bạn đang chờ xử lý' })
  @ApiResponse({ status: 200, type: [FriendResponseDto] })
  async getPendingRequests(
    @CurrentUser('id') userId: string,
  ): Promise<FriendResponseDto[]> {
    const list = await this.friendsService.getPendingRequests(userId);
    return list.map((f) => new FriendResponseDto(f));
  }

  @Post('request/:targetUserId')
  @HttpCode(HttpStatus.CREATED)
  @ResponseMessage('Gửi lời mời kết bạn thành công')
  @ApiOperation({ summary: 'Gửi lời mời kết bạn tới một người dùng' })
  @ApiResponse({ status: 201, type: FriendResponseDto })
  async sendFriendRequest(
    @CurrentUser('id') userId: string,
    @Param('targetUserId') targetUserId: string,
  ): Promise<FriendResponseDto> {
    const result = await this.friendsService.sendFriendRequest(
      userId,
      targetUserId,
    );
    return new FriendResponseDto(result);
  }

  @Patch('requests/:requestId/accept')
  @HttpCode(HttpStatus.OK)
  @ResponseMessage('Đã chấp nhận lời mời kết bạn')
  @ApiOperation({ summary: 'Chấp nhận lời mời kết bạn' })
  @ApiResponse({ status: 200, type: FriendResponseDto })
  async acceptFriendRequest(
    @CurrentUser('id') userId: string,
    @Param('requestId') requestId: string,
  ): Promise<FriendResponseDto> {
    const result = await this.friendsService.acceptFriendRequest(
      userId,
      requestId,
    );
    return new FriendResponseDto(result);
  }

  @Patch('requests/:requestId/decline')
  @HttpCode(HttpStatus.OK)
  @ResponseMessage('Đã từ chối lời mời kết bạn')
  @ApiOperation({ summary: 'Từ chối lời mời kết bạn' })
  @ApiResponse({ status: 200, type: FriendResponseDto })
  async declineFriendRequest(
    @CurrentUser('id') userId: string,
    @Param('requestId') requestId: string,
  ): Promise<FriendResponseDto> {
    const result = await this.friendsService.declineFriendRequest(
      userId,
      requestId,
    );
    return new FriendResponseDto(result);
  }

  @Post('block/:targetUserId')
  @HttpCode(HttpStatus.OK)
  @ResponseMessage('Đã chặn người dùng')
  @ApiOperation({ summary: 'Chặn người dùng' })
  @ApiResponse({ status: 200, type: FriendResponseDto })
  async blockUser(
    @CurrentUser('id') userId: string,
    @Param('targetUserId') targetUserId: string,
  ): Promise<FriendResponseDto> {
    const result = await this.friendsService.blockUser(userId, targetUserId);
    return new FriendResponseDto(result);
  }

  @Delete(':relationshipId')
  @HttpCode(HttpStatus.OK)
  @ResponseMessage('Hủy kết bạn hoặc hủy yêu cầu thành công')
  @ApiOperation({ summary: 'Hủy quan hệ bạn bè hoặc thu hồi yêu cầu kết bạn' })
  @ApiResponse({ status: 200, description: 'Xóa thành công' })
  async unfriend(
    @CurrentUser('id') userId: string,
    @Param('relationshipId') relationshipId: string,
  ): Promise<{ success: boolean }> {
    const success = await this.friendsService.unfriendOrCancel(
      userId,
      relationshipId,
    );
    return { success };
  }
}
