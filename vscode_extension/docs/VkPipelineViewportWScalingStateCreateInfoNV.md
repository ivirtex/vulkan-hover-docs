# VkPipelineViewportWScalingStateCreateInfoNV(3) Manual Page

## Name

VkPipelineViewportWScalingStateCreateInfoNV - Structure specifying parameters of a newly created pipeline viewport W scaling state



## [](#_c_specification)C Specification

The `VkPipelineViewportWScalingStateCreateInfoNV` structure is defined as:

```c++
// Provided by VK_NV_clip_space_w_scaling
typedef struct VkPipelineViewportWScalingStateCreateInfoNV {
    VkStructureType                sType;
    const void*                    pNext;
    VkBool32                       viewportWScalingEnable;
    uint32_t                       viewportCount;
    const VkViewportWScalingNV*    pViewportWScalings;
} VkPipelineViewportWScalingStateCreateInfoNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `viewportWScalingEnable` controls whether viewport **W** scaling is enabled.
- `viewportCount` is the number of viewports used by **W** scaling, and **must** match the number of viewports in the pipeline if viewport **W** scaling is enabled.
- `pViewportWScalings` is a pointer to an array of [VkViewportWScalingNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkViewportWScalingNV.html) structures defining the **W** scaling parameters for the corresponding viewports. If the viewport **W** scaling state is dynamic, this member is ignored.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkPipelineViewportWScalingStateCreateInfoNV-sType-sType)VUID-VkPipelineViewportWScalingStateCreateInfoNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PIPELINE_VIEWPORT_W_SCALING_STATE_CREATE_INFO_NV`
- [](#VUID-VkPipelineViewportWScalingStateCreateInfoNV-viewportCount-arraylength)VUID-VkPipelineViewportWScalingStateCreateInfoNV-viewportCount-arraylength  
  `viewportCount` **must** be greater than `0`

## [](#_see_also)See Also

[VK\_NV\_clip\_space\_w\_scaling](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_clip_space_w_scaling.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkViewportWScalingNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkViewportWScalingNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineViewportWScalingStateCreateInfoNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0