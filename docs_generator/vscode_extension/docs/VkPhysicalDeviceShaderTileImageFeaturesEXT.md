# VkPhysicalDeviceShaderTileImageFeaturesEXT(3) Manual Page

## Name

VkPhysicalDeviceShaderTileImageFeaturesEXT - Structure describing tile image features supported by the implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceShaderTileImageFeaturesEXT` structure is defined as:

```c++
// Provided by VK_EXT_shader_tile_image
typedef struct VkPhysicalDeviceShaderTileImageFeaturesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           shaderTileImageColorReadAccess;
    VkBool32           shaderTileImageDepthReadAccess;
    VkBool32           shaderTileImageStencilReadAccess;
} VkPhysicalDeviceShaderTileImageFeaturesEXT;
```

## [](#_members)Members

The members of the `VkPhysicalDeviceShaderTileImageFeaturesEXT` structure describe the following features:

## [](#_description)Description

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`shaderTileImageColorReadAccess` indicates that the implementation supports the `TileImageColorReadAccessEXT` SPIR-V capability.
- []()`shaderTileImageDepthReadAccess` indicates that the implementation supports the `TileImageDepthReadAccessEXT` SPIR-V capability.
- []()`shaderTileImageStencilReadAccess` indicates that the implementation supports the `TileImageStencilReadAccessEXT` SPIR-V capability.

If the `VkPhysicalDeviceShaderTileImageFeaturesEXT` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceShaderTileImageFeaturesEXT`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceShaderTileImageFeaturesEXT-sType-sType)VUID-VkPhysicalDeviceShaderTileImageFeaturesEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_TILE_IMAGE_FEATURES_EXT`

## [](#_see_also)See Also

[VK\_EXT\_shader\_tile\_image](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_shader_tile_image.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceShaderTileImageFeaturesEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0