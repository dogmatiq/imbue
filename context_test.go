package imbue_test

import (
	"context"
	"errors"

	"github.com/dogmatiq/imbue"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("type Context", func() {
	var container *imbue.Container

	BeforeEach(func() {
		container = imbue.New()
	})

	AfterEach(func() {
		container.Close()
	})

	Describe("func Defer()", func() {
		When("construction and decoration succeeds", func() {
			It("defers calling functions deferred by constructors until the container is closed", func() {
				called := false
				imbue.With0(
					container,
					func(
						ctx *imbue.Context,
					) (Concrete1, error) {
						ctx.Defer(func() error {
							called = true
							return nil
						})
						return "<concrete>", nil
					},
				)

				err := imbue.Invoke1(
					context.Background(),
					container,
					func(
						context.Context,
						Concrete1,
					) error {
						return nil
					},
				)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(called).To(BeFalse())

				err = container.Close()
				Expect(err).ShouldNot(HaveOccurred())
				Expect(called).To(BeTrue())
			})

			It("defers calling functions deferred by decorators until the container is closed", func() {
				imbue.With0(
					container,
					func(
						ctx *imbue.Context,
					) (Concrete1, error) {
						return "<concrete>", nil
					},
				)

				called := false
				imbue.Decorate0(
					container,
					func(
						ctx *imbue.Context,
						dep Concrete1,
					) (Concrete1, error) {
						ctx.Defer(func() error {
							called = true
							return nil
						})
						return dep, nil
					},
				)

				err := imbue.Invoke1(
					context.Background(),
					container,
					func(
						context.Context,
						Concrete1,
					) error {
						return nil
					},
				)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(called).To(BeFalse())

				err = container.Close()
				Expect(err).ShouldNot(HaveOccurred())
				Expect(called).To(BeTrue())
			})

			It("calls deferred functions in reverse order", func() {
				var order []string

				imbue.With0(
					container,
					func(
						ctx *imbue.Context,
					) (Concrete1, error) {
						ctx.Defer(func() error {
							order = append(order, "<concrete-1-constructor>")
							return nil
						})
						return "<concrete-1>", nil
					},
				)

				imbue.Decorate0(
					container,
					func(
						ctx *imbue.Context,
						dep Concrete1,
					) (Concrete1, error) {
						ctx.Defer(func() error {
							order = append(order, "<concrete-1-decorator>")
							return nil
						})
						return dep, nil
					},
				)

				imbue.With1(
					container,
					func(
						ctx *imbue.Context,
						_ Concrete1,
					) (Concrete2, error) {
						ctx.Defer(func() error {
							order = append(order, "<concrete-2-constructor>")
							return nil
						})
						return "<concrete-2>", nil
					},
				)

				err := imbue.Invoke1(
					context.Background(),
					container,
					func(
						context.Context,
						Concrete2,
					) error {
						return nil
					},
				)
				Expect(err).ShouldNot(HaveOccurred())

				err = container.Close()
				Expect(err).ShouldNot(HaveOccurred())
				Expect(order).To(Equal([]string{
					"<concrete-2-constructor>",
					"<concrete-1-decorator>",
					"<concrete-1-constructor>",
				}))
			})

			It("produces an error when one or more deferred functions fails", func() {
				imbue.With0(
					container,
					func(
						ctx *imbue.Context,
					) (Concrete1, error) {
						ctx.Defer(func() error {
							return errors.New("<concrete-1-constructor>")
						})
						return "<concrete-1>", nil
					},
				)

				imbue.Decorate0(
					container,
					func(
						ctx *imbue.Context,
						dep Concrete1,
					) (Concrete1, error) {
						ctx.Defer(func() error {
							return errors.New("<concrete-1-decorator>")
						})
						return dep, nil
					},
				)

				imbue.With1(
					container,
					func(
						ctx *imbue.Context,
						_ Concrete1,
					) (Concrete2, error) {
						ctx.Defer(func() error {
							return errors.New("<concrete-2-constructor>")
						})
						return "<concrete-2>", nil
					},
				)

				err := imbue.Invoke1(
					context.Background(),
					container,
					func(
						context.Context,
						Concrete2,
					) error {
						return nil
					},
				)
				Expect(err).ShouldNot(HaveOccurred())

				err = container.Close()
				Expect(err).To(
					MatchError(
						MatchRegexp(
							`3 error\(s\) occurred while closing the container:`+
								`\n\t1\) function deferred at context_test\.go:\d+ by imbue_test\.Concrete2 constructor \(context_test\.go:\d+\) failed: <concrete-2-constructor>`+
								`\n\t2\) function deferred at context_test\.go:\d+ by imbue_test\.Concrete1 decorator \(context_test\.go:\d+\) failed: <concrete-1-decorator>`+
								`\n\t3\) function deferred at context_test\.go:\d+ by imbue_test\.Concrete1 constructor \(context_test\.go:\d+\) failed: <concrete-1-constructor>`,
						),
					),
					err.Error(),
				)
			})
		})

		It("guarantees all deferred functions are called when there is a panic", func() {
			count := 0

			imbue.With0(
				container,
				func(
					ctx *imbue.Context,
				) (Concrete1, error) {
					ctx.Defer(func() error {
						count++
						return nil
					})
					return "<concrete-1>", nil
				},
			)

			imbue.Decorate0(
				container,
				func(
					ctx *imbue.Context,
					dep Concrete1,
				) (Concrete1, error) {
					ctx.Defer(func() error {
						count++
						panic("<panic>")
					})
					return dep, nil
				},
			)

			imbue.With1(
				container,
				func(
					ctx *imbue.Context,
					_ Concrete1,
				) (Concrete2, error) {
					ctx.Defer(func() error {
						count++
						return nil
					})
					return "<concrete-2>", nil
				},
			)

			err := imbue.Invoke1(
				context.Background(),
				container,
				func(
					context.Context,
					Concrete2,
				) error {
					return nil
				},
			)
			Expect(err).ShouldNot(HaveOccurred())

			Expect(func() {
				container.Close()
			}).To(PanicWith("<panic>"))

			Expect(count).To(BeNumerically("==", 3))
		})

		It("does not call deferred functions again when the container is closed a second time", func() {
			closingAgain := false
			imbue.With0(
				container,
				func(
					ctx *imbue.Context,
				) (Concrete1, error) {
					ctx.Defer(func() error {
						if closingAgain {
							Fail("unexpected call")
						}
						return nil
					})
					return "<concrete>", nil
				},
			)

			imbue.Decorate0(
				container,
				func(
					ctx *imbue.Context,
					dep Concrete1,
				) (Concrete1, error) {
					ctx.Defer(func() error {
						if closingAgain {
							Fail("unexpected call")
						}
						return nil
					})
					return "", nil
				},
			)

			err := imbue.Invoke1(
				context.Background(),
				container,
				func(
					context.Context,
					Concrete1,
				) error {
					return nil
				},
			)
			Expect(err).ShouldNot(HaveOccurred())

			err = container.Close()
			Expect(err).ShouldNot(HaveOccurred())

			closingAgain = true
			err = container.Close()
			Expect(err).ShouldNot(HaveOccurred())
		})
	})

	When("construction and/or decoration fails", func() {
		It("calls functions deferred by constructors immediately", func() {
			called := false
			imbue.With0(
				container,
				func(
					ctx *imbue.Context,
				) (Concrete1, error) {
					ctx.Defer(func() error {
						called = true
						return nil
					})
					return "", errors.New("<error>")
				},
			)

			err := imbue.Invoke1(
				context.Background(),
				container,
				func(
					context.Context,
					Concrete1,
				) error {
					panic("unexpected call")
				},
			)
			Expect(err).To(
				MatchError(
					MatchRegexp(
						`imbue_test\.Concrete1 constructor \(context_test\.go:\d+\) failed: <error>`,
					),
				),
			)
			Expect(called).To(BeTrue())
		})

		It("calls functions deferred by decorators immediately", func() {
			imbue.With0(
				container,
				func(
					ctx *imbue.Context,
				) (Concrete1, error) {
					return "<concrete>", nil
				},
			)

			called := false
			imbue.Decorate0(
				container,
				func(
					ctx *imbue.Context,
					dep Concrete1,
				) (Concrete1, error) {
					ctx.Defer(func() error {
						called = true
						return nil
					})
					return "", errors.New("<error>")
				},
			)

			err := imbue.Invoke1(
				context.Background(),
				container,
				func(
					context.Context,
					Concrete1,
				) error {
					panic("unexpected call")
				},
			)
			Expect(err).To(
				MatchError(
					MatchRegexp(
						`imbue_test\.Concrete1 decorator \(context_test\.go:\d+\) failed: <error>`,
					),
				),
			)
			Expect(called).To(BeTrue())
		})

		It("calls deferred functions in reverse order", func() {
			var order []string

			imbue.With0(
				container,
				func(
					ctx *imbue.Context,
				) (Concrete1, error) {
					ctx.Defer(func() error {
						order = append(order, "<concrete-1-constructor>")
						return nil
					})
					return "<concrete-1>", nil
				},
			)

			imbue.Decorate0(
				container,
				func(
					ctx *imbue.Context,
					dep Concrete1,
				) (Concrete1, error) {
					ctx.Defer(func() error {
						order = append(order, "<concrete-1-decorator-1>")
						return nil
					})
					return "", nil
				},
			)

			imbue.Decorate0(
				container,
				func(
					ctx *imbue.Context,
					dep Concrete1,
				) (Concrete1, error) {
					ctx.Defer(func() error {
						order = append(order, "<concrete-1-decorator-2>")
						return nil
					})
					return "", errors.New("<error>")
				},
			)

			err := imbue.Invoke1(
				context.Background(),
				container,
				func(
					context.Context,
					Concrete1,
				) error {
					panic("unexpected call")
				},
			)
			Expect(err).Should(HaveOccurred())
			Expect(order).To(Equal([]string{
				"<concrete-1-decorator-2>",
				"<concrete-1-decorator-1>",
				"<concrete-1-constructor>",
			}))
		})

		It("guarantees all deferred functions are called when there is a panic", func() {
			count := 0

			imbue.With0(
				container,
				func(
					ctx *imbue.Context,
				) (Concrete1, error) {
					ctx.Defer(func() error {
						count++
						return nil
					})
					return "<concrete-1>", nil
				},
			)

			imbue.Decorate0(
				container,
				func(
					ctx *imbue.Context,
					dep Concrete1,
				) (Concrete1, error) {
					ctx.Defer(func() error {
						count++
						panic("<panic>")
					})
					return dep, errors.New("<error>")
				},
			)

			Expect(func() {
				imbue.Invoke1(
					context.Background(),
					container,
					func(
						context.Context,
						Concrete1,
					) error {
						panic("unexpected call")
					},
				)
			}).To(PanicWith("<panic>"))

			Expect(count).To(BeNumerically("==", 2))
		})

		It("does not call deferred functions again when the container is closed", func() {
			closing := false
			imbue.With0(
				container,
				func(
					ctx *imbue.Context,
				) (Concrete1, error) {
					ctx.Defer(func() error {
						if closing {
							Fail("unexpected call")
						}
						return nil
					})
					return "<concrete>", nil
				},
			)

			imbue.Decorate0(
				container,
				func(
					ctx *imbue.Context,
					dep Concrete1,
				) (Concrete1, error) {
					ctx.Defer(func() error {
						if closing {
							Fail("unexpected call")
						}
						return nil
					})
					return "", errors.New("<error>")
				},
			)

			err := imbue.Invoke1(
				context.Background(),
				container,
				func(
					context.Context,
					Concrete1,
				) error {
					return nil
				},
			)
			Expect(err).Should(HaveOccurred())

			closing = true
			err = container.Close()
			Expect(err).ShouldNot(HaveOccurred())
		})
	})
})
