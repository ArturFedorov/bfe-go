package logger_rate_limiter

import "testing"

func TestLogger(t *testing.T) {
	t.Run("sequential messages within 10s blocked", func(t *testing.T) {
		l := Constructor()
		if got := l.ShouldPrintMessage(1, "foo"); !got {
			t.Errorf("timestamp=1, 'foo': got %v, want true", got)
		}
		if got := l.ShouldPrintMessage(2, "foo"); got {
			t.Errorf("timestamp=2, 'foo': got %v, want false", got)
		}
		if got := l.ShouldPrintMessage(10, "foo"); got {
			t.Errorf("timestamp=10, 'foo': got %v, want false", got)
		}
	})

	t.Run("after 10s allowed", func(t *testing.T) {
		l := Constructor()
		if got := l.ShouldPrintMessage(1, "foo"); !got {
			t.Errorf("timestamp=1, 'foo': got %v, want true", got)
		}
		if got := l.ShouldPrintMessage(11, "foo"); !got {
			t.Errorf("timestamp=11, 'foo': got %v, want true", got)
		}
	})

	t.Run("different messages independent", func(t *testing.T) {
		l := Constructor()
		if got := l.ShouldPrintMessage(1, "foo"); !got {
			t.Errorf("timestamp=1, 'foo': got %v, want true", got)
		}
		if got := l.ShouldPrintMessage(1, "bar"); !got {
			t.Errorf("timestamp=1, 'bar': got %v, want true", got)
		}
		if got := l.ShouldPrintMessage(2, "foo"); got {
			t.Errorf("timestamp=2, 'foo': got %v, want false", got)
		}
		if got := l.ShouldPrintMessage(2, "bar"); got {
			t.Errorf("timestamp=2, 'bar': got %v, want false", got)
		}
	})
}
