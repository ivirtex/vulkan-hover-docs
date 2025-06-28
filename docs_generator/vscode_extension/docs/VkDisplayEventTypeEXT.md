# VkDisplayEventTypeEXT(3) Manual Page

## Name

VkDisplayEventTypeEXT - Events that can occur on a display object



## [](#_c_specification)C Specification

Possible values of [VkDisplayEventInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayEventInfoEXT.html)::`displayEvent`, specifying when a fence will be signaled, are:

```c++
// Provided by VK_EXT_display_control
typedef enum VkDisplayEventTypeEXT {
    VK_DISPLAY_EVENT_TYPE_FIRST_PIXEL_OUT_EXT = 0,
} VkDisplayEventTypeEXT;
```

## [](#_description)Description

- `VK_DISPLAY_EVENT_TYPE_FIRST_PIXEL_OUT_EXT` specifies that the fence is signaled when the first pixel of the next display refresh cycle leaves the display engine for the display.

## [](#_see_also)See Also

[VK\_EXT\_display\_control](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_display_control.html), [VkDisplayEventInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayEventInfoEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDisplayEventTypeEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0