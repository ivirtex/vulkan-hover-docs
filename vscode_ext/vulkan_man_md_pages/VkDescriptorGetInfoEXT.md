# VkDescriptorGetInfoEXT(3) Manual Page

## Name

VkDescriptorGetInfoEXT - Structure specifying parameters of descriptor
to get



## <a href="#_c_specification" class="anchor"></a>C Specification

Information about the descriptor to get is passed in a
`VkDescriptorGetInfoEXT` structure:

``` c
// Provided by VK_EXT_descriptor_buffer
typedef struct VkDescriptorGetInfoEXT {
    VkStructureType        sType;
    const void*            pNext;
    VkDescriptorType       type;
    VkDescriptorDataEXT    data;
} VkDescriptorGetInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `type` is the type of descriptor to get.

- `data` is a structure containing the information needed to get the
  descriptor.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkDescriptorGetInfoEXT-type-08018"
  id="VUID-VkDescriptorGetInfoEXT-type-08018"></a>
  VUID-VkDescriptorGetInfoEXT-type-08018  
  `type` **must** not be `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC`,
  `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC` or
  `VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK`

- <a href="#VUID-VkDescriptorGetInfoEXT-type-08019"
  id="VUID-VkDescriptorGetInfoEXT-type-08019"></a>
  VUID-VkDescriptorGetInfoEXT-type-08019  
  If `type` is `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER`, the
  `pCombinedImageSampler->sampler` member of `data` **must** be a
  [VkSampler](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampler.html) created on `device`

- <a href="#VUID-VkDescriptorGetInfoEXT-type-08020"
  id="VUID-VkDescriptorGetInfoEXT-type-08020"></a>
  VUID-VkDescriptorGetInfoEXT-type-08020  
  If `type` is `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER`, the
  `pCombinedImageSampler->imageView` member of `data` **must** be a
  [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) created on `device`, or
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

- <a href="#VUID-VkDescriptorGetInfoEXT-type-08021"
  id="VUID-VkDescriptorGetInfoEXT-type-08021"></a>
  VUID-VkDescriptorGetInfoEXT-type-08021  
  If `type` is `VK_DESCRIPTOR_TYPE_INPUT_ATTACHMENT`, the
  `pInputAttachmentImage->imageView` member of `data` **must** be a
  [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) created on `device`

- <a href="#VUID-VkDescriptorGetInfoEXT-type-08022"
  id="VUID-VkDescriptorGetInfoEXT-type-08022"></a>
  VUID-VkDescriptorGetInfoEXT-type-08022  
  If `type` is `VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE`, and if
  `pSampledImage` is not `NULL`, the `pSampledImage->imageView` member
  of `data` **must** be a [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) created on
  `device`, or [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

- <a href="#VUID-VkDescriptorGetInfoEXT-type-08023"
  id="VUID-VkDescriptorGetInfoEXT-type-08023"></a>
  VUID-VkDescriptorGetInfoEXT-type-08023  
  If `type` is `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE`, and if
  `pStorageImage` is not `NULL`, the `pStorageImage->imageView` member
  of `data` **must** be a [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) created on
  `device`, or [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

- <a href="#VUID-VkDescriptorGetInfoEXT-type-08024"
  id="VUID-VkDescriptorGetInfoEXT-type-08024"></a>
  VUID-VkDescriptorGetInfoEXT-type-08024  
  If `type` is `VK_DESCRIPTOR_TYPE_UNIFORM_TEXEL_BUFFER`,
  `pUniformTexelBuffer` is not `NULL` and `pUniformTexelBuffer->address`
  is not zero, `pUniformTexelBuffer->address` **must** be an address
  within a [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) created on `device`

- <a href="#VUID-VkDescriptorGetInfoEXT-type-08025"
  id="VUID-VkDescriptorGetInfoEXT-type-08025"></a>
  VUID-VkDescriptorGetInfoEXT-type-08025  
  If `type` is `VK_DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER`,
  `pStorageTexelBuffer` is not `NULL` and `pStorageTexelBuffer->address`
  is not zero, `pStorageTexelBuffer->address` **must** be an address
  within a [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) created on `device`

- <a href="#VUID-VkDescriptorGetInfoEXT-type-08026"
  id="VUID-VkDescriptorGetInfoEXT-type-08026"></a>
  VUID-VkDescriptorGetInfoEXT-type-08026  
  If `type` is `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER`, `pUniformBuffer` is
  not `NULL` and `pUniformBuffer->address` is not zero,
  `pUniformBuffer->address` **must** be an address within a
  [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) created on `device`

- <a href="#VUID-VkDescriptorGetInfoEXT-type-08027"
  id="VUID-VkDescriptorGetInfoEXT-type-08027"></a>
  VUID-VkDescriptorGetInfoEXT-type-08027  
  If `type` is `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER`, `pStorageBuffer` is
  not `NULL` and `pStorageBuffer->address` is not zero,
  `pStorageBuffer->address` **must** be an address within a
  [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) created on `device`

- <a href="#VUID-VkDescriptorGetInfoEXT-type-09427"
  id="VUID-VkDescriptorGetInfoEXT-type-09427"></a>
  VUID-VkDescriptorGetInfoEXT-type-09427  
  If `type` is `VK_DESCRIPTOR_TYPE_UNIFORM_TEXEL_BUFFER`,
  `pUniformBuffer` is not `NULL` , the number of texel buffer elements
  given by (⌊`pUniformBuffer->range` / (texel block size)⌋ × (texels per
  block)) where texel block size and texels per block are as defined in
  the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#formats-compatibility"
  target="_blank" rel="noopener">Compatible Formats</a> table for
  `pUniformBuffer->format`, **must** be less than or equal to
  `VkPhysicalDeviceLimits`::`maxTexelBufferElements`

- <a href="#VUID-VkDescriptorGetInfoEXT-type-09428"
  id="VUID-VkDescriptorGetInfoEXT-type-09428"></a>
  VUID-VkDescriptorGetInfoEXT-type-09428  
  If `type` is `VK_DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER`,
  `pStorageBuffer` is not `NULL` , the number of texel buffer elements
  given by (⌊`pStorageBuffer->range` / (texel block size)⌋ × (texels per
  block)) where texel block size and texels per block are as defined in
  the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#formats-compatibility"
  target="_blank" rel="noopener">Compatible Formats</a> table for
  `pStorageBuffer->format`, **must** be less than or equal to
  `VkPhysicalDeviceLimits`::`maxTexelBufferElements`

- <a href="#VUID-VkDescriptorGetInfoEXT-type-08028"
  id="VUID-VkDescriptorGetInfoEXT-type-08028"></a>
  VUID-VkDescriptorGetInfoEXT-type-08028  
  If `type` is `VK_DESCRIPTOR_TYPE_ACCELERATION_STRUCTURE_KHR` and
  `accelerationStructure` is not `0`, `accelerationStructure` **must**
  contain the address of a
  [VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureKHR.html) created
  on `device`

- <a href="#VUID-VkDescriptorGetInfoEXT-type-08029"
  id="VUID-VkDescriptorGetInfoEXT-type-08029"></a>
  VUID-VkDescriptorGetInfoEXT-type-08029  
  If `type` is `VK_DESCRIPTOR_TYPE_ACCELERATION_STRUCTURE_NV` and
  `accelerationStructure` is not `0`, `accelerationStructure` **must**
  contain the handle of a
  [VkAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureNV.html) created on
  `device`, returned by
  [vkGetAccelerationStructureHandleNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetAccelerationStructureHandleNV.html)

Valid Usage (Implicit)

- <a href="#VUID-VkDescriptorGetInfoEXT-sType-sType"
  id="VUID-VkDescriptorGetInfoEXT-sType-sType"></a>
  VUID-VkDescriptorGetInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DESCRIPTOR_GET_INFO_EXT`

