import React from "react";
import BarGraph from "@/app/components/Transaction/BarGraph";
import Recent from "@/app/components/Transaction/Recent";
import Pagination from "@/app/components/Transaction/Pagination";
import VisaCard from "@/app/components/Card/VisaCard";
import FetchTransaction from "@/app/Services/api/transactionApi";
import {
  TransactionResponse,
  TransactionType,
  ChartData,
} from "@/types/TransactionValue";
import TransactionService from "@/app/Services/api/transactionApi";

const Transaction = async () => {
  const accessToken = process.env.NAHOM_TOKEN as string;

  const transactionData = await TransactionService.getTransactions(accessToken);
  const balance = await TransactionService.balanceHistory(accessToken);
  // console.log(balance, "---");
  const convertToChartData = (data: TransactionType[]): ChartData[] => {
    const dayMap: { [key: string]: number } = {
      Mon: 0,
      Tue: 0,
      Wed: 0,
      Thur: 0,
      Fri: 0,
      Sat: 0,
      Sun: 0,
    };

    data.forEach((transaction) => {
      const day = new Date(transaction.date).toLocaleString("en-US", {
        weekday: "short",
      });
      const formattedDay =
        day.charAt(0).toUpperCase() + day.slice(1, 3).toLowerCase();

      if (dayMap[formattedDay] !== undefined) {
        dayMap[formattedDay] += transaction.amount;
      }
    });

    return Object.keys(dayMap).map((day) => ({
      day,
      amount: dayMap[day],
    }));
  };

  const chartData = convertToChartData(transactionData);

  return (
    <div className="space-y-6 bg-[#F5F7FA] px-4 sm:px-6 md:px-8 lg:px-10">
      <div className="grid grid-cols-1 lg:grid-cols-3 gap-8 w-full">
        <div className="lg:col-span-2 py-4 overflow-x-auto scrollbar-hide">
          <div className="flex justify-between mb-5">
            <p className="font-semibold text-xl sm:text-2xl text-[#343C6A]">
              My Cards
            </p>
            <p className="font-semibold text-base sm:text-lg text-[#343C6A] cursor-pointer">
              + Add Card
            </p>
          </div>
          <div className="overflow-x-auto scrollbar-hide">
            <div className="flex gap-8 min-w-[650px] min-h-[170px]">
              <VisaCard isBlack={false} isFade={false} isSimGray={false} />
              <VisaCard isBlack={true} isFade={false} isSimGray={false} />
            </div>
          </div>
        </div>
        <div className="w-full">
          <p className="font-semibold text-xl sm:text-2xl text-[#343C6A] mb-1 py-4">
            My Expense
          </p>
          <BarGraph chartData={chartData} />
        </div>
      </div>
      {/* <Recent data={transactionData} /> */}
      <Recent />
      <Pagination />
    </div>
  );
};

export default Transaction;
