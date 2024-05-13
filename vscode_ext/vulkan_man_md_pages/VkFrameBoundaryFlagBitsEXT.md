# VkFrameBoundaryFlagBitsEXT(3) Manual Page

## Name

VkFrameBoundaryFlagBitsEXT - Bitmask specifying whether a queue
submission is the last one for a given frame



## <a href="#_c_specification" class="anchor"></a>C Specification

The bit which **can** be set in
[VkFrameBoundaryEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFrameBoundaryEXT.html)::`flags` is:

``` c
// Provided by VK_EXT_frame_boundary
typedef enum VkFrameBoundaryFlagBitsEXT {
    VK_FRAME_BOUNDARY_FRAME_END_BIT_EXT = 0x00000001,
} VkFrameBoundaryFlagBitsEXT;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_FRAME_BOUNDARY_FRAME_END_BIT_EXT` specifies that this queue
  submission is the last one for this frame, i.e. once this queue
  submission has terminated, then the work for this frame is completed.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_frame_boundary](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_frame_boundary.html),
[VkFrameBoundaryFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFrameBoundaryFlagsEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkFrameBoundaryFlagBitsEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
