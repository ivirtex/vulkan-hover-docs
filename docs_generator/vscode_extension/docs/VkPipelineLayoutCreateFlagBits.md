# VkPipelineLayoutCreateFlagBits(3) Manual Page

## Name

VkPipelineLayoutCreateFlagBits - Pipeline layout creation flag bits



## [](#_c_specification)C Specification

```c++
// Provided by VK_EXT_graphics_pipeline_library
typedef enum VkPipelineLayoutCreateFlagBits {
  // Provided by VK_EXT_graphics_pipeline_library
    VK_PIPELINE_LAYOUT_CREATE_INDEPENDENT_SETS_BIT_EXT = 0x00000002,
} VkPipelineLayoutCreateFlagBits;
```

## [](#_description)Description

- `VK_PIPELINE_LAYOUT_CREATE_INDEPENDENT_SETS_BIT_EXT` specifies that implementations **must** ensure that the properties and/or absence of a particular descriptor set do not influence any other properties of the pipeline layout. This allows pipelines libraries linked without `VK_PIPELINE_CREATE_LINK_TIME_OPTIMIZATION_BIT_EXT` to be created with a subset of the total descriptor sets.

## [](#_see_also)See Also

[VK\_EXT\_graphics\_pipeline\_library](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_graphics_pipeline_library.html), [VkPipelineLayoutCreateFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayoutCreateFlags.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineLayoutCreateFlagBits)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0