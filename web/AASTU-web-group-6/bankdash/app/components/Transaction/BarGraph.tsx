"use client";

import { useState } from "react";
import { Bar, BarChart, CartesianGrid, XAxis, Cell } from "recharts";

import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import {
  ChartConfig,
  ChartContainer,
  ChartTooltip,
  ChartTooltipContent,
} from "@/components/ui/chart";

const chartData = [
  { month: "January", desktop: 186 },
  { month: "February", desktop: 305 },
  { month: "March", desktop: 237 },
  { month: "April", desktop: 73 },
  { month: "May", desktop: 209 },
  { month: "June", desktop: 214 },
];

const chartConfig = {
  desktop: {
    label: "Desktop",
    color: "#EDF0F7",
  },
} satisfies ChartConfig;

const BarGraph = () => {
  const [activeIndex, setActiveIndex] = useState<number | null>(null);

  const handleMouseOver = (index: number) => {
    setActiveIndex(index);
  };

  const handleMouseOut = () => {
    setActiveIndex(null);
  };

  return (
    <div className="space-y-3">
      <Card className="rounded-[25px]">
        <CardContent>
          <ChartContainer config={chartConfig} className="h-[200px] w-full">
            <BarChart data={chartData}>
              <CartesianGrid vertical={false} />
              <XAxis
                dataKey="month"
                tickLine={false}
                tickMargin={10}
                axisLine={false}
                tickFormatter={(value) => value.slice(0, 3)}
              />
              <ChartTooltip
                cursor={false}
                content={<ChartTooltipContent hideLabel />}
              />
              <Bar dataKey="desktop" radius={8}>
                {chartData.map((entry, index) => (
                  <Cell
                    key={`cell-${index}`}
                    fill={
                      index === activeIndex
                        ? "#16DBCC" // Hover color
                        : "var(--color-desktop)" // Default color
                    }
                    onMouseOver={() => handleMouseOver(index)}
                    onMouseOut={handleMouseOut}
                  />
                ))}
              </Bar>
            </BarChart>
          </ChartContainer>
        </CardContent>
      </Card>
    </div>
  );
};

export default BarGraph;
