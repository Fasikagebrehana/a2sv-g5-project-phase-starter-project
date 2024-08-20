"use client";
import { useEffect } from "react";
import { useAppDispatch } from "@/app/Redux/store/store";
import TransactionService from "@/app/Services/api/transactionApi";
import {
  setTran,
  setStatus,
  setBalHist,
  setError,
  setExpense,
  setIncome,
} from "../slices/TransactionSlice";

const useTranDispatch = (accessToken: string) => {
  const dispatch = useAppDispatch();

  useEffect(() => {
    const fetchInitialCards = async () => {
      try {
        dispatch(setStatus("loading"));
        const transaction: any = await TransactionService.getTransactions(
          accessToken
        );

        const history: any = await TransactionService.balanceHistory(
          accessToken
        );
        const expense: any = await TransactionService.getExpenseData(
          accessToken
        );
        const income: any = await TransactionService.getIncomeData(accessToken);

        if (transaction) {
          dispatch(setTran(transaction.content));
          dispatch(setBalHist(history));
          dispatch(setExpense(expense.content));
          dispatch(setIncome(income.content));
          dispatch(setStatus("succeeded"));
        }
      } catch (error) {
        dispatch(setError("Failed to fetch transaction"));
        dispatch(setStatus("failed"));
      }
    };

    fetchInitialCards();
  }, [dispatch, accessToken]);
};

export default useTranDispatch;
