const jobSitesEl = document.getElementById("jobSites");
const jobTitlesEl = document.getElementById("jobTitles");
const ignoreEl = document.getElementById("ignore");
const includeEl = document.getElementById("include");
const fromEl = document.getElementById("from");
const textEl = document.getElementById("stat-text");
const searchBtnEl = document.getElementById("search-button");
const jobsListEl = document.getElementById("jobs-list");
const seTypeEl = document.getElementById("mode");

fromEl.setAttribute("max", new Date().toISOString().split("T")[0]);

let jobsData = [];

const calculateDateDayValue = (date) => {
  const aDayInMilliSecs = 86400000;
  const currentDate = new Date().getTime();
  const selectedDate = new Date(date).getTime();

  return Math.ceil((currentDate - selectedDate) / aDayInMilliSecs);
};

const validateSearchInput = (data) => {
  const errors = {};

  if (!data.sites || data.sites.trim() === "") {
    errors.jobSites = "Job Sites field is empty";
  }

  if (!data.titles || data.titles.trim() === "") {
    errors.jobTitles = "Job Titles field is empty";
  }

  if (!data.mode || data.mode.trim() === "") {
    errors.mode = "Select Search API to use";
  }

  if (!data.from) {
    errors.from = "From field is empty";
  }

  return errors;
};

const searchForJobs = async (reqBody) => {
  try {
    const baseUrl = "https://job-searcher.onrender.com";
    let url = `${baseUrl}/api/v1/search`;

    let reqObj = {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(reqBody),
    };

    const resp = await fetch(url, reqObj);

    const result = await resp.json();
    const { status, data } = result;

    if (status) {
      jobsData = data.jobs;

      textEl.innerText = `✅ Jobs search was successful.`;
      textEl.classList.remove("hide");
      searchBtnEl.innerHTML = '<i class="fa fa-clipboard"></i> Search';

      setTimeout(() => {
        textEl.innerText = "";
        textEl.classList.add("hide");
      }, 5000);
    } else {
      textEl.innerText = `❌ Failed! Unable to search for jobs`;
      textEl.classList.remove("hide");
      searchBtnEl.innerHTML = '<i class="fa fa-clipboard"></i> Search';

      setTimeout(() => {
        textEl.innerText = "";
        textEl.classList.add("hide");
      }, 5000);
    }
  } catch (error) {
    textEl.innerText = `❌ Failed! Unable to search for jobs`;
    textEl.classList.remove("hide");
    searchBtnEl.innerHTML = '<i class="fa fa-clipboard"></i> Search';

    console.log(error);

    setTimeout(() => {
      textEl.innerText = "";
      textEl.classList.add("hide");
    }, 5000);
  }
};

document.getElementById("close").addEventListener("click", (evt) => {
  evt.preventDefault();

  jobsData = [];

  document.getElementById("search").classList.remove("hide");
  document.getElementById("modal").classList.add("hide");

  jobsListEl.innerHTML = "";
});

searchBtnEl.addEventListener("click", async (evt) => {
  evt.preventDefault();

  const sites = jobSitesEl.value;
  const titles = jobTitlesEl.value;
  const ignore = ignoreEl.value;
  const include = includeEl.value;
  const from = fromEl.value;
  const mode = seTypeEl.value;

  const validationErrorList = Object.keys(
    validateSearchInput({
      sites,
      titles,
      from,
      mode,
    })
  );

  if (validationErrorList.length > 0) {
    validationErrorList.forEach((keyName) => {
      document.getElementById(keyName).classList.add("search-box-err");
    });

    setTimeout(() => {
      validationErrorList.forEach((keyName) => {
        document.getElementById(keyName).classList.remove("search-box-err");
      });
    }, 5000);

    return;
  }

  searchBtnEl.innerHTML =
    '<i class="fa fa-spinner fa-spin"></i> Searching for jobs...';

  let ignores = ignore.split(",");
  let includes = include.split(",");
  const jobSites = sites.split(",");
  const jobTitles = titles.split(",");
  const numberOfDays = calculateDateDayValue(from);

  if (ignores.length === 1 && ignores[0].trim() === "") {
    ignores = [];
  }

  if (includes.length === 1 && includes[0].trim() === "") {
    includes = [];
  }

  await searchForJobs({
    jobSites,
    jobTitles,
    ignore: ignores,
    include: includes,
    from: numberOfDays,
    mode,
  });

  if (jobsData && jobsData.length > 0) {
    for (let index = 0; index < jobsData.length; index += 1) {
      const { title, description, location, url } = jobsData[index];

      jobsListEl.innerHTML += `<div class="col-md-4" id="dv-${index}">
      <div class="card pd">
        <h4>${title}</h4>
  
        <p>
          ${description}
        </p>
  
        <p>${location}</p>
  
        <a
          href="${url}" id="a-${index}">${url}</a
        >
      </div>
    </div>`;
    }

    document.getElementById("search").classList.add("hide");
    document.getElementById("modal").classList.remove("hide");
  }
});
