# VkDescriptorDataEXT(3) Manual Page

## Name

VkDescriptorDataEXT - Structure specifying descriptor data



## [](#_c_specification)C Specification

Data describing the descriptor is passed in a `VkDescriptorDataEXT` structure:

```c++
// Provided by VK_EXT_descriptor_buffer
typedef union VkDescriptorDataEXT {
    const VkSampler*                     pSampler;
    const VkDescriptorImageInfo*         pCombinedImageSampler;
    const VkDescriptorImageInfo*         pInputAttachmentImage;
    const VkDescriptorImageInfo*         pSampledImage;
    const VkDescriptorImageInfo*         pStorageImage;
    const VkDescriptorAddressInfoEXT*    pUniformTexelBuffer;
    const VkDescriptorAddressInfoEXT*    pStorageTexelBuffer;
    const VkDescriptorAddressInfoEXT*    pUniformBuffer;
    const VkDescriptorAddressInfoEXT*    pStorageBuffer;
    VkDeviceAddress                      accelerationStructure;
} VkDescriptorDataEXT;
```

## [](#_members)Members

- `pSampler` is a pointer to a [VkSampler](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampler.html) handle specifying the parameters of a `VK_DESCRIPTOR_TYPE_SAMPLER` descriptor.
- `pCombinedImageSampler` is a pointer to a [VkDescriptorImageInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorImageInfo.html) structure specifying the parameters of a `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER` descriptor.
- `pInputAttachmentImage` is a pointer to a [VkDescriptorImageInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorImageInfo.html) structure specifying the parameters of a `VK_DESCRIPTOR_TYPE_INPUT_ATTACHMENT` descriptor.
- `pSampledImage` is a pointer to a [VkDescriptorImageInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorImageInfo.html) structure specifying the parameters of a `VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE` descriptor.
- `pStorageImage` is a pointer to a [VkDescriptorImageInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorImageInfo.html) structure specifying the parameters of a `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE` descriptor.
- `pUniformTexelBuffer` is a pointer to a [VkDescriptorAddressInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorAddressInfoEXT.html) structure specifying the parameters of a `VK_DESCRIPTOR_TYPE_UNIFORM_TEXEL_BUFFER` descriptor.
- `pStorageTexelBuffer` is a pointer to a [VkDescriptorAddressInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorAddressInfoEXT.html) structure specifying the parameters of a `VK_DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER` descriptor.
- `pUniformBuffer` is a pointer to a [VkDescriptorAddressInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorAddressInfoEXT.html) structure specifying the parameters of a `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER` descriptor.
- `pStorageBuffer` is a pointer to a [VkDescriptorAddressInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorAddressInfoEXT.html) structure specifying the parameters of a `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER` descriptor.
- `accelerationStructure` is the address of a [VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureKHR.html) specifying the parameters of a `VK_DESCRIPTOR_TYPE_ACCELERATION_STRUCTURE_KHR` descriptor , or a [VkAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureNV.html) handle specifying the parameters of a `VK_DESCRIPTOR_TYPE_ACCELERATION_STRUCTURE_NV` descriptor.

## [](#_description)Description

If the [`nullDescriptor`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-nullDescriptor) feature is enabled, `pSampledImage`, `pStorageImage`, `pUniformTexelBuffer`, `pStorageTexelBuffer`, `pUniformBuffer`, and `pStorageBuffer` **can** each be `NULL`. Loads from a null descriptor return zero values and stores and atomics to a null descriptor are discarded.

If the [`nullDescriptor`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-nullDescriptor) feature is enabled, `accelerationStructure` **can** be `0`. A null acceleration structure descriptor results in the miss shader being invoked.

Valid Usage

