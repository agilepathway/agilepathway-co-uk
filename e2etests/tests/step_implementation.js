/* globals gauge*/
"use strict";
const { openBrowser,write, closeBrowser, goto, press,text, contains } = require('taiko');
const assert = require("assert");
const headless = process.env.headless_chrome.toLowerCase() === 'true';
const site_url = process.env.TEST_SITE_URL

beforeSuite(async () => {
    await openBrowser({ headless: true, args: [
        '--disable-gpu',
         '--disable-dev-shm-usage',
         '--disable-setuid-sandbox',
         '--no-first-run',
         '--no-sandbox',
         '--no-zygote']})
});

afterSuite(async () => {
    await closeBrowser();
});

step("Goto Agile Pathway's home page", async () => {
    await goto(site_url);
});

step("Page contains <content>", async (content) => {
    assert.ok(await text(contains(content)).exists());
});
