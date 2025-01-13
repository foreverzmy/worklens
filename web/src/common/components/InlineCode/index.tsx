import type { CSSProperties, FC, PropsWithChildren } from 'react';

interface InlineCodeProps {
	className?: string;
	style?: CSSProperties;
}

export const InlineCode: FC<PropsWithChildren<InlineCodeProps>> = ({
	className,
	style,
	children,
}) => {
	return (
		<code
			className={`inline-block bg-gray-100 text-red-600 font-mono p-1 rounded ${className}`}
			style={style}
		>
			{children}
		</code>
	);
};
