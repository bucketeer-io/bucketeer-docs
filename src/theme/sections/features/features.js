
import React from 'react';
import useDocusaurusContext from '@docusaurus/useDocusaurusContext';
import useBaseUrl from '@docusaurus/useBaseUrl';

export default function Features() {

	const context = useDocusaurusContext();
	const {siteConfig = {}} = context;

	const component = 'shift-featured';

	return (
		<main className="home">
			
			<section className="featured" data-aos="fade-in">
				<div className="container">
						<div className="row align-items-center">
								<div className="col-lg-6 sub-item">
										<div className="item-container d-none d-sm-inline-block" data-aos="fade-right-small">
												<h2 className="title">Control your features</h2>
												<p>Feature Flags are a software development tool that ensures an efficient, low-risk release cycle by enabling or disabling features in real time without deploying new code.<br />
												Bucketeer offers advanced features, such as dark launch and staged rollouts, that perform limited releases based on user attributes, devices, and other segments.</p>
												<p>With feature flags, you can continuously deploy new features that are still in development without making them visible to the users.<br />
												This makes it possible to separate the deployment from the release, which allows teams to manage the feature's entire lifecycle.</p>
										</div>
										<div className="item-container d-inline-block d-sm-none" data-aos="fade-left-small">
												<h2 className="title">Control your features</h2>
												<div className="col-lg-6 sub-item d-inline-block d-lg-none" data-aos="fade-left-small"><div className="img-holder text-center"><img src="img/pages/home/feature_flag.png" className="img-fluid" /></div></div>
												<p>Feature Flags are a software development tool that ensures an efficient, low-risk release cycle by enabling or disabling features in real time without deploying new code.<br />
												Bucketeer offers advanced features, such as dark launch and staged rollouts, that perform limited releases based on user attributes, devices, and other segments.</p>
												<p>With feature flags, you can continuously deploy new features that are still in development without making them visible to the users.<br />
												This makes it possible to separate the deployment from the release, which allows teams to manage the feature's entire lifecycle.</p>
										</div>
								</div>
								<div className="col-lg-6 sub-item d-none d-lg-inline-block" data-aos="fade-left-small"><div className="img-holder text-center"><img src="img/pages/home/feature_flag.png" className="img-fluid" /></div></div>
						</div>
				</div>
			</section>

			<section className="featured" data-aos="fade-in">
				<div className="container">
						<div className="row align-items-center">
								<div className="col-lg-6 sub-item d-none d-lg-inline-block" data-aos="fade-right-small"><div className="img-holder text-center"><img src="img/pages/home/experimentation.png" className="img-fluid" /></div></div>
								<div className="col-lg-6 sub-item">
										<div className="item-container d-none d-sm-inline-block" data-aos="fade-left-small">
												<h2 className="title">Better decisions with data</h2>
												<p>A/B testing is an experimentation process to compare one or multiple versions of an application. It helps your team analyze what performs better, and make better data-driven decisions without relying on intuition or personal experience.</p>
												<p>Bucketeer uses the Bayesian probabilities to analyze which variable of your A/B test is likely to perform better. Because it requires a smaller sample size, you can get results faster with lower experimentation costs than a Frequentist probabilities.</p>
										</div>
										<div className="item-container d-inline-block d-sm-none" data-aos="fade-left-small">
												<h2 className="title">Better decisions with data</h2>
												<div className="col-lg-6 sub-item d-inline-block d-lg-none" data-aos="fade-left-small"><div className="img-holder text-center"><img src="img/pages/home/experimentation.png" className="img-fluid" /></div></div>
												<p>A/B testing is an experimentation process to compare one or multiple versions of an application. It helps your team analyze what performs better, and make better data-driven decisions without relying on intuition or personal experience.</p>
												<p>Bucketeer uses the Bayesian probabilities to analyze which variable of your A/B test is likely to perform better. Because it requires a smaller sample size, you can get results faster with lower experimentation costs than a Frequentist probabilities.</p>
										</div>
								</div>
						</div>
				</div>
			</section>
			
			<section className="featured" data-aos="fade-in">
				<div className="container">
						<div className="row align-items-center">
								<div className="col-lg-6 sub-item">
										<div className="item-container d-none d-sm-inline-block" data-aos="fade-right-small">
												<h2 className="title">Increased Development Speed</h2>
												<p>Trunk-based development reduces lead time, speeding up the process from code review to release with the use of feature flags.</p>
												<p>Developers can implement a new feature by disabling the flag and deploying it to the main branch at any time.<br />
												This helps prevent merge conflicts caused by long-lived branches and reduces code review costs.</p>
												<p>This practice is essential for large teams to ensure that a shared branch is always releasable without delaying the QA time and affecting the user.</p>
										</div>
										<div className="item-container d-inline-block d-sm-none" data-aos="fade-left-small">
												<h2 className="title">Increased Development</h2>
												<div className="col-lg-6 sub-item d-inline-block d-lg-none" data-aos="fade-left-small"><div className="img-holder text-center"><img src="img/pages/home/trunk-based_development.png" className="img-fluid" /></div></div>
												<p>Trunk-based development reduces lead time, speeding up the process from code review to release with the use of feature flags.</p>
												<p>Developers can implement a new feature by disabling the flag and deploying it to the main branch at any time.<br />
												This helps prevent merge conflicts caused by long-lived branches and reduces code review costs.</p>
												<p>This practice is essential for large teams to ensure that a shared branch is always releasable without delaying the QA time and affecting the user.</p>
										</div>
								</div>
								<div className="col-lg-6 sub-item d-none d-lg-inline-block" data-aos="fade-left-small"><div className="img-holder text-center"><img src="img/pages/home/trunk-based_development.png" className="img-fluid" /></div></div>
						</div>
				</div>
			</section>

		</main>
	);
}