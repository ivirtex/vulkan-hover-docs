# VkDisplayEventTypeEXT(3) Manual Page

## Name

VkDisplayEventTypeEXT - Events that can occur on a display object



## <a href="#_c_specification" class="anchor"></a>C Specification

Possible values of
[VkDisplayEventInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayEventInfoEXT.html)::`displayEvent`,
specifying when a fence will be signaled, are:

``` c
// Provided by VK_EXT_display_control
typedef enum VkDisplayEventTypeEXT {
    VK_DISPLAY_EVENT_TYPE_FIRST_PIXEL_OUT_EXT = 0,
} VkDisplayEventTypeEXT;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_DISPLAY_EVENT_TYPE_FIRST_PIXEL_OUT_EXT` specifies that the fence
  is signaled when the first pixel of the next display refresh cycle
  leaves the display engine for the display.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_display_control](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_display_control.html),
[VkDisplayEventInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayEventInfoEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDisplayEventTypeEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
