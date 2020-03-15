/* globals gauge*/
"use strict";
const { openBrowser,write, closeBrowser, goto, press,text, contains } = require('taiko');
const assert = require("assert");
const headless = process.env.headless_chrome.toLowerCase() === 'true';
const site_url = process.env.TEST_SITE_URL

beforeSuite(async () => {
    await openBrowser({ headless: headless })
});

afterSuite(async () => {
    await closeBrowser();
});

step("Goto Google's search page", async () => {
    await goto(site_url);
});

step("Search for <query>", async (query) => {
    await write(query);
    await press('Enter');
});

step("Page contains <content>", async (content) => {
    assert.ok(await text(contains(content)).exists());
});
