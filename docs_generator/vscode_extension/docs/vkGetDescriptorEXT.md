# vkGetDescriptorEXT(3) Manual Page

## Name

vkGetDescriptorEXT - To get a descriptor to place in a buffer



## [](#_c_specification)C Specification

To get descriptor data to place in a buffer, call:

```c++
// Provided by VK_EXT_descriptor_buffer
void vkGetDescriptorEXT(
    VkDevice                                    device,
    const VkDescriptorGetInfoEXT*               pDescriptorInfo,
    size_t                                      dataSize,
    void*                                       pDescriptor);
```

## [](#_parameters)Parameters

- `device` is the logical device that gets the descriptor.
- `pDescriptorInfo` is a pointer to a [VkDescriptorGetInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorGetInfoEXT.html) structure specifying the parameters of the descriptor to get.
- `dataSize` is the amount of the descriptor data to get in bytes.
- `pDescriptor` is a pointer to an application-allocated buffer where the descriptor will be written.

## [](#_description)Description

The size of the data for each descriptor type is determined by the value in [VkPhysicalDeviceDescriptorBufferPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDescriptorBufferPropertiesEXT.html). This value also defines the stride in bytes for arrays of that descriptor type.

If the [VkPhysicalDeviceDescriptorBufferPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDescriptorBufferPropertiesEXT.html)::`combinedImageSamplerDescriptorSingleArray` property is `VK_FALSE` the implementation requires an array of `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER` descriptors to be written into a descriptor buffer as an array of image descriptors, immediately followed by an array of sampler descriptors. Applications **must** write the first [VkPhysicalDeviceDescriptorBufferPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDescriptorBufferPropertiesEXT.html)::`sampledImageDescriptorSize` bytes of the data returned through `pDescriptor` to the first array, and the remaining [VkPhysicalDeviceDescriptorBufferPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDescriptorBufferPropertiesEXT.html)::`samplerDescriptorSize` bytes of the data to the second array. For variable-sized descriptor bindings of `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER` descriptors, the two arrays each have a size equal to the upper bound `descriptorCount` of that binding.

A descriptor obtained by this command references the underlying [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) or [VkSampler](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampler.html), and these objects **must** not be destroyed before the last time a descriptor is dynamically accessed. For descriptor types which consume an address instead of an object, the underlying [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) is referenced instead.

Valid Usage

- [](#VUID-vkGetDescriptorEXT-None-08015)VUID-vkGetDescriptorEXT-None-08015  
  The [`descriptorBuffer`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-descriptorBuffer) feature **must** be enabled
- [](#VUID-vkGetDescriptorEXT-dataSize-08125)VUID-vkGetDescriptorEXT-dataSize-08125  
  If `pDescriptorInfo->type` is not `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER` or `pDescriptorInfo->data.pCombinedImageSampler` has an `imageView` member that was not created with a `VkSamplerYcbcrConversionInfo` structure in its `pNext` chain, `dataSize` **must** equal the size of a descriptor of type [VkDescriptorGetInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorGetInfoEXT.html)::`type` determined by the value in [VkPhysicalDeviceDescriptorBufferPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDescriptorBufferPropertiesEXT.html) , or determined by [VkPhysicalDeviceDescriptorBufferDensityMapPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDescriptorBufferDensityMapPropertiesEXT.html)::`combinedImageSamplerDensityMapDescriptorSize` if `pDescriptorInfo` specifies a `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER` whose [VkSampler](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampler.html) was created with `VK_SAMPLER_CREATE_SUBSAMPLED_BIT_EXT` set
- [](#VUID-vkGetDescriptorEXT-descriptorType-09469)VUID-vkGetDescriptorEXT-descriptorType-09469  
  If `pDescriptorInfo->type` is `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER` and `pDescriptorInfo->data.pCombinedImageSampler` has an `imageView` member that was created with a `VkSamplerYcbcrConversionInfo` structure in its `pNext` chain, `dataSize` **must** equal the size of [VkPhysicalDeviceDescriptorBufferPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDescriptorBufferPropertiesEXT.html)::`combinedImageSamplerDescriptorSize` times [VkSamplerYcbcrConversionImageFormatProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrConversionImageFormatProperties.html)::`combinedImageSamplerDescriptorCount`
- [](#VUID-vkGetDescriptorEXT-pDescriptorInfo-09507)VUID-vkGetDescriptorEXT-pDescriptorInfo-09507  
  If `pDescriptorInfo->type` is `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER` and it has a `imageView` that is [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) then `dataSize` **must** be equal to the size of [VkPhysicalDeviceDescriptorBufferPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDescriptorBufferPropertiesEXT.html)::`combinedImageSamplerDescriptorSize`

Valid Usage (Implicit)

- [](#VUID-vkGetDescriptorEXT-device-parameter)VUID-vkGetDescriptorEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetDescriptorEXT-pDescriptorInfo-parameter)VUID-vkGetDescriptorEXT-pDescriptorInfo-parameter  
  `pDescriptorInfo` **must** be a valid pointer to a valid [VkDescriptorGetInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorGetInfoEXT.html) structure
- [](#VUID-vkGetDescriptorEXT-pDescriptor-parameter)VUID-vkGetDescriptorEXT-pDescriptor-parameter  
  `pDescriptor` **must** be a valid pointer to an array of `dataSize` bytes
- [](#VUID-vkGetDescriptorEXT-dataSize-arraylength)VUID-vkGetDescriptorEXT-dataSize-arraylength  
  `dataSize` **must** be greater than `0`

## [](#_see_also)See Also

[VK\_EXT\_descriptor\_buffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_descriptor_buffer.html), [VkDescriptorGetInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorGetInfoEXT.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetDescriptorEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0