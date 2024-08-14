import Image from "next/image";
import ImageComponent from "./QuickTransferImage"
export default function Home() {
  return (
    <div className = "bg-white rounded-3xl my-4 mx-4">
      <div className="flex flex-col gap-3 px-5 py-5">
        {/*  Image Component  */}
            <div className="flex gap-6 py-2 justify-between items-center">
              <ImageComponent/>
              <ImageComponent/>
              <ImageComponent/>
              <div className="flex items-center rounded-full min-h-10  min-w-10 justify-center shadow-xl">
                <svg width="6" height="10" viewBox="0 0 9 15" fill="none" xmlns="http://www.w3.org/2000/svg">
                  <path d="M1 1L7.5 7.5L1 14" stroke="#718EBF" stroke-width="2"/>
                </svg>
              </div>
            </div> 

            <div className="flex  text-[#718EBF] text-xs justify-between items-center text-nowrap gap-3 ">
              <p>Write Amount</p>
              <div className="flex gap-6 bg-[#EDF1F7] rounded-full">
                    <div className=" bg-[#EDF1F7] rounded-full flex gap-80 items-center  px-5 py-3">
                      <p>525.50</p>
                    </div>
                    <div >
                      <button className="flex gap-2 bg-[#1814F3] rounded-full px-5 py-3">
                          <p className="text-white text-xs">Send</p>         
                          <svg width="15" height="10" viewBox="0 0 26 23" fill="none" xmlns="http://www.w3.org/2000/svg">
                            <path d="M25.9824 0.923369C26.1091 0.333347 25.5307 -0.164153 24.9664 0.0511577L0.490037 9.39483C0.195457 9.50731 0.000610804 9.78965 1.43342e-06 10.105C-0.000607937 10.4203 0.193121 10.7034 0.487294 10.817L7.36317 13.4726V21.8369C7.36317 22.1897 7.60545 22.4963 7.94873 22.5779C8.28972 22.659 8.64529 22.4967 8.80515 22.1796L11.6489 16.5364L18.5888 21.6868C19.011 22.0001 19.6178 21.8008 19.7714 21.2974C26.251 0.0528342 25.9708 0.97674 25.9824 0.923369ZM19.9404 3.60043L8.01692 12.092L2.88664 10.1106L19.9404 3.60043ZM8.8866 13.3428L19.2798 5.94118C10.3366 15.3758 10.8037 14.8792 10.7647 14.9317C10.7067 15.0096 10.8655 14.7058 8.8866 18.6327V13.3428ZM18.6293 19.8197L12.5206 15.2862L23.566 3.63395L18.6293 19.8197Z" fill="white"/>
                          </svg>
                      </button>
                    </div>
              </div>
            </div>  
        </div>
      </div>
  
  );
}
