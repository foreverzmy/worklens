import {
	type FC,
	type PropsWithChildren,
	createContext,
	useCallback,
	useContext,
	useState,
} from 'react';

const RepoPathContext = createContext(
	{} as {
		repoPath: string;
		setRepoPath: (repoPath: string) => void;
	},
);

export const useRepoPath = () => {
	const ctx = useContext(RepoPathContext);

	return ctx.repoPath;
};

export const useSetRepoPath = () => {
	const ctx = useContext(RepoPathContext);

	const setRepoPath = useCallback(
		(repoPath: string) => {
			if (!repoPath) {
				return;
			}

			const usp = new URLSearchParams(location.search);
			usp.set('repo', repoPath);
			location.search = usp.toString();

			return ctx.setRepoPath(repoPath);
		},
		[ctx.setRepoPath],
	);

	return setRepoPath;
};

export const useResetRepoPath = () => {
	const ctx = useContext(RepoPathContext);

	const resetRepoPath = useCallback(() => {
		const usp = new URLSearchParams(location.search);
		usp.delete('repo');
		location.search = usp.toString();

		return ctx.setRepoPath('');
	}, [ctx.setRepoPath]);

	return resetRepoPath;
};

export const RepoPathProvider: FC<PropsWithChildren> = ({ children }) => {
	const [repoPath, setRepoPath] = useState(() => {
		const usp = new URLSearchParams(location.search);

		return usp.get('repo') || '';
	});

	return (
		<RepoPathContext.Provider value={{ repoPath, setRepoPath }}>
			{children}
		</RepoPathContext.Provider>
	);
};
