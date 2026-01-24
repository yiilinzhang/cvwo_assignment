import { Post } from "../components/Post";
import { useQuery, useQueryClient } from "@tanstack/react-query";
import { useAuth } from "../hooks/useAuth";
import { Typography } from "@mui/material";
import { EnvelopeOpenIcon } from "@phosphor-icons/react";
import { useParams } from "react-router";
import axios from "axios";

type PostItem = {
  post_id: number;
  title: string;
  content: string;
  user_id: number;
};

type PostsResponse = {payload?: {data?: PostItem[]}}

type Topic = {topic_id: number; title: string}
type TopicsResponse = {payload?: {data?: Topic[]}}
export default function Posts() {
  const { id } = useParams();

  const { userID, isLoading: userLoading } = useAuth();
  const queryClient = useQueryClient();
  const { isLoading: postLoading, data: postData } = useQuery<PostsResponse>({
    queryKey: [`posts`, id ?? "all"],
    queryFn: async () => {
      const url = id
        ? `http://localhost:8000/posts/${id}`
        : "http://localhost:8000/posts";
      const response = await axios.get(url);

      return response.data;
    },
  });
  const topics = queryClient.getQueryData<TopicsResponse>(["topics"])?.payload?.data;

  const currTopic = topics?.find((t) => t.topic_id === Number(id))?.title;
  return (
    <div className="flex flex-col items-center gap-8 py-6 px-20 h-screen">
      <Typography >
        {currTopic || "All Topics"}
      </Typography>

      {postData?.payload?.data?.length === 0 ? (
        <div className="w-full flex items-center justify-center flex-col pt-40 ">
          <EnvelopeOpenIcon size={80} />
          <Typography fontSize={25}>
            There is no content for this page.
          </Typography>
        </div>
      ) : (
        postData?.payload?.data?.map((post: PostItem) => (
          <Post
            key={post.post_id}
            id={post.post_id}
            title={post.title}
            content={post.content}
            isOwner={Number(post.user_id) === userID}
            showChat={true}
          />
        ))
      )}
    </div>
  );
}
