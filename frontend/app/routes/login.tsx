export function LoginPage() {
  return (
    <div className="w-screen h-screen bg-[#F5F5F5] flex justify-center pt-40">
      <div className="flex flex-col w-96 h-64 bg-white items-center py-8 gap-5">
        <text className="font-bold text-2xl">Sign In/ Sign Up</text>
        <div className="flex flex-col">
        <text className="text-l">username</text>
        <textarea className="border rounded-2xl py-1" rows={1} cols={30}/></div>
        <button className="bg-[#9BE3FF] p-2 mt-2">
          <text className="font-bold text-xl">Enter!</text>
        </button>
      </div>
    </div>
  );
}
