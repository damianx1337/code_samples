import signal
import time

class GracefulTermination:
	kill_now = False
	def __init__(self):
		signal.signal(signal.SIGINT, self.exit_gracefully)
		signal.signal(signal.SIGTERM, self.exit_gracefully)

	def exit_gracefully(self, signum, frame):
		self.kill_now = True

if __name__ == '__main__':
	terminator = GracefulTermination()
	while not terminator.kill_now:
		time.sleep(1)
		print("doing something...")
	print("Shutting down...")
