import { TransactionProps } from "@/types";


export async function getIncomes(page: number, size: number) {
  try {
    const accessToken =
      "eyJhbGciOiJIUzM4NCJ9.eyJzdWIiOiJlbW5ldC10ZXMiLCJpYXQiOjE3MjQwNDg3NjAsImV4cCI6MTcyNDEzNTE2MH0.baqrlqraepMSM7YMMdUKSUd2j_Z3ui7hyQjvw8b-ENDP9cly77sngGLsVvC3lpC-";

    const res = await fetch(
      `https://bank-dashboard-6acc.onrender.com/transactions/incomes?page=${page}&size=${size}`,
      {
        method: "GET",
        headers: {
          Authorization: `Bearer ${accessToken}`,
        },
        body: null,
      }
    );
    if (!res.ok) {
      throw new Error("faild to fetch data");
    }
    const incomes: TransactionProps = await res.json();
    console.log(incomes);

    return incomes.data;
  } catch (error) {
    console.log("An error occurred:", error);
    // alert("An unexpected error occurred. Please try again later.");
  }
}
