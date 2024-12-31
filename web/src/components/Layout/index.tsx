import type { FC } from "react";
import { useRepoPath } from "../../context";
import { RepoInfo } from "../RepoInfo";
import { RepoInput } from "../RepoInput";

export const Layout: FC = () => {
	const repoPath = useRepoPath();

	if (!repoPath) {
		return <RepoInput />;
	}

	return <RepoInfo />;
};
