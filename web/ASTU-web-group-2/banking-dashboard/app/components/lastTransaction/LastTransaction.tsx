"use client";
import { useGetAllTransactionQuery } from "@/lib/service/TransactionService";
import Image from "next/image";
import { useSession } from "next-auth/react";
import { useEffect } from "react";
import { defaultItems } from "./lastTransactionItems";
import RecentTransactionSkeleton from "../recent-transaction/RecentTransactionSkeleton";
import ErrorImage from "../Error/ErrorImage";

export type Item = {
  transactionId: string;
  type: string;
  description: string;
  date: string;
  amount: number;
  receiverUserName: string;
};

function formatDate(date: Date) {
  const options: Intl.DateTimeFormatOptions = {
    day: "2-digit",
    month: "short",
    year: "numeric",
  };
  return date.toLocaleDateString("en-GB", options);
}

const LastTransaction = () => {
  let access: string = "";
  const { data: session, status } = useSession();
  useEffect(() => {}, [status, session]);
  if (session) {
    access = session?.user?.accessToken;
  }

  const { data, isError, isLoading } = useGetAllTransactionQuery(access);

  if (isLoading) {
    return (
      <RecentTransactionSkeleton />
    );
  }
  if (isError){
    <ErrorImage />
  }

  let transactions: { content: Item[]; totalPages: number } =
    data?.data || defaultItems;

  let items = transactions.content
    .map((transaction) => ({
      ...transaction,
      date: new Date(transaction.date),
    }))
    .sort((a, b) => a.date.getTime() - b.date.getTime())
    .map((transaction) => ({
      ...transaction,
      date: formatDate(transaction.date),
    }));

  if (items.length > 3) {
    items = items.slice(0, 3);
  }

  return (
    // The width depends on the width of container
    <div className=" flex flex-col bg-white gap-5 p-5 rounded-3xl">
      {items.length === 0 && (
        <div className="flex justify-center">
          <img
            src="/assets/bankService/empty-image.png"
            className="w-[35%]  "
          />
        </div>
      )}
      {items.map((item, index) => (
        <div key={index} className="flex items-center justify-between ">
          <div className="flex w-[45px] sm:w-[55px] justify-center mr-5">
            {item.type === "Shopping" || item.type === "shopping" ? (
              <Image
                src={"/assets/lastTransaction/spot-sub.svg"}
                width={55}
                height={55}
                alt={`${item.receiverUserName}-image`}
              />
            ) : item.type === "Service" || item.type === "service" ? (
              <Image
                src={"/assets/lastTransaction/settings.svg"}
                width={55}
                height={55}
                alt={`${item.receiverUserName}-image`}
              />
            ) : (
              <Image
                src={"/assets/lastTransaction/user.svg"}
                width={55}
                height={55}
                alt={`${item.receiverUserName}-image`}
              />
            )}
          </div>

          <div className="w-[137px] sm:w-[117px] md:w-[156px] ">
            <p className="text-[14px] sm:text-[16px] text-[#333B69] font-medium">
              {item.receiverUserName
                ? item.receiverUserName
                : "Spotify Subscription"}
            </p>
            <span className="text-[12px] sm:text-[15px] text-[#718EBF]">
              {item.date}
            </span>
          </div>
          <div className="hidden  sm:w-[100px] sm:flex justify-start items-center text-[16px] text-[#718EBF]">
            {item.type}
          </div>
          <div className="hidden sm:w-[100px] sm:flex justify-start items-center text-[16px] text-[#718EBF]">
            {item.transactionId.slice(0, 4)}****
          </div>
          <div className="hidden  sm:w-[100px] sm:flex justify-start items-center text-[16px] text-[#718EBF]">
            {item.description}
          </div>
          {item.amount < 0 ? (
            <div className="w-[100px] flex justify-end items-center text-[16px] text-[#FE5C73]">
              ${item.amount}
            </div>
          ) : (
            <div className="w-[100px] flex justify-end items-center text-[16px] text-[#16DBAA]">
              ${item.amount}
            </div>
          )}
        </div>
      ))}
    </div>
  );
};

export default LastTransaction;
