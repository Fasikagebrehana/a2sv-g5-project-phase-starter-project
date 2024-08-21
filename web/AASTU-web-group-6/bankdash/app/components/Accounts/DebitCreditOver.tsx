"use client";

import { TrendingUp } from "lucide-react";
import { Bar, BarChart, CartesianGrid, XAxis } from "recharts";
import { useState, useEffect } from "react";

import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import {
  ChartConfig,
  ChartContainer,
  ChartTooltip,
  ChartTooltipContent,
} from "@/components/ui/chart";
import { TransactionType } from "@/app/Redux/slices/TransactionSlice";
import { useAppSelector } from "@/app/Redux/store/store";

const chartConfig = {
  debit: {
    label: "Debit",
    color: "hsl(var(--chart-1))",
  },
  credit: {
    label: "Credit",
    color: "hsl(var(--chart-2))",
  },
} satisfies ChartConfig;
interface chartData {
  day: string;
  debit: number;
  credit: number;
}
function isDateInLast7Days(dateString: string): boolean {
  const currentDate = new Date();
  const sevenDaysAgo = new Date(currentDate);
  sevenDaysAgo.setDate(currentDate.getDate() - 7);
  sevenDaysAgo.setHours(0, 0, 0, 0);

  const transactionDate = new Date(dateString);
  return transactionDate >= sevenDaysAgo && transactionDate <= currentDate;
}
function getDayOfWeek(dateString: string): string {
  const date = new Date(dateString);
  const daysOfWeek = [
    "Sunday",
    "Monday",
    "Tuesday",
    "Wednesday",
    "Thursday",
    "Friday",
    "Saturday",
  ];
  return daysOfWeek[date.getDay()];
}
export function DebitCreditOver() {
  const [data, setData] = useState<chartData[]>([]);
  const [totalIncome, setTotalIncome] = useState(0);
  const [totalExpense, setTotalExpense] = useState(0);
  const expense: TransactionType[] = useAppSelector(
    (state) => state.transactions.expense
  );
  const income: TransactionType[] = useAppSelector(
    (state) => state.transactions.income
  );

  console.log(expense, "from expense");
  console.log(income, "from income");
  useEffect(() => {
    const getData = () => {
      try {
        const chartData: { [day: string]: { debit: number; credit: number } } =
          {
            Sunday: { debit: 0, credit: 0 },
            Monday: { debit: 0, credit: 0 },
            Tuesday: { debit: 0, credit: 0 },
            Wednesday: { debit: 0, credit: 0 },
            Thursday: { debit: 0, credit: 0 },
            Friday: { debit: 0, credit: 0 },
            Saturday: { debit: 0, credit: 0 },
          };
        let incomeSum = 0;
        let expenseSum = 0;
        income.forEach((transaction) => {
          if (isDateInLast7Days(transaction.date)) {
            incomeSum += transaction.amount;
            const dayOfWeek = getDayOfWeek(transaction.date);
            chartData[dayOfWeek].credit += transaction.amount;
          }
        });

        expense.forEach((transaction) => {
          if (isDateInLast7Days(transaction.date)) {
            expenseSum += transaction.amount;
            const dayOfWeek = getDayOfWeek(transaction.date);
            chartData[dayOfWeek].debit += transaction.amount;
          }
        });

        const formattedChartData = Object.keys(chartData).map((day) => ({
          day: day,
          debit: chartData[day].debit,
          credit: chartData[day].credit,
        }));
        const currentDayIndex = new Date().getDay();
        const rotatedChartData = [
          ...formattedChartData.slice(currentDayIndex + 1),
          ...formattedChartData.slice(0, currentDayIndex + 1),]
        setData(rotatedChartData);
        setTotalExpense(expenseSum);
        setTotalIncome(incomeSum);
      } catch (error) {
        // alert("Error Fetching data ");
      }
    };
    getData();
  }, []);
  return (
    <Card className="rounded-3xl shadow-lg dark:bg-[#232328]  ">
      <CardHeader>
        <div className="flex justify-end lg:justify-between ">
          <CardTitle className="hidden gap-2 lg:block lg:text-[12px] xl:text-base text-base font-normal font-inter text-[#718EBF] dark:text-gray-400">
            <span className="font-semibold text-black dark:text-gray-300">
              ${totalExpense}
            </span>{" "}
            Debited &{" "}
            <span className="font-semibold text-black dark:text-gray-300">
              ${totalIncome}
            </span>{" "}
            Credited in this Week
          </CardTitle>
          <div className="flex gap-5 ">
            <div className="flex items-center gap-2">
              <div className="border border-[#4C78FF] w-[15px] h-[15px] rounded-sm bg-[#4C78FF]"></div>
              <p className="font-inter font-normal text-base text-[#718EBF] dark:text-gray-300">
                Debit
              </p>
            </div>
            <div className="flex items-center gap-2">
              <div className="border border-[#FCAA0B] w-[15px] h-[15px] rounded-sm bg-[#FCAA0B]"></div>
              <p className="font-inter font-normal text-base text-[#718EBF] dark:text-gray-300">
                Credit
              </p>
            </div>
          </div>
        </div>
      </CardHeader>
      <CardContent>
        <ChartContainer config={chartConfig} className="lg:h-[350px] w-[100%]">
          <BarChart accessibilityLayer data={data}>
            <CartesianGrid vertical={false} className="h-[50%] lg:h-[70%]" />
            <XAxis
              dataKey="day"
              tickLine={false}
              tickMargin={10}
              axisLine={false}
              fontSize={10}
              tickFormatter={(value) => value.slice(0, 3)}
            />
            <ChartTooltip
              cursor={false}
              content={<ChartTooltipContent indicator="dashed" />}
            />
            <Bar dataKey="debit" fill="#1A16F3" radius={10} />
            <Bar dataKey="credit" fill="#FCAA0B" radius={10} />
          </BarChart>
        </ChartContainer>
      </CardContent>
    </Card>
  );
}
