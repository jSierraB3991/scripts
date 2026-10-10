
from mcp.server.fastmcp import FastMCP

from tools.file_listing import list_files
from tools.largest_files import find_largest_files

mcp = FastMCP("files-server")

mcp.tool()(list_files)
mcp.tool()(find_largest_files)


if __name__ == "__main__":
    mcp.run(transport="stdio")
