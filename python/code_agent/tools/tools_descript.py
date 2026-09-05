from tools.database import database_tool_descript
from tools.file_system import file_system_tool_descript
from tools.test_executors import test_tool_descript


my_tools = database_tool_descript.database_tools + file_system_tool_descript.file_system_tools + test_tool_descript.test_tools
available_tools = database_tool_descript.database_available_tools | file_system_tool_descript.file_system_available_tools | test_tool_descript.test_available_tools

tools_with_question = ["write_file", "remove_file", "remove_folder", "update_memory"]
