import AnimateInView from "@/components/ui/AnimateInView";

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
          <div className="bg-gray-100 dark:bg-gray-800 rounded-2xl p-6 md:p-8 shadow-md">
            <div className="flex flex-col md:flex-row items-center">
              <div className="mb-6 md:mb-0 md:mr-8">
                <div className="bg-primary/10 dark:bg-primary/20 p-3 rounded-full w-16 h-16 flex items-center justify-center mb-4 mx-auto md:mx-0">
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    className="h-8 w-8 text-primary"
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
              <div className="text-center md:text-left md:flex-1">
                <h3 className="text-xl font-bold mb-2">
                  {title}
                </h3>
                <p className="text-gray-600 dark:text-gray-300 mb-6">
                  {description}
                </p>
                <div className="flex flex-col sm:flex-row gap-3">
                  <input
                    type="email"
                    placeholder={placeholderText}
                    className="px-4 py-3 rounded-lg border-0 shadow-sm flex-1"
                    aria-label="メールアドレス入力欄"
                  />
                  <button 
                    className="btn-primary text-white py-3 px-6 rounded-lg shadow-sm font-medium"
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
