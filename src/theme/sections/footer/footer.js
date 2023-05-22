
import React from 'react';
import useDocusaurusContext from '@docusaurus/useDocusaurusContext';
import useBaseUrl from '@docusaurus/useBaseUrl';

export default function Footer() {

	const context = useDocusaurusContext();
	const {siteConfig = {}} = context;

	const component = 'shift-intro';

	return (
		<footer id="footer">
			<div className="credits">
				<div className="container">
					<div className="row justify-content-center">
						<div className="col-md-6 cia text-center text-md-left pt-2 mb-4 mb-md-2">
							<span><div dangerouslySetInnerHTML={{__html: siteConfig.themeConfig.footer.copyright}}></div></span>
						</div>
						<div className="col-md-6 social-footer text-center text-md-right">
							<ul>
								<li><a href="https://github.com/bucketeer-io/bucketeer" target="_blank"><i className="fa fa-github"></i></a></li>
								<li><a href="https://app.slack.com/client/T08PSQ7BQ/C043026BME1" target="_blank"><i className="fa fa-slack"></i></a></li>
								<li><a href="https://twitter.com/bucketeer_io" target="_blank"><i className="fa fa-twitter"></i></a></li>
							</ul>
						</div>
					</div>
				</div>
			</div>
		</footer>
	);
}