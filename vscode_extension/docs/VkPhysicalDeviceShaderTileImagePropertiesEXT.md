# VkPhysicalDeviceShaderTileImagePropertiesEXT(3) Manual Page

## Name

VkPhysicalDeviceShaderTileImagePropertiesEXT - Structure containing information about tile image support for a physical device



## [](#_c_specification)C Specification

The `VkPhysicalDeviceShaderTileImagePropertiesEXT` structure is defined as:

```c++
// Provided by VK_EXT_shader_tile_image
typedef struct VkPhysicalDeviceShaderTileImagePropertiesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           shaderTileImageCoherentReadAccelerated;
    VkBool32           shaderTileImageReadSampleFromPixelRateInvocation;
    VkBool32           shaderTileImageReadFromHelperInvocation;
} VkPhysicalDeviceShaderTileImagePropertiesEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `shaderTileImageCoherentReadAccelerated` is a boolean that will be `VK_TRUE` if coherent reads of tile image data is accelerated.
- `shaderTileImageReadSampleFromPixelRateInvocation` is a boolean that will be `VK_TRUE` if reading from samples from a pixel rate fragment invocation is supported when [VkPipelineMultisampleStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineMultisampleStateCreateInfo.html)::`rasterizationSamples` &gt; 1.
- `shaderTileImageReadFromHelperInvocation` is a boolean that will be `VK_TRUE` if reads of tile image data from helper fragment invocations result in valid values.

## [](#_description)Description

If the `VkPhysicalDeviceShaderTileImagePropertiesEXT` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

These are properties of the tile image information of a physical device.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceShaderTileImagePropertiesEXT-sType-sType)VUID-VkPhysicalDeviceShaderTileImagePropertiesEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_TILE_IMAGE_PROPERTIES_EXT`

## [](#_see_also)See Also

[VK\_EXT\_shader\_tile\_image](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_shader_tile_image.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceShaderTileImagePropertiesEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0