# Go Hacker News Confidence

This project grabs Hacker News, sorts news items according to the lower bound of a Wilson score confidence interval for a Bernoulli parameter, and spits them out as reformatted HTML. We assume points as upvotes and comments as downvotes. 

### Slightly longer explanation
This project is how I weaned myself off Hacker News. I've found that the most
worthwhile HN articles to read are the ones with high-upvote-to-comment ratio.
For example, controversial articles e.g. "Why Rails sucks" tend to have upvote count ~= comments. Whereas high-quality analysis and world-changing news generates little conversation but much upvoting.

Reading just the top 5 stories ranked according to this score saves you a *ton*
of time.

### Original Project
Written with Scrapy, found [here](https://github.com/ejamesc/hacker-news-confidence).
