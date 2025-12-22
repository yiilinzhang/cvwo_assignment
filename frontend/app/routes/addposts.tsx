import { useQuery } from "@tanstack/react-query";
import { useNavigate } from "react-router";
//TODO add typing later
export default function addPosts() {
  const { isLoading, data } = useQuery({
    queryKey: [`topiclist`],
    queryFn: async () => {
      const response = await fetch("http://localhost:8000/topics");
      return await response.json();
    },
  });
  const navigate = useNavigate()

  const handleSubmit = async (e) => {
    e.preventDefault();

    const form = e.target;
    const formData = new FormData(form);
    const body = {
      title: formData.get("title"),
      content : formData.get("content"),
      topic_id: Number(formData.get("topic_id"))
    }
    console.log(body)
    try {
      const response = await fetch("http://localhost:8000/posts", {
        method: "POST",
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(body),
      });
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      console.log(response)
      const result = await response.json();
    } catch (error) {
      console.error("Error:", error);
      alert("Failed to create post.");
      return
    }
    alert("Successfully created post.");
    navigate("/")

  };

  return (
    <form method="post" onSubmit={handleSubmit}>
      <div className="flex flex-col gap-1 items-center">
        <text className="font-medium text-3xl p-6">
          Create a new post today!
        </text>

        <div className="flex flex-col justify-start gap-2">
          <text className="font-medium text-2xl">Topic</text>
          <select className="w-40 h-10 border-2 rounded-2xl p-2" name="topic_id">
            {data?.payload.data.map((title: any) => (
              <option className="text-xl" key={title.topic_id} value={title.topic_id}>{title.title}</option>
            ))}
          </select>

          <text className="font-medium text-2xl ">Title</text>
          <textarea
            className="border-2 rounded-2xl p-2 w-96 h-20"
            name="title"
          />

          <text className="font-medium text-2xl">Content</text>
          <textarea className="border-2 rounded-2xl p-2 h-44" name="content" />
        </div>

        <button
          className="bg-[#9BE3FF] w-34 p-2 m-4 rounded-2xl hover:cursor-pointer"
          type="submit"
        >
          <text className="font-medium text-2xl">Post now!</text>
        </button>
      </div>
    </form>
  );
}
