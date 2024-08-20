"use client";
import React, { useEffect, useState } from "react";
import DescriptionCard from "@/app/components/Card/DescriptionCard";
import ServicesCard from "@/app/components/Card/ServicesCard";
import axios from "axios";

interface BankService {
  id: string;
  name: string;
  details: string;
  numberOfUsers: number;
  status: string;
  type: string;
  icon: string;
  colors: string;
}

const Services = () => {
  const colors = [
    "bg-orange-100",
    "bg-pink-100",
    "bg-blue-100",
    "bg-green-100",
    "bg-pink-100",
  ];
  const [services, setServices] = useState<BankService[]>([]);
  const [pageNumber, setPageNumber] = useState(1);
  const accessToken =
    "eyJhbGciOiJIUzM4NCJ9.eyJzdWIiOiJiZXRzZWxvdCIsImlhdCI6MTcyNDE1NjE5MywiZXhwIjoxNzI0MjQyNTkzfQ.x6tV4QOyUslrkXYJrRfWZjwD8MBgQaRx5UyV0qR0byV68wwf2rVPkWkgu3VZRFpA";

  async function fetchData(accessToken: string) {
    try {
      const response = await axios.get(
        `https://bank-dashboard-6acc.onrender.com/bank-services?page=0&size=5`,
        {
          headers: {
            Authorization: `Bearer ${accessToken}`,
          },
        }
      );
      setServices(response.data.data.content);
      console.log(services);
    } catch (error) {
      console.error("There was a problem with the axios request:", error);
    }
  }

  useEffect(() => {
    fetchData(accessToken);
  }, []);

  return (
    <div className="w-[96%] xxs:pt-4 xs:pt-20 md:pt-5 lg:pt-0">
      <div className="ml-5 lg:ml-0 ">
        <div className="mr-5 lg:mr-0 flex gap-4 overflow-x-auto lg:overflow-x-visible [&::-webkit-scrollbar]:hidden [-ms-overflow-style:none] [scrollbar-width:none] lg:pl-10 lg:pt-10 w-full">
          {/* <div className="w-[100%] lg:w-[350px] flex-shrink-0"> */}
          <ServicesCard
            img="/assets/lifeInsurance.svg"
            title="Life Insurance"
            desc="Unlimited Protection"
          />
          {/* </div> */}
          {/* <div className="w-[100%] lg:w-[350px] flex-shrink-0"> */}
          <ServicesCard
            img="/assets/shoppingBag.svg"
            title="Shopping"
            desc="Buy. Think. Grow"
          />
          {/* </div> */}
          {/* <div className="w-[100%] lg:w-[350px] flex-shrink-0"> */}
          <ServicesCard
            img="/assets/safety.svg"
            title="Safety"
            desc="We are your allies"
          />
        </div>
        {/* </div> */}

        <div>
          <p className="font-semibold text-[22px] text-[#343C6A] pt-5 pb-5 lg:p-10 ">
            Bank Services List
          </p>
          <div className="w-full flex flex-col grow items-start px-4">
            {services.length > 0 ? (
              services.map((service, index) => (
                <DescriptionCard
                  key={service.id}
                  img={service.icon}
                  title={service.name}
                  desc={service.details}
                  colOne="Number of Users"
                  descOne={service.numberOfUsers}
                  colTwo="Status"
                  descTwo={service.status}
                  colThree="Type"
                  descThree={service.type}
                  btn="View Details"
                  color={colors[index]}
                />
              ))
            ) : (
              <p className="pl-10">No services available</p>
            )}
          </div>
        </div>
      </div>
    </div>
  );
};

export default Services;
