import { Comments } from "~/components/comments";
import { Post } from "../components/post";
import { useQuery, useQueryClient } from "@tanstack/react-query";
import { useAuth } from "~/hooks/useAuth";
import { Button, TextField, Typography } from "@mui/material";
import { PlusIcon } from "@phosphor-icons/react";
import { useState } from "react";
import axios from "axios";

//Should accept the post details so i can abstract the post id and selet comments
//TODO chaneg this from cache to a new query
export default function postComments({ params }) {
  const { user, isLoading: userLoading } = useAuth();
  const queryClient = useQueryClient();
  const [isEditing, setIsEditing] = useState(false);
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
  const post = queryClient.getQueryData(["posts", "all"])?.payload.data;
  console.log(post);
  const currPost = post.find((p) => p.post_id === Number(postId));
  console.log(currPost);
  const handleSubmit =  (e) => {
    e.preventDefault();

    const form = e.target;
    const formData = new FormData(form);
    const body = {
      content: formData.get("comment"),
    };
    console.log(body);

    //TODO change to axios
    axios.post(`http://localhost:8000/comments/${postId}`, body, { withCredentials: true }).then(() => {
        alert("Successfully created comment.");
        setIsEditing(false)
        queryClient.invalidateQueries({ queryKey: [`comments`, params.id] })
      });
    }
    return (
      <div className="flex flex-col gap-6 py-6 px-20">
        <Post
          id={currPost.postId}
          title={currPost.title}
          content={currPost.content}
          isOwner={Number(currPost.user_id) === user?.payload.data}
          showChat={false}
        />
        <hr />
        {/* TODO check how i should pass in post prop requery or drill */}
        <div className="flex flex-row items-center gap-2">
          {/* <Typography fontWeight={500} fontSize="1.5rem">
        Comments
      </Typography> */}
          {isEditing ? (
            <form method="post" onSubmit={handleSubmit} className="w-full flex flex-col gap-2">
                <TextField
                  fullWidth={true}
                  multiline={true}
                  rows={3}
                  name="comment"
                />
                <Button
                  variant="outlined"
                  disableRipple
                  sx={{
                    color: "black",
                    borderColor: "black",
                    borderRadius: 20,
                    width: 200
                  }}
                  type="submit"
                >
                  <Typography>Post Comment</Typography>
                </Button>
            </form>
          ) : (
            <Button
              variant="outlined"
              disableRipple
              sx={{ color: "black", borderColor: "black", borderRadius: 20 }}
              onClick={() => setIsEditing(true)}
            >
              <PlusIcon />
              <Typography>Add a Comment</Typography>
            </Button>
          )}
        </div>
        {data?.payload.data.map((comment) => (
          <Comments
            username={comment.name}
            content={comment.content}
            owner={Number(comment.user_id) === user?.payload.data}
          />
        ))}
      </div>
    );
  };