- [](#VUID-VkDescriptorDataEXT-type-08030)VUID-VkDescriptorDataEXT-type-08030  
  If [VkDescriptorGetInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorGetInfoEXT.html):`type` is `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER`, and `pUniformBuffer->address` is the address of a non-sparse buffer, then that buffer **must** be bound completely and contiguously to a single `VkDeviceMemory` object
- [](#VUID-VkDescriptorDataEXT-type-08031)VUID-VkDescriptorDataEXT-type-08031  
  If [VkDescriptorGetInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorGetInfoEXT.html):`type` is `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER`, and `pStorageBuffer->address` is the address of a non-sparse buffer, then that buffer **must** be bound completely and contiguously to a single `VkDeviceMemory` object
- [](#VUID-VkDescriptorDataEXT-type-08032)VUID-VkDescriptorDataEXT-type-08032  
  If [VkDescriptorGetInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorGetInfoEXT.html):`type` is `VK_DESCRIPTOR_TYPE_UNIFORM_TEXEL_BUFFER`, and `pUniformTexelBuffer->address` is the address of a non-sparse buffer, then that buffer **must** be bound completely and contiguously to a single `VkDeviceMemory` object
- [](#VUID-VkDescriptorDataEXT-type-08033)VUID-VkDescriptorDataEXT-type-08033  
  If [VkDescriptorGetInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorGetInfoEXT.html):`type` is `VK_DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER`, and `pStorageTexelBuffer->address` is the address of a non-sparse buffer, then that buffer **must** be bound completely and contiguously to a single `VkDeviceMemory` object
- [](#VUID-VkDescriptorDataEXT-type-08034)VUID-VkDescriptorDataEXT-type-08034  
  If [VkDescriptorGetInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorGetInfoEXT.html):`type` is `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER`, and the [`nullDescriptor`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-nullDescriptor) feature is not enabled, `pCombinedImageSampler->imageView` **must** not be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html)
- [](#VUID-VkDescriptorDataEXT-type-08035)VUID-VkDescriptorDataEXT-type-08035  
  If [VkDescriptorGetInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorGetInfoEXT.html):`type` is `VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE`, and the [`nullDescriptor`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-nullDescriptor) feature is not enabled, `pSampledImage` **must** not be `NULL` and `pSampledImage->imageView` **must** not be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html)
- [](#VUID-VkDescriptorDataEXT-type-08036)VUID-VkDescriptorDataEXT-type-08036  
  If [VkDescriptorGetInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorGetInfoEXT.html):`type` is `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE`, and the [`nullDescriptor`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-nullDescriptor) feature is not enabled, `pStorageImage` **must** not be `NULL` and `pStorageImage->imageView` **must** not be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html)
- [](#VUID-VkDescriptorDataEXT-type-08037)VUID-VkDescriptorDataEXT-type-08037  
  If [VkDescriptorGetInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorGetInfoEXT.html):`type` is `VK_DESCRIPTOR_TYPE_UNIFORM_TEXEL_BUFFER`, and the [`nullDescriptor`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-nullDescriptor) feature is not enabled, `pUniformTexelBuffer` **must** not be `NULL`
- [](#VUID-VkDescriptorDataEXT-type-08038)VUID-VkDescriptorDataEXT-type-08038  
  If [VkDescriptorGetInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorGetInfoEXT.html):`type` is `VK_DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER`, and the [`nullDescriptor`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-nullDescriptor) feature is not enabled, `pStorageTexelBuffer` **must** not be `NULL`
- [](#VUID-VkDescriptorDataEXT-type-08039)VUID-VkDescriptorDataEXT-type-08039  
  If [VkDescriptorGetInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorGetInfoEXT.html):`type` is `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER`, and the [`nullDescriptor`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-nullDescriptor) feature is not enabled, `pUniformBuffer` **must** not be `NULL`
- [](#VUID-VkDescriptorDataEXT-type-08040)VUID-VkDescriptorDataEXT-type-08040  
  If [VkDescriptorGetInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorGetInfoEXT.html):`type` is `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER`, and the [`nullDescriptor`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-nullDescriptor) feature is not enabled, `pStorageBuffer` **must** not be `NULL`
- [](#VUID-VkDescriptorDataEXT-type-08041)VUID-VkDescriptorDataEXT-type-08041  
  If [VkDescriptorGetInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorGetInfoEXT.html):`type` is `VK_DESCRIPTOR_TYPE_ACCELERATION_STRUCTURE_KHR`, and the [`nullDescriptor`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-nullDescriptor) feature is not enabled, `accelerationStructure` **must** not be `0`
- [](#VUID-VkDescriptorDataEXT-type-08042)VUID-VkDescriptorDataEXT-type-08042  
  If [VkDescriptorGetInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorGetInfoEXT.html):`type` is `VK_DESCRIPTOR_TYPE_ACCELERATION_STRUCTURE_NV`, and the [`nullDescriptor`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-nullDescriptor) feature is not enabled, `accelerationStructure` **must** not be `0`

## [](#_see_also)See Also

[VK\_EXT\_descriptor\_buffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_descriptor_buffer.html), [VkDescriptorAddressInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorAddressInfoEXT.html), [VkDescriptorGetInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorGetInfoEXT.html), [VkDescriptorImageInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorImageInfo.html), [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html), [VkSampler](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampler.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDescriptorDataEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0