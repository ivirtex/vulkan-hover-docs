# VkDescriptorGetInfoEXT(3) Manual Page

## Name

VkDescriptorGetInfoEXT - Structure specifying parameters of descriptor to get



## [](#_c_specification)C Specification

Information about the descriptor to get is passed in a `VkDescriptorGetInfoEXT` structure:

```c++
// Provided by VK_EXT_descriptor_buffer
typedef struct VkDescriptorGetInfoEXT {
    VkStructureType        sType;
    const void*            pNext;
    VkDescriptorType       type;
    VkDescriptorDataEXT    data;
} VkDescriptorGetInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `type` is the type of descriptor to get.
- `data` is a structure containing the information needed to get the descriptor.

## [](#_description)Description

Valid Usage

- [](#VUID-VkDescriptorGetInfoEXT-type-08018)VUID-VkDescriptorGetInfoEXT-type-08018  
  `type` **must** not be `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC`, `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC` or `VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK`
- [](#VUID-VkDescriptorGetInfoEXT-type-08019)VUID-VkDescriptorGetInfoEXT-type-08019  
  If `type` is `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER`, the `pCombinedImageSampler->sampler` member of `data` **must** be a [VkSampler](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampler.html) created on `device`
- [](#VUID-VkDescriptorGetInfoEXT-type-08020)VUID-VkDescriptorGetInfoEXT-type-08020  
  If `type` is `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER`, the `pCombinedImageSampler->imageView` member of `data` **must** be a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) created on `device`, or [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html)
- [](#VUID-VkDescriptorGetInfoEXT-type-08021)VUID-VkDescriptorGetInfoEXT-type-08021  
  If `type` is `VK_DESCRIPTOR_TYPE_INPUT_ATTACHMENT`, the `pInputAttachmentImage->imageView` member of `data` **must** be a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) created on `device`
- [](#VUID-VkDescriptorGetInfoEXT-type-08022)VUID-VkDescriptorGetInfoEXT-type-08022  
  If `type` is `VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE`, and if `pSampledImage` is not `NULL`, the `pSampledImage->imageView` member of `data` **must** be a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) created on `device`, or [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html)
- [](#VUID-VkDescriptorGetInfoEXT-type-08023)VUID-VkDescriptorGetInfoEXT-type-08023  
  If `type` is `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE`, and if `pStorageImage` is not `NULL`, the `pStorageImage->imageView` member of `data` **must** be a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) created on `device`, or [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html)
- [](#VUID-VkDescriptorGetInfoEXT-type-08024)VUID-VkDescriptorGetInfoEXT-type-08024  
  If `type` is `VK_DESCRIPTOR_TYPE_UNIFORM_TEXEL_BUFFER`, `pUniformTexelBuffer` is not `NULL` and `pUniformTexelBuffer->address` is not zero, `pUniformTexelBuffer->address` **must** be an address within a [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) created on `device`
- [](#VUID-VkDescriptorGetInfoEXT-type-08025)VUID-VkDescriptorGetInfoEXT-type-08025  
  If `type` is `VK_DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER`, `pStorageTexelBuffer` is not `NULL` and `pStorageTexelBuffer->address` is not zero, `pStorageTexelBuffer->address` **must** be an address within a [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) created on `device`
- [](#VUID-VkDescriptorGetInfoEXT-type-08026)VUID-VkDescriptorGetInfoEXT-type-08026  
  If `type` is `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER`, `pUniformBuffer` is not `NULL` and `pUniformBuffer->address` is not zero, `pUniformBuffer->address` **must** be an address within a [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) created on `device`
- [](#VUID-VkDescriptorGetInfoEXT-type-08027)VUID-VkDescriptorGetInfoEXT-type-08027  
  If `type` is `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER`, `pStorageBuffer` is not `NULL` and `pStorageBuffer->address` is not zero, `pStorageBuffer->address` **must** be an address within a [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) created on `device`
- [](#VUID-VkDescriptorGetInfoEXT-type-09427)VUID-VkDescriptorGetInfoEXT-type-09427  
  If `type` is `VK_DESCRIPTOR_TYPE_UNIFORM_TEXEL_BUFFER`, `pUniformBuffer` is not `NULL` , the number of texel buffer elements given by (⌊`pUniformBuffer->range` / (texel block size)⌋ × (texels per block)) where texel block size and texels per block are as defined in the [Compatible Formats](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-compatibility) table for `pUniformBuffer->format`, **must** be less than or equal to `VkPhysicalDeviceLimits`::`maxTexelBufferElements`
- [](#VUID-VkDescriptorGetInfoEXT-type-09428)VUID-VkDescriptorGetInfoEXT-type-09428  
  If `type` is `VK_DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER`, `pStorageBuffer` is not `NULL` , the number of texel buffer elements given by (⌊`pStorageBuffer->range` / (texel block size)⌋ × (texels per block)) where texel block size and texels per block are as defined in the [Compatible Formats](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-compatibility) table for `pStorageBuffer->format`, **must** be less than or equal to `VkPhysicalDeviceLimits`::`maxTexelBufferElements`
- [](#VUID-VkDescriptorGetInfoEXT-type-08028)VUID-VkDescriptorGetInfoEXT-type-08028  
  If `type` is `VK_DESCRIPTOR_TYPE_ACCELERATION_STRUCTURE_KHR` and `accelerationStructure` is not `0`, `accelerationStructure` **must** contain the address of a [VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureKHR.html) created on `device`
- [](#VUID-VkDescriptorGetInfoEXT-type-08029)VUID-VkDescriptorGetInfoEXT-type-08029  
  If `type` is `VK_DESCRIPTOR_TYPE_ACCELERATION_STRUCTURE_NV` and `accelerationStructure` is not `0`, `accelerationStructure` **must** contain the handle of a [VkAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureNV.html) created on `device`, returned by [vkGetAccelerationStructureHandleNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetAccelerationStructureHandleNV.html)
- [](#VUID-VkDescriptorGetInfoEXT-type-09701)VUID-VkDescriptorGetInfoEXT-type-09701  
  If `type` is `VK_DESCRIPTOR_TYPE_TENSOR_ARM`, a [VkDescriptorGetTensorInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorGetTensorInfoARM.html) structure **must** be included in the `pNext` chain and `data` is ignored

Valid Usage (Implicit)

- [](#VUID-VkDescriptorGetInfoEXT-sType-sType)VUID-VkDescriptorGetInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DESCRIPTOR_GET_INFO_EXT`
- [](#VUID-VkDescriptorGetInfoEXT-pNext-pNext)VUID-VkDescriptorGetInfoEXT-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of [VkDescriptorGetTensorInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorGetTensorInfoARM.html)
- [](#VUID-VkDescriptorGetInfoEXT-sType-unique)VUID-VkDescriptorGetInfoEXT-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkDescriptorGetInfoEXT-type-parameter)VUID-VkDescriptorGetInfoEXT-type-parameter  
  `type` **must** be a valid [VkDescriptorType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorType.html) value
- [](#VUID-VkDescriptorGetInfoEXT-pSampler-parameter)VUID-VkDescriptorGetInfoEXT-pSampler-parameter  
  If `type` is `VK_DESCRIPTOR_TYPE_SAMPLER`, the `pSampler` member of `data` **must** be a valid pointer to a valid [VkSampler](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampler.html) handle
- [](#VUID-VkDescriptorGetInfoEXT-pCombinedImageSampler-parameter)VUID-VkDescriptorGetInfoEXT-pCombinedImageSampler-parameter  
  If `type` is `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER`, the `pCombinedImageSampler` member of `data` **must** be a valid pointer to a valid [VkDescriptorImageInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorImageInfo.html) structure
- [](#VUID-VkDescriptorGetInfoEXT-pInputAttachmentImage-parameter)VUID-VkDescriptorGetInfoEXT-pInputAttachmentImage-parameter  
  If `type` is `VK_DESCRIPTOR_TYPE_INPUT_ATTACHMENT`, the `pInputAttachmentImage` member of `data` **must** be a valid pointer to a valid [VkDescriptorImageInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorImageInfo.html) structure
- [](#VUID-VkDescriptorGetInfoEXT-pSampledImage-parameter)VUID-VkDescriptorGetInfoEXT-pSampledImage-parameter  
  If `type` is `VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE`, and if `pSampledImage` is not `NULL`, the `pSampledImage` member of `data` **must** be a valid pointer to a valid [VkDescriptorImageInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorImageInfo.html) structure
- [](#VUID-VkDescriptorGetInfoEXT-pStorageImage-parameter)VUID-VkDescriptorGetInfoEXT-pStorageImage-parameter  
  If `type` is `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE`, and if `pStorageImage` is not `NULL`, the `pStorageImage` member of `data` **must** be a valid pointer to a valid [VkDescriptorImageInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorImageInfo.html) structure
- [](#VUID-VkDescriptorGetInfoEXT-pUniformTexelBuffer-parameter)VUID-VkDescriptorGetInfoEXT-pUniformTexelBuffer-parameter  
  If `type` is `VK_DESCRIPTOR_TYPE_UNIFORM_TEXEL_BUFFER`, and if `pUniformTexelBuffer` is not `NULL`, the `pUniformTexelBuffer` member of `data` **must** be a valid pointer to a valid [VkDescriptorAddressInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorAddressInfoEXT.html) structure
- [](#VUID-VkDescriptorGetInfoEXT-pStorageTexelBuffer-parameter)VUID-VkDescriptorGetInfoEXT-pStorageTexelBuffer-parameter  
  If `type` is `VK_DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER`, and if `pStorageTexelBuffer` is not `NULL`, the `pStorageTexelBuffer` member of `data` **must** be a valid pointer to a valid [VkDescriptorAddressInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorAddressInfoEXT.html) structure
- [](#VUID-VkDescriptorGetInfoEXT-pUniformBuffer-parameter)VUID-VkDescriptorGetInfoEXT-pUniformBuffer-parameter  
  If `type` is `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER`, and if `pUniformBuffer` is not `NULL`, the `pUniformBuffer` member of `data` **must** be a valid pointer to a valid [VkDescriptorAddressInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorAddressInfoEXT.html) structure
- [](#VUID-VkDescriptorGetInfoEXT-pStorageBuffer-parameter)VUID-VkDescriptorGetInfoEXT-pStorageBuffer-parameter  
  If `type` is `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER`, and if `pStorageBuffer` is not `NULL`, the `pStorageBuffer` member of `data` **must** be a valid pointer to a valid [VkDescriptorAddressInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorAddressInfoEXT.html) structure

## [](#_see_also)See Also

[VK\_EXT\_descriptor\_buffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_descriptor_buffer.html), [VkDescriptorDataEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorDataEXT.html), [VkDescriptorType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorType.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetDescriptorEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDescriptorEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDescriptorGetInfoEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0