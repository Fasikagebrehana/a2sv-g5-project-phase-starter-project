import React from "react";
import Card from "./LoanComponents/Card";
import { TableDemo } from "./LoanComponents/Table/columns";
import { loanTypes } from "@/constants";
const Loans = () => {
  return (
    <div className="bg-transparent md:px-4">
      <div className="flex gap-2 overflow-x-scroll scrollbar-hidden my-3 scroll md:w-[90%] md:mx-auto">
        {loanTypes.map((item) => (
          <div className="my-2" key={item.name}>
            <Card {...item} />
          </div>
          ))}
      </div>
      <div className="mt-5 md:ml-4">
        <h1 className="text-md font-[500] md:font-[600] text-[#333B69] md:text-left md:pl-20 md:text-[35px] pl-4">Active Loans Overview</h1>
        <TableDemo />
      </div>

    </div>

  );
};

export default Loans;
