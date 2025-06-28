# VkGraphicsPipelineLibraryCreateInfoEXT(3) Manual Page

## Name

VkGraphicsPipelineLibraryCreateInfoEXT - Structure specifying the subsets of the graphics pipeline being compiled



## [](#_c_specification)C Specification

The `VkGraphicsPipelineLibraryCreateInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_graphics_pipeline_library
typedef struct VkGraphicsPipelineLibraryCreateInfoEXT {
    VkStructureType                      sType;
    const void*                          pNext;
    VkGraphicsPipelineLibraryFlagsEXT    flags;
} VkGraphicsPipelineLibraryCreateInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is a bitmask of [VkGraphicsPipelineLibraryFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineLibraryFlagBitsEXT.html) specifying the subsets of the graphics pipeline that are being compiled.

## [](#_description)Description

If a `VkGraphicsPipelineLibraryCreateInfoEXT` structure is included in the `pNext` chain of [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineCreateInfo.html), it specifies the [subsets of the graphics pipeline](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-graphics-subsets) being created, excluding any subsets from linked pipeline libraries. If the pipeline is created with pipeline libraries, state from those libraries is aggregated with said subset.

If this structure is omitted, and either [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineCreateInfo.html)::`flags` includes `VK_PIPELINE_CREATE_LIBRARY_BIT_KHR` or the [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineCreateInfo.html)::`pNext` chain includes a [VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLibraryCreateInfoKHR.html) structure with a `libraryCount` greater than `0`, it is as if `flags` is `0`. Otherwise if this structure is omitted, it is as if `flags` includes all possible subsets of the graphics pipeline (i.e. a [complete graphics pipeline](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-graphics-subsets-complete)).

Valid Usage (Implicit)

- [](#VUID-VkGraphicsPipelineLibraryCreateInfoEXT-sType-sType)VUID-VkGraphicsPipelineLibraryCreateInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_GRAPHICS_PIPELINE_LIBRARY_CREATE_INFO_EXT`
- [](#VUID-VkGraphicsPipelineLibraryCreateInfoEXT-flags-parameter)VUID-VkGraphicsPipelineLibraryCreateInfoEXT-flags-parameter  
  `flags` **must** be a valid combination of [VkGraphicsPipelineLibraryFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineLibraryFlagBitsEXT.html) values
- [](#VUID-VkGraphicsPipelineLibraryCreateInfoEXT-flags-requiredbitmask)VUID-VkGraphicsPipelineLibraryCreateInfoEXT-flags-requiredbitmask  
  `flags` **must** not be `0`

## [](#_see_also)See Also

[VK\_EXT\_graphics\_pipeline\_library](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_graphics_pipeline_library.html), [VkGraphicsPipelineLibraryFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineLibraryFlagsEXT.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkGraphicsPipelineLibraryCreateInfoEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0