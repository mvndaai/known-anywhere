/*
    Format for now
    "key": "<string value>" | <{"single": "<single value>", "multi": "<multi value (things with 's')>"}>
    // Strings can have a %s for variables
*/
export default {
    test: {
        string: "This is a string",
        withCount : {
            single: "There is {0} thing",
            multi: "There are {0} things"
        }
    },
    search : {
        search: "Search",
        select:{
            username: "Username",
            coupon: "Coupon",
            name: "Name/Alias",
        },
    },
    home: {
        welcome: "Welcome",
        overview: "Known Socially is a platform to let find your favorite people on any social platform and support them by using their coupon codes.",
        howItWorks: "How it works",
        createGroup: "Create a 'group' and then add links and coupons to that group",
        logIn: "When you log in you can create a group for yourself and verify that your the links",
        shareGroup: "Share the group with your followers and they can support you by using your coupon codes",
        features: "Features",
        searchPeople: "Search for people by username, name or coupon code",
        searchCoupons: "Search for coupon codes by domain",
        // followGroups: "Follow groups",
        // showCoupons: "Show coupons from those you follow when",
        // notifications: "Notifications when a group you follow adds a new coupon or link",
        // linkGroups: "Groups not linked to a user can be linked to other groups"
    }
}