import { useQuery } from "@tanstack/react-query";

export default function addPosts() {
  const { isLoading, data } = useQuery({
    queryKey: [`topiclist`],
    queryFn: async () => {
      const response = await fetch("http://localhost:8000/topics");
      return await response.json();
    },
  });
  return (
    <div className="flex flex-col gap-1 items-center">
      <text className="font-medium text-3xl p-6">Create a new post today!</text>

      <div className="flex flex-col justify-start gap-2">
        <text className="font-medium text-2xl">Topic</text>
        <select className="w-40 h-10 border-2 rounded-2xl p-2">
          {data?.payload.data.map((title) => (
            <option className="text-xl">{title.title}</option>
          ))}
        </select>

        <text className="font-medium text-2xl">Title</text>
        <textarea className="border-2 rounded-2xl p-2 w-96 h-20" />

        <text className="font-medium text-2xl">Content</text>
        <textarea className="border-2 rounded-2xl p-2 h-44" />
      </div>

      <button className="bg-[#9BE3FF] w-34 p-2 m-4 rounded-2xl hover:cursor-pointer">
        <text className="font-medium text-2xl">Post now!</text>
      </button>
    </div>
  );
}
