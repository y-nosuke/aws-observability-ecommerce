"use client";

import { ReactNode, useEffect, useRef, useState } from "react";

type AnimationDirection =
  | "up"
  | "down"
  | "left"
  | "right"
  | "scale"
  | "opacity";

interface AnimateInViewProps {
  children: ReactNode;
  direction?: AnimationDirection;
  delay?: number;
  duration?: number;
  once?: boolean;
  className?: string;
}

export default function AnimateInView({
  children,
  direction = "up",
  delay = 0,
  duration = 500,
  once = true,
  className = "",
}: AnimateInViewProps) {
  const [isVisible, setIsVisible] = useState(false);
  const [hasAnimated, setHasAnimated] = useState(false);
  const ref = useRef<HTMLDivElement>(null);

  useEffect(() => {
    const currentRef = ref.current;
    if (!currentRef) return;

    const observer = new IntersectionObserver(
      ([entry]) => {
        if (entry.isIntersecting) {
          setIsVisible(true);
          if (once) {
            setHasAnimated(true);
          }
        } else if (!once && !hasAnimated) {
          setIsVisible(false);
        }
      },
      {
        threshold: 0.1,
      }
    );

    observer.observe(currentRef);

    return () => {
      if (currentRef) {
        observer.unobserve(currentRef);
      }
    };
  }, [once, hasAnimated]);

  // 方向に基づいて初期スタイルと最終スタイルを決定
  const getAnimationStyles = () => {
    const baseStyles = {
      opacity: 0,
      transform: "translateY(0) translateX(0) scale(1)",
    };

    switch (direction) {
      case "up":
        baseStyles.transform = "translateY(30px)";
        break;
      case "down":
        baseStyles.transform = "translateY(-30px)";
        break;
      case "left":
        baseStyles.transform = "translateX(30px)";
        break;
      case "right":
        baseStyles.transform = "translateX(-30px)";
        break;
      case "scale":
        baseStyles.transform = "scale(0.9)";
        break;
      case "opacity":
        // opacity のみを変更
        break;
    }

    return baseStyles;
  };

  const animationStyles = isVisible
    ? {
        opacity: 1,
        transform: "translateY(0) translateX(0) scale(1)",
        transition: `opacity ${duration}ms, transform ${duration}ms ease-out`,
        transitionDelay: `${delay}ms`,
      }
    : {
        ...getAnimationStyles(),
        transition: `opacity ${duration}ms, transform ${duration}ms ease-out`,
        transitionDelay: "0ms",
      };

  return (
    <div ref={ref} className={className} style={animationStyles}>
      {children}
    </div>
  );
}
