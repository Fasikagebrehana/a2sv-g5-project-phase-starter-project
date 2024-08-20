'use client'
import Image from "next/image";
import { useEffect, useState } from "react";
import ImageComponent from "./components/ImageComponent";
import Reviving from "./components/QuickTransfer";
import { BalanceHistory } from "./components/BalanceHistory";
import { WeeklyActivity } from "./components/WeeklyActivity";
import { ExpenseStatistics } from "./components/ExpenseStatistics";
import RecentTransaction from "./components/RecentTransaction";
import CreditCard from "./components/CreditCard";
import { getSession } from "next-auth/react";
import { useRouter } from "next/navigation";
// import {RecentTransaction} from "@/components/RecentTransaction"

export default function Home() {
  const [session, setSession] = useState(false);
  const [loading, setLoading] = useState(true);
  const route = useRouter();
  useEffect(() => {
    const fetchSession = async () => {
      const sessionData = await getSession();
      if (sessionData?.user) {
        setSession(true);
      } else {
        route.push(`./api/auth/signin?callbackUrl=${encodeURIComponent('/accounts')}`);
      }
      setLoading(false); // Set loading to false after session check
    };

    fetchSession();
  }, [route]); // Add router as a dependency
  // getting the session ends here

  
  return (
    <div className="h-screen w-screen ">

    <div className="flex flex-col">
      {/* Mobile Version */}
      <div className="flex flex-col md:hidden">
      <div className="flex items-center justify-between">
              <h1 className="mx-4 my-4 font-bold text-[#343C6A] text-2xl">My Cards</h1>
              <h1 className="mx-4 my-4 font-bold text-[#343C6A] text-lg">See All</h1>
            </div>
        <div className="flex overflow-x-auto [&::-webkit-scrollbar]:hidden">

          <div className="flex-col">

            <div className="flex">
              <div className="min-w-max min-h-max">
                <CreditCard
                  balance="$5,756"
                  cardHolder="Eddy Cusuma"
                  validThru="12/22"
                  cardNumber="3778 **** **** 1234"
                  filterClass=""
                  bgColor="from-[#4C49ED] to-[#0A06F4]"
                  textColor="text-white"
                  iconBgColor="bg-opacity-10"
                  showIcon={true}
                />
              </div>
              <div className="min-w-max min-h-max [&::-webkit-scrollbar]:hidden">
                <CreditCard
                  balance="$5,756"
                  cardHolder="Eddy Cusuma"
                  validThru="12/22"
                  cardNumber="3778 **** **** 1234"
                  filterClass=""
                  bgColor="bg-white"
                  textColor="text-black"
                  iconBgColor="bg-black"
                  showIcon={true}
                />
              </div>
            </div>
          </div>
        </div>
        <RecentTransaction />
        <WeeklyActivity />
        <ExpenseStatistics />
        <Reviving />
        <BalanceHistory />
      </div>

      {/* Web Version */}
      <div className="hidden md:flex flex-col  px-6 py-4 bg-[#f5f7fa] h-[130vh]">
           {/* <div className="flex items-center justify-between">
              <h1 className="mx-4 my-4 font-bold text-[#343C6A] text-2xl">My Cards</h1>
              <h1 className="mx-4 my-4 font-bold text-[#343C6A] text-lg">See All</h1>
            </div> */}
        <div className="flex">
          <div className="flex flex-col w-1/2">
            {/* My Cards Section */}
            <div className="flex items-center justify-between">
              <h1 className="mx-4 my-4 font-bold text-[#343C6A] text-2xl">My Cards</h1>
              <h1 className="mx-4 my-4 font-bold text-[#343C6A] text-lg">See All</h1>
            </div>
            <div className="flex space-x-6 overflow-x-auto [&::-webkit-scrollbar]:hidden">
              <div className="flex-shrink-0">
                <CreditCard
                  balance="$5,756"
                  cardHolder="Eddy Cusuma"
                  validThru="12/22"
                  cardNumber="3778 **** **** 1234"
                  bgColor="from-[#4C49ED] to-[#0A06F4]"
                  textColor="text-white"
                  iconBgColor="bg-opacity-10"
                  showIcon={true}
                />
              </div>
              <div className="flex-shrink-0">
                <CreditCard
                  balance="$5,756"
                  cardHolder="Eddy Cusuma"
                  validThru="12/22"
                  cardNumber="3778 **** **** 1234"
                  bgColor="bg-white"
                  textColor="text-black"
                  iconBgColor="bg-black"
                  showIcon={true}
                />
              </div>
              <div className="flex-shrink-0">
                <CreditCard
                  balance="$5,756"
                  cardHolder="Eddy Cusuma"
                  validThru="12/22"
                  cardNumber="3778 **** **** 1234"
                  bgColor="from-[#4C49ED] to-[#0A06F4]"
                  textColor="text-white"
                  iconBgColor="bg-opacity-10"
                  showIcon={true}
                />
              </div>
              <div className="flex-shrink-0">
                <CreditCard
                  balance="$5,756"
                  cardHolder="Eddy Cusuma"
                  validThru="12/22"
                  cardNumber="3778 **** **** 1234"
                  bgColor="from-[#4C49ED] to-[#0A06F4]"
                  textColor="text-white"
                  iconBgColor="bg-opacity-10"
                  showIcon={true}
                />
              </div>
            </div>
          </div>
  
  <div className="flex flex-col justify-between w-1/2 flex-grow-0">
    <h1 className="mx-4 my-4 font-bold text-[#343C6A] text-2xl">Recent Transaction</h1>
    <RecentTransaction />
  </div>
</div>

        <div className="flex space-x-6">
            <div className=" w-1/2">
              <h1 className="flex mx-4 my-4 font-bold  text-[#343C6A] text-2xl"> Weekly Activity</h1>
              <WeeklyActivity />
            </div>
            <div className=" w-1/3">
              <h1 className="flex mx-4 my-4 font-bold  text-[#343C6A] text-2xl"> Expense Statistics</h1>
              <ExpenseStatistics  />
            </div>
        </div>
        <div className="flex justify-between space-x-6 w-full h-24">
          <div className=" w-1/3 ">
          <h1 className="flex mx-4 my-4 font-bold  text-[#343C6A] text-2xl">Quick Transfers</h1>
            <Reviving />
          </div>
          <div className="w-2/3 h-5" >
          <h1 className="flex mx-4 my-4 font-bold  text-[#343C6A] text-2xl">Balance History</h1>
          <BalanceHistory />
          </div>
        </div>
      </div>
    </div>
    </div>

  
  );

}