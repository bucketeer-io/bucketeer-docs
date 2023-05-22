
import React from 'react';
import useDocusaurusContext from '@docusaurus/useDocusaurusContext';
import useBaseUrl from '@docusaurus/useBaseUrl';


export default function Intro() {

	const context = useDocusaurusContext();
	const {siteConfig = {}} = context;

	const component = 'shift-social';

	return (
	    <section className="calltoaction">
	        <div className="container container-fluid">
	        	<div className="calltoaction__bottom text--center">
	        		
	        		<div className="row">
						<div className="col-md-4 mb-5 mb-md-0 social-call-white text-center item">
							<i className="github-icon"></i>
							<span className="title">Contributions welcome!</span>
							<p>We value every contribution.</p>
							<p>Join us on GitHub to learn how you can help.</p>
							<a href="https://docs.bucketeer.io/contribution-guide">Read more...</a>
						</div>
						<div className="col-md-4 mb-5 mb-md-0 social-call-white text-center item">
							<i className="slack-icon"></i>
							<span className="title">Join the conversation!</span>
							<p>Have a question?</p>
							<p>Learn more by talking with other contributors in Cloud Native Slack via <a href="https://app.slack.com/client/T08PSQ7BQ/C043026BME1" className="bucketeer-link" target="_blank">#bucketeer</a> channel.</p>
							<a href="https://app.slack.com/client/T08PSQ7BQ/C043026BME1" target="_blank">Read more...</a>
						</div>
						<div className="col-md-4 mb-1 mb-md-0 social-call-white text-center item">
							<i className="twitter-icon"></i>
							<span className="title">Follow us on Twitter!</span>
							<p>Don’t miss out!</p>
							<p>Follow us for feature announcements and other news.</p>
							<a href="https://twitter.com/bucketeer_io" target="_blank">Read more...</a>
						</div>
					</div>

	        	</div>
	        </div>
	    </section>
	);
}