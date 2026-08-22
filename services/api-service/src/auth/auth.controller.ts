import { Body, Controller, HttpCode, HttpStatus, Post } from '@nestjs/common';
import { AuthService } from './auth.service';
import { ResponseMessage } from 'src/common/decorators/response-message.decorator';
import { RegisterDto } from './dto/register.dto';
import { UserResponseDto } from 'src/users/dto/uset-response.dto';
import { LoginDto, LoginResponseDto } from './dto/login.dto';
import { ApiOperation, ApiResponse, ApiTags } from '@nestjs/swagger';

@ApiTags('Auth')
@Controller('auth')
export class AuthController {
  constructor(private readonly authService: AuthService) {}

  @Post('register')
  @HttpCode(HttpStatus.CREATED)
  @ResponseMessage('Register successfull')
  @ApiOperation({ summary: 'Đăng ký tài khoản mới' })
  @ApiResponse({ status: 201, description: 'Đăng ký thành công' })
  @ApiResponse({ status: 409, description: 'Username đã tồn tại' })
  async register(@Body() dto: RegisterDto): Promise<UserResponseDto> {
    const user = await this.authService.register(dto);
    return new UserResponseDto(user);
  }

  @Post('login')
  @HttpCode(HttpStatus.OK)
  @ResponseMessage('Login successfull')
  @ApiOperation({ summary: 'Đăng nhập lấy JWT Token' })
  @ApiResponse({ status: 200, description: 'Đăng nhập thành công' })
  @ApiResponse({ status: 401, description: 'Sai tài khoản hoặc mật khẩu' })
  async login(@Body() dto: LoginDto): Promise<LoginResponseDto> {
    return await this.authService.login(dto);
  }
}
