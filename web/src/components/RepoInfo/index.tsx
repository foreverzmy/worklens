import { useBehaviorSubjectState } from '@service/useBehaviorSubjectState';
import { useStream } from '@service/useStream';
import type { RepoReference } from '@typings/message';
import { Button } from 'primereact/button';
import { Card } from 'primereact/card';
import { Tag } from 'primereact/tag';
import type { FC } from 'react';
import { scan } from 'rxjs';
import { useRepoPath, useResetRepoPath } from '../../context';

export const RepoInfo: FC = () => {
	const repoPath = useRepoPath();
	const resetRepoPath = useResetRepoPath();
	const [remotes, { sub$Ref: remoteSub$Ref }] = useBehaviorSubjectState<any[]>(
		[],
	);
	const [head, { sub$Ref: headSub$Ref }] =
		useBehaviorSubjectState<RepoReference | null>(null);

	const [branches, { sub$Ref: branchSub$Ref }] = useBehaviorSubjectState<
		RepoReference,
		RepoReference[]
	>(
		[],
		scan((acc, cur) => [...acc, cur], [] as any[]),
	);

	useStream('/repo/info', {
		remote: remoteSub$Ref.current,
		branch: branchSub$Ref.current,
		head: headSub$Ref.current,
	});

	return (
		<div className="m-4 flex flex-col gap-4">
			<Card title="仓库信息">
				<p className="m-0">
					{repoPath}
					<Button size="small" onClick={resetRepoPath}>
						切换仓库
					</Button>
				</p>
				<p>
					{remotes.map((remote) => (
						<p key={remote.name}>
							{remote.name}：{remote.urls.join(',')}
						</p>
					))}
				</p>
			</Card>
			<Card title="分支信息">
				<p>
					{head !== null && (
						<p key={head.name}>
							{head.name} {head.name === head?.name && <Tag value="head" />}
						</p>
					)}
					{branches
						.filter((b) => !head || b.name !== head.name)
						.map((branch) => (
							<p key={branch.name}>
								{branch.shortName}
								{branch.name === head?.name && <Tag value="head" />}
							</p>
						))}
				</p>
			</Card>
		</div>
	);
};
