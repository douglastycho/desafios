import { Controller, Get } from '@nestjs/common';

import { AppService } from './app.service';
import { Route } from './types';

@Controller()
export class AppController {
  constructor(private readonly appService: AppService) {}

  @Get('routes')
  getHello(): Route[] {
    return this.appService.getRoutes();
  }
}
