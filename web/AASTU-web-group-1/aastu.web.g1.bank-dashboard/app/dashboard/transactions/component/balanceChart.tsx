"use client";

import { Card, CardContent } from "@/components/ui/card";
import {
  ChartConfig,
  ChartContainer,
  ChartTooltip,
  ChartTooltipContent,
} from "@/components/ui/chart";
import { Area, AreaChart, CartesianGrid, Legend, XAxis, YAxis } from "recharts";

import { BalanceData } from "@/types";
import { useEffect, useState } from "react";
import { formatMonth } from "./utils";

const chartConfig = {
  value: {
    label: "Value",
    color: "#2D60FF",
  },
} satisfies ChartConfig;

export function BalanceAreachart({
  balanceHistory,
}: {
  balanceHistory: BalanceData[];
}) {
  const [chartData, setChartData] = useState<
    { month: string; value: number }[]
  >([]);

  useEffect(() => {
    const newChartData: { month: string; value: number }[] = [];
    balanceHistory.map((balance: BalanceData) => {
      newChartData.push({
        month: formatMonth(balance.time),
        value: Math.round(balance.value),
      });
    });
    setChartData(newChartData);
    console.log(newChartData);
  }, [balanceHistory]);
  return (
    <Card>
      <CardContent className="py-2">
        <ChartContainer config={chartConfig} className="h-40 w-full ">
          <AreaChart
            data={chartData}
            margin={{
              top: 10,
              right: 30,
              left: 0,
              bottom: 0,
            }}
            barSize={20}
          >
            <defs>
              <linearGradient id="colorValue" x1="0" y1="0" x2="0" y2="1">
                <stop offset="5%" stopColor="#2D60FF" stopOpacity={0.8} />
                <stop offset="95%" stopColor="#2D60FF" stopOpacity={0.01} />
              </linearGradient>
            </defs>
            <CartesianGrid strokeDasharray="3 3" vertical={false} />
            <XAxis
              dataKey="month"
              tickLine={false}
              axisLine={false}
              tickMargin={8}
              tickFormatter={(value) => value.slice(0, 3)}
            />
            <YAxis tickLine={false} axisLine={false} tickMargin={8} />
            <ChartTooltip
              cursor={false}
              content={<ChartTooltipContent indicator="dot" />}
            />

            <Legend />
            <Area
              type="natural"
              dataKey="value"
              stroke="#2D60FF"
              fillOpacity={1}
              fill="url(#colorValue)"
            />
          </AreaChart>
        </ChartContainer>
      </CardContent>
    </Card>
  );
}
