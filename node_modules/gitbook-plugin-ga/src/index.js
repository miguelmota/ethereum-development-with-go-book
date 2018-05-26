const GitBook = require('gitbook-core');
const { React } = GitBook;

/**
 * Ga wrapper to track page view.
 * @type {ReactClass}
 */
let GAWrapper = React.createClass({
    propTypes: {
        children: React.PropTypes.node,
        config:   GitBook.PropTypes.map.isRequired,
        location: GitBook.PropTypes.Location.isRequired
    },

    componentDidMount() {
        const { config } = this.props;

        // Load ga
        (function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
        (i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
        m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
        })(window,document,'script','//www.google-analytics.com/analytics.js','ga');

        // Initialize ga
        ga('create', config.get('token'), config.get('configuration'));
    },

    componentDidUpdate(prevProps) {
        const { location: prevLocation } = prevProps;
        const { location } = this.props;

        // We do not track hash change
        const hasChanged = !prevLocation.delete('hash').equals(location.delete('hash'));

        if (hasChanged) {
            ga('send', 'pageview', location.toString());
        }
    },

    render() {
        const { children } = this.props;
        return children ? React.Children.only(children) : null;
    }
});
GAWrapper = GitBook.connect(
    GAWrapper,
    ({ history, config }) => ({
        location: history.location,
        config: config.getForPlugin('ga')
    })
);

module.exports = GitBook.createPlugin({
    activate: (dispatch, getState, { Components }) => {
        dispatch(Components.registerComponent(GAWrapper, { role: 'website:body' }));
    }
});
