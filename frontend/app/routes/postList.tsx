import { Post } from "../components/post";
import { useQuery } from "@tanstack/react-query";

export function PostList() {
  const {isLoading, data} = useQuery({
              queryKey: [`posts`], 
              queryFn: getPosts
          })
  return (<div className="flex flex-col items-center gap-8 py-6 px-20">
    <text className="text-4xl font-bold w-full ">All topics</text>
    {data?.payload.data.map((post) => <Post title={post.title} content={post.content} owner={true}/>)}
  </div>
);
}

const getPosts = async () => {
    const response = await fetch("http://localhost:8000/posts") 
    return await response.json()
}