import {
  Body,
  Controller,
  Delete,
  Get,
  HttpCode,
  HttpStatus,
  Param,
  Post,
  UseGuards,
} from '@nestjs/common';
import {
  ApiBearerAuth,
  ApiOperation,
  ApiResponse,
  ApiTags,
} from '@nestjs/swagger';
import { ConversationsService } from './conversations.service';
import { CreateDirectConversationDto } from './dto/create-direct-conversation.dto';
import { CreateGroupConversationDto } from './dto/create-group-conversation.dto';
import { AddMemberDto } from './dto/add-member.dto';
import { UpdateReadReceiptDto } from './dto/update-read-receipt.dto';
import {
  ConversationResponseDto,
  MemberResponseDto,
  ReadReceiptResponseDto,
} from './dto/conversation-response.dto';
import { JwtAuthGuard } from 'src/auth/guards/jwt-auth.guard';
import { CurrentUser } from 'src/common/decorators/current-user.decorator';
import { ResponseMessage } from 'src/common/decorators/response-message.decorator';

@ApiTags('Conversations')
@ApiBearerAuth()
@UseGuards(JwtAuthGuard)
@Controller('conversations')
export class ConversationsController {
  constructor(private readonly conversationsService: ConversationsService) {}

  @Get()
  @HttpCode(HttpStatus.OK)
  @ResponseMessage('Lấy danh sách cuộc hội thoại thành công')
  @ApiOperation({ summary: 'Lấy danh sách các cuộc hội thoại của tôi' })
  @ApiResponse({ status: 200, type: [ConversationResponseDto] })
  async getMyConversations(
    @CurrentUser('id') userId: string,
  ): Promise<ConversationResponseDto[]> {
    const convos = await this.conversationsService.getUserConversations(userId);
    return convos.map((c) => new ConversationResponseDto(c));
  }

  @Post('direct')
  @HttpCode(HttpStatus.OK)
  @ResponseMessage('Khởi tạo cuộc hội thoại 1-1 thành công')
  @ApiOperation({ summary: 'Tạo hoặc lấy cuộc hội thoại 1-1 với người khác' })
  @ApiResponse({ status: 200, type: ConversationResponseDto })
  async getOrCreateDirect(
    @CurrentUser('id') userId: string,
    @Body() dto: CreateDirectConversationDto,
  ): Promise<ConversationResponseDto> {
    const convo = await this.conversationsService.getOrCreateDirectConversation(
      userId,
      dto.partnerId,
    );
    return new ConversationResponseDto(convo);
  }

  @Post('group')
  @HttpCode(HttpStatus.CREATED)
  @ResponseMessage('Tạo nhóm chat thành công')
  @ApiOperation({ summary: 'Tạo nhóm chat mới' })
  @ApiResponse({ status: 201, type: ConversationResponseDto })
  async createGroup(
    @CurrentUser('id') userId: string,
    @Body() dto: CreateGroupConversationDto,
  ): Promise<ConversationResponseDto> {
    const convo = await this.conversationsService.createGroupConversation(
      userId,
      dto,
    );
    return new ConversationResponseDto(convo);
  }

  @Get(':id')
  @HttpCode(HttpStatus.OK)
  @ResponseMessage('Lấy chi tiết cuộc hội thoại thành công')
  @ApiOperation({
    summary: 'Lấy chi tiết cuộc hội thoại và danh sách thành viên',
  })
  @ApiResponse({ status: 200, type: ConversationResponseDto })
  async getConversationById(
    @CurrentUser('id') userId: string,
    @Param('id') id: string,
  ): Promise<ConversationResponseDto> {
    const convo = await this.conversationsService.getConversationById(
      id,
      userId,
    );
    return new ConversationResponseDto(convo);
  }

  @Post(':id/members')
  @HttpCode(HttpStatus.CREATED)
  @ResponseMessage('Thêm thành viên vào nhóm thành công')
  @ApiOperation({ summary: 'Thêm thành viên vào nhóm chat' })
  @ApiResponse({ status: 201, type: MemberResponseDto })
  async addMember(
    @CurrentUser('id') userId: string,
    @Param('id') conversationId: string,
    @Body() dto: AddMemberDto,
  ): Promise<MemberResponseDto> {
    const member = await this.conversationsService.addMember(
      conversationId,
      userId,
      dto,
    );
    return {
      userId: member.userId,
      role: member.role,
      joinedAt: member.joinedAt,
      user: member.user ? { ...member.user, passwordHash: '' } : undefined,
    };
  }

  @Delete(':id/members/:targetUserId')
  @HttpCode(HttpStatus.OK)
  @ResponseMessage('Xóa thành viên hoặc rời nhóm thành công')
  @ApiOperation({ summary: 'Xóa thành viên khỏi nhóm hoặc tự rời nhóm' })
  @ApiResponse({ status: 200, description: 'Thao tác thành công' })
  async removeMember(
    @CurrentUser('id') userId: string,
    @Param('id') conversationId: string,
    @Param('targetUserId') targetUserId: string,
  ): Promise<{ success: boolean }> {
    return await this.conversationsService.removeMember(
      conversationId,
      userId,
      targetUserId,
    );
  }

  @Post(':id/read')
  @HttpCode(HttpStatus.OK)
  @ResponseMessage('Cập nhật trạng thái đã đọc thành công')
  @ApiOperation({
    summary: 'Cập nhật tin nhắn đã đọc gần nhất trong hội thoại',
  })
  @ApiResponse({ status: 200, type: ReadReceiptResponseDto })
  async updateReadReceipt(
    @CurrentUser('id') userId: string,
    @Param('id') conversationId: string,
    @Body() dto: UpdateReadReceiptDto,
  ): Promise<ReadReceiptResponseDto> {
    const receipt = await this.conversationsService.updateReadReceipt(
      conversationId,
      userId,
      dto.lastMessageSeen,
    );
    return {
      userId: receipt.userId,
      lastMessageSeen: receipt.lastMessageSeen,
      lastReadAt: receipt.lastReadAt,
    };
  }
}
