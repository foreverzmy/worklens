import { Button } from "primereact/button";
import { InputText } from "primereact/inputtext";
import { type FC, useState } from "react";
import { useSetRepoPath } from "../../context";

export const RepoInput: FC = () => {
	const [value, setValue] = useState("");
	const setRepoPath = useSetRepoPath();

	const handleSubmit = () => {
		setRepoPath(value);
	};

	return (
		<div className="w-screen h-screen flex flex-col items-center">
			<h1 className="text-8xl" style={{ marginTop: "15vh" }}>
				Work Lens
			</h1>
			<InputText
				className="p-inputtext-lg"
				style={{
					marginTop: "10vh",
					width: "61.8vw",
					minWidth: "600px",
					textAlign: "center",
				}}
				placeholder="Please input the absolute path to the repository's root directory, starting with '/'."
				value={value}
				onChange={(e) => setValue(e.target.value)}
			/>
			<Button className="mt-20" onClick={handleSubmit}>
				Submit
			</Button>
		</div>
	);
};
