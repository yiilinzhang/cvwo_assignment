import { Comments } from "~/components/comments";
import { Post } from "../components/post";
import { useQuery } from "@tanstack/react-query";

//Should accept the post details so i can abstract the post id and selet comments
export default function postComments({ params }) {
  const postId = params?.id;
  //Fetch all comments for the particular post
  const { isLoading, data } = useQuery({
    queryKey: [`comments`, params.id],
    queryFn: async () => {
      const url = `http://localhost:8000/comments/${postId}`;
      const response = await fetch(url);
      return await response.json();
    },
  });
  return (
    <div className="flex flex-col items-center gap-8 py-6 px-20">
      {/* TODO check how i should pass in post prop requery or drill */}
      {data?.payload.data.map((comment) => (
        <Comments username={comment.name} content={comment.content} owner={true} />
      ))}
    </div>
  );
}

