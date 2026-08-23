import {
  Body,
  Controller,
  Get,
  HttpCode,
  HttpStatus,
  NotFoundException,
  Param,
  Patch,
  Query,
  UseGuards,
} from '@nestjs/common';
import {
  ApiBearerAuth,
  ApiOperation,
  ApiQuery,
  ApiResponse,
  ApiTags,
} from '@nestjs/swagger';
import { UserService } from './users.service';
import { UserResponseDto } from './dto/user-response.dto';
import { UpdateUserDto } from './dto/update-user.dto';
import { JwtAuthGuard } from 'src/auth/guards/jwt-auth.guard';
import { CurrentUser } from 'src/common/decorators/current-user.decorator';
import { User } from 'src/domain/entities/user.entity';
import { ResponseMessage } from 'src/common/decorators/response-message.decorator';

@ApiTags('Users')
@ApiBearerAuth()
@UseGuards(JwtAuthGuard)
@Controller('users')
export class UsersController {
  constructor(private readonly usersService: UserService) {}

  @Get('me')
  @HttpCode(HttpStatus.OK)
  @ResponseMessage('Lấy thông tin cá nhân thành công')
  @ApiOperation({ summary: 'Lấy thông tin tài khoản hiện tại' })
  @ApiResponse({ status: 200, type: UserResponseDto })
  getProfile(@CurrentUser() user: User): UserResponseDto {
    return new UserResponseDto(user);
  }

  @Patch('me')
  @HttpCode(HttpStatus.OK)
  @ResponseMessage('Cập nhật thông tin cá nhân thành công')
  @ApiOperation({ summary: 'Cập nhật thông tin tài khoản hiện tại' })
  @ApiResponse({ status: 200, type: UserResponseDto })
  async updateProfile(
    @CurrentUser('id') userId: string,
    @Body() dto: UpdateUserDto,
  ): Promise<UserResponseDto> {
    const updated = await this.usersService.updateProfile(userId, dto);
    return new UserResponseDto(updated);
  }

  @Get('search')
  @HttpCode(HttpStatus.OK)
  @ResponseMessage('Tìm kiếm người dùng thành công')
  @ApiOperation({ summary: 'Tìm kiếm người dùng theo username, tên hoặc SĐT' })
  @ApiQuery({ name: 'q', required: true, description: 'Từ khóa tìm kiếm' })
  @ApiResponse({ status: 200, type: [UserResponseDto] })
  async searchUsers(@Query('q') query: string): Promise<UserResponseDto[]> {
    const users = await this.usersService.searchUsers(query);
    return users.map((u) => new UserResponseDto(u));
  }

  @Get(':id')
  @HttpCode(HttpStatus.OK)
  @ResponseMessage('Lấy thông tin người dùng thành công')
  @ApiOperation({ summary: 'Lấy thông tin người dùng theo ID' })
  @ApiResponse({ status: 200, type: UserResponseDto })
  @ApiResponse({ status: 404, description: 'Không tìm thấy người dùng' })
  async getUserById(@Param('id') id: string): Promise<UserResponseDto> {
    const user = await this.usersService.findById(id);
    if (!user) {
      throw new NotFoundException('Không tìm thấy người dùng');
    }
    return new UserResponseDto(user);
  }
}
