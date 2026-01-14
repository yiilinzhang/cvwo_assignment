import { Post } from "../components/post";
import { useQuery } from "@tanstack/react-query";
import { useAuth } from "../hooks/useAuth"

export default function Posts({ params }) {
    const {user, isLoading : userLoading} = useAuth()
  const topicId = params?.id;
  const { isLoading: postLoading, data: postData } = useQuery({
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
      {/* TODO change this so it shows the curr topic */}
      <text className="text-4xl font-bold w-full ">All Topic</text>
      {postData?.payload.data.map((post) => (
        <Post key={post.post_id} id={post.post_id} title={post.title} content={post.content} isOwner={Number(post.user_id) === user?.payload.data} />
      ))}
    </div>
  );
}
