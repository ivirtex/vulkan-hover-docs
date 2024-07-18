# WarpMaxIDARM(3) Manual Page

## Name

WarpMaxIDARM - Max ID for a warp on the core running a shader invovation



## <a href="#_description" class="anchor"></a>Description

`WarpMaxIDARM`  
Decorating a variable with the `WarpMaxIDARM` built-in decoration will
make that variable contain the maximum warp ID for the core on which the
current invocation is running.

Valid Usage

- <a href="#VUID-WarpMaxIDARM-WarpMaxIDARM-07601"
  id="VUID-WarpMaxIDARM-WarpMaxIDARM-07601"></a>
  VUID-WarpMaxIDARM-WarpMaxIDARM-07601  
  The variable decorated with `WarpMaxIDARM` **must** be declared using
  the `Input` `Storage` `Class`

- <a href="#VUID-WarpMaxIDARM-WarpMaxIDARM-07602"
  id="VUID-WarpMaxIDARM-WarpMaxIDARM-07602"></a>
  VUID-WarpMaxIDARM-WarpMaxIDARM-07602  
  The variable decorated with `WarpMaxIDARM` **must** be declared as a
  scalar 32-bit integer value

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#WarpMaxIDARM"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
