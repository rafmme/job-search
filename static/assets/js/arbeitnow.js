const jobSearchKeywordEl = document.getElementById("jobKeywords");
const apiDataLengthEl = document.getElementById("apiDataLength");
const textEl = document.getElementById("stat-text");
const searchBtnEl = document.getElementById("search-button");
const jobsListEl = document.getElementById("jobs-list");

const fetchJobs = async (i) => {
  try {
    const baseUrl = "https://www.arbeitnow.com/api/job-board-api";
    const url = `${baseUrl}?page=${i}`;

    const resp = await fetch(url);
    const result = await resp.json();

    return result.data;
  } catch (error) {
    console.log(error);

    textEl.innerText = `❌ Failed! Unable to search for jobs`;
    textEl.classList.remove("hide");
    searchBtnEl.innerHTML = '<i class="fa fa-clipboard"></i> Search';

    setTimeout(() => {
      textEl.innerText = "";
      textEl.classList.add("hide");
    }, 5000);
  }
};

const getArbeitnowJobsPosting = async (keywords, apiDataLength) => {
  let jobsList = [];

  for (let i = 1; i <= apiDataLength; i += 1) {
    const jobData = await fetchJobs(i);
    jobsList = jobsList.concat(jobData);
  }

  if (keywords.trim() === "" || !keywords) {
    return jobsList;
  }

  const filteredJobs = jobsList.filter((job) => {
    return (
      job.title.toLowerCase().includes(keywords) ||
      job.slug.toLowerCase().includes(keywords) ||
      job.description.toLowerCase().includes(keywords) ||
      job.tags.join(",").toLowerCase().includes(keywords)
    );
  });

  return filteredJobs;
};

document.getElementById("close").addEventListener("click", (evt) => {
  evt.preventDefault();
  document.getElementById("search").classList.remove("hide");
  document.getElementById("modal").classList.add("hide");
});

searchBtnEl.addEventListener("click", async (evt) => {
  evt.preventDefault();

  searchBtnEl.innerHTML =
  '<i class="fa fa-spinner fa-spin"></i> Searching for jobs...';

  const keywords = jobSearchKeywordEl.value;
  let apiDataLength = Number.parseInt(apiDataLengthEl.value, 10);

  if (!apiDataLength) {
    apiDataLength = 2;
  }

  const jobsData = await getArbeitnowJobsPosting(
    keywords.toLowerCase(),
    apiDataLength
  );

  if (jobsData && jobsData.length > 0) {
    textEl.innerText = `✅ Jobs search was successful.`;
    textEl.classList.remove("hide");
    searchBtnEl.innerHTML = '<i class="fa fa-clipboard"></i> Search';

    setTimeout(() => {
        textEl.innerText = "";
        textEl.classList.add("hide");
      }, 5000);

    for (let index = 0; index < jobsData.length; index += 1) {
      const {
        title,
        description,
        location,
        url,
        tags,
        created_at,
        company_name,
        remote,
        job_types
      } = jobsData[index];

      jobsListEl.innerHTML += `<div class="col-md-4" id="dv-${index}">
      <div class="card pd">
        <h4>${title}</h4>
  
        <p>
          ${description}
        </p>

        <p><b>Company:</b> ${company_name}</p>

        <p><b>Location:</b> ${location}</p>
    
        <p><b>Remote:</b> ${remote}</p>

        <p><b>Posted At:</b> ${new Date(created_at * 1000)}</p>

        <p><b>Job Types:</b> ${job_types.join(",")}</p>

        <p><b>Tags:</b> ${tags}</p>
  
        <a
          href="${url}" id="a-${index}"><b>${url}</b></a
        >
      </div>
    </div>`;
    }

    document.getElementById("search").classList.add("hide");
    document.getElementById("modal").classList.remove("hide");
  }
});
