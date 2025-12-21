import { Post } from "../components/post";
import { useQuery } from "@tanstack/react-query";

export default function Posts({ params }) {
  const topicId = params?.id;
  const { isLoading, data } = useQuery({
    queryKey: [`posts`, params.id ?? "all"],
    queryFn: async () => {
      const url = topicId
        ? `http://localhost:8000/posts/${params.id}`
        : "http://localhost:8000/posts";
      const response = await fetch(url);
      return await response.json();
    },
  });
  return (
    <div className="flex flex-col items-center gap-8 py-6 px-20">
      <text className="text-4xl font-bold w-full ">All topics</text>
      {data?.payload.data.map((post) => (
        <Post title={post.title} content={post.content} owner={true} />
      ))}
    </div>
  );
}
