from springboot2beat import BaseTest

import os


class Test(BaseTest):

    def test_base(self):
        """
        Basic test with exiting Springboot2beat normally
        """
        self.render_config_template(
            path=os.path.abspath(self.working_dir) + "/log/*"
        )

        springboot2beat_proc = self.start_beat()
        self.wait_until(lambda: self.log_contains("springboot2beat is running"))
        exit_code = springboot2beat_proc.kill_and_wait()
        assert exit_code == 0
