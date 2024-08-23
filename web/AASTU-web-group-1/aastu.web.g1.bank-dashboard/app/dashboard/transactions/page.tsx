"use client";

import { useState, useEffect } from "react";

import { useUser } from "@/contexts/UserContext";
import { Loading } from "../_components/Loading";
import { TransactionCards } from "./component/transactionCards";
import { TransactionTable } from "./component/TransactionTable";

const Transactions = () => {
  const { isDarkMode } = useUser();
  const [dataFetched, setDataFetched] = useState(false);
  const [transactionCardsLoaded, setTransactionCardsLoaded] = useState(false);
  const [transactionTableLoaded, setTransactionTableLoaded] = useState(false);

  useEffect(() => {
    if (transactionCardsLoaded && transactionTableLoaded) {
      setDataFetched(true);
    }
  }, [transactionCardsLoaded, transactionTableLoaded]);

  return (
    <div
      className={`space-y-5 ${dataFetched?"p-10":""} ${
        isDarkMode ? "bg-gray-700 text-gray-200" : "bg-[#F5F7FA] text-gray-900"
      }`}
    >
      {!dataFetched && <Loading />}
      {/* First row */}
      <TransactionCards
        onLoadingComplete={() => setTransactionCardsLoaded(true)}
      />
      {/* Second row */}
      <TransactionTable
        onLoadingComplete={() => setTransactionTableLoaded(true)}
      />
    </div>
  );
};

export default Transactions;
