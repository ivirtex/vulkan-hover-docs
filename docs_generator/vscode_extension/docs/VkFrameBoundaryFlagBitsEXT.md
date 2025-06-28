# VkFrameBoundaryFlagBitsEXT(3) Manual Page

## Name

VkFrameBoundaryFlagBitsEXT - Bitmask specifying whether a queue submission is the last one for a given frame



## [](#_c_specification)C Specification

The bit which **can** be set in [VkFrameBoundaryEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFrameBoundaryEXT.html)::`flags` is:

```c++
// Provided by VK_EXT_frame_boundary
typedef enum VkFrameBoundaryFlagBitsEXT {
    VK_FRAME_BOUNDARY_FRAME_END_BIT_EXT = 0x00000001,
} VkFrameBoundaryFlagBitsEXT;
```

## [](#_description)Description

- `VK_FRAME_BOUNDARY_FRAME_END_BIT_EXT` specifies that this queue submission is the last one for this frame, i.e. once this queue submission has terminated, then the work for this frame is completed.

## [](#_see_also)See Also

[VK\_EXT\_frame\_boundary](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_frame_boundary.html), [VkFrameBoundaryFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFrameBoundaryFlagsEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkFrameBoundaryFlagBitsEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0