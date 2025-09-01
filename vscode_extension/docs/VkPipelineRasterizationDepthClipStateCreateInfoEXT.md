# VkPipelineRasterizationDepthClipStateCreateInfoEXT(3) Manual Page

## Name

VkPipelineRasterizationDepthClipStateCreateInfoEXT - Structure specifying depth clipping state



## [](#_c_specification)C Specification

If the `pNext` chain of [VkPipelineRasterizationStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRasterizationStateCreateInfo.html) includes a `VkPipelineRasterizationDepthClipStateCreateInfoEXT` structure, then that structure controls whether depth clipping is enabled or disabled.

The `VkPipelineRasterizationDepthClipStateCreateInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_depth_clip_enable
typedef struct VkPipelineRasterizationDepthClipStateCreateInfoEXT {
    VkStructureType                                        sType;
    const void*                                            pNext;
    VkPipelineRasterizationDepthClipStateCreateFlagsEXT    flags;
    VkBool32                                               depthClipEnable;
} VkPipelineRasterizationDepthClipStateCreateInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is reserved for future use.
- `depthClipEnable` controls whether depth clipping is enabled as described in [Primitive Clipping](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vertexpostproc-clipping).

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkPipelineRasterizationDepthClipStateCreateInfoEXT-sType-sType)VUID-VkPipelineRasterizationDepthClipStateCreateInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PIPELINE_RASTERIZATION_DEPTH_CLIP_STATE_CREATE_INFO_EXT`
- [](#VUID-VkPipelineRasterizationDepthClipStateCreateInfoEXT-flags-zerobitmask)VUID-VkPipelineRasterizationDepthClipStateCreateInfoEXT-flags-zerobitmask  
  `flags` **must** be `0`

## [](#_see_also)See Also

[VK\_EXT\_depth\_clip\_enable](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_depth_clip_enable.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkPipelineRasterizationDepthClipStateCreateFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRasterizationDepthClipStateCreateFlagsEXT.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineRasterizationDepthClipStateCreateInfoEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0