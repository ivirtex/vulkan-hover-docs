# VkViewportWScalingNV(3) Manual Page

## Name

VkViewportWScalingNV - Structure specifying a viewport



## [](#_c_specification)C Specification

The `VkViewportWScalingNV` structure is defined as:

```c++
// Provided by VK_NV_clip_space_w_scaling
typedef struct VkViewportWScalingNV {
    float    xcoeff;
    float    ycoeff;
} VkViewportWScalingNV;
```

## [](#_members)Members

- `xcoeff` and `ycoeff` are the viewport’s W scaling factor for x and y respectively.

## [](#_description)Description

## [](#_see_also)See Also

[VK\_NV\_clip\_space\_w\_scaling](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_clip_space_w_scaling.html), [VkPipelineViewportWScalingStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineViewportWScalingStateCreateInfoNV.html), [vkCmdSetViewportWScalingNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetViewportWScalingNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkViewportWScalingNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0