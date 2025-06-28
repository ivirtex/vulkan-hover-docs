# VkPhysicalDeviceDescriptorBufferDensityMapPropertiesEXT(3) Manual Page

## Name

VkPhysicalDeviceDescriptorBufferDensityMapPropertiesEXT - Structure describing descriptor buffer density map properties supported by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceDescriptorBufferDensityMapPropertiesEXT` structure is defined as:

```c++
// Provided by VK_EXT_descriptor_buffer
typedef struct VkPhysicalDeviceDescriptorBufferDensityMapPropertiesEXT {
    VkStructureType    sType;
    void*              pNext;
    size_t             combinedImageSamplerDensityMapDescriptorSize;
} VkPhysicalDeviceDescriptorBufferDensityMapPropertiesEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`combinedImageSamplerDensityMapDescriptorSize` indicates the size in bytes of a `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER` descriptor when creating the descriptor with `VK_SAMPLER_CREATE_SUBSAMPLED_BIT_EXT` set.

## [](#_description)Description

If the `VkPhysicalDeviceDescriptorBufferDensityMapPropertiesEXT` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceDescriptorBufferDensityMapPropertiesEXT-sType-sType)VUID-VkPhysicalDeviceDescriptorBufferDensityMapPropertiesEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DESCRIPTOR_BUFFER_DENSITY_MAP_PROPERTIES_EXT`

## [](#_see_also)See Also

[VK\_EXT\_descriptor\_buffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_descriptor_buffer.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceDescriptorBufferDensityMapPropertiesEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0