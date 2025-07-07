import AnimateInView from '@/components/ui/AnimateInView';

interface NewsletterSignupProps {
  title: string;
  description: string;
  buttonText: string;
  placeholderText: string;
}

export default function NewsletterSignup({
  title,
  description,
  buttonText,
  placeholderText,
}: NewsletterSignupProps) {
  return (
    <section className="mb-16">
      <AnimateInView direction="up" delay={300}>
        <div className="container mx-auto px-6">
          <div className="rounded-2xl bg-gray-100 p-6 shadow-md md:p-8 dark:bg-gray-800">
            <div className="flex flex-col items-center md:flex-row">
              <div className="mb-6 md:mr-8 md:mb-0">
                <div className="bg-primary/10 dark:bg-primary/20 mx-auto mb-4 flex h-16 w-16 items-center justify-center rounded-full p-3 md:mx-0">
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    className="text-primary h-8 w-8"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke="currentColor"
                  >
                    <path
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      strokeWidth={2}
                      d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"
                    />
                  </svg>
                </div>
              </div>
              <div className="text-center md:flex-1 md:text-left">
                <h3 className="mb-2 text-xl font-bold">{title}</h3>
                <p className="mb-6 text-gray-600 dark:text-gray-300">{description}</p>
                <div className="flex flex-col gap-3 sm:flex-row">
                  <input
                    type="email"
                    placeholder={placeholderText}
                    className="flex-1 rounded-lg border-0 px-4 py-3 shadow-sm"
                    aria-label="メールアドレス入力欄"
                  />
                  <button
                    className="btn-primary rounded-lg px-6 py-3 font-medium text-white shadow-sm"
                    aria-label="メールマガジン登録ボタン"
                  >
                    {buttonText}
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </AnimateInView>
    </section>
  );
}