- <a href="#VUID-VkDescriptorGetInfoEXT-pNext-pNext"
  id="VUID-VkDescriptorGetInfoEXT-pNext-pNext"></a>
  VUID-VkDescriptorGetInfoEXT-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkDescriptorGetInfoEXT-type-parameter"
  id="VUID-VkDescriptorGetInfoEXT-type-parameter"></a>
  VUID-VkDescriptorGetInfoEXT-type-parameter  
  `type` **must** be a valid [VkDescriptorType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorType.html)
  value

- <a href="#VUID-VkDescriptorGetInfoEXT-pSampler-parameter"
  id="VUID-VkDescriptorGetInfoEXT-pSampler-parameter"></a>
  VUID-VkDescriptorGetInfoEXT-pSampler-parameter  
  If `type` is `VK_DESCRIPTOR_TYPE_SAMPLER`, the `pSampler` member of
  `data` **must** be a valid pointer to a valid
  [VkSampler](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampler.html) handle

- <a href="#VUID-VkDescriptorGetInfoEXT-pCombinedImageSampler-parameter"
  id="VUID-VkDescriptorGetInfoEXT-pCombinedImageSampler-parameter"></a>
  VUID-VkDescriptorGetInfoEXT-pCombinedImageSampler-parameter  
  If `type` is `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER`, the
  `pCombinedImageSampler` member of `data` **must** be a valid pointer
  to a valid [VkDescriptorImageInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorImageInfo.html)
  structure

