import { IsNotEmpty, IsString } from 'class-validator';

export class LoginDto {
  @IsString()
  @IsNotEmpty({ message: 'Username or Email is required' })
  userNameOrEmail: string;

  @IsString()
  @IsNotEmpty({ message: 'Password is required' })
  password: string;
}
