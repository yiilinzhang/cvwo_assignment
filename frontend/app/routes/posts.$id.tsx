import { Post } from "../components/post";
import { useQuery, useQueryClient } from "@tanstack/react-query";
import { useAuth } from "../hooks/useAuth";
import { Typography } from "@mui/material";
import { EnvelopeOpenIcon } from "@phosphor-icons/react";
import { useParams } from "react-router";
import axios from "axios";

export default function Posts() {
  const { id } = useParams();

  const { user, isLoading: userLoading } = useAuth();
  const queryClient = useQueryClient();
  const { isLoading: postLoading, data: postData } = useQuery({
    queryKey: [`posts`, id?? "all"],
    queryFn: async () => {
      const url = id
        ? `http://localhost:8000/posts/${id}`
        : "http://localhost:8000/posts";
      const response = await axios.get(url)

      return response.data;
    },
  });
  const topics = queryClient.getQueryData(["topics"])?.payload?.data;

  const currTopic = topics?.find((t) => t.topic_id === Number(id))
    ?.title;
  return (
    <div className="flex flex-col items-center gap-8 py-6 px-20 h-screen">
      {/* TODO change this so it shows the curr topic */}
      <text className="text-3xl font-bold w-full ">
        {currTopic || "All Topics"}
      </text>

      {postData?.payload.data.length === 0 ? (
        <div className="w-full flex items-center justify-center flex-col pt-40 ">
        <EnvelopeOpenIcon size={80}/>
        <Typography fontSize={25}>There is no content for this page.</Typography></div>
      ) : (
        postData?.payload.data.map((post) => (
          <Post
            key={post.post_id}
            id={post.post_id}
            title={post.title}
            content={post.content}
            isOwner={Number(post.user_id) === user?.payload.data}
            showChat={true}
          />
        ))
      )}
    </div>
  );
}
