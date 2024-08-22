
export default function Topleft({ topic }: { topic: string }) {
  return (
    
    <div className="flex gap-20 w-1/4 text-black mt-3 ml-5 items-center">
      <p className="font-normal text-3xl ">{topic}</p>
    </div>
  );
}
