# CoreMaxIDARM(3) Manual Page

## Name

CoreMaxIDARM - Max core ID that can be observed on the device running
the invovation reading CoreMaxIDARM



## <a href="#_description" class="anchor"></a>Description

`CoreMaxIDARM`  
Decorating a variable with the `CoreMaxIDARM` built-in decoration will
make that variable contain the max ID of any shader core on the device
on which the current shader invocation is running.

Valid Usage

- <a href="#VUID-CoreMaxIDARM-CoreMaxIDARM-07597"
  id="VUID-CoreMaxIDARM-CoreMaxIDARM-07597"></a>
  VUID-CoreMaxIDARM-CoreMaxIDARM-07597  
  The variable decorated with `CoreMaxIDARM` **must** be declared using
  the `Input` `Storage` `Class`

- <a href="#VUID-CoreMaxIDARM-CoreMaxIDARM-07598"
  id="VUID-CoreMaxIDARM-CoreMaxIDARM-07598"></a>
  VUID-CoreMaxIDARM-CoreMaxIDARM-07598  
  The variable decorated with `CoreMaxIDARM` **must** be declared as a
  scalar 32-bit integer value

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#CoreMaxIDARM"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
