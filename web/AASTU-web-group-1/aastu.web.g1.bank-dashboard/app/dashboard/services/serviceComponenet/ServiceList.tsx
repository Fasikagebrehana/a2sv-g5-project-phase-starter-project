import { useUser } from "@/contexts/UserContext";
import { LoansIcon } from "../serviceIcons/icons";
import Image from "next/image";

interface itemProp {
  icon: string;
  name: string;
}
const ServiceList = ({ icon, name }: itemProp) => {
  const {isDarkMode} = useUser();
  return (
    
    <div className={`flex justify-between items-center p-3 md:p5 rounded-xl ${isDarkMode ? "bg-gray-800":"bg-white border-2 "} `}>
      <div className="flex ml-1  gap-1">
        <div>
          <Image src={icon} alt="" width={30} height={30} />
        </div>
        <div className="px-2">
          <div className="mt-1 font-medium">{name}</div>
          <div className="text-xs text-[#718EBF]">
            it&apos;s a long established
          </div>{" "}
        </div>
      </div>
      <div className="hidden md:flex justify-between gap-4 w-1/2">
        <div className="px-2">
          <div className="mt-1 font-medium">Business loans</div>
          <div className="text-xs text-[#718EBF]">
            it&apos;s a long established
          </div>{" "}
        </div>
        <div className="px-2">
          <div className="mt-1 font-medium">Business loans</div>
          <div className="text-xs text-[#718EBF]">
            it&apos;s a long established
          </div>{" "}
        </div>
        <div className="px-2">
          <div className="mt-1 font-medium">Business loans</div>
          <div className="text-xs text-[#718EBF]">
            it&apos;s a long established
          </div>{" "}
        </div>
      </div>
      <div className="text-[#1814F3] text-sm font-bold lg:border-2 lg:px-6 lg:rounded-full lg:my-auto lg:py-1 ">
        View Detail
      </div>
    </div>
  );
};

export default ServiceList;
