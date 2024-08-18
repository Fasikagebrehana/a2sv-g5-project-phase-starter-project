"use client";

import React, { useState, useEffect } from "react";
import axios from "axios";
import { BarChartComponent } from "./components/BarChartComponent";
import { TableComponent } from "./components/TableComponent";
import Card from "../components/common/card";
import columns from "./components/columns";
import TableCard from "./components/TableComponentMobile";
import Chip_card1 from "@/public/assets/image/Chip_Card1.png";
import Chip_card2 from "@/public/assets/image/Chip_Card2.png";
import Chip_card3 from "@/public/assets/image/Chip_Card3.png";

const creditCardColor = {
  cardOne: {
    cardBgColor: "bg-blue-500 rounded-3xl text-white",
    bottomBgColor: "flex justify-between p-4 bg-blue-400 rounded-bl-3xl rounded-br-3xl",
    imageCreditCard: Chip_card1,
    grayCircleColor: false,
  },
  cardTwo: {
    cardBgColor: "bg-blue-700 rounded-3xl text-white",
    bottomBgColor: "flex justify-between p-4 bg-blue-600 rounded-bl-3xl rounded-br-3xl",
    imageCreditCard: Chip_card2,
    grayCircleColor: false,
  },
  cardThree: {
    cardBgColor: "bg-white rounded-3xl textblack",
    bottomBgColor: "",
    imageCreditCard: Chip_card3,
    grayCircleColor: true,
  },
};

const Transactions: React.FC = () => {
  const [data, setData] = useState<any[]>([]);
  const [error, setError] = useState<string | null>(null);
  const [activeLink, setActiveLink] = useState<string>('');

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await axios.get('https://bank-dashboard-6acc.onrender.com/transactions', {
          headers: {
            Authorization: `Bearer your_actual_bearer_token_here`, // Replace with your actual token
          },
        });
        
        // Transform API data to match the format expected by the table
        const transformedData = response.data.data.map((item: any) => ({
          column1: item.description,
          column2: item.transactionId,
          column3: item.type,
          column4: "N/A", // Update this if you have card info
          column5: new Date(item.date).toLocaleDateString(),
          column6: `$${item.amount.toFixed(2)}`, // Format amount as currency
          column7: "N/A", // Update this if you have receipt info
        }));

        setData(transformedData);
      } catch (error) {
        console.error("Failed to fetch data:", error);
        setError("Failed to fetch data. Please check the console for more details.");
      }
    };

    fetchData();
  }, []);

  const handleLinkClick = (linkName: string) => {
    setActiveLink(linkName);
  };

  return (
    <div className="flex flex-col gap-6 p-6">
      {/* First Row: Cards and BarChart */}
      <div className="flex flex-col lg:flex-row gap-6">
        <div className="flex-1 lg:w-1/2 overflow-hidden">
          <div className="flex justify-between items-center mb-4">
            <h2 className="text-lg font-semibold font-Inter text-[#343C6A]">
              My Cards
            </h2>
            <button className="text-[#343C6A] font-Inter font-medium">
              + Add Card
            </button>
          </div>
          {/* Flex container for cards */}
          <div className="flex gap-4">
            <Card creditCardColor={creditCardColor.cardOne} />
            <Card creditCardColor={creditCardColor.cardThree} />
          </div>
        </div>
        <div className="flex-1 lg:w-1/2">
          <h2 className="text-lg font-semibold mb-4 font-Inter text-[#343C6A]">
            My Expenses
          </h2>
          <BarChartComponent />
        </div>
      </div>

      {/* Second Row: Links and Conditional Rendering Based on Device Size */}
      <div className="flex flex-col w-full">
        <div className="flex flex-row justify-start items-center mb-4 overflow-x-auto">
          <a
            href="#"
            className={`text-lg font-normal text-[#343C6A] mx-2 transition-all ${activeLink === 'recent' ? 'font-bold' : ''}`}
            onClick={() => handleLinkClick('recent')}
          >
            Recent Transactions
          </a>
          <a
            href="#"
            className={`text-lg font-normal text-[#343C6A] mx-2 transition-all ${activeLink === 'income' ? 'font-bold' : ''}`}
            onClick={() => handleLinkClick('income')}
          >
            Income
          </a>
          <a
            href="#"
            className={`text-lg font-normal text-[#343C6A] mx-2 transition-all ${activeLink === 'expenses' ? 'font-bold' : ''}`}
            onClick={() => handleLinkClick('expenses')}
          >
            Expenses
          </a>
        </div>

        <div className="hidden lg:flex flex-col w-full">
          {/* Render TableComponent for desktop and tablet */}
          {error ? <div>{error}</div> : <TableComponent columns={columns} data={data} />}
        </div>

        <div className="lg:hidden flex flex-col w-full">
          {/* Render TableCard for mobile */}
          <TableCard />
        </div>
      </div>
    </div>
  );
}

export default Transactions;
