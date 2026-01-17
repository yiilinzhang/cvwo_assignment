import { Post } from "../components/post";
import { useQuery, useQueryClient } from "@tanstack/react-query";
import { useAuth } from "../hooks/useAuth";
import axios from "axios";
import { Typography } from "@mui/material";
import { EnvelopeOpenIcon } from "@phosphor-icons/react";

export default function Posts({ params }) {
  const { user, isLoading: userLoading } = useAuth();
  const queryClient = useQueryClient();
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
  const topics = queryClient.getQueryData(["topics"])?.payload?.data;

  const currTopic = topics?.find((t) => t.topic_id === Number(params.id))
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
