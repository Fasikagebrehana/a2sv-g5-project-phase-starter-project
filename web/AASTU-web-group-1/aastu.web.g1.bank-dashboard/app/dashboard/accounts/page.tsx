"use client"
import CreditCard from "../_components/Credit_Card";
import InfoCard from "./components/InfoCard";
import LastTransaction from "./components/LastTransaction";
import { ChartWeekly } from "@/components/ui/BarchartWeekly";
import Invoices from "./components/Invoices";
import { useUser } from "@/contexts/UserContext";

const Accounts = () => {
  const { isDarkMode } = useUser();
  return (
    <div
      className={`p-5 md:pr-20 ${
        isDarkMode ? "bg-gray-700 text-gray-100" : "bg-[#F5F7FA] text-gray-800"
      }`}
    >
      <div
        className={`grid grid-cols-2 md:grid-cols-4 gap-4 md:flex-row p-3 ${
          isDarkMode ? "bg-gray-700" : "bg-white"
        }`}
      >
        <div
          className={`p-4 rounded-3xl ${
            isDarkMode ? "bg-gray-800" : "bg-white"
          } shadow-md`}
        >
          <InfoCard
            title="My balance"
            amount={12750}
            image="/icons/money.svg"
            color={isDarkMode ? "bg-yellow-300" : "bg-yellow-500"}
          />
        </div>
        <div
          className={`p-4 rounded-3xl ${
            isDarkMode ? "bg-gray-800" : "bg-white"
          } shadow-md`}
        >
          <InfoCard
            title="Income"
            amount={5600}
            image="/icons/handmoney.svg"
            color={isDarkMode ? "bg-blue-300" : "bg-blue-500"}
          />
        </div>
        <div
          className={`p-4 rounded-3xl ${
            isDarkMode ? "bg-gray-800" : "bg-white"
          } shadow-md`}
        >
          <InfoCard
            title="Expense"
            amount={3460}
            image="/icons/001-medical.svg"
            color={isDarkMode ? "bg-pink-300" : "bg-pink-500"}
          />
        </div>
        <div
          className={`p-4 rounded-3xl ${
            isDarkMode ? "bg-gray-800" : "bg-white"
          } shadow-md`}
        >
          <InfoCard
            title="Total Saving"
            amount={7920}
            image="/icons/003-saving.svg"
            color={isDarkMode ? "bg-green-300" : "bg-green-500"}
          />
        </div>
      </div>
      <div
        className={`md:flex md:gap-12 ${
          isDarkMode ? "bg-gray-700" : "bg-white"
        }`}
      >
        <div className="w-[70%]">
          <h1 className="text-xl mb-4">Last Transactions</h1>
          <div className="">
            <LastTransaction
              image="/icons/Bell.svg"
              alttext="bell"
              description="Spotify Subscription"
              transaction={-150}
              colorimg={isDarkMode ? "bg-green-400" : "bg-green-500"}
              date="25 Jan 2021"
              type="Shopping"
              account="1234 ****"
              status="Pending"
            />
            <LastTransaction
              image="/icons/tools.svg"
              alttext="bell"
              description="Mobile Service"
              transaction={-340}
              colorimg={isDarkMode ? "bg-blue-400" : "bg-blue-500"}
              date="25 Jan 2021"
              type="Service"
              account="1234 ****"
              status="Completed"
            />
            <LastTransaction
              image="/icons/user.svg"
              alttext="settings"
              description="Emilly Wilson"
              transaction={780}
              colorimg={isDarkMode ? "bg-pink-400" : "bg-pink-500"}
              date="25 Jan 2021"
              type="Transfer"
              account="1234 ****"
              status="Completed"
            />
          </div>
        </div>
        <div className="md:w-[30%]">
          <div className="flex justify-between font-inter text-[16px] font-semibold mb-4">
            <h4>My Cards</h4>
            <h4
              className={`text-blue-500 ${isDarkMode ? "text-blue-300" : ""}`}
            >
              See All
            </h4>
          </div>
          <div className="mb-4">
            <CreditCard
              id="1234"
              balance={5894}
              semiCardNumber="37781234"
              cardHolder="Ediy Cusuma"
              expiryDate="2024-08-20T07:06:50.283Z"
              cardType={"Visa"}
            />
          </div>
        </div>
      </div>
      <div
        className={`md:flex gap-6 mb-5 ${
          isDarkMode ? "bg-gray-700" : "bg-white"
        }`}
      >
        <div>
          <h1 className="text-xl mb-4"> Debit & Credit Overview </h1>
          <div className="mb-4">
            <ChartWeekly />
          </div>
        </div>
        <div>
          <div>
            <h1 className="text-xl mb-4">Invoices Sent</h1>
            <div
              className={`rounded-xl ${isDarkMode ? "bg-gray-800" : "bg-white"}`}
            >
              <Invoices
                image="/icons/apple.svg"
                title="Apple Store"
                date="5h ago"
                expense={450}
                color={isDarkMode ? "bg-green-300" : "bg-green-500"}
              />
              <Invoices
                image="/icons/useryello.svg"
                title="Michael"
                date="2 days ago"
                expense={450}
                color={isDarkMode ? "bg-yellow-300" : "bg-yellow-500"}
              />
              <Invoices
                image="/icons/playstation.svg"
                title="Apple Store"
                date="2 days ago"
                expense={1085}
                color={isDarkMode ? "bg-blue-300" : "bg-blue-500"}
              />
              <Invoices
                image="/icons/user.svg"
                title="William"
                date="10 days ago"
                expense={90}
                color={isDarkMode ? "bg-pink-300" : "bg-pink-500"}
              />
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Accounts;
