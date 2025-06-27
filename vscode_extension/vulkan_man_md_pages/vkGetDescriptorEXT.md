# vkGetDescriptorEXT(3) Manual Page

## Name

vkGetDescriptorEXT - To get a descriptor to place in a buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

To get descriptor data to place in a buffer, call:

``` c
// Provided by VK_EXT_descriptor_buffer
void vkGetDescriptorEXT(
    VkDevice                                    device,
    const VkDescriptorGetInfoEXT*               pDescriptorInfo,
    size_t                                      dataSize,
    void*                                       pDescriptor);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that gets the descriptor.

- `pDescriptorInfo` is a pointer to a
  [VkDescriptorGetInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorGetInfoEXT.html) structure
  specifying the parameters of the descriptor to get.

- `dataSize` is the amount of the descriptor data to get in bytes.

- `pDescriptor` is a pointer to an application-allocated buffer where
  the descriptor will be written.

## <a href="#_description" class="anchor"></a>Description

The size of the data for each descriptor type is determined by the value
in
[VkPhysicalDeviceDescriptorBufferPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDescriptorBufferPropertiesEXT.html).
This value also defines the stride in bytes for arrays of that
descriptor type.

If the
[VkPhysicalDeviceDescriptorBufferPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDescriptorBufferPropertiesEXT.html)::`combinedImageSamplerDescriptorSingleArray`
property is `VK_FALSE` the implementation requires an array of
`VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER` descriptors to be written
into a descriptor buffer as an array of image descriptors, immediately
followed by an array of sampler descriptors. Applications **must** write
the first
[VkPhysicalDeviceDescriptorBufferPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDescriptorBufferPropertiesEXT.html)::`sampledImageDescriptorSize`
bytes of the data returned through `pDescriptor` to the first array, and
the remaining
[VkPhysicalDeviceDescriptorBufferPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDescriptorBufferPropertiesEXT.html)::`samplerDescriptorSize`
bytes of the data to the second array. For variable-sized descriptor
bindings of `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER` descriptors, the
two arrays each have a size equal to the upper bound `descriptorCount`
of that binding.

A descriptor obtained by this command references the underlying
[VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) or [VkSampler](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampler.html), and
these objects **must** not be destroyed before the last time a
descriptor is dynamically accessed. For descriptor types which consume
an address instead of an object, the underlying
[VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) is referenced instead.

Valid Usage

- <a href="#VUID-vkGetDescriptorEXT-None-08015"
  id="VUID-vkGetDescriptorEXT-None-08015"></a>
  VUID-vkGetDescriptorEXT-None-08015  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-descriptorBuffer"
  target="_blank" rel="noopener"><code>descriptorBuffer</code></a>
  feature **must** be enabled

- <a href="#VUID-vkGetDescriptorEXT-dataSize-08125"
  id="VUID-vkGetDescriptorEXT-dataSize-08125"></a>
  VUID-vkGetDescriptorEXT-dataSize-08125  
  If `pDescriptorInfo->type` is not
  `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER` or
  `pDescriptorInfo->data.pCombinedImageSampler` has an `imageView`
  member that was not created with a `VkSamplerYcbcrConversionInfo`
  structure in its `pNext` chain, `dataSize` **must** equal the size of
  a descriptor of type
  [VkDescriptorGetInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorGetInfoEXT.html)::`type`
  determined by the value in
  [VkPhysicalDeviceDescriptorBufferPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDescriptorBufferPropertiesEXT.html)
  , or determined by
  [VkPhysicalDeviceDescriptorBufferDensityMapPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDescriptorBufferDensityMapPropertiesEXT.html)::`combinedImageSamplerDensityMapDescriptorSize`
  if `pDescriptorInfo` specifies a
  `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER` whose
  [VkSampler](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampler.html) was created with
  `VK_SAMPLER_CREATE_SUBSAMPLED_BIT_EXT` set

- <a href="#VUID-vkGetDescriptorEXT-descriptorType-09469"
  id="VUID-vkGetDescriptorEXT-descriptorType-09469"></a>
  VUID-vkGetDescriptorEXT-descriptorType-09469  
  If `pDescriptorInfo->type` is
  `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER` and
  `pDescriptorInfo->data.pCombinedImageSampler` has an `imageView`
  member that was created with a `VkSamplerYcbcrConversionInfo`
  structure in its `pNext` chain, `dataSize` **must** equal the size of
  [VkPhysicalDeviceDescriptorBufferPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDescriptorBufferPropertiesEXT.html)::`combinedImageSamplerDescriptorSize`
  times
  [VkSamplerYcbcrConversionImageFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversionImageFormatProperties.html)::`combinedImageSamplerDescriptorCount`

- <a href="#VUID-vkGetDescriptorEXT-pDescriptorInfo-09507"
  id="VUID-vkGetDescriptorEXT-pDescriptorInfo-09507"></a>
  VUID-vkGetDescriptorEXT-pDescriptorInfo-09507  
  If `pDescriptorInfo->type` is
  `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER` and it has a `imageView`
  that is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) then `dataSize` **must**
  be equal to the size of
  [VkPhysicalDeviceDescriptorBufferPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDescriptorBufferPropertiesEXT.html)::`combinedImageSamplerDescriptorSize`

Valid Usage (Implicit)

- <a href="#VUID-vkGetDescriptorEXT-device-parameter"
  id="VUID-vkGetDescriptorEXT-device-parameter"></a>
  VUID-vkGetDescriptorEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkGetDescriptorEXT-pDescriptorInfo-parameter"
  id="VUID-vkGetDescriptorEXT-pDescriptorInfo-parameter"></a>
  VUID-vkGetDescriptorEXT-pDescriptorInfo-parameter  
  `pDescriptorInfo` **must** be a valid pointer to a valid
  [VkDescriptorGetInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorGetInfoEXT.html) structure

- <a href="#VUID-vkGetDescriptorEXT-pDescriptor-parameter"
  id="VUID-vkGetDescriptorEXT-pDescriptor-parameter"></a>
  VUID-vkGetDescriptorEXT-pDescriptor-parameter  
  `pDescriptor` **must** be a valid pointer to an array of `dataSize`
  bytes

- <a href="#VUID-vkGetDescriptorEXT-dataSize-arraylength"
  id="VUID-vkGetDescriptorEXT-dataSize-arraylength"></a>
  VUID-vkGetDescriptorEXT-dataSize-arraylength  
  `dataSize` **must** be greater than `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_descriptor_buffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_descriptor_buffer.html),
[VkDescriptorGetInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorGetInfoEXT.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetDescriptorEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