- <a href="#VUID-VkDescriptorGetInfoEXT-pInputAttachmentImage-parameter"
  id="VUID-VkDescriptorGetInfoEXT-pInputAttachmentImage-parameter"></a>
  VUID-VkDescriptorGetInfoEXT-pInputAttachmentImage-parameter  
  If `type` is `VK_DESCRIPTOR_TYPE_INPUT_ATTACHMENT`, the
  `pInputAttachmentImage` member of `data` **must** be a valid pointer
  to a valid [VkDescriptorImageInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorImageInfo.html)
  structure

- <a href="#VUID-VkDescriptorGetInfoEXT-pSampledImage-parameter"
  id="VUID-VkDescriptorGetInfoEXT-pSampledImage-parameter"></a>
  VUID-VkDescriptorGetInfoEXT-pSampledImage-parameter  
  If `type` is `VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE`, and if
  `pSampledImage` is not `NULL`, the `pSampledImage` member of `data`
  **must** be a valid pointer to a valid
  [VkDescriptorImageInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorImageInfo.html) structure

- <a href="#VUID-VkDescriptorGetInfoEXT-pStorageImage-parameter"
  id="VUID-VkDescriptorGetInfoEXT-pStorageImage-parameter"></a>
  VUID-VkDescriptorGetInfoEXT-pStorageImage-parameter  
  If `type` is `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE`, and if
  `pStorageImage` is not `NULL`, the `pStorageImage` member of `data`
  **must** be a valid pointer to a valid
  [VkDescriptorImageInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorImageInfo.html) structure

- <a href="#VUID-VkDescriptorGetInfoEXT-pUniformTexelBuffer-parameter"
  id="VUID-VkDescriptorGetInfoEXT-pUniformTexelBuffer-parameter"></a>
  VUID-VkDescriptorGetInfoEXT-pUniformTexelBuffer-parameter  
  If `type` is `VK_DESCRIPTOR_TYPE_UNIFORM_TEXEL_BUFFER`, and if
  `pUniformTexelBuffer` is not `NULL`, the `pUniformTexelBuffer` member
  of `data` **must** be a valid pointer to a valid
  [VkDescriptorAddressInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorAddressInfoEXT.html)
  structure

- <a href="#VUID-VkDescriptorGetInfoEXT-pStorageTexelBuffer-parameter"
  id="VUID-VkDescriptorGetInfoEXT-pStorageTexelBuffer-parameter"></a>
  VUID-VkDescriptorGetInfoEXT-pStorageTexelBuffer-parameter  
  If `type` is `VK_DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER`, and if
  `pStorageTexelBuffer` is not `NULL`, the `pStorageTexelBuffer` member
  of `data` **must** be a valid pointer to a valid
  [VkDescriptorAddressInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorAddressInfoEXT.html)
  structure

- <a href="#VUID-VkDescriptorGetInfoEXT-pUniformBuffer-parameter"
  id="VUID-VkDescriptorGetInfoEXT-pUniformBuffer-parameter"></a>
  VUID-VkDescriptorGetInfoEXT-pUniformBuffer-parameter  
  If `type` is `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER`, and if
  `pUniformBuffer` is not `NULL`, the `pUniformBuffer` member of `data`
  **must** be a valid pointer to a valid
  [VkDescriptorAddressInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorAddressInfoEXT.html)
  structure

- <a href="#VUID-VkDescriptorGetInfoEXT-pStorageBuffer-parameter"
  id="VUID-VkDescriptorGetInfoEXT-pStorageBuffer-parameter"></a>
  VUID-VkDescriptorGetInfoEXT-pStorageBuffer-parameter  
  If `type` is `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER`, and if
  `pStorageBuffer` is not `NULL`, the `pStorageBuffer` member of `data`
  **must** be a valid pointer to a valid
  [VkDescriptorAddressInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorAddressInfoEXT.html)
  structure

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_descriptor_buffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_descriptor_buffer.html),
[VkDescriptorDataEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorDataEXT.html),
[VkDescriptorType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorType.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetDescriptorEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDescriptorEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDescriptorGetInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
