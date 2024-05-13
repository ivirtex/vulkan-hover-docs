# VkPhysicalDeviceShaderTileImagePropertiesEXT(3) Manual Page

## Name

VkPhysicalDeviceShaderTileImagePropertiesEXT - Structure containing
information about tile image support for a physical device



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceShaderTileImagePropertiesEXT` structure is defined
as:

``` c
// Provided by VK_EXT_shader_tile_image
typedef struct VkPhysicalDeviceShaderTileImagePropertiesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           shaderTileImageCoherentReadAccelerated;
    VkBool32           shaderTileImageReadSampleFromPixelRateInvocation;
    VkBool32           shaderTileImageReadFromHelperInvocation;
} VkPhysicalDeviceShaderTileImagePropertiesEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `shaderTileImageCoherentReadAccelerated` is a boolean that will be
  `VK_TRUE` if coherent reads of tile image data is accelerated.

- `shaderTileImageReadSampleFromPixelRateInvocation` is a boolean that
  will be `VK_TRUE` if reading from samples from a pixel rate fragment
  invocation is supported when
  [VkPipelineMultisampleStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineMultisampleStateCreateInfo.html)::`rasterizationSamples`
  \> 1.

- `shaderTileImageReadFromHelperInvocation` is a boolean that will be
  `VK_TRUE` if reads of tile image data from helper fragment invocations
  result in valid values.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceShaderTileImagePropertiesEXT` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

These are properties of the tile image information of a physical device.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceShaderTileImagePropertiesEXT-sType-sType"
  id="VUID-VkPhysicalDeviceShaderTileImagePropertiesEXT-sType-sType"></a>
  VUID-VkPhysicalDeviceShaderTileImagePropertiesEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_TILE_IMAGE_PROPERTIES_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_shader_tile_image](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_shader_tile_image.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceShaderTileImagePropertiesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
