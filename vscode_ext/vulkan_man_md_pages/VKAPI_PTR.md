# VKAPI_PTR(3) Manual Page

## Name

VKAPI_PTR - Vulkan function pointer calling conventions macro



## <a href="#_description" class="anchor"></a>Description

`VKAPI_PTR` is a macro placed between the '(' and '\*' in Vulkan API
function pointer declarations. This macro also controls calling
conventions, and typically has the same definition as `VKAPI_ATTR` or
`VKAPI_CALL`, depending on the compiler.

## <a href="#_see_also" class="anchor"></a>See Also

[VKAPI_ATTR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VKAPI_ATTR.html), [VKAPI_CALL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VKAPI_CALL.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#boilerplate-platform-specific-calling-conventions"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
