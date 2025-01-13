import { useBehaviorSubjectState } from '@service/useBehaviorSubjectState';
import { useStream } from '@service/useStream';
import type { RepoReference } from '@typings/message';
import { Button } from 'primereact/button';
import { Card } from 'primereact/card';
import { Tag } from 'primereact/tag';
import type { FC } from 'react';
import { scan } from 'rxjs';
import { InlineCode } from '../../common/components/InlineCode';
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
		<div className="flex flex-col gap-2 p-2">
			<Card
				title={
					<div className="flex justify-between">
						<span>Repo Path</span>
						<Button className="py-2" size="small" onClick={resetRepoPath}>
							Reset
						</Button>
					</div>
				}
			>
				<InlineCode className="ml-2">{repoPath}</InlineCode>
			</Card>
			<Card title="Remote">
				<div>
					{remotes.map((remote) => (
						<p key={remote.name} className="m-0">
							{remote.name}：
							{remote.urls.map((url: string) => (
								<InlineCode key={url}>{url}</InlineCode>
							))}
						</p>
					))}
				</div>
			</Card>
			<Card title="Branches">
				<div>
					{head !== null && (
						<p key={head.name}>
							{head.shortName}
							{head.name === head?.name && (
								<Tag className="ml-1" value="head" />
							)}
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
				</div>
			</Card>
		</div>
	);
};
