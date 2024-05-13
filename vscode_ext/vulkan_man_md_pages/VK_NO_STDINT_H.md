# VK_NO_STDINT_H(3) Manual Page

## Name

VK_NO_STDINT_H - Control definition of types



## <a href="#_description" class="anchor"></a>Description

If the `VK_NO_STDINT_H` macro is defined by the application at compile
time, extended integer types used by the Vulkan API, such as `uint8_t`,
**must** also be defined by the application. Otherwise, the Vulkan
headers will not compile. If `VK_NO_STDINT_H` is not defined, the system
`<stdint.h>` is used to define these types. There is a fallback path
when Microsoft Visual Studio version 2008 and earlier versions are
detected at compile time.

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#boilerplate-platform-specific-header-control"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
