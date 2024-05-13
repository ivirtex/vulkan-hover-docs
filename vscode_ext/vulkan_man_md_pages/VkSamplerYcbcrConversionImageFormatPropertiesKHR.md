# VkSamplerYcbcrConversionImageFormatProperties(3) Manual Page

## Name

VkSamplerYcbcrConversionImageFormatProperties - Structure specifying
combined image sampler descriptor count for multi-planar images



## <a href="#_c_specification" class="anchor"></a>C Specification

To determine the number of combined image samplers required to support a
multi-planar format, add
[VkSamplerYcbcrConversionImageFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversionImageFormatProperties.html)
to the `pNext` chain of the
[VkImageFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatProperties2.html) structure in a
call to `vkGetPhysicalDeviceImageFormatProperties2`.

The `VkSamplerYcbcrConversionImageFormatProperties` structure is defined
as:

``` c
// Provided by VK_VERSION_1_1
typedef struct VkSamplerYcbcrConversionImageFormatProperties {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           combinedImageSamplerDescriptorCount;
} VkSamplerYcbcrConversionImageFormatProperties;
```

or the equivalent

``` c
// Provided by VK_KHR_sampler_ycbcr_conversion
typedef VkSamplerYcbcrConversionImageFormatProperties VkSamplerYcbcrConversionImageFormatPropertiesKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `combinedImageSamplerDescriptorCount` is the number of combined image
  sampler descriptors that the implementation uses to access the format.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a
  href="#VUID-VkSamplerYcbcrConversionImageFormatProperties-sType-sType"
  id="VUID-VkSamplerYcbcrConversionImageFormatProperties-sType-sType"></a>
  VUID-VkSamplerYcbcrConversionImageFormatProperties-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_IMAGE_FORMAT_PROPERTIES`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSamplerYcbcrConversionImageFormatProperties"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
