# VkPhysicalDeviceGraphicsPipelineLibraryPropertiesEXT(3) Manual Page

## Name

VkPhysicalDeviceGraphicsPipelineLibraryPropertiesEXT - Structure
describing additional properties of graphics pipeline libraries



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceGraphicsPipelineLibraryPropertiesEXT` structure is
defined as:

``` c
// Provided by VK_EXT_graphics_pipeline_library
typedef struct VkPhysicalDeviceGraphicsPipelineLibraryPropertiesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           graphicsPipelineLibraryFastLinking;
    VkBool32           graphicsPipelineLibraryIndependentInterpolationDecoration;
} VkPhysicalDeviceGraphicsPipelineLibraryPropertiesEXT;
```

## <a href="#_members" class="anchor"></a>Members

- <span id="limits-graphicsPipelineLibraryFastLinking"></span>
  `graphicsPipelineLibraryFastLinking` indicates whether fast linking of
  graphics pipelines is supported. If it is `VK_TRUE`, creating a
  graphics pipeline entirely from pipeline libraries without
  `VK_PIPELINE_CREATE_LINK_TIME_OPTIMIZATION_BIT_EXT` is comparable in
  cost to recording a command in a command buffer.

- <span id="limits-graphicsPipelineLibraryIndependentInterpolationDecoration"></span>
  `graphicsPipelineLibraryIndependentInterpolationDecoration` indicates
  whether `NoPerspective` and `Flat` interpolation decorations in the
  last vertex processing stage and the fragment shader are required to
  match when using graphics pipeline libraries. If it is `VK_TRUE`, the
  interpolation decorations do not need to match. If it is `VK_FALSE`,
  these decorations **must** either be present in both stages or neither
  stage in order for a given interface variable to match.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceGraphicsPipelineLibraryPropertiesEXT` structure
is included in the `pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceGraphicsPipelineLibraryPropertiesEXT-sType-sType"
  id="VUID-VkPhysicalDeviceGraphicsPipelineLibraryPropertiesEXT-sType-sType"></a>
  VUID-VkPhysicalDeviceGraphicsPipelineLibraryPropertiesEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_GRAPHICS_PIPELINE_LIBRARY_PROPERTIES_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_graphics_pipeline_library](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_graphics_pipeline_library.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceGraphicsPipelineLibraryPropertiesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
