package link

var html1 = `
<html>
<body>
  <h1>Hello!</h1>
  <a href="/other-page">A link to another page</a>
</body>
</html>
`

var html2 = `
<html>
<head>
  <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css">
</head>
<body>
  <h1>Social stuffs</h1>
  <div>
    <a href="https://www.twitter.com/joncalhoun">
      Check me out on twitter
      <i class="fa fa-twitter" aria-hidden="true"></i>
    </a>
    <a href="https://github.com/gophercises">
      Gophercises is on <strong>Github</strong>!
    </a>
  </div>
</body>
</html>
`

var html3 = `
<!DOCTYPE html>
<!--[if lt IE 7]> <html class="ie ie6 lt-ie9 lt-ie8 lt-ie7" lang="en"> <![endif]-->
<!--[if IE 7]>    <html class="ie ie7 lt-ie9 lt-ie8"        lang="en"> <![endif]-->
<!--[if IE 8]>    <html class="ie ie8 lt-ie9"               lang="en"> <![endif]-->
<!--[if IE 9]>    <html class="ie ie9"                      lang="en"> <![endif]-->
<!--[if !IE]><!-->
<html lang="en" class="no-ie">
<!--<![endif]-->

<head>
  <title>Gophercises - Coding exercises for budding gophers</title>
</head>

<body>
  <section class="header-section">
    <div class="jumbo-content">
      <div class="pull-right login-section">
        Already have an account?
        <a href="#" class="btn btn-login">Login <i class="fa fa-sign-in" aria-hidden="true"></i></a>
      </div>
      <center>
        <img src="https://gophercises.com/img/gophercises_logo.png" style="max-width: 85%; z-index: 3;">
        <h1>coding exercises for budding gophers</h1>
        <br/>
        <form action="/do-stuff" method="post">
          <div class="input-group">
            <input type="email" id="drip-email" name="fields[email]" class="btn-input" placeholder="Email Address" required>
            <button class="btn btn-success btn-lg" type="submit">Sign me up!</button>
            <a href="/lost">Lost? Need help?</a>
          </div>
        </form>
        <p class="disclaimer disclaimer-box">Gophercises is 100% FREE, but is currently in beta. There will be bugs, and things will be changing significantly over the coming weeks.</p>
      </center>
    </div>
  </section>
  <section class="footer-section">
    <div class="row">
      <div class="col-md-6 col-md-offset-1 vcenter">
        <div class="quote">
          "Success is no accident. It is hard work, perseverance, learning, studying, sacrifice and most of all, love of what you are doing or learning to do." - Pele
        </div>
      </div>
      <div class="col-md-4 col-md-offset-0 vcenter">
        <center>
          <img src="https://gophercises.com/img/gophercises_lifting.gif" style="width: 80%">
          <br/>
          <br/>
        </center>
      </div>
    </div>
    <div class="row">
      <div class="col-md-10 col-md-offset-1">
        <center>
          <p class="disclaimer">
            Artwork created by Marcus Olsson (<a href="https://twitter.com/marcusolsson">@marcusolsson</a>), animated by Jon Calhoun (that's me!), and inspired by the original Go Gopher created by Renee French.
          </p>
        </center>
      </div>
    </div>
  </section>
</body>
</html>
`

var html4 = `
<html>
<body>
  <a href="/dog-cat">dog cat <!-- commented text SHOULD NOT be included! --></a>
</body>
</html>
`

