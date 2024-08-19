"use client";
import React from "react";
import Image from "next/image";
import x from "../../../public/assets/next-icon.svg";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { library } from "@fortawesome/fontawesome-svg-core";
import { faPaperPlane } from "@fortawesome/free-regular-svg-icons";
import { AreaComp } from "../Charts/AreaComp";
import { useAppSelector } from "@/app/Redux/store/store";
import {
  faArrowRight,
  faGreaterThan,
  faLessThan,
} from "@fortawesome/free-solid-svg-icons";
import { BalanceType } from "@/app/Redux/slices/TransactionSlice";
const Bottom = () => {
  const people: any = [1, 2, 3];
  const BalanceData: BalanceType[] = useAppSelector(
    (state) => state.transactions.balanceHist
  );

  return (
    <section className="Botom flex gap-6 xs:flex-col lg:flex-row ">
      <div className="cards-container sm:w-full lg:w-[45%]  center-content flex flex-col gap-6">
        <h1 className="text-xl font-semibold text-colorBody-1 dark:text-gray-300">
          Expense Statistics
        </h1>
        <div className="flex gap-6 bg-white dark:bg-[#232328] rounded-3xl  p-6">
          <div className="profle-box w-full flex flex-col gap-4">
            <div className="w-full flex gap-2 items-center ">
              <div className="profile-item flex flex-col gap-1 p-6 items-center justify-center dark:text-gray-300">
                <Image
                  className=" rounded-full "
                  src={"/assets/profile-1.png"}
                  width={70}
                  height={70}
                  alt=""
                />
                <div className="name text-base font-semibold">LIvia Bator</div>
                <div className="role text-base text-[#718EBF]">CEO</div>
              </div>
              <div className="profile-item flex flex-col gap-1 p-6 items-center justify-center dark:text-gray-300">
                <Image
                  className=" rounded-full "
                  src={"/assets/profile-1.png"}
                  width={70}
                  height={70}
                  alt=""
                />
                <div className="name text-base font-semibold">LIvia Bator</div>
                <div className="role text-base text-[#718EBF]">CEO</div>
              </div>
              <div className="profile-item flex flex-col gap-1 p-6 items-center justify-center dark:text-gray-300">
                <Image
                  className=" rounded-full "
                  src={"/assets/profile-1.png"}
                  width={70}
                  height={70}
                  alt=""
                />
                <div className="name text-base font-semibold">LIvia Bator</div>
                <div className="role text-base text-[#718EBF]">CEO</div>
              </div>
              <button className="relative flex p-6 py-7 items-center justify-center bg-white dark:bg-gray-700 dark:shadow-gray-500 shadow-sm shadow-blue-300 rounded-full">
                <FontAwesomeIcon
                  icon={faGreaterThan}
                  className="w-5 font-normal dark:text-gray-300"
                />
              </button>
            </div>

            <div className="bottom flex gap-4 items-center">
              <div className="text-gray- text-lg text-[#718EBF] dark:text-gray-400">
                Write Amount
              </div>
              <div className="flex items-center text-base text-[#718EBF] bg-[#EDF1F7] dark:bg-gray-700 dark:text-gray-400 rounded-[50px] py-1 pl-6 grow justify-end">
                <div className="flex w-full grow">525.50</div>
                <button className="flex gap-2 w-full p-4 rounded-[50px] bg-[#1814F3] text-white grow px-6 text-6 items-center justify-center">
                  Send
                  <FontAwesomeIcon
                    icon={faPaperPlane}
                    className="w-5 font-normal"
                  />
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div className="cards-container center-content flex flex-col gap-6 xs:w-full lg:w-[55%]">
        <h1 className="page text-xl font-semibold text-colorBody-1 dark:text-gray-300">
          Balance History
        </h1>

        <div className="flex w-full gap-6 p-8 bg-white dark:bg-[#232328] dark:border-gray-500 rounded-3xl border-solid border-gray-200 border-[0.5px] shadow-sm xs:w-[85%] sm:w-full">
          <div className="leftCanva pb-10 flex flex-col items-end justify-between text-sm text-[#718EBF]">
            <span>400</span>
            <span>300</span>
            <span>200</span>
            <span>100</span>
            <span>0</span>
          </div>
          <AreaComp data={BalanceData} />
        </div>
      </div>
    </section>
  );
};

export default Bottom;
