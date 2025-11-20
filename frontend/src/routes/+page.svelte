<script lang="ts">
  import { fly } from 'svelte/transition';
  import { cubicOut } from 'svelte/easing';
  import { onMount } from 'svelte';

  let show = false;

  onMount(() => {
    show = true;
  });

  function enterSite() {
    // TODO: replace with your real navigation / state
    // e.g. use goto('/') or set some `hasAccepted = true`
    console.log('User accepted disclaimer');
    show = false;
  }

  function leaveSite() {
    // TODO: point this somewhere else if you want
    window.location.href = 'https://www.google.com';
  }
</script>

<main class="min-h-screen flex items-center justify-center px-6">
  {#if show}
    <section
      role="dialog"
      aria-modal="true"
      in:fly={{ y: 50, duration: 1000, easing: cubicOut }}
      class="w-full max-w-3xl rounded-2xl bg-zinc-900/95 border border-white/5 shadow-2xl px-10 py-12 text-center backdrop-blur"
    >
      <h1 class="mb-3 text-4xl sm:text-5xl font-medium">
        <span class="text-white">echo</span><span class="text-brand-500">2k</span>
      </h1>

      <!-- Legal disclaimer scroll box -->
      <div
        class="mt-6 text-left max-h-48 overflow-y-auto rounded-xl border border-white/10 bg-black/20 px-4 py-3 text-xs sm:text-sm text-zinc-300 space-y-3"
      >
        <p class="font-semibold text-zinc-100">
		Copyright Notice &amp; Important Legal Disclaimer
		</p>

		<p class="mt-2">
		This website is an independent, student-run platform. It is not affiliated with, endorsed by, or operated on behalf of the University of Warwick or Echo360 / EchoVideo.
		</p>

		<p class="mt-4 font-semibold text-zinc-100">
		Intellectual Property &amp; Ownership
		</p>

		<p>
		All lecture recordings, slides, transcripts, and teaching materials accessible through this site are the intellectual property of their respective rights holders. In most cases, this includes the University of Warwick and/or individual lecturers. No ownership or copyright is claimed over any such material.
		</p>

		<p class="mt-4 font-semibold text-zinc-100">
		Copyright Law (UK)
		</p>

		<p>
		All materials made available through this site may be protected under the Copyright, Designs and Patents Act 1988 (CDPA). Copyright law restricts the copying, distribution, public display, and communication of protected works without permission from the rights holder.
		</p>

		<p class="mt-2">
		Users are reminded that:
		</p>

		<ul class="list-disc list-inside space-y-1">
		<li>Section 16 CDPA gives copyright owners exclusive rights to copy and distribute their works.</li>
		<li>Section 17 defines “copying” to include digital reproduction.</li>
		<li>Section 18 prohibits unauthorised distribution.</li>
		<li>Section 20 restricts communication to the public without permission.</li>
		<li>Section 29 permits limited copying for private study, but this does NOT include uploading, sharing, or redistributing copies to others.</li>
		<li>Section 32 allows “illustration for instruction” only under strict conditions, and does NOT permit broad redistribution of recordings.</li>
		</ul>

		<p class="mt-4">
		By continuing, you agree that:
		</p>

		<ul class="list-disc list-inside space-y-1">
		<li>You are solely responsible for ensuring you have the legal right to access any material you view through this site.</li>
		<li>You will not copy, download, redistribute, publish, or share any material unless explicitly permitted by law or by the rights holder.</li>
		<li>You acknowledge that unauthorised distribution of copyright-protected material may expose you to legal or disciplinary consequences.</li>
		</ul>

		<p class="mt-4 font-semibold text-zinc-100">
		Role of This Site
		</p>

		<p>
		This site functions only as a technical platform. Responsibility for lawful use of any material rests entirely with the user. The operator of this site does not verify user permissions and does not authorise or encourage copyright infringement of any kind.
		</p>

		<p class="mt-4 font-semibold text-zinc-100">
		Takedown Contact
		</p>

		<p>
		If you are a rights holder and believe any material accessible through this site infringes your rights, please contact the webmaster at:
		<a href="mailto:echo2k_webmaster@protonmail.com" class="underline">
			echo2k_webmaster@protonmail.com
		</a>.
		We will investigate and, where appropriate, remove or restrict access to the material.
		</p>

      </div>

      <!-- Buttons -->
      <div class="mt-8 flex flex-col sm:flex-row gap-3">
        <button
          type="button"
          on:click={enterSite}
          class="flex-1 rounded-full bg-brand-500 text-white text-sm font-semibold py-3 px-4 tracking-wide
                 hover:bg-brand-400 focus:outline-none focus:ring-2 focus:ring-brand-500/70
                 transition-colors"
        >
          I have read and agree – enter site
        </button>

        <button
          type="button"
          on:click={leaveSite}
          class="flex-1 rounded-full border border-zinc-700 bg-zinc-900/80 text-zinc-200 text-sm font-medium py-3 px-4
                 hover:bg-zinc-800 focus:outline-none focus:ring-2 focus:ring-zinc-600/60
                 transition-colors"
        >
          Leave site
        </button>
      </div>
    </section>
  {/if}
</main>