var html5 = `

<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8" />
    
    <link rel="shortcut icon" href="/favicon.ico" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <meta name="theme-color" content="#000000" />

    <link
      href="https://fonts.googleapis.com/css?family=Lato|Satisfy|Kristi&display=swap"
      rel="stylesheet"
    />
    <link href="/assets/styles.css" rel="stylesheet">

    <title>Extracting Node Text | courses.calhoun.io</title>
  </head>
  <body class="bg-grey-100">
    
    
  
  <div class="w-full bg-grey-grad">
    <div class="flex items-center w-full text-white px-6 py-6">
      <div class="flex items-baseline">
        <div class="text-xl md:text-4xl text-white font-cursive">
          
          <a class="md:flex items-center" href="/courses">
            Go Courses
            <div class="text-sm md:text-lg text-teal-100 md:pl-2">
              by
              <span class="text-lg md:text-2xl font-signature">
                Jon Calhoun
              </span>
            </div>
          </a>
        </div>
        
      </div>
      <div class="flex-grow"></div>
      <div>
        
          <div class="flex">
            <a class="flex px-3 h-10 items-center rounded-l hover:text-white text-teal-100 bg-teal-400 border-r border-teal-500 hover:bg-teal-500 no-underline" href="/user/edit">
              <img src="https://www.gravatar.com/avatar/e7fca7528751b6d803949184446af2f6?d=mp&amp;s=24" class="hidden md:block md:w-6 md:mr-2 rounded-full opacity-75">
              
              mertgnksn@gmail.com
            </a>
            <form class="" action="/signout" method="POST">
              <button class="flex items-center px-3 h-10 rounded-r hover:text-white text-teal-100 bg-teal-400 border-l border-teal-500 hover:bg-teal-500" type="submit">
                <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" class="w-6">
                  <circle cx="12" cy="12" r="10" class="stroke-current fill-none"></circle>
                  <path class="fill-current"
                    d="M13.41 12l2.83 2.83a1 1 0 0 1-1.41 1.41L12 13.41l-2.83 2.83a1 1 0 1 1-1.41-1.41L10.59 12 7.76 9.17a1 1 0 0 1 1.41-1.41L12 10.59l2.83-2.83a1 1 0 0 1 1.41 1.41L13.41 12z">
                  </path>
                </svg>
              </button>
            </form>
          </div>
        
      </div>
    </div>
  </div>
  <div class="w-full h-1 bg-teal-400"></div>
  
    <div class="w-full bg-teal-100">
      <div class="py-2 px-6 text-md">
        <div class="flex items-center overflow-x-auto">
          
            <a class="flex-none no-underline text-teal-500" href="/courses">
              Courses
            </a>
            
<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" class="w-8 mx-2 pt-1 flex-none text-teal-200">
  <path class="fill-current" d="M10.3 8.7a1 1 0 0 1 1.4-1.4l4 4a1 1 0 0 1 0 1.4l-4 4a1 1 0 0 1-1.4-1.4l3.29-3.3-3.3-3.3z"></path>
</svg>

          
            <a class="flex-none no-underline text-teal-500" href="/courses/cor_gophercises">
              Gophercises
            </a>
            
<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" class="w-8 mx-2 pt-1 flex-none text-teal-200">
  <path class="fill-current" d="M10.3 8.7a1 1 0 0 1 1.4-1.4l4 4a1 1 0 0 1 0 1.4l-4 4a1 1 0 0 1-1.4-1.4l3.29-3.3-3.3-3.3z"></path>
</svg>

          
          
            <a class="flex-none no-underline text-teal-600" href="/lessons/les_goph_22">
              Extracting Node Text
            </a>
          
        </div>
      </div>
    </div>
  

  <div class="w-full mb-4 pt-8">
    <div class="flex-row-reverse md:flex">
      <div class="py-6 md:py-0 flex-1 max-w-1920 px-6 mx-auto">
        
          <div class="relative font-white">
            <div class="embed-container rounded">
              <iframe class="embed-video" src="https://player.vimeo.com/video/242407321" frameborder="0" webkitallowfullscreen mozallowfullscreen allowfullscreen></iframe>
            </div>
          </div>
        

        

        <div class="pt-6">
          
<div class="rounded shadow bg-white">
  
  <div class="px-4 py-6 md:flex items-center rounded bg-grey-050">
    
      <div class="pb-4 md:pb-0">
        <a class="no-underline text-grey-600 hover:text-grey-800"
        href="/lessons/les_goph_21">
          <div class="flex items-center">
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" class="w-6 flex-none">
              <path
                class="fill-current"
                d="M9.41 11H17a1 1 0 0 1 0 2H9.41l2.3 2.3a1 1 0 1 1-1.42 1.4l-4-4a1 1 0 0 1 0-1.4l4-4a1 1 0 0 1 1.42 1.4L9.4 11z"
              ></path>
            </svg>
            <span class="flex-grow">
              Building Link Structs
            </span>
          </div>
        </a>
      </div>
    
    <div class="md:flex-grow"></div>
    
      <a class="no-underline text-grey-600 hover:text-grey-800"
      href="/lessons/les_goph_23">
        <div class="flex items-center">
          <span class="md:flex-grow">
            Parsing the Examples
          </span>
          <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" class="w-6 flex-none">
            <path
              class="fill-current"
              d="M14.59 13H7a1 1 0 0 1 0-2h7.59l-2.3-2.3a1 1 0 1 1 1.42-1.4l4 4a1 1 0 0 1 0 1.4l-4 4a1 1 0 0 1-1.42-1.4l2.3-2.3z"
            ></path>
          </svg>
        </div>
      </a>
    
  </div>
</div>

        </div>
        
        <div class="pt-6 px-6">
          <h4 class="text-grey-600 pb-2">Additional Links</h4>
          <ul class="list-disc text-grey-500 underline">
            
              <li class="ml-6"><a href="https://www.github.com/gophercises/link">Github repo with a writeup of the exercise details</a></li>
            
              <li class="ml-6"><a href="https://github.com/gophercises/link/tree/solution">Solution Code</a></li>
            
          </ul>
        </div>
        
        
          <div class="text-grey-700 px-6 py-8 mt-8 mb-16 bg-white rounded shadow">
            <div class="prose max-w-full overflow-scroll">
              <h1>Exercise #4: HTML Link Parser</h1>

<p><a href="https://gophercises.com/exercises/link" rel="nofollow"><img src="https://img.shields.io/badge/exercise%20status-released-green.svg?style=for-the-badge"/></a></p>

<h3>Exercise details</h3><hr/>

<p>In this exercise your goal is create a package that makes it easy to parse an HTML file and extract all of the links (<code>&lt;a href=&#34;&#34;&gt;...&lt;/a&gt;</code> tags). For each extracted link you should return a data structure that includes both the <code>href</code>, as well as the text inside the link. Any HTML inside of the link can be stripped out, along with any extra whitespace including newlines, back-to-back spaces, etc.</p>

<p>Links will be nested in different HTML elements, and it is very possible that you will have to deal with HTML similar to code below.</p>

<pre><code>&lt;a href=&#34;/dog&#34;&gt;
  &lt;span&gt;Something in a span&lt;/span&gt;
  Text not in a span
  &lt;b&gt;Bold text!&lt;/b&gt;
&lt;/a&gt;
</code></pre>

<p>In situations like these we want to get output that looks roughly like:</p>

<pre><code>Link{
  Href: &#34;/dog&#34;,
  Text: &#34;Something in a span Text not in a span Bold text!&#34;,
}
</code></pre>

<p>Once you have a working program, try to write some tests for it to practice using the testing package in go.</p>

<h4>Notes</h4>

<p><strong>1. Use the x/net/html package</strong></p>

<p>I recommend checking out the <a href="https://godoc.org/golang.org/x/net/html" rel="nofollow">x/net/html</a> package for this task, which you will need to <code>go get</code>. It is provided by the Go team, but isn’t included in the standard library. This makes it a little easier to parse HTML files.</p>

<p><strong>2. Ignore nested links</strong></p>

<p>You can ignore any links nested inside of another link. Eg with following HTML:</p>

<pre><code>&lt;a href=&#34;#&#34;&gt;
  Something here &lt;a href=&#34;/dog&#34;&gt;nested dog link&lt;/a&gt;
&lt;/a&gt;
</code></pre>

<p>It is okay if your code returns only the outside link.</p>

<p><strong>3. Get something working before focusing on edge-cases</strong></p>

<p>Don’t worry about having perfect code. Chances are there will be a lot of edge cases here that will be kinda tricky to handle. Just try to cover the most basic use cases first and then improve on that.</p>

<p><strong>4. A few HTML examples have been provided</strong></p>

<p>I created a few simpler HTML files and included them in this repo to help with testing. They won’t cover all potential use cases, but should help you start testing out your code.</p>

<p><strong>5. The fourth example will help you remove comments from your link text</strong></p>

<p>Chances are your first version will include the text from comments inside a link tag. Mine did. Use <a href="ex4.html" rel="nofollow">ex4.html</a> to test that case out and fix the bug.</p>

<p><em>Hint: See <a href="https://godoc.org/golang.org/x/net/html#NodeType" rel="nofollow">NodeType</a> constants and look for the types that you can ignore.</em></p>

<h3>External Resources</h3>

<p>In the solution for this exercise I end up using a DFS, which is a graph theory algorithm. If you want to learn a little more about that, I have discussed it on YouTube here - <a href="https://www.youtube.com/watch?v=zboCGDMnU3I" rel="nofollow">https://www.youtube.com/watch?v=zboCGDMnU3I</a></p>

<p>There is a complete series on algorithms and graph theory, though at this time it is somewhat incomplete. I never have enough time in the day 🙁. Hopefully one day <em>Let’s Learn Algorithms</em> will be its own series like <em>Gophercises</em>.</p>

<h4>Bonus</h4>

<p>The only bonuses here are to improve your tests and edge-case coverage.</p>

            </div>
          </div>
        

      </div>
      <div id="showSidebarButton" class="hidden">
        <button class="px-2 pb-1 border border-grey-200 rounded-r-lg text-xl text-grey-400 uppercase" onclick="showSidebar()">»</button>
      </div>


      <div class="w-full md:w-80 xl:w-96 flex-none text-md" id="sidebar">
        <div class="pl-6 pb-2 text-grey-300 uppercase tracking-wider text-sm font-bold">
          HTML Link Parser
        </div>
        <div class="mx-4 md:mr-0 border border-grey-200">
          <div class="h-2 bg-grey-300" style="width: 75%;">
          </div>
        </div>
        <div class="py-2">
          <button class="text-xs text-grey-300 px-4 uppercase" onclick="hideSidebar()">[ Click to Hide Sidebar ]</button>
        </div>
        <div class="py-2 max-h-900 overflow-y-auto">
          
    <div class="px-4 md:pr-0">
      <a href="/lessons/les_goph_16" class="flex items-center no-underline pl-2 py-1 border-r-4 hover:border-solid hover:text-grey-600 hover:border-orange-400 border-transparent text-grey-400 text-base leading-tight">
        
          
<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" class="flex-none w-6 py-1 text-green-500"><circle cx="12" cy="12" r="10" class="stroke-current fill-none"></circle><path class="fill-current" d="M10 14.59l6.3-6.3a1 1 0 0 1 1.4 1.42l-7 7a1 1 0 0 1-1.4 0l-3-3a1 1 0 0 1 1.4-1.42l2.3 2.3z"></path></svg>

        
        <div class="flex-grow py-2 px-2">
          Overview
        </div>
      </a>
    </div>
    <div class="px-4 md:pr-0">
      <a href="/lessons/les_goph_17" class="flex items-center no-underline pl-2 py-1 border-r-4 hover:border-solid hover:text-grey-600 hover:border-orange-400 border-transparent text-grey-400 text-base leading-tight">
        
          
<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" class="flex-none w-6 py-1 text-green-500"><circle cx="12" cy="12" r="10" class="stroke-current fill-none"></circle><path class="fill-current" d="M10 14.59l6.3-6.3a1 1 0 0 1 1.4 1.42l-7 7a1 1 0 0 1-1.4 0l-3-3a1 1 0 0 1 1.4-1.42l2.3 2.3z"></path></svg>

        
        <div class="flex-grow py-2 px-2">
          HTML Docs are Trees
        </div>
      </a>
    </div>

    <div class="px-4 md:pr-0">
      <a href="/lessons/les_goph_18" class="flex items-center no-underline pl-2 py-1 border-r-4 hover:border-solid hover:text-grey-600 hover:border-orange-400 border-transparent text-grey-400 text-base leading-tight">
        
          
<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" class="flex-none w-6 py-1 text-green-500"><circle cx="12" cy="12" r="10" class="stroke-current fill-none"></circle><path class="fill-current" d="M10 14.59l6.3-6.3a1 1 0 0 1 1.4 1.42l-7 7a1 1 0 0 1-1.4 0l-3-3a1 1 0 0 1 1.4-1.42l2.3 2.3z"></path></svg>

        
        <div class="flex-grow py-2 px-2">
          Defining the API
        </div>
      </a>
    </div>
    <div class="px-4 md:pr-0">
      <a href="/lessons/les_goph_19" class="flex items-center no-underline pl-2 py-1 border-r-4 hover:border-solid hover:text-grey-600 hover:border-orange-400 border-transparent text-grey-400 text-base leading-tight">
<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" class="flex-none w-6 py-1 text-green-500"><circle cx="12" cy="12" r="10" class="stroke-current fill-none"></circle><path class="fill-current" d="M10 14.59l6.3-6.3a1 1 0 0 1 1.4 1.42l-7 7a1 1 0 0 1-1.4 0l-3-3a1 1 0 0 1 1.4-1.42l2.3 2.3z"></path></svg>

        
        <div class="flex-grow py-2 px-2">
          DFSing the HTML
        </div>
      </a>
    </div>
    <div class="px-4 md:pr-0">
      <a href="/lessons/les_goph_20" class="flex items-center no-underline pl-2 py-1 border-r-4 hover:border-solid hover:text-grey-600 hover:border-orange-400 border-transparent text-grey-400 text-base leading-tight">
<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" class="flex-none w-6 py-1 text-green-500"><circle cx="12" cy="12" r="10" class="stroke-current fill-none"></circle><path class="fill-current" d="M10 14.59l6.3-6.3a1 1 0 0 1 1.4 1.42l-7 7a1 1 0 0 1-1.4 0l-3-3a1 1 0 0 1 1.4-1.42l2.3 2.3z"></path></svg>

        
        <div class="flex-grow py-2 px-2">
          Finding Link Nodes
        </div>
      </a>
    </div>
    <div class="px-4 md:pr-0">
      <a href="/lessons/les_goph_21" class="flex items-center no-underline pl-2 py-1 border-r-4 hover:border-solid hover:text-grey-600 hover:border-orange-400 border-transparent text-grey-400 text-base leading-tight">
        
          
<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" class="flex-none w-6 py-1 text-green-500"><circle cx="12" cy="12" r="10" class="stroke-current fill-none"></circle><path class="fill-current" d="M10 14.59l6.3-6.3a1 1 0 0 1 1.4 1.42l-7 7a1 1 0 0 1-1.4 0l-3-3a1 1 0 0 1 1.4-1.42l2.3 2.3z"></path></svg>

        
        <div class="flex-grow py-2 px-2">
          Building Link Structs
        </div>
      </a>
    </div>
    <div class="px-4 md:pr-0">
      <a href="/lessons/les_goph_22" class="flex items-center no-underline pl-2 py-1 border-r-4 hover:border-solid hover:text-grey-600 hover:border-orange-400 border-orange-300 text-grey-500 bg-orange-100 text-base leading-tight">
<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" class="flex-none w-6 py-1 text-green-500">
  <circle cx="12" cy="12" r="10" class="stroke-current fill-none"></circle>
</svg>
        <div class="flex-grow py-2 px-2">
          Extracting Node Text
        </div>
      </a>
    </div><div class="px-4 md:pr-0">
      <a href="/lessons/les_goph_23" class="flex items-center no-underline pl-2 py-1 border-r-4 hover:border-solid hover:text-grey-600 hover:border-orange-400  text-base leading-tight">
<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" class="flex-none w-6 py-1 text-green-500">
  <circle cx="12" cy="12" r="10" class="stroke-current fill-none"></circle>
</svg>
        <div class="flex-grow py-2 px-2">
          Parsing the Examples
        </div>
      </a>
    </div>
  <div class="py-8 px-6 text-sm text-grey-400">
    The "completed" checkmarks and the progress bar are faked for the time being until I actually start tracking video watches.
  </div>
        </div>
        <div class="w-full"><center>
      </div>
    </div>
  </div>
<script>
function hideSidebar() {
  const sidebar = document.getElementById("sidebar");
  sidebar.classList.add("hidden");
  const showButton = document.getElementById("showSidebarButton");
  showButton.classList.remove("hidden");
}
function showSidebar() {
  const sidebar = document.getElementById("sidebar");
  sidebar.classList.remove("hidden");
  const showButton = document.getElementById("showSidebarButton");
  showButton.classList.add("hidden");
}
</script>

  </body>
</html>
`

func GetHtmlTemplate(num int8) string {
	switch num {
	case 1:
		return html1
	case 2:
		return html2
	case 3:
		return html3
	case 4:
		return html4
	case 5:
		return html5
	default:
		return html1
	}
}
