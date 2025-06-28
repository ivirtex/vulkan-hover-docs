# VkPhysicalDeviceGraphicsPipelineLibraryPropertiesEXT(3) Manual Page

## Name

VkPhysicalDeviceGraphicsPipelineLibraryPropertiesEXT - Structure describing additional properties of graphics pipeline libraries



## [](#_c_specification)C Specification

The `VkPhysicalDeviceGraphicsPipelineLibraryPropertiesEXT` structure is defined as:

```c++
// Provided by VK_EXT_graphics_pipeline_library
typedef struct VkPhysicalDeviceGraphicsPipelineLibraryPropertiesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           graphicsPipelineLibraryFastLinking;
    VkBool32           graphicsPipelineLibraryIndependentInterpolationDecoration;
} VkPhysicalDeviceGraphicsPipelineLibraryPropertiesEXT;
```

## [](#_members)Members

- []()`graphicsPipelineLibraryFastLinking` indicates whether fast linking of graphics pipelines is supported. If it is `VK_TRUE`, creating a graphics pipeline entirely from pipeline libraries without `VK_PIPELINE_CREATE_LINK_TIME_OPTIMIZATION_BIT_EXT` is comparable in cost to recording a command in a command buffer.
- []()`graphicsPipelineLibraryIndependentInterpolationDecoration` indicates whether `NoPerspective` and `Flat` interpolation decorations in the last vertex processing stage and the fragment shader are required to match when using graphics pipeline libraries. If it is `VK_TRUE`, the interpolation decorations do not need to match. If it is `VK_FALSE`, these decorations **must** either be present in both stages or neither stage in order for a given interface variable to match.

## [](#_description)Description

If the `VkPhysicalDeviceGraphicsPipelineLibraryPropertiesEXT` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceGraphicsPipelineLibraryPropertiesEXT-sType-sType)VUID-VkPhysicalDeviceGraphicsPipelineLibraryPropertiesEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_GRAPHICS_PIPELINE_LIBRARY_PROPERTIES_EXT`

## [](#_see_also)See Also

[VK\_EXT\_graphics\_pipeline\_library](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_graphics_pipeline_library.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceGraphicsPipelineLibraryPropertiesEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0