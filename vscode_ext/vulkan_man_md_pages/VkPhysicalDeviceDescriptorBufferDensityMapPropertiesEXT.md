# VkPhysicalDeviceDescriptorBufferDensityMapPropertiesEXT(3) Manual Page

## Name

VkPhysicalDeviceDescriptorBufferDensityMapPropertiesEXT - Structure
describing descriptor buffer density map properties supported by an
implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceDescriptorBufferDensityMapPropertiesEXT` structure
is defined as:

``` c
// Provided by VK_EXT_descriptor_buffer
typedef struct VkPhysicalDeviceDescriptorBufferDensityMapPropertiesEXT {
    VkStructureType    sType;
    void*              pNext;
    size_t             combinedImageSamplerDensityMapDescriptorSize;
} VkPhysicalDeviceDescriptorBufferDensityMapPropertiesEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="limits-combinedImageSamplerDensityMapDescriptorSize"></span>
  `combinedImageSamplerDensityMapDescriptorSize` indicates the size in
  bytes of a `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER` descriptor when
  creating the descriptor with `VK_SAMPLER_CREATE_SUBSAMPLED_BIT_EXT`
  set.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceDescriptorBufferDensityMapPropertiesEXT`
structure is included in the `pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceDescriptorBufferDensityMapPropertiesEXT-sType-sType"
  id="VUID-VkPhysicalDeviceDescriptorBufferDensityMapPropertiesEXT-sType-sType"></a>
  VUID-VkPhysicalDeviceDescriptorBufferDensityMapPropertiesEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DESCRIPTOR_BUFFER_DENSITY_MAP_PROPERTIES_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_descriptor_buffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_descriptor_buffer.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceDescriptorBufferDensityMapPropertiesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
