import { ConfigService } from '@nestjs/config';
import {
  FraudDetectionResult,
  FraudSpecificationContext,
  IFraudSpecification,
} from './fraud-specification.interface';
import { PrismaService } from 'src/prisma/prisma.service';
import { FraudReason } from '@prisma/client';
import { Injectable } from '@nestjs/common';

@Injectable()
export class FrequentHighValueSpecification implements IFraudSpecification {
  constructor(
    private prisma: PrismaService,
    private configService: ConfigService,
  ) {}

  async detectFraud(
    context: FraudSpecificationContext,
  ): Promise<FraudDetectionResult> {
    const { account } = context;
    const suspiciousInvoicesCount = this.configService.getOrThrow<number>(
      'SUSPICIOUS_INVOICES_COUNT',
    );
    const suspiciousTimeframeHours = this.configService.getOrThrow<number>(
      'SUSPICIOUS_TIMEFRAME_HOURS',
    );

    const recentDate = new Date();
    recentDate.setHours(recentDate.getHours() - suspiciousTimeframeHours);

    const recentInvoices = await this.prisma.invoice.findMany({
      where: {
        accountId: account.id,
        createdAt: { gte: recentDate },
      },
    });

    if (recentInvoices.length >= suspiciousInvoicesCount) {
      await this.prisma.account.update({
        where: { id: account.id },
        data: { isSuspicious: true },
      });

      return {
        hasFraud: true,
        reason: FraudReason.FREQUENT_HIGH_VALUE,
        description: `${recentInvoices.length} high-value invoices in the last ${suspiciousTimeframeHours} hours`,
      };
    }

    return { hasFraud: false };
  }
}
