# VkPipelineLayoutCreateFlagBits(3) Manual Page

## Name

VkPipelineLayoutCreateFlagBits - Pipeline layout creation flag bits



## <a href="#_c_specification" class="anchor"></a>C Specification

``` c
// Provided by VK_EXT_graphics_pipeline_library
typedef enum VkPipelineLayoutCreateFlagBits {
  // Provided by VK_EXT_graphics_pipeline_library
    VK_PIPELINE_LAYOUT_CREATE_INDEPENDENT_SETS_BIT_EXT = 0x00000002,
} VkPipelineLayoutCreateFlagBits;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_PIPELINE_LAYOUT_CREATE_INDEPENDENT_SETS_BIT_EXT` specifies that
  implementations **must** ensure that the properties and/or absence of
  a particular descriptor set do not influence any other properties of
  the pipeline layout. This allows pipelines libraries linked without
  `VK_PIPELINE_CREATE_LINK_TIME_OPTIMIZATION_BIT_EXT` to be created with
  a subset of the total descriptor sets.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_graphics_pipeline_library](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_graphics_pipeline_library.html),
[VkPipelineLayoutCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayoutCreateFlags.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPipelineLayoutCreateFlagBits"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
